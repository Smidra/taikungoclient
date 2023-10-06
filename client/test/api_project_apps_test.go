/*
Taikun - WebApi

Testing ProjectAppsAPIService

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

func Test_taikuncore_ProjectAppsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAppsAPIService ProjectappAutosync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectAppsAPI.ProjectappAutosync(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectAppId int32

		httpRes, err := apiClient.ProjectAppsAPI.ProjectappDelete(context.Background(), projectAppId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ProjectAppsAPI.ProjectappDetails(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappInstall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectAppsAPI.ProjectappInstall(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectAppsAPI.ProjectappList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappLockManager", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectAppsAPI.ProjectappLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAppsAPIService ProjectappSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectAppsAPI.ProjectappSync(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
