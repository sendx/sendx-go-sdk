# LinkStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The link clicked | [optional] 
**Count** | Pointer to **int32** | Total number of times the link was clicked | [optional] 

## Methods

### NewLinkStat

`func NewLinkStat() *LinkStat`

NewLinkStat instantiates a new LinkStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkStatWithDefaults

`func NewLinkStatWithDefaults() *LinkStat`

NewLinkStatWithDefaults instantiates a new LinkStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LinkStat) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkStat) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkStat) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LinkStat) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCount

`func (o *LinkStat) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LinkStat) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LinkStat) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LinkStat) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


