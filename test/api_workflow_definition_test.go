/*
Keyfactor-v1

Testing WorkflowDefinitionApiService

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

func Test_keyfactor_WorkflowDefinitionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionConfigureDefinitionSteps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionConfigureDefinitionSteps(context.Background(), definitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionCreateNewDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionCreateNewDefinition(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string

		httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionDelete(context.Background(), definitionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionGet(context.Background(), definitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionGetStepSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionGetStepSchema(context.Background(), extensionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionPublishDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionPublishDefinition(context.Background(), definitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQuery(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionQueryAvailableSteps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQueryAvailableSteps(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionQueryWorkflowTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQueryWorkflowTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkflowDefinitionApiService WorkflowDefinitionUpdateExistingDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string

		resp, httpRes, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionUpdateExistingDefinition(context.Background(), definitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}