# RestRCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Encrypted ID of the campaign | [optional] 
**Name** | **string** | Internal campaign name | 
**Subject** | **string** | Email subject line | 
**Preheader** | Pointer to **string** | Email preview text | [optional] 
**Sender** | **string** | Sender id | 
**HtmlContent** | Pointer to **string** | HTML email content | [optional] 
**TextContent** | Pointer to **string** | Plain text email content | [optional] 
**ScheduleType** | **int32** | Campaign scheduling type.  **Values:** - &#x60;0&#x60; - Schedule later - &#x60;1&#x60; - Send Now  | [default to 0]
**ScheduleCondition** | **string** | datetime for scheduled campaigns (required if scheduleType&#x3D;1) | 
**Status** | Pointer to **int32** | Campaign status.  **Values:** - &#x60;0&#x60; - Draft - &#x60;1&#x60; - Scheduled - &#x60;2&#x60; - Sending - &#x60;3&#x60; - Sent - &#x60;4&#x60; - Quarantined - &#x60;5&#x60; - Evaluating - &#x60;6&#x60; - Evaluation Failed - &#x60;7&#x60; - Warming Up  | [optional] 
**TimeCondition** | Pointer to **string** | Time-related condition for the campaign | [optional] 
**Timezone** | Pointer to **string** | Campaign timezone | [optional] 
**SmartSend** | Pointer to **bool** | Timezone for the scheduled send | [optional] 
**SendInContactsTimezone** | Pointer to **bool** | Send at specified time in each contact&#39;s timezone | [optional] 
**PreferredTimeCondition** | Pointer to **string** | Preferred time condition, in case of smartSend and sendInContactTimeZone | [optional] 
**PreferredTimezone** | Pointer to **string** | Preferred timezone for smart send optimization | [optional] 
**Strategy** | Pointer to **string** | Campaign delivery strategy | [optional] 
**IncludedSegments** | Pointer to **[]string** | Included segment IDs | [optional] 
**IncludedLists** | **[]string** | Included list IDs with prefix | 
**IncludedTags** | Pointer to **[]string** | Included tag IDs with prefix | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Excluded segment IDs | [optional] 
**ExcludedLists** | **[]string** | Excluded list IDs with prefix | 
**ExcludedTags** | Pointer to **[]string** | Excluded tag IDs with prefix | [optional] 
**IsArchived** | Pointer to **bool** | Whether the campaign is archived | [optional] 
**CampaignScreenshotUrl** | Pointer to **string** | URL to the campaign screenshot | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRestRCampaign

`func NewRestRCampaign(name string, subject string, sender string, scheduleType int32, scheduleCondition string, includedLists []string, excludedLists []string, ) *RestRCampaign`

NewRestRCampaign instantiates a new RestRCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRCampaignWithDefaults

`func NewRestRCampaignWithDefaults() *RestRCampaign`

NewRestRCampaignWithDefaults instantiates a new RestRCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *RestRCampaign) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RestRCampaign) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RestRCampaign) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetPreheader

`func (o *RestRCampaign) GetPreheader() string`

GetPreheader returns the Preheader field if non-nil, zero value otherwise.

### GetPreheaderOk

`func (o *RestRCampaign) GetPreheaderOk() (*string, bool)`

GetPreheaderOk returns a tuple with the Preheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreheader

`func (o *RestRCampaign) SetPreheader(v string)`

SetPreheader sets Preheader field to given value.

### HasPreheader

`func (o *RestRCampaign) HasPreheader() bool`

HasPreheader returns a boolean if a field has been set.

### GetSender

`func (o *RestRCampaign) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *RestRCampaign) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *RestRCampaign) SetSender(v string)`

SetSender sets Sender field to given value.


### GetHtmlContent

`func (o *RestRCampaign) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *RestRCampaign) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *RestRCampaign) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *RestRCampaign) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetTextContent

`func (o *RestRCampaign) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *RestRCampaign) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *RestRCampaign) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *RestRCampaign) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetScheduleType

`func (o *RestRCampaign) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *RestRCampaign) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *RestRCampaign) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetScheduleCondition

`func (o *RestRCampaign) GetScheduleCondition() string`

GetScheduleCondition returns the ScheduleCondition field if non-nil, zero value otherwise.

### GetScheduleConditionOk

`func (o *RestRCampaign) GetScheduleConditionOk() (*string, bool)`

GetScheduleConditionOk returns a tuple with the ScheduleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCondition

`func (o *RestRCampaign) SetScheduleCondition(v string)`

SetScheduleCondition sets ScheduleCondition field to given value.


### GetStatus

`func (o *RestRCampaign) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestRCampaign) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestRCampaign) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestRCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeCondition

`func (o *RestRCampaign) GetTimeCondition() string`

GetTimeCondition returns the TimeCondition field if non-nil, zero value otherwise.

### GetTimeConditionOk

`func (o *RestRCampaign) GetTimeConditionOk() (*string, bool)`

GetTimeConditionOk returns a tuple with the TimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCondition

`func (o *RestRCampaign) SetTimeCondition(v string)`

SetTimeCondition sets TimeCondition field to given value.

### HasTimeCondition

`func (o *RestRCampaign) HasTimeCondition() bool`

HasTimeCondition returns a boolean if a field has been set.

### GetTimezone

`func (o *RestRCampaign) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RestRCampaign) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RestRCampaign) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RestRCampaign) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSmartSend

`func (o *RestRCampaign) GetSmartSend() bool`

GetSmartSend returns the SmartSend field if non-nil, zero value otherwise.

### GetSmartSendOk

`func (o *RestRCampaign) GetSmartSendOk() (*bool, bool)`

GetSmartSendOk returns a tuple with the SmartSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSend

`func (o *RestRCampaign) SetSmartSend(v bool)`

SetSmartSend sets SmartSend field to given value.

### HasSmartSend

`func (o *RestRCampaign) HasSmartSend() bool`

HasSmartSend returns a boolean if a field has been set.

### GetSendInContactsTimezone

`func (o *RestRCampaign) GetSendInContactsTimezone() bool`

GetSendInContactsTimezone returns the SendInContactsTimezone field if non-nil, zero value otherwise.

### GetSendInContactsTimezoneOk

`func (o *RestRCampaign) GetSendInContactsTimezoneOk() (*bool, bool)`

GetSendInContactsTimezoneOk returns a tuple with the SendInContactsTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInContactsTimezone

`func (o *RestRCampaign) SetSendInContactsTimezone(v bool)`

SetSendInContactsTimezone sets SendInContactsTimezone field to given value.

### HasSendInContactsTimezone

`func (o *RestRCampaign) HasSendInContactsTimezone() bool`

HasSendInContactsTimezone returns a boolean if a field has been set.

### GetPreferredTimeCondition

`func (o *RestRCampaign) GetPreferredTimeCondition() string`

GetPreferredTimeCondition returns the PreferredTimeCondition field if non-nil, zero value otherwise.

### GetPreferredTimeConditionOk

`func (o *RestRCampaign) GetPreferredTimeConditionOk() (*string, bool)`

GetPreferredTimeConditionOk returns a tuple with the PreferredTimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimeCondition

`func (o *RestRCampaign) SetPreferredTimeCondition(v string)`

SetPreferredTimeCondition sets PreferredTimeCondition field to given value.

### HasPreferredTimeCondition

`func (o *RestRCampaign) HasPreferredTimeCondition() bool`

HasPreferredTimeCondition returns a boolean if a field has been set.

### GetPreferredTimezone

`func (o *RestRCampaign) GetPreferredTimezone() string`

GetPreferredTimezone returns the PreferredTimezone field if non-nil, zero value otherwise.

### GetPreferredTimezoneOk

`func (o *RestRCampaign) GetPreferredTimezoneOk() (*string, bool)`

GetPreferredTimezoneOk returns a tuple with the PreferredTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimezone

`func (o *RestRCampaign) SetPreferredTimezone(v string)`

SetPreferredTimezone sets PreferredTimezone field to given value.

### HasPreferredTimezone

`func (o *RestRCampaign) HasPreferredTimezone() bool`

HasPreferredTimezone returns a boolean if a field has been set.

### GetStrategy

`func (o *RestRCampaign) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *RestRCampaign) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *RestRCampaign) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *RestRCampaign) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetIncludedSegments

`func (o *RestRCampaign) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *RestRCampaign) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *RestRCampaign) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *RestRCampaign) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetIncludedLists

`func (o *RestRCampaign) GetIncludedLists() []string`

GetIncludedLists returns the IncludedLists field if non-nil, zero value otherwise.

### GetIncludedListsOk

`func (o *RestRCampaign) GetIncludedListsOk() (*[]string, bool)`

GetIncludedListsOk returns a tuple with the IncludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLists

`func (o *RestRCampaign) SetIncludedLists(v []string)`

SetIncludedLists sets IncludedLists field to given value.


### GetIncludedTags

`func (o *RestRCampaign) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *RestRCampaign) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *RestRCampaign) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *RestRCampaign) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *RestRCampaign) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *RestRCampaign) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *RestRCampaign) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *RestRCampaign) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetExcludedLists

`func (o *RestRCampaign) GetExcludedLists() []string`

GetExcludedLists returns the ExcludedLists field if non-nil, zero value otherwise.

### GetExcludedListsOk

`func (o *RestRCampaign) GetExcludedListsOk() (*[]string, bool)`

GetExcludedListsOk returns a tuple with the ExcludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedLists

`func (o *RestRCampaign) SetExcludedLists(v []string)`

SetExcludedLists sets ExcludedLists field to given value.


### GetExcludedTags

`func (o *RestRCampaign) GetExcludedTags() []string`

GetExcludedTags returns the ExcludedTags field if non-nil, zero value otherwise.

### GetExcludedTagsOk

`func (o *RestRCampaign) GetExcludedTagsOk() (*[]string, bool)`

GetExcludedTagsOk returns a tuple with the ExcludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTags

`func (o *RestRCampaign) SetExcludedTags(v []string)`

SetExcludedTags sets ExcludedTags field to given value.

### HasExcludedTags

`func (o *RestRCampaign) HasExcludedTags() bool`

HasExcludedTags returns a boolean if a field has been set.

### GetIsArchived

`func (o *RestRCampaign) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *RestRCampaign) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *RestRCampaign) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *RestRCampaign) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetCampaignScreenshotUrl

`func (o *RestRCampaign) GetCampaignScreenshotUrl() string`

GetCampaignScreenshotUrl returns the CampaignScreenshotUrl field if non-nil, zero value otherwise.

### GetCampaignScreenshotUrlOk

`func (o *RestRCampaign) GetCampaignScreenshotUrlOk() (*string, bool)`

GetCampaignScreenshotUrlOk returns a tuple with the CampaignScreenshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignScreenshotUrl

`func (o *RestRCampaign) SetCampaignScreenshotUrl(v string)`

SetCampaignScreenshotUrl sets CampaignScreenshotUrl field to given value.

### HasCampaignScreenshotUrl

`func (o *RestRCampaign) HasCampaignScreenshotUrl() bool`

HasCampaignScreenshotUrl returns a boolean if a field has been set.

### GetCreated

`func (o *RestRCampaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRCampaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRCampaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRCampaign) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRCampaign) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRCampaign) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRCampaign) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRCampaign) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


