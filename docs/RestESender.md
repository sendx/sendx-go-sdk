# RestESender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Sender display name | 
**Email** | **string** | Sender email address (must be verified) | 

## Methods

### NewRestESender

`func NewRestESender(name string, email string, ) *RestESender`

NewRestESender instantiates a new RestESender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestESenderWithDefaults

`func NewRestESenderWithDefaults() *RestESender`

NewRestESenderWithDefaults instantiates a new RestESender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestESender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestESender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestESender) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *RestESender) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RestESender) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RestESender) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


