# RevenueEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Unique identifier for the contact. | 
**Amount** | **float32** | Recognized revenue amount associated with the event. | 
**Currency** | **string** | Currency code (ISO 4217) for the revenue (e.g., &#39;USD&#39;, &#39;EUR&#39;, &#39;INR&#39;). | 
**Source** | **string** | Source of the revenue (e.g., &#39;website&#39;, &#39;mobile app&#39;, &#39;partner referral&#39;). | 
**Time** | **int32** | Unix timestamp indicating when the revenue event occurred. | 

## Methods

### NewRevenueEventRequest

`func NewRevenueEventRequest(identifier string, amount float32, currency string, source string, time int32, ) *RevenueEventRequest`

NewRevenueEventRequest instantiates a new RevenueEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueEventRequestWithDefaults

`func NewRevenueEventRequestWithDefaults() *RevenueEventRequest`

NewRevenueEventRequestWithDefaults instantiates a new RevenueEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *RevenueEventRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RevenueEventRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RevenueEventRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetAmount

`func (o *RevenueEventRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RevenueEventRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RevenueEventRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *RevenueEventRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RevenueEventRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RevenueEventRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSource

`func (o *RevenueEventRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RevenueEventRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RevenueEventRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetTime

`func (o *RevenueEventRequest) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RevenueEventRequest) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RevenueEventRequest) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


