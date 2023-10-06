/*
Taikun - WebApi

Testing PartnersAPIService

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

func Test_taikuncore_PartnersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PartnersAPIService PartnerAddWhitelistDomain", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerAddWhitelistDomain(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerBecomeAPartner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerBecomeAPartner(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerBindOrganizations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerBindOrganizations(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerContactUs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerContactUs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerDeleteWhitelistDomain", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PartnersAPI.PartnerDeleteWhitelistDomain(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PartnersAPI.PartnerDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerDropdown", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PartnersAPI.PartnerDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PartnersAPI.PartnerInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PartnersAPI.PartnerList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartnersAPIService PartnerUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.PartnersAPI.PartnerUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
