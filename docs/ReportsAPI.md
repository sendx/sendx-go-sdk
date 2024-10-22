# \ReportsAPI

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignReport**](ReportsAPI.md#GetCampaignReport) | **Get** /report/campaign/{campaignId} | Get CampaignReport Data



## GetCampaignReport

> ReportData GetCampaignReport(ctx, campaignId).IntegrationType(integrationType).Execute()

Get CampaignReport Data



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
	campaignId := "campaignId_example" // string | The ID of the campaign to retrieve the report data for
	integrationType := "integrationType_example" // string | Type of integration for the report data (optional) (optional)

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetCampaignReport(context.Background(), campaignId).IntegrationType(integrationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetCampaignReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignReport`: ReportData
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetCampaignReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign to retrieve the report data for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationType** | **string** | Type of integration for the report data (optional) | 

### Return type

[**ReportData**](ReportData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

