/*
SendX REST API

Testing ContactAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sendx

import (
	"context"
	"testing"

	sendx "github.com/sendx/sendx-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_sendx_ContactAPIService(t *testing.T) {

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)

	t.Run("Test ContactAPIService CreateContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContactAPI.CreateContact(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactAPIService DeleteContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId string

		resp, httpRes, err := apiClient.ContactAPI.DeleteContact(context.Background(), contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactAPIService GetAllContacts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContactAPI.GetAllContacts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactAPIService GetContactById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId string

		resp, httpRes, err := apiClient.ContactAPI.GetContactById(context.Background(), contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactAPIService UnsubscribeContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId string

		resp, httpRes, err := apiClient.ContactAPI.UnsubscribeContact(context.Background(), contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactAPIService UpdateContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId string

		resp, httpRes, err := apiClient.ContactAPI.UpdateContact(context.Background(), contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
