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

Keyfactor-v1

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

	t.Run("Test SecurityRolesApiService SecurityRolesDeleteSecurityRole", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesDeleteSecurityRole_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesDeleteSecurityRole_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesDeleteSecurityRole_payload: <none>")
		httpRes, err := apiClient.SecurityRolesApi.SecurityRolesDeleteSecurityRole(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesGetIdentitiesWithRole", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesGetIdentitiesWithRole_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesGetIdentitiesWithRole_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesGetIdentitiesWithRole_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesGetIdentitiesWithRole(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesGetSecurityRole", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesGetSecurityRole_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesGetSecurityRole_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesGetSecurityRole_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesGetSecurityRole(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SecurityRolesApiService SecurityRolesUpdateIdentitiesWithRole", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("SecurityRolesApi_SecurityRolesUpdateIdentitiesWithRole_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("SecurityRolesApi_SecurityRolesUpdateIdentitiesWithRole_id: %v", id)

		t.Log("SecurityRolesApi_SecurityRolesUpdateIdentitiesWithRole_payload: <none>")
		resp, httpRes, err := apiClient.SecurityRolesApi.SecurityRolesUpdateIdentitiesWithRole(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}