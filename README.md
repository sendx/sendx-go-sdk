# Go API client for sendx

## 🚀 Introduction

The SendX API is organized around REST principles. Our API has predictable resource-oriented URLs, accepts JSON-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.

**Key Features:**
- 🔒 **Security**: Team-based authentication with optional member-level access
- 🎯 **Resource-Oriented**: RESTful design with clear resource boundaries
- 📊 **Rich Data Models**: Three-layer model system (Input/Output/Internal)
- 🔗 **Relationships**: Automatic prefix handling for resource relationships
- 📈 **Scalable**: Built for high-volume email marketing operations

## 🏗️ Architecture Overview

SendX uses a three-layer model architecture:

1. **Input Models** (`RestE*`): For API requests
2. **Output Models** (`RestR*`): For API responses with prefixed IDs
3. **Internal Models**: Core business logic (not exposed in API)

## 🔐 Security & Authentication

SendX uses API key authentication:

### Team API Key
```http
X-Team-ApiKey: YOUR_TEAM_API_KEY
```
- **Required for all requests**
- Team-level access to resources
- Available in SendX Settings → Team API Key

## 🆔 Encrypted ID System

SendX uses encrypted IDs for security and better developer experience:

- **Internal IDs**: Sequential integers (not exposed)
- **Encrypted IDs**: 22-character alphanumeric strings
- **Prefixed IDs**: Resource-type prefixes in API responses (`contact_<22-char-id>`)

### ID Format

**All resource IDs follow this pattern:**
```
<resource_prefix>_<22_character_alphanumeric_string>
```

**Example:**
```json
{
  "id": "contact_BnKjkbBBS500CoBCP0oChQ",
  "lists": ["list_OcuxJHdiAvujmwQVJfd3ss", "list_0tOFLp5RgV7s3LNiHrjGYs"],
  "tags": ["tag_UhsDkjL772Qbj5lWtT62VK", "tag_fL7t9lsnZ9swvx2HrtQ9wM"]
}
```

## 📚 Resource Prefixes

| Resource | Prefix | Example |
|----------|--------|---------|
| Contact | `contact_` | `contact_BnKjkbBBS500CoBCP0oChQ` |
| Campaign | `campaign_` | `campaign_LUE9BTxmksSmqHWbh96zsn` |
| List | `list_` | `list_OcuxJHdiAvujmwQVJfd3ss` |
| Tag | `tag_` | `tag_UhsDkjL772Qbj5lWtT62VK` |
| Sender | `sender_` | `sender_4vK3WFhMgvOwUNyaL4QxCD` |
| Template | `template_` | `template_f3lJvTEhSjKGVb5Lwc5SWS` |
| Custom Field | `field_` | `field_MnuqBAG2NPLm7PZMWbjQxt` |
| Webhook | `webhook_` | `webhook_9l154iiXlZoPo7vngmamee` |
| Post | `post_` | `post_XyZ123aBc456DeF789GhI` |
| Post Category | `post_category_` | `post_category_YzS1wOU20yw87UUHKxMzwn` |
| Post Tag | `post_tag_` | `post_tag_123XyZ456AbC` |
| Member | `member_` | `member_JkL012MnO345PqR678` |

## 🎯 Best Practices

### Error Handling
- **Always check status codes**: 2xx = success, 4xx = client error, 5xx = server error
- **Read error messages**: Descriptive messages help debug issues
- **Handle rate limits**: Respect API rate limits for optimal performance

### Data Validation
- **Email format**: Must be valid email addresses
- **Required fields**: Check documentation for mandatory fields
- **Field lengths**: Respect maximum length constraints

### Performance
- **Pagination**: Use offset/limit for large datasets
- **Batch operations**: Process multiple items when supported
- **Caching**: Cache responses when appropriate

## 🛠️ SDKs & Integration

Official SDKs available for:
- [Golang](https://github.com/sendx/sendx-go-sdk)
- [Python](https://github.com/sendx/sendx-python-sdk)
- [Ruby](https://github.com/sendx/sendx-ruby-sdk)
- [Java](https://github.com/sendx/sendx-java-sdk)
- [PHP](https://github.com/sendx/sendx-php-sdk)
- [JavaScript](https://github.com/sendx/sendx-javascript-sdk)

## 📞 Support

Need help? Contact us:
- 💬 **Website Chat**: Available on sendx.io
- 📧 **Email**: hello@sendx.io
- 📚 **Documentation**: Full guides at help.sendx.io

---

**API Endpoint:** `https://api.sendx.io/api/v1/rest`

[<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-44b198b0-5219-4619-a01f-cfc24d573885?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-44b198b0-5219-4619-a01f-cfc24d573885%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e)


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

Note, enum values are always validated and all unused variables are silently ignored.


## Documentation for API Endpoints

All URIs are relative to *https://api.sendx.io/api/v1/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignAPI* | [**CreateCampaign**](docs/CampaignAPI.md#createcampaign) | **Post** /campaign | Create campaign
*CampaignAPI* | [**DeleteCampaign**](docs/CampaignAPI.md#deletecampaign) | **Delete** /campaign/{identifier} | Delete campaign
*CampaignAPI* | [**GetAllCampaigns**](docs/CampaignAPI.md#getallcampaigns) | **Get** /campaign | Get all campaigns
*CampaignAPI* | [**GetCampaign**](docs/CampaignAPI.md#getcampaign) | **Get** /campaign/{identifier} | Get campaign by ID
*ContactAPI* | [**CreateContact**](docs/ContactAPI.md#createcontact) | **Post** /contact | Create a new contact
*ContactAPI* | [**DeleteContact**](docs/ContactAPI.md#deletecontact) | **Delete** /contact/{identifier} | Delete contact
*ContactAPI* | [**GetAllContacts**](docs/ContactAPI.md#getallcontacts) | **Get** /contact | Get all contacts
*ContactAPI* | [**GetContact**](docs/ContactAPI.md#getcontact) | **Get** /contact/{identifier} | Get contact by ID
*ContactAPI* | [**UnsubscribeContact**](docs/ContactAPI.md#unsubscribecontact) | **Post** /contact/unsubscribe/{identifier} | Unsubscribe contact
*ContactAPI* | [**UpdateContact**](docs/ContactAPI.md#updatecontact) | **Put** /contact/{identifier} | Update contact
*CustomFieldAPI* | [**CreateCustomField**](docs/CustomFieldAPI.md#createcustomfield) | **Post** /customfield | Create custom field
*CustomFieldAPI* | [**DeleteCustomField**](docs/CustomFieldAPI.md#deletecustomfield) | **Delete** /customfield/{identifier} | Delete custom field
*CustomFieldAPI* | [**GetAllCustomFields**](docs/CustomFieldAPI.md#getallcustomfields) | **Get** /customfield | Get all custom fields
*CustomFieldAPI* | [**GetCustomField**](docs/CustomFieldAPI.md#getcustomfield) | **Get** /customfield/{identifier} | Get custom field by ID
*CustomFieldAPI* | [**UpdateCustomField**](docs/CustomFieldAPI.md#updatecustomfield) | **Put** /customfield/{identifier} | Update custom field
*EmailSendingAPI* | [**SendEmail**](docs/EmailSendingAPI.md#sendemail) | **Post** /send/email | Send transactional email
*EmailSendingAPI* | [**SendEmailWithTemplate**](docs/EmailSendingAPI.md#sendemailwithtemplate) | **Post** /send/template | Send email using template
*EventAPI* | [**EventsCustomPostbackGet**](docs/EventAPI.md#eventscustompostbackget) | **Get** /events/custom/postback | Custom Event Postback URL
*EventAPI* | [**EventsRevenuePostbackGet**](docs/EventAPI.md#eventsrevenuepostbackget) | **Get** /events/revenue/postback | Revenue Event Postback URL
*EventsAPI* | [**TrackCustomEvent**](docs/EventsAPI.md#trackcustomevent) | **Post** /events/custom | Track custom event
*EventsAPI* | [**TrackRevenueEvent**](docs/EventsAPI.md#trackrevenueevent) | **Post** /events/revenue | Track revenue event
*ListAPI* | [**CreateList**](docs/ListAPI.md#createlist) | **Post** /list | Create list
*ListAPI* | [**DeleteList**](docs/ListAPI.md#deletelist) | **Delete** /list/{identifier} | Delete list
*ListAPI* | [**GetAllLists**](docs/ListAPI.md#getalllists) | **Get** /list | Get all lists
*ListAPI* | [**GetList**](docs/ListAPI.md#getlist) | **Get** /list/{identifier} | Get list by ID
*ListAPI* | [**UpdateList**](docs/ListAPI.md#updatelist) | **Put** /list/{identifier} | Update list
*PostAPI* | [**CreatePost**](docs/PostAPI.md#createpost) | **Post** /post | Create blog post
*PostAPI* | [**DeletePost**](docs/PostAPI.md#deletepost) | **Delete** /post/{identifier} | Delete post
*PostAPI* | [**GetAllPosts**](docs/PostAPI.md#getallposts) | **Get** /post | Get all posts
*PostAPI* | [**GetPost**](docs/PostAPI.md#getpost) | **Get** /post/{identifier} | Get post by ID
*PostAPI* | [**UpdatePost**](docs/PostAPI.md#updatepost) | **Put** /post/{identifier} | Update post
*PostCategoryAPI* | [**CreatePostCategory**](docs/PostCategoryAPI.md#createpostcategory) | **Post** /post/category | Create post category
*PostCategoryAPI* | [**DeletePostCategory**](docs/PostCategoryAPI.md#deletepostcategory) | **Delete** /post/category/{identifier} | Delete post category
*PostCategoryAPI* | [**GetAllPostCategories**](docs/PostCategoryAPI.md#getallpostcategories) | **Get** /post/category | Get all post categories
*PostCategoryAPI* | [**GetPostCategory**](docs/PostCategoryAPI.md#getpostcategory) | **Get** /post/category/{identifier} | Get post category by ID
*PostCategoryAPI* | [**UpdatePostCategory**](docs/PostCategoryAPI.md#updatepostcategory) | **Put** /post/category/{identifier} | Update post category
*PostTagAPI* | [**CreatePostTag**](docs/PostTagAPI.md#createposttag) | **Post** /post/tag | Create post tag
*PostTagAPI* | [**DeletePostTag**](docs/PostTagAPI.md#deleteposttag) | **Delete** /post/tag/{identifier} | Delete post tag
*PostTagAPI* | [**GetAllPostTags**](docs/PostTagAPI.md#getallposttags) | **Get** /post/tag | Get all post tags
*PostTagAPI* | [**GetPostTag**](docs/PostTagAPI.md#getposttag) | **Get** /post/tag/{identifier} | Get post tag by ID
*PostTagAPI* | [**UpdatePostTag**](docs/PostTagAPI.md#updateposttag) | **Put** /post/tag/{identifier} | Update post tag
*ReportAPI* | [**GetCampaignReport**](docs/ReportAPI.md#getcampaignreport) | **Get** /report/campaign/{identifier} | Get campaign report
*SenderAPI* | [**CreateSender**](docs/SenderAPI.md#createsender) | **Post** /sender | Create sender
*SenderAPI* | [**GetAllSenders**](docs/SenderAPI.md#getallsenders) | **Get** /sender | Get all senders
*TagAPI* | [**CreateTag**](docs/TagAPI.md#createtag) | **Post** /tag | Create tag
*TagAPI* | [**DeleteTag**](docs/TagAPI.md#deletetag) | **Delete** /tag/{identifier} | Delete tag
*TagAPI* | [**GetAllTags**](docs/TagAPI.md#getalltags) | **Get** /tag | Get all tags
*TagAPI* | [**GetTag**](docs/TagAPI.md#gettag) | **Get** /tag/{identifier} | Get tag by ID
*TagAPI* | [**UpdateTag**](docs/TagAPI.md#updatetag) | **Put** /tag/{identifier} | Update tag
*TeamMemberAPI* | [**GetAllTeamMembers**](docs/TeamMemberAPI.md#getallteammembers) | **Get** /team/member | Get all team members
*TeamMemberAPI* | [**GetTeamMember**](docs/TeamMemberAPI.md#getteammember) | **Get** /team/member/{identifier} | Get a team member by ID
*TemplateAPI* | [**CreateEmailTemplate**](docs/TemplateAPI.md#createemailtemplate) | **Post** /template/email | Create email template
*TemplateAPI* | [**DeleteEmailTemplate**](docs/TemplateAPI.md#deleteemailtemplate) | **Delete** /template/email/{identifier} | Delete template
*TemplateAPI* | [**GetAllEmailTemplates**](docs/TemplateAPI.md#getallemailtemplates) | **Get** /template/email | Get all templates
*TemplateAPI* | [**GetEmailTemplate**](docs/TemplateAPI.md#getemailtemplate) | **Get** /template/email/{identifier} | Get template by ID
*TemplateAPI* | [**UpdateEmailTemplate**](docs/TemplateAPI.md#updateemailtemplate) | **Put** /template/email/{identifier} | Update template
*TrackingAPI* | [**IdentifyContact**](docs/TrackingAPI.md#identifycontact) | **Post** /contact/identify | Identify contact
*TrackingAPI* | [**TrackContact**](docs/TrackingAPI.md#trackcontact) | **Post** /contact/track | Track contact
*WebhookAPI* | [**CreateWebhook**](docs/WebhookAPI.md#createwebhook) | **Post** /webhook | Create webhook
*WebhookAPI* | [**DeleteWebhook**](docs/WebhookAPI.md#deletewebhook) | **Delete** /webhook/{identifier} | Delete webhook
*WebhookAPI* | [**GetAllWebhooks**](docs/WebhookAPI.md#getallwebhooks) | **Get** /webhook | Get all webhooks
*WebhookAPI* | [**GetWebhook**](docs/WebhookAPI.md#getwebhook) | **Get** /webhook/{identifier} | Get webhook by ID
*WebhookAPI* | [**UpdateWebhook**](docs/WebhookAPI.md#updatewebhook) | **Put** /webhook/{identifier} | Update webhook


## Documentation For Models

 - [CustomEventRequest](docs/CustomEventRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EventResponse](docs/EventResponse.md)
 - [EventsRevenuePostbackGet200Response](docs/EventsRevenuePostbackGet200Response.md)
 - [EventsRevenuePostbackGet400Response](docs/EventsRevenuePostbackGet400Response.md)
 - [EventsRevenuePostbackGet500Response](docs/EventsRevenuePostbackGet500Response.md)
 - [IdentifyRequest](docs/IdentifyRequest.md)
 - [IdentifyResponse](docs/IdentifyResponse.md)
 - [LinkStat](docs/LinkStat.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [PostbackResponse](docs/PostbackResponse.md)
 - [RestECampaign](docs/RestECampaign.md)
 - [RestEContact](docs/RestEContact.md)
 - [RestECustomField](docs/RestECustomField.md)
 - [RestEList](docs/RestEList.md)
 - [RestEPost](docs/RestEPost.md)
 - [RestEPostCategory](docs/RestEPostCategory.md)
 - [RestEPostTag](docs/RestEPostTag.md)
 - [RestESender](docs/RestESender.md)
 - [RestETag](docs/RestETag.md)
 - [RestETemplate](docs/RestETemplate.md)
 - [RestEWebhook](docs/RestEWebhook.md)
 - [RestRCampaign](docs/RestRCampaign.md)
 - [RestRContact](docs/RestRContact.md)
 - [RestRCustomField](docs/RestRCustomField.md)
 - [RestRList](docs/RestRList.md)
 - [RestRMember](docs/RestRMember.md)
 - [RestRPost](docs/RestRPost.md)
 - [RestRPostCategory](docs/RestRPostCategory.md)
 - [RestRPostTag](docs/RestRPostTag.md)
 - [RestRSender](docs/RestRSender.md)
 - [RestRTag](docs/RestRTag.md)
 - [RestRTemplate](docs/RestRTemplate.md)
 - [RestRWebhook](docs/RestRWebhook.md)
 - [RestReportData](docs/RestReportData.md)
 - [RevenueEventRequest](docs/RevenueEventRequest.md)
 - [TemplateEmailMessage](docs/TemplateEmailMessage.md)
 - [TrackRequest](docs/TrackRequest.md)
 - [TrackResponse](docs/TrackResponse.md)
 - [WebhookObject](docs/WebhookObject.md)
 - [XAttachment](docs/XAttachment.md)
 - [XEmailMessage](docs/XEmailMessage.md)
 - [XEmailResponse](docs/XEmailResponse.md)
 - [XFrom](docs/XFrom.md)
 - [XReplyTo](docs/XReplyTo.md)
 - [XTo](docs/XTo.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### TeamApiKey

- **Type**: API key
- **API key parameter name**: X-Team-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: TeamApiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sendx.ContextAPIKeys,
		map[string]sendx.APIKey{
			"TeamApiKey": {Key: "API_KEY_STRING"},
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

hello@sendx.io

