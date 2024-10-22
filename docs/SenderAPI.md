# \SenderAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSender**](SenderAPI.md#CreateSender) | **Post** /sender | Create Sender
[**GetAllSenders**](SenderAPI.md#GetAllSenders) | **Get** /sender | Get All Senders



## CreateSender

> Sender CreateSender(ctx).SenderRequest(senderRequest).Execute()

Create Sender



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
	senderRequest := *sendx.NewSenderRequest("Linus", "linus@linux.org") // SenderRequest | 

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.SenderAPI.CreateSender(context.Background()).SenderRequest(senderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenderAPI.CreateSender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSender`: Sender
	fmt.Fprintf(os.Stdout, "Response from `SenderAPI.CreateSender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderRequest** | [**SenderRequest**](SenderRequest.md) |  | 

### Return type

[**Sender**](Sender.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSenders

> []SenderResponse GetAllSenders(ctx).Offset(offset).Limit(limit).Search(search).Execute()

Get All Senders



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
	offset := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 10)
	search := "search_example" // string | Search keyword to filter senders by name or email (optional)

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.SenderAPI.GetAllSenders(context.Background()).Offset(offset).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenderAPI.GetAllSenders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSenders`: []SenderResponse
	fmt.Fprintf(os.Stdout, "Response from `SenderAPI.GetAllSenders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSendersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 10]
 **search** | **string** | Search keyword to filter senders by name or email | 

### Return type

[**[]SenderResponse**](SenderResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

