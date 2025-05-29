package keyfactor

import (
	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
	v1 "github.com/Keyfactor/keyfactor-go-client-sdk/v25/api/keyfactor/v1"
	v2 "github.com/Keyfactor/keyfactor-go-client-sdk/v25/api/keyfactor/v2"
)

type APIClient struct {
	V1 *v1.APIClient
	V2 *v2.APIClient
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
