/*
Copyright 2023 Keyfactor
Licensed under the Apache License, Version 2.0 (the License); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an AS IS BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

Keyfactor API Reference and Utility

Testing SslApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package command

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_command_SslApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test SslApiService SSLEndpointsIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLEndpointsIdGet_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLEndpointsIdGet_id: %v", id)

		t.Log("SslApi_SSLEndpointsIdGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLEndpointsIdGet(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLEndpointsIdHistoryGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLEndpointsIdHistoryGet_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLEndpointsIdHistoryGet_id: %v", id)

		t.Log("SslApi_SSLEndpointsIdHistoryGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLEndpointsIdHistoryGet(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLEndpointsMonitorAllPut", func(t *testing.T) {

		t.Log("SslApi_SSLEndpointsMonitorAllPut_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLEndpointsMonitorAllPut(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLEndpointsMonitorStatusPut", func(t *testing.T) {

		t.Log("SslApi_SSLEndpointsMonitorStatusPut_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLEndpointsMonitorStatusPut(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLEndpointsReviewAllPut", func(t *testing.T) {

		t.Log("SslApi_SSLEndpointsReviewAllPut_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLEndpointsReviewAllPut(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLEndpointsReviewStatusPut", func(t *testing.T) {

		t.Log("SslApi_SSLEndpointsReviewStatusPut_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLEndpointsReviewStatusPut(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLGet", func(t *testing.T) {

		t.Log("SslApi_SSLGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworkRangesIdDelete", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworkRangesIdDelete_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworkRangesIdDelete_id: %v", id)

		t.Log("SslApi_SSLNetworkRangesIdDelete_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworkRangesIdDelete(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworkRangesIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworkRangesIdGet_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworkRangesIdGet_id: %v", id)

		t.Log("SslApi_SSLNetworkRangesIdGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworkRangesIdGet(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworkRangesPost", func(t *testing.T) {

		t.Log("SslApi_SSLNetworkRangesPost_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworkRangesPost(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworkRangesPut", func(t *testing.T) {

		t.Log("SslApi_SSLNetworkRangesPut_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworkRangesPut(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworkRangesValidatePost", func(t *testing.T) {

		t.Log("SslApi_SSLNetworkRangesValidatePost_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworkRangesValidatePost(context.Background()).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksGet", func(t *testing.T) {

		t.Log("SslApi_SSLNetworksGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworksGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksIdDelete", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworksIdDelete_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworksIdDelete_id: %v", id)

		t.Log("SslApi_SSLNetworksIdDelete_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworksIdDelete(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksIdPartsGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworksIdPartsGet_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworksIdPartsGet_id: %v", id)

		t.Log("SslApi_SSLNetworksIdPartsGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworksIdPartsGet(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksIdResetPost", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworksIdResetPost_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworksIdResetPost_id: %v", id)

		t.Log("SslApi_SSLNetworksIdResetPost_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworksIdResetPost(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksIdScanPost", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLNetworksIdScanPost_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLNetworksIdScanPost_id: %v", id)

		t.Log("SslApi_SSLNetworksIdScanPost_payload: <none>")
		httpRes, err := apiClient.SslApi.SSLNetworksIdScanPost(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksIdentifierGet", func(t *testing.T) {

		var identifier interface{}

		identifier = os.Getenv("SslApi_SSLNetworksIdentifierGet_identifier")
		identifier, _ = convertParamInterface(identifier, "string")
		t.Logf("SslApi_SSLNetworksIdentifierGet_identifier: %v", identifier)

		t.Log("SslApi_SSLNetworksIdentifierGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworksIdentifierGet(context.Background(), identifier.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksPost", func(t *testing.T) {

		t.Log("SslApi_SSLNetworksPost_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworksPost(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLNetworksPut", func(t *testing.T) {

		t.Log("SslApi_SSLNetworksPut_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLNetworksPut(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SslApiService SSLPartsIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SslApi_SSLPartsIdGet_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("SslApi_SSLPartsIdGet_id: %v", id)

		t.Log("SslApi_SSLPartsIdGet_payload: <none>")
		resp, httpRes, err := apiClient.SslApi.SSLPartsIdGet(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}