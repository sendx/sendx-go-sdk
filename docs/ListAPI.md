# \ListAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateList**](ListAPI.md#CreateList) | **Post** /list | Create List
[**DeleteList**](ListAPI.md#DeleteList) | **Delete** /list/{listId} | Delete List
[**GetAllLists**](ListAPI.md#GetAllLists) | **Get** /list | Get All Lists
[**GetListById**](ListAPI.md#GetListById) | **Get** /list/{listId} | Get List
[**UpdateList**](ListAPI.md#UpdateList) | **Put** /list/{listId} | Update List



## CreateList

> CreateResponse CreateList(ctx).ListRequest(listRequest).Execute()

Create List



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
	listRequest := *sendx.NewListRequest() // ListRequest | 

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.CreateList(context.Background()).ListRequest(listRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.CreateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateList`: CreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.CreateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listRequest** | [**ListRequest**](ListRequest.md) |  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteList

> DeleteResponse DeleteList(ctx, listId).Execute()

Delete List



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
	listId := "sendx123" // string | The ID of the list you want to delete

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.DeleteList(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.DeleteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteList`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.DeleteList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The ID of the list you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLists

> []ListModel GetAllLists(ctx).Offset(offset).Limit(limit).Search(search).Execute()

Get All Lists



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
	offset := int32(0) // int32 | Offset for pagination. (optional)
	limit := int32(10) // int32 | Limit the number of results returned. (optional)
	search := "Marketing" // string | Search term to filter lists. (optional)

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.GetAllLists(context.Background()).Offset(offset).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.GetAllLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLists`: []ListModel
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.GetAllLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset for pagination. | 
 **limit** | **int32** | Limit the number of results returned. | 
 **search** | **string** | Search term to filter lists. | 

### Return type

[**[]ListModel**](ListModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListById

> ListModel GetListById(ctx, listId).Execute()

Get List



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
	listId := "sendx123" // string | The ID of the list you want to retrieve

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.GetListById(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.GetListById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListById`: ListModel
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.GetListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The ID of the list you want to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListModel**](ListModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateList

> Response UpdateList(ctx, listId).ListRequest(listRequest).Execute()

Update List



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
	listRequest := *sendx.NewListRequest() // ListRequest | 
	listId := "listId_example" // string | The ID of the list to be updated.

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.UpdateList(context.Background(), listId).ListRequest(listRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.UpdateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateList`: Response
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.UpdateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The ID of the list to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listRequest** | [**ListRequest**](ListRequest.md) |  | 


### Return type

[**Response**](Response.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

