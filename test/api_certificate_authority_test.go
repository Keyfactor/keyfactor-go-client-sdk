/*
Keyfactor-v1

Testing CertificateAuthorityApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keyfactor

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Keyfactor/keyfactor-go-client-sdk"
)

func Test_keyfactor_CertificateAuthorityApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityCreateCA", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityCreateCA(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityDeleteCA", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityDeleteCA(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityGetCa", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGetCa(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityGetCas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGetCas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityPublishCRL", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityPublishCRL(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityTestCertificateAuthority", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityTestCertificateAuthority(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificateAuthorityApiService CertificateAuthorityUpdateCA", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificateAuthorityApi.CertificateAuthorityUpdateCA(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
