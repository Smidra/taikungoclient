# UserGroupDetailsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Users** | Pointer to [**[]UserListDto**](UserListDto.md) |  | [optional] 
**ProjectGroups** | Pointer to [**[]ProjectGroupEntityListDto**](ProjectGroupEntityListDto.md) |  | [optional] 

## Methods

### NewUserGroupDetailsListDto

`func NewUserGroupDetailsListDto() *UserGroupDetailsListDto`

NewUserGroupDetailsListDto instantiates a new UserGroupDetailsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupDetailsListDtoWithDefaults

`func NewUserGroupDetailsListDtoWithDefaults() *UserGroupDetailsListDto`

NewUserGroupDetailsListDtoWithDefaults instantiates a new UserGroupDetailsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupDetailsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupDetailsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupDetailsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupDetailsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserGroupDetailsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupDetailsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupDetailsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroupDetailsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserGroupDetailsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserGroupDetailsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *UserGroupDetailsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserGroupDetailsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserGroupDetailsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UserGroupDetailsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *UserGroupDetailsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *UserGroupDetailsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *UserGroupDetailsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *UserGroupDetailsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *UserGroupDetailsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *UserGroupDetailsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetUsers

`func (o *UserGroupDetailsListDto) GetUsers() []UserListDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserGroupDetailsListDto) GetUsersOk() (*[]UserListDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserGroupDetailsListDto) SetUsers(v []UserListDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserGroupDetailsListDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *UserGroupDetailsListDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *UserGroupDetailsListDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetProjectGroups

`func (o *UserGroupDetailsListDto) GetProjectGroups() []ProjectGroupEntityListDto`

GetProjectGroups returns the ProjectGroups field if non-nil, zero value otherwise.

### GetProjectGroupsOk

`func (o *UserGroupDetailsListDto) GetProjectGroupsOk() (*[]ProjectGroupEntityListDto, bool)`

GetProjectGroupsOk returns a tuple with the ProjectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroups

`func (o *UserGroupDetailsListDto) SetProjectGroups(v []ProjectGroupEntityListDto)`

SetProjectGroups sets ProjectGroups field to given value.

### HasProjectGroups

`func (o *UserGroupDetailsListDto) HasProjectGroups() bool`

HasProjectGroups returns a boolean if a field has been set.

### SetProjectGroupsNil

`func (o *UserGroupDetailsListDto) SetProjectGroupsNil(b bool)`

 SetProjectGroupsNil sets the value for ProjectGroups to be an explicit nil

### UnsetProjectGroups
`func (o *UserGroupDetailsListDto) UnsetProjectGroups()`

UnsetProjectGroups ensures that no value is present for ProjectGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


