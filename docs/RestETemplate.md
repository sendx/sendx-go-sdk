# RestETemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Template name | 
**Subject** | **string** | Email subject line | 
**HtmlContent** | Pointer to **string** | HTML email content | [optional] 
**TextContent** | Pointer to **string** | Plain text content | [optional] 
**Preheader** | Pointer to **string** | Preview text | [optional] 
**Tags** | Pointer to **[]string** | Template tags for organization | [optional] 

## Methods

### NewRestETemplate

`func NewRestETemplate(name string, subject string, ) *RestETemplate`

NewRestETemplate instantiates a new RestETemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestETemplateWithDefaults

`func NewRestETemplateWithDefaults() *RestETemplate`

NewRestETemplateWithDefaults instantiates a new RestETemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestETemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestETemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestETemplate) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *RestETemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RestETemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RestETemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtmlContent

`func (o *RestETemplate) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *RestETemplate) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *RestETemplate) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *RestETemplate) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetTextContent

`func (o *RestETemplate) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *RestETemplate) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *RestETemplate) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *RestETemplate) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetPreheader

`func (o *RestETemplate) GetPreheader() string`

GetPreheader returns the Preheader field if non-nil, zero value otherwise.

### GetPreheaderOk

`func (o *RestETemplate) GetPreheaderOk() (*string, bool)`

GetPreheaderOk returns a tuple with the Preheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreheader

`func (o *RestETemplate) SetPreheader(v string)`

SetPreheader sets Preheader field to given value.

### HasPreheader

`func (o *RestETemplate) HasPreheader() bool`

HasPreheader returns a boolean if a field has been set.

### GetTags

`func (o *RestETemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestETemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestETemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestETemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


