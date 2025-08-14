# \TrackingAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentifyContact**](TrackingAPI.md#IdentifyContact) | **Post** /contact/identify | Identify contact
[**TrackContact**](TrackingAPI.md#TrackContact) | **Post** /contact/track | Track contact



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
	identifyRequest := *openapiclient.NewIdentifyRequest("Email_example") // IdentifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingAPI.IdentifyContact(context.Background()).IdentifyRequest(identifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingAPI.IdentifyContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentifyContact`: IdentifyResponse
	fmt.Fprintf(os.Stdout, "Response from `TrackingAPI.IdentifyContact`: %v\n", resp)
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

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackContact

> TrackResponse TrackContact(ctx).TrackRequest(trackRequest).Execute()

Track contact



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
	resp, r, err := apiClient.TrackingAPI.TrackContact(context.Background()).TrackRequest(trackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingAPI.TrackContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackContact`: TrackResponse
	fmt.Fprintf(os.Stdout, "Response from `TrackingAPI.TrackContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackRequest** | [**TrackRequest**](TrackRequest.md) |  | 

### Return type

[**TrackResponse**](TrackResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

