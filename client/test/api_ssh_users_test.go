/*
Taikun - WebApi

Testing SshUsersAPIService

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

func Test_taikuncore_SshUsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SshUsersAPIService SshusersCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SshUsersAPI.SshusersCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersAPIService SshusersDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SshUsersAPI.SshusersDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersAPIService SshusersEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SshUsersAPI.SshusersEdit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SshUsersAPIService SshusersList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessProfileId int32

		resp, httpRes, err := apiClient.SshUsersAPI.SshusersList(context.Background(), accessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
