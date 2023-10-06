# ProjectGroupDetailsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to [**[]ProjectListDto**](ProjectListDto.md) |  | [optional] 
**UserGroups** | Pointer to [**[]UserGroupEntityListDto**](UserGroupEntityListDto.md) |  | [optional] 

## Methods

### NewProjectGroupDetailsListDto

`func NewProjectGroupDetailsListDto() *ProjectGroupDetailsListDto`

NewProjectGroupDetailsListDto instantiates a new ProjectGroupDetailsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectGroupDetailsListDtoWithDefaults

`func NewProjectGroupDetailsListDtoWithDefaults() *ProjectGroupDetailsListDto`

NewProjectGroupDetailsListDtoWithDefaults instantiates a new ProjectGroupDetailsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectGroupDetailsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectGroupDetailsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectGroupDetailsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectGroupDetailsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectGroupDetailsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectGroupDetailsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectGroupDetailsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectGroupDetailsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectGroupDetailsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectGroupDetailsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *ProjectGroupDetailsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectGroupDetailsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectGroupDetailsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectGroupDetailsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProjectGroupDetailsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectGroupDetailsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectGroupDetailsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectGroupDetailsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ProjectGroupDetailsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectGroupDetailsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetProjects

`func (o *ProjectGroupDetailsListDto) GetProjects() []ProjectListDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectGroupDetailsListDto) GetProjectsOk() (*[]ProjectListDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectGroupDetailsListDto) SetProjects(v []ProjectListDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectGroupDetailsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *ProjectGroupDetailsListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *ProjectGroupDetailsListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetUserGroups

`func (o *ProjectGroupDetailsListDto) GetUserGroups() []UserGroupEntityListDto`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ProjectGroupDetailsListDto) GetUserGroupsOk() (*[]UserGroupEntityListDto, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ProjectGroupDetailsListDto) SetUserGroups(v []UserGroupEntityListDto)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ProjectGroupDetailsListDto) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroupsNil

`func (o *ProjectGroupDetailsListDto) SetUserGroupsNil(b bool)`

 SetUserGroupsNil sets the value for UserGroups to be an explicit nil

### UnsetUserGroups
`func (o *ProjectGroupDetailsListDto) UnsetUserGroups()`

UnsetUserGroups ensures that no value is present for UserGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


