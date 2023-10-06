/*
Taikun - WebApi

Testing TicketAPIService

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

func Test_taikuncore_TicketAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TicketAPIService TicketArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketArchive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketClose", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketClose(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ticketId string

		httpRes, err := apiClient.TicketAPI.TicketDelete(context.Background(), ticketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketDeleteMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		httpRes, err := apiClient.TicketAPI.TicketDeleteMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketEdit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketEditMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketEditMessage(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TicketAPI.TicketList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketMessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ticketId string

		resp, httpRes, err := apiClient.TicketAPI.TicketMessages(context.Background(), ticketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketOpen", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketOpen(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketReply", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketReply(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketSetPriority", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketSetPriority(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketTransfer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TicketAPI.TicketTransfer(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TicketAPIService TicketTransferList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TicketAPI.TicketTransferList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
