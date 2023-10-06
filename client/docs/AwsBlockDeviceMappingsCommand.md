# AwsBlockDeviceMappingsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAwsBlockDeviceMappingsCommand

`func NewAwsBlockDeviceMappingsCommand() *AwsBlockDeviceMappingsCommand`

NewAwsBlockDeviceMappingsCommand instantiates a new AwsBlockDeviceMappingsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsBlockDeviceMappingsCommandWithDefaults

`func NewAwsBlockDeviceMappingsCommandWithDefaults() *AwsBlockDeviceMappingsCommand`

NewAwsBlockDeviceMappingsCommandWithDefaults instantiates a new AwsBlockDeviceMappingsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AwsBlockDeviceMappingsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AwsBlockDeviceMappingsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AwsBlockDeviceMappingsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AwsBlockDeviceMappingsCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetImageId

`func (o *AwsBlockDeviceMappingsCommand) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AwsBlockDeviceMappingsCommand) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AwsBlockDeviceMappingsCommand) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AwsBlockDeviceMappingsCommand) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *AwsBlockDeviceMappingsCommand) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *AwsBlockDeviceMappingsCommand) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


