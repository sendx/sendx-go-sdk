# DashboardStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastEmailCampaignStat** | Pointer to [**LastSentCampaignStat**](LastSentCampaignStat.md) |  | [optional] 
**NewsletterCount** | Pointer to **int32** | Number of newsletters sent. | [optional] 
**AutomationCount** | Pointer to **int32** | Number of automations set up. | [optional] 

## Methods

### NewDashboardStats

`func NewDashboardStats() *DashboardStats`

NewDashboardStats instantiates a new DashboardStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardStatsWithDefaults

`func NewDashboardStatsWithDefaults() *DashboardStats`

NewDashboardStatsWithDefaults instantiates a new DashboardStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastEmailCampaignStat

`func (o *DashboardStats) GetLastEmailCampaignStat() LastSentCampaignStat`

GetLastEmailCampaignStat returns the LastEmailCampaignStat field if non-nil, zero value otherwise.

### GetLastEmailCampaignStatOk

`func (o *DashboardStats) GetLastEmailCampaignStatOk() (*LastSentCampaignStat, bool)`

GetLastEmailCampaignStatOk returns a tuple with the LastEmailCampaignStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEmailCampaignStat

`func (o *DashboardStats) SetLastEmailCampaignStat(v LastSentCampaignStat)`

SetLastEmailCampaignStat sets LastEmailCampaignStat field to given value.

### HasLastEmailCampaignStat

`func (o *DashboardStats) HasLastEmailCampaignStat() bool`

HasLastEmailCampaignStat returns a boolean if a field has been set.

### GetNewsletterCount

`func (o *DashboardStats) GetNewsletterCount() int32`

GetNewsletterCount returns the NewsletterCount field if non-nil, zero value otherwise.

### GetNewsletterCountOk

`func (o *DashboardStats) GetNewsletterCountOk() (*int32, bool)`

GetNewsletterCountOk returns a tuple with the NewsletterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletterCount

`func (o *DashboardStats) SetNewsletterCount(v int32)`

SetNewsletterCount sets NewsletterCount field to given value.

### HasNewsletterCount

`func (o *DashboardStats) HasNewsletterCount() bool`

HasNewsletterCount returns a boolean if a field has been set.

### GetAutomationCount

`func (o *DashboardStats) GetAutomationCount() int32`

GetAutomationCount returns the AutomationCount field if non-nil, zero value otherwise.

### GetAutomationCountOk

`func (o *DashboardStats) GetAutomationCountOk() (*int32, bool)`

GetAutomationCountOk returns a tuple with the AutomationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationCount

`func (o *DashboardStats) SetAutomationCount(v int32)`

SetAutomationCount sets AutomationCount field to given value.

### HasAutomationCount

`func (o *DashboardStats) HasAutomationCount() bool`

HasAutomationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


