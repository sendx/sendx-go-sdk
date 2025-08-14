# RestECustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Custom field name (must be unique within team) | 
**Type** | **int32** | Field data type.  **Values:** - &#x60;0&#x60; - Text (max 255 characters) - &#x60;1&#x60; - Number (integer or decimal) - &#x60;2&#x60; - Date (YYYY-MM-DD format) - &#x60;3&#x60; - Boolean (true/false) - &#x60;4&#x60; - Phone number (international format)  | 
**Description** | Pointer to **string** | Field description for documentation | [optional] 

## Methods

### NewRestECustomField

`func NewRestECustomField(name string, type_ int32, ) *RestECustomField`

NewRestECustomField instantiates a new RestECustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestECustomFieldWithDefaults

`func NewRestECustomFieldWithDefaults() *RestECustomField`

NewRestECustomFieldWithDefaults instantiates a new RestECustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestECustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestECustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestECustomField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RestECustomField) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestECustomField) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestECustomField) SetType(v int32)`

SetType sets Type field to given value.


### GetDescription

`func (o *RestECustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestECustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestECustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestECustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


