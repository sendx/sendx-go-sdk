# IdentifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**Contact**](Contact.md) |  | [optional] 

## Methods

### NewIdentifyResponse

`func NewIdentifyResponse() *IdentifyResponse`

NewIdentifyResponse instantiates a new IdentifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifyResponseWithDefaults

`func NewIdentifyResponseWithDefaults() *IdentifyResponse`

NewIdentifyResponseWithDefaults instantiates a new IdentifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IdentifyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentifyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentifyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentifyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *IdentifyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IdentifyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IdentifyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IdentifyResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *IdentifyResponse) GetData() Contact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IdentifyResponse) GetDataOk() (*Contact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IdentifyResponse) SetData(v Contact)`

SetData sets Data field to given value.

### HasData

`func (o *IdentifyResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


