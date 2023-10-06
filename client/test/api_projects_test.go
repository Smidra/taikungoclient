/*
Taikun - WebApi

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	openapiclient "github.com/smidra/taikungoclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikuncore_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService ProjectsAiAnalyzer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsAiAnalyzer(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsAlerts(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsChatCompletions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsChatCompletions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsCommit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsAPI.ProjectsCommit(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDeleteWholeProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsDeleteWholeProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDescribe", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsDescribe(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsDetails(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDropdown", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsAPI.ProjectsEdit(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsEditHealth", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsEditHealth(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsEditStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsEditStatus(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsExtendLifetime", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsExtendLifetime(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsForAlerting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsForAlerting(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsForBilling", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsForBilling(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsForPoller", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsForPoller(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsLockManager", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsLokiLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsLokiLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsMonitoring", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsMonitoring(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsMonitoringAlerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsMonitoringAlerts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsPrometheusMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsPrometheusMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsPurge", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsPurge(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsPurgeWholeProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsPurgeWholeProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsRepair", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsAPI.ProjectsRepair(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsToggleFullSpot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsToggleFullSpot(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsToggleSpotVms", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsToggleSpotVms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsToggleSpotWorkers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectsAPI.ProjectsToggleSpotWorkers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsUpgrade", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectsAPI.ProjectsUpgrade(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsVisibility", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsVisibility(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
