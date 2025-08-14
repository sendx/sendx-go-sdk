# \CustomFieldAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomField**](CustomFieldAPI.md#CreateCustomField) | **Post** /customfield | Create custom field
[**DeleteCustomField**](CustomFieldAPI.md#DeleteCustomField) | **Delete** /customfield/{identifier} | Delete custom field
[**GetAllCustomFields**](CustomFieldAPI.md#GetAllCustomFields) | **Get** /customfield | Get all custom fields
[**GetCustomField**](CustomFieldAPI.md#GetCustomField) | **Get** /customfield/{identifier} | Get custom field by ID
[**UpdateCustomField**](CustomFieldAPI.md#UpdateCustomField) | **Put** /customfield/{identifier} | Update custom field



## CreateCustomField

> RestRCustomField CreateCustomField(ctx).RestECustomField(restECustomField).Execute()

Create custom field



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
	restECustomField := *openapiclient.NewRestECustomField("Account Type", int32(123)) // RestECustomField | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.CreateCustomField(context.Background()).RestECustomField(restECustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: RestRCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restECustomField** | [**RestECustomField**](RestECustomField.md) |  | 

### Return type

[**RestRCustomField**](RestRCustomField.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteResponse DeleteCustomField(ctx, identifier).Execute()

Delete custom field



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
	identifier := "identifier_example" // string | Custom field identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.DeleteCustomField(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomField`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.DeleteCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Custom field identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


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


## GetAllCustomFields

> []RestRCustomField GetAllCustomFields(ctx).Offset(offset).Limit(limit).Search(search).Execute()

Get all custom fields



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
	offset := int32(56) // int32 | Number of fields to skip for pagination (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of fields to return (optional) (default to 10)
	search := "search_example" // string | Search custom fields by name (case-insensitive partial matching).  **Examples:** - `points` - Finds \"Loyalty points\", \"Reward points\"  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.GetAllCustomFields(context.Background()).Offset(offset).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.GetAllCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCustomFields`: []RestRCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.GetAllCustomFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of fields to skip for pagination | [default to 0]
 **limit** | **int32** | Maximum number of fields to return | [default to 10]
 **search** | **string** | Search custom fields by name (case-insensitive partial matching).  **Examples:** - &#x60;points&#x60; - Finds \&quot;Loyalty points\&quot;, \&quot;Reward points\&quot;  | 

### Return type

[**[]RestRCustomField**](RestRCustomField.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomField

> RestRCustomField GetCustomField(ctx, identifier).Execute()

Get custom field by ID



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
	identifier := "identifier_example" // string | Custom field identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.GetCustomField(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.GetCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomField`: RestRCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.GetCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Custom field identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRCustomField**](RestRCustomField.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> RestRCustomField UpdateCustomField(ctx, identifier).RestECustomField(restECustomField).Execute()

Update custom field



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
	restECustomField := *openapiclient.NewRestECustomField("Account Type", int32(123)) // RestECustomField | 
	identifier := "identifier_example" // string | Custom field identifier to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.UpdateCustomField(context.Background(), identifier).RestECustomField(restECustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomField`: RestRCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Custom field identifier to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restECustomField** | [**RestECustomField**](RestECustomField.md) |  | 


### Return type

[**RestRCustomField**](RestRCustomField.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

