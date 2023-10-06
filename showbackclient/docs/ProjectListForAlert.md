# ProjectListForAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Health** | Pointer to **NullableString** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**HasKubeConfigFile** | Pointer to **bool** |  | [optional] 
**MonitoringCredential** | Pointer to [**MonitoringCredentialsListDto**](MonitoringCredentialsListDto.md) |  | [optional] 
**KubernetesAlerts** | Pointer to [**[]KubernetesAlertDto**](KubernetesAlertDto.md) |  | [optional] 

## Methods

### NewProjectListForAlert

`func NewProjectListForAlert() *ProjectListForAlert`

NewProjectListForAlert instantiates a new ProjectListForAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListForAlertWithDefaults

`func NewProjectListForAlertWithDefaults() *ProjectListForAlert`

NewProjectListForAlertWithDefaults instantiates a new ProjectListForAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectListForAlert) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectListForAlert) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectListForAlert) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectListForAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectListForAlert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectListForAlert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectListForAlert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectListForAlert) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectListForAlert) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectListForAlert) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetToken

`func (o *ProjectListForAlert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProjectListForAlert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProjectListForAlert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProjectListForAlert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ProjectListForAlert) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ProjectListForAlert) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetStatus

`func (o *ProjectListForAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListForAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListForAlert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectListForAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProjectListForAlert) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectListForAlert) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOrganizationId

`func (o *ProjectListForAlert) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectListForAlert) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectListForAlert) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectListForAlert) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectListForAlert) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectListForAlert) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectListForAlert) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectListForAlert) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### SetHealthNil

`func (o *ProjectListForAlert) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *ProjectListForAlert) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil
### GetIsKubernetes

`func (o *ProjectListForAlert) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectListForAlert) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectListForAlert) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *ProjectListForAlert) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectListForAlert) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectListForAlert) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectListForAlert) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectListForAlert) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectListForAlert) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectListForAlert) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectListForAlert) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectListForAlert) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetHasKubeConfigFile

`func (o *ProjectListForAlert) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectListForAlert) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectListForAlert) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.

### HasHasKubeConfigFile

`func (o *ProjectListForAlert) HasHasKubeConfigFile() bool`

HasHasKubeConfigFile returns a boolean if a field has been set.

### GetMonitoringCredential

`func (o *ProjectListForAlert) GetMonitoringCredential() MonitoringCredentialsListDto`

GetMonitoringCredential returns the MonitoringCredential field if non-nil, zero value otherwise.

### GetMonitoringCredentialOk

`func (o *ProjectListForAlert) GetMonitoringCredentialOk() (*MonitoringCredentialsListDto, bool)`

GetMonitoringCredentialOk returns a tuple with the MonitoringCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringCredential

`func (o *ProjectListForAlert) SetMonitoringCredential(v MonitoringCredentialsListDto)`

SetMonitoringCredential sets MonitoringCredential field to given value.

### HasMonitoringCredential

`func (o *ProjectListForAlert) HasMonitoringCredential() bool`

HasMonitoringCredential returns a boolean if a field has been set.

### GetKubernetesAlerts

`func (o *ProjectListForAlert) GetKubernetesAlerts() []KubernetesAlertDto`

GetKubernetesAlerts returns the KubernetesAlerts field if non-nil, zero value otherwise.

### GetKubernetesAlertsOk

`func (o *ProjectListForAlert) GetKubernetesAlertsOk() (*[]KubernetesAlertDto, bool)`

GetKubernetesAlertsOk returns a tuple with the KubernetesAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesAlerts

`func (o *ProjectListForAlert) SetKubernetesAlerts(v []KubernetesAlertDto)`

SetKubernetesAlerts sets KubernetesAlerts field to given value.

### HasKubernetesAlerts

`func (o *ProjectListForAlert) HasKubernetesAlerts() bool`

HasKubernetesAlerts returns a boolean if a field has been set.

### SetKubernetesAlertsNil

`func (o *ProjectListForAlert) SetKubernetesAlertsNil(b bool)`

 SetKubernetesAlertsNil sets the value for KubernetesAlerts to be an explicit nil

### UnsetKubernetesAlerts
`func (o *ProjectListForAlert) UnsetKubernetesAlerts()`

UnsetKubernetesAlerts ensures that no value is present for KubernetesAlerts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


