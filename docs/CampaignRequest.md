# CampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the campaign. | [optional] 
**HtmlCode** | Pointer to **string** | The HTML code of the campaign. | [optional] 
**Subject** | Pointer to **string** | The subject of the campaign. | [optional] 
**Sender** | Pointer to **string** | Sender unique identifier. | [optional] 
**PreviewText** | Pointer to **string** | The preview text shown in email clients. | [optional] 
**ScheduleType** | Pointer to **int32** | The type of scheduling for the campaign &lt;br&gt; 0: Send Now &lt;br&gt; 1: Send Later  | [optional] 
**ScheduleCondition** | Pointer to **string** | The condition for scheduling the campaign. | [optional] 
**TimeCondition** | Pointer to **string** | The specific time to send the campaign. | [optional] 
**Timezone** | Pointer to **string** | The timezone for the campaign scheduling. | [optional] 
**PreferredTimezone** | Pointer to **string** | Preferred timezone for scheduling. | [optional] 
**PreferredTimeCondition** | Pointer to **string** | Specific time preference for sending the campaign. | [optional] 
**SendInContactsTimezone** | Pointer to **bool** | Whether to send emails in each contact&#39;s timezone. | [optional] 
**SmartSend** | Pointer to **bool** | Whether to enable smart send features (e.g., optimizing send time). | [optional] 
**IncludedSegments** | Pointer to **[]string** | List of segment IDs to include. | [optional] 
**IncludedLists** | Pointer to **[]string** | List of list IDs to include. | [optional] 
**IncludedTags** | Pointer to **[]string** | List of tag IDs to include. | [optional] 
**ExcludedSegments** | Pointer to **[]string** | List of segment IDs to exclude. | [optional] 
**ExcludedLists** | Pointer to **[]string** | List of list IDs to exclude. | [optional] 
**ExcludedTags** | Pointer to **[]string** | List of tag IDs to exclude. | [optional] 

## Methods

### NewCampaignRequest

`func NewCampaignRequest() *CampaignRequest`

NewCampaignRequest instantiates a new CampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRequestWithDefaults

`func NewCampaignRequestWithDefaults() *CampaignRequest`

NewCampaignRequestWithDefaults instantiates a new CampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHtmlCode

`func (o *CampaignRequest) GetHtmlCode() string`

GetHtmlCode returns the HtmlCode field if non-nil, zero value otherwise.

### GetHtmlCodeOk

`func (o *CampaignRequest) GetHtmlCodeOk() (*string, bool)`

GetHtmlCodeOk returns a tuple with the HtmlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlCode

`func (o *CampaignRequest) SetHtmlCode(v string)`

SetHtmlCode sets HtmlCode field to given value.

### HasHtmlCode

`func (o *CampaignRequest) HasHtmlCode() bool`

HasHtmlCode returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSender

`func (o *CampaignRequest) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CampaignRequest) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CampaignRequest) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *CampaignRequest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetPreviewText

`func (o *CampaignRequest) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *CampaignRequest) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *CampaignRequest) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *CampaignRequest) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### GetScheduleType

`func (o *CampaignRequest) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CampaignRequest) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CampaignRequest) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *CampaignRequest) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleCondition

`func (o *CampaignRequest) GetScheduleCondition() string`

GetScheduleCondition returns the ScheduleCondition field if non-nil, zero value otherwise.

### GetScheduleConditionOk

`func (o *CampaignRequest) GetScheduleConditionOk() (*string, bool)`

GetScheduleConditionOk returns a tuple with the ScheduleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCondition

`func (o *CampaignRequest) SetScheduleCondition(v string)`

SetScheduleCondition sets ScheduleCondition field to given value.

### HasScheduleCondition

`func (o *CampaignRequest) HasScheduleCondition() bool`

HasScheduleCondition returns a boolean if a field has been set.

### GetTimeCondition

`func (o *CampaignRequest) GetTimeCondition() string`

GetTimeCondition returns the TimeCondition field if non-nil, zero value otherwise.

### GetTimeConditionOk

`func (o *CampaignRequest) GetTimeConditionOk() (*string, bool)`

GetTimeConditionOk returns a tuple with the TimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCondition

`func (o *CampaignRequest) SetTimeCondition(v string)`

SetTimeCondition sets TimeCondition field to given value.

### HasTimeCondition

`func (o *CampaignRequest) HasTimeCondition() bool`

HasTimeCondition returns a boolean if a field has been set.

### GetTimezone

`func (o *CampaignRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CampaignRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CampaignRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CampaignRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPreferredTimezone

`func (o *CampaignRequest) GetPreferredTimezone() string`

GetPreferredTimezone returns the PreferredTimezone field if non-nil, zero value otherwise.

### GetPreferredTimezoneOk

`func (o *CampaignRequest) GetPreferredTimezoneOk() (*string, bool)`

GetPreferredTimezoneOk returns a tuple with the PreferredTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimezone

`func (o *CampaignRequest) SetPreferredTimezone(v string)`

SetPreferredTimezone sets PreferredTimezone field to given value.

### HasPreferredTimezone

`func (o *CampaignRequest) HasPreferredTimezone() bool`

HasPreferredTimezone returns a boolean if a field has been set.

### GetPreferredTimeCondition

`func (o *CampaignRequest) GetPreferredTimeCondition() string`

GetPreferredTimeCondition returns the PreferredTimeCondition field if non-nil, zero value otherwise.

### GetPreferredTimeConditionOk

`func (o *CampaignRequest) GetPreferredTimeConditionOk() (*string, bool)`

GetPreferredTimeConditionOk returns a tuple with the PreferredTimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimeCondition

`func (o *CampaignRequest) SetPreferredTimeCondition(v string)`

SetPreferredTimeCondition sets PreferredTimeCondition field to given value.

### HasPreferredTimeCondition

`func (o *CampaignRequest) HasPreferredTimeCondition() bool`

HasPreferredTimeCondition returns a boolean if a field has been set.

### GetSendInContactsTimezone

`func (o *CampaignRequest) GetSendInContactsTimezone() bool`

GetSendInContactsTimezone returns the SendInContactsTimezone field if non-nil, zero value otherwise.

### GetSendInContactsTimezoneOk

`func (o *CampaignRequest) GetSendInContactsTimezoneOk() (*bool, bool)`

GetSendInContactsTimezoneOk returns a tuple with the SendInContactsTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInContactsTimezone

`func (o *CampaignRequest) SetSendInContactsTimezone(v bool)`

SetSendInContactsTimezone sets SendInContactsTimezone field to given value.

### HasSendInContactsTimezone

`func (o *CampaignRequest) HasSendInContactsTimezone() bool`

HasSendInContactsTimezone returns a boolean if a field has been set.

### GetSmartSend

`func (o *CampaignRequest) GetSmartSend() bool`

GetSmartSend returns the SmartSend field if non-nil, zero value otherwise.

### GetSmartSendOk

`func (o *CampaignRequest) GetSmartSendOk() (*bool, bool)`

GetSmartSendOk returns a tuple with the SmartSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSend

`func (o *CampaignRequest) SetSmartSend(v bool)`

SetSmartSend sets SmartSend field to given value.

### HasSmartSend

`func (o *CampaignRequest) HasSmartSend() bool`

HasSmartSend returns a boolean if a field has been set.

### GetIncludedSegments

`func (o *CampaignRequest) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *CampaignRequest) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *CampaignRequest) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *CampaignRequest) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetIncludedLists

`func (o *CampaignRequest) GetIncludedLists() []string`

GetIncludedLists returns the IncludedLists field if non-nil, zero value otherwise.

### GetIncludedListsOk

`func (o *CampaignRequest) GetIncludedListsOk() (*[]string, bool)`

GetIncludedListsOk returns a tuple with the IncludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLists

`func (o *CampaignRequest) SetIncludedLists(v []string)`

SetIncludedLists sets IncludedLists field to given value.

### HasIncludedLists

`func (o *CampaignRequest) HasIncludedLists() bool`

HasIncludedLists returns a boolean if a field has been set.

### GetIncludedTags

`func (o *CampaignRequest) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *CampaignRequest) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *CampaignRequest) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *CampaignRequest) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *CampaignRequest) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *CampaignRequest) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *CampaignRequest) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *CampaignRequest) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetExcludedLists

`func (o *CampaignRequest) GetExcludedLists() []string`

GetExcludedLists returns the ExcludedLists field if non-nil, zero value otherwise.

### GetExcludedListsOk

`func (o *CampaignRequest) GetExcludedListsOk() (*[]string, bool)`

GetExcludedListsOk returns a tuple with the ExcludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedLists

`func (o *CampaignRequest) SetExcludedLists(v []string)`

SetExcludedLists sets ExcludedLists field to given value.

### HasExcludedLists

`func (o *CampaignRequest) HasExcludedLists() bool`

HasExcludedLists returns a boolean if a field has been set.

### GetExcludedTags

`func (o *CampaignRequest) GetExcludedTags() []string`

GetExcludedTags returns the ExcludedTags field if non-nil, zero value otherwise.

### GetExcludedTagsOk

`func (o *CampaignRequest) GetExcludedTagsOk() (*[]string, bool)`

GetExcludedTagsOk returns a tuple with the ExcludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTags

`func (o *CampaignRequest) SetExcludedTags(v []string)`

SetExcludedTags sets ExcludedTags field to given value.

### HasExcludedTags

`func (o *CampaignRequest) HasExcludedTags() bool`

HasExcludedTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


