# RestRPostCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id for the post category | [optional] 
**Name** | Pointer to **string** | Name for the post category | [optional] 
**Created** | Pointer to **time.Time** | Date and time when the post category was created | [optional] 
**Updated** | Pointer to **time.Time** | Date and time when the post category was last updated | [optional] 

## Methods

### NewRestRPostCategory

`func NewRestRPostCategory() *RestRPostCategory`

NewRestRPostCategory instantiates a new RestRPostCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRPostCategoryWithDefaults

`func NewRestRPostCategoryWithDefaults() *RestRPostCategory`

NewRestRPostCategoryWithDefaults instantiates a new RestRPostCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRPostCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRPostCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRPostCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRPostCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRPostCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRPostCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRPostCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRPostCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *RestRPostCategory) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRPostCategory) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRPostCategory) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRPostCategory) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRPostCategory) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRPostCategory) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRPostCategory) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRPostCategory) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


