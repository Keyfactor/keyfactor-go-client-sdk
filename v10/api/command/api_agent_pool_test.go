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

Testing AgentPoolApiService

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

func Test_command_AgentPoolApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test AgentPoolApiService AgentPoolCreateAgentPool", func(t *testing.T) {

		t.Log("AgentPoolApi_AgentPoolCreateAgentPool_payload: <none>")
		resp, httpRes, err := apiClient.AgentPoolApi.AgentPoolCreateAgentPool(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AgentPoolApiService AgentPoolDeleteAgentPool", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("AgentPoolApi_AgentPoolDeleteAgentPool_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("AgentPoolApi_AgentPoolDeleteAgentPool_id: %v", id)

		t.Log("AgentPoolApi_AgentPoolDeleteAgentPool_payload: <none>")
		httpRes, err := apiClient.AgentPoolApi.AgentPoolDeleteAgentPool(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AgentPoolApiService AgentPoolGetAgentPoolById", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("AgentPoolApi_AgentPoolGetAgentPoolById_id")
		id, _ = convertParamInterface(id, "string")
		t.Logf("AgentPoolApi_AgentPoolGetAgentPoolById_id: %v", id)

		t.Log("AgentPoolApi_AgentPoolGetAgentPoolById_payload: <none>")
		resp, httpRes, err := apiClient.AgentPoolApi.AgentPoolGetAgentPoolById(context.Background(), id.(string)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AgentPoolApiService AgentPoolGetAgentPools", func(t *testing.T) {

		t.Log("AgentPoolApi_AgentPoolGetAgentPools_payload: <none>")
		resp, httpRes, err := apiClient.AgentPoolApi.AgentPoolGetAgentPools(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AgentPoolApiService AgentPoolGetDefaultAgentPoolAgents", func(t *testing.T) {

		t.Log("AgentPoolApi_AgentPoolGetDefaultAgentPoolAgents_payload: <none>")
		resp, httpRes, err := apiClient.AgentPoolApi.AgentPoolGetDefaultAgentPoolAgents(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AgentPoolApiService AgentPoolUpdateAgentPool", func(t *testing.T) {

		t.Log("AgentPoolApi_AgentPoolUpdateAgentPool_payload: <none>")
		resp, httpRes, err := apiClient.AgentPoolApi.AgentPoolUpdateAgentPool(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}