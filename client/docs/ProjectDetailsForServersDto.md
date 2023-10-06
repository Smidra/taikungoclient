# ProjectDetailsForServersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsTotalCount** | Pointer to **int32** |  | [optional] 
**Worker** | Pointer to **int32** |  | [optional] 
**Bastion** | Pointer to **int32** |  | [optional] 
**ProjectStatus** | Pointer to [**ProjectStatus**](ProjectStatus.md) |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**AccessIp** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**MasterReady** | Pointer to **int32** |  | [optional] 
**CloudType** | Pointer to [**CloudType**](CloudType.md) |  | [optional] 
**CloudName** | Pointer to **NullableString** |  | [optional] 
**CloudId** | Pointer to **int32** |  | [optional] 
**QuotaId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**AiEnabled** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**IsOpaEnabled** | Pointer to **bool** |  | [optional] 
**HasKubeConfigFile** | Pointer to **bool** |  | [optional] 
**HasSelectedFlavors** | Pointer to **bool** |  | [optional] 
**IsMaintenanceModeEnabled** | Pointer to **bool** |  | [optional] 
**IsDeprecated** | Pointer to **bool** |  | [optional] 
**CpuLimit** | Pointer to **int64** |  | [optional] 
**RamLimit** | Pointer to **int64** |  | [optional] 
**DiskSizeLimit** | Pointer to **int64** |  | [optional] 
**UsedCpu** | Pointer to **int64** |  | [optional] 
**UsedRam** | Pointer to **int64** |  | [optional] 
**UsedDiskSize** | Pointer to **int64** |  | [optional] 
**VmCpuLimit** | Pointer to **int64** |  | [optional] 
**VmRamLimit** | Pointer to **int64** |  | [optional] 
**VmVolumeSizeLimit** | Pointer to **int64** |  | [optional] 
**VmUsedCpu** | Pointer to **int64** |  | [optional] 
**VmUsedRam** | Pointer to **int64** |  | [optional] 
**VmUsedVolumeSize** | Pointer to **int64** |  | [optional] 
**AccessProfileName** | Pointer to **NullableString** |  | [optional] 
**AccessProfileId** | Pointer to **NullableInt32** |  | [optional] 
**KubernetesProfileName** | Pointer to **NullableString** |  | [optional] 
**KubernetesProfileId** | Pointer to **NullableInt32** |  | [optional] 
**AlertingProfileName** | Pointer to **NullableString** |  | [optional] 
**ProjectHealth** | Pointer to [**ProjectHealth**](ProjectHealth.md) |  | [optional] 
**AlertingProfileId** | Pointer to **NullableInt32** |  | [optional] 
**S3CredentialId** | Pointer to **NullableInt32** |  | [optional] 
**AiCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**ExpiredAt** | Pointer to **NullableString** |  | [optional] 
**CertificationExpiredAt** | Pointer to **NullableString** |  | [optional] 
**OpaProfileId** | Pointer to **NullableInt32** |  | [optional] 
**OpaProfileName** | Pointer to **NullableString** |  | [optional] 
**AllowFullSpotKubernetes** | Pointer to **bool** |  | [optional] 
**AllowSpotWorkers** | Pointer to **bool** |  | [optional] 
**AllowSpotVMs** | Pointer to **bool** |  | [optional] 
**TotalHourlyCost** | Pointer to **float64** |  | [optional] 
**AutoscalingGroupName** | Pointer to **NullableString** |  | [optional] 
**MinSize** | Pointer to **NullableInt32** |  | [optional] 
**MaxSize** | Pointer to **NullableInt32** |  | [optional] 
**DiskSize** | Pointer to **NullableFloat64** |  | [optional] 
**Flavor** | Pointer to **NullableString** |  | [optional] 
**SpotEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsAutoscalingEnabled** | Pointer to **bool** |  | [optional] 
**HasNfsServer** | Pointer to **bool** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**Hypervisors** | Pointer to **[]string** |  | [optional] 
**ProxmoxStorage** | Pointer to [**ProxmoxStorage**](ProxmoxStorage.md) |  | [optional] 

## Methods

### NewProjectDetailsForServersDto

`func NewProjectDetailsForServersDto() *ProjectDetailsForServersDto`

NewProjectDetailsForServersDto instantiates a new ProjectDetailsForServersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsForServersDtoWithDefaults

`func NewProjectDetailsForServersDtoWithDefaults() *ProjectDetailsForServersDto`

NewProjectDetailsForServersDtoWithDefaults instantiates a new ProjectDetailsForServersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsTotalCount

`func (o *ProjectDetailsForServersDto) GetAlertsTotalCount() int32`

GetAlertsTotalCount returns the AlertsTotalCount field if non-nil, zero value otherwise.

### GetAlertsTotalCountOk

`func (o *ProjectDetailsForServersDto) GetAlertsTotalCountOk() (*int32, bool)`

GetAlertsTotalCountOk returns a tuple with the AlertsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsTotalCount

`func (o *ProjectDetailsForServersDto) SetAlertsTotalCount(v int32)`

SetAlertsTotalCount sets AlertsTotalCount field to given value.

### HasAlertsTotalCount

`func (o *ProjectDetailsForServersDto) HasAlertsTotalCount() bool`

HasAlertsTotalCount returns a boolean if a field has been set.

### GetWorker

`func (o *ProjectDetailsForServersDto) GetWorker() int32`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ProjectDetailsForServersDto) GetWorkerOk() (*int32, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ProjectDetailsForServersDto) SetWorker(v int32)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ProjectDetailsForServersDto) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetBastion

`func (o *ProjectDetailsForServersDto) GetBastion() int32`

GetBastion returns the Bastion field if non-nil, zero value otherwise.

### GetBastionOk

`func (o *ProjectDetailsForServersDto) GetBastionOk() (*int32, bool)`

GetBastionOk returns a tuple with the Bastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastion

`func (o *ProjectDetailsForServersDto) SetBastion(v int32)`

SetBastion sets Bastion field to given value.

### HasBastion

`func (o *ProjectDetailsForServersDto) HasBastion() bool`

HasBastion returns a boolean if a field has been set.

### GetProjectStatus

`func (o *ProjectDetailsForServersDto) GetProjectStatus() ProjectStatus`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *ProjectDetailsForServersDto) GetProjectStatusOk() (*ProjectStatus, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *ProjectDetailsForServersDto) SetProjectStatus(v ProjectStatus)`

SetProjectStatus sets ProjectStatus field to given value.

### HasProjectStatus

`func (o *ProjectDetailsForServersDto) HasProjectStatus() bool`

HasProjectStatus returns a boolean if a field has been set.

### GetProjectName

`func (o *ProjectDetailsForServersDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectDetailsForServersDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectDetailsForServersDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectDetailsForServersDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ProjectDetailsForServersDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectDetailsForServersDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetAccessIp

`func (o *ProjectDetailsForServersDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectDetailsForServersDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectDetailsForServersDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectDetailsForServersDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### SetAccessIpNil

`func (o *ProjectDetailsForServersDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *ProjectDetailsForServersDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetProjectId

`func (o *ProjectDetailsForServersDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectDetailsForServersDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectDetailsForServersDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectDetailsForServersDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetMasterReady

`func (o *ProjectDetailsForServersDto) GetMasterReady() int32`

GetMasterReady returns the MasterReady field if non-nil, zero value otherwise.

### GetMasterReadyOk

`func (o *ProjectDetailsForServersDto) GetMasterReadyOk() (*int32, bool)`

GetMasterReadyOk returns a tuple with the MasterReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterReady

`func (o *ProjectDetailsForServersDto) SetMasterReady(v int32)`

SetMasterReady sets MasterReady field to given value.

### HasMasterReady

`func (o *ProjectDetailsForServersDto) HasMasterReady() bool`

HasMasterReady returns a boolean if a field has been set.

### GetCloudType

`func (o *ProjectDetailsForServersDto) GetCloudType() CloudType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectDetailsForServersDto) GetCloudTypeOk() (*CloudType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectDetailsForServersDto) SetCloudType(v CloudType)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectDetailsForServersDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudName

`func (o *ProjectDetailsForServersDto) GetCloudName() string`

GetCloudName returns the CloudName field if non-nil, zero value otherwise.

### GetCloudNameOk

`func (o *ProjectDetailsForServersDto) GetCloudNameOk() (*string, bool)`

GetCloudNameOk returns a tuple with the CloudName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudName

`func (o *ProjectDetailsForServersDto) SetCloudName(v string)`

SetCloudName sets CloudName field to given value.

### HasCloudName

`func (o *ProjectDetailsForServersDto) HasCloudName() bool`

HasCloudName returns a boolean if a field has been set.

### SetCloudNameNil

`func (o *ProjectDetailsForServersDto) SetCloudNameNil(b bool)`

 SetCloudNameNil sets the value for CloudName to be an explicit nil

### UnsetCloudName
`func (o *ProjectDetailsForServersDto) UnsetCloudName()`

UnsetCloudName ensures that no value is present for CloudName, not even an explicit nil
### GetCloudId

`func (o *ProjectDetailsForServersDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ProjectDetailsForServersDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ProjectDetailsForServersDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *ProjectDetailsForServersDto) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetQuotaId

`func (o *ProjectDetailsForServersDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectDetailsForServersDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectDetailsForServersDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *ProjectDetailsForServersDto) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ProjectDetailsForServersDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectDetailsForServersDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectDetailsForServersDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectDetailsForServersDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ProjectDetailsForServersDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectDetailsForServersDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ProjectDetailsForServersDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectDetailsForServersDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectDetailsForServersDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectDetailsForServersDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetKubernetesCurrentVersion

`func (o *ProjectDetailsForServersDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectDetailsForServersDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectDetailsForServersDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectDetailsForServersDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *ProjectDetailsForServersDto) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *ProjectDetailsForServersDto) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil
### GetIsBackupEnabled

`func (o *ProjectDetailsForServersDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectDetailsForServersDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectDetailsForServersDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetAiEnabled

`func (o *ProjectDetailsForServersDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ProjectDetailsForServersDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ProjectDetailsForServersDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *ProjectDetailsForServersDto) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectDetailsForServersDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectDetailsForServersDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectDetailsForServersDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectDetailsForServersDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *ProjectDetailsForServersDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectDetailsForServersDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectDetailsForServersDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectDetailsForServersDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectDetailsForServersDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectDetailsForServersDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectDetailsForServersDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetIsOpaEnabled

`func (o *ProjectDetailsForServersDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectDetailsForServersDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.

### HasIsOpaEnabled

`func (o *ProjectDetailsForServersDto) HasIsOpaEnabled() bool`

HasIsOpaEnabled returns a boolean if a field has been set.

### GetHasKubeConfigFile

`func (o *ProjectDetailsForServersDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectDetailsForServersDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectDetailsForServersDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.

### HasHasKubeConfigFile

`func (o *ProjectDetailsForServersDto) HasHasKubeConfigFile() bool`

HasHasKubeConfigFile returns a boolean if a field has been set.

### GetHasSelectedFlavors

`func (o *ProjectDetailsForServersDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectDetailsForServersDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectDetailsForServersDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.

### HasHasSelectedFlavors

`func (o *ProjectDetailsForServersDto) HasHasSelectedFlavors() bool`

HasHasSelectedFlavors returns a boolean if a field has been set.

### GetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.

### HasIsMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) HasIsMaintenanceModeEnabled() bool`

HasIsMaintenanceModeEnabled returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *ProjectDetailsForServersDto) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ProjectDetailsForServersDto) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ProjectDetailsForServersDto) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *ProjectDetailsForServersDto) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetCpuLimit

`func (o *ProjectDetailsForServersDto) GetCpuLimit() int64`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ProjectDetailsForServersDto) GetCpuLimitOk() (*int64, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ProjectDetailsForServersDto) SetCpuLimit(v int64)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *ProjectDetailsForServersDto) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetRamLimit

`func (o *ProjectDetailsForServersDto) GetRamLimit() int64`

GetRamLimit returns the RamLimit field if non-nil, zero value otherwise.

### GetRamLimitOk

`func (o *ProjectDetailsForServersDto) GetRamLimitOk() (*int64, bool)`

GetRamLimitOk returns a tuple with the RamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamLimit

`func (o *ProjectDetailsForServersDto) SetRamLimit(v int64)`

SetRamLimit sets RamLimit field to given value.

### HasRamLimit

`func (o *ProjectDetailsForServersDto) HasRamLimit() bool`

HasRamLimit returns a boolean if a field has been set.

### GetDiskSizeLimit

`func (o *ProjectDetailsForServersDto) GetDiskSizeLimit() int64`

GetDiskSizeLimit returns the DiskSizeLimit field if non-nil, zero value otherwise.

### GetDiskSizeLimitOk

`func (o *ProjectDetailsForServersDto) GetDiskSizeLimitOk() (*int64, bool)`

GetDiskSizeLimitOk returns a tuple with the DiskSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeLimit

`func (o *ProjectDetailsForServersDto) SetDiskSizeLimit(v int64)`

SetDiskSizeLimit sets DiskSizeLimit field to given value.

### HasDiskSizeLimit

`func (o *ProjectDetailsForServersDto) HasDiskSizeLimit() bool`

HasDiskSizeLimit returns a boolean if a field has been set.

### GetUsedCpu

`func (o *ProjectDetailsForServersDto) GetUsedCpu() int64`

GetUsedCpu returns the UsedCpu field if non-nil, zero value otherwise.

### GetUsedCpuOk

`func (o *ProjectDetailsForServersDto) GetUsedCpuOk() (*int64, bool)`

GetUsedCpuOk returns a tuple with the UsedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCpu

`func (o *ProjectDetailsForServersDto) SetUsedCpu(v int64)`

SetUsedCpu sets UsedCpu field to given value.

### HasUsedCpu

`func (o *ProjectDetailsForServersDto) HasUsedCpu() bool`

HasUsedCpu returns a boolean if a field has been set.

### GetUsedRam

`func (o *ProjectDetailsForServersDto) GetUsedRam() int64`

GetUsedRam returns the UsedRam field if non-nil, zero value otherwise.

### GetUsedRamOk

`func (o *ProjectDetailsForServersDto) GetUsedRamOk() (*int64, bool)`

GetUsedRamOk returns a tuple with the UsedRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRam

`func (o *ProjectDetailsForServersDto) SetUsedRam(v int64)`

SetUsedRam sets UsedRam field to given value.

### HasUsedRam

`func (o *ProjectDetailsForServersDto) HasUsedRam() bool`

HasUsedRam returns a boolean if a field has been set.

### GetUsedDiskSize

`func (o *ProjectDetailsForServersDto) GetUsedDiskSize() int64`

GetUsedDiskSize returns the UsedDiskSize field if non-nil, zero value otherwise.

### GetUsedDiskSizeOk

`func (o *ProjectDetailsForServersDto) GetUsedDiskSizeOk() (*int64, bool)`

GetUsedDiskSizeOk returns a tuple with the UsedDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDiskSize

`func (o *ProjectDetailsForServersDto) SetUsedDiskSize(v int64)`

SetUsedDiskSize sets UsedDiskSize field to given value.

### HasUsedDiskSize

`func (o *ProjectDetailsForServersDto) HasUsedDiskSize() bool`

HasUsedDiskSize returns a boolean if a field has been set.

### GetVmCpuLimit

`func (o *ProjectDetailsForServersDto) GetVmCpuLimit() int64`

GetVmCpuLimit returns the VmCpuLimit field if non-nil, zero value otherwise.

### GetVmCpuLimitOk

`func (o *ProjectDetailsForServersDto) GetVmCpuLimitOk() (*int64, bool)`

GetVmCpuLimitOk returns a tuple with the VmCpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuLimit

`func (o *ProjectDetailsForServersDto) SetVmCpuLimit(v int64)`

SetVmCpuLimit sets VmCpuLimit field to given value.

### HasVmCpuLimit

`func (o *ProjectDetailsForServersDto) HasVmCpuLimit() bool`

HasVmCpuLimit returns a boolean if a field has been set.

### GetVmRamLimit

`func (o *ProjectDetailsForServersDto) GetVmRamLimit() int64`

GetVmRamLimit returns the VmRamLimit field if non-nil, zero value otherwise.

### GetVmRamLimitOk

`func (o *ProjectDetailsForServersDto) GetVmRamLimitOk() (*int64, bool)`

GetVmRamLimitOk returns a tuple with the VmRamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRamLimit

`func (o *ProjectDetailsForServersDto) SetVmRamLimit(v int64)`

SetVmRamLimit sets VmRamLimit field to given value.

### HasVmRamLimit

`func (o *ProjectDetailsForServersDto) HasVmRamLimit() bool`

HasVmRamLimit returns a boolean if a field has been set.

### GetVmVolumeSizeLimit

`func (o *ProjectDetailsForServersDto) GetVmVolumeSizeLimit() int64`

GetVmVolumeSizeLimit returns the VmVolumeSizeLimit field if non-nil, zero value otherwise.

### GetVmVolumeSizeLimitOk

`func (o *ProjectDetailsForServersDto) GetVmVolumeSizeLimitOk() (*int64, bool)`

GetVmVolumeSizeLimitOk returns a tuple with the VmVolumeSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVolumeSizeLimit

`func (o *ProjectDetailsForServersDto) SetVmVolumeSizeLimit(v int64)`

SetVmVolumeSizeLimit sets VmVolumeSizeLimit field to given value.

### HasVmVolumeSizeLimit

`func (o *ProjectDetailsForServersDto) HasVmVolumeSizeLimit() bool`

HasVmVolumeSizeLimit returns a boolean if a field has been set.

### GetVmUsedCpu

`func (o *ProjectDetailsForServersDto) GetVmUsedCpu() int64`

GetVmUsedCpu returns the VmUsedCpu field if non-nil, zero value otherwise.

### GetVmUsedCpuOk

`func (o *ProjectDetailsForServersDto) GetVmUsedCpuOk() (*int64, bool)`

GetVmUsedCpuOk returns a tuple with the VmUsedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedCpu

`func (o *ProjectDetailsForServersDto) SetVmUsedCpu(v int64)`

SetVmUsedCpu sets VmUsedCpu field to given value.

### HasVmUsedCpu

`func (o *ProjectDetailsForServersDto) HasVmUsedCpu() bool`

HasVmUsedCpu returns a boolean if a field has been set.

### GetVmUsedRam

`func (o *ProjectDetailsForServersDto) GetVmUsedRam() int64`

GetVmUsedRam returns the VmUsedRam field if non-nil, zero value otherwise.

### GetVmUsedRamOk

`func (o *ProjectDetailsForServersDto) GetVmUsedRamOk() (*int64, bool)`

GetVmUsedRamOk returns a tuple with the VmUsedRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedRam

`func (o *ProjectDetailsForServersDto) SetVmUsedRam(v int64)`

SetVmUsedRam sets VmUsedRam field to given value.

### HasVmUsedRam

`func (o *ProjectDetailsForServersDto) HasVmUsedRam() bool`

HasVmUsedRam returns a boolean if a field has been set.

### GetVmUsedVolumeSize

`func (o *ProjectDetailsForServersDto) GetVmUsedVolumeSize() int64`

GetVmUsedVolumeSize returns the VmUsedVolumeSize field if non-nil, zero value otherwise.

### GetVmUsedVolumeSizeOk

`func (o *ProjectDetailsForServersDto) GetVmUsedVolumeSizeOk() (*int64, bool)`

GetVmUsedVolumeSizeOk returns a tuple with the VmUsedVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedVolumeSize

`func (o *ProjectDetailsForServersDto) SetVmUsedVolumeSize(v int64)`

SetVmUsedVolumeSize sets VmUsedVolumeSize field to given value.

### HasVmUsedVolumeSize

`func (o *ProjectDetailsForServersDto) HasVmUsedVolumeSize() bool`

HasVmUsedVolumeSize returns a boolean if a field has been set.

### GetAccessProfileName

`func (o *ProjectDetailsForServersDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *ProjectDetailsForServersDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *ProjectDetailsForServersDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.

### HasAccessProfileName

`func (o *ProjectDetailsForServersDto) HasAccessProfileName() bool`

HasAccessProfileName returns a boolean if a field has been set.

### SetAccessProfileNameNil

`func (o *ProjectDetailsForServersDto) SetAccessProfileNameNil(b bool)`

 SetAccessProfileNameNil sets the value for AccessProfileName to be an explicit nil

### UnsetAccessProfileName
`func (o *ProjectDetailsForServersDto) UnsetAccessProfileName()`

UnsetAccessProfileName ensures that no value is present for AccessProfileName, not even an explicit nil
### GetAccessProfileId

`func (o *ProjectDetailsForServersDto) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *ProjectDetailsForServersDto) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *ProjectDetailsForServersDto) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.

### HasAccessProfileId

`func (o *ProjectDetailsForServersDto) HasAccessProfileId() bool`

HasAccessProfileId returns a boolean if a field has been set.

### SetAccessProfileIdNil

`func (o *ProjectDetailsForServersDto) SetAccessProfileIdNil(b bool)`

 SetAccessProfileIdNil sets the value for AccessProfileId to be an explicit nil

### UnsetAccessProfileId
`func (o *ProjectDetailsForServersDto) UnsetAccessProfileId()`

UnsetAccessProfileId ensures that no value is present for AccessProfileId, not even an explicit nil
### GetKubernetesProfileName

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileName() string`

GetKubernetesProfileName returns the KubernetesProfileName field if non-nil, zero value otherwise.

### GetKubernetesProfileNameOk

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileNameOk() (*string, bool)`

GetKubernetesProfileNameOk returns a tuple with the KubernetesProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfileName

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileName(v string)`

SetKubernetesProfileName sets KubernetesProfileName field to given value.

### HasKubernetesProfileName

`func (o *ProjectDetailsForServersDto) HasKubernetesProfileName() bool`

HasKubernetesProfileName returns a boolean if a field has been set.

### SetKubernetesProfileNameNil

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileNameNil(b bool)`

 SetKubernetesProfileNameNil sets the value for KubernetesProfileName to be an explicit nil

### UnsetKubernetesProfileName
`func (o *ProjectDetailsForServersDto) UnsetKubernetesProfileName()`

UnsetKubernetesProfileName ensures that no value is present for KubernetesProfileName, not even an explicit nil
### GetKubernetesProfileId

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileId() int32`

GetKubernetesProfileId returns the KubernetesProfileId field if non-nil, zero value otherwise.

### GetKubernetesProfileIdOk

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileIdOk() (*int32, bool)`

GetKubernetesProfileIdOk returns a tuple with the KubernetesProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfileId

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileId(v int32)`

SetKubernetesProfileId sets KubernetesProfileId field to given value.

### HasKubernetesProfileId

`func (o *ProjectDetailsForServersDto) HasKubernetesProfileId() bool`

HasKubernetesProfileId returns a boolean if a field has been set.

### SetKubernetesProfileIdNil

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileIdNil(b bool)`

 SetKubernetesProfileIdNil sets the value for KubernetesProfileId to be an explicit nil

### UnsetKubernetesProfileId
`func (o *ProjectDetailsForServersDto) UnsetKubernetesProfileId()`

UnsetKubernetesProfileId ensures that no value is present for KubernetesProfileId, not even an explicit nil
### GetAlertingProfileName

`func (o *ProjectDetailsForServersDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *ProjectDetailsForServersDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *ProjectDetailsForServersDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.

### HasAlertingProfileName

`func (o *ProjectDetailsForServersDto) HasAlertingProfileName() bool`

HasAlertingProfileName returns a boolean if a field has been set.

### SetAlertingProfileNameNil

`func (o *ProjectDetailsForServersDto) SetAlertingProfileNameNil(b bool)`

 SetAlertingProfileNameNil sets the value for AlertingProfileName to be an explicit nil

### UnsetAlertingProfileName
`func (o *ProjectDetailsForServersDto) UnsetAlertingProfileName()`

UnsetAlertingProfileName ensures that no value is present for AlertingProfileName, not even an explicit nil
### GetProjectHealth

`func (o *ProjectDetailsForServersDto) GetProjectHealth() ProjectHealth`

GetProjectHealth returns the ProjectHealth field if non-nil, zero value otherwise.

### GetProjectHealthOk

`func (o *ProjectDetailsForServersDto) GetProjectHealthOk() (*ProjectHealth, bool)`

GetProjectHealthOk returns a tuple with the ProjectHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectHealth

`func (o *ProjectDetailsForServersDto) SetProjectHealth(v ProjectHealth)`

SetProjectHealth sets ProjectHealth field to given value.

### HasProjectHealth

`func (o *ProjectDetailsForServersDto) HasProjectHealth() bool`

HasProjectHealth returns a boolean if a field has been set.

### GetAlertingProfileId

`func (o *ProjectDetailsForServersDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *ProjectDetailsForServersDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *ProjectDetailsForServersDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.

### HasAlertingProfileId

`func (o *ProjectDetailsForServersDto) HasAlertingProfileId() bool`

HasAlertingProfileId returns a boolean if a field has been set.

### SetAlertingProfileIdNil

`func (o *ProjectDetailsForServersDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *ProjectDetailsForServersDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetS3CredentialId

`func (o *ProjectDetailsForServersDto) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *ProjectDetailsForServersDto) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *ProjectDetailsForServersDto) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.

### HasS3CredentialId

`func (o *ProjectDetailsForServersDto) HasS3CredentialId() bool`

HasS3CredentialId returns a boolean if a field has been set.

### SetS3CredentialIdNil

`func (o *ProjectDetailsForServersDto) SetS3CredentialIdNil(b bool)`

 SetS3CredentialIdNil sets the value for S3CredentialId to be an explicit nil

### UnsetS3CredentialId
`func (o *ProjectDetailsForServersDto) UnsetS3CredentialId()`

UnsetS3CredentialId ensures that no value is present for S3CredentialId, not even an explicit nil
### GetAiCredentialId

`func (o *ProjectDetailsForServersDto) GetAiCredentialId() int32`

GetAiCredentialId returns the AiCredentialId field if non-nil, zero value otherwise.

### GetAiCredentialIdOk

`func (o *ProjectDetailsForServersDto) GetAiCredentialIdOk() (*int32, bool)`

GetAiCredentialIdOk returns a tuple with the AiCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiCredentialId

`func (o *ProjectDetailsForServersDto) SetAiCredentialId(v int32)`

SetAiCredentialId sets AiCredentialId field to given value.

### HasAiCredentialId

`func (o *ProjectDetailsForServersDto) HasAiCredentialId() bool`

HasAiCredentialId returns a boolean if a field has been set.

### SetAiCredentialIdNil

`func (o *ProjectDetailsForServersDto) SetAiCredentialIdNil(b bool)`

 SetAiCredentialIdNil sets the value for AiCredentialId to be an explicit nil

### UnsetAiCredentialId
`func (o *ProjectDetailsForServersDto) UnsetAiCredentialId()`

UnsetAiCredentialId ensures that no value is present for AiCredentialId, not even an explicit nil
### GetExpiredAt

`func (o *ProjectDetailsForServersDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectDetailsForServersDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectDetailsForServersDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *ProjectDetailsForServersDto) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *ProjectDetailsForServersDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ProjectDetailsForServersDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetCertificationExpiredAt

`func (o *ProjectDetailsForServersDto) GetCertificationExpiredAt() string`

GetCertificationExpiredAt returns the CertificationExpiredAt field if non-nil, zero value otherwise.

### GetCertificationExpiredAtOk

`func (o *ProjectDetailsForServersDto) GetCertificationExpiredAtOk() (*string, bool)`

GetCertificationExpiredAtOk returns a tuple with the CertificationExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationExpiredAt

`func (o *ProjectDetailsForServersDto) SetCertificationExpiredAt(v string)`

SetCertificationExpiredAt sets CertificationExpiredAt field to given value.

### HasCertificationExpiredAt

`func (o *ProjectDetailsForServersDto) HasCertificationExpiredAt() bool`

HasCertificationExpiredAt returns a boolean if a field has been set.

### SetCertificationExpiredAtNil

`func (o *ProjectDetailsForServersDto) SetCertificationExpiredAtNil(b bool)`

 SetCertificationExpiredAtNil sets the value for CertificationExpiredAt to be an explicit nil

### UnsetCertificationExpiredAt
`func (o *ProjectDetailsForServersDto) UnsetCertificationExpiredAt()`

UnsetCertificationExpiredAt ensures that no value is present for CertificationExpiredAt, not even an explicit nil
### GetOpaProfileId

`func (o *ProjectDetailsForServersDto) GetOpaProfileId() int32`

GetOpaProfileId returns the OpaProfileId field if non-nil, zero value otherwise.

### GetOpaProfileIdOk

`func (o *ProjectDetailsForServersDto) GetOpaProfileIdOk() (*int32, bool)`

GetOpaProfileIdOk returns a tuple with the OpaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileId

`func (o *ProjectDetailsForServersDto) SetOpaProfileId(v int32)`

SetOpaProfileId sets OpaProfileId field to given value.

### HasOpaProfileId

`func (o *ProjectDetailsForServersDto) HasOpaProfileId() bool`

HasOpaProfileId returns a boolean if a field has been set.

### SetOpaProfileIdNil

`func (o *ProjectDetailsForServersDto) SetOpaProfileIdNil(b bool)`

 SetOpaProfileIdNil sets the value for OpaProfileId to be an explicit nil

### UnsetOpaProfileId
`func (o *ProjectDetailsForServersDto) UnsetOpaProfileId()`

UnsetOpaProfileId ensures that no value is present for OpaProfileId, not even an explicit nil
### GetOpaProfileName

`func (o *ProjectDetailsForServersDto) GetOpaProfileName() string`

GetOpaProfileName returns the OpaProfileName field if non-nil, zero value otherwise.

### GetOpaProfileNameOk

`func (o *ProjectDetailsForServersDto) GetOpaProfileNameOk() (*string, bool)`

GetOpaProfileNameOk returns a tuple with the OpaProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileName

`func (o *ProjectDetailsForServersDto) SetOpaProfileName(v string)`

SetOpaProfileName sets OpaProfileName field to given value.

### HasOpaProfileName

`func (o *ProjectDetailsForServersDto) HasOpaProfileName() bool`

HasOpaProfileName returns a boolean if a field has been set.

### SetOpaProfileNameNil

`func (o *ProjectDetailsForServersDto) SetOpaProfileNameNil(b bool)`

 SetOpaProfileNameNil sets the value for OpaProfileName to be an explicit nil

### UnsetOpaProfileName
`func (o *ProjectDetailsForServersDto) UnsetOpaProfileName()`

UnsetOpaProfileName ensures that no value is present for OpaProfileName, not even an explicit nil
### GetAllowFullSpotKubernetes

`func (o *ProjectDetailsForServersDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectDetailsForServersDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectDetailsForServersDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.

### HasAllowFullSpotKubernetes

`func (o *ProjectDetailsForServersDto) HasAllowFullSpotKubernetes() bool`

HasAllowFullSpotKubernetes returns a boolean if a field has been set.

### GetAllowSpotWorkers

`func (o *ProjectDetailsForServersDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectDetailsForServersDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectDetailsForServersDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.

### HasAllowSpotWorkers

`func (o *ProjectDetailsForServersDto) HasAllowSpotWorkers() bool`

HasAllowSpotWorkers returns a boolean if a field has been set.

### GetAllowSpotVMs

`func (o *ProjectDetailsForServersDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectDetailsForServersDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectDetailsForServersDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.

### HasAllowSpotVMs

`func (o *ProjectDetailsForServersDto) HasAllowSpotVMs() bool`

HasAllowSpotVMs returns a boolean if a field has been set.

### GetTotalHourlyCost

`func (o *ProjectDetailsForServersDto) GetTotalHourlyCost() float64`

GetTotalHourlyCost returns the TotalHourlyCost field if non-nil, zero value otherwise.

### GetTotalHourlyCostOk

`func (o *ProjectDetailsForServersDto) GetTotalHourlyCostOk() (*float64, bool)`

GetTotalHourlyCostOk returns a tuple with the TotalHourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHourlyCost

`func (o *ProjectDetailsForServersDto) SetTotalHourlyCost(v float64)`

SetTotalHourlyCost sets TotalHourlyCost field to given value.

### HasTotalHourlyCost

`func (o *ProjectDetailsForServersDto) HasTotalHourlyCost() bool`

HasTotalHourlyCost returns a boolean if a field has been set.

### GetAutoscalingGroupName

`func (o *ProjectDetailsForServersDto) GetAutoscalingGroupName() string`

GetAutoscalingGroupName returns the AutoscalingGroupName field if non-nil, zero value otherwise.

### GetAutoscalingGroupNameOk

`func (o *ProjectDetailsForServersDto) GetAutoscalingGroupNameOk() (*string, bool)`

GetAutoscalingGroupNameOk returns a tuple with the AutoscalingGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroupName

`func (o *ProjectDetailsForServersDto) SetAutoscalingGroupName(v string)`

SetAutoscalingGroupName sets AutoscalingGroupName field to given value.

### HasAutoscalingGroupName

`func (o *ProjectDetailsForServersDto) HasAutoscalingGroupName() bool`

HasAutoscalingGroupName returns a boolean if a field has been set.

### SetAutoscalingGroupNameNil

`func (o *ProjectDetailsForServersDto) SetAutoscalingGroupNameNil(b bool)`

 SetAutoscalingGroupNameNil sets the value for AutoscalingGroupName to be an explicit nil

### UnsetAutoscalingGroupName
`func (o *ProjectDetailsForServersDto) UnsetAutoscalingGroupName()`

UnsetAutoscalingGroupName ensures that no value is present for AutoscalingGroupName, not even an explicit nil
### GetMinSize

`func (o *ProjectDetailsForServersDto) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *ProjectDetailsForServersDto) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *ProjectDetailsForServersDto) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *ProjectDetailsForServersDto) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### SetMinSizeNil

`func (o *ProjectDetailsForServersDto) SetMinSizeNil(b bool)`

 SetMinSizeNil sets the value for MinSize to be an explicit nil

### UnsetMinSize
`func (o *ProjectDetailsForServersDto) UnsetMinSize()`

UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
### GetMaxSize

`func (o *ProjectDetailsForServersDto) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *ProjectDetailsForServersDto) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *ProjectDetailsForServersDto) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *ProjectDetailsForServersDto) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *ProjectDetailsForServersDto) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *ProjectDetailsForServersDto) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetDiskSize

`func (o *ProjectDetailsForServersDto) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ProjectDetailsForServersDto) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ProjectDetailsForServersDto) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ProjectDetailsForServersDto) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### SetDiskSizeNil

`func (o *ProjectDetailsForServersDto) SetDiskSizeNil(b bool)`

 SetDiskSizeNil sets the value for DiskSize to be an explicit nil

### UnsetDiskSize
`func (o *ProjectDetailsForServersDto) UnsetDiskSize()`

UnsetDiskSize ensures that no value is present for DiskSize, not even an explicit nil
### GetFlavor

`func (o *ProjectDetailsForServersDto) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ProjectDetailsForServersDto) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ProjectDetailsForServersDto) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *ProjectDetailsForServersDto) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *ProjectDetailsForServersDto) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *ProjectDetailsForServersDto) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetSpotEnabled

`func (o *ProjectDetailsForServersDto) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *ProjectDetailsForServersDto) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *ProjectDetailsForServersDto) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.

### HasSpotEnabled

`func (o *ProjectDetailsForServersDto) HasSpotEnabled() bool`

HasSpotEnabled returns a boolean if a field has been set.

### SetSpotEnabledNil

`func (o *ProjectDetailsForServersDto) SetSpotEnabledNil(b bool)`

 SetSpotEnabledNil sets the value for SpotEnabled to be an explicit nil

### UnsetSpotEnabled
`func (o *ProjectDetailsForServersDto) UnsetSpotEnabled()`

UnsetSpotEnabled ensures that no value is present for SpotEnabled, not even an explicit nil
### GetIsAutoscalingEnabled

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectDetailsForServersDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.

### HasIsAutoscalingEnabled

`func (o *ProjectDetailsForServersDto) HasIsAutoscalingEnabled() bool`

HasIsAutoscalingEnabled returns a boolean if a field has been set.

### GetHasNfsServer

`func (o *ProjectDetailsForServersDto) GetHasNfsServer() bool`

GetHasNfsServer returns the HasNfsServer field if non-nil, zero value otherwise.

### GetHasNfsServerOk

`func (o *ProjectDetailsForServersDto) GetHasNfsServerOk() (*bool, bool)`

GetHasNfsServerOk returns a tuple with the HasNfsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNfsServer

`func (o *ProjectDetailsForServersDto) SetHasNfsServer(v bool)`

SetHasNfsServer sets HasNfsServer field to given value.

### HasHasNfsServer

`func (o *ProjectDetailsForServersDto) HasHasNfsServer() bool`

HasHasNfsServer returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *ProjectDetailsForServersDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectDetailsForServersDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectDetailsForServersDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *ProjectDetailsForServersDto) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *ProjectDetailsForServersDto) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *ProjectDetailsForServersDto) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetHypervisors

`func (o *ProjectDetailsForServersDto) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *ProjectDetailsForServersDto) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *ProjectDetailsForServersDto) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *ProjectDetailsForServersDto) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.

### SetHypervisorsNil

`func (o *ProjectDetailsForServersDto) SetHypervisorsNil(b bool)`

 SetHypervisorsNil sets the value for Hypervisors to be an explicit nil

### UnsetHypervisors
`func (o *ProjectDetailsForServersDto) UnsetHypervisors()`

UnsetHypervisors ensures that no value is present for Hypervisors, not even an explicit nil
### GetProxmoxStorage

`func (o *ProjectDetailsForServersDto) GetProxmoxStorage() ProxmoxStorage`

GetProxmoxStorage returns the ProxmoxStorage field if non-nil, zero value otherwise.

### GetProxmoxStorageOk

`func (o *ProjectDetailsForServersDto) GetProxmoxStorageOk() (*ProxmoxStorage, bool)`

GetProxmoxStorageOk returns a tuple with the ProxmoxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxStorage

`func (o *ProjectDetailsForServersDto) SetProxmoxStorage(v ProxmoxStorage)`

SetProxmoxStorage sets ProxmoxStorage field to given value.

### HasProxmoxStorage

`func (o *ProjectDetailsForServersDto) HasProxmoxStorage() bool`

HasProxmoxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


