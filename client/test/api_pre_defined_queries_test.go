/*
Taikun - WebApi

Testing PreDefinedQueriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	openapiclient "github.com/Smidra/taikungoclient/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikuncore_PreDefinedQueriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PreDefinedQueriesAPIService PredefinedqueriesCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreDefinedQueriesAPIService PredefinedqueriesDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreDefinedQueriesAPIService PredefinedqueriesList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreDefinedQueriesAPIService PredefinedqueriesPrometheusDashboardCommon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesPrometheusDashboardCommon(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreDefinedQueriesAPIService PredefinedqueriesUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
