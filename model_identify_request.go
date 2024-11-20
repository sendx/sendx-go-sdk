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

// checks if the IdentifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentifyRequest{}

// IdentifyRequest struct for IdentifyRequest
type IdentifyRequest struct {
	// First name of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// Last name of the contact.
	LastName *string `json:"lastName,omitempty"`
	// Email address of the contact.
	Email string `json:"email"`
	// New email address of the contact.
	NewEmail *string `json:"newEmail,omitempty"`
	// Company of the contact.
	Company *string `json:"company,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"customFields,omitempty"`
}

type _IdentifyRequest IdentifyRequest

// NewIdentifyRequest instantiates a new IdentifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifyRequest(email string) *IdentifyRequest {
	this := IdentifyRequest{}
	this.Email = email
	return &this
}

// NewIdentifyRequestWithDefaults instantiates a new IdentifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifyRequestWithDefaults() *IdentifyRequest {
	this := IdentifyRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IdentifyRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IdentifyRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IdentifyRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IdentifyRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IdentifyRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IdentifyRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value
func (o *IdentifyRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IdentifyRequest) SetEmail(v string) {
	o.Email = v
}

// GetNewEmail returns the NewEmail field value if set, zero value otherwise.
func (o *IdentifyRequest) GetNewEmail() string {
	if o == nil || IsNil(o.NewEmail) {
		var ret string
		return ret
	}
	return *o.NewEmail
}

// GetNewEmailOk returns a tuple with the NewEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetNewEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NewEmail) {
		return nil, false
	}
	return o.NewEmail, true
}

// HasNewEmail returns a boolean if a field has been set.
func (o *IdentifyRequest) HasNewEmail() bool {
	if o != nil && !IsNil(o.NewEmail) {
		return true
	}

	return false
}

// SetNewEmail gets a reference to the given string and assigns it to the NewEmail field.
func (o *IdentifyRequest) SetNewEmail(v string) {
	o.NewEmail = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *IdentifyRequest) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *IdentifyRequest) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *IdentifyRequest) SetCompany(v string) {
	o.Company = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IdentifyRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IdentifyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IdentifyRequest) SetTags(v []string) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IdentifyRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifyRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IdentifyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *IdentifyRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o IdentifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.NewEmail) {
		toSerialize["newEmail"] = o.NewEmail
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	return toSerialize, nil
}

func (o *IdentifyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varIdentifyRequest := _IdentifyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentifyRequest)

	if err != nil {
		return err
	}

	*o = IdentifyRequest(varIdentifyRequest)

	return err
}

type NullableIdentifyRequest struct {
	value *IdentifyRequest
	isSet bool
}

func (v NullableIdentifyRequest) Get() *IdentifyRequest {
	return v.value
}

func (v *NullableIdentifyRequest) Set(val *IdentifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifyRequest(val *IdentifyRequest) *NullableIdentifyRequest {
	return &NullableIdentifyRequest{value: val, isSet: true}
}

func (v NullableIdentifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

