# CustomEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the custom event (e.g., &#39;abandoned_cart&#39;). | 
**Identifier** | **string** | Unique identifier for the contact (e.g., contact&#39;s email). | 
**Data** | **map[string]string** | Map of property-value pairs associated with the event, where both key and value are strings. | 
**Time** | **int32** | Unix timestamp of when the event occurred. | 

## Methods

### NewCustomEventRequest

`func NewCustomEventRequest(name string, identifier string, data map[string]string, time int32, ) *CustomEventRequest`

NewCustomEventRequest instantiates a new CustomEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventRequestWithDefaults

`func NewCustomEventRequestWithDefaults() *CustomEventRequest`

NewCustomEventRequestWithDefaults instantiates a new CustomEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEventRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *CustomEventRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CustomEventRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CustomEventRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetData

`func (o *CustomEventRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomEventRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomEventRequest) SetData(v map[string]string)`

SetData sets Data field to given value.


### GetTime

`func (o *CustomEventRequest) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CustomEventRequest) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CustomEventRequest) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


