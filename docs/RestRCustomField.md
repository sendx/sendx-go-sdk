# RestRCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique field identifier with field_ prefix | [optional] 
**Name** | Pointer to **string** | Custom field name | [optional] 
**Type** | Pointer to **int32** | Field data type.  **Values:** - &#x60;0&#x60; - Text (max 255 characters) - &#x60;1&#x60; - Number (integer or decimal) - &#x60;2&#x60; - Date (YYYY-MM-DD format) - &#x60;3&#x60; - Boolean (true/false) - &#x60;4&#x60; - Phone number (international format)  | [optional] 
**Description** | Pointer to **string** | Field description for documentation | [optional] 

## Methods

### NewRestRCustomField

`func NewRestRCustomField() *RestRCustomField`

NewRestRCustomField instantiates a new RestRCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRCustomFieldWithDefaults

`func NewRestRCustomFieldWithDefaults() *RestRCustomField`

NewRestRCustomFieldWithDefaults instantiates a new RestRCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRCustomField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRCustomField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRCustomField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRCustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RestRCustomField) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestRCustomField) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestRCustomField) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestRCustomField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *RestRCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestRCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestRCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestRCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


