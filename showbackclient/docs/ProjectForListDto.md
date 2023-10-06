# ProjectForListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**IsOpaEnabled** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**AiEnabled** | Pointer to **bool** |  | [optional] 
**S3BucketName** | Pointer to **NullableString** |  | [optional] 
**HasKubeConfigFile** | Pointer to **bool** |  | [optional] 
**HasSelectedFlavors** | Pointer to **bool** |  | [optional] 
**QuotaId** | Pointer to **int32** |  | [optional] 
**CloudCredentialName** | Pointer to **NullableString** |  | [optional] 
**CloudCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**PartnerId** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Health** | Pointer to **NullableString** |  | [optional] 
**AccessIp** | Pointer to **NullableString** |  | [optional] 
**CloudType** | Pointer to **NullableString** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**KubesprayTargetVersion** | Pointer to **NullableString** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**KubernetesTargetVersion** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**MonitoringCredential** | Pointer to [**MonitoringCredentialsListDto**](MonitoringCredentialsListDto.md) |  | [optional] 
**IsAutoscalingEnabled** | Pointer to **bool** |  | [optional] 
**Flavors** | Pointer to **[]string** |  | [optional] 
**AccessProfile** | Pointer to [**AccessProfilesForProjectListDto**](AccessProfilesForProjectListDto.md) |  | [optional] 
**KubernetesProfiles** | Pointer to [**KubernetesProfilesLisForPollerDto**](KubernetesProfilesLisForPollerDto.md) |  | [optional] 
**OpaProfile** | Pointer to [**OpaProfileListDto**](OpaProfileListDto.md) |  | [optional] 
**KubernetesAlerts** | Pointer to [**[]KubernetesAlertDto**](KubernetesAlertDto.md) |  | [optional] 
**IsDeleteCluster** | Pointer to **bool** |  | [optional] 
**TaikunPrivateSSHKey** | Pointer to **NullableString** |  | [optional] 
**TaikunPublicSSHKey** | Pointer to **NullableString** |  | [optional] 
**GoogleProjectId** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**NetMask** | Pointer to **NullableInt32** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**IsKubevapEnabled** | Pointer to **bool** |  | [optional] 
**TanzuReleaseVersion** | Pointer to **NullableString** |  | [optional] 
**NfsDiskSize** | Pointer to **int32** |  | [optional] 
**KubevapEnabeledKubernetesVersions** | Pointer to **[]string** |  | [optional] 
**AwsProjectAZSubnets** | Pointer to [**[]AwsProjectAZSubnetDto**](AwsProjectAZSubnetDto.md) |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**WorkersCount** | Pointer to **int32** |  | [optional] 
**TaikunLB** | Pointer to [**TaikunLbDto**](TaikunLbDto.md) |  | [optional] 
**S3Credential** | Pointer to [**S3CredentialForProjectDto**](S3CredentialForProjectDto.md) |  | [optional] 
**ProjectRevision** | Pointer to [**ProjectRevisionDto**](ProjectRevisionDto.md) |  | [optional] 
**ProjectActionDto** | Pointer to [**ProjectActionDto**](ProjectActionDto.md) |  | [optional] 
**AiCredential** | Pointer to [**AiCredentialDto**](AiCredentialDto.md) |  | [optional] 

## Methods

### NewProjectForListDto

`func NewProjectForListDto() *ProjectForListDto`

NewProjectForListDto instantiates a new ProjectForListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectForListDtoWithDefaults

`func NewProjectForListDtoWithDefaults() *ProjectForListDto`

NewProjectForListDtoWithDefaults instantiates a new ProjectForListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectForListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectForListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectForListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectForListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectForListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectForListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectForListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectForListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectForListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectForListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageName

`func (o *ProjectForListDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProjectForListDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProjectForListDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProjectForListDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *ProjectForListDto) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *ProjectForListDto) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetToken

`func (o *ProjectForListDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProjectForListDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProjectForListDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProjectForListDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ProjectForListDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ProjectForListDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetIsKubernetes

`func (o *ProjectForListDto) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectForListDto) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectForListDto) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *ProjectForListDto) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectForListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectForListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectForListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectForListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ProjectForListDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectForListDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectForListDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *ProjectForListDto) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *ProjectForListDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectForListDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectForListDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *ProjectForListDto) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetIsOpaEnabled

`func (o *ProjectForListDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectForListDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectForListDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.

### HasIsOpaEnabled

`func (o *ProjectForListDto) HasIsOpaEnabled() bool`

HasIsOpaEnabled returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *ProjectForListDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectForListDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectForListDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *ProjectForListDto) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetAiEnabled

`func (o *ProjectForListDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ProjectForListDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ProjectForListDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *ProjectForListDto) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetS3BucketName

`func (o *ProjectForListDto) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *ProjectForListDto) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *ProjectForListDto) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.

### HasS3BucketName

`func (o *ProjectForListDto) HasS3BucketName() bool`

HasS3BucketName returns a boolean if a field has been set.

### SetS3BucketNameNil

`func (o *ProjectForListDto) SetS3BucketNameNil(b bool)`

 SetS3BucketNameNil sets the value for S3BucketName to be an explicit nil

### UnsetS3BucketName
`func (o *ProjectForListDto) UnsetS3BucketName()`

UnsetS3BucketName ensures that no value is present for S3BucketName, not even an explicit nil
### GetHasKubeConfigFile

`func (o *ProjectForListDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectForListDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectForListDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.

### HasHasKubeConfigFile

`func (o *ProjectForListDto) HasHasKubeConfigFile() bool`

HasHasKubeConfigFile returns a boolean if a field has been set.

### GetHasSelectedFlavors

`func (o *ProjectForListDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectForListDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectForListDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.

### HasHasSelectedFlavors

`func (o *ProjectForListDto) HasHasSelectedFlavors() bool`

HasHasSelectedFlavors returns a boolean if a field has been set.

### GetQuotaId

`func (o *ProjectForListDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectForListDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectForListDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *ProjectForListDto) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *ProjectForListDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ProjectForListDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ProjectForListDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ProjectForListDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### SetCloudCredentialNameNil

`func (o *ProjectForListDto) SetCloudCredentialNameNil(b bool)`

 SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil

### UnsetCloudCredentialName
`func (o *ProjectForListDto) UnsetCloudCredentialName()`

UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
### GetCloudCredentialId

`func (o *ProjectForListDto) GetCloudCredentialId() int32`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *ProjectForListDto) GetCloudCredentialIdOk() (*int32, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *ProjectForListDto) SetCloudCredentialId(v int32)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *ProjectForListDto) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### SetCloudCredentialIdNil

`func (o *ProjectForListDto) SetCloudCredentialIdNil(b bool)`

 SetCloudCredentialIdNil sets the value for CloudCredentialId to be an explicit nil

### UnsetCloudCredentialId
`func (o *ProjectForListDto) UnsetCloudCredentialId()`

UnsetCloudCredentialId ensures that no value is present for CloudCredentialId, not even an explicit nil
### GetOrganizationName

`func (o *ProjectForListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectForListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectForListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectForListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ProjectForListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectForListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ProjectForListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectForListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectForListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectForListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPartnerId

`func (o *ProjectForListDto) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ProjectForListDto) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ProjectForListDto) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ProjectForListDto) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *ProjectForListDto) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *ProjectForListDto) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetStatus

`func (o *ProjectForListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectForListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectForListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectForListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProjectForListDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectForListDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHealth

`func (o *ProjectForListDto) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectForListDto) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectForListDto) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectForListDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### SetHealthNil

`func (o *ProjectForListDto) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *ProjectForListDto) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil
### GetAccessIp

`func (o *ProjectForListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectForListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectForListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.

### HasAccessIp

`func (o *ProjectForListDto) HasAccessIp() bool`

HasAccessIp returns a boolean if a field has been set.

### SetAccessIpNil

`func (o *ProjectForListDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *ProjectForListDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetCloudType

`func (o *ProjectForListDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectForListDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectForListDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectForListDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### SetCloudTypeNil

`func (o *ProjectForListDto) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *ProjectForListDto) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetKubesprayCurrentVersion

`func (o *ProjectForListDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectForListDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectForListDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *ProjectForListDto) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### SetKubesprayCurrentVersionNil

`func (o *ProjectForListDto) SetKubesprayCurrentVersionNil(b bool)`

 SetKubesprayCurrentVersionNil sets the value for KubesprayCurrentVersion to be an explicit nil

### UnsetKubesprayCurrentVersion
`func (o *ProjectForListDto) UnsetKubesprayCurrentVersion()`

UnsetKubesprayCurrentVersion ensures that no value is present for KubesprayCurrentVersion, not even an explicit nil
### GetKubesprayTargetVersion

`func (o *ProjectForListDto) GetKubesprayTargetVersion() string`

GetKubesprayTargetVersion returns the KubesprayTargetVersion field if non-nil, zero value otherwise.

### GetKubesprayTargetVersionOk

`func (o *ProjectForListDto) GetKubesprayTargetVersionOk() (*string, bool)`

GetKubesprayTargetVersionOk returns a tuple with the KubesprayTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayTargetVersion

`func (o *ProjectForListDto) SetKubesprayTargetVersion(v string)`

SetKubesprayTargetVersion sets KubesprayTargetVersion field to given value.

### HasKubesprayTargetVersion

`func (o *ProjectForListDto) HasKubesprayTargetVersion() bool`

HasKubesprayTargetVersion returns a boolean if a field has been set.

### SetKubesprayTargetVersionNil

`func (o *ProjectForListDto) SetKubesprayTargetVersionNil(b bool)`

 SetKubesprayTargetVersionNil sets the value for KubesprayTargetVersion to be an explicit nil

### UnsetKubesprayTargetVersion
`func (o *ProjectForListDto) UnsetKubesprayTargetVersion()`

UnsetKubesprayTargetVersion ensures that no value is present for KubesprayTargetVersion, not even an explicit nil
### GetKubernetesCurrentVersion

`func (o *ProjectForListDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectForListDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectForListDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectForListDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *ProjectForListDto) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *ProjectForListDto) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil
### GetKubernetesTargetVersion

`func (o *ProjectForListDto) GetKubernetesTargetVersion() string`

GetKubernetesTargetVersion returns the KubernetesTargetVersion field if non-nil, zero value otherwise.

### GetKubernetesTargetVersionOk

`func (o *ProjectForListDto) GetKubernetesTargetVersionOk() (*string, bool)`

GetKubernetesTargetVersionOk returns a tuple with the KubernetesTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetVersion

`func (o *ProjectForListDto) SetKubernetesTargetVersion(v string)`

SetKubernetesTargetVersion sets KubernetesTargetVersion field to given value.

### HasKubernetesTargetVersion

`func (o *ProjectForListDto) HasKubernetesTargetVersion() bool`

HasKubernetesTargetVersion returns a boolean if a field has been set.

### SetKubernetesTargetVersionNil

`func (o *ProjectForListDto) SetKubernetesTargetVersionNil(b bool)`

 SetKubernetesTargetVersionNil sets the value for KubernetesTargetVersion to be an explicit nil

### UnsetKubernetesTargetVersion
`func (o *ProjectForListDto) UnsetKubernetesTargetVersion()`

UnsetKubernetesTargetVersion ensures that no value is present for KubernetesTargetVersion, not even an explicit nil
### GetUpdatedAt

`func (o *ProjectForListDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectForListDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectForListDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectForListDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProjectForListDto) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProjectForListDto) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetMonitoringCredential

`func (o *ProjectForListDto) GetMonitoringCredential() MonitoringCredentialsListDto`

GetMonitoringCredential returns the MonitoringCredential field if non-nil, zero value otherwise.

### GetMonitoringCredentialOk

`func (o *ProjectForListDto) GetMonitoringCredentialOk() (*MonitoringCredentialsListDto, bool)`

GetMonitoringCredentialOk returns a tuple with the MonitoringCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringCredential

`func (o *ProjectForListDto) SetMonitoringCredential(v MonitoringCredentialsListDto)`

SetMonitoringCredential sets MonitoringCredential field to given value.

### HasMonitoringCredential

`func (o *ProjectForListDto) HasMonitoringCredential() bool`

HasMonitoringCredential returns a boolean if a field has been set.

### GetIsAutoscalingEnabled

`func (o *ProjectForListDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectForListDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectForListDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.

### HasIsAutoscalingEnabled

`func (o *ProjectForListDto) HasIsAutoscalingEnabled() bool`

HasIsAutoscalingEnabled returns a boolean if a field has been set.

### GetFlavors

`func (o *ProjectForListDto) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *ProjectForListDto) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *ProjectForListDto) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *ProjectForListDto) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### SetFlavorsNil

`func (o *ProjectForListDto) SetFlavorsNil(b bool)`

 SetFlavorsNil sets the value for Flavors to be an explicit nil

### UnsetFlavors
`func (o *ProjectForListDto) UnsetFlavors()`

UnsetFlavors ensures that no value is present for Flavors, not even an explicit nil
### GetAccessProfile

`func (o *ProjectForListDto) GetAccessProfile() AccessProfilesForProjectListDto`

GetAccessProfile returns the AccessProfile field if non-nil, zero value otherwise.

### GetAccessProfileOk

`func (o *ProjectForListDto) GetAccessProfileOk() (*AccessProfilesForProjectListDto, bool)`

GetAccessProfileOk returns a tuple with the AccessProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfile

`func (o *ProjectForListDto) SetAccessProfile(v AccessProfilesForProjectListDto)`

SetAccessProfile sets AccessProfile field to given value.

### HasAccessProfile

`func (o *ProjectForListDto) HasAccessProfile() bool`

HasAccessProfile returns a boolean if a field has been set.

### GetKubernetesProfiles

`func (o *ProjectForListDto) GetKubernetesProfiles() KubernetesProfilesLisForPollerDto`

GetKubernetesProfiles returns the KubernetesProfiles field if non-nil, zero value otherwise.

### GetKubernetesProfilesOk

`func (o *ProjectForListDto) GetKubernetesProfilesOk() (*KubernetesProfilesLisForPollerDto, bool)`

GetKubernetesProfilesOk returns a tuple with the KubernetesProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfiles

`func (o *ProjectForListDto) SetKubernetesProfiles(v KubernetesProfilesLisForPollerDto)`

SetKubernetesProfiles sets KubernetesProfiles field to given value.

### HasKubernetesProfiles

`func (o *ProjectForListDto) HasKubernetesProfiles() bool`

HasKubernetesProfiles returns a boolean if a field has been set.

### GetOpaProfile

`func (o *ProjectForListDto) GetOpaProfile() OpaProfileListDto`

GetOpaProfile returns the OpaProfile field if non-nil, zero value otherwise.

### GetOpaProfileOk

`func (o *ProjectForListDto) GetOpaProfileOk() (*OpaProfileListDto, bool)`

GetOpaProfileOk returns a tuple with the OpaProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfile

`func (o *ProjectForListDto) SetOpaProfile(v OpaProfileListDto)`

SetOpaProfile sets OpaProfile field to given value.

### HasOpaProfile

`func (o *ProjectForListDto) HasOpaProfile() bool`

HasOpaProfile returns a boolean if a field has been set.

### GetKubernetesAlerts

`func (o *ProjectForListDto) GetKubernetesAlerts() []KubernetesAlertDto`

GetKubernetesAlerts returns the KubernetesAlerts field if non-nil, zero value otherwise.

### GetKubernetesAlertsOk

`func (o *ProjectForListDto) GetKubernetesAlertsOk() (*[]KubernetesAlertDto, bool)`

GetKubernetesAlertsOk returns a tuple with the KubernetesAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesAlerts

`func (o *ProjectForListDto) SetKubernetesAlerts(v []KubernetesAlertDto)`

SetKubernetesAlerts sets KubernetesAlerts field to given value.

### HasKubernetesAlerts

`func (o *ProjectForListDto) HasKubernetesAlerts() bool`

HasKubernetesAlerts returns a boolean if a field has been set.

### SetKubernetesAlertsNil

`func (o *ProjectForListDto) SetKubernetesAlertsNil(b bool)`

 SetKubernetesAlertsNil sets the value for KubernetesAlerts to be an explicit nil

### UnsetKubernetesAlerts
`func (o *ProjectForListDto) UnsetKubernetesAlerts()`

UnsetKubernetesAlerts ensures that no value is present for KubernetesAlerts, not even an explicit nil
### GetIsDeleteCluster

`func (o *ProjectForListDto) GetIsDeleteCluster() bool`

GetIsDeleteCluster returns the IsDeleteCluster field if non-nil, zero value otherwise.

### GetIsDeleteClusterOk

`func (o *ProjectForListDto) GetIsDeleteClusterOk() (*bool, bool)`

GetIsDeleteClusterOk returns a tuple with the IsDeleteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteCluster

`func (o *ProjectForListDto) SetIsDeleteCluster(v bool)`

SetIsDeleteCluster sets IsDeleteCluster field to given value.

### HasIsDeleteCluster

`func (o *ProjectForListDto) HasIsDeleteCluster() bool`

HasIsDeleteCluster returns a boolean if a field has been set.

### GetTaikunPrivateSSHKey

`func (o *ProjectForListDto) GetTaikunPrivateSSHKey() string`

GetTaikunPrivateSSHKey returns the TaikunPrivateSSHKey field if non-nil, zero value otherwise.

### GetTaikunPrivateSSHKeyOk

`func (o *ProjectForListDto) GetTaikunPrivateSSHKeyOk() (*string, bool)`

GetTaikunPrivateSSHKeyOk returns a tuple with the TaikunPrivateSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPrivateSSHKey

`func (o *ProjectForListDto) SetTaikunPrivateSSHKey(v string)`

SetTaikunPrivateSSHKey sets TaikunPrivateSSHKey field to given value.

### HasTaikunPrivateSSHKey

`func (o *ProjectForListDto) HasTaikunPrivateSSHKey() bool`

HasTaikunPrivateSSHKey returns a boolean if a field has been set.

### SetTaikunPrivateSSHKeyNil

`func (o *ProjectForListDto) SetTaikunPrivateSSHKeyNil(b bool)`

 SetTaikunPrivateSSHKeyNil sets the value for TaikunPrivateSSHKey to be an explicit nil

### UnsetTaikunPrivateSSHKey
`func (o *ProjectForListDto) UnsetTaikunPrivateSSHKey()`

UnsetTaikunPrivateSSHKey ensures that no value is present for TaikunPrivateSSHKey, not even an explicit nil
### GetTaikunPublicSSHKey

`func (o *ProjectForListDto) GetTaikunPublicSSHKey() string`

GetTaikunPublicSSHKey returns the TaikunPublicSSHKey field if non-nil, zero value otherwise.

### GetTaikunPublicSSHKeyOk

`func (o *ProjectForListDto) GetTaikunPublicSSHKeyOk() (*string, bool)`

GetTaikunPublicSSHKeyOk returns a tuple with the TaikunPublicSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunPublicSSHKey

`func (o *ProjectForListDto) SetTaikunPublicSSHKey(v string)`

SetTaikunPublicSSHKey sets TaikunPublicSSHKey field to given value.

### HasTaikunPublicSSHKey

`func (o *ProjectForListDto) HasTaikunPublicSSHKey() bool`

HasTaikunPublicSSHKey returns a boolean if a field has been set.

### SetTaikunPublicSSHKeyNil

`func (o *ProjectForListDto) SetTaikunPublicSSHKeyNil(b bool)`

 SetTaikunPublicSSHKeyNil sets the value for TaikunPublicSSHKey to be an explicit nil

### UnsetTaikunPublicSSHKey
`func (o *ProjectForListDto) UnsetTaikunPublicSSHKey()`

UnsetTaikunPublicSSHKey ensures that no value is present for TaikunPublicSSHKey, not even an explicit nil
### GetGoogleProjectId

`func (o *ProjectForListDto) GetGoogleProjectId() string`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *ProjectForListDto) GetGoogleProjectIdOk() (*string, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *ProjectForListDto) SetGoogleProjectId(v string)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *ProjectForListDto) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### SetGoogleProjectIdNil

`func (o *ProjectForListDto) SetGoogleProjectIdNil(b bool)`

 SetGoogleProjectIdNil sets the value for GoogleProjectId to be an explicit nil

### UnsetGoogleProjectId
`func (o *ProjectForListDto) UnsetGoogleProjectId()`

UnsetGoogleProjectId ensures that no value is present for GoogleProjectId, not even an explicit nil
### GetCidr

`func (o *ProjectForListDto) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ProjectForListDto) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ProjectForListDto) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ProjectForListDto) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *ProjectForListDto) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *ProjectForListDto) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetNetMask

`func (o *ProjectForListDto) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *ProjectForListDto) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *ProjectForListDto) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *ProjectForListDto) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### SetNetMaskNil

`func (o *ProjectForListDto) SetNetMaskNil(b bool)`

 SetNetMaskNil sets the value for NetMask to be an explicit nil

### UnsetNetMask
`func (o *ProjectForListDto) UnsetNetMask()`

UnsetNetMask ensures that no value is present for NetMask, not even an explicit nil
### GetPrivateIp

`func (o *ProjectForListDto) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ProjectForListDto) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ProjectForListDto) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ProjectForListDto) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *ProjectForListDto) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *ProjectForListDto) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetPublicIp

`func (o *ProjectForListDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ProjectForListDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ProjectForListDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ProjectForListDto) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *ProjectForListDto) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *ProjectForListDto) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetIsKubevapEnabled

`func (o *ProjectForListDto) GetIsKubevapEnabled() bool`

GetIsKubevapEnabled returns the IsKubevapEnabled field if non-nil, zero value otherwise.

### GetIsKubevapEnabledOk

`func (o *ProjectForListDto) GetIsKubevapEnabledOk() (*bool, bool)`

GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubevapEnabled

`func (o *ProjectForListDto) SetIsKubevapEnabled(v bool)`

SetIsKubevapEnabled sets IsKubevapEnabled field to given value.

### HasIsKubevapEnabled

`func (o *ProjectForListDto) HasIsKubevapEnabled() bool`

HasIsKubevapEnabled returns a boolean if a field has been set.

### GetTanzuReleaseVersion

`func (o *ProjectForListDto) GetTanzuReleaseVersion() string`

GetTanzuReleaseVersion returns the TanzuReleaseVersion field if non-nil, zero value otherwise.

### GetTanzuReleaseVersionOk

`func (o *ProjectForListDto) GetTanzuReleaseVersionOk() (*string, bool)`

GetTanzuReleaseVersionOk returns a tuple with the TanzuReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuReleaseVersion

`func (o *ProjectForListDto) SetTanzuReleaseVersion(v string)`

SetTanzuReleaseVersion sets TanzuReleaseVersion field to given value.

### HasTanzuReleaseVersion

`func (o *ProjectForListDto) HasTanzuReleaseVersion() bool`

HasTanzuReleaseVersion returns a boolean if a field has been set.

### SetTanzuReleaseVersionNil

`func (o *ProjectForListDto) SetTanzuReleaseVersionNil(b bool)`

 SetTanzuReleaseVersionNil sets the value for TanzuReleaseVersion to be an explicit nil

### UnsetTanzuReleaseVersion
`func (o *ProjectForListDto) UnsetTanzuReleaseVersion()`

UnsetTanzuReleaseVersion ensures that no value is present for TanzuReleaseVersion, not even an explicit nil
### GetNfsDiskSize

`func (o *ProjectForListDto) GetNfsDiskSize() int32`

GetNfsDiskSize returns the NfsDiskSize field if non-nil, zero value otherwise.

### GetNfsDiskSizeOk

`func (o *ProjectForListDto) GetNfsDiskSizeOk() (*int32, bool)`

GetNfsDiskSizeOk returns a tuple with the NfsDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsDiskSize

`func (o *ProjectForListDto) SetNfsDiskSize(v int32)`

SetNfsDiskSize sets NfsDiskSize field to given value.

### HasNfsDiskSize

`func (o *ProjectForListDto) HasNfsDiskSize() bool`

HasNfsDiskSize returns a boolean if a field has been set.

### GetKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) GetKubevapEnabeledKubernetesVersions() []string`

GetKubevapEnabeledKubernetesVersions returns the KubevapEnabeledKubernetesVersions field if non-nil, zero value otherwise.

### GetKubevapEnabeledKubernetesVersionsOk

`func (o *ProjectForListDto) GetKubevapEnabeledKubernetesVersionsOk() (*[]string, bool)`

GetKubevapEnabeledKubernetesVersionsOk returns a tuple with the KubevapEnabeledKubernetesVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) SetKubevapEnabeledKubernetesVersions(v []string)`

SetKubevapEnabeledKubernetesVersions sets KubevapEnabeledKubernetesVersions field to given value.

### HasKubevapEnabeledKubernetesVersions

`func (o *ProjectForListDto) HasKubevapEnabeledKubernetesVersions() bool`

HasKubevapEnabeledKubernetesVersions returns a boolean if a field has been set.

### SetKubevapEnabeledKubernetesVersionsNil

`func (o *ProjectForListDto) SetKubevapEnabeledKubernetesVersionsNil(b bool)`

 SetKubevapEnabeledKubernetesVersionsNil sets the value for KubevapEnabeledKubernetesVersions to be an explicit nil

### UnsetKubevapEnabeledKubernetesVersions
`func (o *ProjectForListDto) UnsetKubevapEnabeledKubernetesVersions()`

UnsetKubevapEnabeledKubernetesVersions ensures that no value is present for KubevapEnabeledKubernetesVersions, not even an explicit nil
### GetAwsProjectAZSubnets

`func (o *ProjectForListDto) GetAwsProjectAZSubnets() []AwsProjectAZSubnetDto`

GetAwsProjectAZSubnets returns the AwsProjectAZSubnets field if non-nil, zero value otherwise.

### GetAwsProjectAZSubnetsOk

`func (o *ProjectForListDto) GetAwsProjectAZSubnetsOk() (*[]AwsProjectAZSubnetDto, bool)`

GetAwsProjectAZSubnetsOk returns a tuple with the AwsProjectAZSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsProjectAZSubnets

`func (o *ProjectForListDto) SetAwsProjectAZSubnets(v []AwsProjectAZSubnetDto)`

SetAwsProjectAZSubnets sets AwsProjectAZSubnets field to given value.

### HasAwsProjectAZSubnets

`func (o *ProjectForListDto) HasAwsProjectAZSubnets() bool`

HasAwsProjectAZSubnets returns a boolean if a field has been set.

### SetAwsProjectAZSubnetsNil

`func (o *ProjectForListDto) SetAwsProjectAZSubnetsNil(b bool)`

 SetAwsProjectAZSubnetsNil sets the value for AwsProjectAZSubnets to be an explicit nil

### UnsetAwsProjectAZSubnets
`func (o *ProjectForListDto) UnsetAwsProjectAZSubnets()`

UnsetAwsProjectAZSubnets ensures that no value is present for AwsProjectAZSubnets, not even an explicit nil
### GetAvailabilityZones

`func (o *ProjectForListDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectForListDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectForListDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *ProjectForListDto) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *ProjectForListDto) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *ProjectForListDto) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetWorkersCount

`func (o *ProjectForListDto) GetWorkersCount() int32`

GetWorkersCount returns the WorkersCount field if non-nil, zero value otherwise.

### GetWorkersCountOk

`func (o *ProjectForListDto) GetWorkersCountOk() (*int32, bool)`

GetWorkersCountOk returns a tuple with the WorkersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkersCount

`func (o *ProjectForListDto) SetWorkersCount(v int32)`

SetWorkersCount sets WorkersCount field to given value.

### HasWorkersCount

`func (o *ProjectForListDto) HasWorkersCount() bool`

HasWorkersCount returns a boolean if a field has been set.

### GetTaikunLB

`func (o *ProjectForListDto) GetTaikunLB() TaikunLbDto`

GetTaikunLB returns the TaikunLB field if non-nil, zero value otherwise.

### GetTaikunLBOk

`func (o *ProjectForListDto) GetTaikunLBOk() (*TaikunLbDto, bool)`

GetTaikunLBOk returns a tuple with the TaikunLB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLB

`func (o *ProjectForListDto) SetTaikunLB(v TaikunLbDto)`

SetTaikunLB sets TaikunLB field to given value.

### HasTaikunLB

`func (o *ProjectForListDto) HasTaikunLB() bool`

HasTaikunLB returns a boolean if a field has been set.

### GetS3Credential

`func (o *ProjectForListDto) GetS3Credential() S3CredentialForProjectDto`

GetS3Credential returns the S3Credential field if non-nil, zero value otherwise.

### GetS3CredentialOk

`func (o *ProjectForListDto) GetS3CredentialOk() (*S3CredentialForProjectDto, bool)`

GetS3CredentialOk returns a tuple with the S3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Credential

`func (o *ProjectForListDto) SetS3Credential(v S3CredentialForProjectDto)`

SetS3Credential sets S3Credential field to given value.

### HasS3Credential

`func (o *ProjectForListDto) HasS3Credential() bool`

HasS3Credential returns a boolean if a field has been set.

### GetProjectRevision

`func (o *ProjectForListDto) GetProjectRevision() ProjectRevisionDto`

GetProjectRevision returns the ProjectRevision field if non-nil, zero value otherwise.

### GetProjectRevisionOk

`func (o *ProjectForListDto) GetProjectRevisionOk() (*ProjectRevisionDto, bool)`

GetProjectRevisionOk returns a tuple with the ProjectRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRevision

`func (o *ProjectForListDto) SetProjectRevision(v ProjectRevisionDto)`

SetProjectRevision sets ProjectRevision field to given value.

### HasProjectRevision

`func (o *ProjectForListDto) HasProjectRevision() bool`

HasProjectRevision returns a boolean if a field has been set.

### GetProjectActionDto

`func (o *ProjectForListDto) GetProjectActionDto() ProjectActionDto`

GetProjectActionDto returns the ProjectActionDto field if non-nil, zero value otherwise.

### GetProjectActionDtoOk

`func (o *ProjectForListDto) GetProjectActionDtoOk() (*ProjectActionDto, bool)`

GetProjectActionDtoOk returns a tuple with the ProjectActionDto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectActionDto

`func (o *ProjectForListDto) SetProjectActionDto(v ProjectActionDto)`

SetProjectActionDto sets ProjectActionDto field to given value.

### HasProjectActionDto

`func (o *ProjectForListDto) HasProjectActionDto() bool`

HasProjectActionDto returns a boolean if a field has been set.

### GetAiCredential

`func (o *ProjectForListDto) GetAiCredential() AiCredentialDto`

GetAiCredential returns the AiCredential field if non-nil, zero value otherwise.

### GetAiCredentialOk

`func (o *ProjectForListDto) GetAiCredentialOk() (*AiCredentialDto, bool)`

GetAiCredentialOk returns a tuple with the AiCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiCredential

`func (o *ProjectForListDto) SetAiCredential(v AiCredentialDto)`

SetAiCredential sets AiCredential field to given value.

### HasAiCredential

`func (o *ProjectForListDto) HasAiCredential() bool`

HasAiCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


