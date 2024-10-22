# ListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the list | [optional] 
**Type** | Pointer to **int32** | Type of the list representing opt-in methods &lt;br&gt; 0: Single opt-in list &lt;br&gt; 1: Double opt-in list  | [optional] 
**SendThankYouMail** | Pointer to **bool** | Indicates if a thank-you email should be sent | [optional] 
**ThankYouFromName** | Pointer to **string** | Name displayed as the sender in the thank-you email | [optional] 
**ThankYouFromEmail** | Pointer to **string** | Email address from which the thank-you email is sent | [optional] 
**ThankYouMailSubject** | Pointer to **string** | Subject line of the thank-you email | [optional] 
**ThankYouMailMessage** | Pointer to **string** | Plain text message body of the thank-you email | [optional] 
**ThankYouSender** | Pointer to **string** | Sender ID for the thank-you email | [optional] 
**ConfirmFromName** | Pointer to **string** | Name displayed as the sender in the confirmation email | [optional] 
**ConfirmFromEmail** | Pointer to **string** | Email address from which the confirmation email is sent | [optional] 
**ConfirmMailSubject** | Pointer to **string** | Subject line of the confirmation email | [optional] 
**ConfirmMailMessage** | Pointer to **string** | Plain text message body of the confirmation email | [optional] 
**ConfirmSuccessPage** | Pointer to **string** | URL of the success page after confirmation | [optional] 
**ConfirmSender** | Pointer to **string** | Sender ID for the confirmation email | [optional] 

## Methods

### NewListRequest

`func NewListRequest() *ListRequest`

NewListRequest instantiates a new ListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRequestWithDefaults

`func NewListRequestWithDefaults() *ListRequest`

NewListRequestWithDefaults instantiates a new ListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ListRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSendThankYouMail

`func (o *ListRequest) GetSendThankYouMail() bool`

GetSendThankYouMail returns the SendThankYouMail field if non-nil, zero value otherwise.

### GetSendThankYouMailOk

`func (o *ListRequest) GetSendThankYouMailOk() (*bool, bool)`

GetSendThankYouMailOk returns a tuple with the SendThankYouMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendThankYouMail

`func (o *ListRequest) SetSendThankYouMail(v bool)`

SetSendThankYouMail sets SendThankYouMail field to given value.

### HasSendThankYouMail

`func (o *ListRequest) HasSendThankYouMail() bool`

HasSendThankYouMail returns a boolean if a field has been set.

### GetThankYouFromName

`func (o *ListRequest) GetThankYouFromName() string`

GetThankYouFromName returns the ThankYouFromName field if non-nil, zero value otherwise.

### GetThankYouFromNameOk

`func (o *ListRequest) GetThankYouFromNameOk() (*string, bool)`

GetThankYouFromNameOk returns a tuple with the ThankYouFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouFromName

`func (o *ListRequest) SetThankYouFromName(v string)`

SetThankYouFromName sets ThankYouFromName field to given value.

### HasThankYouFromName

`func (o *ListRequest) HasThankYouFromName() bool`

HasThankYouFromName returns a boolean if a field has been set.

### GetThankYouFromEmail

`func (o *ListRequest) GetThankYouFromEmail() string`

GetThankYouFromEmail returns the ThankYouFromEmail field if non-nil, zero value otherwise.

### GetThankYouFromEmailOk

`func (o *ListRequest) GetThankYouFromEmailOk() (*string, bool)`

GetThankYouFromEmailOk returns a tuple with the ThankYouFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouFromEmail

`func (o *ListRequest) SetThankYouFromEmail(v string)`

SetThankYouFromEmail sets ThankYouFromEmail field to given value.

### HasThankYouFromEmail

`func (o *ListRequest) HasThankYouFromEmail() bool`

HasThankYouFromEmail returns a boolean if a field has been set.

### GetThankYouMailSubject

`func (o *ListRequest) GetThankYouMailSubject() string`

GetThankYouMailSubject returns the ThankYouMailSubject field if non-nil, zero value otherwise.

### GetThankYouMailSubjectOk

`func (o *ListRequest) GetThankYouMailSubjectOk() (*string, bool)`

GetThankYouMailSubjectOk returns a tuple with the ThankYouMailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMailSubject

`func (o *ListRequest) SetThankYouMailSubject(v string)`

SetThankYouMailSubject sets ThankYouMailSubject field to given value.

### HasThankYouMailSubject

`func (o *ListRequest) HasThankYouMailSubject() bool`

HasThankYouMailSubject returns a boolean if a field has been set.

### GetThankYouMailMessage

`func (o *ListRequest) GetThankYouMailMessage() string`

GetThankYouMailMessage returns the ThankYouMailMessage field if non-nil, zero value otherwise.

### GetThankYouMailMessageOk

`func (o *ListRequest) GetThankYouMailMessageOk() (*string, bool)`

GetThankYouMailMessageOk returns a tuple with the ThankYouMailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouMailMessage

`func (o *ListRequest) SetThankYouMailMessage(v string)`

SetThankYouMailMessage sets ThankYouMailMessage field to given value.

### HasThankYouMailMessage

`func (o *ListRequest) HasThankYouMailMessage() bool`

HasThankYouMailMessage returns a boolean if a field has been set.

### GetThankYouSender

`func (o *ListRequest) GetThankYouSender() string`

GetThankYouSender returns the ThankYouSender field if non-nil, zero value otherwise.

### GetThankYouSenderOk

`func (o *ListRequest) GetThankYouSenderOk() (*string, bool)`

GetThankYouSenderOk returns a tuple with the ThankYouSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouSender

`func (o *ListRequest) SetThankYouSender(v string)`

SetThankYouSender sets ThankYouSender field to given value.

### HasThankYouSender

`func (o *ListRequest) HasThankYouSender() bool`

HasThankYouSender returns a boolean if a field has been set.

### GetConfirmFromName

`func (o *ListRequest) GetConfirmFromName() string`

GetConfirmFromName returns the ConfirmFromName field if non-nil, zero value otherwise.

### GetConfirmFromNameOk

`func (o *ListRequest) GetConfirmFromNameOk() (*string, bool)`

GetConfirmFromNameOk returns a tuple with the ConfirmFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmFromName

`func (o *ListRequest) SetConfirmFromName(v string)`

SetConfirmFromName sets ConfirmFromName field to given value.

### HasConfirmFromName

`func (o *ListRequest) HasConfirmFromName() bool`

HasConfirmFromName returns a boolean if a field has been set.

### GetConfirmFromEmail

`func (o *ListRequest) GetConfirmFromEmail() string`

GetConfirmFromEmail returns the ConfirmFromEmail field if non-nil, zero value otherwise.

### GetConfirmFromEmailOk

`func (o *ListRequest) GetConfirmFromEmailOk() (*string, bool)`

GetConfirmFromEmailOk returns a tuple with the ConfirmFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmFromEmail

`func (o *ListRequest) SetConfirmFromEmail(v string)`

SetConfirmFromEmail sets ConfirmFromEmail field to given value.

### HasConfirmFromEmail

`func (o *ListRequest) HasConfirmFromEmail() bool`

HasConfirmFromEmail returns a boolean if a field has been set.

### GetConfirmMailSubject

`func (o *ListRequest) GetConfirmMailSubject() string`

GetConfirmMailSubject returns the ConfirmMailSubject field if non-nil, zero value otherwise.

### GetConfirmMailSubjectOk

`func (o *ListRequest) GetConfirmMailSubjectOk() (*string, bool)`

GetConfirmMailSubjectOk returns a tuple with the ConfirmMailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmMailSubject

`func (o *ListRequest) SetConfirmMailSubject(v string)`

SetConfirmMailSubject sets ConfirmMailSubject field to given value.

### HasConfirmMailSubject

`func (o *ListRequest) HasConfirmMailSubject() bool`

HasConfirmMailSubject returns a boolean if a field has been set.

### GetConfirmMailMessage

`func (o *ListRequest) GetConfirmMailMessage() string`

GetConfirmMailMessage returns the ConfirmMailMessage field if non-nil, zero value otherwise.

### GetConfirmMailMessageOk

`func (o *ListRequest) GetConfirmMailMessageOk() (*string, bool)`

GetConfirmMailMessageOk returns a tuple with the ConfirmMailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmMailMessage

`func (o *ListRequest) SetConfirmMailMessage(v string)`

SetConfirmMailMessage sets ConfirmMailMessage field to given value.

### HasConfirmMailMessage

`func (o *ListRequest) HasConfirmMailMessage() bool`

HasConfirmMailMessage returns a boolean if a field has been set.

### GetConfirmSuccessPage

`func (o *ListRequest) GetConfirmSuccessPage() string`

GetConfirmSuccessPage returns the ConfirmSuccessPage field if non-nil, zero value otherwise.

### GetConfirmSuccessPageOk

`func (o *ListRequest) GetConfirmSuccessPageOk() (*string, bool)`

GetConfirmSuccessPageOk returns a tuple with the ConfirmSuccessPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmSuccessPage

`func (o *ListRequest) SetConfirmSuccessPage(v string)`

SetConfirmSuccessPage sets ConfirmSuccessPage field to given value.

### HasConfirmSuccessPage

`func (o *ListRequest) HasConfirmSuccessPage() bool`

HasConfirmSuccessPage returns a boolean if a field has been set.

### GetConfirmSender

`func (o *ListRequest) GetConfirmSender() string`

GetConfirmSender returns the ConfirmSender field if non-nil, zero value otherwise.

### GetConfirmSenderOk

`func (o *ListRequest) GetConfirmSenderOk() (*string, bool)`

GetConfirmSenderOk returns a tuple with the ConfirmSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmSender

`func (o *ListRequest) SetConfirmSender(v string)`

SetConfirmSender sets ConfirmSender field to given value.

### HasConfirmSender

`func (o *ListRequest) HasConfirmSender() bool`

HasConfirmSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


