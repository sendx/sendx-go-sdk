# \ContactAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](ContactAPI.md#CreateContact) | **Post** /contact | Create a contact
[**DeleteContact**](ContactAPI.md#DeleteContact) | **Delete** /contact/{identifier} | Delete Contact
[**GetAllContacts**](ContactAPI.md#GetAllContacts) | **Get** /contact | Get All Contacts
[**GetContactById**](ContactAPI.md#GetContactById) | **Get** /contact/{identifier} | Get Contact by Identifier
[**UnsubscribeContact**](ContactAPI.md#UnsubscribeContact) | **Post** /contact/unsubscribe/{identifier} | Unsubscribe Contact
[**UpdateContact**](ContactAPI.md#UpdateContact) | **Put** /contact/{identifier} | Update Contact



## CreateContact

> Response CreateContact(ctx).ContactRequest(contactRequest).Execute()

Create a contact



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
	contactRequest := *openapiclient.NewContactRequest() // ContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.CreateContact(context.Background()).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: Response
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactRequest** | [**ContactRequest**](ContactRequest.md) |  | 

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


## DeleteContact

> Response DeleteContact(ctx, identifier).Execute()

Delete Contact



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
	identifier := "identifier_example" // string | The Contact ID/ Email to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.DeleteContact(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.DeleteContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContact`: Response
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.DeleteContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The Contact ID/ Email to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllContacts

> []Contact GetAllContacts(ctx).Offset(offset).Limit(limit).ContactType(contactType).Search(search).Execute()

Get All Contacts



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
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)
	limit := int32(56) // int32 | Limit for pagination (optional) (default to 10)
	contactType := "contactType_example" // string | Filter contacts by type (optional)
	search := "search_example" // string | Search term to filter contacts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetAllContacts(context.Background()).Offset(offset).Limit(limit).ContactType(contactType).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetAllContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllContacts`: []Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetAllContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset for pagination | [default to 0]
 **limit** | **int32** | Limit for pagination | [default to 10]
 **contactType** | **string** | Filter contacts by type | 
 **search** | **string** | Search term to filter contacts | 

### Return type

[**[]Contact**](Contact.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactById

> Contact GetContactById(ctx, identifier).Execute()

Get Contact by Identifier



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
	identifier := "john@doe.com" // string | The ID or Email of the contact to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetContactById(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetContactById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactById`: Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetContactById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The ID or Email of the contact to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeContact

> Response UnsubscribeContact(ctx, identifier).Execute()

Unsubscribe Contact



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
	identifier := "sendx123" // string | The Contact ID or email to unsubscribe

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.UnsubscribeContact(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.UnsubscribeContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeContact`: Response
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.UnsubscribeContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The Contact ID or email to unsubscribe | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> Contact UpdateContact(ctx, identifier).ContactRequest(contactRequest).Execute()

Update Contact



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
	contactRequest := *openapiclient.NewContactRequest() // ContactRequest | 
	identifier := "sendxid123" // string | The ID or email of the Contact to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.UpdateContact(context.Background(), identifier).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The ID or email of the Contact to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactRequest** | [**ContactRequest**](ContactRequest.md) |  | 


### Return type

[**Contact**](Contact.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

