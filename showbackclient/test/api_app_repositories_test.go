/*
Taikun - WebApi

Testing AppRepositoriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikunshowback

import (
	"context"
	openapiclient "github.com/Smidra/taikungoclient/showbackclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikunshowback_AppRepositoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppRepositoriesAPIService RepositoryAvailableList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppRepositoriesAPI.RepositoryAvailableList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryBind", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppRepositoriesAPI.RepositoryBind(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppRepositoriesAPI.RepositoryDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryDisable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppRepositoriesAPI.RepositoryDisable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryImport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppRepositoriesAPI.RepositoryImport(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryRecommendedList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppRepositoriesAPI.RepositoryRecommendedList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRepositoriesAPIService RepositoryUnbind", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppRepositoriesAPI.RepositoryUnbind(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}