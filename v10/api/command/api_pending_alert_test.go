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

Testing PendingAlertApiService

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

func Test_command_PendingAlertApiService(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Logf("Working directory: %s", cwd)
	config := GetEnvConfiguration()

	configuration, configErr := NewConfiguration(config)
	require.Nil(t, configErr)

	apiClient := NewAPIClient(configuration)

	t.Run("Test PendingAlertApiService PendingAlertAddPendingAlert", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertAddPendingAlert_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertAddPendingAlert(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertDeletePendingAlert", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("PendingAlertApi_PendingAlertDeletePendingAlert_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("PendingAlertApi_PendingAlertDeletePendingAlert_id: %v", id)

		t.Log("PendingAlertApi_PendingAlertDeletePendingAlert_payload: <none>")
		httpRes, err := apiClient.PendingAlertApi.PendingAlertDeletePendingAlert(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertEditPendingAlert", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertEditPendingAlert_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertEditPendingAlert(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertEditSchedule", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertEditSchedule_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertEditSchedule(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertGetPendingAlert", func(t *testing.T) {

		var id interface{}

		id = os.Getenv("PendingAlertApi_PendingAlertGetPendingAlert_id")
		id, _ = convertParamInterface(id, "int32")
		t.Logf("PendingAlertApi_PendingAlertGetPendingAlert_id: %v", id)

		t.Log("PendingAlertApi_PendingAlertGetPendingAlert_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertGetPendingAlert(context.Background(), id.(int32)).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertGetPendingAlerts", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertGetPendingAlerts_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertGetPendingAlerts(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertGetSchedule", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertGetSchedule_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertGetSchedule(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertTestAllPendingAlert", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertTestAllPendingAlert_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertTestAllPendingAlert(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PendingAlertApiService PendingAlertTestPendingAlert", func(t *testing.T) {

		t.Log("PendingAlertApi_PendingAlertTestPendingAlert_payload: <none>")
		resp, httpRes, err := apiClient.PendingAlertApi.PendingAlertTestPendingAlert(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}