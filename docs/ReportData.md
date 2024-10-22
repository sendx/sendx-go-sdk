# ReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **string** | The ID of the campaign | [optional] 
**LinkStats** | Pointer to **map[string]int32** | Statistics about the links clicked within the campaign | [optional] 
**ClickedContactCount** | Pointer to **int32** | The total number of contacts that clicked on any link | [optional] 
**OpenedContactCount** | Pointer to **int32** | The total number of contacts that opened the campaign email | [optional] 
**RepliedContactCount** | Pointer to **int32** | The total number of contacts that replied to the campaign email | [optional] 
**ClickedUniqueContactCount** | Pointer to **int32** | The unique number of contacts that clicked on any link | [optional] 
**OpenedUniqueContactCount** | Pointer to **int32** | The unique number of contacts that opened the campaign email | [optional] 
**RepliedUniqueContactCount** | Pointer to **int32** | The unique number of contacts that replied to the campaign email | [optional] 
**SentContactCount** | Pointer to **int32** | The total number of contacts the campaign was sent to | [optional] 
**UnsubscribeContactCount** | Pointer to **int32** | The total number of contacts that unsubscribed | [optional] 
**BounceContactCount** | Pointer to **int32** | The total number of bounced contacts | [optional] 
**SpamContactCount** | Pointer to **int32** | The total number of contacts that marked the email as spam | [optional] 
**EmailRevenue** | Pointer to **string** | The total revenue generated from the campaign email | [optional] 
**RevenuePerRecipient** | Pointer to **string** | The average revenue generated per recipient | [optional] 
**Currency** | Pointer to **string** | The currency in which the revenue is measured | [optional] 

## Methods

### NewReportData

`func NewReportData() *ReportData`

NewReportData instantiates a new ReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportDataWithDefaults

`func NewReportDataWithDefaults() *ReportData`

NewReportDataWithDefaults instantiates a new ReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *ReportData) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ReportData) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ReportData) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ReportData) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetLinkStats

`func (o *ReportData) GetLinkStats() map[string]int32`

GetLinkStats returns the LinkStats field if non-nil, zero value otherwise.

### GetLinkStatsOk

`func (o *ReportData) GetLinkStatsOk() (*map[string]int32, bool)`

GetLinkStatsOk returns a tuple with the LinkStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStats

`func (o *ReportData) SetLinkStats(v map[string]int32)`

SetLinkStats sets LinkStats field to given value.

### HasLinkStats

`func (o *ReportData) HasLinkStats() bool`

HasLinkStats returns a boolean if a field has been set.

### GetClickedContactCount

`func (o *ReportData) GetClickedContactCount() int32`

GetClickedContactCount returns the ClickedContactCount field if non-nil, zero value otherwise.

### GetClickedContactCountOk

`func (o *ReportData) GetClickedContactCountOk() (*int32, bool)`

GetClickedContactCountOk returns a tuple with the ClickedContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedContactCount

`func (o *ReportData) SetClickedContactCount(v int32)`

SetClickedContactCount sets ClickedContactCount field to given value.

### HasClickedContactCount

`func (o *ReportData) HasClickedContactCount() bool`

HasClickedContactCount returns a boolean if a field has been set.

### GetOpenedContactCount

`func (o *ReportData) GetOpenedContactCount() int32`

GetOpenedContactCount returns the OpenedContactCount field if non-nil, zero value otherwise.

### GetOpenedContactCountOk

`func (o *ReportData) GetOpenedContactCountOk() (*int32, bool)`

GetOpenedContactCountOk returns a tuple with the OpenedContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedContactCount

`func (o *ReportData) SetOpenedContactCount(v int32)`

SetOpenedContactCount sets OpenedContactCount field to given value.

### HasOpenedContactCount

`func (o *ReportData) HasOpenedContactCount() bool`

HasOpenedContactCount returns a boolean if a field has been set.

### GetRepliedContactCount

`func (o *ReportData) GetRepliedContactCount() int32`

GetRepliedContactCount returns the RepliedContactCount field if non-nil, zero value otherwise.

### GetRepliedContactCountOk

`func (o *ReportData) GetRepliedContactCountOk() (*int32, bool)`

GetRepliedContactCountOk returns a tuple with the RepliedContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedContactCount

`func (o *ReportData) SetRepliedContactCount(v int32)`

SetRepliedContactCount sets RepliedContactCount field to given value.

### HasRepliedContactCount

`func (o *ReportData) HasRepliedContactCount() bool`

HasRepliedContactCount returns a boolean if a field has been set.

### GetClickedUniqueContactCount

`func (o *ReportData) GetClickedUniqueContactCount() int32`

GetClickedUniqueContactCount returns the ClickedUniqueContactCount field if non-nil, zero value otherwise.

### GetClickedUniqueContactCountOk

`func (o *ReportData) GetClickedUniqueContactCountOk() (*int32, bool)`

GetClickedUniqueContactCountOk returns a tuple with the ClickedUniqueContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedUniqueContactCount

`func (o *ReportData) SetClickedUniqueContactCount(v int32)`

SetClickedUniqueContactCount sets ClickedUniqueContactCount field to given value.

### HasClickedUniqueContactCount

`func (o *ReportData) HasClickedUniqueContactCount() bool`

HasClickedUniqueContactCount returns a boolean if a field has been set.

### GetOpenedUniqueContactCount

`func (o *ReportData) GetOpenedUniqueContactCount() int32`

GetOpenedUniqueContactCount returns the OpenedUniqueContactCount field if non-nil, zero value otherwise.

### GetOpenedUniqueContactCountOk

`func (o *ReportData) GetOpenedUniqueContactCountOk() (*int32, bool)`

GetOpenedUniqueContactCountOk returns a tuple with the OpenedUniqueContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedUniqueContactCount

`func (o *ReportData) SetOpenedUniqueContactCount(v int32)`

SetOpenedUniqueContactCount sets OpenedUniqueContactCount field to given value.

### HasOpenedUniqueContactCount

`func (o *ReportData) HasOpenedUniqueContactCount() bool`

HasOpenedUniqueContactCount returns a boolean if a field has been set.

### GetRepliedUniqueContactCount

`func (o *ReportData) GetRepliedUniqueContactCount() int32`

GetRepliedUniqueContactCount returns the RepliedUniqueContactCount field if non-nil, zero value otherwise.

### GetRepliedUniqueContactCountOk

`func (o *ReportData) GetRepliedUniqueContactCountOk() (*int32, bool)`

GetRepliedUniqueContactCountOk returns a tuple with the RepliedUniqueContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedUniqueContactCount

`func (o *ReportData) SetRepliedUniqueContactCount(v int32)`

SetRepliedUniqueContactCount sets RepliedUniqueContactCount field to given value.

### HasRepliedUniqueContactCount

`func (o *ReportData) HasRepliedUniqueContactCount() bool`

HasRepliedUniqueContactCount returns a boolean if a field has been set.

### GetSentContactCount

`func (o *ReportData) GetSentContactCount() int32`

GetSentContactCount returns the SentContactCount field if non-nil, zero value otherwise.

### GetSentContactCountOk

`func (o *ReportData) GetSentContactCountOk() (*int32, bool)`

GetSentContactCountOk returns a tuple with the SentContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentContactCount

`func (o *ReportData) SetSentContactCount(v int32)`

SetSentContactCount sets SentContactCount field to given value.

### HasSentContactCount

`func (o *ReportData) HasSentContactCount() bool`

HasSentContactCount returns a boolean if a field has been set.

### GetUnsubscribeContactCount

`func (o *ReportData) GetUnsubscribeContactCount() int32`

GetUnsubscribeContactCount returns the UnsubscribeContactCount field if non-nil, zero value otherwise.

### GetUnsubscribeContactCountOk

`func (o *ReportData) GetUnsubscribeContactCountOk() (*int32, bool)`

GetUnsubscribeContactCountOk returns a tuple with the UnsubscribeContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeContactCount

`func (o *ReportData) SetUnsubscribeContactCount(v int32)`

SetUnsubscribeContactCount sets UnsubscribeContactCount field to given value.

### HasUnsubscribeContactCount

`func (o *ReportData) HasUnsubscribeContactCount() bool`

HasUnsubscribeContactCount returns a boolean if a field has been set.

### GetBounceContactCount

`func (o *ReportData) GetBounceContactCount() int32`

GetBounceContactCount returns the BounceContactCount field if non-nil, zero value otherwise.

### GetBounceContactCountOk

`func (o *ReportData) GetBounceContactCountOk() (*int32, bool)`

GetBounceContactCountOk returns a tuple with the BounceContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceContactCount

`func (o *ReportData) SetBounceContactCount(v int32)`

SetBounceContactCount sets BounceContactCount field to given value.

### HasBounceContactCount

`func (o *ReportData) HasBounceContactCount() bool`

HasBounceContactCount returns a boolean if a field has been set.

### GetSpamContactCount

`func (o *ReportData) GetSpamContactCount() int32`

GetSpamContactCount returns the SpamContactCount field if non-nil, zero value otherwise.

### GetSpamContactCountOk

`func (o *ReportData) GetSpamContactCountOk() (*int32, bool)`

GetSpamContactCountOk returns a tuple with the SpamContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamContactCount

`func (o *ReportData) SetSpamContactCount(v int32)`

SetSpamContactCount sets SpamContactCount field to given value.

### HasSpamContactCount

`func (o *ReportData) HasSpamContactCount() bool`

HasSpamContactCount returns a boolean if a field has been set.

### GetEmailRevenue

`func (o *ReportData) GetEmailRevenue() string`

GetEmailRevenue returns the EmailRevenue field if non-nil, zero value otherwise.

### GetEmailRevenueOk

`func (o *ReportData) GetEmailRevenueOk() (*string, bool)`

GetEmailRevenueOk returns a tuple with the EmailRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRevenue

`func (o *ReportData) SetEmailRevenue(v string)`

SetEmailRevenue sets EmailRevenue field to given value.

### HasEmailRevenue

`func (o *ReportData) HasEmailRevenue() bool`

HasEmailRevenue returns a boolean if a field has been set.

### GetRevenuePerRecipient

`func (o *ReportData) GetRevenuePerRecipient() string`

GetRevenuePerRecipient returns the RevenuePerRecipient field if non-nil, zero value otherwise.

### GetRevenuePerRecipientOk

`func (o *ReportData) GetRevenuePerRecipientOk() (*string, bool)`

GetRevenuePerRecipientOk returns a tuple with the RevenuePerRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerRecipient

`func (o *ReportData) SetRevenuePerRecipient(v string)`

SetRevenuePerRecipient sets RevenuePerRecipient field to given value.

### HasRevenuePerRecipient

`func (o *ReportData) HasRevenuePerRecipient() bool`

HasRevenuePerRecipient returns a boolean if a field has been set.

### GetCurrency

`func (o *ReportData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReportData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReportData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReportData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


