/*
Taikun - WebApi

Testing ProjectAppParamsAPIService

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

func Test_taikuncore_ProjectAppParamsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAppParamsAPIService ProjectappparamEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectAppId int32

		httpRes, err := apiClient.ProjectAppParamsAPI.ProjectappparamEdit(context.Background(), projectAppId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}