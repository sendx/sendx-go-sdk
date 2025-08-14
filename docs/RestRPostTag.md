# RestRPostTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Tag name | [optional] 
**Created** | Pointer to **time.Time** | Date and time when tag was created | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when tag was updated | [optional] 

## Methods

### NewRestRPostTag

`func NewRestRPostTag() *RestRPostTag`

NewRestRPostTag instantiates a new RestRPostTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRPostTagWithDefaults

`func NewRestRPostTagWithDefaults() *RestRPostTag`

NewRestRPostTagWithDefaults instantiates a new RestRPostTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRPostTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRPostTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRPostTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRPostTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRPostTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRPostTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRPostTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRPostTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *RestRPostTag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRPostTag) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRPostTag) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRPostTag) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRPostTag) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRPostTag) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRPostTag) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRPostTag) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


