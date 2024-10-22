# ContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address of the contact. Must be a valid email. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact. | [optional] 
**Company** | Pointer to **string** | The company where the contact works. | [optional] 
**LastTrackedIp** | Pointer to **string** | The last known IP address of the contact. | [optional] 
**CustomFields** | Pointer to **map[string]string** | Custom fields specific to the contact, which may vary by account. | [optional] 
**Lists** | Pointer to **[]string** | A list of &#x60;lists&#x60; ids the contact is subscribed to. | [optional] 
**Tags** | Pointer to **[]string** | Tag ids associated with the contact for segmentation or categorization. | [optional] 

## Methods

### NewContactRequest

`func NewContactRequest() *ContactRequest`

NewContactRequest instantiates a new ContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestWithDefaults

`func NewContactRequestWithDefaults() *ContactRequest`

NewContactRequestWithDefaults instantiates a new ContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ContactRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ContactRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *ContactRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ContactRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ContactRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ContactRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLastTrackedIp

`func (o *ContactRequest) GetLastTrackedIp() string`

GetLastTrackedIp returns the LastTrackedIp field if non-nil, zero value otherwise.

### GetLastTrackedIpOk

`func (o *ContactRequest) GetLastTrackedIpOk() (*string, bool)`

GetLastTrackedIpOk returns a tuple with the LastTrackedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrackedIp

`func (o *ContactRequest) SetLastTrackedIp(v string)`

SetLastTrackedIp sets LastTrackedIp field to given value.

### HasLastTrackedIp

`func (o *ContactRequest) HasLastTrackedIp() bool`

HasLastTrackedIp returns a boolean if a field has been set.

### GetCustomFields

`func (o *ContactRequest) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ContactRequest) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ContactRequest) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ContactRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLists

`func (o *ContactRequest) GetLists() []string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ContactRequest) GetListsOk() (*[]string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ContactRequest) SetLists(v []string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *ContactRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetTags

`func (o *ContactRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


