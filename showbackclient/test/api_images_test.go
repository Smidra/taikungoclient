/*
Taikun - WebApi

Testing ImagesAPIService

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

func Test_taikunshowback_ImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesAPIService ImagesAwsCommonImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAwsCommonImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesAwsImagesList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAwsImagesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesAwsPersonalImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAwsPersonalImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesAzureCommonImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAzureCommonImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesAzureImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32
		var publisherName string
		var offer string
		var sku string

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAzureImages(context.Background(), cloudId, publisherName, offer, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesAzurePersonalImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesAzurePersonalImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesBindImagesToProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ImagesAPI.ImagesBindImagesToProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesCommonGoogleImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesCommonGoogleImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesGoogleImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32
		var type_ string

		resp, httpRes, err := apiClient.ImagesAPI.ImagesGoogleImages(context.Background(), cloudId, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesImageDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImagesAPI.ImagesImageDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesOpenstackImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesOpenstackImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesProxmoxImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesProxmoxImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesSelectedImagesForProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImagesAPI.ImagesSelectedImagesForProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesTanzuImages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesAPI.ImagesTanzuImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImagesUnbindImagesFromProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ImagesAPI.ImagesUnbindImagesFromProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
