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

Testing UserApiService

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

func Test_command_UserApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test UserApiService SSHUsersAccessPost", func(t *testing.T) {

		t.Log("UserApi_SSHUsersAccessPost_payload: <none>")
		resp, httpRes, err := apiClient.UserApi.SSHUsersAccessPost(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UserApiService SSHUsersGet", func(t *testing.T) {

		t.Log("UserApi_SSHUsersGet_payload: <none>")
		resp, httpRes, err := apiClient.UserApi.SSHUsersGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UserApiService SSHUsersIdDelete", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("UserApi_SSHUsersIdDelete_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("UserApi_SSHUsersIdDelete_id: %v", id)

		t.Log("UserApi_SSHUsersIdDelete_payload: <none>")
		httpRes, err := apiClient.UserApi.SSHUsersIdDelete(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UserApiService SSHUsersIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("UserApi_SSHUsersIdGet_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("UserApi_SSHUsersIdGet_id: %v", id)

		t.Log("UserApi_SSHUsersIdGet_payload: <none>")
		resp, httpRes, err := apiClient.UserApi.SSHUsersIdGet(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UserApiService SSHUsersPost", func(t *testing.T) {

		t.Log("UserApi_SSHUsersPost_payload: <none>")
		resp, httpRes, err := apiClient.UserApi.SSHUsersPost(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test UserApiService SSHUsersPut", func(t *testing.T) {

		t.Log("UserApi_SSHUsersPut_payload: <none>")
		resp, httpRes, err := apiClient.UserApi.SSHUsersPut(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}