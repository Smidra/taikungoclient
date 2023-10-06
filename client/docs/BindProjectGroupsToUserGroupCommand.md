# BindProjectGroupsToUserGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectGroups** | Pointer to [**[]UpdateUserProjectGroupDto**](UpdateUserProjectGroupDto.md) |  | [optional] 
**UserGroupId** | Pointer to **int32** |  | [optional] 
**UserGroupName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBindProjectGroupsToUserGroupCommand

`func NewBindProjectGroupsToUserGroupCommand() *BindProjectGroupsToUserGroupCommand`

NewBindProjectGroupsToUserGroupCommand instantiates a new BindProjectGroupsToUserGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindProjectGroupsToUserGroupCommandWithDefaults

`func NewBindProjectGroupsToUserGroupCommandWithDefaults() *BindProjectGroupsToUserGroupCommand`

NewBindProjectGroupsToUserGroupCommandWithDefaults instantiates a new BindProjectGroupsToUserGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectGroups

`func (o *BindProjectGroupsToUserGroupCommand) GetProjectGroups() []UpdateUserProjectGroupDto`

GetProjectGroups returns the ProjectGroups field if non-nil, zero value otherwise.

### GetProjectGroupsOk

`func (o *BindProjectGroupsToUserGroupCommand) GetProjectGroupsOk() (*[]UpdateUserProjectGroupDto, bool)`

GetProjectGroupsOk returns a tuple with the ProjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroups

`func (o *BindProjectGroupsToUserGroupCommand) SetProjectGroups(v []UpdateUserProjectGroupDto)`

SetProjectGroups sets ProjectGroups field to given value.

### HasProjectGroups

`func (o *BindProjectGroupsToUserGroupCommand) HasProjectGroups() bool`

HasProjectGroups returns a boolean if a field has been set.

### SetProjectGroupsNil

`func (o *BindProjectGroupsToUserGroupCommand) SetProjectGroupsNil(b bool)`

 SetProjectGroupsNil sets the value for ProjectGroups to be an explicit nil

### UnsetProjectGroups
`func (o *BindProjectGroupsToUserGroupCommand) UnsetProjectGroups()`

UnsetProjectGroups ensures that no value is present for ProjectGroups, not even an explicit nil
### GetUserGroupId

`func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupId() int32`

GetUserGroupId returns the UserGroupId field if non-nil, zero value otherwise.

### GetUserGroupIdOk

`func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupIdOk() (*int32, bool)`

GetUserGroupIdOk returns a tuple with the UserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupId

`func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupId(v int32)`

SetUserGroupId sets UserGroupId field to given value.

### HasUserGroupId

`func (o *BindProjectGroupsToUserGroupCommand) HasUserGroupId() bool`

HasUserGroupId returns a boolean if a field has been set.

### GetUserGroupName

`func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *BindProjectGroupsToUserGroupCommand) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### SetUserGroupNameNil

`func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupNameNil(b bool)`

 SetUserGroupNameNil sets the value for UserGroupName to be an explicit nil

### UnsetUserGroupName
`func (o *BindProjectGroupsToUserGroupCommand) UnsetUserGroupName()`

UnsetUserGroupName ensures that no value is present for UserGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


