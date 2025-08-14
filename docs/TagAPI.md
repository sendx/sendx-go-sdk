# \TagAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagAPI.md#CreateTag) | **Post** /tag | Create tag
[**DeleteTag**](TagAPI.md#DeleteTag) | **Delete** /tag/{identifier} | Delete tag
[**GetAllTags**](TagAPI.md#GetAllTags) | **Get** /tag | Get all tags
[**GetTag**](TagAPI.md#GetTag) | **Get** /tag/{identifier} | Get tag by ID
[**UpdateTag**](TagAPI.md#UpdateTag) | **Put** /tag/{identifier} | Update tag



## CreateTag

> RestRTag CreateTag(ctx).RestETag(restETag).Execute()

Create tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func main() {
	restETag := *openapiclient.NewRestETag("VIP Customer") // RestETag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.CreateTag(context.Background()).RestETag(restETag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: RestRTag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restETag** | [**RestETag**](RestETag.md) |  | 

### Return type

[**RestRTag**](RestRTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteResponse DeleteTag(ctx, identifier).Execute()

Delete tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func main() {
	identifier := "identifier_example" // string | Tag identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.DeleteTag(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTag`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Tag identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTags

> []RestRTag GetAllTags(ctx).Offset(offset).Limit(limit).Execute()

Get all tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func main() {
	offset := int32(56) // int32 | Number of tags to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of tags to return (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.GetAllTags(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.GetAllTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTags`: []RestRTag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.GetAllTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of tags to skip | [default to 0]
 **limit** | **int32** | Maximum number of tags to return | [default to 10]

### Return type

[**[]RestRTag**](RestRTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> RestRTag GetTag(ctx, identifier).Execute()

Get tag by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func main() {
	identifier := "identifier_example" // string | Tag identifier to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.GetTag(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: RestRTag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Tag identifier to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRTag**](RestRTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> RestRTag UpdateTag(ctx, identifier).RestETag(restETag).Execute()

Update tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendx/sendx-go-sdk"
)

func main() {
	restETag := *openapiclient.NewRestETag("VIP Customer") // RestETag | 
	identifier := "identifier_example" // string | Tag identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAPI.UpdateTag(context.Background(), identifier).RestETag(restETag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: RestRTag
	fmt.Fprintf(os.Stdout, "Response from `TagAPI.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Tag identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restETag** | [**RestETag**](RestETag.md) |  | 


### Return type

[**RestRTag**](RestRTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

