/*
SendX REST API

Testing ListAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sendx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func Test_sendx_ListAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListAPIService CreateList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListAPI.CreateList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListAPIService DeleteList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListAPI.DeleteList(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListAPIService GetAllLists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListAPI.GetAllLists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListAPIService GetListById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListAPI.GetListById(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListAPIService UpdateList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListAPI.UpdateList(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}