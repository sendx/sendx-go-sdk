# \EmailSendingAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmail**](EmailSendingAPI.md#SendEmail) | **Post** /send/email | Send transactional email
[**SendEmailWithTemplate**](EmailSendingAPI.md#SendEmailWithTemplate) | **Post** /send/template | Send email using template



## SendEmail

> []XEmailResponse SendEmail(ctx).XEmailMessage(xEmailMessage).Execute()

Send transactional email



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
	xEmailMessage := *openapiclient.NewXEmailMessage(*openapiclient.NewXFrom("from@example.com"), []openapiclient.XTo{*openapiclient.NewXTo("to@example.com")}, "Your Subject Here", "<h1>Your HTML Content</h1>") // XEmailMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSendingAPI.SendEmail(context.Background()).XEmailMessage(xEmailMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSendingAPI.SendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmail`: []XEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailSendingAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEmailMessage** | [**XEmailMessage**](XEmailMessage.md) |  | 

### Return type

[**[]XEmailResponse**](XEmailResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailWithTemplate

> []XEmailResponse SendEmailWithTemplate(ctx).TemplateEmailMessage(templateEmailMessage).Execute()

Send email using template



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
	templateEmailMessage := *openapiclient.NewTemplateEmailMessage(*openapiclient.NewXFrom("from@example.com"), []openapiclient.XTo{*openapiclient.NewXTo("to@example.com")}, "Your Subject Here", "template_f3lJvTEhSjKGVb5Lwc5SWS") // TemplateEmailMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSendingAPI.SendEmailWithTemplate(context.Background()).TemplateEmailMessage(templateEmailMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSendingAPI.SendEmailWithTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailWithTemplate`: []XEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailSendingAPI.SendEmailWithTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailWithTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateEmailMessage** | [**TemplateEmailMessage**](TemplateEmailMessage.md) |  | 

### Return type

[**[]XEmailResponse**](XEmailResponse.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

