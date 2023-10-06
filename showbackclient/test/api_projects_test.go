/*
Taikun - Showback API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikunshowback

import (
	"context"
	openapiclient "github.com/smidra/taikungoclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikunshowback_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService ProjectsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
