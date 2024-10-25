/*
SendX REST API

# Introduction SendX is an email marketing product. It helps you convert website visitors to customers, send them promotional emails, engage with them using drip sequences and craft custom journeys using powerful but simple automations. The SendX API is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs. The SendX Rest API doesn’t support bulk updates. You can work on only one object per request. <br> 

API version: 1.0.0
Contact: support@sendx.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ReportsAPIService ReportsAPI service
type ReportsAPIService service

type ApiGetCampaignReportRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
	campaignId string
	integrationType *string
}

// Type of integration for the report data (optional)
func (r ApiGetCampaignReportRequest) IntegrationType(integrationType string) ApiGetCampaignReportRequest {
	r.integrationType = &integrationType
	return r
}

func (r ApiGetCampaignReportRequest) Execute() (*ReportData, *http.Response, error) {
	return r.ApiService.GetCampaignReportExecute(r)
}

/*
GetCampaignReport Get CampaignReport Data

Retrieve the campaign report data based on the provided campaign id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignId The ID of the campaign to retrieve the report data for
 @return ApiGetCampaignReportRequest
*/
func (a *ReportsAPIService) GetCampaignReport(ctx context.Context, campaignId string) ApiGetCampaignReportRequest {
	return ApiGetCampaignReportRequest{
		ApiService: a,
		ctx: ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//  @return ReportData
func (a *ReportsAPIService) GetCampaignReportExecute(r ApiGetCampaignReportRequest) (*ReportData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetCampaignReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/report/campaign/{campaignId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.integrationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "integrationType", r.integrationType, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Team-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
