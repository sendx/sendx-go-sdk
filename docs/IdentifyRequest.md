# IdentifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of the contact. | [optional] 
**LastName** | Pointer to **string** | Last name of the contact. | [optional] 
**Email** | **string** | Email address of the contact. | 
**NewEmail** | Pointer to **string** | New email address of the contact. | [optional] 
**Company** | Pointer to **string** | Company of the contact. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIdentifyRequest

`func NewIdentifyRequest(email string, ) *IdentifyRequest`

NewIdentifyRequest instantiates a new IdentifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifyRequestWithDefaults

`func NewIdentifyRequestWithDefaults() *IdentifyRequest`

NewIdentifyRequestWithDefaults instantiates a new IdentifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *IdentifyRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentifyRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentifyRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentifyRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *IdentifyRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentifyRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentifyRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentifyRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *IdentifyRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentifyRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentifyRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewEmail

`func (o *IdentifyRequest) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *IdentifyRequest) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *IdentifyRequest) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.

### HasNewEmail

`func (o *IdentifyRequest) HasNewEmail() bool`

HasNewEmail returns a boolean if a field has been set.

### GetCompany

`func (o *IdentifyRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *IdentifyRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *IdentifyRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *IdentifyRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTags

`func (o *IdentifyRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IdentifyRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IdentifyRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IdentifyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IdentifyRequest) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IdentifyRequest) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IdentifyRequest) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IdentifyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


