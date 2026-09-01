package keyfactor

import (
	"net/http"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
	v1 "github.com/Keyfactor/keyfactor-go-client-sdk/v25/api/keyfactor/v1"
	v2 "github.com/Keyfactor/keyfactor-go-client-sdk/v25/api/keyfactor/v2"
)

type APIClient struct {
	V1 *v1.APIClient
	V2 *v2.APIClient
}

// AuthConfig is a common interface implemented by VCR auth stubs and real auth configs.
// Structurally equivalent to v1.AuthConfig and v2.AuthConfig.
type AuthConfig interface {
	Authenticate() error
	GetHttpClient() (*http.Client, error)
	GetServerConfig() *auth_providers.Server
}

// NewAPIClientWithAuth creates an APIClient with a pre-built AuthConfig, bypassing
// the Authenticate() network call. Used in unit tests with VCR cassettes.
func NewAPIClientWithAuth(auth AuthConfig) *APIClient {
	var v1a v1.AuthConfig = auth
	var v2a v2.AuthConfig = auth
	return &APIClient{
		V1: v1.NewAPIClientWithAuth(v1a),
		V2: v2.NewAPIClientWithAuth(v2a),
	}
}

func NewAPIClient(cfg *auth_providers.Server) (*APIClient, error) {
	var err error

	clientV1, err := v1.NewAPIClient(cfg)
	if err != nil {
		return nil, err
	}
	clientV2, err := v2.NewAPIClient(cfg)
	if err != nil {
		return nil, err
	}

	return &APIClient{
		V1: clientV1,
		V2: clientV2,
	}, nil
}
