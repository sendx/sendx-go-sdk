# SendX Go SDK

# Introduction 

SendX is an email marketing product. It helps you convert website visitors to customers, send them promotional emails, engage with them using drip sequences and craft custom journeys using powerful but simple automations. 

The SendX API is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.
The SendX Rest API doesn’t support bulk updates. You can work on only one object per request. <br>


## Installation

Install the following dependencies:

```sh
go get github.com/sendx/sendx-go-sdk
```

## Examples

```go
package main

import (
	"context"
	"fmt"
	"os"
	
	sendx "github.com/sendx/sendx-go-sdk"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		sendx.ContextAPIKeys,
		map[string]sendx.APIKey{
			"apiKeyAuth": {Key: "YOUR_API_KEY"},
		},
	)

	contactRequest := *sendx.NewContactRequest() // ContactRequest |
	contactRequest.Email = sendx.PtrString("jane@doe.com")
	contactRequest.FirstName = sendx.PtrString("Jane")
	contactRequest.LastName = sendx.PtrString("Doe")
	contactRequest.Company = sendx.PtrString("Tech Solutions Inc.")
	contactRequest.LastTrackedIp = sendx.PtrString("34.94.159.140")
	contactRequest.CustomFields = &map[string]string{"K2mxBVReqBhbwx9e0ItSea": "VIP", "7o3Tl1aY2yKp2X1aflRjOL": "Special Offer Subscriber"}
	contactRequest.Lists = []string{"1244"}
	contactRequest.Tags = []string{"MKdhTovsTJDetCyrJmRySL"}

	configuration := sendx.NewConfiguration()
	apiClient := sendx.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.CreateContact(ctx).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: Response
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.CreateContact`: %v\n", resp)

}

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

Note, each API key must be added to a map of `map[string]APIKey` where the key is the team api key you get via SendX dashboard and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sendx.ContextAPIKeys,
		map[string]sendx.APIKey{
			"apiKeyAuth": {Key: "Your API Key"},
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

