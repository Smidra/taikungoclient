# ProjectActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **NullableString** |  | [optional] 
**JobUrl** | Pointer to **NullableString** |  | [optional] 
**TopicName** | Pointer to **NullableString** |  | [optional] 
**EstimatedTime** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectActionDto

`func NewProjectActionDto() *ProjectActionDto`

NewProjectActionDto instantiates a new ProjectActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectActionDtoWithDefaults

`func NewProjectActionDtoWithDefaults() *ProjectActionDto`

NewProjectActionDtoWithDefaults instantiates a new ProjectActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ProjectActionDto) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectActionDto) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectActionDto) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ProjectActionDto) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *ProjectActionDto) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *ProjectActionDto) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetJobUrl

`func (o *ProjectActionDto) GetJobUrl() string`

GetJobUrl returns the JobUrl field if non-nil, zero value otherwise.

### GetJobUrlOk

`func (o *ProjectActionDto) GetJobUrlOk() (*string, bool)`

GetJobUrlOk returns a tuple with the JobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUrl

`func (o *ProjectActionDto) SetJobUrl(v string)`

SetJobUrl sets JobUrl field to given value.

### HasJobUrl

`func (o *ProjectActionDto) HasJobUrl() bool`

HasJobUrl returns a boolean if a field has been set.

### SetJobUrlNil

`func (o *ProjectActionDto) SetJobUrlNil(b bool)`

 SetJobUrlNil sets the value for JobUrl to be an explicit nil

### UnsetJobUrl
`func (o *ProjectActionDto) UnsetJobUrl()`

UnsetJobUrl ensures that no value is present for JobUrl, not even an explicit nil
### GetTopicName

`func (o *ProjectActionDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ProjectActionDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ProjectActionDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *ProjectActionDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *ProjectActionDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *ProjectActionDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetEstimatedTime

`func (o *ProjectActionDto) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *ProjectActionDto) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *ProjectActionDto) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *ProjectActionDto) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### SetEstimatedTimeNil

`func (o *ProjectActionDto) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *ProjectActionDto) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


