# \PostAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](PostAPI.md#CreatePost) | **Post** /post | Create blog post
[**DeletePost**](PostAPI.md#DeletePost) | **Delete** /post/{identifier} | Delete post
[**GetAllPosts**](PostAPI.md#GetAllPosts) | **Get** /post | Get all posts
[**GetPost**](PostAPI.md#GetPost) | **Get** /post/{identifier} | Get post by ID
[**UpdatePost**](PostAPI.md#UpdatePost) | **Put** /post/{identifier} | Update post



## CreatePost

> RestRPost CreatePost(ctx).RestEPost(restEPost).Execute()

Create blog post



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
	restEPost := *openapiclient.NewRestEPost("Summer Product Launch", "Introducing Our New Summer Collection") // RestEPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.CreatePost(context.Background()).RestEPost(restEPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.CreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePost`: RestRPost
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.CreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPost** | [**RestEPost**](RestEPost.md) |  | 

### Return type

[**RestRPost**](RestRPost.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePost

> MessageResponse DeletePost(ctx, identifier).Execute()

Delete post



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
	identifier := "contact_BnKjkbBBS500CoBCP0oChQ" // string | Resource identifier with prefix (e.g., `contact_BnKjkbBBS500CoBCP0oChQ`)  **Format:** `<prefix>_<22-character-id>` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.DeletePost(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.DeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePost`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.DeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Resource identifier with prefix (e.g., &#x60;contact_BnKjkbBBS500CoBCP0oChQ&#x60;)  **Format:** &#x60;&lt;prefix&gt;_&lt;22-character-id&gt;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostRequest struct via the builder pattern


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


## GetAllPosts

> []RestRPost GetAllPosts(ctx).Offset(offset).Limit(limit).Execute()

Get all posts



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
	offset := int32(56) // int32 | Number of posts to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of posts to return (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.GetAllPosts(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.GetAllPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPosts`: []RestRPost
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.GetAllPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of posts to skip | [default to 0]
 **limit** | **int32** | Maximum number of posts to return | [default to 10]

### Return type

[**[]RestRPost**](RestRPost.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPost

> RestRPost GetPost(ctx, identifier).Execute()

Get post by ID



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
	identifier := "contact_BnKjkbBBS500CoBCP0oChQ" // string | Resource identifier with prefix (e.g., `contact_BnKjkbBBS500CoBCP0oChQ`)  **Format:** `<prefix>_<22-character-id>` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.GetPost(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.GetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPost`: RestRPost
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.GetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Resource identifier with prefix (e.g., &#x60;contact_BnKjkbBBS500CoBCP0oChQ&#x60;)  **Format:** &#x60;&lt;prefix&gt;_&lt;22-character-id&gt;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRPost**](RestRPost.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePost

> RestRPost UpdatePost(ctx, identifier).RestEPost(restEPost).Execute()

Update post



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
	restEPost := *openapiclient.NewRestEPost("Summer Product Launch", "Introducing Our New Summer Collection") // RestEPost | 
	identifier := "contact_BnKjkbBBS500CoBCP0oChQ" // string | Resource identifier with prefix (e.g., `contact_BnKjkbBBS500CoBCP0oChQ`)  **Format:** `<prefix>_<22-character-id>` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.UpdatePost(context.Background(), identifier).RestEPost(restEPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.UpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePost`: RestRPost
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.UpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Resource identifier with prefix (e.g., &#x60;contact_BnKjkbBBS500CoBCP0oChQ&#x60;)  **Format:** &#x60;&lt;prefix&gt;_&lt;22-character-id&gt;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEPost** | [**RestEPost**](RestEPost.md) |  | 


### Return type

[**RestRPost**](RestRPost.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

