/*
Taikun - WebApi

Testing ProjectGroupsAPIService

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

func Test_taikuncore_ProjectGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectGroupsAPIService ProjectgroupsBind", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsBind(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectGroupId int32

		httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsEdit(context.Background(), projectGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsListByUserId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userGroupId int32

		resp, httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsListByUserId(context.Background(), userGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectGroupsAPIService ProjectgroupsProjectList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectGroupId int32

		resp, httpRes, err := apiClient.ProjectGroupsAPI.ProjectgroupsProjectList(context.Background(), projectGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
