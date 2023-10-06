/*
Taikun - WebApi

Testing SecurityGroupAPIService

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

func Test_taikuncore_SecurityGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityGroupAPIService SecuritygroupCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SecurityGroupAPI.SecuritygroupCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityGroupAPIService SecuritygroupDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.SecurityGroupAPI.SecuritygroupDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityGroupAPIService SecuritygroupEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SecurityGroupAPI.SecuritygroupEdit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityGroupAPIService SecuritygroupList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var standAloneProfileId int32

		resp, httpRes, err := apiClient.SecurityGroupAPI.SecuritygroupList(context.Background(), standAloneProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
