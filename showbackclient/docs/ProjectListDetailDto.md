# ProjectListDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**HasKubeConfigFile** | Pointer to **bool** |  | [optional] 
**CloudCredentialName** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Health** | Pointer to **NullableString** |  | [optional] 
**CloudType** | Pointer to **NullableString** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**KubesprayTargetVersion** | Pointer to **NullableString** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**KubernetesTargetVersion** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**AlertsCount** | Pointer to **int32** |  | [optional] 
**TotalServersCount** | Pointer to **int32** |  | [optional] 
**TotalStandaloneVmsCount** | Pointer to **int32** |  | [optional] 
**BoundUsers** | Pointer to [**[]UserDto**](UserDto.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**ExpiredAt** | Pointer to **NullableString** |  | [optional] 
**DeleteOnExpiration** | Pointer to **bool** |  | [optional] 
**CertificateExpiredAt** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**QuotaId** | Pointer to **int32** |  | [optional] 
**AllowFullSpotKubernetes** | Pointer to **bool** |  | [optional] 
**AllowSpotWorkers** | Pointer to **bool** |  | [optional] 
**AllowSpotVMs** | Pointer to **bool** |  | [optional] 
**MaxSpotPrice** | Pointer to **NullableFloat64** |  | [optional] 
**ProjectAction** | Pointer to **bool** |  | [optional] 
**HasExpirationWarning** | Pointer to **bool** |  | [optional] 
**TotalHourlyCost** | Pointer to **float64** |  | [optional] 
**IsAutoscalingEnabled** | Pointer to **bool** |  | [optional] 
**AiEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectListDetailDto

`func NewProjectListDetailDto() *ProjectListDetailDto`

NewProjectListDetailDto instantiates a new ProjectListDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListDetailDtoWithDefaults

`func NewProjectListDetailDtoWithDefaults() *ProjectListDetailDto`

NewProjectListDetailDtoWithDefaults instantiates a new ProjectListDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectListDetailDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectListDetailDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectListDetailDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectListDetailDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectListDetailDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectListDetailDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectListDetailDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectListDetailDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectListDetailDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectListDetailDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsKubernetes

`func (o *ProjectListDetailDto) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectListDetailDto) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectListDetailDto) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *ProjectListDetailDto) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectListDetailDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectListDetailDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectListDetailDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectListDetailDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetHasKubeConfigFile

`func (o *ProjectListDetailDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectListDetailDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectListDetailDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.

### HasHasKubeConfigFile

`func (o *ProjectListDetailDto) HasHasKubeConfigFile() bool`

HasHasKubeConfigFile returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *ProjectListDetailDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ProjectListDetailDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ProjectListDetailDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ProjectListDetailDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### SetCloudCredentialNameNil

`func (o *ProjectListDetailDto) SetCloudCredentialNameNil(b bool)`

 SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil

### UnsetCloudCredentialName
`func (o *ProjectListDetailDto) UnsetCloudCredentialName()`

UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
### GetOrganizationName

`func (o *ProjectListDetailDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectListDetailDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectListDetailDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectListDetailDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ProjectListDetailDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectListDetailDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ProjectListDetailDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectListDetailDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectListDetailDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectListDetailDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectListDetailDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListDetailDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListDetailDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectListDetailDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProjectListDetailDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectListDetailDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHealth

`func (o *ProjectListDetailDto) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectListDetailDto) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectListDetailDto) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectListDetailDto) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### SetHealthNil

`func (o *ProjectListDetailDto) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *ProjectListDetailDto) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil
### GetCloudType

`func (o *ProjectListDetailDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectListDetailDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectListDetailDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectListDetailDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### SetCloudTypeNil

`func (o *ProjectListDetailDto) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *ProjectListDetailDto) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetKubesprayCurrentVersion

`func (o *ProjectListDetailDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectListDetailDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectListDetailDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *ProjectListDetailDto) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### SetKubesprayCurrentVersionNil

`func (o *ProjectListDetailDto) SetKubesprayCurrentVersionNil(b bool)`

 SetKubesprayCurrentVersionNil sets the value for KubesprayCurrentVersion to be an explicit nil

### UnsetKubesprayCurrentVersion
`func (o *ProjectListDetailDto) UnsetKubesprayCurrentVersion()`

UnsetKubesprayCurrentVersion ensures that no value is present for KubesprayCurrentVersion, not even an explicit nil
### GetKubesprayTargetVersion

`func (o *ProjectListDetailDto) GetKubesprayTargetVersion() string`

GetKubesprayTargetVersion returns the KubesprayTargetVersion field if non-nil, zero value otherwise.

### GetKubesprayTargetVersionOk

`func (o *ProjectListDetailDto) GetKubesprayTargetVersionOk() (*string, bool)`

GetKubesprayTargetVersionOk returns a tuple with the KubesprayTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayTargetVersion

`func (o *ProjectListDetailDto) SetKubesprayTargetVersion(v string)`

SetKubesprayTargetVersion sets KubesprayTargetVersion field to given value.

### HasKubesprayTargetVersion

`func (o *ProjectListDetailDto) HasKubesprayTargetVersion() bool`

HasKubesprayTargetVersion returns a boolean if a field has been set.

### SetKubesprayTargetVersionNil

`func (o *ProjectListDetailDto) SetKubesprayTargetVersionNil(b bool)`

 SetKubesprayTargetVersionNil sets the value for KubesprayTargetVersion to be an explicit nil

### UnsetKubesprayTargetVersion
`func (o *ProjectListDetailDto) UnsetKubesprayTargetVersion()`

UnsetKubesprayTargetVersion ensures that no value is present for KubesprayTargetVersion, not even an explicit nil
### GetKubernetesCurrentVersion

`func (o *ProjectListDetailDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectListDetailDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectListDetailDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ProjectListDetailDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *ProjectListDetailDto) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *ProjectListDetailDto) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil
### GetKubernetesTargetVersion

`func (o *ProjectListDetailDto) GetKubernetesTargetVersion() string`

GetKubernetesTargetVersion returns the KubernetesTargetVersion field if non-nil, zero value otherwise.

### GetKubernetesTargetVersionOk

`func (o *ProjectListDetailDto) GetKubernetesTargetVersionOk() (*string, bool)`

GetKubernetesTargetVersionOk returns a tuple with the KubernetesTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetVersion

`func (o *ProjectListDetailDto) SetKubernetesTargetVersion(v string)`

SetKubernetesTargetVersion sets KubernetesTargetVersion field to given value.

### HasKubernetesTargetVersion

`func (o *ProjectListDetailDto) HasKubernetesTargetVersion() bool`

HasKubernetesTargetVersion returns a boolean if a field has been set.

### SetKubernetesTargetVersionNil

`func (o *ProjectListDetailDto) SetKubernetesTargetVersionNil(b bool)`

 SetKubernetesTargetVersionNil sets the value for KubernetesTargetVersion to be an explicit nil

### UnsetKubernetesTargetVersion
`func (o *ProjectListDetailDto) UnsetKubernetesTargetVersion()`

UnsetKubernetesTargetVersion ensures that no value is present for KubernetesTargetVersion, not even an explicit nil
### GetCreatedAt

`func (o *ProjectListDetailDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectListDetailDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectListDetailDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectListDetailDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ProjectListDetailDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ProjectListDetailDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetAlertsCount

`func (o *ProjectListDetailDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ProjectListDetailDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ProjectListDetailDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.

### HasAlertsCount

`func (o *ProjectListDetailDto) HasAlertsCount() bool`

HasAlertsCount returns a boolean if a field has been set.

### GetTotalServersCount

`func (o *ProjectListDetailDto) GetTotalServersCount() int32`

GetTotalServersCount returns the TotalServersCount field if non-nil, zero value otherwise.

### GetTotalServersCountOk

`func (o *ProjectListDetailDto) GetTotalServersCountOk() (*int32, bool)`

GetTotalServersCountOk returns a tuple with the TotalServersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServersCount

`func (o *ProjectListDetailDto) SetTotalServersCount(v int32)`

SetTotalServersCount sets TotalServersCount field to given value.

### HasTotalServersCount

`func (o *ProjectListDetailDto) HasTotalServersCount() bool`

HasTotalServersCount returns a boolean if a field has been set.

### GetTotalStandaloneVmsCount

`func (o *ProjectListDetailDto) GetTotalStandaloneVmsCount() int32`

GetTotalStandaloneVmsCount returns the TotalStandaloneVmsCount field if non-nil, zero value otherwise.

### GetTotalStandaloneVmsCountOk

`func (o *ProjectListDetailDto) GetTotalStandaloneVmsCountOk() (*int32, bool)`

GetTotalStandaloneVmsCountOk returns a tuple with the TotalStandaloneVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStandaloneVmsCount

`func (o *ProjectListDetailDto) SetTotalStandaloneVmsCount(v int32)`

SetTotalStandaloneVmsCount sets TotalStandaloneVmsCount field to given value.

### HasTotalStandaloneVmsCount

`func (o *ProjectListDetailDto) HasTotalStandaloneVmsCount() bool`

HasTotalStandaloneVmsCount returns a boolean if a field has been set.

### GetBoundUsers

`func (o *ProjectListDetailDto) GetBoundUsers() []UserDto`

GetBoundUsers returns the BoundUsers field if non-nil, zero value otherwise.

### GetBoundUsersOk

`func (o *ProjectListDetailDto) GetBoundUsersOk() (*[]UserDto, bool)`

GetBoundUsersOk returns a tuple with the BoundUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUsers

`func (o *ProjectListDetailDto) SetBoundUsers(v []UserDto)`

SetBoundUsers sets BoundUsers field to given value.

### HasBoundUsers

`func (o *ProjectListDetailDto) HasBoundUsers() bool`

HasBoundUsers returns a boolean if a field has been set.

### SetBoundUsersNil

`func (o *ProjectListDetailDto) SetBoundUsersNil(b bool)`

 SetBoundUsersNil sets the value for BoundUsers to be an explicit nil

### UnsetBoundUsers
`func (o *ProjectListDetailDto) UnsetBoundUsers()`

UnsetBoundUsers ensures that no value is present for BoundUsers, not even an explicit nil
### GetCreatedBy

`func (o *ProjectListDetailDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectListDetailDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectListDetailDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectListDetailDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectListDetailDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectListDetailDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *ProjectListDetailDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ProjectListDetailDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ProjectListDetailDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ProjectListDetailDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *ProjectListDetailDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ProjectListDetailDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetExpiredAt

`func (o *ProjectListDetailDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectListDetailDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectListDetailDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *ProjectListDetailDto) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *ProjectListDetailDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ProjectListDetailDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetDeleteOnExpiration

`func (o *ProjectListDetailDto) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *ProjectListDetailDto) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *ProjectListDetailDto) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.

### HasDeleteOnExpiration

`func (o *ProjectListDetailDto) HasDeleteOnExpiration() bool`

HasDeleteOnExpiration returns a boolean if a field has been set.

### GetCertificateExpiredAt

`func (o *ProjectListDetailDto) GetCertificateExpiredAt() string`

GetCertificateExpiredAt returns the CertificateExpiredAt field if non-nil, zero value otherwise.

### GetCertificateExpiredAtOk

`func (o *ProjectListDetailDto) GetCertificateExpiredAtOk() (*string, bool)`

GetCertificateExpiredAtOk returns a tuple with the CertificateExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpiredAt

`func (o *ProjectListDetailDto) SetCertificateExpiredAt(v string)`

SetCertificateExpiredAt sets CertificateExpiredAt field to given value.

### HasCertificateExpiredAt

`func (o *ProjectListDetailDto) HasCertificateExpiredAt() bool`

HasCertificateExpiredAt returns a boolean if a field has been set.

### SetCertificateExpiredAtNil

`func (o *ProjectListDetailDto) SetCertificateExpiredAtNil(b bool)`

 SetCertificateExpiredAtNil sets the value for CertificateExpiredAt to be an explicit nil

### UnsetCertificateExpiredAt
`func (o *ProjectListDetailDto) UnsetCertificateExpiredAt()`

UnsetCertificateExpiredAt ensures that no value is present for CertificateExpiredAt, not even an explicit nil
### GetLastModifiedBy

`func (o *ProjectListDetailDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProjectListDetailDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProjectListDetailDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ProjectListDetailDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *ProjectListDetailDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *ProjectListDetailDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetQuotaId

`func (o *ProjectListDetailDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectListDetailDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectListDetailDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *ProjectListDetailDto) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetAllowFullSpotKubernetes

`func (o *ProjectListDetailDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectListDetailDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectListDetailDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.

### HasAllowFullSpotKubernetes

`func (o *ProjectListDetailDto) HasAllowFullSpotKubernetes() bool`

HasAllowFullSpotKubernetes returns a boolean if a field has been set.

### GetAllowSpotWorkers

`func (o *ProjectListDetailDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectListDetailDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectListDetailDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.

### HasAllowSpotWorkers

`func (o *ProjectListDetailDto) HasAllowSpotWorkers() bool`

HasAllowSpotWorkers returns a boolean if a field has been set.

### GetAllowSpotVMs

`func (o *ProjectListDetailDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectListDetailDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectListDetailDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.

### HasAllowSpotVMs

`func (o *ProjectListDetailDto) HasAllowSpotVMs() bool`

HasAllowSpotVMs returns a boolean if a field has been set.

### GetMaxSpotPrice

`func (o *ProjectListDetailDto) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *ProjectListDetailDto) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *ProjectListDetailDto) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.

### HasMaxSpotPrice

`func (o *ProjectListDetailDto) HasMaxSpotPrice() bool`

HasMaxSpotPrice returns a boolean if a field has been set.

### SetMaxSpotPriceNil

`func (o *ProjectListDetailDto) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *ProjectListDetailDto) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetProjectAction

`func (o *ProjectListDetailDto) GetProjectAction() bool`

GetProjectAction returns the ProjectAction field if non-nil, zero value otherwise.

### GetProjectActionOk

`func (o *ProjectListDetailDto) GetProjectActionOk() (*bool, bool)`

GetProjectActionOk returns a tuple with the ProjectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAction

`func (o *ProjectListDetailDto) SetProjectAction(v bool)`

SetProjectAction sets ProjectAction field to given value.

### HasProjectAction

`func (o *ProjectListDetailDto) HasProjectAction() bool`

HasProjectAction returns a boolean if a field has been set.

### GetHasExpirationWarning

`func (o *ProjectListDetailDto) GetHasExpirationWarning() bool`

GetHasExpirationWarning returns the HasExpirationWarning field if non-nil, zero value otherwise.

### GetHasExpirationWarningOk

`func (o *ProjectListDetailDto) GetHasExpirationWarningOk() (*bool, bool)`

GetHasExpirationWarningOk returns a tuple with the HasExpirationWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpirationWarning

`func (o *ProjectListDetailDto) SetHasExpirationWarning(v bool)`

SetHasExpirationWarning sets HasExpirationWarning field to given value.

### HasHasExpirationWarning

`func (o *ProjectListDetailDto) HasHasExpirationWarning() bool`

HasHasExpirationWarning returns a boolean if a field has been set.

### GetTotalHourlyCost

`func (o *ProjectListDetailDto) GetTotalHourlyCost() float64`

GetTotalHourlyCost returns the TotalHourlyCost field if non-nil, zero value otherwise.

### GetTotalHourlyCostOk

`func (o *ProjectListDetailDto) GetTotalHourlyCostOk() (*float64, bool)`

GetTotalHourlyCostOk returns a tuple with the TotalHourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHourlyCost

`func (o *ProjectListDetailDto) SetTotalHourlyCost(v float64)`

SetTotalHourlyCost sets TotalHourlyCost field to given value.

### HasTotalHourlyCost

`func (o *ProjectListDetailDto) HasTotalHourlyCost() bool`

HasTotalHourlyCost returns a boolean if a field has been set.

### GetIsAutoscalingEnabled

`func (o *ProjectListDetailDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectListDetailDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectListDetailDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.

### HasIsAutoscalingEnabled

`func (o *ProjectListDetailDto) HasIsAutoscalingEnabled() bool`

HasIsAutoscalingEnabled returns a boolean if a field has been set.

### GetAiEnabled

`func (o *ProjectListDetailDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ProjectListDetailDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ProjectListDetailDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *ProjectListDetailDto) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


