# RestRTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique template identifier with template_ prefix | [optional] 
**Name** | Pointer to **string** | Name of the template | [optional] 
**Subject** | Pointer to **string** | Email subject line (if applicable) | [optional] 
**HtmlCode** | Pointer to **string** | HTML content of the template | [optional] 
**TemplateCode** | Pointer to **string** | Template code for visual editors (JSON structure) | [optional] 
**Type** | Pointer to **int32** | Template type.  **Values:** - &#x60;0&#x60; - Email template - &#x60;1&#x60; - Other types  | [optional] 
**Thumbnail** | Pointer to **string** | URL to template thumbnail image | [optional] 
**EditorType** | Pointer to **int32** | Editor type used to create the template.  **Values:** - &#x60;0&#x60; - PlainText - &#x60;1&#x60; - DragDrop - &#x60;2&#x60; - SendxEditor  | [optional] 
**Created** | Pointer to **time.Time** | Template creation timestamp | [optional] 
**Updated** | Pointer to **time.Time** | Template last update timestamp | [optional] 

## Methods

### NewRestRTemplate

`func NewRestRTemplate() *RestRTemplate`

NewRestRTemplate instantiates a new RestRTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRTemplateWithDefaults

`func NewRestRTemplateWithDefaults() *RestRTemplate`

NewRestRTemplateWithDefaults instantiates a new RestRTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *RestRTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RestRTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RestRTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RestRTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHtmlCode

`func (o *RestRTemplate) GetHtmlCode() string`

GetHtmlCode returns the HtmlCode field if non-nil, zero value otherwise.

### GetHtmlCodeOk

`func (o *RestRTemplate) GetHtmlCodeOk() (*string, bool)`

GetHtmlCodeOk returns a tuple with the HtmlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlCode

`func (o *RestRTemplate) SetHtmlCode(v string)`

SetHtmlCode sets HtmlCode field to given value.

### HasHtmlCode

`func (o *RestRTemplate) HasHtmlCode() bool`

HasHtmlCode returns a boolean if a field has been set.

### GetTemplateCode

`func (o *RestRTemplate) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *RestRTemplate) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *RestRTemplate) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *RestRTemplate) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetType

`func (o *RestRTemplate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestRTemplate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestRTemplate) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestRTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetThumbnail

`func (o *RestRTemplate) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *RestRTemplate) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *RestRTemplate) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *RestRTemplate) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetEditorType

`func (o *RestRTemplate) GetEditorType() int32`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *RestRTemplate) GetEditorTypeOk() (*int32, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *RestRTemplate) SetEditorType(v int32)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *RestRTemplate) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetCreated

`func (o *RestRTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRTemplate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRTemplate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRTemplate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


