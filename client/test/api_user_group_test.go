/*
Taikun - WebApi

Testing UserGroupAPIService

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

func Test_taikuncore_UserGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserGroupAPIService UsergroupsBindProjectsGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserGroupAPI.UsergroupsBindProjectsGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserGroupAPI.UsergroupsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserGroupAPI.UsergroupsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserGroupAPI.UsergroupsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsListByProjectGroupId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectGroupId int32

		resp, httpRes, err := apiClient.UserGroupAPI.UsergroupsListByProjectGroupId(context.Background(), projectGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userGroupId int32

		httpRes, err := apiClient.UserGroupAPI.UsergroupsUpdate(context.Background(), userGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UsergroupsUserGroupUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userGroupId int32

		resp, httpRes, err := apiClient.UserGroupAPI.UsergroupsUserGroupUsers(context.Background(), userGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}