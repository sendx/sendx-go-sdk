# \PostCategoryAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostCategory**](PostCategoryAPI.md#CreatePostCategory) | **Post** /post/category | Create post category
[**DeletePostCategory**](PostCategoryAPI.md#DeletePostCategory) | **Delete** /post/category/{identifier} | Delete post category
[**GetAllPostCategories**](PostCategoryAPI.md#GetAllPostCategories) | **Get** /post/category | Get all post categories
[**GetPostCategory**](PostCategoryAPI.md#GetPostCategory) | **Get** /post/category/{identifier} | Get post category by ID
[**UpdatePostCategory**](PostCategoryAPI.md#UpdatePostCategory) | **Put** /post/category/{identifier} | Update post category



## CreatePostCategory

> RestRPostCategory CreatePostCategory(ctx).RestEPostCategory(restEPostCategory).Execute()

Create post category



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
	restEPostCategory := *openapiclient.NewRestEPostCategory("Product Updates") // RestEPostCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostCategoryAPI.CreatePostCategory(context.Background()).RestEPostCategory(restEPostCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostCategoryAPI.CreatePostCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostCategory`: RestRPostCategory
	fmt.Fprintf(os.Stdout, "Response from `PostCategoryAPI.CreatePostCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPostCategory** | [**RestEPostCategory**](RestEPostCategory.md) |  | 

### Return type

[**RestRPostCategory**](RestRPostCategory.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostCategory

> MessageResponse DeletePostCategory(ctx, identifier).Execute()

Delete post category



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
	identifier := "post_category_YzS1wOU20yw87UUHKxMzwn" // string | The unique post category identifier to retrieve. - `post_category_YzS1wOU20yw87UUHKxMzwn` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostCategoryAPI.DeletePostCategory(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostCategoryAPI.DeletePostCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePostCategory`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `PostCategoryAPI.DeletePostCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post category identifier to retrieve. - &#x60;post_category_YzS1wOU20yw87UUHKxMzwn&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostCategoryRequest struct via the builder pattern


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


## GetAllPostCategories

> []RestRPostCategory GetAllPostCategories(ctx).Execute()

Get all post categories



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
	resp, r, err := apiClient.PostCategoryAPI.GetAllPostCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostCategoryAPI.GetAllPostCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPostCategories`: []RestRPostCategory
	fmt.Fprintf(os.Stdout, "Response from `PostCategoryAPI.GetAllPostCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPostCategoriesRequest struct via the builder pattern


### Return type

[**[]RestRPostCategory**](RestRPostCategory.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostCategory

> RestRPostCategory GetPostCategory(ctx, identifier).Execute()

Get post category by ID



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
	identifier := "post_category_YzS1wOU20yw87UUHKxMzwn" // string | The unique post category identifier to retrieve. - `post_category_YzS1wOU20yw87UUHKxMzwn` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostCategoryAPI.GetPostCategory(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostCategoryAPI.GetPostCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostCategory`: RestRPostCategory
	fmt.Fprintf(os.Stdout, "Response from `PostCategoryAPI.GetPostCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post category identifier to retrieve. - &#x60;post_category_YzS1wOU20yw87UUHKxMzwn&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRPostCategory**](RestRPostCategory.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostCategory

> RestRPostCategory UpdatePostCategory(ctx, identifier).RestEPostCategory(restEPostCategory).Execute()

Update post category



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
	restEPostCategory := *openapiclient.NewRestEPostCategory("Product Updates") // RestEPostCategory | 
	identifier := "post_category_YzS1wOU20yw87UUHKxMzwn" // string | The unique post category identifier to retrieve. - `post_category_YzS1wOU20yw87UUHKxMzwn` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostCategoryAPI.UpdatePostCategory(context.Background(), identifier).RestEPostCategory(restEPostCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostCategoryAPI.UpdatePostCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostCategory`: RestRPostCategory
	fmt.Fprintf(os.Stdout, "Response from `PostCategoryAPI.UpdatePostCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique post category identifier to retrieve. - &#x60;post_category_YzS1wOU20yw87UUHKxMzwn&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPostCategory** | [**RestEPostCategory**](RestEPostCategory.md) |  | 


### Return type

[**RestRPostCategory**](RestRPostCategory.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

