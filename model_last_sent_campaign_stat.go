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
)

// checks if the LastSentCampaignStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastSentCampaignStat{}

// LastSentCampaignStat struct for LastSentCampaignStat
type LastSentCampaignStat struct {
	Campaign *CampaignDashboardData `json:"campaign,omitempty"`
	// Number of emails sent.
	Sent *int32 `json:"sent,omitempty"`
	// Number of emails delivered.
	Delivered *int32 `json:"delivered,omitempty"`
	// Number of new subscriptions.
	Subscribed *int32 `json:"subscribed,omitempty"`
	// Number of unsubscribes.
	Unsubscribed *int32 `json:"unsubscribed,omitempty"`
	// Number of emails opened.
	Open *int32 `json:"open,omitempty"`
	// Number of link clicks.
	LinkClicked *int32 `json:"linkClicked,omitempty"`
	// Number of replies received.
	Replied *int32 `json:"replied,omitempty"`
}

// NewLastSentCampaignStat instantiates a new LastSentCampaignStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastSentCampaignStat() *LastSentCampaignStat {
	this := LastSentCampaignStat{}
	return &this
}

// NewLastSentCampaignStatWithDefaults instantiates a new LastSentCampaignStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastSentCampaignStatWithDefaults() *LastSentCampaignStat {
	this := LastSentCampaignStat{}
	return &this
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetCampaign() CampaignDashboardData {
	if o == nil || IsNil(o.Campaign) {
		var ret CampaignDashboardData
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetCampaignOk() (*CampaignDashboardData, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given CampaignDashboardData and assigns it to the Campaign field.
func (o *LastSentCampaignStat) SetCampaign(v CampaignDashboardData) {
	o.Campaign = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetSent() int32 {
	if o == nil || IsNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetSentOk() (*int32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *LastSentCampaignStat) SetSent(v int32) {
	o.Sent = &v
}

// GetDelivered returns the Delivered field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetDelivered() int32 {
	if o == nil || IsNil(o.Delivered) {
		var ret int32
		return ret
	}
	return *o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetDeliveredOk() (*int32, bool) {
	if o == nil || IsNil(o.Delivered) {
		return nil, false
	}
	return o.Delivered, true
}

// HasDelivered returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasDelivered() bool {
	if o != nil && !IsNil(o.Delivered) {
		return true
	}

	return false
}

// SetDelivered gets a reference to the given int32 and assigns it to the Delivered field.
func (o *LastSentCampaignStat) SetDelivered(v int32) {
	o.Delivered = &v
}

// GetSubscribed returns the Subscribed field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetSubscribed() int32 {
	if o == nil || IsNil(o.Subscribed) {
		var ret int32
		return ret
	}
	return *o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetSubscribedOk() (*int32, bool) {
	if o == nil || IsNil(o.Subscribed) {
		return nil, false
	}
	return o.Subscribed, true
}

// HasSubscribed returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasSubscribed() bool {
	if o != nil && !IsNil(o.Subscribed) {
		return true
	}

	return false
}

// SetSubscribed gets a reference to the given int32 and assigns it to the Subscribed field.
func (o *LastSentCampaignStat) SetSubscribed(v int32) {
	o.Subscribed = &v
}

// GetUnsubscribed returns the Unsubscribed field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetUnsubscribed() int32 {
	if o == nil || IsNil(o.Unsubscribed) {
		var ret int32
		return ret
	}
	return *o.Unsubscribed
}

// GetUnsubscribedOk returns a tuple with the Unsubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetUnsubscribedOk() (*int32, bool) {
	if o == nil || IsNil(o.Unsubscribed) {
		return nil, false
	}
	return o.Unsubscribed, true
}

// HasUnsubscribed returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasUnsubscribed() bool {
	if o != nil && !IsNil(o.Unsubscribed) {
		return true
	}

	return false
}

// SetUnsubscribed gets a reference to the given int32 and assigns it to the Unsubscribed field.
func (o *LastSentCampaignStat) SetUnsubscribed(v int32) {
	o.Unsubscribed = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetOpen() int32 {
	if o == nil || IsNil(o.Open) {
		var ret int32
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetOpenOk() (*int32, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given int32 and assigns it to the Open field.
func (o *LastSentCampaignStat) SetOpen(v int32) {
	o.Open = &v
}

// GetLinkClicked returns the LinkClicked field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetLinkClicked() int32 {
	if o == nil || IsNil(o.LinkClicked) {
		var ret int32
		return ret
	}
	return *o.LinkClicked
}

// GetLinkClickedOk returns a tuple with the LinkClicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetLinkClickedOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkClicked) {
		return nil, false
	}
	return o.LinkClicked, true
}

// HasLinkClicked returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasLinkClicked() bool {
	if o != nil && !IsNil(o.LinkClicked) {
		return true
	}

	return false
}

// SetLinkClicked gets a reference to the given int32 and assigns it to the LinkClicked field.
func (o *LastSentCampaignStat) SetLinkClicked(v int32) {
	o.LinkClicked = &v
}

// GetReplied returns the Replied field value if set, zero value otherwise.
func (o *LastSentCampaignStat) GetReplied() int32 {
	if o == nil || IsNil(o.Replied) {
		var ret int32
		return ret
	}
	return *o.Replied
}

// GetRepliedOk returns a tuple with the Replied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastSentCampaignStat) GetRepliedOk() (*int32, bool) {
	if o == nil || IsNil(o.Replied) {
		return nil, false
	}
	return o.Replied, true
}

// HasReplied returns a boolean if a field has been set.
func (o *LastSentCampaignStat) HasReplied() bool {
	if o != nil && !IsNil(o.Replied) {
		return true
	}

	return false
}

// SetReplied gets a reference to the given int32 and assigns it to the Replied field.
func (o *LastSentCampaignStat) SetReplied(v int32) {
	o.Replied = &v
}

func (o LastSentCampaignStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastSentCampaignStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Delivered) {
		toSerialize["delivered"] = o.Delivered
	}
	if !IsNil(o.Subscribed) {
		toSerialize["subscribed"] = o.Subscribed
	}
	if !IsNil(o.Unsubscribed) {
		toSerialize["unsubscribed"] = o.Unsubscribed
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.LinkClicked) {
		toSerialize["linkClicked"] = o.LinkClicked
	}
	if !IsNil(o.Replied) {
		toSerialize["replied"] = o.Replied
	}
	return toSerialize, nil
}

type NullableLastSentCampaignStat struct {
	value *LastSentCampaignStat
	isSet bool
}

func (v NullableLastSentCampaignStat) Get() *LastSentCampaignStat {
	return v.value
}

func (v *NullableLastSentCampaignStat) Set(val *LastSentCampaignStat) {
	v.value = val
	v.isSet = true
}

func (v NullableLastSentCampaignStat) IsSet() bool {
	return v.isSet
}

func (v *NullableLastSentCampaignStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastSentCampaignStat(val *LastSentCampaignStat) *NullableLastSentCampaignStat {
	return &NullableLastSentCampaignStat{value: val, isSet: true}
}

func (v NullableLastSentCampaignStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastSentCampaignStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


