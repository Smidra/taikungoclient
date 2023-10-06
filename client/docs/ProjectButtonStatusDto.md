# ProjectButtonStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Reasons** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectButtonStatusDto

`func NewProjectButtonStatusDto() *ProjectButtonStatusDto`

NewProjectButtonStatusDto instantiates a new ProjectButtonStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectButtonStatusDtoWithDefaults

`func NewProjectButtonStatusDtoWithDefaults() *ProjectButtonStatusDto`

NewProjectButtonStatusDtoWithDefaults instantiates a new ProjectButtonStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ProjectButtonStatusDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ProjectButtonStatusDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ProjectButtonStatusDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ProjectButtonStatusDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetReasons

`func (o *ProjectButtonStatusDto) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ProjectButtonStatusDto) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ProjectButtonStatusDto) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ProjectButtonStatusDto) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *ProjectButtonStatusDto) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *ProjectButtonStatusDto) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


