# BindUserGroupsToProjectGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroups** | Pointer to [**[]UpdateProjectUserGroupDto**](UpdateProjectUserGroupDto.md) |  | [optional] 
**ProjectGroupId** | Pointer to **int32** |  | [optional] 
**ProjectGroupName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBindUserGroupsToProjectGroupCommand

`func NewBindUserGroupsToProjectGroupCommand() *BindUserGroupsToProjectGroupCommand`

NewBindUserGroupsToProjectGroupCommand instantiates a new BindUserGroupsToProjectGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindUserGroupsToProjectGroupCommandWithDefaults

`func NewBindUserGroupsToProjectGroupCommandWithDefaults() *BindUserGroupsToProjectGroupCommand`

NewBindUserGroupsToProjectGroupCommandWithDefaults instantiates a new BindUserGroupsToProjectGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) GetUserGroups() []UpdateProjectUserGroupDto`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *BindUserGroupsToProjectGroupCommand) GetUserGroupsOk() (*[]UpdateProjectUserGroupDto, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) SetUserGroups(v []UpdateProjectUserGroupDto)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *BindUserGroupsToProjectGroupCommand) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroupsNil

`func (o *BindUserGroupsToProjectGroupCommand) SetUserGroupsNil(b bool)`

 SetUserGroupsNil sets the value for UserGroups to be an explicit nil

### UnsetUserGroups
`func (o *BindUserGroupsToProjectGroupCommand) UnsetUserGroups()`

UnsetUserGroups ensures that no value is present for UserGroups, not even an explicit nil
### GetProjectGroupId

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupId() int32`

GetProjectGroupId returns the ProjectGroupId field if non-nil, zero value otherwise.

### GetProjectGroupIdOk

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupIdOk() (*int32, bool)`

GetProjectGroupIdOk returns a tuple with the ProjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupId

`func (o *BindUserGroupsToProjectGroupCommand) SetProjectGroupId(v int32)`

SetProjectGroupId sets ProjectGroupId field to given value.

### HasProjectGroupId

`func (o *BindUserGroupsToProjectGroupCommand) HasProjectGroupId() bool`

HasProjectGroupId returns a boolean if a field has been set.

### GetProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupName() string`

GetProjectGroupName returns the ProjectGroupName field if non-nil, zero value otherwise.

### GetProjectGroupNameOk

`func (o *BindUserGroupsToProjectGroupCommand) GetProjectGroupNameOk() (*string, bool)`

GetProjectGroupNameOk returns a tuple with the ProjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) SetProjectGroupName(v string)`

SetProjectGroupName sets ProjectGroupName field to given value.

### HasProjectGroupName

`func (o *BindUserGroupsToProjectGroupCommand) HasProjectGroupName() bool`

HasProjectGroupName returns a boolean if a field has been set.

### SetProjectGroupNameNil

`func (o *BindUserGroupsToProjectGroupCommand) SetProjectGroupNameNil(b bool)`

 SetProjectGroupNameNil sets the value for ProjectGroupName to be an explicit nil

### UnsetProjectGroupName
`func (o *BindUserGroupsToProjectGroupCommand) UnsetProjectGroupName()`

UnsetProjectGroupName ensures that no value is present for ProjectGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


