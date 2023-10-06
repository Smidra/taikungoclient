# DescribeNodeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDescribeNodeCommand

`func NewDescribeNodeCommand() *DescribeNodeCommand`

NewDescribeNodeCommand instantiates a new DescribeNodeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNodeCommandWithDefaults

`func NewDescribeNodeCommandWithDefaults() *DescribeNodeCommand`

NewDescribeNodeCommandWithDefaults instantiates a new DescribeNodeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DescribeNodeCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeNodeCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeNodeCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DescribeNodeCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *DescribeNodeCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeNodeCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeNodeCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeNodeCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DescribeNodeCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DescribeNodeCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


