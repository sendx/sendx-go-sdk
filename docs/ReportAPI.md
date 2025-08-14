# \ReportAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignReport**](ReportAPI.md#GetCampaignReport) | **Get** /report/campaign/{identifier} | Get campaign report



## GetCampaignReport

> RestReportData GetCampaignReport(ctx, identifier).Execute()

Get campaign report



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
	identifier := "contact_BnKjkbBBS500CoBCP0oChQ" // string | Resource identifier with prefix (e.g., `contact_BnKjkbBBS500CoBCP0oChQ`)  **Format:** `<prefix>_<22-character-id>` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportAPI.GetCampaignReport(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.GetCampaignReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignReport`: RestReportData
	fmt.Fprintf(os.Stdout, "Response from `ReportAPI.GetCampaignReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Resource identifier with prefix (e.g., &#x60;contact_BnKjkbBBS500CoBCP0oChQ&#x60;)  **Format:** &#x60;&lt;prefix&gt;_&lt;22-character-id&gt;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestReportData**](RestReportData.md)

### Authorization

[TeamApiKey](../README.md#TeamApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

