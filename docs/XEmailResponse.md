# XEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** | Recipient email address | [optional] 
**SubmittedAt** | Pointer to **int64** | Unix timestamp of submission | [optional] 
**MessageId** | Pointer to **string** | Unique message identifier | [optional] 
**ErrorCode** | Pointer to **int32** | Error code (0 &#x3D; success) | [optional] 
**Message** | Pointer to **string** | Status message | [optional] 

## Methods

### NewXEmailResponse

`func NewXEmailResponse() *XEmailResponse`

NewXEmailResponse instantiates a new XEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXEmailResponseWithDefaults

`func NewXEmailResponseWithDefaults() *XEmailResponse`

NewXEmailResponseWithDefaults instantiates a new XEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *XEmailResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *XEmailResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *XEmailResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *XEmailResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *XEmailResponse) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *XEmailResponse) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *XEmailResponse) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *XEmailResponse) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetMessageId

`func (o *XEmailResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *XEmailResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *XEmailResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *XEmailResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetErrorCode

`func (o *XEmailResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *XEmailResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *XEmailResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *XEmailResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *XEmailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *XEmailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *XEmailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *XEmailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


