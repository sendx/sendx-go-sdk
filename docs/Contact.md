# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for the contact. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact. | [optional] 
**Email** | Pointer to **string** | The email address of the contact. | [optional] 
**Company** | Pointer to **string** | The company of the contact. | [optional] 
**CustomFields** | Pointer to **map[string]string** | Custom fields and their values | [optional] 
**Unsubscribed** | Pointer to **bool** | Indicates if the contact has unsubscribed from emails. | [optional] 
**Bounced** | Pointer to **bool** | Indicates if the contact&#39;s email has bounced. | [optional] 
**Spam** | Pointer to **bool** | Indicates if the contact marked the email as spam. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the contact was created. | [optional] 
**Updated** | Pointer to **time.Time** | The date and time when the contact was last updated. | [optional] 
**Blocked** | Pointer to **bool** | Indicates if the contact is blocked from receiving emails. | [optional] 
**Dropped** | Pointer to **bool** | Indicates if emails to this contact were dropped. | [optional] 
**LTV** | Pointer to **int32** | Lifetime value (LTV) of the contact in currency units. | [optional] 
**ContactSource** | Pointer to **int32** | The source from which the contact was added. Possible values:  | [optional] 
**LastTrackedIp** | Pointer to **string** | The last known IP address tracked for the contact. | [optional] 
**Lists** | Pointer to **[]string** | A list of &#x60;lists&#x60; ids the contact is subscribed to. | [optional] 
**Tags** | Pointer to **[]string** | &#x60;Tag&#x60; ids associated with the contact for segmentation or categorization. | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *Contact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Contact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Contact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Contact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCustomFields

`func (o *Contact) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Contact) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Contact) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Contact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *Contact) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *Contact) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *Contact) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *Contact) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetBounced

`func (o *Contact) GetBounced() bool`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *Contact) GetBouncedOk() (*bool, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *Contact) SetBounced(v bool)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *Contact) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetSpam

`func (o *Contact) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *Contact) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *Contact) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *Contact) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetCreated

`func (o *Contact) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Contact) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Contact) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Contact) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Contact) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Contact) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Contact) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Contact) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetBlocked

`func (o *Contact) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Contact) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Contact) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *Contact) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetDropped

`func (o *Contact) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *Contact) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *Contact) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *Contact) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetLTV

`func (o *Contact) GetLTV() int32`

GetLTV returns the LTV field if non-nil, zero value otherwise.

### GetLTVOk

`func (o *Contact) GetLTVOk() (*int32, bool)`

GetLTVOk returns a tuple with the LTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLTV

`func (o *Contact) SetLTV(v int32)`

SetLTV sets LTV field to given value.

### HasLTV

`func (o *Contact) HasLTV() bool`

HasLTV returns a boolean if a field has been set.

### GetContactSource

`func (o *Contact) GetContactSource() int32`

GetContactSource returns the ContactSource field if non-nil, zero value otherwise.

### GetContactSourceOk

`func (o *Contact) GetContactSourceOk() (*int32, bool)`

GetContactSourceOk returns a tuple with the ContactSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSource

`func (o *Contact) SetContactSource(v int32)`

SetContactSource sets ContactSource field to given value.

### HasContactSource

`func (o *Contact) HasContactSource() bool`

HasContactSource returns a boolean if a field has been set.

### GetLastTrackedIp

`func (o *Contact) GetLastTrackedIp() string`

GetLastTrackedIp returns the LastTrackedIp field if non-nil, zero value otherwise.

### GetLastTrackedIpOk

`func (o *Contact) GetLastTrackedIpOk() (*string, bool)`

GetLastTrackedIpOk returns a tuple with the LastTrackedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrackedIp

`func (o *Contact) SetLastTrackedIp(v string)`

SetLastTrackedIp sets LastTrackedIp field to given value.

### HasLastTrackedIp

`func (o *Contact) HasLastTrackedIp() bool`

HasLastTrackedIp returns a boolean if a field has been set.

### GetLists

`func (o *Contact) GetLists() []string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Contact) GetListsOk() (*[]string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Contact) SetLists(v []string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *Contact) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetTags

`func (o *Contact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Contact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Contact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Contact) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


