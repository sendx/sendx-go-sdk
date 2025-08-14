# \TemplateAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailTemplate**](TemplateAPI.md#CreateEmailTemplate) | **Post** /template/email | Create email template
[**DeleteEmailTemplate**](TemplateAPI.md#DeleteEmailTemplate) | **Delete** /template/email/{identifier} | Delete template
[**GetAllEmailTemplates**](TemplateAPI.md#GetAllEmailTemplates) | **Get** /template/email | Get all templates
[**GetEmailTemplate**](TemplateAPI.md#GetEmailTemplate) | **Get** /template/email/{identifier} | Get template by ID
[**UpdateEmailTemplate**](TemplateAPI.md#UpdateEmailTemplate) | **Put** /template/email/{identifier} | Update template



## CreateEmailTemplate

> RestRTemplate CreateEmailTemplate(ctx).RestETemplate(restETemplate).Execute()

Create email template



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
	restETemplate := *openapiclient.NewRestETemplate("Welcome Email Template", "Welcome to {{company_name}}!") // RestETemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.CreateEmailTemplate(context.Background()).RestETemplate(restETemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.CreateEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailTemplate`: RestRTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.CreateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restETemplate** | [**RestETemplate**](RestETemplate.md) |  | 

### Return type

[**RestRTemplate**](RestRTemplate.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailTemplate

> DeleteResponse DeleteEmailTemplate(ctx, identifier).Execute()

Delete template



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
	identifier := "template_f3lJvTEhSjKGVb5Lwc5SWS" // string | The unique template identifier to update.  - `template_f3lJvTEhSjKGVb5Lwc5SWS` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.DeleteEmailTemplate(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.DeleteEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmailTemplate`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.DeleteEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique template identifier to update.  - &#x60;template_f3lJvTEhSjKGVb5Lwc5SWS&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailTemplateRequest struct via the builder pattern


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


## GetAllEmailTemplates

> []RestRTemplate GetAllEmailTemplates(ctx).Offset(offset).Limit(limit).Search(search).Execute()

Get all templates



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
	offset := int32(56) // int32 | Number of records to skip for pagination.  **Examples:** - `0` - First page (default) - `25` - Second page (with limit=25) - `50` - Third page (with limit=25)  (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of templates to return per page.  **Guidelines:** - Default: 10 templates - Maximum: 100 templates - Recommended: 25-100 for optimal performance  (optional) (default to 10)
	search := "search_example" // string | Search templates by name (case-insensitive partial matching).  **Examples:** - `newsletter` - Finds \"Weekly Newsletter\", \"Monthly Newsletter\" - `welcome` - Finds \"Welcome Email\", \"New User Welcome\" - `product` - Finds \"Product Launch\", \"Product Update\"  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.GetAllEmailTemplates(context.Background()).Offset(offset).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.GetAllEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllEmailTemplates`: []RestRTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.GetAllEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of records to skip for pagination.  **Examples:** - &#x60;0&#x60; - First page (default) - &#x60;25&#x60; - Second page (with limit&#x3D;25) - &#x60;50&#x60; - Third page (with limit&#x3D;25)  | [default to 0]
 **limit** | **int32** | Maximum number of templates to return per page.  **Guidelines:** - Default: 10 templates - Maximum: 100 templates - Recommended: 25-100 for optimal performance  | [default to 10]
 **search** | **string** | Search templates by name (case-insensitive partial matching).  **Examples:** - &#x60;newsletter&#x60; - Finds \&quot;Weekly Newsletter\&quot;, \&quot;Monthly Newsletter\&quot; - &#x60;welcome&#x60; - Finds \&quot;Welcome Email\&quot;, \&quot;New User Welcome\&quot; - &#x60;product&#x60; - Finds \&quot;Product Launch\&quot;, \&quot;Product Update\&quot;  | 

### Return type

[**[]RestRTemplate**](RestRTemplate.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplate

> RestRTemplate GetEmailTemplate(ctx, identifier).Execute()

Get template by ID



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
	identifier := "template_f3lJvTEhSjKGVb5Lwc5SWS" // string | The unique template identifier.  - `template_f3lJvTEhSjKGVb5Lwc5SWS` - Standard prefixed ID 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.GetEmailTemplate(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.GetEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailTemplate`: RestRTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique template identifier.  - &#x60;template_f3lJvTEhSjKGVb5Lwc5SWS&#x60; - Standard prefixed ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestRTemplate**](RestRTemplate.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailTemplate

> RestRTemplate UpdateEmailTemplate(ctx, identifier).RestETemplate(restETemplate).Execute()

Update template



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
	restETemplate := *openapiclient.NewRestETemplate("Welcome Email Template", "Welcome to {{company_name}}!") // RestETemplate | 
	identifier := "template_f3lJvTEhSjKGVb5Lwc5SWS" // string | The unique template identifier to update.  - `template_f3lJvTEhSjKGVb5Lwc5SWS` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.UpdateEmailTemplate(context.Background(), identifier).RestETemplate(restETemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.UpdateEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailTemplate`: RestRTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.UpdateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The unique template identifier to update.  - &#x60;template_f3lJvTEhSjKGVb5Lwc5SWS&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restETemplate** | [**RestETemplate**](RestETemplate.md) |  | 


### Return type

[**RestRTemplate**](RestRTemplate.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

