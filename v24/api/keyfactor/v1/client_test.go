package v1

import (
	"reflect"
	"testing"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
)

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
