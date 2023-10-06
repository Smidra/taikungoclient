# DeleteRepositoryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeleteRepositoryCommand

`func NewDeleteRepositoryCommand() *DeleteRepositoryCommand`

NewDeleteRepositoryCommand instantiates a new DeleteRepositoryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRepositoryCommandWithDefaults

`func NewDeleteRepositoryCommandWithDefaults() *DeleteRepositoryCommand`

NewDeleteRepositoryCommandWithDefaults instantiates a new DeleteRepositoryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoName

`func (o *DeleteRepositoryCommand) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *DeleteRepositoryCommand) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *DeleteRepositoryCommand) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *DeleteRepositoryCommand) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### SetRepoNameNil

`func (o *DeleteRepositoryCommand) SetRepoNameNil(b bool)`

 SetRepoNameNil sets the value for RepoName to be an explicit nil

### UnsetRepoName
`func (o *DeleteRepositoryCommand) UnsetRepoName()`

UnsetRepoName ensures that no value is present for RepoName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


