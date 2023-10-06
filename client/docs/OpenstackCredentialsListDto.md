# OpenstackCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectCount** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**PublicNetwork** | Pointer to **NullableString** |  | [optional] 
**ImportNetwork** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**InternalSubnetId** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**ContinentName** | Pointer to **NullableString** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsInfra** | Pointer to **bool** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenstackCredentialsListDto

`func NewOpenstackCredentialsListDto() *OpenstackCredentialsListDto`

NewOpenstackCredentialsListDto instantiates a new OpenstackCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackCredentialsListDtoWithDefaults

`func NewOpenstackCredentialsListDtoWithDefaults() *OpenstackCredentialsListDto`

NewOpenstackCredentialsListDtoWithDefaults instantiates a new OpenstackCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenstackCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenstackCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenstackCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpenstackCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectCount

`func (o *OpenstackCredentialsListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *OpenstackCredentialsListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *OpenstackCredentialsListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *OpenstackCredentialsListDto) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### GetIsLocked

`func (o *OpenstackCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OpenstackCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OpenstackCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *OpenstackCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetName

`func (o *OpenstackCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenstackCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenstackCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenstackCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenstackCredentialsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenstackCredentialsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUser

`func (o *OpenstackCredentialsListDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OpenstackCredentialsListDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OpenstackCredentialsListDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OpenstackCredentialsListDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *OpenstackCredentialsListDto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *OpenstackCredentialsListDto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrl

`func (o *OpenstackCredentialsListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenstackCredentialsListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenstackCredentialsListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenstackCredentialsListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *OpenstackCredentialsListDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *OpenstackCredentialsListDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetProject

`func (o *OpenstackCredentialsListDto) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OpenstackCredentialsListDto) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OpenstackCredentialsListDto) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *OpenstackCredentialsListDto) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *OpenstackCredentialsListDto) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *OpenstackCredentialsListDto) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetDomain

`func (o *OpenstackCredentialsListDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OpenstackCredentialsListDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OpenstackCredentialsListDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OpenstackCredentialsListDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *OpenstackCredentialsListDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *OpenstackCredentialsListDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetRegion

`func (o *OpenstackCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OpenstackCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OpenstackCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OpenstackCredentialsListDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *OpenstackCredentialsListDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *OpenstackCredentialsListDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPublicNetwork

`func (o *OpenstackCredentialsListDto) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *OpenstackCredentialsListDto) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *OpenstackCredentialsListDto) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *OpenstackCredentialsListDto) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### SetPublicNetworkNil

`func (o *OpenstackCredentialsListDto) SetPublicNetworkNil(b bool)`

 SetPublicNetworkNil sets the value for PublicNetwork to be an explicit nil

### UnsetPublicNetwork
`func (o *OpenstackCredentialsListDto) UnsetPublicNetwork()`

UnsetPublicNetwork ensures that no value is present for PublicNetwork, not even an explicit nil
### GetImportNetwork

`func (o *OpenstackCredentialsListDto) GetImportNetwork() bool`

GetImportNetwork returns the ImportNetwork field if non-nil, zero value otherwise.

### GetImportNetworkOk

`func (o *OpenstackCredentialsListDto) GetImportNetworkOk() (*bool, bool)`

GetImportNetworkOk returns a tuple with the ImportNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNetwork

`func (o *OpenstackCredentialsListDto) SetImportNetwork(v bool)`

SetImportNetwork sets ImportNetwork field to given value.

### HasImportNetwork

`func (o *OpenstackCredentialsListDto) HasImportNetwork() bool`

HasImportNetwork returns a boolean if a field has been set.

### GetTenantId

`func (o *OpenstackCredentialsListDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OpenstackCredentialsListDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OpenstackCredentialsListDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OpenstackCredentialsListDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *OpenstackCredentialsListDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *OpenstackCredentialsListDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAvailabilityZone

`func (o *OpenstackCredentialsListDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *OpenstackCredentialsListDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *OpenstackCredentialsListDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *OpenstackCredentialsListDto) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *OpenstackCredentialsListDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *OpenstackCredentialsListDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetVolumeType

`func (o *OpenstackCredentialsListDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *OpenstackCredentialsListDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *OpenstackCredentialsListDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *OpenstackCredentialsListDto) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *OpenstackCredentialsListDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *OpenstackCredentialsListDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetInternalSubnetId

`func (o *OpenstackCredentialsListDto) GetInternalSubnetId() string`

GetInternalSubnetId returns the InternalSubnetId field if non-nil, zero value otherwise.

### GetInternalSubnetIdOk

`func (o *OpenstackCredentialsListDto) GetInternalSubnetIdOk() (*string, bool)`

GetInternalSubnetIdOk returns a tuple with the InternalSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSubnetId

`func (o *OpenstackCredentialsListDto) SetInternalSubnetId(v string)`

SetInternalSubnetId sets InternalSubnetId field to given value.

### HasInternalSubnetId

`func (o *OpenstackCredentialsListDto) HasInternalSubnetId() bool`

HasInternalSubnetId returns a boolean if a field has been set.

### SetInternalSubnetIdNil

`func (o *OpenstackCredentialsListDto) SetInternalSubnetIdNil(b bool)`

 SetInternalSubnetIdNil sets the value for InternalSubnetId to be an explicit nil

### UnsetInternalSubnetId
`func (o *OpenstackCredentialsListDto) UnsetInternalSubnetId()`

UnsetInternalSubnetId ensures that no value is present for InternalSubnetId, not even an explicit nil
### GetProjects

`func (o *OpenstackCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OpenstackCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OpenstackCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OpenstackCredentialsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *OpenstackCredentialsListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *OpenstackCredentialsListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetCreatedBy

`func (o *OpenstackCredentialsListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OpenstackCredentialsListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OpenstackCredentialsListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OpenstackCredentialsListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OpenstackCredentialsListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OpenstackCredentialsListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *OpenstackCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *OpenstackCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *OpenstackCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *OpenstackCredentialsListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *OpenstackCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *OpenstackCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *OpenstackCredentialsListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *OpenstackCredentialsListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *OpenstackCredentialsListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *OpenstackCredentialsListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *OpenstackCredentialsListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *OpenstackCredentialsListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetIsDefault

`func (o *OpenstackCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OpenstackCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OpenstackCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OpenstackCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OpenstackCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpenstackCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpenstackCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OpenstackCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OpenstackCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OpenstackCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OpenstackCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OpenstackCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OpenstackCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OpenstackCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCreatedAt

`func (o *OpenstackCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenstackCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenstackCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpenstackCredentialsListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *OpenstackCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OpenstackCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetContinentName

`func (o *OpenstackCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *OpenstackCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *OpenstackCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *OpenstackCredentialsListDto) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.

### SetContinentNameNil

`func (o *OpenstackCredentialsListDto) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *OpenstackCredentialsListDto) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
### GetIsAdmin

`func (o *OpenstackCredentialsListDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenstackCredentialsListDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenstackCredentialsListDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OpenstackCredentialsListDto) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsInfra

`func (o *OpenstackCredentialsListDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *OpenstackCredentialsListDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *OpenstackCredentialsListDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.

### HasIsInfra

`func (o *OpenstackCredentialsListDto) HasIsInfra() bool`

HasIsInfra returns a boolean if a field has been set.

### GetApplicationCredEnabled

`func (o *OpenstackCredentialsListDto) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenstackCredentialsListDto) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenstackCredentialsListDto) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *OpenstackCredentialsListDto) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


