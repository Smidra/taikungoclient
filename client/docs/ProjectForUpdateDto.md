# ProjectForUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | Pointer to **NullableString** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**AccessIp** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] 
**MonitoringIp** | Pointer to **NullableString** |  | [optional] 
**TanzuReleaseCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**ProjectStatus**](ProjectStatus.md) |  | [optional] 
**Health** | Pointer to [**ProjectHealth**](ProjectHealth.md) |  | [optional] 
**IsBackupEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsOpaEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsAiEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **NullableBool** |  | [optional] 
**AppEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsKubevapEnabled** | Pointer to **NullableBool** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectForUpdateDto

`func NewProjectForUpdateDto() *ProjectForUpdateDto`

NewProjectForUpdateDto instantiates a new ProjectForUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectForUpdateDtoWithDefaults

`func NewProjectForUpdateDtoWithDefaults() *ProjectForUpdateDto`

NewProjectForUpdateDtoWithDefaults instantiates a new ProjectForUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *ProjectForUpdateDto) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *ProjectForUpdateDto) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *ProjectForUpdateDto) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *ProjectForUpdateDto) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *ProjectForUpdateDto) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *ProjectForUpdateDto) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectForUpdateDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *ProjectForUpdateDto) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### SetKubesprayCurrentVersionNil

`func (o *ProjectForUpdateDto) SetKubesprayCurrentVersionNil(b bool)`

 SetKubesprayCurrentVersionNil sets the value for KubesprayCurrentVersion to be an explicit nil

### UnsetKubesprayCurrentVersion
`func (o *ProjectForUpdateDto) UnsetKubesprayCurrentVersion()`

UnsetKubesprayCurrentVersion ensures that no value is present for KubesprayCurrentVersion, not even an explicit nil
### GetAccessIp

`func (o *ProjectForUpdateDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectForUpdateDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectForUpdateDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectForUpdateDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### SetAccessIpNil

`func (o *ProjectForUpdateDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *ProjectForUpdateDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetImageName

`func (o *ProjectForUpdateDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProjectForUpdateDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProjectForUpdateDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProjectForUpdateDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *ProjectForUpdateDto) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *ProjectForUpdateDto) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetMonitoringIp

`func (o *ProjectForUpdateDto) GetMonitoringIp() string`

GetMonitoringIp returns the MonitoringIp field if non-nil, zero value otherwise.

### GetMonitoringIpOk

`func (o *ProjectForUpdateDto) GetMonitoringIpOk() (*string, bool)`

GetMonitoringIpOk returns a tuple with the MonitoringIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringIp

`func (o *ProjectForUpdateDto) SetMonitoringIp(v string)`

SetMonitoringIp sets MonitoringIp field to given value.

### HasMonitoringIp

`func (o *ProjectForUpdateDto) HasMonitoringIp() bool`

HasMonitoringIp returns a boolean if a field has been set.

### SetMonitoringIpNil

`func (o *ProjectForUpdateDto) SetMonitoringIpNil(b bool)`

 SetMonitoringIpNil sets the value for MonitoringIp to be an explicit nil

### UnsetMonitoringIp
`func (o *ProjectForUpdateDto) UnsetMonitoringIp()`

UnsetMonitoringIp ensures that no value is present for MonitoringIp, not even an explicit nil
### GetTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) GetTanzuReleaseCurrentVersion() string`

GetTanzuReleaseCurrentVersion returns the TanzuReleaseCurrentVersion field if non-nil, zero value otherwise.

### GetTanzuReleaseCurrentVersionOk

`func (o *ProjectForUpdateDto) GetTanzuReleaseCurrentVersionOk() (*string, bool)`

GetTanzuReleaseCurrentVersionOk returns a tuple with the TanzuReleaseCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) SetTanzuReleaseCurrentVersion(v string)`

SetTanzuReleaseCurrentVersion sets TanzuReleaseCurrentVersion field to given value.

### HasTanzuReleaseCurrentVersion

`func (o *ProjectForUpdateDto) HasTanzuReleaseCurrentVersion() bool`

HasTanzuReleaseCurrentVersion returns a boolean if a field has been set.

### SetTanzuReleaseCurrentVersionNil

`func (o *ProjectForUpdateDto) SetTanzuReleaseCurrentVersionNil(b bool)`

 SetTanzuReleaseCurrentVersionNil sets the value for TanzuReleaseCurrentVersion to be an explicit nil

### UnsetTanzuReleaseCurrentVersion
`func (o *ProjectForUpdateDto) UnsetTanzuReleaseCurrentVersion()`

UnsetTanzuReleaseCurrentVersion ensures that no value is present for TanzuReleaseCurrentVersion, not even an explicit nil
### GetStatus

`func (o *ProjectForUpdateDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectForUpdateDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectForUpdateDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectForUpdateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHealth

`func (o *ProjectForUpdateDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectForUpdateDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectForUpdateDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectForUpdateDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ProjectForUpdateDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectForUpdateDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectForUpdateDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectForUpdateDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### SetIsBackupEnabledNil

`func (o *ProjectForUpdateDto) SetIsBackupEnabledNil(b bool)`

 SetIsBackupEnabledNil sets the value for IsBackupEnabled to be an explicit nil

### UnsetIsBackupEnabled
`func (o *ProjectForUpdateDto) UnsetIsBackupEnabled()`

UnsetIsBackupEnabled ensures that no value is present for IsBackupEnabled, not even an explicit nil
### GetIsMonitoringEnabled

`func (o *ProjectForUpdateDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectForUpdateDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectForUpdateDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectForUpdateDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### SetIsMonitoringEnabledNil

`func (o *ProjectForUpdateDto) SetIsMonitoringEnabledNil(b bool)`

 SetIsMonitoringEnabledNil sets the value for IsMonitoringEnabled to be an explicit nil

### UnsetIsMonitoringEnabled
`func (o *ProjectForUpdateDto) UnsetIsMonitoringEnabled()`

UnsetIsMonitoringEnabled ensures that no value is present for IsMonitoringEnabled, not even an explicit nil
### GetIsOpaEnabled

`func (o *ProjectForUpdateDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectForUpdateDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectForUpdateDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.

### HasIsOpaEnabled

`func (o *ProjectForUpdateDto) HasIsOpaEnabled() bool`

HasIsOpaEnabled returns a boolean if a field has been set.

### SetIsOpaEnabledNil

`func (o *ProjectForUpdateDto) SetIsOpaEnabledNil(b bool)`

 SetIsOpaEnabledNil sets the value for IsOpaEnabled to be an explicit nil

### UnsetIsOpaEnabled
`func (o *ProjectForUpdateDto) UnsetIsOpaEnabled()`

UnsetIsOpaEnabled ensures that no value is present for IsOpaEnabled, not even an explicit nil
### GetIsAiEnabled

`func (o *ProjectForUpdateDto) GetIsAiEnabled() bool`

GetIsAiEnabled returns the IsAiEnabled field if non-nil, zero value otherwise.

### GetIsAiEnabledOk

`func (o *ProjectForUpdateDto) GetIsAiEnabledOk() (*bool, bool)`

GetIsAiEnabledOk returns a tuple with the IsAiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiEnabled

`func (o *ProjectForUpdateDto) SetIsAiEnabled(v bool)`

SetIsAiEnabled sets IsAiEnabled field to given value.

### HasIsAiEnabled

`func (o *ProjectForUpdateDto) HasIsAiEnabled() bool`

HasIsAiEnabled returns a boolean if a field has been set.

### SetIsAiEnabledNil

`func (o *ProjectForUpdateDto) SetIsAiEnabledNil(b bool)`

 SetIsAiEnabledNil sets the value for IsAiEnabled to be an explicit nil

### UnsetIsAiEnabled
`func (o *ProjectForUpdateDto) UnsetIsAiEnabled()`

UnsetIsAiEnabled ensures that no value is present for IsAiEnabled, not even an explicit nil
### GetIsAutoUpgrade

`func (o *ProjectForUpdateDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectForUpdateDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectForUpdateDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectForUpdateDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### SetIsAutoUpgradeNil

`func (o *ProjectForUpdateDto) SetIsAutoUpgradeNil(b bool)`

 SetIsAutoUpgradeNil sets the value for IsAutoUpgrade to be an explicit nil

### UnsetIsAutoUpgrade
`func (o *ProjectForUpdateDto) UnsetIsAutoUpgrade()`

UnsetIsAutoUpgrade ensures that no value is present for IsAutoUpgrade, not even an explicit nil
### GetAppEnabled

`func (o *ProjectForUpdateDto) GetAppEnabled() bool`

GetAppEnabled returns the AppEnabled field if non-nil, zero value otherwise.

### GetAppEnabledOk

`func (o *ProjectForUpdateDto) GetAppEnabledOk() (*bool, bool)`

GetAppEnabledOk returns a tuple with the AppEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnabled

`func (o *ProjectForUpdateDto) SetAppEnabled(v bool)`

SetAppEnabled sets AppEnabled field to given value.

### HasAppEnabled

`func (o *ProjectForUpdateDto) HasAppEnabled() bool`

HasAppEnabled returns a boolean if a field has been set.

### SetAppEnabledNil

`func (o *ProjectForUpdateDto) SetAppEnabledNil(b bool)`

 SetAppEnabledNil sets the value for AppEnabled to be an explicit nil

### UnsetAppEnabled
`func (o *ProjectForUpdateDto) UnsetAppEnabled()`

UnsetAppEnabled ensures that no value is present for AppEnabled, not even an explicit nil
### GetIsKubevapEnabled

`func (o *ProjectForUpdateDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *ProjectForUpdateDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *ProjectForUpdateDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *ProjectForUpdateDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.

### SetIsKubevapEnabledNil

`func (o *ProjectForUpdateDto) SetIsKubevapEnabledNil(b bool)`

 SetIsKubevapEnabledNil sets the value for IsKubevapEnabled to be an explicit nil

### UnsetIsKubevapEnabled
`func (o *ProjectForUpdateDto) UnsetIsKubevapEnabled()`

UnsetIsKubevapEnabled ensures that no value is present for IsKubevapEnabled, not even an explicit nil
### GetKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectForUpdateDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectForUpdateDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *ProjectForUpdateDto) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *ProjectForUpdateDto) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil
### GetFailureReason

`func (o *ProjectForUpdateDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProjectForUpdateDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProjectForUpdateDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProjectForUpdateDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ProjectForUpdateDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ProjectForUpdateDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


