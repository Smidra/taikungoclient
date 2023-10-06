/*
Taikun - WebApi

Testing FlavorsAPIService

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

func Test_taikuncore_FlavorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FlavorsAPIService FlavorsAwsInstanceTypes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsAwsInstanceTypes(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsAzureVmSizes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsAzureVmSizes(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsBindToProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.FlavorsAPI.FlavorsBindToProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsDropdownFlavors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsDropdownFlavors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsGoogleMachineTypes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsGoogleMachineTypes(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsOpenshiftFlavors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsOpenshiftFlavors(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsOpenstackFlavors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsOpenstackFlavors(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsProxmoxFlavors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsProxmoxFlavors(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsSelectedFlavorsForProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsSelectedFlavorsForProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsTanzuFlavors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.FlavorsAPI.FlavorsTanzuFlavors(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlavorsAPIService FlavorsUnbindFromProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.FlavorsAPI.FlavorsUnbindFromProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
