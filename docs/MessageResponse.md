# MessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageResponse

`func NewMessageResponse() *MessageResponse`

NewMessageResponse instantiates a new MessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseWithDefaults

`func NewMessageResponseWithDefaults() *MessageResponse`

NewMessageResponseWithDefaults instantiates a new MessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *MessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


