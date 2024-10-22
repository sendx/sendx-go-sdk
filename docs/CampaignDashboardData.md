# CampaignDashboardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier for the campaign. | [optional] 
**Name** | Pointer to **string** | Name of the campaign. | [optional] 
**Subject** | Pointer to **string** | Subject of the campaign. | [optional] 
**SentTime** | Pointer to **time.Time** | The time the campaign was sent. | [optional] 
**CampaignScreenshotUrl** | Pointer to **string** | URL to a screenshot of the campaign. | [optional] 

## Methods

### NewCampaignDashboardData

`func NewCampaignDashboardData() *CampaignDashboardData`

NewCampaignDashboardData instantiates a new CampaignDashboardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDashboardDataWithDefaults

`func NewCampaignDashboardDataWithDefaults() *CampaignDashboardData`

NewCampaignDashboardDataWithDefaults instantiates a new CampaignDashboardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignDashboardData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignDashboardData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignDashboardData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignDashboardData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignDashboardData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignDashboardData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignDashboardData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignDashboardData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignDashboardData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignDashboardData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignDashboardData) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignDashboardData) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSentTime

`func (o *CampaignDashboardData) GetSentTime() time.Time`

GetSentTime returns the SentTime field if non-nil, zero value otherwise.

### GetSentTimeOk

`func (o *CampaignDashboardData) GetSentTimeOk() (*time.Time, bool)`

GetSentTimeOk returns a tuple with the SentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTime

`func (o *CampaignDashboardData) SetSentTime(v time.Time)`

SetSentTime sets SentTime field to given value.

### HasSentTime

`func (o *CampaignDashboardData) HasSentTime() bool`

HasSentTime returns a boolean if a field has been set.

### GetCampaignScreenshotUrl

`func (o *CampaignDashboardData) GetCampaignScreenshotUrl() string`

GetCampaignScreenshotUrl returns the CampaignScreenshotUrl field if non-nil, zero value otherwise.

### GetCampaignScreenshotUrlOk

`func (o *CampaignDashboardData) GetCampaignScreenshotUrlOk() (*string, bool)`

GetCampaignScreenshotUrlOk returns a tuple with the CampaignScreenshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignScreenshotUrl

`func (o *CampaignDashboardData) SetCampaignScreenshotUrl(v string)`

SetCampaignScreenshotUrl sets CampaignScreenshotUrl field to given value.

### HasCampaignScreenshotUrl

`func (o *CampaignDashboardData) HasCampaignScreenshotUrl() bool`

HasCampaignScreenshotUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


