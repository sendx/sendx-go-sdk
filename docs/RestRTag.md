# RestRTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique tag identifier with tag_ prefix | [optional] 
**Name** | Pointer to **string** | Tag name | [optional] 
**Created** | Pointer to **time.Time** | Tag creation timestamp | [optional] 
**Updated** | Pointer to **time.Time** | Tag last update timestamp | [optional] 

## Methods

### NewRestRTag

`func NewRestRTag() *RestRTag`

NewRestRTag instantiates a new RestRTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRTagWithDefaults

`func NewRestRTagWithDefaults() *RestRTag`

NewRestRTagWithDefaults instantiates a new RestRTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *RestRTag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRTag) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRTag) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRTag) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRTag) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRTag) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRTag) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRTag) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


