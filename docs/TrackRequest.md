# TrackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**Email**](email.md) | Email address of the contact to track. | [optional] 
**AddTags** | Pointer to **[]string** |  | [optional] 
**RemoveTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackRequest

`func NewTrackRequest() *TrackRequest`

NewTrackRequest instantiates a new TrackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackRequestWithDefaults

`func NewTrackRequestWithDefaults() *TrackRequest`

NewTrackRequestWithDefaults instantiates a new TrackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TrackRequest) GetEmail() Email`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrackRequest) GetEmailOk() (*Email, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrackRequest) SetEmail(v Email)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrackRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddTags

`func (o *TrackRequest) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *TrackRequest) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *TrackRequest) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *TrackRequest) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *TrackRequest) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *TrackRequest) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *TrackRequest) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *TrackRequest) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


