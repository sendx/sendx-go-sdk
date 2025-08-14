# RestEPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Internal post name | 
**PostTitle** | **string** | Public post title | 
**PostDescription** | Pointer to **string** | Post excerpt/description | [optional] 
**PostCategory** | Pointer to **string** | Category ID (with or without prefix) | [optional] 
**Member** | Pointer to **string** | Author member ID | [optional] 
**PostThumbnail** | Pointer to **string** | Thumbnail image URL | [optional] 
**PostHtml** | Pointer to **string** | Post HTML content | [optional] 
**PostTemplate** | Pointer to **string** | Post template | [optional] 
**IsPublished** | Pointer to **bool** | Publication status | [optional] [default to false]
**IncludedTags** | Pointer to **[]string** | Post tag IDs | [optional] 
**EditorType** | Pointer to **int32** | Editor type used | [optional] [default to 1]
**PostSlug** | Pointer to **string** | URL slug | [optional] 
**Status** | Pointer to **int32** | Post status | [optional] [default to 1]
**PageTitle** | Pointer to **string** | SEO page title | [optional] 
**PageDescription** | Pointer to **string** | SEO meta description | [optional] 
**PageKeywords** | Pointer to **string** | SEO keywords | [optional] 
**SocialTitle** | Pointer to **string** | Social media title | [optional] 
**SocialDescription** | Pointer to **string** | Social media description | [optional] 
**SocialImageUrl** | Pointer to **string** | Social media image URL | [optional] 

## Methods

### NewRestEPost

`func NewRestEPost(name string, postTitle string, ) *RestEPost`

NewRestEPost instantiates a new RestEPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestEPostWithDefaults

`func NewRestEPostWithDefaults() *RestEPost`

NewRestEPostWithDefaults instantiates a new RestEPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestEPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestEPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestEPost) SetName(v string)`

SetName sets Name field to given value.


### GetPostTitle

`func (o *RestEPost) GetPostTitle() string`

GetPostTitle returns the PostTitle field if non-nil, zero value otherwise.

### GetPostTitleOk

`func (o *RestEPost) GetPostTitleOk() (*string, bool)`

GetPostTitleOk returns a tuple with the PostTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTitle

`func (o *RestEPost) SetPostTitle(v string)`

SetPostTitle sets PostTitle field to given value.


### GetPostDescription

`func (o *RestEPost) GetPostDescription() string`

GetPostDescription returns the PostDescription field if non-nil, zero value otherwise.

### GetPostDescriptionOk

`func (o *RestEPost) GetPostDescriptionOk() (*string, bool)`

GetPostDescriptionOk returns a tuple with the PostDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDescription

`func (o *RestEPost) SetPostDescription(v string)`

SetPostDescription sets PostDescription field to given value.

### HasPostDescription

`func (o *RestEPost) HasPostDescription() bool`

HasPostDescription returns a boolean if a field has been set.

### GetPostCategory

`func (o *RestEPost) GetPostCategory() string`

GetPostCategory returns the PostCategory field if non-nil, zero value otherwise.

### GetPostCategoryOk

`func (o *RestEPost) GetPostCategoryOk() (*string, bool)`

GetPostCategoryOk returns a tuple with the PostCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCategory

`func (o *RestEPost) SetPostCategory(v string)`

SetPostCategory sets PostCategory field to given value.

### HasPostCategory

`func (o *RestEPost) HasPostCategory() bool`

HasPostCategory returns a boolean if a field has been set.

### GetMember

`func (o *RestEPost) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *RestEPost) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *RestEPost) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *RestEPost) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPostThumbnail

`func (o *RestEPost) GetPostThumbnail() string`

GetPostThumbnail returns the PostThumbnail field if non-nil, zero value otherwise.

### GetPostThumbnailOk

`func (o *RestEPost) GetPostThumbnailOk() (*string, bool)`

GetPostThumbnailOk returns a tuple with the PostThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostThumbnail

`func (o *RestEPost) SetPostThumbnail(v string)`

SetPostThumbnail sets PostThumbnail field to given value.

### HasPostThumbnail

`func (o *RestEPost) HasPostThumbnail() bool`

HasPostThumbnail returns a boolean if a field has been set.

### GetPostHtml

`func (o *RestEPost) GetPostHtml() string`

GetPostHtml returns the PostHtml field if non-nil, zero value otherwise.

### GetPostHtmlOk

`func (o *RestEPost) GetPostHtmlOk() (*string, bool)`

GetPostHtmlOk returns a tuple with the PostHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostHtml

`func (o *RestEPost) SetPostHtml(v string)`

SetPostHtml sets PostHtml field to given value.

### HasPostHtml

`func (o *RestEPost) HasPostHtml() bool`

HasPostHtml returns a boolean if a field has been set.

### GetPostTemplate

`func (o *RestEPost) GetPostTemplate() string`

GetPostTemplate returns the PostTemplate field if non-nil, zero value otherwise.

### GetPostTemplateOk

`func (o *RestEPost) GetPostTemplateOk() (*string, bool)`

GetPostTemplateOk returns a tuple with the PostTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTemplate

`func (o *RestEPost) SetPostTemplate(v string)`

SetPostTemplate sets PostTemplate field to given value.

### HasPostTemplate

`func (o *RestEPost) HasPostTemplate() bool`

HasPostTemplate returns a boolean if a field has been set.

### GetIsPublished

`func (o *RestEPost) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *RestEPost) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *RestEPost) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *RestEPost) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetIncludedTags

`func (o *RestEPost) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *RestEPost) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *RestEPost) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *RestEPost) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetEditorType

`func (o *RestEPost) GetEditorType() int32`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *RestEPost) GetEditorTypeOk() (*int32, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *RestEPost) SetEditorType(v int32)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *RestEPost) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetPostSlug

`func (o *RestEPost) GetPostSlug() string`

GetPostSlug returns the PostSlug field if non-nil, zero value otherwise.

### GetPostSlugOk

`func (o *RestEPost) GetPostSlugOk() (*string, bool)`

GetPostSlugOk returns a tuple with the PostSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSlug

`func (o *RestEPost) SetPostSlug(v string)`

SetPostSlug sets PostSlug field to given value.

### HasPostSlug

`func (o *RestEPost) HasPostSlug() bool`

HasPostSlug returns a boolean if a field has been set.

### GetStatus

`func (o *RestEPost) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestEPost) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestEPost) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestEPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPageTitle

`func (o *RestEPost) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *RestEPost) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *RestEPost) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *RestEPost) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetPageDescription

`func (o *RestEPost) GetPageDescription() string`

GetPageDescription returns the PageDescription field if non-nil, zero value otherwise.

### GetPageDescriptionOk

`func (o *RestEPost) GetPageDescriptionOk() (*string, bool)`

GetPageDescriptionOk returns a tuple with the PageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDescription

`func (o *RestEPost) SetPageDescription(v string)`

SetPageDescription sets PageDescription field to given value.

### HasPageDescription

`func (o *RestEPost) HasPageDescription() bool`

HasPageDescription returns a boolean if a field has been set.

### GetPageKeywords

`func (o *RestEPost) GetPageKeywords() string`

GetPageKeywords returns the PageKeywords field if non-nil, zero value otherwise.

### GetPageKeywordsOk

`func (o *RestEPost) GetPageKeywordsOk() (*string, bool)`

GetPageKeywordsOk returns a tuple with the PageKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageKeywords

`func (o *RestEPost) SetPageKeywords(v string)`

SetPageKeywords sets PageKeywords field to given value.

### HasPageKeywords

`func (o *RestEPost) HasPageKeywords() bool`

HasPageKeywords returns a boolean if a field has been set.

### GetSocialTitle

`func (o *RestEPost) GetSocialTitle() string`

GetSocialTitle returns the SocialTitle field if non-nil, zero value otherwise.

### GetSocialTitleOk

`func (o *RestEPost) GetSocialTitleOk() (*string, bool)`

GetSocialTitleOk returns a tuple with the SocialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialTitle

`func (o *RestEPost) SetSocialTitle(v string)`

SetSocialTitle sets SocialTitle field to given value.

### HasSocialTitle

`func (o *RestEPost) HasSocialTitle() bool`

HasSocialTitle returns a boolean if a field has been set.

### GetSocialDescription

`func (o *RestEPost) GetSocialDescription() string`

GetSocialDescription returns the SocialDescription field if non-nil, zero value otherwise.

### GetSocialDescriptionOk

`func (o *RestEPost) GetSocialDescriptionOk() (*string, bool)`

GetSocialDescriptionOk returns a tuple with the SocialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialDescription

`func (o *RestEPost) SetSocialDescription(v string)`

SetSocialDescription sets SocialDescription field to given value.

### HasSocialDescription

`func (o *RestEPost) HasSocialDescription() bool`

HasSocialDescription returns a boolean if a field has been set.

### GetSocialImageUrl

`func (o *RestEPost) GetSocialImageUrl() string`

GetSocialImageUrl returns the SocialImageUrl field if non-nil, zero value otherwise.

### GetSocialImageUrlOk

`func (o *RestEPost) GetSocialImageUrlOk() (*string, bool)`

GetSocialImageUrlOk returns a tuple with the SocialImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialImageUrl

`func (o *RestEPost) SetSocialImageUrl(v string)`

SetSocialImageUrl sets SocialImageUrl field to given value.

### HasSocialImageUrl

`func (o *RestEPost) HasSocialImageUrl() bool`

HasSocialImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


