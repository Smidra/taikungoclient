# UpdateProjectGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUpdateProjectGroupDto

`func NewUpdateProjectGroupDto() *UpdateProjectGroupDto`

NewUpdateProjectGroupDto instantiates a new UpdateProjectGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectGroupDtoWithDefaults

`func NewUpdateProjectGroupDtoWithDefaults() *UpdateProjectGroupDto`

NewUpdateProjectGroupDtoWithDefaults instantiates a new UpdateProjectGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProjectGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProjectGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateProjectGroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateProjectGroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectIds

`func (o *UpdateProjectGroupDto) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateProjectGroupDto) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateProjectGroupDto) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *UpdateProjectGroupDto) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *UpdateProjectGroupDto) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *UpdateProjectGroupDto) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


