# ProjectAppDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**CatalogName** | Pointer to **NullableString** |  | [optional] 
**CatalogAppName** | Pointer to **NullableString** |  | [optional] 
**AppRepoName** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**Values** | Pointer to **NullableString** |  | [optional] 
**AutoSync** | Pointer to **bool** |  | [optional] 
**ReleaseNotes** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**HelmResult** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**HasJsonSchema** | Pointer to **bool** |  | [optional] 
**ProjectAppParams** | Pointer to [**[]ProjectAppParamDto**](ProjectAppParamDto.md) |  | [optional] 

## Methods

### NewProjectAppDetailsDto

`func NewProjectAppDetailsDto() *ProjectAppDetailsDto`

NewProjectAppDetailsDto instantiates a new ProjectAppDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppDetailsDtoWithDefaults

`func NewProjectAppDetailsDtoWithDefaults() *ProjectAppDetailsDto`

NewProjectAppDetailsDtoWithDefaults instantiates a new ProjectAppDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectAppDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectAppDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectAppDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectAppDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectAppDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAppDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAppDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectAppDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectAppDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectAppDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *ProjectAppDetailsDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProjectAppDetailsDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProjectAppDetailsDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ProjectAppDetailsDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ProjectAppDetailsDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ProjectAppDetailsDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStatus

`func (o *ProjectAppDetailsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectAppDetailsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectAppDetailsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectAppDetailsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProjectAppDetailsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectAppDetailsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVersion

`func (o *ProjectAppDetailsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectAppDetailsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectAppDetailsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectAppDetailsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProjectAppDetailsDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProjectAppDetailsDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCatalogId

`func (o *ProjectAppDetailsDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ProjectAppDetailsDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ProjectAppDetailsDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *ProjectAppDetailsDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogName

`func (o *ProjectAppDetailsDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *ProjectAppDetailsDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *ProjectAppDetailsDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *ProjectAppDetailsDto) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### SetCatalogNameNil

`func (o *ProjectAppDetailsDto) SetCatalogNameNil(b bool)`

 SetCatalogNameNil sets the value for CatalogName to be an explicit nil

### UnsetCatalogName
`func (o *ProjectAppDetailsDto) UnsetCatalogName()`

UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
### GetCatalogAppName

`func (o *ProjectAppDetailsDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *ProjectAppDetailsDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *ProjectAppDetailsDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.

### HasCatalogAppName

`func (o *ProjectAppDetailsDto) HasCatalogAppName() bool`

HasCatalogAppName returns a boolean if a field has been set.

### SetCatalogAppNameNil

`func (o *ProjectAppDetailsDto) SetCatalogAppNameNil(b bool)`

 SetCatalogAppNameNil sets the value for CatalogAppName to be an explicit nil

### UnsetCatalogAppName
`func (o *ProjectAppDetailsDto) UnsetCatalogAppName()`

UnsetCatalogAppName ensures that no value is present for CatalogAppName, not even an explicit nil
### GetAppRepoName

`func (o *ProjectAppDetailsDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *ProjectAppDetailsDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *ProjectAppDetailsDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.

### HasAppRepoName

`func (o *ProjectAppDetailsDto) HasAppRepoName() bool`

HasAppRepoName returns a boolean if a field has been set.

### SetAppRepoNameNil

`func (o *ProjectAppDetailsDto) SetAppRepoNameNil(b bool)`

 SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil

### UnsetAppRepoName
`func (o *ProjectAppDetailsDto) UnsetAppRepoName()`

UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
### GetLogo

`func (o *ProjectAppDetailsDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ProjectAppDetailsDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ProjectAppDetailsDto) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ProjectAppDetailsDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *ProjectAppDetailsDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *ProjectAppDetailsDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetValues

`func (o *ProjectAppDetailsDto) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProjectAppDetailsDto) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProjectAppDetailsDto) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProjectAppDetailsDto) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ProjectAppDetailsDto) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ProjectAppDetailsDto) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetAutoSync

`func (o *ProjectAppDetailsDto) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *ProjectAppDetailsDto) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *ProjectAppDetailsDto) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *ProjectAppDetailsDto) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *ProjectAppDetailsDto) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ProjectAppDetailsDto) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ProjectAppDetailsDto) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *ProjectAppDetailsDto) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### SetReleaseNotesNil

`func (o *ProjectAppDetailsDto) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *ProjectAppDetailsDto) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil
### GetProjectName

`func (o *ProjectAppDetailsDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectAppDetailsDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectAppDetailsDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectAppDetailsDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ProjectAppDetailsDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectAppDetailsDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetHelmResult

`func (o *ProjectAppDetailsDto) GetHelmResult() string`

GetHelmResult returns the HelmResult field if non-nil, zero value otherwise.

### GetHelmResultOk

`func (o *ProjectAppDetailsDto) GetHelmResultOk() (*string, bool)`

GetHelmResultOk returns a tuple with the HelmResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmResult

`func (o *ProjectAppDetailsDto) SetHelmResult(v string)`

SetHelmResult sets HelmResult field to given value.

### HasHelmResult

`func (o *ProjectAppDetailsDto) HasHelmResult() bool`

HasHelmResult returns a boolean if a field has been set.

### SetHelmResultNil

`func (o *ProjectAppDetailsDto) SetHelmResultNil(b bool)`

 SetHelmResultNil sets the value for HelmResult to be an explicit nil

### UnsetHelmResult
`func (o *ProjectAppDetailsDto) UnsetHelmResult()`

UnsetHelmResult ensures that no value is present for HelmResult, not even an explicit nil
### GetProjectId

`func (o *ProjectAppDetailsDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectAppDetailsDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectAppDetailsDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectAppDetailsDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetHasJsonSchema

`func (o *ProjectAppDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *ProjectAppDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *ProjectAppDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.

### HasHasJsonSchema

`func (o *ProjectAppDetailsDto) HasHasJsonSchema() bool`

HasHasJsonSchema returns a boolean if a field has been set.

### GetProjectAppParams

`func (o *ProjectAppDetailsDto) GetProjectAppParams() []ProjectAppParamDto`

GetProjectAppParams returns the ProjectAppParams field if non-nil, zero value otherwise.

### GetProjectAppParamsOk

`func (o *ProjectAppDetailsDto) GetProjectAppParamsOk() (*[]ProjectAppParamDto, bool)`

GetProjectAppParamsOk returns a tuple with the ProjectAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppParams

`func (o *ProjectAppDetailsDto) SetProjectAppParams(v []ProjectAppParamDto)`

SetProjectAppParams sets ProjectAppParams field to given value.

### HasProjectAppParams

`func (o *ProjectAppDetailsDto) HasProjectAppParams() bool`

HasProjectAppParams returns a boolean if a field has been set.

### SetProjectAppParamsNil

`func (o *ProjectAppDetailsDto) SetProjectAppParamsNil(b bool)`

 SetProjectAppParamsNil sets the value for ProjectAppParams to be an explicit nil

### UnsetProjectAppParams
`func (o *ProjectAppDetailsDto) UnsetProjectAppParams()`

UnsetProjectAppParams ensures that no value is present for ProjectAppParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


