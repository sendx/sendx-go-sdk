# RestEWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Webhook endpoint URL | 
**Enabled** | **bool** | Whether webhook is enabled | [default to true]
**Unsubscribed** | Pointer to **bool** | Trigger webhook when a contact unsubscribes | [optional] [default to false]
**Dropped** | Pointer to **bool** | Trigger webhook when an email is dropped | [optional] [default to false]
**Bounced** | Pointer to **bool** | Trigger webhook when an email bounces | [optional] [default to false]
**MarkedSpam** | Pointer to **bool** | Trigger webhook when an email is marked as spam | [optional] [default to false]
**Clicked** | Pointer to **bool** | Trigger webhook when a link in the email is clicked | [optional] [default to false]
**Opened** | Pointer to **bool** | Trigger webhook when an email is opened | [optional] [default to false]
**ContactCreated** | Pointer to **bool** | Trigger webhook when a new contact is created | [optional] [default to false]

## Methods

### NewRestEWebhook

`func NewRestEWebhook(url string, enabled bool, ) *RestEWebhook`

NewRestEWebhook instantiates a new RestEWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestEWebhookWithDefaults

`func NewRestEWebhookWithDefaults() *RestEWebhook`

NewRestEWebhookWithDefaults instantiates a new RestEWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RestEWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RestEWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RestEWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *RestEWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RestEWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RestEWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUnsubscribed

`func (o *RestEWebhook) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *RestEWebhook) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *RestEWebhook) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *RestEWebhook) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetDropped

`func (o *RestEWebhook) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *RestEWebhook) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *RestEWebhook) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *RestEWebhook) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetBounced

`func (o *RestEWebhook) GetBounced() bool`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *RestEWebhook) GetBouncedOk() (*bool, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *RestEWebhook) SetBounced(v bool)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *RestEWebhook) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetMarkedSpam

`func (o *RestEWebhook) GetMarkedSpam() bool`

GetMarkedSpam returns the MarkedSpam field if non-nil, zero value otherwise.

### GetMarkedSpamOk

`func (o *RestEWebhook) GetMarkedSpamOk() (*bool, bool)`

GetMarkedSpamOk returns a tuple with the MarkedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedSpam

`func (o *RestEWebhook) SetMarkedSpam(v bool)`

SetMarkedSpam sets MarkedSpam field to given value.

### HasMarkedSpam

`func (o *RestEWebhook) HasMarkedSpam() bool`

HasMarkedSpam returns a boolean if a field has been set.

### GetClicked

`func (o *RestEWebhook) GetClicked() bool`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *RestEWebhook) GetClickedOk() (*bool, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *RestEWebhook) SetClicked(v bool)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *RestEWebhook) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetOpened

`func (o *RestEWebhook) GetOpened() bool`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *RestEWebhook) GetOpenedOk() (*bool, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *RestEWebhook) SetOpened(v bool)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *RestEWebhook) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetContactCreated

`func (o *RestEWebhook) GetContactCreated() bool`

GetContactCreated returns the ContactCreated field if non-nil, zero value otherwise.

### GetContactCreatedOk

`func (o *RestEWebhook) GetContactCreatedOk() (*bool, bool)`

GetContactCreatedOk returns a tuple with the ContactCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCreated

`func (o *RestEWebhook) SetContactCreated(v bool)`

SetContactCreated sets ContactCreated field to given value.

### HasContactCreated

`func (o *RestEWebhook) HasContactCreated() bool`

HasContactCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


