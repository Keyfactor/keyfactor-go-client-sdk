/*
Keyfactor-v1

Testing OrchestratorJobApiService

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

func Test_keyfactor_OrchestratorJobApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrchestratorJobApiService OrchestratorJobAcknowledgeJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobAcknowledgeJobs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobGetCustomJobResultData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobGetCustomJobResultData(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobGetJobHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobGetJobHistory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobGetScheduledJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobGetScheduledJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobRescheduleJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobRescheduleJobs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobScheduleBulkJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobScheduleBulkJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobScheduleJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobScheduleJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrchestratorJobApiService OrchestratorJobUnscheduleJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrchestratorJobApi.OrchestratorJobUnscheduleJobs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}