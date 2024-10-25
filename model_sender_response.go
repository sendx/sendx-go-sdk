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

// checks if the SenderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SenderResponse{}

// SenderResponse struct for SenderResponse
type SenderResponse struct {
	// Encrypted ID of the sender
	Id string `json:"id"`
	// Name of the sender
	Name string `json:"name"`
	// Email address of the sender
	Email string `json:"email"`
	// Indicates if the sender is whitelisted
	IsWhitelisted bool `json:"isWhitelisted"`
}

type _SenderResponse SenderResponse

// NewSenderResponse instantiates a new SenderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSenderResponse(id string, name string, email string, isWhitelisted bool) *SenderResponse {
	this := SenderResponse{}
	this.Id = id
	this.Name = name
	this.Email = email
	this.IsWhitelisted = isWhitelisted
	return &this
}

// NewSenderResponseWithDefaults instantiates a new SenderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderResponseWithDefaults() *SenderResponse {
	this := SenderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SenderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SenderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SenderResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SenderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SenderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SenderResponse) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SenderResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SenderResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SenderResponse) SetEmail(v string) {
	o.Email = v
}

// GetIsWhitelisted returns the IsWhitelisted field value
func (o *SenderResponse) GetIsWhitelisted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWhitelisted
}

// GetIsWhitelistedOk returns a tuple with the IsWhitelisted field value
// and a boolean to check if the value has been set.
func (o *SenderResponse) GetIsWhitelistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWhitelisted, true
}

// SetIsWhitelisted sets field value
func (o *SenderResponse) SetIsWhitelisted(v bool) {
	o.IsWhitelisted = v
}

func (o SenderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SenderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["isWhitelisted"] = o.IsWhitelisted
	return toSerialize, nil
}

func (o *SenderResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"email",
		"isWhitelisted",
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

	varSenderResponse := _SenderResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSenderResponse)

	if err != nil {
		return err
	}

	*o = SenderResponse(varSenderResponse)

	return err
}

type NullableSenderResponse struct {
	value *SenderResponse
	isSet bool
}

func (v NullableSenderResponse) Get() *SenderResponse {
	return v.value
}

func (v *NullableSenderResponse) Set(val *SenderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSenderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSenderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSenderResponse(val *SenderResponse) *NullableSenderResponse {
	return &NullableSenderResponse{value: val, isSet: true}
}

func (v NullableSenderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSenderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


