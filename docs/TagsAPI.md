# \TagsAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsAPI.md#CreateTag) | **Post** /tag | Create a Tag
[**DeleteTag**](TagsAPI.md#DeleteTag) | **Delete** /tag/{tagId} | Delete a Tag
[**GetAllTags**](TagsAPI.md#GetAllTags) | **Get** /tag | Get All Tags
[**GetTagById**](TagsAPI.md#GetTagById) | **Get** /tag/{tagId} | Get a Tag by ID
[**UpdateTag**](TagsAPI.md#UpdateTag) | **Put** /tag/{tagId} | Update a Tag



## CreateTag

> CreateResponse CreateTag(ctx).TagRequest(tagRequest).Execute()

Create a Tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	tagRequest := *sendx.NewTagRequest() // TagRequest | The tag content

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.CreateTag(context.Background()).TagRequest(tagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: CreateResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagRequest** | [**TagRequest**](TagRequest.md) | The tag content | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteResponse DeleteTag(ctx, tagId).Execute()

Delete a Tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	tagId := "tagId_example" // string | ID of the tag to delete

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.DeleteTag(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTag`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | ID of the tag to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTags

> []Tag GetAllTags(ctx).Offset(offset).Limit(limit).Search(search).Execute()

Get All Tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	offset := int32(56) // int32 | Offset for pagination (optional)
	limit := int32(56) // int32 | Limit the number of results (optional)
	search := "search_example" // string | Search term to filter tags (optional)

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetAllTags(context.Background()).Offset(offset).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetAllTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTags`: []Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetAllTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset for pagination | 
 **limit** | **int32** | Limit the number of results | 
 **search** | **string** | Search term to filter tags | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagById

> Tag GetTagById(ctx, tagId).Execute()

Get a Tag by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	tagId := "tagId_example" // string | ID of the tag you want to fetch

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagById(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagById`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | ID of the tag you want to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> Response UpdateTag(ctx, tagId).TagRequest(tagRequest).Execute()

Update a Tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	tagRequest := *sendx.NewTagRequest() // TagRequest | The tag content
	tagId := "tagId_example" // string | ID of the tag to update

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.UpdateTag(context.Background(), tagId).TagRequest(tagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | ID of the tag to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagRequest** | [**TagRequest**](TagRequest.md) | The tag content | 


### Return type

[**Response**](Response.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

