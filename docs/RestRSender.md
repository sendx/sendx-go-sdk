# RestRSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique sender identifier with sender_ prefix | [optional] 
**Name** | Pointer to **string** | Sender display name | [optional] 
**Email** | Pointer to **string** | Sender email address | [optional] 
**IsWhitelisted** | Pointer to **bool** | Sender whitelist status | [optional] 

## Methods

### NewRestRSender

`func NewRestRSender() *RestRSender`

NewRestRSender instantiates a new RestRSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRSenderWithDefaults

`func NewRestRSenderWithDefaults() *RestRSender`

NewRestRSenderWithDefaults instantiates a new RestRSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRSender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRSender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRSender) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRSender) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *RestRSender) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RestRSender) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RestRSender) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RestRSender) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsWhitelisted

`func (o *RestRSender) GetIsWhitelisted() bool`

GetIsWhitelisted returns the IsWhitelisted field if non-nil, zero value otherwise.

### GetIsWhitelistedOk

`func (o *RestRSender) GetIsWhitelistedOk() (*bool, bool)`

GetIsWhitelistedOk returns a tuple with the IsWhitelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhitelisted

`func (o *RestRSender) SetIsWhitelisted(v bool)`

SetIsWhitelisted sets IsWhitelisted field to given value.

### HasIsWhitelisted

`func (o *RestRSender) HasIsWhitelisted() bool`

HasIsWhitelisted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


