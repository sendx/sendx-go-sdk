# WebhookObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **int64** | The type of the event. | [optional] 
**Time** | Pointer to **int64** | The timestamp of the event in milliseconds since the epoch. | [optional] 
**Data** | Pointer to **string** | Arbitrary data associated with the event. | [optional] 
**ProviderMessageId** | Pointer to **string** | Optional provider message ID. | [optional] 
**CampaignId** | Pointer to **string** | Optional campaign ID. | [optional] 
**DripStepId** | Pointer to **string** | Optional drip step ID. | [optional] 
**RssExecId** | Pointer to **string** | Optional RSS execution ID. | [optional] 
**TagId** | Pointer to **string** | Optional tag ID. | [optional] 
**Link** | Pointer to **string** | Optional link associated with the event. | [optional] 
**ListId** | Pointer to **string** | Optional list ID. | [optional] 
**ContactId** | Pointer to **string** | Optional contact ID. | [optional] 
**CustomFieldId** | Pointer to **string** | Optional custom field ID. | [optional] 
**TemplateId** | Pointer to **string** | Optional template ID. | [optional] 
**PopupId** | Pointer to **string** | Optional popup ID. | [optional] 
**LandingPageId** | Pointer to **string** | Optional landing page ID. | [optional] 
**FormId** | Pointer to **string** | Optional form ID. | [optional] 
**SegmentId** | Pointer to **string** | Optional segment ID. | [optional] 
**AutomationId** | Pointer to **string** | Optional automation ID. | [optional] 
**DripId** | Pointer to **string** | Optional drip ID. | [optional] 
**RssId** | Pointer to **string** | Optional RSS ID. | [optional] 
**AbTestId** | Pointer to **string** | Optional A/B test ID. | [optional] 
**WorkflowId** | Pointer to **string** | Optional workflow ID. | [optional] 
**WorkflowNodeId** | Pointer to **string** | Optional workflow node ID. | [optional] 
**WorkflowEmailId** | Pointer to **string** | Optional workflow email ID. | [optional] 

## Methods

### NewWebhookObject

`func NewWebhookObject() *WebhookObject`

NewWebhookObject instantiates a new WebhookObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookObjectWithDefaults

`func NewWebhookObjectWithDefaults() *WebhookObject`

NewWebhookObjectWithDefaults instantiates a new WebhookObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookObject) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookObject) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookObject) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTime

`func (o *WebhookObject) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WebhookObject) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WebhookObject) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *WebhookObject) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetData

`func (o *WebhookObject) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookObject) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookObject) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookObject) HasData() bool`

HasData returns a boolean if a field has been set.

### GetProviderMessageId

`func (o *WebhookObject) GetProviderMessageId() string`

GetProviderMessageId returns the ProviderMessageId field if non-nil, zero value otherwise.

### GetProviderMessageIdOk

`func (o *WebhookObject) GetProviderMessageIdOk() (*string, bool)`

GetProviderMessageIdOk returns a tuple with the ProviderMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMessageId

`func (o *WebhookObject) SetProviderMessageId(v string)`

SetProviderMessageId sets ProviderMessageId field to given value.

### HasProviderMessageId

`func (o *WebhookObject) HasProviderMessageId() bool`

HasProviderMessageId returns a boolean if a field has been set.

### GetCampaignId

`func (o *WebhookObject) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *WebhookObject) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *WebhookObject) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *WebhookObject) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetDripStepId

`func (o *WebhookObject) GetDripStepId() string`

GetDripStepId returns the DripStepId field if non-nil, zero value otherwise.

### GetDripStepIdOk

`func (o *WebhookObject) GetDripStepIdOk() (*string, bool)`

GetDripStepIdOk returns a tuple with the DripStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripStepId

`func (o *WebhookObject) SetDripStepId(v string)`

SetDripStepId sets DripStepId field to given value.

### HasDripStepId

`func (o *WebhookObject) HasDripStepId() bool`

HasDripStepId returns a boolean if a field has been set.

### GetRssExecId

`func (o *WebhookObject) GetRssExecId() string`

GetRssExecId returns the RssExecId field if non-nil, zero value otherwise.

### GetRssExecIdOk

`func (o *WebhookObject) GetRssExecIdOk() (*string, bool)`

GetRssExecIdOk returns a tuple with the RssExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssExecId

`func (o *WebhookObject) SetRssExecId(v string)`

SetRssExecId sets RssExecId field to given value.

### HasRssExecId

`func (o *WebhookObject) HasRssExecId() bool`

HasRssExecId returns a boolean if a field has been set.

### GetTagId

`func (o *WebhookObject) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *WebhookObject) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *WebhookObject) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *WebhookObject) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetLink

`func (o *WebhookObject) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WebhookObject) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WebhookObject) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *WebhookObject) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetListId

`func (o *WebhookObject) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *WebhookObject) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *WebhookObject) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *WebhookObject) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetContactId

`func (o *WebhookObject) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *WebhookObject) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *WebhookObject) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *WebhookObject) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetCustomFieldId

`func (o *WebhookObject) GetCustomFieldId() string`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *WebhookObject) GetCustomFieldIdOk() (*string, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *WebhookObject) SetCustomFieldId(v string)`

SetCustomFieldId sets CustomFieldId field to given value.

### HasCustomFieldId

`func (o *WebhookObject) HasCustomFieldId() bool`

HasCustomFieldId returns a boolean if a field has been set.

### GetTemplateId

`func (o *WebhookObject) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *WebhookObject) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *WebhookObject) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *WebhookObject) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetPopupId

`func (o *WebhookObject) GetPopupId() string`

GetPopupId returns the PopupId field if non-nil, zero value otherwise.

### GetPopupIdOk

`func (o *WebhookObject) GetPopupIdOk() (*string, bool)`

GetPopupIdOk returns a tuple with the PopupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopupId

`func (o *WebhookObject) SetPopupId(v string)`

SetPopupId sets PopupId field to given value.

### HasPopupId

`func (o *WebhookObject) HasPopupId() bool`

HasPopupId returns a boolean if a field has been set.

### GetLandingPageId

`func (o *WebhookObject) GetLandingPageId() string`

GetLandingPageId returns the LandingPageId field if non-nil, zero value otherwise.

### GetLandingPageIdOk

`func (o *WebhookObject) GetLandingPageIdOk() (*string, bool)`

GetLandingPageIdOk returns a tuple with the LandingPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageId

`func (o *WebhookObject) SetLandingPageId(v string)`

SetLandingPageId sets LandingPageId field to given value.

### HasLandingPageId

`func (o *WebhookObject) HasLandingPageId() bool`

HasLandingPageId returns a boolean if a field has been set.

### GetFormId

`func (o *WebhookObject) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *WebhookObject) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *WebhookObject) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *WebhookObject) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetSegmentId

`func (o *WebhookObject) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *WebhookObject) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *WebhookObject) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *WebhookObject) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### GetAutomationId

`func (o *WebhookObject) GetAutomationId() string`

GetAutomationId returns the AutomationId field if non-nil, zero value otherwise.

### GetAutomationIdOk

`func (o *WebhookObject) GetAutomationIdOk() (*string, bool)`

GetAutomationIdOk returns a tuple with the AutomationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationId

`func (o *WebhookObject) SetAutomationId(v string)`

SetAutomationId sets AutomationId field to given value.

### HasAutomationId

`func (o *WebhookObject) HasAutomationId() bool`

HasAutomationId returns a boolean if a field has been set.

### GetDripId

`func (o *WebhookObject) GetDripId() string`

GetDripId returns the DripId field if non-nil, zero value otherwise.

### GetDripIdOk

`func (o *WebhookObject) GetDripIdOk() (*string, bool)`

GetDripIdOk returns a tuple with the DripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripId

`func (o *WebhookObject) SetDripId(v string)`

SetDripId sets DripId field to given value.

### HasDripId

`func (o *WebhookObject) HasDripId() bool`

HasDripId returns a boolean if a field has been set.

### GetRssId

`func (o *WebhookObject) GetRssId() string`

GetRssId returns the RssId field if non-nil, zero value otherwise.

### GetRssIdOk

`func (o *WebhookObject) GetRssIdOk() (*string, bool)`

GetRssIdOk returns a tuple with the RssId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssId

`func (o *WebhookObject) SetRssId(v string)`

SetRssId sets RssId field to given value.

### HasRssId

`func (o *WebhookObject) HasRssId() bool`

HasRssId returns a boolean if a field has been set.

### GetAbTestId

`func (o *WebhookObject) GetAbTestId() string`

GetAbTestId returns the AbTestId field if non-nil, zero value otherwise.

### GetAbTestIdOk

`func (o *WebhookObject) GetAbTestIdOk() (*string, bool)`

GetAbTestIdOk returns a tuple with the AbTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestId

`func (o *WebhookObject) SetAbTestId(v string)`

SetAbTestId sets AbTestId field to given value.

### HasAbTestId

`func (o *WebhookObject) HasAbTestId() bool`

HasAbTestId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WebhookObject) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WebhookObject) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WebhookObject) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WebhookObject) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowNodeId

`func (o *WebhookObject) GetWorkflowNodeId() string`

GetWorkflowNodeId returns the WorkflowNodeId field if non-nil, zero value otherwise.

### GetWorkflowNodeIdOk

`func (o *WebhookObject) GetWorkflowNodeIdOk() (*string, bool)`

GetWorkflowNodeIdOk returns a tuple with the WorkflowNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNodeId

`func (o *WebhookObject) SetWorkflowNodeId(v string)`

SetWorkflowNodeId sets WorkflowNodeId field to given value.

### HasWorkflowNodeId

`func (o *WebhookObject) HasWorkflowNodeId() bool`

HasWorkflowNodeId returns a boolean if a field has been set.

### GetWorkflowEmailId

`func (o *WebhookObject) GetWorkflowEmailId() string`

GetWorkflowEmailId returns the WorkflowEmailId field if non-nil, zero value otherwise.

### GetWorkflowEmailIdOk

`func (o *WebhookObject) GetWorkflowEmailIdOk() (*string, bool)`

GetWorkflowEmailIdOk returns a tuple with the WorkflowEmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEmailId

`func (o *WebhookObject) SetWorkflowEmailId(v string)`

SetWorkflowEmailId sets WorkflowEmailId field to given value.

### HasWorkflowEmailId

`func (o *WebhookObject) HasWorkflowEmailId() bool`

HasWorkflowEmailId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


