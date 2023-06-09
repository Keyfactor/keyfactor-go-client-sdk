/*
Keyfactor-v1

Testing ReportsApiService

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

func Test_keyfactor_ReportsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReportsApiService ReportsCreateCustomReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsApi.ReportsCreateCustomReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsCreateReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsCreateReportSchedule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsDeleteReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.ReportsApi.ReportsDeleteReport(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsDeleteReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.ReportsApi.ReportsDeleteReportSchedule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsGetCustomReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsGetCustomReport(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsGetReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsGetReport(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsGetReportParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsGetReportParameters(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsGetReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsGetReportSchedule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsGetReportSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsGetReportSchedules(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsQueryCustomReports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsApi.ReportsQueryCustomReports(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsQueryReports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsApi.ReportsQueryReports(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsUpdateCustomReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsApi.ReportsUpdateCustomReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsUpdateReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReportsApi.ReportsUpdateReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsUpdateReportParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsUpdateReportParameters(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportsApiService ReportsUpdateReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ReportsApi.ReportsUpdateReportSchedule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
