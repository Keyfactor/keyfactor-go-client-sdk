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

Testing AuditLogApiService

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

func Test_command_AuditLogApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test AuditLogApiService AuditDownloadGet", func(t *testing.T) {

		t.Log("AuditLogApi_AuditDownloadGet_payload: <none>")
		resp, httpRes, err := apiClient.AuditLogApi.AuditDownloadGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AuditLogApiService AuditGet", func(t *testing.T) {

		t.Log("AuditLogApi_AuditGet_payload: <none>")
		resp, httpRes, err := apiClient.AuditLogApi.AuditGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AuditLogApiService AuditIdGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("AuditLogApi_AuditIdGet_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("AuditLogApi_AuditIdGet_id: %v", id)

		t.Log("AuditLogApi_AuditIdGet_payload: <none>")
		httpRes, err := apiClient.AuditLogApi.AuditIdGet(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AuditLogApiService AuditIdValidateGet", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("AuditLogApi_AuditIdValidateGet_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("AuditLogApi_AuditIdValidateGet_id: %v", id)

		t.Log("AuditLogApi_AuditIdValidateGet_payload: <none>")
		resp, httpRes, err := apiClient.AuditLogApi.AuditIdValidateGet(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AuditLogApiService AuditRelatedEntitiesGet", func(t *testing.T) {

		t.Log("AuditLogApi_AuditRelatedEntitiesGet_payload: <none>")
		resp, httpRes, err := apiClient.AuditLogApi.AuditRelatedEntitiesGet(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}