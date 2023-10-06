# CreateProjectFromTemplateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**CanCommit** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateProjectFromTemplateCommand

`func NewCreateProjectFromTemplateCommand() *CreateProjectFromTemplateCommand`

NewCreateProjectFromTemplateCommand instantiates a new CreateProjectFromTemplateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectFromTemplateCommandWithDefaults

`func NewCreateProjectFromTemplateCommandWithDefaults() *CreateProjectFromTemplateCommand`

NewCreateProjectFromTemplateCommandWithDefaults instantiates a new CreateProjectFromTemplateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateProjectFromTemplateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateProjectFromTemplateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateProjectFromTemplateCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateProjectFromTemplateCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateProjectFromTemplateCommand) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateProjectFromTemplateCommand) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateProjectFromTemplateCommand) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *CreateProjectFromTemplateCommand) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *CreateProjectFromTemplateCommand) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *CreateProjectFromTemplateCommand) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetCanCommit

`func (o *CreateProjectFromTemplateCommand) GetCanCommit() bool`

GetCanCommit returns the CanCommit field if non-nil, zero value otherwise.

### GetCanCommitOk

`func (o *CreateProjectFromTemplateCommand) GetCanCommitOk() (*bool, bool)`

GetCanCommitOk returns a tuple with the CanCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCommit

`func (o *CreateProjectFromTemplateCommand) SetCanCommit(v bool)`

SetCanCommit sets CanCommit field to given value.

### HasCanCommit

`func (o *CreateProjectFromTemplateCommand) HasCanCommit() bool`

HasCanCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


