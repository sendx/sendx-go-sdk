# RestReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **string** | Unique identifier of the campaign | [optional] 
**LinkStats** | Pointer to [**[]LinkStat**](LinkStat.md) |  | [optional] 
**ClickedUniqueContactCount** | Pointer to **int32** | Total number of unique contacts who clicked on the link | [optional] 
**OpenedUniqueContactCount** | Pointer to **int32** | Total number of unique contacts who opened the link | [optional] 
**SentContactCount** | Pointer to **int32** | Total number of contacts who sent the link | [optional] 
**UnsubscribeContactCount** | Pointer to **int32** | Total number of contacts who unsubscribed from the link | [optional] 
**BounceContactCount** | Pointer to **int32** | Total number of contacts who bounced the link | [optional] 
**SpamContactCount** | Pointer to **int32** | Total number of contacts who marked the link as spam | [optional] 
**ClickedContactCount** | Pointer to **int32** | Total number of contacts who clicked on the link | [optional] 
**OpenedContactCount** | Pointer to **int32** | Total number of contacts who opened the link | [optional] 

## Methods

### NewRestReportData

`func NewRestReportData() *RestReportData`

NewRestReportData instantiates a new RestReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestReportDataWithDefaults

`func NewRestReportDataWithDefaults() *RestReportData`

NewRestReportDataWithDefaults instantiates a new RestReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *RestReportData) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *RestReportData) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *RestReportData) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *RestReportData) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetLinkStats

`func (o *RestReportData) GetLinkStats() []LinkStat`

GetLinkStats returns the LinkStats field if non-nil, zero value otherwise.

### GetLinkStatsOk

`func (o *RestReportData) GetLinkStatsOk() (*[]LinkStat, bool)`

GetLinkStatsOk returns a tuple with the LinkStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStats

`func (o *RestReportData) SetLinkStats(v []LinkStat)`

SetLinkStats sets LinkStats field to given value.

### HasLinkStats

`func (o *RestReportData) HasLinkStats() bool`

HasLinkStats returns a boolean if a field has been set.

### GetClickedUniqueContactCount

`func (o *RestReportData) GetClickedUniqueContactCount() int32`

GetClickedUniqueContactCount returns the ClickedUniqueContactCount field if non-nil, zero value otherwise.

### GetClickedUniqueContactCountOk

`func (o *RestReportData) GetClickedUniqueContactCountOk() (*int32, bool)`

GetClickedUniqueContactCountOk returns a tuple with the ClickedUniqueContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedUniqueContactCount

`func (o *RestReportData) SetClickedUniqueContactCount(v int32)`

SetClickedUniqueContactCount sets ClickedUniqueContactCount field to given value.

### HasClickedUniqueContactCount

`func (o *RestReportData) HasClickedUniqueContactCount() bool`

HasClickedUniqueContactCount returns a boolean if a field has been set.

### GetOpenedUniqueContactCount

`func (o *RestReportData) GetOpenedUniqueContactCount() int32`

GetOpenedUniqueContactCount returns the OpenedUniqueContactCount field if non-nil, zero value otherwise.

### GetOpenedUniqueContactCountOk

`func (o *RestReportData) GetOpenedUniqueContactCountOk() (*int32, bool)`

GetOpenedUniqueContactCountOk returns a tuple with the OpenedUniqueContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedUniqueContactCount

`func (o *RestReportData) SetOpenedUniqueContactCount(v int32)`

SetOpenedUniqueContactCount sets OpenedUniqueContactCount field to given value.

### HasOpenedUniqueContactCount

`func (o *RestReportData) HasOpenedUniqueContactCount() bool`

HasOpenedUniqueContactCount returns a boolean if a field has been set.

### GetSentContactCount

`func (o *RestReportData) GetSentContactCount() int32`

GetSentContactCount returns the SentContactCount field if non-nil, zero value otherwise.

### GetSentContactCountOk

`func (o *RestReportData) GetSentContactCountOk() (*int32, bool)`

GetSentContactCountOk returns a tuple with the SentContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentContactCount

`func (o *RestReportData) SetSentContactCount(v int32)`

SetSentContactCount sets SentContactCount field to given value.

### HasSentContactCount

`func (o *RestReportData) HasSentContactCount() bool`

HasSentContactCount returns a boolean if a field has been set.

### GetUnsubscribeContactCount

`func (o *RestReportData) GetUnsubscribeContactCount() int32`

GetUnsubscribeContactCount returns the UnsubscribeContactCount field if non-nil, zero value otherwise.

### GetUnsubscribeContactCountOk

`func (o *RestReportData) GetUnsubscribeContactCountOk() (*int32, bool)`

GetUnsubscribeContactCountOk returns a tuple with the UnsubscribeContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeContactCount

`func (o *RestReportData) SetUnsubscribeContactCount(v int32)`

SetUnsubscribeContactCount sets UnsubscribeContactCount field to given value.

### HasUnsubscribeContactCount

`func (o *RestReportData) HasUnsubscribeContactCount() bool`

HasUnsubscribeContactCount returns a boolean if a field has been set.

### GetBounceContactCount

`func (o *RestReportData) GetBounceContactCount() int32`

GetBounceContactCount returns the BounceContactCount field if non-nil, zero value otherwise.

### GetBounceContactCountOk

`func (o *RestReportData) GetBounceContactCountOk() (*int32, bool)`

GetBounceContactCountOk returns a tuple with the BounceContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceContactCount

`func (o *RestReportData) SetBounceContactCount(v int32)`

SetBounceContactCount sets BounceContactCount field to given value.

### HasBounceContactCount

`func (o *RestReportData) HasBounceContactCount() bool`

HasBounceContactCount returns a boolean if a field has been set.

### GetSpamContactCount

`func (o *RestReportData) GetSpamContactCount() int32`

GetSpamContactCount returns the SpamContactCount field if non-nil, zero value otherwise.

### GetSpamContactCountOk

`func (o *RestReportData) GetSpamContactCountOk() (*int32, bool)`

GetSpamContactCountOk returns a tuple with the SpamContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamContactCount

`func (o *RestReportData) SetSpamContactCount(v int32)`

SetSpamContactCount sets SpamContactCount field to given value.

### HasSpamContactCount

`func (o *RestReportData) HasSpamContactCount() bool`

HasSpamContactCount returns a boolean if a field has been set.

### GetClickedContactCount

`func (o *RestReportData) GetClickedContactCount() int32`

GetClickedContactCount returns the ClickedContactCount field if non-nil, zero value otherwise.

### GetClickedContactCountOk

`func (o *RestReportData) GetClickedContactCountOk() (*int32, bool)`

GetClickedContactCountOk returns a tuple with the ClickedContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedContactCount

`func (o *RestReportData) SetClickedContactCount(v int32)`

SetClickedContactCount sets ClickedContactCount field to given value.

### HasClickedContactCount

`func (o *RestReportData) HasClickedContactCount() bool`

HasClickedContactCount returns a boolean if a field has been set.

### GetOpenedContactCount

`func (o *RestReportData) GetOpenedContactCount() int32`

GetOpenedContactCount returns the OpenedContactCount field if non-nil, zero value otherwise.

### GetOpenedContactCountOk

`func (o *RestReportData) GetOpenedContactCountOk() (*int32, bool)`

GetOpenedContactCountOk returns a tuple with the OpenedContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedContactCount

`func (o *RestReportData) SetOpenedContactCount(v int32)`

SetOpenedContactCount sets OpenedContactCount field to given value.

### HasOpenedContactCount

`func (o *RestReportData) HasOpenedContactCount() bool`

HasOpenedContactCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


