# DeleteUserGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewDeleteUserGroupCommand

`func NewDeleteUserGroupCommand() *DeleteUserGroupCommand`

NewDeleteUserGroupCommand instantiates a new DeleteUserGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserGroupCommandWithDefaults

`func NewDeleteUserGroupCommandWithDefaults() *DeleteUserGroupCommand`

NewDeleteUserGroupCommandWithDefaults instantiates a new DeleteUserGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroupIds

`func (o *DeleteUserGroupCommand) GetUserGroupIds() []int32`

GetUserGroupIds returns the UserGroupIds field if non-nil, zero value otherwise.

### GetUserGroupIdsOk

`func (o *DeleteUserGroupCommand) GetUserGroupIdsOk() (*[]int32, bool)`

GetUserGroupIdsOk returns a tuple with the UserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupIds

`func (o *DeleteUserGroupCommand) SetUserGroupIds(v []int32)`

SetUserGroupIds sets UserGroupIds field to given value.

### HasUserGroupIds

`func (o *DeleteUserGroupCommand) HasUserGroupIds() bool`

HasUserGroupIds returns a boolean if a field has been set.

### SetUserGroupIdsNil

`func (o *DeleteUserGroupCommand) SetUserGroupIdsNil(b bool)`

 SetUserGroupIdsNil sets the value for UserGroupIds to be an explicit nil

### UnsetUserGroupIds
`func (o *DeleteUserGroupCommand) UnsetUserGroupIds()`

UnsetUserGroupIds ensures that no value is present for UserGroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


