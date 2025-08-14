# XEmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**XFrom**](XFrom.md) |  | 
**To** | [**[]XTo**](XTo.md) |  | 
**ReplyTo** | Pointer to [**XReplyTo**](XReplyTo.md) |  | [optional] 
**Subject** | **string** |  | 
**HtmlBody** | **string** |  | 
**TextBody** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Template** | Pointer to **string** | Template identifier | [optional] 

## Methods

### NewXEmailMessage

`func NewXEmailMessage(from XFrom, to []XTo, subject string, htmlBody string, ) *XEmailMessage`

NewXEmailMessage instantiates a new XEmailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXEmailMessageWithDefaults

`func NewXEmailMessageWithDefaults() *XEmailMessage`

NewXEmailMessageWithDefaults instantiates a new XEmailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *XEmailMessage) GetFrom() XFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *XEmailMessage) GetFromOk() (*XFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *XEmailMessage) SetFrom(v XFrom)`

SetFrom sets From field to given value.


### GetTo

`func (o *XEmailMessage) GetTo() []XTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *XEmailMessage) GetToOk() (*[]XTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *XEmailMessage) SetTo(v []XTo)`

SetTo sets To field to given value.


### GetReplyTo

`func (o *XEmailMessage) GetReplyTo() XReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *XEmailMessage) GetReplyToOk() (*XReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *XEmailMessage) SetReplyTo(v XReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *XEmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *XEmailMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *XEmailMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *XEmailMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtmlBody

`func (o *XEmailMessage) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *XEmailMessage) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *XEmailMessage) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.


### GetTextBody

`func (o *XEmailMessage) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *XEmailMessage) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *XEmailMessage) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *XEmailMessage) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetHeaders

`func (o *XEmailMessage) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *XEmailMessage) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *XEmailMessage) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *XEmailMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTemplate

`func (o *XEmailMessage) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *XEmailMessage) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *XEmailMessage) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *XEmailMessage) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


