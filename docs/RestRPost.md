# RestRPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PostTitle** | Pointer to **string** |  | [optional] 
**PostDescription** | Pointer to **string** |  | [optional] 
**PostCategory** | Pointer to **string** |  | [optional] 
**Member** | Pointer to **string** |  | [optional] 
**PostThumbnail** | Pointer to **string** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**IncludedTags** | Pointer to **[]string** |  | [optional] 
**PostSlug** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**PageTitle** | Pointer to **string** |  | [optional] 
**PageDescription** | Pointer to **string** |  | [optional] 
**PageKeywords** | Pointer to **string** |  | [optional] 
**SocialTitle** | Pointer to **string** |  | [optional] 
**SocialDescription** | Pointer to **string** |  | [optional] 
**SocialImageUrl** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRestRPost

`func NewRestRPost() *RestRPost`

NewRestRPost instantiates a new RestRPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestRPostWithDefaults

`func NewRestRPostWithDefaults() *RestRPost`

NewRestRPostWithDefaults instantiates a new RestRPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestRPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestRPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestRPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestRPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestRPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestRPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestRPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestRPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostTitle

`func (o *RestRPost) GetPostTitle() string`

GetPostTitle returns the PostTitle field if non-nil, zero value otherwise.

### GetPostTitleOk

`func (o *RestRPost) GetPostTitleOk() (*string, bool)`

GetPostTitleOk returns a tuple with the PostTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTitle

`func (o *RestRPost) SetPostTitle(v string)`

SetPostTitle sets PostTitle field to given value.

### HasPostTitle

`func (o *RestRPost) HasPostTitle() bool`

HasPostTitle returns a boolean if a field has been set.

### GetPostDescription

`func (o *RestRPost) GetPostDescription() string`

GetPostDescription returns the PostDescription field if non-nil, zero value otherwise.

### GetPostDescriptionOk

`func (o *RestRPost) GetPostDescriptionOk() (*string, bool)`

GetPostDescriptionOk returns a tuple with the PostDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDescription

`func (o *RestRPost) SetPostDescription(v string)`

SetPostDescription sets PostDescription field to given value.

### HasPostDescription

`func (o *RestRPost) HasPostDescription() bool`

HasPostDescription returns a boolean if a field has been set.

### GetPostCategory

`func (o *RestRPost) GetPostCategory() string`

GetPostCategory returns the PostCategory field if non-nil, zero value otherwise.

### GetPostCategoryOk

`func (o *RestRPost) GetPostCategoryOk() (*string, bool)`

GetPostCategoryOk returns a tuple with the PostCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCategory

`func (o *RestRPost) SetPostCategory(v string)`

SetPostCategory sets PostCategory field to given value.

### HasPostCategory

`func (o *RestRPost) HasPostCategory() bool`

HasPostCategory returns a boolean if a field has been set.

### GetMember

`func (o *RestRPost) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *RestRPost) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *RestRPost) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *RestRPost) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPostThumbnail

`func (o *RestRPost) GetPostThumbnail() string`

GetPostThumbnail returns the PostThumbnail field if non-nil, zero value otherwise.

### GetPostThumbnailOk

`func (o *RestRPost) GetPostThumbnailOk() (*string, bool)`

GetPostThumbnailOk returns a tuple with the PostThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostThumbnail

`func (o *RestRPost) SetPostThumbnail(v string)`

SetPostThumbnail sets PostThumbnail field to given value.

### HasPostThumbnail

`func (o *RestRPost) HasPostThumbnail() bool`

HasPostThumbnail returns a boolean if a field has been set.

### GetIsPublished

`func (o *RestRPost) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *RestRPost) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *RestRPost) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *RestRPost) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetIncludedTags

`func (o *RestRPost) GetIncludedTags() []string`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *RestRPost) GetIncludedTagsOk() (*[]string, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *RestRPost) SetIncludedTags(v []string)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *RestRPost) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetPostSlug

`func (o *RestRPost) GetPostSlug() string`

GetPostSlug returns the PostSlug field if non-nil, zero value otherwise.

### GetPostSlugOk

`func (o *RestRPost) GetPostSlugOk() (*string, bool)`

GetPostSlugOk returns a tuple with the PostSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSlug

`func (o *RestRPost) SetPostSlug(v string)`

SetPostSlug sets PostSlug field to given value.

### HasPostSlug

`func (o *RestRPost) HasPostSlug() bool`

HasPostSlug returns a boolean if a field has been set.

### GetStatus

`func (o *RestRPost) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestRPost) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestRPost) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestRPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPageTitle

`func (o *RestRPost) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *RestRPost) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *RestRPost) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *RestRPost) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetPageDescription

`func (o *RestRPost) GetPageDescription() string`

GetPageDescription returns the PageDescription field if non-nil, zero value otherwise.

### GetPageDescriptionOk

`func (o *RestRPost) GetPageDescriptionOk() (*string, bool)`

GetPageDescriptionOk returns a tuple with the PageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDescription

`func (o *RestRPost) SetPageDescription(v string)`

SetPageDescription sets PageDescription field to given value.

### HasPageDescription

`func (o *RestRPost) HasPageDescription() bool`

HasPageDescription returns a boolean if a field has been set.

### GetPageKeywords

`func (o *RestRPost) GetPageKeywords() string`

GetPageKeywords returns the PageKeywords field if non-nil, zero value otherwise.

### GetPageKeywordsOk

`func (o *RestRPost) GetPageKeywordsOk() (*string, bool)`

GetPageKeywordsOk returns a tuple with the PageKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageKeywords

`func (o *RestRPost) SetPageKeywords(v string)`

SetPageKeywords sets PageKeywords field to given value.

### HasPageKeywords

`func (o *RestRPost) HasPageKeywords() bool`

HasPageKeywords returns a boolean if a field has been set.

### GetSocialTitle

`func (o *RestRPost) GetSocialTitle() string`

GetSocialTitle returns the SocialTitle field if non-nil, zero value otherwise.

### GetSocialTitleOk

`func (o *RestRPost) GetSocialTitleOk() (*string, bool)`

GetSocialTitleOk returns a tuple with the SocialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialTitle

`func (o *RestRPost) SetSocialTitle(v string)`

SetSocialTitle sets SocialTitle field to given value.

### HasSocialTitle

`func (o *RestRPost) HasSocialTitle() bool`

HasSocialTitle returns a boolean if a field has been set.

### GetSocialDescription

`func (o *RestRPost) GetSocialDescription() string`

GetSocialDescription returns the SocialDescription field if non-nil, zero value otherwise.

### GetSocialDescriptionOk

`func (o *RestRPost) GetSocialDescriptionOk() (*string, bool)`

GetSocialDescriptionOk returns a tuple with the SocialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialDescription

`func (o *RestRPost) SetSocialDescription(v string)`

SetSocialDescription sets SocialDescription field to given value.

### HasSocialDescription

`func (o *RestRPost) HasSocialDescription() bool`

HasSocialDescription returns a boolean if a field has been set.

### GetSocialImageUrl

`func (o *RestRPost) GetSocialImageUrl() string`

GetSocialImageUrl returns the SocialImageUrl field if non-nil, zero value otherwise.

### GetSocialImageUrlOk

`func (o *RestRPost) GetSocialImageUrlOk() (*string, bool)`

GetSocialImageUrlOk returns a tuple with the SocialImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialImageUrl

`func (o *RestRPost) SetSocialImageUrl(v string)`

SetSocialImageUrl sets SocialImageUrl field to given value.

### HasSocialImageUrl

`func (o *RestRPost) HasSocialImageUrl() bool`

HasSocialImageUrl returns a boolean if a field has been set.

### GetCreated

`func (o *RestRPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestRPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestRPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestRPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RestRPost) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RestRPost) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RestRPost) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RestRPost) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


