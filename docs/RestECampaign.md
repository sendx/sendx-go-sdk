# RestECampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Campaign name for internal organization | 
**Subject** | **string** | Email subject line with personalization support.  **Features:** - Supports template variables ({{contact.firstName}}) - Emoji support for better engagement - A/B testing variations supported  | 
**Sender** | **string** | Sender identifier.  **Note:** Sender must be verified before use  | 
**HtmlCode** | **string** | HTML content of the email campaign | 
**PreviewText** | Pointer to **string** | Preview text shown in email clients | [optional] 
**PlainText** | Pointer to **string** | Plain text version for better deliverability | [optional] 
**ScheduleType** | Pointer to **int32** | Campaign scheduling type.  **Values:** - &#x60;0&#x60; - Schedule for specific date/time - &#x60;1&#x60; - Send immediately  | [optional] 
**ScheduleCondition** | Pointer to **string** | datetime for scheduled campaigns (required if scheduleType&#x3D;0) | [optional] 
**TimeCondition** | Pointer to **string** | Time condition for scheduled campaigns in HH:MM PM/AM format | [optional] 
**Timezone** | Pointer to **string** | Timezone for scheduled campaigns (IANA format) | [optional] 
**PreferredTimezone** | Pointer to **string** | Preferred timezone for smart send optimization (required for smartSend and sendInContactsTimezone) | [optional] 
**PreferredTimeCondition** | Pointer to **string** | Preferred time optimization setting (required for smartSend and sendInContactsTimezone) | [optional] 
**SendInContactsTimezone** | Pointer to **bool** | Send at specified time in each contact&#39;s timezone | [optional] 
**SmartSend** | Pointer to **bool** | Enable AI-powered send time optimization | [optional] 
**IncludedSegments** | Pointer to **[]string** | Segment IDs to include | [optional] 
**IncludedLists** | Pointer to **[]string** | List IDs to include | [optional] 
**IncludedTags** | Pointer to **[]string** | Tag IDs to include | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment IDs to exclude | [optional] 
**ExcludedLists** | Pointer to **[]string** | List IDs to exclude | [optional] 
**ExcludedTags** | Pointer to **[]string** | Tag IDs to exclude (prefix automatically stripped) | [optional] 

## Methods

### NewRestECampaign

`func NewRestECampaign(name string, subject string, sender string, htmlCode string, ) *RestECampaign`

NewRestECampaign instantiates a new RestECampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestECampaignWithDefaults

`func NewRestECampaignWithDefaults() *RestECampaign`

NewRestECampaignWithDefaults instantiates a new RestECampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestECampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestECampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestECampaign) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *RestECampaign) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RestECampaign) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RestECampaign) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSender

`func (o *RestECampaign) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *RestECampaign) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *RestECampaign) SetSender(v string)`

SetSender sets Sender field to given value.


### GetHtmlCode

`func (o *RestECampaign) GetHtmlCode() string`

GetHtmlCode returns the HtmlCode field if non-nil, zero value otherwise.

### GetHtmlCodeOk

`func (o *RestECampaign) GetHtmlCodeOk() (*string, bool)`

GetHtmlCodeOk returns a tuple with the HtmlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlCode

`func (o *RestECampaign) SetHtmlCode(v string)`

SetHtmlCode sets HtmlCode field to given value.


### GetPreviewText

`func (o *RestECampaign) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *RestECampaign) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *RestECampaign) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *RestECampaign) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### GetPlainText

`func (o *RestECampaign) GetPlainText() string`

GetPlainText returns the PlainText field if non-nil, zero value otherwise.

### GetPlainTextOk

`func (o *RestECampaign) GetPlainTextOk() (*string, bool)`

GetPlainTextOk returns a tuple with the PlainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainText

`func (o *RestECampaign) SetPlainText(v string)`

SetPlainText sets PlainText field to given value.

### HasPlainText

`func (o *RestECampaign) HasPlainText() bool`

HasPlainText returns a boolean if a field has been set.

### GetScheduleType

`func (o *RestECampaign) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *RestECampaign) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *RestECampaign) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *RestECampaign) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleCondition

`func (o *RestECampaign) GetScheduleCondition() string`

GetScheduleCondition returns the ScheduleCondition field if non-nil, zero value otherwise.

### GetScheduleConditionOk

`func (o *RestECampaign) GetScheduleConditionOk() (*string, bool)`

GetScheduleConditionOk returns a tuple with the ScheduleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCondition

`func (o *RestECampaign) SetScheduleCondition(v string)`

SetScheduleCondition sets ScheduleCondition field to given value.

### HasScheduleCondition

`func (o *RestECampaign) HasScheduleCondition() bool`

HasScheduleCondition returns a boolean if a field has been set.

### GetTimeCondition

`func (o *RestECampaign) GetTimeCondition() string`

GetTimeCondition returns the TimeCondition field if non-nil, zero value otherwise.

### GetTimeConditionOk

`func (o *RestECampaign) GetTimeConditionOk() (*string, bool)`

GetTimeConditionOk returns a tuple with the TimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCondition

`func (o *RestECampaign) SetTimeCondition(v string)`

SetTimeCondition sets TimeCondition field to given value.

### HasTimeCondition

`func (o *RestECampaign) HasTimeCondition() bool`

HasTimeCondition returns a boolean if a field has been set.

### GetTimezone

`func (o *RestECampaign) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RestECampaign) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RestECampaign) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RestECampaign) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPreferredTimezone

`func (o *RestECampaign) GetPreferredTimezone() string`

GetPreferredTimezone returns the PreferredTimezone field if non-nil, zero value otherwise.

### GetPreferredTimezoneOk

`func (o *RestECampaign) GetPreferredTimezoneOk() (*string, bool)`

GetPreferredTimezoneOk returns a tuple with the PreferredTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimezone

`func (o *RestECampaign) SetPreferredTimezone(v string)`

SetPreferredTimezone sets PreferredTimezone field to given value.

### HasPreferredTimezone

`func (o *RestECampaign) HasPreferredTimezone() bool`

HasPreferredTimezone returns a boolean if a field has been set.

### GetPreferredTimeCondition

`func (o *RestECampaign) GetPreferredTimeCondition() string`

GetPreferredTimeCondition returns the PreferredTimeCondition field if non-nil, zero value otherwise.

### GetPreferredTimeConditionOk

`func (o *RestECampaign) GetPreferredTimeConditionOk() (*string, bool)`

GetPreferredTimeConditionOk returns a tuple with the PreferredTimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimeCondition

`func (o *RestECampaign) SetPreferredTimeCondition(v string)`

SetPreferredTimeCondition sets PreferredTimeCondition field to given value.

### HasPreferredTimeCondition

`func (o *RestECampaign) HasPreferredTimeCondition() bool`

HasPreferredTimeCondition returns a boolean if a field has been set.

### GetSendInContactsTimezone

`func (o *RestECampaign) GetSendInContactsTimezone() bool`

GetSendInContactsTimezone returns the SendInContactsTimezone field if non-nil, zero value otherwise.

### GetSendInContactsTimezoneOk

`func (o *RestECampaign) GetSendInContactsTimezoneOk() (*bool, bool)`

GetSendInContactsTimezoneOk returns a tuple with the SendInContactsTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInContactsTimezone

`func (o *RestECampaign) SetSendInContactsTimezone(v bool)`

SetSendInContactsTimezone sets SendInContactsTimezone field to given value.

### HasSendInContactsTimezone

`func (o *RestECampaign) HasSendInContactsTimezone() bool`

HasSendInContactsTimezone returns a boolean if a field has been set.

### GetSmartSend

`func (o *RestECampaign) GetSmartSend() bool`

GetSmartSend returns the SmartSend field if non-nil, zero value otherwise.

### GetSmartSendOk

`func (o *RestECampaign) GetSmartSendOk() (*bool, bool)`

GetSmartSendOk returns a tuple with the SmartSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSend

`func (o *RestECampaign) SetSmartSend(v bool)`

SetSmartSend sets SmartSend field to given value.

### HasSmartSend

`func (o *RestECampaign) HasSmartSend() bool`

HasSmartSend returns a boolean if a field has been set.

### GetIncludedSegments

`func (o *RestECampaign) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *RestECampaign) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *RestECampaign) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *RestECampaign) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetIncludedLists

`func (o *RestECampaign) GetIncludedLists() []string`

GetIncludedLists returns the IncludedLists field if non-nil, zero value otherwise.

### GetIncludedListsOk

`func (o *RestECampaign) GetIncludedListsOk() (*[]string, bool)`

GetIncludedListsOk returns a tuple with the IncludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLists

`func (o *RestECampaign) SetIncludedLists(v []string)`

SetIncludedLists sets IncludedLists field to given value.

### HasIncludedLists

`func (o *RestECampaign) HasIncludedLists() bool`

HasIncludedLists returns a boolean if a field has been set.

### GetIncludedTags

`func (o *RestECampaign) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *RestECampaign) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *RestECampaign) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *RestECampaign) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *RestECampaign) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *RestECampaign) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *RestECampaign) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *RestECampaign) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetExcludedLists

`func (o *RestECampaign) GetExcludedLists() []string`

GetExcludedLists returns the ExcludedLists field if non-nil, zero value otherwise.

### GetExcludedListsOk

`func (o *RestECampaign) GetExcludedListsOk() (*[]string, bool)`

GetExcludedListsOk returns a tuple with the ExcludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedLists

`func (o *RestECampaign) SetExcludedLists(v []string)`

SetExcludedLists sets ExcludedLists field to given value.

### HasExcludedLists

`func (o *RestECampaign) HasExcludedLists() bool`

HasExcludedLists returns a boolean if a field has been set.

### GetExcludedTags

`func (o *RestECampaign) GetExcludedTags() []string`

GetExcludedTags returns the ExcludedTags field if non-nil, zero value otherwise.

### GetExcludedTagsOk

`func (o *RestECampaign) GetExcludedTagsOk() (*[]string, bool)`

GetExcludedTagsOk returns a tuple with the ExcludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTags

`func (o *RestECampaign) SetExcludedTags(v []string)`

SetExcludedTags sets ExcludedTags field to given value.

### HasExcludedTags

`func (o *RestECampaign) HasExcludedTags() bool`

HasExcludedTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


