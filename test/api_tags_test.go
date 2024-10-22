/*
SendX REST API

Testing TagsAPIService

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

func Test_sendx_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService CreateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.CreateTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService DeleteTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.DeleteTag(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetAllTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.GetAllTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetTagById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.GetTagById(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService UpdateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.UpdateTag(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}