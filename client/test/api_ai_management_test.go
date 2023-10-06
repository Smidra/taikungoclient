/*
Taikun - WebApi

Testing AiManagementAPIService

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

func Test_taikuncore_AiManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AiManagementAPIService AiManagementDisable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AiManagementAPI.AiManagementDisable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AiManagementAPIService AiManagementEnable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AiManagementAPI.AiManagementEnable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
