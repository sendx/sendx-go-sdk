# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Encrypted ID of the campaign | [optional] 
**Name** | Pointer to **string** | Name of the campaign | [optional] 
**TrackReply** | Pointer to **bool** | Indicates if replies to the campaign should be tracked | [optional] 
**Status** | Pointer to **int32** | Campaign status: 0 (Draft), 1 (Scheduled), 2 (Sending), 3 (Sent), 4 (Quarantined) | [optional] 
**ScheduleType** | Pointer to **int32** | Schedule type: 0 (Schedule later), 1 (Send Now), 2 (Trigger via automation), 3 (Recurring) | [optional] 
**ScheduleCondition** | Pointer to **string** | Condition for scheduling the campaign | [optional] 
**TimeCondition** | Pointer to **string** | Time-related condition for the campaign | [optional] 
**Timezone** | Pointer to **string** | Timezone for the scheduled send | [optional] 
**PreferredTimeCondition** | Pointer to **string** | Preferred time condition for the campaign | [optional] 
**PreferredTimezone** | Pointer to **string** | Preferred timezone for sending the campaign | [optional] 
**Strategy** | Pointer to **string** | Strategy for the campaign | [optional] 
**SendInContactsTimezone** | Pointer to **bool** | Indicates if the campaign should be sent in the recipient&#39;s timezone | [optional] 
**SmartSend** | Pointer to **bool** | Indicates if smart sending should be used | [optional] 
**IsArchived** | Pointer to **bool** | Indicates if the campaign is archived | [optional] 
**Sender** | Pointer to **string** | Information about the sender of the campaign | [optional] 
**CampaignScreenshotUrl** | Pointer to **string** | URL of the campaign&#39;s screenshot | [optional] 
**IncludedSegments** | Pointer to **[]string** | Segments to be included in the campaign | [optional] 
**IncludedLists** | Pointer to **[]string** | Lists to be included in the campaign | [optional] 
**IncludedTags** | Pointer to **[]string** | Tags to be included in the campaign | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segments to be excluded from the campaign | [optional] 
**ExcludedLists** | Pointer to **[]string** | Lists to be excluded from the campaign | [optional] 
**ExcludedTags** | Pointer to **[]string** | Tags to be excluded from the campaign | [optional] 

## Methods

### NewCampaign

`func NewCampaign() *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Campaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Campaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrackReply

`func (o *Campaign) GetTrackReply() bool`

GetTrackReply returns the TrackReply field if non-nil, zero value otherwise.

### GetTrackReplyOk

`func (o *Campaign) GetTrackReplyOk() (*bool, bool)`

GetTrackReplyOk returns a tuple with the TrackReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackReply

`func (o *Campaign) SetTrackReply(v bool)`

SetTrackReply sets TrackReply field to given value.

### HasTrackReply

`func (o *Campaign) HasTrackReply() bool`

HasTrackReply returns a boolean if a field has been set.

### GetStatus

`func (o *Campaign) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Campaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduleType

`func (o *Campaign) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *Campaign) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *Campaign) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *Campaign) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleCondition

`func (o *Campaign) GetScheduleCondition() string`

GetScheduleCondition returns the ScheduleCondition field if non-nil, zero value otherwise.

### GetScheduleConditionOk

`func (o *Campaign) GetScheduleConditionOk() (*string, bool)`

GetScheduleConditionOk returns a tuple with the ScheduleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCondition

`func (o *Campaign) SetScheduleCondition(v string)`

SetScheduleCondition sets ScheduleCondition field to given value.

### HasScheduleCondition

`func (o *Campaign) HasScheduleCondition() bool`

HasScheduleCondition returns a boolean if a field has been set.

### GetTimeCondition

`func (o *Campaign) GetTimeCondition() string`

GetTimeCondition returns the TimeCondition field if non-nil, zero value otherwise.

### GetTimeConditionOk

`func (o *Campaign) GetTimeConditionOk() (*string, bool)`

GetTimeConditionOk returns a tuple with the TimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCondition

`func (o *Campaign) SetTimeCondition(v string)`

SetTimeCondition sets TimeCondition field to given value.

### HasTimeCondition

`func (o *Campaign) HasTimeCondition() bool`

HasTimeCondition returns a boolean if a field has been set.

### GetTimezone

`func (o *Campaign) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Campaign) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Campaign) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Campaign) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPreferredTimeCondition

`func (o *Campaign) GetPreferredTimeCondition() string`

GetPreferredTimeCondition returns the PreferredTimeCondition field if non-nil, zero value otherwise.

### GetPreferredTimeConditionOk

`func (o *Campaign) GetPreferredTimeConditionOk() (*string, bool)`

GetPreferredTimeConditionOk returns a tuple with the PreferredTimeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimeCondition

`func (o *Campaign) SetPreferredTimeCondition(v string)`

SetPreferredTimeCondition sets PreferredTimeCondition field to given value.

### HasPreferredTimeCondition

`func (o *Campaign) HasPreferredTimeCondition() bool`

HasPreferredTimeCondition returns a boolean if a field has been set.

### GetPreferredTimezone

`func (o *Campaign) GetPreferredTimezone() string`

GetPreferredTimezone returns the PreferredTimezone field if non-nil, zero value otherwise.

### GetPreferredTimezoneOk

`func (o *Campaign) GetPreferredTimezoneOk() (*string, bool)`

GetPreferredTimezoneOk returns a tuple with the PreferredTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimezone

`func (o *Campaign) SetPreferredTimezone(v string)`

SetPreferredTimezone sets PreferredTimezone field to given value.

### HasPreferredTimezone

`func (o *Campaign) HasPreferredTimezone() bool`

HasPreferredTimezone returns a boolean if a field has been set.

### GetStrategy

`func (o *Campaign) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Campaign) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Campaign) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Campaign) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSendInContactsTimezone

`func (o *Campaign) GetSendInContactsTimezone() bool`

GetSendInContactsTimezone returns the SendInContactsTimezone field if non-nil, zero value otherwise.

### GetSendInContactsTimezoneOk

`func (o *Campaign) GetSendInContactsTimezoneOk() (*bool, bool)`

GetSendInContactsTimezoneOk returns a tuple with the SendInContactsTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInContactsTimezone

`func (o *Campaign) SetSendInContactsTimezone(v bool)`

SetSendInContactsTimezone sets SendInContactsTimezone field to given value.

### HasSendInContactsTimezone

`func (o *Campaign) HasSendInContactsTimezone() bool`

HasSendInContactsTimezone returns a boolean if a field has been set.

### GetSmartSend

`func (o *Campaign) GetSmartSend() bool`

GetSmartSend returns the SmartSend field if non-nil, zero value otherwise.

### GetSmartSendOk

`func (o *Campaign) GetSmartSendOk() (*bool, bool)`

GetSmartSendOk returns a tuple with the SmartSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSend

`func (o *Campaign) SetSmartSend(v bool)`

SetSmartSend sets SmartSend field to given value.

### HasSmartSend

`func (o *Campaign) HasSmartSend() bool`

HasSmartSend returns a boolean if a field has been set.

### GetIsArchived

`func (o *Campaign) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Campaign) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Campaign) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Campaign) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetSender

`func (o *Campaign) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Campaign) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Campaign) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Campaign) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetCampaignScreenshotUrl

`func (o *Campaign) GetCampaignScreenshotUrl() string`

GetCampaignScreenshotUrl returns the CampaignScreenshotUrl field if non-nil, zero value otherwise.

### GetCampaignScreenshotUrlOk

`func (o *Campaign) GetCampaignScreenshotUrlOk() (*string, bool)`

GetCampaignScreenshotUrlOk returns a tuple with the CampaignScreenshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignScreenshotUrl

`func (o *Campaign) SetCampaignScreenshotUrl(v string)`

SetCampaignScreenshotUrl sets CampaignScreenshotUrl field to given value.

### HasCampaignScreenshotUrl

`func (o *Campaign) HasCampaignScreenshotUrl() bool`

HasCampaignScreenshotUrl returns a boolean if a field has been set.

### GetIncludedSegments

`func (o *Campaign) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *Campaign) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *Campaign) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *Campaign) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetIncludedLists

`func (o *Campaign) GetIncludedLists() []string`

GetIncludedLists returns the IncludedLists field if non-nil, zero value otherwise.

### GetIncludedListsOk

`func (o *Campaign) GetIncludedListsOk() (*[]string, bool)`

GetIncludedListsOk returns a tuple with the IncludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLists

`func (o *Campaign) SetIncludedLists(v []string)`

SetIncludedLists sets IncludedLists field to given value.

### HasIncludedLists

`func (o *Campaign) HasIncludedLists() bool`

HasIncludedLists returns a boolean if a field has been set.

### GetIncludedTags

`func (o *Campaign) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *Campaign) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *Campaign) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *Campaign) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *Campaign) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *Campaign) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *Campaign) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *Campaign) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetExcludedLists

`func (o *Campaign) GetExcludedLists() []string`

GetExcludedLists returns the ExcludedLists field if non-nil, zero value otherwise.

### GetExcludedListsOk

`func (o *Campaign) GetExcludedListsOk() (*[]string, bool)`

GetExcludedListsOk returns a tuple with the ExcludedLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedLists

`func (o *Campaign) SetExcludedLists(v []string)`

SetExcludedLists sets ExcludedLists field to given value.

### HasExcludedLists

`func (o *Campaign) HasExcludedLists() bool`

HasExcludedLists returns a boolean if a field has been set.

### GetExcludedTags

`func (o *Campaign) GetExcludedTags() []string`

GetExcludedTags returns the ExcludedTags field if non-nil, zero value otherwise.

### GetExcludedTagsOk

`func (o *Campaign) GetExcludedTagsOk() (*[]string, bool)`

GetExcludedTagsOk returns a tuple with the ExcludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTags

`func (o *Campaign) SetExcludedTags(v []string)`

SetExcludedTags sets ExcludedTags field to given value.

### HasExcludedTags

`func (o *Campaign) HasExcludedTags() bool`

HasExcludedTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


