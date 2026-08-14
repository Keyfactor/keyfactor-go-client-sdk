package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
)

// unsetEnvForTest removes an environment variable for the duration of the
// test and restores its original value (or absence) on cleanup. Unlike
// t.Setenv(key, ""), which leaves the variable "present" with an empty
// value (still visible to os.LookupEnv), this genuinely unsets it so code
// that branches on presence -- e.g. ValidateAuthConfig's
// KEYFACTOR_CLIENT_TIMEOUT handling, which treats a present-but-unparseable
// value as "leave HttpClientTimeout at its current value" rather than
// falling through to the 60s default -- behaves as if the caller's shell
// never exported it at all.
func unsetEnvForTest(t *testing.T, key string) {
	t.Helper()
	if orig, ok := os.LookupEnv(key); ok {
		t.Cleanup(func() {
			_ = os.Setenv(key, orig)
		})
		_ = os.Unsetenv(key)
	}
}

// TestCommandConfigOauth_AccessTokenFieldPropagation is a compilation + correctness
// regression test for the v2.8.0 bug where AccessToken, Audience, and Scopes were
// silently dropped when constructing CommandConfigOauth from auth_providers.Server
// in buildHttpClientV2. If any of those three fields are ever removed from either
// struct, this test fails to compile.
func TestCommandConfigOauth_AccessTokenFieldPropagation(t *testing.T) {
	srv := &auth_providers.Server{
		Host:        "test.example.com",
		AccessToken: "mytoken-abc123",
		Audience:    "https://my.audience.example.com",
		Scopes:      []string{"read", "write", "admin"},
	}

	// Step 1: Verify Server struct holds the fields correctly
	if srv.AccessToken != "mytoken-abc123" {
		t.Errorf("Server.AccessToken = %q, want %q", srv.AccessToken, "mytoken-abc123")
	}
	if srv.Audience != "https://my.audience.example.com" {
		t.Errorf("Server.Audience = %q, want %q", srv.Audience, "https://my.audience.example.com")
	}
	if !reflect.DeepEqual(srv.Scopes, []string{"read", "write", "admin"}) {
		t.Errorf("Server.Scopes = %v, want %v", srv.Scopes, []string{"read", "write", "admin"})
	}

	// Step 2: Construct CommandConfigOauth the same way buildHttpClientV2 does
	// (minus the Authenticate() call which requires network). This mirrors lines
	// 344-351 of client.go exactly.
	baseConfig := auth_providers.CommandAuthConfig{
		CommandHostName: srv.Host,
		CommandPort:     srv.Port,
		CommandAPIPath:  srv.APIPath,
		CommandCACert:   srv.CACertPath,
		SkipVerify:      srv.SkipTLSVerify,
	}
	oauthCfg := auth_providers.CommandConfigOauth{
		CommandAuthConfig: baseConfig,
		ClientID:          srv.ClientID,
		ClientSecret:      srv.ClientSecret,
		TokenURL:          srv.OAuthTokenUrl,
		AccessToken:       srv.AccessToken,
		Audience:          srv.Audience,
		Scopes:            srv.Scopes,
	}

	// Step 3: Verify the three fields that were missing in the v2.8.0 regression
	if oauthCfg.AccessToken != "mytoken-abc123" {
		t.Errorf("CommandConfigOauth.AccessToken = %q, want %q", oauthCfg.AccessToken, "mytoken-abc123")
	}
	if oauthCfg.Audience != "https://my.audience.example.com" {
		t.Errorf("CommandConfigOauth.Audience = %q, want %q", oauthCfg.Audience, "https://my.audience.example.com")
	}
	if !reflect.DeepEqual(oauthCfg.Scopes, []string{"read", "write", "admin"}) {
		t.Errorf("CommandConfigOauth.Scopes = %v, want %v", oauthCfg.Scopes, []string{"read", "write", "admin"})
	}

	// Step 4: Verify GetAuthType returns "oauth" for access_token-only config
	authType := srv.GetAuthType()
	if authType != "oauth" {
		t.Errorf("Server.GetAuthType() = %q, want %q (access_token-only should be oauth)", authType, "oauth")
	}
}

// TestCommandConfigOauth_AccessTokenOnlyNoClientCreds verifies that a Server
// configured with only Host + AccessToken (no ClientID/ClientSecret/TokenURL)
// is classified as "oauth" auth type and the token propagates correctly.
func TestCommandConfigOauth_AccessTokenOnlyNoClientCreds(t *testing.T) {
	srv := &auth_providers.Server{
		Host:        "command.example.com",
		AccessToken: "pre-fetched-bearer-token",
		// Deliberately omitting ClientID, ClientSecret, OAuthTokenUrl
	}

	if got := srv.GetAuthType(); got != "oauth" {
		t.Fatalf("GetAuthType() = %q, want %q for access_token-only", got, "oauth")
	}

	oauthCfg := auth_providers.CommandConfigOauth{
		AccessToken: srv.AccessToken,
	}

	if oauthCfg.AccessToken != "pre-fetched-bearer-token" {
		t.Errorf("AccessToken = %q, want %q", oauthCfg.AccessToken, "pre-fetched-bearer-token")
	}
	if oauthCfg.ClientID != "" {
		t.Errorf("ClientID = %q, want empty", oauthCfg.ClientID)
	}
	if oauthCfg.ClientSecret != "" {
		t.Errorf("ClientSecret = %q, want empty", oauthCfg.ClientSecret)
	}
}

// newFakeCommandServer stands in for a Keyfactor Command instance for
// CommandAuthConfigBasic.Authenticate(), which performs a real GET against
// {host}/{apiPath}/Status/Endpoints as part of authentication. It always
// returns 200 with a valid JSON string array, regardless of credentials.
func newFakeCommandServer(t *testing.T) *httptest.Server {
	t.Helper()
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["endpoint1"]`))
	}))
	t.Cleanup(server.Close)
	return server
}

// TestBuildHttpClientV2_ClientTimeoutPropagation is a regression test for the
// bug where Server.ClientTimeout was silently dropped when buildHttpClientV2
// rebuilt its own CommandAuthConfig, causing every caller (including the
// Terraform provider's request_timeout setting) to fall back to
// auth_providers.DefaultClientTimeout (60s) regardless of what was
// configured -- surfacing as "net/http: timeout awaiting response headers" on
// long-running calls such as PFX enrollment. Unlike
// TestCommandConfigOauth_AccessTokenFieldPropagation above (which mirrors
// buildHttpClientV2's lines to avoid the network call inside Authenticate()),
// this test exercises buildHttpClientV2 itself against a fake Command server.
func TestBuildHttpClientV2_ClientTimeoutPropagation(t *testing.T) {
	// Hermetic against the ambient environment: ValidateAuthConfig (in
	// keyfactor-auth-client-go's auth_core.go) unconditionally overwrites
	// SkipVerify from KEYFACTOR_SKIP_VERIFY whenever the variable is merely
	// *present*, regardless of its value -- so anything other than exactly
	// "true"/"1" (e.g. "TRUE", "false", "0", or an empty string sourced from
	// a lab env file) flips SkipTLSVerify back to false and this test fails
	// against its own self-signed httptest server with "x509: certificate
	// signed by unknown authority". Pin it explicitly rather than relying on
	// it being unset in whatever shell runs `go test`. (The upstream
	// unconditional-overwrite behavior itself is tracked and fixed
	// separately in keyfactor-auth-client-go; this only needs to make our
	// test hermetic against it.)
	t.Setenv(auth_providers.EnvKeyfactorSkipVerify, "true")
	// A stale/bad KEYFACTOR_CA_CERT path in the ambient environment would
	// make BuildTransport() below treat the value as literal PEM bytes and
	// fail with "failed to append custom CA cert to pool". Neutralize it.
	unsetEnvForTest(t, auth_providers.EnvKeyfactorCACert)
	// KEYFACTOR_CLIENT_TIMEOUT only matters when HttpClientTimeout is <= 0
	// going in (it isn't here -- srv.ClientTimeout is 300 below), but pin it
	// too for defense-in-depth. t.Setenv(..., "") would NOT achieve this: an
	// empty value is still "present" to os.LookupEnv, so ValidateAuthConfig
	// would see ok=true, fail to strconv.Atoi(""), and leave
	// HttpClientTimeout at whatever it currently is instead of falling
	// through to the 60s default -- the "zero-timeout case".
	unsetEnvForTest(t, auth_providers.EnvKeyfactorClientTimeout)

	server := newFakeCommandServer(t)
	u, uErr := url.Parse(server.URL)
	if uErr != nil {
		t.Fatalf("failed to parse test server URL: %v", uErr)
	}

	srv := &auth_providers.Server{
		Host:          u.Host,
		Username:      "user",
		Password:      "pass",
		APIPath:       "api",
		SkipTLSVerify: true,
		ClientTimeout: 300,
	}

	authCfg, err := buildHttpClientV2(srv)
	if err != nil {
		t.Fatalf("buildHttpClientV2() returned unexpected error: %v", err)
	}

	basicCfg, ok := authCfg.(*auth_providers.CommandAuthConfigBasic)
	if !ok {
		t.Fatalf("expected AuthConfig to be *auth_providers.CommandAuthConfigBasic, got %T", authCfg)
	}

	if basicCfg.HttpClientTimeout != 300 {
		t.Errorf("CommandAuthConfigBasic.HttpClientTimeout = %d, want %d", basicCfg.HttpClientTimeout, 300)
	}

	transport, tErr := basicCfg.CommandAuthConfig.BuildTransport()
	if tErr != nil {
		t.Fatalf("BuildTransport() returned unexpected error: %v", tErr)
	}

	expected := 300 * time.Second
	if transport.ResponseHeaderTimeout != expected {
		t.Errorf("ResponseHeaderTimeout = %v, want %v", transport.ResponseHeaderTimeout, expected)
	}
}

// newFakeCommandServerCapturingAuth is like newFakeCommandServer but also
// records the Authorization header of the last request it received, so a
// test can assert which credential actually reached the wire (rather than
// only asserting the shape of a struct literal that never leaves the test).
func newFakeCommandServerCapturingAuth(t *testing.T) (*httptest.Server, *string) {
	t.Helper()
	var gotAuth string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["endpoint1"]`))
	}))
	t.Cleanup(server.Close)
	return server, &gotAuth
}

// TestBuildHttpClientV2_OAuthAccessTokenPropagation is a regression test for
// the hand-edit cataloged in HAND_EDITS.md (v24 entry #2) that restores
// AccessToken, Audience, and Scopes to the auth_providers.CommandConfigOauth{}
// literal inside buildHttpClientV2's OAuth branch. Those three fields were
// silently dropped once already (2b88eb2) and restored (229db7d); this test
// -- unlike TestCommandConfigOauth_AccessTokenFieldPropagation above, which
// only constructs its own local CommandConfigOauth literal and would keep
// passing even if buildHttpClientV2's actual literal dropped these fields --
// drives buildHttpClientV2 itself end-to-end against a fake Command server
// and asserts the pre-fetched access token actually reaches the wire as a
// Bearer credential.
func TestBuildHttpClientV2_OAuthAccessTokenPropagation(t *testing.T) {
	// See TestBuildHttpClientV2_ClientTimeoutPropagation for why each of
	// these is pinned/unset: makes the test hermetic against the ambient
	// shell environment.
	t.Setenv(auth_providers.EnvKeyfactorSkipVerify, "true")
	unsetEnvForTest(t, auth_providers.EnvKeyfactorCACert)
	unsetEnvForTest(t, auth_providers.EnvKeyfactorClientTimeout)
	unsetEnvForTest(t, auth_providers.EnvKeyfactorAccessToken)
	unsetEnvForTest(t, auth_providers.EnvKeyfactorClientID)

	server, gotAuth := newFakeCommandServerCapturingAuth(t)
	u, uErr := url.Parse(server.URL)
	if uErr != nil {
		t.Fatalf("failed to parse test server URL: %v", uErr)
	}

	srv := &auth_providers.Server{
		Host: u.Host,
		// Deliberately no ClientID/ClientSecret/OAuthTokenUrl: this is the
		// pre-fetched access_token-only auth path that 2b88eb2 broke.
		AccessToken:   "mytoken-abc123",
		Audience:      "https://my.audience.example.com",
		Scopes:        []string{"read", "write"},
		APIPath:       "api",
		SkipTLSVerify: true,
		ClientTimeout: 60,
	}

	authCfg, err := buildHttpClientV2(srv)
	if err != nil {
		t.Fatalf("buildHttpClientV2() returned unexpected error: %v", err)
	}

	oauthCfg, ok := authCfg.(*auth_providers.CommandConfigOauth)
	if !ok {
		t.Fatalf("expected AuthConfig to be *auth_providers.CommandConfigOauth, got %T", authCfg)
	}

	if oauthCfg.AccessToken != "mytoken-abc123" {
		t.Errorf("CommandConfigOauth.AccessToken = %q, want %q", oauthCfg.AccessToken, "mytoken-abc123")
	}
	if oauthCfg.Audience != "https://my.audience.example.com" {
		t.Errorf("CommandConfigOauth.Audience = %q, want %q", oauthCfg.Audience, "https://my.audience.example.com")
	}
	if !reflect.DeepEqual(oauthCfg.Scopes, []string{"read", "write"}) {
		t.Errorf("CommandConfigOauth.Scopes = %v, want %v", oauthCfg.Scopes, []string{"read", "write"})
	}

	// The real assertion: buildHttpClientV2's internal Authenticate() call
	// made an actual HTTP request to the fake Command server, and that
	// request must carry the access token as a Bearer credential. If
	// AccessToken were dropped from buildHttpClientV2's literal, the OAuth
	// branch would fall through to the client-credentials grant with no
	// ClientID/ClientSecret/TokenURL and buildHttpClientV2 would return an
	// error above instead of ever reaching this assertion.
	wantAuth := "Bearer mytoken-abc123"
	if *gotAuth != wantAuth {
		t.Errorf("request Authorization header = %q, want %q", *gotAuth, wantAuth)
	}
}

// TestPrepareRequest_Port443Guard is a regression test for the hand-edit in
// commit 229db7d that added "&& serverConfig.Port != 443" to prepareRequest's
// port guard. Without it, a Server configured with Port: 443 (the default
// HTTPS port, and what many callers -- including the Terraform provider --
// set explicitly) produces request URLs like "https://host:443/..." instead
// of "https://host/...". Both are technically valid HTTPS URLs, but the
// explicit ":443" broke servers/proxies that match on Host header exactly
// (no port suffix) and was reported as a functional regression. This test
// was previously unprotected: reverting the guard would fail nothing in CI.
func TestPrepareRequest_Port443Guard(t *testing.T) {
	// Constructed directly against the exported AuthClient field (rather
	// than via NewAPIClientWithAuth) so this test exercises prepareRequest
	// in isolation without depending on an unrelated hand-edit.
	newClientWithPort := func(t *testing.T, port int) *APIClient {
		t.Helper()
		return &APIClient{
			AuthClient: &auth_providers.CommandAuthConfigBasic{
				CommandAuthConfig: auth_providers.CommandAuthConfig{
					CommandHostName: "command.example.com",
					CommandPort:     port,
				},
			},
		}
	}

	prepare := func(t *testing.T, c *APIClient) *http.Request {
		t.Helper()
		req, err := c.prepareRequest(
			context.Background(),
			"https://placeholder.invalid/api/Status/Endpoints",
			"GET",
			nil,
			map[string]string{},
			nil,
			nil,
			nil,
		)
		if err != nil {
			t.Fatalf("prepareRequest() returned unexpected error: %v", err)
		}
		return req
	}

	t.Run("port 443 is omitted from the request host", func(t *testing.T) {
		req := prepare(t, newClientWithPort(t, 443))
		if got, want := req.URL.Host, "command.example.com"; got != want {
			t.Errorf("URL.Host = %q, want %q", got, want)
		}
	})

	t.Run("non-443 port is still appended to the request host", func(t *testing.T) {
		req := prepare(t, newClientWithPort(t, 8443))
		if got, want := req.URL.Host, "command.example.com:8443"; got != want {
			t.Errorf("URL.Host = %q, want %q", got, want)
		}
	})
}
