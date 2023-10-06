# PatchCrdCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Yaml** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchCrdCommand

`func NewPatchCrdCommand() *PatchCrdCommand`

NewPatchCrdCommand instantiates a new PatchCrdCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCrdCommandWithDefaults

`func NewPatchCrdCommandWithDefaults() *PatchCrdCommand`

NewPatchCrdCommandWithDefaults instantiates a new PatchCrdCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PatchCrdCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PatchCrdCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PatchCrdCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PatchCrdCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetYaml

`func (o *PatchCrdCommand) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *PatchCrdCommand) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *PatchCrdCommand) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *PatchCrdCommand) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### SetYamlNil

`func (o *PatchCrdCommand) SetYamlNil(b bool)`

 SetYamlNil sets the value for Yaml to be an explicit nil

### UnsetYaml
`func (o *PatchCrdCommand) UnsetYaml()`

UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
### GetName

`func (o *PatchCrdCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchCrdCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchCrdCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchCrdCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchCrdCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchCrdCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


