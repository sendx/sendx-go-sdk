# DeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteContacts** | Pointer to **bool** | Indicates whether to delete associated contacts. | [optional] 

## Methods

### NewDeleteRequest

`func NewDeleteRequest() *DeleteRequest`

NewDeleteRequest instantiates a new DeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRequestWithDefaults

`func NewDeleteRequestWithDefaults() *DeleteRequest`

NewDeleteRequestWithDefaults instantiates a new DeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteContacts

`func (o *DeleteRequest) GetDeleteContacts() bool`

GetDeleteContacts returns the DeleteContacts field if non-nil, zero value otherwise.

### GetDeleteContactsOk

`func (o *DeleteRequest) GetDeleteContactsOk() (*bool, bool)`

GetDeleteContactsOk returns a tuple with the DeleteContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteContacts

`func (o *DeleteRequest) SetDeleteContacts(v bool)`

SetDeleteContacts sets DeleteContacts field to given value.

### HasDeleteContacts

`func (o *DeleteRequest) HasDeleteContacts() bool`

HasDeleteContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


