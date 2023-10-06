/*
Taikun - WebApi

Testing BackupPolicyAPIService

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

func Test_taikuncore_BackupPolicyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BackupPolicyAPIService BackupByName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupByName(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupClearProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupClearProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDeleteBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupDeleteBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDeleteBackupLocation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupDeleteBackupLocation(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDeleteRestore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupDeleteRestore(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDeleteSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupDeleteSchedule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDescribeBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupDescribeBackup(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDescribeRestore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupDescribeRestore(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDescribeSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupDescribeSchedule(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupDisableBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupDisableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupEnableBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupEnableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupImportBackupStorage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupImportBackupStorage(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupListAllBackupStorages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupListAllBackupStorages(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupListAllBackups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupListAllBackups(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupListAllDeleteBackupRequests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupListAllDeleteBackupRequests(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupListAllRestores", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupListAllRestores(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupListAllSchedules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyAPI.BackupListAllSchedules(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyAPIService BackupRestoreBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BackupPolicyAPI.BackupRestoreBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
