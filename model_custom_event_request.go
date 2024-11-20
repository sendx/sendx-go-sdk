/*
SendX REST API

# Introduction SendX is an email marketing product. It helps you convert website visitors to customers, send them promotional emails, engage with them using drip sequences and craft custom journeys using powerful but simple automations. The SendX API is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs. The SendX Rest API doesn’t support bulk updates. You can work on only one object per request. <br> 

API version: 1.0.0
Contact: support@sendx.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendx

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEventRequest{}

// CustomEventRequest struct for CustomEventRequest
type CustomEventRequest struct {
	// Name of the custom event (e.g., 'abandoned_cart').
	Name string `json:"name"`
	// Unique identifier for the contact (e.g., contact's email).
	Identifier string `json:"identifier"`
	// Map of property-value pairs associated with the event, where both key and value are strings.
	Data map[string]string `json:"data"`
	// Unix timestamp of when the event occurred.
	Time int32 `json:"time"`
}

type _CustomEventRequest CustomEventRequest

// NewCustomEventRequest instantiates a new CustomEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEventRequest(name string, identifier string, data map[string]string, time int32) *CustomEventRequest {
	this := CustomEventRequest{}
	this.Name = name
	this.Identifier = identifier
	this.Data = data
	this.Time = time
	return &this
}

// NewCustomEventRequestWithDefaults instantiates a new CustomEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEventRequestWithDefaults() *CustomEventRequest {
	this := CustomEventRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CustomEventRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomEventRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomEventRequest) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value
func (o *CustomEventRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *CustomEventRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *CustomEventRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetData returns the Data field value
func (o *CustomEventRequest) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomEventRequest) GetDataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomEventRequest) SetData(v map[string]string) {
	o.Data = v
}

// GetTime returns the Time field value
func (o *CustomEventRequest) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *CustomEventRequest) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *CustomEventRequest) SetTime(v int32) {
	o.Time = v
}

func (o CustomEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["identifier"] = o.Identifier
	toSerialize["data"] = o.Data
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

func (o *CustomEventRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"identifier",
		"data",
		"time",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomEventRequest := _CustomEventRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEventRequest)

	if err != nil {
		return err
	}

	*o = CustomEventRequest(varCustomEventRequest)

	return err
}

type NullableCustomEventRequest struct {
	value *CustomEventRequest
	isSet bool
}

func (v NullableCustomEventRequest) Get() *CustomEventRequest {
	return v.value
}

func (v *NullableCustomEventRequest) Set(val *CustomEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEventRequest(val *CustomEventRequest) *NullableCustomEventRequest {
	return &NullableCustomEventRequest{value: val, isSet: true}
}

func (v NullableCustomEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


