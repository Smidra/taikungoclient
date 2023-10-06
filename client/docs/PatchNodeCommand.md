# PatchNodeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]PatchNodeLabelsDto**](PatchNodeLabelsDto.md) |  | [optional] 

## Methods

### NewPatchNodeCommand

`func NewPatchNodeCommand() *PatchNodeCommand`

NewPatchNodeCommand instantiates a new PatchNodeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNodeCommandWithDefaults

`func NewPatchNodeCommandWithDefaults() *PatchNodeCommand`

NewPatchNodeCommandWithDefaults instantiates a new PatchNodeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PatchNodeCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PatchNodeCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PatchNodeCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PatchNodeCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *PatchNodeCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchNodeCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchNodeCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchNodeCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchNodeCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchNodeCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParameters

`func (o *PatchNodeCommand) GetParameters() []PatchNodeLabelsDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchNodeCommand) GetParametersOk() (*[]PatchNodeLabelsDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchNodeCommand) SetParameters(v []PatchNodeLabelsDto)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchNodeCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchNodeCommand) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchNodeCommand) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


