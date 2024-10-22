# ListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Encrypted ID of the list | [optional] 
**Name** | Pointer to **string** | Name of the list | [optional] 
**Type** | Pointer to **int32** | Type of the list representing opt-in methods (1: Single Opt-In, 2: Double Opt-In) | [optional] 
**SendThankYouMail** | Pointer to **bool** | Indicates if a thank-you email should be sent | [optional] 
**ThankYouFromName** | Pointer to **string** | Name displayed as the sender in the thank-you email | [optional] 
**ThankYouFromEmail** | Pointer to **string** | Email address from which the thank-you email is sent | [optional] 
**ThankYouMailSubject** | Pointer to **string** | Subject line of the thank-you email | [optional] 
**ThankYouMailMessage** | Pointer to **string** | Plain text message body of the thank-you email | [optional] 
**ThankYouSender** | Pointer to **string** | Details of the sender of the thank-you email | [optional] 
**ConfirmFromName** | Pointer to **string** | Name displayed as the sender in the confirmation email | [optional] 
**ConfirmFromEmail** | Pointer to **string** | Email address from which the confirmation email is sent | [optional] 
**ConfirmMailSubject** | Pointer to **string** | Subject line of the confirmation email | [optional] 
**ConfirmMailMessage** | Pointer to **string** | Plain text message body of the confirmation email | [optional] 
**ConfirmSuccessPage** | Pointer to **string** | URL of the success page after confirmation | [optional] 
**Created** | Pointer to **time.Time** | Date and time when the list was created | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when the list was last updated | [optional] 
**ConfirmSender** | Pointer to **NullableString** | Details of the sender of the confirmation email | [optional] 

## Methods

### NewListModel

`func NewListModel() *ListModel`

NewListModel instantiates a new ListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListModelWithDefaults

`func NewListModelWithDefaults() *ListModel`

NewListModelWithDefaults instantiates a new ListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListModel) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListModel) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListModel) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ListModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSendThankYouMail

`func (o *ListModel) GetSendThankYouMail() bool`

GetSendThankYouMail returns the SendThankYouMail field if non-nil, zero value otherwise.

### GetSendThankYouMailOk

`func (o *ListModel) GetSendThankYouMailOk() (*bool, bool)`

GetSendThankYouMailOk returns a tuple with the SendThankYouMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendThankYouMail

`func (o *ListModel) SetSendThankYouMail(v bool)`

SetSendThankYouMail sets SendThankYouMail field to given value.

### HasSendThankYouMail

`func (o *ListModel) HasSendThankYouMail() bool`

HasSendThankYouMail returns a boolean if a field has been set.

### GetThankYouFromName

`func (o *ListModel) GetThankYouFromName() string`

GetThankYouFromName returns the ThankYouFromName field if non-nil, zero value otherwise.

### GetThankYouFromNameOk

`func (o *ListModel) GetThankYouFromNameOk() (*string, bool)`

GetThankYouFromNameOk returns a tuple with the ThankYouFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouFromName

`func (o *ListModel) SetThankYouFromName(v string)`

SetThankYouFromName sets ThankYouFromName field to given value.

### HasThankYouFromName

`func (o *ListModel) HasThankYouFromName() bool`

HasThankYouFromName returns a boolean if a field has been set.

### GetThankYouFromEmail

`func (o *ListModel) GetThankYouFromEmail() string`

GetThankYouFromEmail returns the ThankYouFromEmail field if non-nil, zero value otherwise.

### GetThankYouFromEmailOk

`func (o *ListModel) GetThankYouFromEmailOk() (*string, bool)`

GetThankYouFromEmailOk returns a tuple with the ThankYouFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouFromEmail

`func (o *ListModel) SetThankYouFromEmail(v string)`

SetThankYouFromEmail sets ThankYouFromEmail field to given value.

### HasThankYouFromEmail

`func (o *ListModel) HasThankYouFromEmail() bool`

HasThankYouFromEmail returns a boolean if a field has been set.

### GetThankYouMailSubject

`func (o *ListModel) GetThankYouMailSubject() string`

GetThankYouMailSubject returns the ThankYouMailSubject field if non-nil, zero value otherwise.

### GetThankYouMailSubjectOk

`func (o *ListModel) GetThankYouMailSubjectOk() (*string, bool)`

GetThankYouMailSubjectOk returns a tuple with the ThankYouMailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMailSubject

`func (o *ListModel) SetThankYouMailSubject(v string)`

SetThankYouMailSubject sets ThankYouMailSubject field to given value.

### HasThankYouMailSubject

`func (o *ListModel) HasThankYouMailSubject() bool`

HasThankYouMailSubject returns a boolean if a field has been set.

### GetThankYouMailMessage

`func (o *ListModel) GetThankYouMailMessage() string`

GetThankYouMailMessage returns the ThankYouMailMessage field if non-nil, zero value otherwise.

### GetThankYouMailMessageOk

`func (o *ListModel) GetThankYouMailMessageOk() (*string, bool)`

GetThankYouMailMessageOk returns a tuple with the ThankYouMailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMailMessage

`func (o *ListModel) SetThankYouMailMessage(v string)`

SetThankYouMailMessage sets ThankYouMailMessage field to given value.

### HasThankYouMailMessage

`func (o *ListModel) HasThankYouMailMessage() bool`

HasThankYouMailMessage returns a boolean if a field has been set.

### GetThankYouSender

`func (o *ListModel) GetThankYouSender() string`

GetThankYouSender returns the ThankYouSender field if non-nil, zero value otherwise.

### GetThankYouSenderOk

`func (o *ListModel) GetThankYouSenderOk() (*string, bool)`

GetThankYouSenderOk returns a tuple with the ThankYouSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouSender

`func (o *ListModel) SetThankYouSender(v string)`

SetThankYouSender sets ThankYouSender field to given value.

### HasThankYouSender

`func (o *ListModel) HasThankYouSender() bool`

HasThankYouSender returns a boolean if a field has been set.

### GetConfirmFromName

`func (o *ListModel) GetConfirmFromName() string`

GetConfirmFromName returns the ConfirmFromName field if non-nil, zero value otherwise.

### GetConfirmFromNameOk

`func (o *ListModel) GetConfirmFromNameOk() (*string, bool)`

GetConfirmFromNameOk returns a tuple with the ConfirmFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmFromName

`func (o *ListModel) SetConfirmFromName(v string)`

SetConfirmFromName sets ConfirmFromName field to given value.

### HasConfirmFromName

`func (o *ListModel) HasConfirmFromName() bool`

HasConfirmFromName returns a boolean if a field has been set.

### GetConfirmFromEmail

`func (o *ListModel) GetConfirmFromEmail() string`

GetConfirmFromEmail returns the ConfirmFromEmail field if non-nil, zero value otherwise.

### GetConfirmFromEmailOk

`func (o *ListModel) GetConfirmFromEmailOk() (*string, bool)`

GetConfirmFromEmailOk returns a tuple with the ConfirmFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmFromEmail

`func (o *ListModel) SetConfirmFromEmail(v string)`

SetConfirmFromEmail sets ConfirmFromEmail field to given value.

### HasConfirmFromEmail

`func (o *ListModel) HasConfirmFromEmail() bool`

HasConfirmFromEmail returns a boolean if a field has been set.

### GetConfirmMailSubject

`func (o *ListModel) GetConfirmMailSubject() string`

GetConfirmMailSubject returns the ConfirmMailSubject field if non-nil, zero value otherwise.

### GetConfirmMailSubjectOk

`func (o *ListModel) GetConfirmMailSubjectOk() (*string, bool)`

GetConfirmMailSubjectOk returns a tuple with the ConfirmMailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmMailSubject

`func (o *ListModel) SetConfirmMailSubject(v string)`

SetConfirmMailSubject sets ConfirmMailSubject field to given value.

### HasConfirmMailSubject

`func (o *ListModel) HasConfirmMailSubject() bool`

HasConfirmMailSubject returns a boolean if a field has been set.

### GetConfirmMailMessage

`func (o *ListModel) GetConfirmMailMessage() string`

GetConfirmMailMessage returns the ConfirmMailMessage field if non-nil, zero value otherwise.

### GetConfirmMailMessageOk

`func (o *ListModel) GetConfirmMailMessageOk() (*string, bool)`

GetConfirmMailMessageOk returns a tuple with the ConfirmMailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmMailMessage

`func (o *ListModel) SetConfirmMailMessage(v string)`

SetConfirmMailMessage sets ConfirmMailMessage field to given value.

### HasConfirmMailMessage

`func (o *ListModel) HasConfirmMailMessage() bool`

HasConfirmMailMessage returns a boolean if a field has been set.

### GetConfirmSuccessPage

`func (o *ListModel) GetConfirmSuccessPage() string`

GetConfirmSuccessPage returns the ConfirmSuccessPage field if non-nil, zero value otherwise.

### GetConfirmSuccessPageOk

`func (o *ListModel) GetConfirmSuccessPageOk() (*string, bool)`

GetConfirmSuccessPageOk returns a tuple with the ConfirmSuccessPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmSuccessPage

`func (o *ListModel) SetConfirmSuccessPage(v string)`

SetConfirmSuccessPage sets ConfirmSuccessPage field to given value.

### HasConfirmSuccessPage

`func (o *ListModel) HasConfirmSuccessPage() bool`

HasConfirmSuccessPage returns a boolean if a field has been set.

### GetCreated

`func (o *ListModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListModel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ListModel) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ListModel) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ListModel) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ListModel) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetConfirmSender

`func (o *ListModel) GetConfirmSender() string`

GetConfirmSender returns the ConfirmSender field if non-nil, zero value otherwise.

### GetConfirmSenderOk

`func (o *ListModel) GetConfirmSenderOk() (*string, bool)`

GetConfirmSenderOk returns a tuple with the ConfirmSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmSender

`func (o *ListModel) SetConfirmSender(v string)`

SetConfirmSender sets ConfirmSender field to given value.

### HasConfirmSender

`func (o *ListModel) HasConfirmSender() bool`

HasConfirmSender returns a boolean if a field has been set.

### SetConfirmSenderNil

`func (o *ListModel) SetConfirmSenderNil(b bool)`

 SetConfirmSenderNil sets the value for ConfirmSender to be an explicit nil

### UnsetConfirmSender
`func (o *ListModel) UnsetConfirmSender()`

UnsetConfirmSender ensures that no value is present for ConfirmSender, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


