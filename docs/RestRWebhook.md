# RestRWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Webhook endpoint URL | [optional] 
**Enabled** | Pointer to **bool** | Whether webhook is enabled | [optional] [default to true]
**Unsubscribed** | Pointer to **bool** | Trigger webhook when a contact unsubscribes | [optional] [default to false]
**Dropped** | Pointer to **bool** | Trigger webhook when an email is dropped | [optional] [default to false]
**Bounced** | Pointer to **bool** | Trigger webhook when an email bounces | [optional] [default to false]
**MarkedSpam** | Pointer to **bool** | Trigger webhook when an email is marked as spam | [optional] [default to false]
**Clicked** | Pointer to **bool** | Trigger webhook when a link in the email is clicked | [optional] [default to false]
**Opened** | Pointer to **bool** | Trigger webhook when an email is opened | [optional] [default to false]
**ContactCreated** | Pointer to **bool** | Trigger webhook when a new contact is created | [optional] [default to false]

## Methods

### NewRestRWebhook

`func NewRestRWebhook() *RestRWebhook`

NewRestRWebhook instantiates a new RestRWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRWebhookWithDefaults

`func NewRestRWebhookWithDefaults() *RestRWebhook`

NewRestRWebhookWithDefaults instantiates a new RestRWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *RestRWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RestRWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RestRWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RestRWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *RestRWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RestRWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RestRWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RestRWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *RestRWebhook) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *RestRWebhook) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *RestRWebhook) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *RestRWebhook) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetDropped

`func (o *RestRWebhook) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *RestRWebhook) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *RestRWebhook) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *RestRWebhook) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetBounced

`func (o *RestRWebhook) GetBounced() bool`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *RestRWebhook) GetBouncedOk() (*bool, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *RestRWebhook) SetBounced(v bool)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *RestRWebhook) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetMarkedSpam

`func (o *RestRWebhook) GetMarkedSpam() bool`

GetMarkedSpam returns the MarkedSpam field if non-nil, zero value otherwise.

### GetMarkedSpamOk

`func (o *RestRWebhook) GetMarkedSpamOk() (*bool, bool)`

GetMarkedSpamOk returns a tuple with the MarkedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedSpam

`func (o *RestRWebhook) SetMarkedSpam(v bool)`

SetMarkedSpam sets MarkedSpam field to given value.

### HasMarkedSpam

`func (o *RestRWebhook) HasMarkedSpam() bool`

HasMarkedSpam returns a boolean if a field has been set.

### GetClicked

`func (o *RestRWebhook) GetClicked() bool`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *RestRWebhook) GetClickedOk() (*bool, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *RestRWebhook) SetClicked(v bool)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *RestRWebhook) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetOpened

`func (o *RestRWebhook) GetOpened() bool`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *RestRWebhook) GetOpenedOk() (*bool, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *RestRWebhook) SetOpened(v bool)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *RestRWebhook) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetContactCreated

`func (o *RestRWebhook) GetContactCreated() bool`

GetContactCreated returns the ContactCreated field if non-nil, zero value otherwise.

### GetContactCreatedOk

`func (o *RestRWebhook) GetContactCreatedOk() (*bool, bool)`

GetContactCreatedOk returns a tuple with the ContactCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCreated

`func (o *RestRWebhook) SetContactCreated(v bool)`

SetContactCreated sets ContactCreated field to given value.

### HasContactCreated

`func (o *RestRWebhook) HasContactCreated() bool`

HasContactCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


