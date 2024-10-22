# SendX Go SDK

# Introduction 

SendX is an email marketing product. It helps you convert website visitors to customers, send them promotional emails, engage with them using drip sequences and craft custom journeys using powerful but simple automations. The SendX API is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.
The SendX Rest API doesn’t support bulk updates. You can work on only one object per request. <br>


## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sendx "github.com/sendx/sendx-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sendx.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sendx.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sendx.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sendx.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sendx.ContextOperationServerIndices` and `sendx.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sendx.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sendx.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignAPI* | [**CreateCampaign**](docs/CampaignAPI.md#createcampaign) | **Post** /campaign | Create Campaign
*CampaignAPI* | [**DeleteCampaign**](docs/CampaignAPI.md#deletecampaign) | **Delete** /campaign/{campaignId} | Delete Campaign
*CampaignAPI* | [**EditCampaign**](docs/CampaignAPI.md#editcampaign) | **Put** /campaign/{campaignId} | Edit Campaign
*CampaignAPI* | [**GetAllCampaigns**](docs/CampaignAPI.md#getallcampaigns) | **Get** /campaign | Get All Campaigns
*CampaignAPI* | [**GetCampaignById**](docs/CampaignAPI.md#getcampaignbyid) | **Get** /campaign/{campaignId} | Get Campaign By Id
*ContactAPI* | [**CreateContact**](docs/ContactAPI.md#createcontact) | **Post** /contact | Create a contact
*ContactAPI* | [**DeleteContact**](docs/ContactAPI.md#deletecontact) | **Delete** /contact/{contactId} | Delete Contact
*ContactAPI* | [**GetAllContacts**](docs/ContactAPI.md#getallcontacts) | **Get** /contact | Get All Contacts
*ContactAPI* | [**GetContactById**](docs/ContactAPI.md#getcontactbyid) | **Get** /contact/{contactId} | Get Contact by ID
*ContactAPI* | [**UnsubscribeContact**](docs/ContactAPI.md#unsubscribecontact) | **Post** /contact/unsubscribe/{contactId} | Unsubscribe Contact
*ContactAPI* | [**UpdateContact**](docs/ContactAPI.md#updatecontact) | **Put** /contact/{contactId} | Update Contact
*ListAPI* | [**CreateList**](docs/ListAPI.md#createlist) | **Post** /list | Create List
*ListAPI* | [**DeleteList**](docs/ListAPI.md#deletelist) | **Delete** /list/{listId} | Delete List
*ListAPI* | [**GetAllLists**](docs/ListAPI.md#getalllists) | **Get** /list | Get All Lists
*ListAPI* | [**GetListById**](docs/ListAPI.md#getlistbyid) | **Get** /list/{listId} | Get List
*ListAPI* | [**UpdateList**](docs/ListAPI.md#updatelist) | **Put** /list/{listId} | Update List
*ReportsAPI* | [**GetCampaignReport**](docs/ReportsAPI.md#getcampaignreport) | **Get** /report/campaign/{campaignId} | Get CampaignReport Data
*SenderAPI* | [**CreateSender**](docs/SenderAPI.md#createsender) | **Post** /sender | Create Sender
*SenderAPI* | [**GetAllSenders**](docs/SenderAPI.md#getallsenders) | **Get** /sender | Get All Senders
*TagsAPI* | [**CreateTag**](docs/TagsAPI.md#createtag) | **Post** /tag | Create a Tag
*TagsAPI* | [**DeleteTag**](docs/TagsAPI.md#deletetag) | **Delete** /tag/{tagId} | Delete a Tag
*TagsAPI* | [**GetAllTags**](docs/TagsAPI.md#getalltags) | **Get** /tag | Get All Tags
*TagsAPI* | [**GetTagById**](docs/TagsAPI.md#gettagbyid) | **Get** /tag/{tagId} | Get a Tag by ID
*TagsAPI* | [**UpdateTag**](docs/TagsAPI.md#updatetag) | **Put** /tag/{tagId} | Update a Tag


## Documentation For Models

 - [Campaign](docs/Campaign.md)
 - [CampaignDashboardData](docs/CampaignDashboardData.md)
 - [CampaignRequest](docs/CampaignRequest.md)
 - [Contact](docs/Contact.md)
 - [ContactRequest](docs/ContactRequest.md)
 - [CreateResponse](docs/CreateResponse.md)
 - [DashboardStats](docs/DashboardStats.md)
 - [DeleteCampaign200Response](docs/DeleteCampaign200Response.md)
 - [DeleteRequest](docs/DeleteRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [LastSentCampaignStat](docs/LastSentCampaignStat.md)
 - [ListModel](docs/ListModel.md)
 - [ListRequest](docs/ListRequest.md)
 - [ReportData](docs/ReportData.md)
 - [Response](docs/Response.md)
 - [Sender](docs/Sender.md)
 - [SenderRequest](docs/SenderRequest.md)
 - [SenderResponse](docs/SenderResponse.md)
 - [Tag](docs/Tag.md)
 - [TagRequest](docs/TagRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### apiKeyAuth

- **Type**: API key
- **API key parameter name**: X-Team-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Team-ApiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sendx.ContextAPIKeys,
		map[string]sendx.APIKey{
			"X-Team-ApiKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@sendx.io

