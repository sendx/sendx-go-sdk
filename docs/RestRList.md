# RestRList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | List type.  **Values:** - &#x60;0&#x60; - Regular list (single opt-in) - &#x60;1&#x60; - Double opt-in list (requires email confirmation)  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRestRList

`func NewRestRList() *RestRList`

NewRestRList instantiates a new RestRList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRListWithDefaults

`func NewRestRListWithDefaults() *RestRList`

NewRestRListWithDefaults instantiates a new RestRList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RestRList) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestRList) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestRList) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestRList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *RestRList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRList) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRList) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRList) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRList) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


