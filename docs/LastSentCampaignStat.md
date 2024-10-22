# LastSentCampaignStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | Pointer to [**CampaignDashboardData**](CampaignDashboardData.md) |  | [optional] 
**Sent** | Pointer to **int32** | Number of emails sent. | [optional] 
**Delivered** | Pointer to **int32** | Number of emails delivered. | [optional] 
**Subscribed** | Pointer to **int32** | Number of new subscriptions. | [optional] 
**Unsubscribed** | Pointer to **int32** | Number of unsubscribes. | [optional] 
**Open** | Pointer to **int32** | Number of emails opened. | [optional] 
**LinkClicked** | Pointer to **int32** | Number of link clicks. | [optional] 
**Replied** | Pointer to **int32** | Number of replies received. | [optional] 

## Methods

### NewLastSentCampaignStat

`func NewLastSentCampaignStat() *LastSentCampaignStat`

NewLastSentCampaignStat instantiates a new LastSentCampaignStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastSentCampaignStatWithDefaults

`func NewLastSentCampaignStatWithDefaults() *LastSentCampaignStat`

NewLastSentCampaignStatWithDefaults instantiates a new LastSentCampaignStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *LastSentCampaignStat) GetCampaign() CampaignDashboardData`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *LastSentCampaignStat) GetCampaignOk() (*CampaignDashboardData, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *LastSentCampaignStat) SetCampaign(v CampaignDashboardData)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *LastSentCampaignStat) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetSent

`func (o *LastSentCampaignStat) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *LastSentCampaignStat) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *LastSentCampaignStat) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *LastSentCampaignStat) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *LastSentCampaignStat) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *LastSentCampaignStat) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *LastSentCampaignStat) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *LastSentCampaignStat) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetSubscribed

`func (o *LastSentCampaignStat) GetSubscribed() int32`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *LastSentCampaignStat) GetSubscribedOk() (*int32, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *LastSentCampaignStat) SetSubscribed(v int32)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *LastSentCampaignStat) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *LastSentCampaignStat) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *LastSentCampaignStat) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *LastSentCampaignStat) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *LastSentCampaignStat) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetOpen

`func (o *LastSentCampaignStat) GetOpen() int32`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *LastSentCampaignStat) GetOpenOk() (*int32, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *LastSentCampaignStat) SetOpen(v int32)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *LastSentCampaignStat) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetLinkClicked

`func (o *LastSentCampaignStat) GetLinkClicked() int32`

GetLinkClicked returns the LinkClicked field if non-nil, zero value otherwise.

### GetLinkClickedOk

`func (o *LastSentCampaignStat) GetLinkClickedOk() (*int32, bool)`

GetLinkClickedOk returns a tuple with the LinkClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkClicked

`func (o *LastSentCampaignStat) SetLinkClicked(v int32)`

SetLinkClicked sets LinkClicked field to given value.

### HasLinkClicked

`func (o *LastSentCampaignStat) HasLinkClicked() bool`

HasLinkClicked returns a boolean if a field has been set.

### GetReplied

`func (o *LastSentCampaignStat) GetReplied() int32`

GetReplied returns the Replied field if non-nil, zero value otherwise.

### GetRepliedOk

`func (o *LastSentCampaignStat) GetRepliedOk() (*int32, bool)`

GetRepliedOk returns a tuple with the Replied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplied

`func (o *LastSentCampaignStat) SetReplied(v int32)`

SetReplied sets Replied field to given value.

### HasReplied

`func (o *LastSentCampaignStat) HasReplied() bool`

HasReplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


