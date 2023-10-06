# ServerCommonRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**Names** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerCommonRecordDto

`func NewServerCommonRecordDto() *ServerCommonRecordDto`

NewServerCommonRecordDto instantiates a new ServerCommonRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCommonRecordDtoWithDefaults

`func NewServerCommonRecordDtoWithDefaults() *ServerCommonRecordDto`

NewServerCommonRecordDtoWithDefaults instantiates a new ServerCommonRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ServerCommonRecordDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServerCommonRecordDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServerCommonRecordDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServerCommonRecordDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *ServerCommonRecordDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ServerCommonRecordDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ServerCommonRecordDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ServerCommonRecordDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ServerCommonRecordDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ServerCommonRecordDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetNames

`func (o *ServerCommonRecordDto) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ServerCommonRecordDto) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ServerCommonRecordDto) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *ServerCommonRecordDto) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNamesNil

`func (o *ServerCommonRecordDto) SetNamesNil(b bool)`

 SetNamesNil sets the value for Names to be an explicit nil

### UnsetNames
`func (o *ServerCommonRecordDto) UnsetNames()`

UnsetNames ensures that no value is present for Names, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


