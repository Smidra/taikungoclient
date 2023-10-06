# BindUsersCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]UpdateProjectUserDto**](UpdateProjectUserDto.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBindUsersCommand

`func NewBindUsersCommand() *BindUsersCommand`

NewBindUsersCommand instantiates a new BindUsersCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindUsersCommandWithDefaults

`func NewBindUsersCommandWithDefaults() *BindUsersCommand`

NewBindUsersCommandWithDefaults instantiates a new BindUsersCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *BindUsersCommand) GetUsers() []UpdateProjectUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BindUsersCommand) GetUsersOk() (*[]UpdateProjectUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BindUsersCommand) SetUsers(v []UpdateProjectUserDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BindUsersCommand) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *BindUsersCommand) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *BindUsersCommand) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetProjectId

`func (o *BindUsersCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BindUsersCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BindUsersCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BindUsersCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


