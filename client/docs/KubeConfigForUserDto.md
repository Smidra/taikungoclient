# KubeConfigForUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**IsAccessibleForAll** | Pointer to **bool** |  | [optional] 
**IsAccessibleForManager** | Pointer to **bool** |  | [optional] 
**KubeConfigRoleName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**ExpirationDate** | Pointer to **NullableString** |  | [optional] 
**CanDownload** | Pointer to **bool** |  | [optional] 
**CanAccessTerminal** | Pointer to **bool** |  | [optional] 
**CanDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubeConfigForUserDto

`func NewKubeConfigForUserDto() *KubeConfigForUserDto`

NewKubeConfigForUserDto instantiates a new KubeConfigForUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigForUserDtoWithDefaults

`func NewKubeConfigForUserDtoWithDefaults() *KubeConfigForUserDto`

NewKubeConfigForUserDtoWithDefaults instantiates a new KubeConfigForUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubeConfigForUserDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubeConfigForUserDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubeConfigForUserDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubeConfigForUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *KubeConfigForUserDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KubeConfigForUserDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KubeConfigForUserDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KubeConfigForUserDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *KubeConfigForUserDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *KubeConfigForUserDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDisplayName

`func (o *KubeConfigForUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KubeConfigForUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KubeConfigForUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KubeConfigForUserDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KubeConfigForUserDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KubeConfigForUserDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetProjectId

`func (o *KubeConfigForUserDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubeConfigForUserDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubeConfigForUserDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KubeConfigForUserDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *KubeConfigForUserDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KubeConfigForUserDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KubeConfigForUserDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *KubeConfigForUserDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectName

`func (o *KubeConfigForUserDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KubeConfigForUserDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KubeConfigForUserDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *KubeConfigForUserDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *KubeConfigForUserDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *KubeConfigForUserDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsAccessibleForAll

`func (o *KubeConfigForUserDto) GetIsAccessibleForAll() bool`

GetIsAccessibleForAll returns the IsAccessibleForAll field if non-nil, zero value otherwise.

### GetIsAccessibleForAllOk

`func (o *KubeConfigForUserDto) GetIsAccessibleForAllOk() (*bool, bool)`

GetIsAccessibleForAllOk returns a tuple with the IsAccessibleForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForAll

`func (o *KubeConfigForUserDto) SetIsAccessibleForAll(v bool)`

SetIsAccessibleForAll sets IsAccessibleForAll field to given value.

### HasIsAccessibleForAll

`func (o *KubeConfigForUserDto) HasIsAccessibleForAll() bool`

HasIsAccessibleForAll returns a boolean if a field has been set.

### GetIsAccessibleForManager

`func (o *KubeConfigForUserDto) GetIsAccessibleForManager() bool`

GetIsAccessibleForManager returns the IsAccessibleForManager field if non-nil, zero value otherwise.

### GetIsAccessibleForManagerOk

`func (o *KubeConfigForUserDto) GetIsAccessibleForManagerOk() (*bool, bool)`

GetIsAccessibleForManagerOk returns a tuple with the IsAccessibleForManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForManager

`func (o *KubeConfigForUserDto) SetIsAccessibleForManager(v bool)`

SetIsAccessibleForManager sets IsAccessibleForManager field to given value.

### HasIsAccessibleForManager

`func (o *KubeConfigForUserDto) HasIsAccessibleForManager() bool`

HasIsAccessibleForManager returns a boolean if a field has been set.

### GetKubeConfigRoleName

`func (o *KubeConfigForUserDto) GetKubeConfigRoleName() string`

GetKubeConfigRoleName returns the KubeConfigRoleName field if non-nil, zero value otherwise.

### GetKubeConfigRoleNameOk

`func (o *KubeConfigForUserDto) GetKubeConfigRoleNameOk() (*string, bool)`

GetKubeConfigRoleNameOk returns a tuple with the KubeConfigRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigRoleName

`func (o *KubeConfigForUserDto) SetKubeConfigRoleName(v string)`

SetKubeConfigRoleName sets KubeConfigRoleName field to given value.

### HasKubeConfigRoleName

`func (o *KubeConfigForUserDto) HasKubeConfigRoleName() bool`

HasKubeConfigRoleName returns a boolean if a field has been set.

### SetKubeConfigRoleNameNil

`func (o *KubeConfigForUserDto) SetKubeConfigRoleNameNil(b bool)`

 SetKubeConfigRoleNameNil sets the value for KubeConfigRoleName to be an explicit nil

### UnsetKubeConfigRoleName
`func (o *KubeConfigForUserDto) UnsetKubeConfigRoleName()`

UnsetKubeConfigRoleName ensures that no value is present for KubeConfigRoleName, not even an explicit nil
### GetCreatedBy

`func (o *KubeConfigForUserDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KubeConfigForUserDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KubeConfigForUserDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KubeConfigForUserDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *KubeConfigForUserDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *KubeConfigForUserDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *KubeConfigForUserDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubeConfigForUserDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubeConfigForUserDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubeConfigForUserDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *KubeConfigForUserDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *KubeConfigForUserDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetNamespace

`func (o *KubeConfigForUserDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubeConfigForUserDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubeConfigForUserDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubeConfigForUserDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *KubeConfigForUserDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *KubeConfigForUserDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetExpirationDate

`func (o *KubeConfigForUserDto) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KubeConfigForUserDto) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KubeConfigForUserDto) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KubeConfigForUserDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *KubeConfigForUserDto) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *KubeConfigForUserDto) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCanDownload

`func (o *KubeConfigForUserDto) GetCanDownload() bool`

GetCanDownload returns the CanDownload field if non-nil, zero value otherwise.

### GetCanDownloadOk

`func (o *KubeConfigForUserDto) GetCanDownloadOk() (*bool, bool)`

GetCanDownloadOk returns a tuple with the CanDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDownload

`func (o *KubeConfigForUserDto) SetCanDownload(v bool)`

SetCanDownload sets CanDownload field to given value.

### HasCanDownload

`func (o *KubeConfigForUserDto) HasCanDownload() bool`

HasCanDownload returns a boolean if a field has been set.

### GetCanAccessTerminal

`func (o *KubeConfigForUserDto) GetCanAccessTerminal() bool`

GetCanAccessTerminal returns the CanAccessTerminal field if non-nil, zero value otherwise.

### GetCanAccessTerminalOk

`func (o *KubeConfigForUserDto) GetCanAccessTerminalOk() (*bool, bool)`

GetCanAccessTerminalOk returns a tuple with the CanAccessTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessTerminal

`func (o *KubeConfigForUserDto) SetCanAccessTerminal(v bool)`

SetCanAccessTerminal sets CanAccessTerminal field to given value.

### HasCanAccessTerminal

`func (o *KubeConfigForUserDto) HasCanAccessTerminal() bool`

HasCanAccessTerminal returns a boolean if a field has been set.

### GetCanDelete

`func (o *KubeConfigForUserDto) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *KubeConfigForUserDto) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *KubeConfigForUserDto) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *KubeConfigForUserDto) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


