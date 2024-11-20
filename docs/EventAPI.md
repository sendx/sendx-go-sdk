# \EventAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRevenueEvent**](EventAPI.md#CreateRevenueEvent) | **Post** /events/revenue | Record a revenue event for a specific contact
[**PushCustomEvent**](EventAPI.md#PushCustomEvent) | **Post** /events/custom | Push a custom event associated with a contact



## CreateRevenueEvent

> EventResponse CreateRevenueEvent(ctx).RevenueEventRequest(revenueEventRequest).Execute()

Record a revenue event for a specific contact



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
	revenueEventRequest := *openapiclient.NewRevenueEventRequest("john.doe@example.com", float32(123.23), "INR", "app", int32(1669990400)) // RevenueEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.CreateRevenueEvent(context.Background()).RevenueEventRequest(revenueEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.CreateRevenueEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRevenueEvent`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.CreateRevenueEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRevenueEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revenueEventRequest** | [**RevenueEventRequest**](RevenueEventRequest.md) |  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushCustomEvent

> EventResponse PushCustomEvent(ctx).CustomEventRequest(customEventRequest).Execute()

Push a custom event associated with a contact



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
	customEventRequest := *openapiclient.NewCustomEventRequest("abandoned_cart", "john.doe@example.com", map[string]string{"key": "Inner_example"}, int32(1669990400)) // CustomEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.PushCustomEvent(context.Background()).CustomEventRequest(customEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.PushCustomEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PushCustomEvent`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.PushCustomEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushCustomEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEventRequest** | [**CustomEventRequest**](CustomEventRequest.md) |  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

