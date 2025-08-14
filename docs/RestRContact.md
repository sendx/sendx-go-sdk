# RestRContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique contact identifier with contact_ prefix.  **Format:** &#x60;contact_&#x60; + 22 alphanumeric characters  **Usage:** - Use this ID for all subsequent API calls - Unique across the entire SendX platform - Never changes once created  | [optional] 
**FirstName** | Pointer to **string** | First name of the contact | [optional] 
**LastName** | Pointer to **string** | Last name of the contact | [optional] 
**Email** | Pointer to **string** | Email address of the contact (unique within team) | [optional] 
**Company** | Pointer to **string** | Company name of the contact | [optional] 
**CustomFields** | Pointer to **map[string]string** | Custom fields with field_ prefixed keys.  **Format:** All keys have &#x60;field_&#x60; prefix in responses  **Example Structure:** &#x60;&#x60;&#x60;json {   \&quot;field_MnuqBAG2NPLm7PZMWbjQxt\&quot;: \&quot;Engineering\&quot;,   \&quot;field_QqfhckbdcvQinLPlduIbHq\&quot;: \&quot;Senior\&quot;,   \&quot;field_MnuqBAG2NPLm7PZMWbjQxt\&quot;: \&quot;$75000\&quot; } &#x60;&#x60;&#x60;  | [optional] 
**Lists** | Pointer to **[]string** | Associated lists with list_ prefixed identifiers.  **Format:** All IDs have &#x60;list_&#x60; prefix in responses  | [optional] 
**Tags** | Pointer to **[]string** | Associated tags with tag_ prefixed identifiers.  **Format:** All IDs have &#x60;tag_&#x60; prefix in responses  | [optional] 
**Unsubscribed** | Pointer to **bool** | Whether the contact has unsubscribed from emails | [optional] 
**Bounced** | Pointer to **bool** | Whether emails to this contact have bounced | [optional] 
**Spam** | Pointer to **bool** | Whether the contact has marked emails as spam | [optional] 
**Blocked** | Pointer to **bool** | Whether the contact is blocked from receiving emails | [optional] 
**Dropped** | Pointer to **bool** | Whether emails to this contact have been dropped | [optional] 
**Created** | Pointer to **time.Time** | Contact creation timestamp (ISO 8601 format) | [optional] 
**Updated** | Pointer to **time.Time** | Contact last update timestamp (ISO 8601 format) | [optional] 
**TrackData** | Pointer to **string** | Email tracking data and UTM parameters.  **Contains:** - UTM parameters from campaigns - Attribution data - Custom tracking parameters  | [optional] 
**ContactSource** | Pointer to **int32** | Source type of the contact creation.  **Values:** - &#x60;1&#x60; - API/Manual - &#x60;2&#x60; - Import - &#x60;3&#x60; - Form - &#x60;4&#x60; - Integration  | [optional] 
**PageSource** | Pointer to **string** | URL of the page where contact was created | [optional] 
**LastTrackedIp** | Pointer to **string** | Last tracked IP address of the contact | [optional] 
**LTV** | Pointer to **int32** | Lifetime Value of the contact in cents.  **Example:** 7500 &#x3D; $75.00  | [optional] 

## Methods

### NewRestRContact

`func NewRestRContact() *RestRContact`

NewRestRContact instantiates a new RestRContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRContactWithDefaults

`func NewRestRContactWithDefaults() *RestRContact`

NewRestRContactWithDefaults instantiates a new RestRContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *RestRContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RestRContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RestRContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RestRContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RestRContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RestRContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RestRContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RestRContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *RestRContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RestRContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RestRContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RestRContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *RestRContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *RestRContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *RestRContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *RestRContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCustomFields

`func (o *RestRContact) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RestRContact) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RestRContact) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RestRContact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLists

`func (o *RestRContact) GetLists() []string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *RestRContact) GetListsOk() (*[]string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *RestRContact) SetLists(v []string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *RestRContact) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetTags

`func (o *RestRContact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestRContact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestRContact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestRContact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *RestRContact) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *RestRContact) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *RestRContact) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *RestRContact) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetBounced

`func (o *RestRContact) GetBounced() bool`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *RestRContact) GetBouncedOk() (*bool, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *RestRContact) SetBounced(v bool)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *RestRContact) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetSpam

`func (o *RestRContact) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *RestRContact) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *RestRContact) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *RestRContact) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetBlocked

`func (o *RestRContact) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RestRContact) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RestRContact) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RestRContact) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetDropped

`func (o *RestRContact) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *RestRContact) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *RestRContact) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *RestRContact) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetCreated

`func (o *RestRContact) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRContact) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRContact) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRContact) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRContact) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRContact) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRContact) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRContact) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetTrackData

`func (o *RestRContact) GetTrackData() string`

GetTrackData returns the TrackData field if non-nil, zero value otherwise.

### GetTrackDataOk

`func (o *RestRContact) GetTrackDataOk() (*string, bool)`

GetTrackDataOk returns a tuple with the TrackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackData

`func (o *RestRContact) SetTrackData(v string)`

SetTrackData sets TrackData field to given value.

### HasTrackData

`func (o *RestRContact) HasTrackData() bool`

HasTrackData returns a boolean if a field has been set.

### GetContactSource

`func (o *RestRContact) GetContactSource() int32`

GetContactSource returns the ContactSource field if non-nil, zero value otherwise.

### GetContactSourceOk

`func (o *RestRContact) GetContactSourceOk() (*int32, bool)`

GetContactSourceOk returns a tuple with the ContactSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSource

`func (o *RestRContact) SetContactSource(v int32)`

SetContactSource sets ContactSource field to given value.

### HasContactSource

`func (o *RestRContact) HasContactSource() bool`

HasContactSource returns a boolean if a field has been set.

### GetPageSource

`func (o *RestRContact) GetPageSource() string`

GetPageSource returns the PageSource field if non-nil, zero value otherwise.

### GetPageSourceOk

`func (o *RestRContact) GetPageSourceOk() (*string, bool)`

GetPageSourceOk returns a tuple with the PageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSource

`func (o *RestRContact) SetPageSource(v string)`

SetPageSource sets PageSource field to given value.

### HasPageSource

`func (o *RestRContact) HasPageSource() bool`

HasPageSource returns a boolean if a field has been set.

### GetLastTrackedIp

`func (o *RestRContact) GetLastTrackedIp() string`

GetLastTrackedIp returns the LastTrackedIp field if non-nil, zero value otherwise.

### GetLastTrackedIpOk

`func (o *RestRContact) GetLastTrackedIpOk() (*string, bool)`

GetLastTrackedIpOk returns a tuple with the LastTrackedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrackedIp

`func (o *RestRContact) SetLastTrackedIp(v string)`

SetLastTrackedIp sets LastTrackedIp field to given value.

### HasLastTrackedIp

`func (o *RestRContact) HasLastTrackedIp() bool`

HasLastTrackedIp returns a boolean if a field has been set.

### GetLTV

`func (o *RestRContact) GetLTV() int32`

GetLTV returns the LTV field if non-nil, zero value otherwise.

### GetLTVOk

`func (o *RestRContact) GetLTVOk() (*int32, bool)`

GetLTVOk returns a tuple with the LTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLTV

`func (o *RestRContact) SetLTV(v int32)`

SetLTV sets LTV field to given value.

### HasLTV

`func (o *RestRContact) HasLTV() bool`

HasLTV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


