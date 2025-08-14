# \EventsAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackCustomEvent**](EventsAPI.md#TrackCustomEvent) | **Post** /events/custom | Track custom event
[**TrackRevenueEvent**](EventsAPI.md#TrackRevenueEvent) | **Post** /events/revenue | Track revenue event



## TrackCustomEvent

> EventResponse TrackCustomEvent(ctx).CustomEventRequest(customEventRequest).Execute()

Track custom event



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
	customEventRequest := *openapiclient.NewCustomEventRequest("video_watched") // CustomEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.TrackCustomEvent(context.Background()).CustomEventRequest(customEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.TrackCustomEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackCustomEvent`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.TrackCustomEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackCustomEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEventRequest** | [**CustomEventRequest**](CustomEventRequest.md) |  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackRevenueEvent

> EventResponse TrackRevenueEvent(ctx).RevenueEventRequest(revenueEventRequest).Execute()

Track revenue event



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
	revenueEventRequest := *openapiclient.NewRevenueEventRequest("customer@example.com", float32(99.99)) // RevenueEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.TrackRevenueEvent(context.Background()).RevenueEventRequest(revenueEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.TrackRevenueEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackRevenueEvent`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.TrackRevenueEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackRevenueEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revenueEventRequest** | [**RevenueEventRequest**](RevenueEventRequest.md) |  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

