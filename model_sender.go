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

// checks if the Sender type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sender{}

// Sender struct for Sender
type Sender struct {
	// ID of the sender
	Id string `json:"id"`
	// Name of the sender
	Name string `json:"name"`
	// Email address of the sender
	Email string `json:"email"`
}

type _Sender Sender

// NewSender instantiates a new Sender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSender(id string, name string, email string) *Sender {
	this := Sender{}
	this.Id = id
	this.Name = name
	this.Email = email
	return &this
}

// NewSenderWithDefaults instantiates a new Sender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderWithDefaults() *Sender {
	this := Sender{}
	return &this
}

// GetId returns the Id field value
func (o *Sender) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Sender) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Sender) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Sender) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Sender) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Sender) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *Sender) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Sender) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Sender) SetEmail(v string) {
	o.Email = v
}

func (o Sender) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sender) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *Sender) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"email",
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

	varSender := _Sender{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSender)

	if err != nil {
		return err
	}

	*o = Sender(varSender)

	return err
}

type NullableSender struct {
	value *Sender
	isSet bool
}

func (v NullableSender) Get() *Sender {
	return v.value
}

func (v *NullableSender) Set(val *Sender) {
	v.value = val
	v.isSet = true
}

func (v NullableSender) IsSet() bool {
	return v.isSet
}

func (v *NullableSender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSender(val *Sender) *NullableSender {
	return &NullableSender{value: val, isSet: true}
}

func (v NullableSender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

