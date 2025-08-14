# \PostTagAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostTag**](PostTagAPI.md#CreatePostTag) | **Post** /post/tag | Create post tag
[**DeletePostTag**](PostTagAPI.md#DeletePostTag) | **Delete** /post/tag/{identifier} | Delete post tag
[**GetAllPostTags**](PostTagAPI.md#GetAllPostTags) | **Get** /post/tag | Get all post tags
[**GetPostTag**](PostTagAPI.md#GetPostTag) | **Get** /post/tag/{identifier} | Get post tag by ID
[**UpdatePostTag**](PostTagAPI.md#UpdatePostTag) | **Put** /post/tag/{identifier} | Update post tag



## CreatePostTag

> RestRPostTag CreatePostTag(ctx).RestEPostTag(restEPostTag).Execute()

Create post tag



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
	restEPostTag := *openapiclient.NewRestEPostTag("Summer Sale") // RestEPostTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostTagAPI.CreatePostTag(context.Background()).RestEPostTag(restEPostTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostTagAPI.CreatePostTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostTag`: RestRPostTag
	fmt.Fprintf(os.Stdout, "Response from `PostTagAPI.CreatePostTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPostTag** | [**RestEPostTag**](RestEPostTag.md) |  | 

### Return type

[**RestRPostTag**](RestRPostTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostTag

> MessageResponse DeletePostTag(ctx, identifier).Execute()

Delete post tag



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
	identifier := "post_tag_leBDiFdrUnRmRz4nfopSrv" // string | The unique post tag identifier to retrieve. - `post_tag_leBDiFdrUnRmRz4nfopSrv` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostTagAPI.DeletePostTag(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostTagAPI.DeletePostTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePostTag`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `PostTagAPI.DeletePostTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post tag identifier to retrieve. - &#x60;post_tag_leBDiFdrUnRmRz4nfopSrv&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPostTags

> []RestRPostTag GetAllPostTags(ctx).Execute()

Get all post tags



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostTagAPI.GetAllPostTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostTagAPI.GetAllPostTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPostTags`: []RestRPostTag
	fmt.Fprintf(os.Stdout, "Response from `PostTagAPI.GetAllPostTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPostTagsRequest struct via the builder pattern


### Return type

[**[]RestRPostTag**](RestRPostTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostTag

> RestRPostTag GetPostTag(ctx, identifier).Execute()

Get post tag by ID



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
	identifier := "post_tag_leBDiFdrUnRmRz4nfopSrv" // string | The unique post tag identifier to retrieve. - `post_tag_leBDiFdrUnRmRz4nfopSrv` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostTagAPI.GetPostTag(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostTagAPI.GetPostTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostTag`: RestRPostTag
	fmt.Fprintf(os.Stdout, "Response from `PostTagAPI.GetPostTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post tag identifier to retrieve. - &#x60;post_tag_leBDiFdrUnRmRz4nfopSrv&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRPostTag**](RestRPostTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostTag

> RestRPostTag UpdatePostTag(ctx, identifier).RestEPostTag(restEPostTag).Execute()

Update post tag



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
	restEPostTag := *openapiclient.NewRestEPostTag("Summer Sale") // RestEPostTag | 
	identifier := "post_tag_leBDiFdrUnRmRz4nfopSrv" // string | The unique post tag identifier to retrieve. - `post_tag_leBDiFdrUnRmRz4nfopSrv` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostTagAPI.UpdatePostTag(context.Background(), identifier).RestEPostTag(restEPostTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostTagAPI.UpdatePostTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostTag`: RestRPostTag
	fmt.Fprintf(os.Stdout, "Response from `PostTagAPI.UpdatePostTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post tag identifier to retrieve. - &#x60;post_tag_leBDiFdrUnRmRz4nfopSrv&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPostTag** | [**RestEPostTag**](RestEPostTag.md) |  | 


### Return type

[**RestRPostTag**](RestRPostTag.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

