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

Testing SecurityRolesApiService

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

func Test_command_SecurityRolesApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test SecurityRolesApiService SecurityRolesGet", func(t *testing.T) {

		t.Log("SecurityRolesApi_SecurityRolesGet_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesIdCopyPost", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesIdCopyPost_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesIdCopyPost_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesIdCopyPost_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesIdCopyPost(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesIdDelete", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesIdDelete_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesIdDelete_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesIdDelete_payload: <none>")
		httpRes, err := apiClient.SecurityRolesApi.SecurityRolesIdDelete(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesIdGet_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesIdGet_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesIdGet_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesIdGet(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesIdIdentitiesGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesIdIdentitiesGet_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesIdIdentitiesGet_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesIdIdentitiesGet_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesIdIdentitiesGet(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesIdIdentitiesPut", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesIdIdentitiesPut_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesIdIdentitiesPut_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesIdIdentitiesPut_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesIdIdentitiesPut(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesPost", func(t *testing.T) {

		t.Log("SecurityRolesApi_SecurityRolesPost_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesPost(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesPut", func(t *testing.T) {

		t.Log("SecurityRolesApi_SecurityRolesPut_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesPut(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}