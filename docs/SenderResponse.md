# SenderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Encrypted ID of the sender | 
**Name** | **string** | Name of the sender | 
**Email** | **string** | Email address of the sender | 
**IsWhitelisted** | **bool** | Indicates if the sender is whitelisted | 

## Methods

### NewSenderResponse

`func NewSenderResponse(id string, name string, email string, isWhitelisted bool, ) *SenderResponse`

NewSenderResponse instantiates a new SenderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderResponseWithDefaults

`func NewSenderResponseWithDefaults() *SenderResponse`

NewSenderResponseWithDefaults instantiates a new SenderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SenderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SenderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SenderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SenderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SenderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SenderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *SenderResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SenderResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SenderResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsWhitelisted

`func (o *SenderResponse) GetIsWhitelisted() bool`

GetIsWhitelisted returns the IsWhitelisted field if non-nil, zero value otherwise.

### GetIsWhitelistedOk

`func (o *SenderResponse) GetIsWhitelistedOk() (*bool, bool)`

GetIsWhitelistedOk returns a tuple with the IsWhitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhitelisted

`func (o *SenderResponse) SetIsWhitelisted(v bool)`

SetIsWhitelisted sets IsWhitelisted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


