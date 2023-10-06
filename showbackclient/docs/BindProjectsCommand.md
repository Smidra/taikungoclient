# BindProjectsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]UpdateUserProjectDto**](UpdateUserProjectDto.md) |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBindProjectsCommand

`func NewBindProjectsCommand() *BindProjectsCommand`

NewBindProjectsCommand instantiates a new BindProjectsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindProjectsCommandWithDefaults

`func NewBindProjectsCommandWithDefaults() *BindProjectsCommand`

NewBindProjectsCommandWithDefaults instantiates a new BindProjectsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *BindProjectsCommand) GetProjects() []UpdateUserProjectDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BindProjectsCommand) GetProjectsOk() (*[]UpdateUserProjectDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BindProjectsCommand) SetProjects(v []UpdateUserProjectDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *BindProjectsCommand) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *BindProjectsCommand) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *BindProjectsCommand) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetUserId

`func (o *BindProjectsCommand) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BindProjectsCommand) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BindProjectsCommand) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BindProjectsCommand) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BindProjectsCommand) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BindProjectsCommand) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *BindProjectsCommand) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *BindProjectsCommand) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *BindProjectsCommand) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *BindProjectsCommand) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *BindProjectsCommand) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *BindProjectsCommand) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


