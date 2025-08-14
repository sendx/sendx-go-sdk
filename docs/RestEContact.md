# RestEContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of the contact | [optional] 
**LastName** | Pointer to **string** | Last name of the contact | [optional] 
**Email** | **string** | Email address of the contact (required and must be unique within team).  **Validation:** - Must be a valid email format - Must be unique within the team - Cannot be empty or null  | 
**Company** | Pointer to **string** | Company name of the contact | [optional] 
**CustomFields** | Pointer to **map[string]string** | Custom fields as key-value pairs. Keys should use &#x60;field_&#x60; prefix.  **Processing:** - Keys are case-sensitive - Values are stored as strings  **Examples:** - &#x60;\&quot;field_MnuqBAG2NPLm7PZMWbjQxt\&quot;: \&quot;Engineering\&quot;&#x60; → stored as &#x60;department: Engineering&#x60;  | [optional] 
**Lists** | Pointer to **[]string** | Array of list identifiers to associate with the contact.  Identifiers should use &#x60;list_&#x60; prefix.  **Processing:** - Invalid list IDs will send our 400 error - Duplicates will be removed  **Examples:** - &#x60;\&quot;list_OcuxJHdiAvujmwQVJfd3ss\&quot;&#x60; → valid prefixed format  | [optional] 
**Tags** | Pointer to **[]string** | Array of tag identifiers to associate with the contact. Identifiers should use &#x60;tag_&#x60; prefix.  **Processing:** - Invalid tag IDs will be ignored - Duplicates will be removed  **Examples:** - &#x60;\&quot;tag_UhsDkjL772Qbj5lWtT62VK\&quot;&#x60; → valid prefixed format  | [optional] 
**LastTrackedIp** | Pointer to **string** | Last tracked IP address of the contact for analytics purposes.  **Usage:** - Used for geographic analytics - Helps with spam detection - Optional field  | [optional] 

## Methods

### NewRestEContact

`func NewRestEContact(email string, ) *RestEContact`

NewRestEContact instantiates a new RestEContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestEContactWithDefaults

`func NewRestEContactWithDefaults() *RestEContact`

NewRestEContactWithDefaults instantiates a new RestEContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *RestEContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RestEContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RestEContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RestEContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RestEContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RestEContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RestEContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RestEContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *RestEContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RestEContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RestEContact) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCompany

`func (o *RestEContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *RestEContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *RestEContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *RestEContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCustomFields

`func (o *RestEContact) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RestEContact) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RestEContact) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RestEContact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLists

`func (o *RestEContact) GetLists() []string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *RestEContact) GetListsOk() (*[]string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *RestEContact) SetLists(v []string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *RestEContact) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetTags

`func (o *RestEContact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestEContact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestEContact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestEContact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastTrackedIp

`func (o *RestEContact) GetLastTrackedIp() string`

GetLastTrackedIp returns the LastTrackedIp field if non-nil, zero value otherwise.

### GetLastTrackedIpOk

`func (o *RestEContact) GetLastTrackedIpOk() (*string, bool)`

GetLastTrackedIpOk returns a tuple with the LastTrackedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrackedIp

`func (o *RestEContact) SetLastTrackedIp(v string)`

SetLastTrackedIp sets LastTrackedIp field to given value.

### HasLastTrackedIp

`func (o *RestEContact) HasLastTrackedIp() bool`

HasLastTrackedIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


