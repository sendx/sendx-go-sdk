# \WebhookAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookAPI.md#CreateWebhook) | **Post** /webhook | Create webhook
[**DeleteWebhook**](WebhookAPI.md#DeleteWebhook) | **Delete** /webhook/{identifier} | Delete webhook
[**GetAllWebhooks**](WebhookAPI.md#GetAllWebhooks) | **Get** /webhook | Get all webhooks
[**GetWebhook**](WebhookAPI.md#GetWebhook) | **Get** /webhook/{identifier} | Get webhook by ID
[**UpdateWebhook**](WebhookAPI.md#UpdateWebhook) | **Put** /webhook/{identifier} | Update webhook



## CreateWebhook

> RestRWebhook CreateWebhook(ctx).RestEWebhook(restEWebhook).Execute()

Create webhook



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
	restEWebhook := *openapiclient.NewRestEWebhook("https://api.example.com/webhooks/sendx", false) // RestEWebhook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.CreateWebhook(context.Background()).RestEWebhook(restEWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: RestRWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEWebhook** | [**RestEWebhook**](RestEWebhook.md) |  | 

### Return type

[**RestRWebhook**](RestRWebhook.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteResponse DeleteWebhook(ctx, identifier).Execute()

Delete webhook



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
	identifier := "identifier_example" // string | Webhook identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.DeleteWebhook(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Webhook identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetAllWebhooks

> []RestRWebhook GetAllWebhooks(ctx).Execute()

Get all webhooks



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
	resp, r, err := apiClient.WebhookAPI.GetAllWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetAllWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWebhooks`: []RestRWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetAllWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWebhooksRequest struct via the builder pattern


### Return type

[**[]RestRWebhook**](RestRWebhook.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> RestRWebhook GetWebhook(ctx, identifier).Execute()

Get webhook by ID



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
	identifier := "identifier_example" // string | Webhook identifier to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.GetWebhook(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: RestRWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Webhook identifier to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRWebhook**](RestRWebhook.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> RestRWebhook UpdateWebhook(ctx, identifier).RestEWebhook(restEWebhook).Execute()

Update webhook



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
	restEWebhook := *openapiclient.NewRestEWebhook("https://api.example.com/webhooks/sendx", false) // RestEWebhook | 
	identifier := "identifier_example" // string | Webhook identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateWebhook(context.Background(), identifier).RestEWebhook(restEWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: RestRWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Webhook identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restEWebhook** | [**RestEWebhook**](RestEWebhook.md) |  | 


### Return type

[**RestRWebhook**](RestRWebhook.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

