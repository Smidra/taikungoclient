/*
Taikun - WebApi

Testing SlackAPIService

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

func Test_taikuncore_SlackAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlackAPIService SlackCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlackAPI.SlackCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackAPIService SlackDeleteMultiple", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SlackAPI.SlackDeleteMultiple(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackAPIService SlackDropdown", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlackAPI.SlackDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackAPIService SlackList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlackAPI.SlackList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackAPIService SlackUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.SlackAPI.SlackUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlackAPIService SlackVerify", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SlackAPI.SlackVerify(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}