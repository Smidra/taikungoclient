/*
Taikun - WebApi

Testing CatalogAppAPIService

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

func Test_taikunshowback_CatalogAppAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CatalogAppAPIService CatalogAppCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CatalogAppAPI.CatalogAppCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.CatalogAppAPI.CatalogAppDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogAppId int32

		resp, httpRes, err := apiClient.CatalogAppAPI.CatalogAppDetails(context.Background(), catalogAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppEditParams", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CatalogAppAPI.CatalogAppEditParams(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppEditVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CatalogAppAPI.CatalogAppEditVersion(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppLockManager", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CatalogAppAPI.CatalogAppLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppAPIService CatalogAppParamDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogAppId int32

		resp, httpRes, err := apiClient.CatalogAppAPI.CatalogAppParamDetails(context.Background(), catalogAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
