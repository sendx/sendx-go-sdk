# \GettingStartedAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentifyContact**](GettingStartedAPI.md#IdentifyContact) | **Post** /contact/identify | Identify contact
[**TrackingContact**](GettingStartedAPI.md#TrackingContact) | **Post** /contact/track | Add Tracking info



## IdentifyContact

> IdentifyResponse IdentifyContact(ctx).IdentifyRequest(identifyRequest).Execute()

Identify contact



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
	identifyRequest := *openapiclient.NewIdentifyRequest("user@example.com") // IdentifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GettingStartedAPI.IdentifyContact(context.Background()).IdentifyRequest(identifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GettingStartedAPI.IdentifyContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentifyContact`: IdentifyResponse
	fmt.Fprintf(os.Stdout, "Response from `GettingStartedAPI.IdentifyContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentifyContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifyRequest** | [**IdentifyRequest**](IdentifyRequest.md) |  | 

### Return type

[**IdentifyResponse**](IdentifyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackingContact

> TrackResponse TrackingContact(ctx).TrackRequest(trackRequest).Execute()

Add Tracking info



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
	trackRequest := *openapiclient.NewTrackRequest() // TrackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GettingStartedAPI.TrackingContact(context.Background()).TrackRequest(trackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GettingStartedAPI.TrackingContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackingContact`: TrackResponse
	fmt.Fprintf(os.Stdout, "Response from `GettingStartedAPI.TrackingContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackingContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackRequest** | [**TrackRequest**](TrackRequest.md) |  | 

### Return type

[**TrackResponse**](TrackResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

