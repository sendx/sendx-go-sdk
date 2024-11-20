# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Unique identifier for the created event. | [optional] 
**Status** | Pointer to **string** | Status of the event creation (e.g., &#39;success&#39;). | [optional] 
**Message** | Pointer to **string** | Additional message about the event creation. | [optional] 

## Methods

### NewEventResponse

`func NewEventResponse() *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EventResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EventResponse) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetStatus

`func (o *EventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *EventResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


