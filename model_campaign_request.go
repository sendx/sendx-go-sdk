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

// checks if the CampaignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRequest{}

// CampaignRequest struct for CampaignRequest
type CampaignRequest struct {
	// The name of the campaign.
	Name *string `json:"name,omitempty"`
	// The HTML code of the campaign.
	HtmlCode *string `json:"htmlCode,omitempty"`
	// The subject of the campaign.
	Subject *string `json:"subject,omitempty"`
	// Sender unique identifier.
	Sender *string `json:"sender,omitempty"`
	// The preview text shown in email clients.
	PreviewText *string `json:"previewText,omitempty"`
	// The type of scheduling for the campaign <br> 0: Send Now <br> 1: Send Later 
	ScheduleType *int32 `json:"scheduleType,omitempty"`
	// The condition for scheduling the campaign.
	ScheduleCondition *string `json:"scheduleCondition,omitempty"`
	// The specific time to send the campaign.
	TimeCondition *string `json:"timeCondition,omitempty"`
	// The timezone for the campaign scheduling.
	Timezone *string `json:"timezone,omitempty"`
	// Preferred timezone for scheduling.
	PreferredTimezone *string `json:"preferredTimezone,omitempty"`
	// Specific time preference for sending the campaign.
	PreferredTimeCondition *string `json:"preferredTimeCondition,omitempty"`
	// Whether to send emails in each contact's timezone.
	SendInContactsTimezone *bool `json:"sendInContactsTimezone,omitempty"`
	// Whether to enable smart send features (e.g., optimizing send time).
	SmartSend *bool `json:"smartSend,omitempty"`
	// List of segment IDs to include.
	IncludedSegments []string `json:"includedSegments,omitempty"`
	// List of list IDs to include.
	IncludedLists []string `json:"includedLists,omitempty"`
	// List of tag IDs to include.
	IncludedTags []string `json:"includedTags,omitempty"`
	// List of segment IDs to exclude.
	ExcludedSegments []string `json:"excludedSegments,omitempty"`
	// List of list IDs to exclude.
	ExcludedLists []string `json:"excludedLists,omitempty"`
	// List of tag IDs to exclude.
	ExcludedTags []string `json:"excludedTags,omitempty"`
}

// NewCampaignRequest instantiates a new CampaignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRequest() *CampaignRequest {
	this := CampaignRequest{}
	return &this
}

// NewCampaignRequestWithDefaults instantiates a new CampaignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRequestWithDefaults() *CampaignRequest {
	this := CampaignRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignRequest) SetName(v string) {
	o.Name = &v
}

// GetHtmlCode returns the HtmlCode field value if set, zero value otherwise.
func (o *CampaignRequest) GetHtmlCode() string {
	if o == nil || IsNil(o.HtmlCode) {
		var ret string
		return ret
	}
	return *o.HtmlCode
}

// GetHtmlCodeOk returns a tuple with the HtmlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetHtmlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlCode) {
		return nil, false
	}
	return o.HtmlCode, true
}

// HasHtmlCode returns a boolean if a field has been set.
func (o *CampaignRequest) HasHtmlCode() bool {
	if o != nil && !IsNil(o.HtmlCode) {
		return true
	}

	return false
}

// SetHtmlCode gets a reference to the given string and assigns it to the HtmlCode field.
func (o *CampaignRequest) SetHtmlCode(v string) {
	o.HtmlCode = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CampaignRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CampaignRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CampaignRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *CampaignRequest) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *CampaignRequest) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *CampaignRequest) SetSender(v string) {
	o.Sender = &v
}

// GetPreviewText returns the PreviewText field value if set, zero value otherwise.
func (o *CampaignRequest) GetPreviewText() string {
	if o == nil || IsNil(o.PreviewText) {
		var ret string
		return ret
	}
	return *o.PreviewText
}

// GetPreviewTextOk returns a tuple with the PreviewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetPreviewTextOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewText) {
		return nil, false
	}
	return o.PreviewText, true
}

// HasPreviewText returns a boolean if a field has been set.
func (o *CampaignRequest) HasPreviewText() bool {
	if o != nil && !IsNil(o.PreviewText) {
		return true
	}

	return false
}

// SetPreviewText gets a reference to the given string and assigns it to the PreviewText field.
func (o *CampaignRequest) SetPreviewText(v string) {
	o.PreviewText = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *CampaignRequest) GetScheduleType() int32 {
	if o == nil || IsNil(o.ScheduleType) {
		var ret int32
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetScheduleTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ScheduleType) {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *CampaignRequest) HasScheduleType() bool {
	if o != nil && !IsNil(o.ScheduleType) {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given int32 and assigns it to the ScheduleType field.
func (o *CampaignRequest) SetScheduleType(v int32) {
	o.ScheduleType = &v
}

// GetScheduleCondition returns the ScheduleCondition field value if set, zero value otherwise.
func (o *CampaignRequest) GetScheduleCondition() string {
	if o == nil || IsNil(o.ScheduleCondition) {
		var ret string
		return ret
	}
	return *o.ScheduleCondition
}

// GetScheduleConditionOk returns a tuple with the ScheduleCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetScheduleConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleCondition) {
		return nil, false
	}
	return o.ScheduleCondition, true
}

// HasScheduleCondition returns a boolean if a field has been set.
func (o *CampaignRequest) HasScheduleCondition() bool {
	if o != nil && !IsNil(o.ScheduleCondition) {
		return true
	}

	return false
}

// SetScheduleCondition gets a reference to the given string and assigns it to the ScheduleCondition field.
func (o *CampaignRequest) SetScheduleCondition(v string) {
	o.ScheduleCondition = &v
}

// GetTimeCondition returns the TimeCondition field value if set, zero value otherwise.
func (o *CampaignRequest) GetTimeCondition() string {
	if o == nil || IsNil(o.TimeCondition) {
		var ret string
		return ret
	}
	return *o.TimeCondition
}

// GetTimeConditionOk returns a tuple with the TimeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetTimeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.TimeCondition) {
		return nil, false
	}
	return o.TimeCondition, true
}

// HasTimeCondition returns a boolean if a field has been set.
func (o *CampaignRequest) HasTimeCondition() bool {
	if o != nil && !IsNil(o.TimeCondition) {
		return true
	}

	return false
}

// SetTimeCondition gets a reference to the given string and assigns it to the TimeCondition field.
func (o *CampaignRequest) SetTimeCondition(v string) {
	o.TimeCondition = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CampaignRequest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CampaignRequest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CampaignRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetPreferredTimezone returns the PreferredTimezone field value if set, zero value otherwise.
func (o *CampaignRequest) GetPreferredTimezone() string {
	if o == nil || IsNil(o.PreferredTimezone) {
		var ret string
		return ret
	}
	return *o.PreferredTimezone
}

// GetPreferredTimezoneOk returns a tuple with the PreferredTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetPreferredTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredTimezone) {
		return nil, false
	}
	return o.PreferredTimezone, true
}

// HasPreferredTimezone returns a boolean if a field has been set.
func (o *CampaignRequest) HasPreferredTimezone() bool {
	if o != nil && !IsNil(o.PreferredTimezone) {
		return true
	}

	return false
}

// SetPreferredTimezone gets a reference to the given string and assigns it to the PreferredTimezone field.
func (o *CampaignRequest) SetPreferredTimezone(v string) {
	o.PreferredTimezone = &v
}

// GetPreferredTimeCondition returns the PreferredTimeCondition field value if set, zero value otherwise.
func (o *CampaignRequest) GetPreferredTimeCondition() string {
	if o == nil || IsNil(o.PreferredTimeCondition) {
		var ret string
		return ret
	}
	return *o.PreferredTimeCondition
}

// GetPreferredTimeConditionOk returns a tuple with the PreferredTimeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetPreferredTimeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredTimeCondition) {
		return nil, false
	}
	return o.PreferredTimeCondition, true
}

// HasPreferredTimeCondition returns a boolean if a field has been set.
func (o *CampaignRequest) HasPreferredTimeCondition() bool {
	if o != nil && !IsNil(o.PreferredTimeCondition) {
		return true
	}

	return false
}

// SetPreferredTimeCondition gets a reference to the given string and assigns it to the PreferredTimeCondition field.
func (o *CampaignRequest) SetPreferredTimeCondition(v string) {
	o.PreferredTimeCondition = &v
}

// GetSendInContactsTimezone returns the SendInContactsTimezone field value if set, zero value otherwise.
func (o *CampaignRequest) GetSendInContactsTimezone() bool {
	if o == nil || IsNil(o.SendInContactsTimezone) {
		var ret bool
		return ret
	}
	return *o.SendInContactsTimezone
}

// GetSendInContactsTimezoneOk returns a tuple with the SendInContactsTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetSendInContactsTimezoneOk() (*bool, bool) {
	if o == nil || IsNil(o.SendInContactsTimezone) {
		return nil, false
	}
	return o.SendInContactsTimezone, true
}

// HasSendInContactsTimezone returns a boolean if a field has been set.
func (o *CampaignRequest) HasSendInContactsTimezone() bool {
	if o != nil && !IsNil(o.SendInContactsTimezone) {
		return true
	}

	return false
}

// SetSendInContactsTimezone gets a reference to the given bool and assigns it to the SendInContactsTimezone field.
func (o *CampaignRequest) SetSendInContactsTimezone(v bool) {
	o.SendInContactsTimezone = &v
}

// GetSmartSend returns the SmartSend field value if set, zero value otherwise.
func (o *CampaignRequest) GetSmartSend() bool {
	if o == nil || IsNil(o.SmartSend) {
		var ret bool
		return ret
	}
	return *o.SmartSend
}

// GetSmartSendOk returns a tuple with the SmartSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetSmartSendOk() (*bool, bool) {
	if o == nil || IsNil(o.SmartSend) {
		return nil, false
	}
	return o.SmartSend, true
}

// HasSmartSend returns a boolean if a field has been set.
func (o *CampaignRequest) HasSmartSend() bool {
	if o != nil && !IsNil(o.SmartSend) {
		return true
	}

	return false
}

// SetSmartSend gets a reference to the given bool and assigns it to the SmartSend field.
func (o *CampaignRequest) SetSmartSend(v bool) {
	o.SmartSend = &v
}

// GetIncludedSegments returns the IncludedSegments field value if set, zero value otherwise.
func (o *CampaignRequest) GetIncludedSegments() []string {
	if o == nil || IsNil(o.IncludedSegments) {
		var ret []string
		return ret
	}
	return o.IncludedSegments
}

// GetIncludedSegmentsOk returns a tuple with the IncludedSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetIncludedSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedSegments) {
		return nil, false
	}
	return o.IncludedSegments, true
}

// HasIncludedSegments returns a boolean if a field has been set.
func (o *CampaignRequest) HasIncludedSegments() bool {
	if o != nil && !IsNil(o.IncludedSegments) {
		return true
	}

	return false
}

// SetIncludedSegments gets a reference to the given []string and assigns it to the IncludedSegments field.
func (o *CampaignRequest) SetIncludedSegments(v []string) {
	o.IncludedSegments = v
}

// GetIncludedLists returns the IncludedLists field value if set, zero value otherwise.
func (o *CampaignRequest) GetIncludedLists() []string {
	if o == nil || IsNil(o.IncludedLists) {
		var ret []string
		return ret
	}
	return o.IncludedLists
}

// GetIncludedListsOk returns a tuple with the IncludedLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetIncludedListsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLists) {
		return nil, false
	}
	return o.IncludedLists, true
}

// HasIncludedLists returns a boolean if a field has been set.
func (o *CampaignRequest) HasIncludedLists() bool {
	if o != nil && !IsNil(o.IncludedLists) {
		return true
	}

	return false
}

// SetIncludedLists gets a reference to the given []string and assigns it to the IncludedLists field.
func (o *CampaignRequest) SetIncludedLists(v []string) {
	o.IncludedLists = v
}

// GetIncludedTags returns the IncludedTags field value if set, zero value otherwise.
func (o *CampaignRequest) GetIncludedTags() []string {
	if o == nil || IsNil(o.IncludedTags) {
		var ret []string
		return ret
	}
	return o.IncludedTags
}

// GetIncludedTagsOk returns a tuple with the IncludedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetIncludedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedTags) {
		return nil, false
	}
	return o.IncludedTags, true
}

// HasIncludedTags returns a boolean if a field has been set.
func (o *CampaignRequest) HasIncludedTags() bool {
	if o != nil && !IsNil(o.IncludedTags) {
		return true
	}

	return false
}

// SetIncludedTags gets a reference to the given []string and assigns it to the IncludedTags field.
func (o *CampaignRequest) SetIncludedTags(v []string) {
	o.IncludedTags = v
}

// GetExcludedSegments returns the ExcludedSegments field value if set, zero value otherwise.
func (o *CampaignRequest) GetExcludedSegments() []string {
	if o == nil || IsNil(o.ExcludedSegments) {
		var ret []string
		return ret
	}
	return o.ExcludedSegments
}

// GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetExcludedSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedSegments) {
		return nil, false
	}
	return o.ExcludedSegments, true
}

// HasExcludedSegments returns a boolean if a field has been set.
func (o *CampaignRequest) HasExcludedSegments() bool {
	if o != nil && !IsNil(o.ExcludedSegments) {
		return true
	}

	return false
}

// SetExcludedSegments gets a reference to the given []string and assigns it to the ExcludedSegments field.
func (o *CampaignRequest) SetExcludedSegments(v []string) {
	o.ExcludedSegments = v
}

// GetExcludedLists returns the ExcludedLists field value if set, zero value otherwise.
func (o *CampaignRequest) GetExcludedLists() []string {
	if o == nil || IsNil(o.ExcludedLists) {
		var ret []string
		return ret
	}
	return o.ExcludedLists
}

// GetExcludedListsOk returns a tuple with the ExcludedLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetExcludedListsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedLists) {
		return nil, false
	}
	return o.ExcludedLists, true
}

// HasExcludedLists returns a boolean if a field has been set.
func (o *CampaignRequest) HasExcludedLists() bool {
	if o != nil && !IsNil(o.ExcludedLists) {
		return true
	}

	return false
}

// SetExcludedLists gets a reference to the given []string and assigns it to the ExcludedLists field.
func (o *CampaignRequest) SetExcludedLists(v []string) {
	o.ExcludedLists = v
}

// GetExcludedTags returns the ExcludedTags field value if set, zero value otherwise.
func (o *CampaignRequest) GetExcludedTags() []string {
	if o == nil || IsNil(o.ExcludedTags) {
		var ret []string
		return ret
	}
	return o.ExcludedTags
}

// GetExcludedTagsOk returns a tuple with the ExcludedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRequest) GetExcludedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedTags) {
		return nil, false
	}
	return o.ExcludedTags, true
}

// HasExcludedTags returns a boolean if a field has been set.
func (o *CampaignRequest) HasExcludedTags() bool {
	if o != nil && !IsNil(o.ExcludedTags) {
		return true
	}

	return false
}

// SetExcludedTags gets a reference to the given []string and assigns it to the ExcludedTags field.
func (o *CampaignRequest) SetExcludedTags(v []string) {
	o.ExcludedTags = v
}

func (o CampaignRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HtmlCode) {
		toSerialize["htmlCode"] = o.HtmlCode
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.PreviewText) {
		toSerialize["previewText"] = o.PreviewText
	}
	if !IsNil(o.ScheduleType) {
		toSerialize["scheduleType"] = o.ScheduleType
	}
	if !IsNil(o.ScheduleCondition) {
		toSerialize["scheduleCondition"] = o.ScheduleCondition
	}
	if !IsNil(o.TimeCondition) {
		toSerialize["timeCondition"] = o.TimeCondition
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.PreferredTimezone) {
		toSerialize["preferredTimezone"] = o.PreferredTimezone
	}
	if !IsNil(o.PreferredTimeCondition) {
		toSerialize["preferredTimeCondition"] = o.PreferredTimeCondition
	}
	if !IsNil(o.SendInContactsTimezone) {
		toSerialize["sendInContactsTimezone"] = o.SendInContactsTimezone
	}
	if !IsNil(o.SmartSend) {
		toSerialize["smartSend"] = o.SmartSend
	}
	if !IsNil(o.IncludedSegments) {
		toSerialize["includedSegments"] = o.IncludedSegments
	}
	if !IsNil(o.IncludedLists) {
		toSerialize["includedLists"] = o.IncludedLists
	}
	if !IsNil(o.IncludedTags) {
		toSerialize["includedTags"] = o.IncludedTags
	}
	if !IsNil(o.ExcludedSegments) {
		toSerialize["excludedSegments"] = o.ExcludedSegments
	}
	if !IsNil(o.ExcludedLists) {
		toSerialize["excludedLists"] = o.ExcludedLists
	}
	if !IsNil(o.ExcludedTags) {
		toSerialize["excludedTags"] = o.ExcludedTags
	}
	return toSerialize, nil
}

type NullableCampaignRequest struct {
	value *CampaignRequest
	isSet bool
}

func (v NullableCampaignRequest) Get() *CampaignRequest {
	return v.value
}

func (v *NullableCampaignRequest) Set(val *CampaignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRequest(val *CampaignRequest) *NullableCampaignRequest {
	return &NullableCampaignRequest{value: val, isSet: true}
}

func (v NullableCampaignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


