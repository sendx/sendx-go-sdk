# XTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewXTo

`func NewXTo(email string, ) *XTo`

NewXTo instantiates a new XTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXToWithDefaults

`func NewXToWithDefaults() *XTo`

NewXToWithDefaults instantiates a new XTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *XTo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *XTo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *XTo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *XTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCc

`func (o *XTo) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *XTo) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *XTo) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *XTo) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *XTo) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *XTo) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *XTo) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *XTo) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCustomFields

`func (o *XTo) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *XTo) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *XTo) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *XTo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


