/*
Taikun - WebApi

Testing KubeConfigRoleAPIService

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

func Test_taikunshowback_KubeConfigRoleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubeConfigRoleAPIService KubeconfigroleList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubeConfigRoleAPI.KubeconfigroleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
