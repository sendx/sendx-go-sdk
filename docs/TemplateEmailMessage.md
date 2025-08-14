# TemplateEmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**XFrom**](XFrom.md) |  | 
**To** | [**[]XTo**](XTo.md) |  | 
**ReplyTo** | Pointer to [**XReplyTo**](XReplyTo.md) |  | [optional] 
**Subject** | **string** | Override template subject | 
**Template** | **string** | Template identifier | 

## Methods

### NewTemplateEmailMessage

`func NewTemplateEmailMessage(from XFrom, to []XTo, subject string, template string, ) *TemplateEmailMessage`

NewTemplateEmailMessage instantiates a new TemplateEmailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEmailMessageWithDefaults

`func NewTemplateEmailMessageWithDefaults() *TemplateEmailMessage`

NewTemplateEmailMessageWithDefaults instantiates a new TemplateEmailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TemplateEmailMessage) GetFrom() XFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateEmailMessage) GetFromOk() (*XFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateEmailMessage) SetFrom(v XFrom)`

SetFrom sets From field to given value.


### GetTo

`func (o *TemplateEmailMessage) GetTo() []XTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TemplateEmailMessage) GetToOk() (*[]XTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TemplateEmailMessage) SetTo(v []XTo)`

SetTo sets To field to given value.


### GetReplyTo

`func (o *TemplateEmailMessage) GetReplyTo() XReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TemplateEmailMessage) GetReplyToOk() (*XReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TemplateEmailMessage) SetReplyTo(v XReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TemplateEmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateEmailMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateEmailMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateEmailMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTemplate

`func (o *TemplateEmailMessage) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *TemplateEmailMessage) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *TemplateEmailMessage) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


