# KubernetesProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Cni** | Pointer to **NullableString** |  | [optional] 
**OctaviaEnabled** | Pointer to **bool** |  | [optional] 
**ExposeNodePortOnBastion** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**TaikunLBEnabled** | Pointer to **bool** |  | [optional] 
**AllowSchedulingOnMaster** | Pointer to **bool** |  | [optional] 
**UniqueClusterName** | Pointer to **bool** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**ProxmoxStorage** | Pointer to **NullableString** |  | [optional] 
**NvidiaGpuOperatorEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubernetesProfilesListDto

`func NewKubernetesProfilesListDto() *KubernetesProfilesListDto`

NewKubernetesProfilesListDto instantiates a new KubernetesProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProfilesListDtoWithDefaults

`func NewKubernetesProfilesListDtoWithDefaults() *KubernetesProfilesListDto`

NewKubernetesProfilesListDtoWithDefaults instantiates a new KubernetesProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesProfilesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesProfilesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesProfilesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubernetesProfilesListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesProfilesListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *KubernetesProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KubernetesProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KubernetesProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *KubernetesProfilesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *KubernetesProfilesListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *KubernetesProfilesListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *KubernetesProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *KubernetesProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *KubernetesProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *KubernetesProfilesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *KubernetesProfilesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *KubernetesProfilesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCni

`func (o *KubernetesProfilesListDto) GetCni() string`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *KubernetesProfilesListDto) GetCniOk() (*string, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *KubernetesProfilesListDto) SetCni(v string)`

SetCni sets Cni field to given value.

### HasCni

`func (o *KubernetesProfilesListDto) HasCni() bool`

HasCni returns a boolean if a field has been set.

### SetCniNil

`func (o *KubernetesProfilesListDto) SetCniNil(b bool)`

 SetCniNil sets the value for Cni to be an explicit nil

### UnsetCni
`func (o *KubernetesProfilesListDto) UnsetCni()`

UnsetCni ensures that no value is present for Cni, not even an explicit nil
### GetOctaviaEnabled

`func (o *KubernetesProfilesListDto) GetOctaviaEnabled() bool`

GetOctaviaEnabled returns the OctaviaEnabled field if non-nil, zero value otherwise.

### GetOctaviaEnabledOk

`func (o *KubernetesProfilesListDto) GetOctaviaEnabledOk() (*bool, bool)`

GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctaviaEnabled

`func (o *KubernetesProfilesListDto) SetOctaviaEnabled(v bool)`

SetOctaviaEnabled sets OctaviaEnabled field to given value.

### HasOctaviaEnabled

`func (o *KubernetesProfilesListDto) HasOctaviaEnabled() bool`

HasOctaviaEnabled returns a boolean if a field has been set.

### GetExposeNodePortOnBastion

`func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastion() bool`

GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field if non-nil, zero value otherwise.

### GetExposeNodePortOnBastionOk

`func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastionOk() (*bool, bool)`

GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeNodePortOnBastion

`func (o *KubernetesProfilesListDto) SetExposeNodePortOnBastion(v bool)`

SetExposeNodePortOnBastion sets ExposeNodePortOnBastion field to given value.

### HasExposeNodePortOnBastion

`func (o *KubernetesProfilesListDto) HasExposeNodePortOnBastion() bool`

HasExposeNodePortOnBastion returns a boolean if a field has been set.

### GetIsLocked

`func (o *KubernetesProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *KubernetesProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *KubernetesProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *KubernetesProfilesListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetTaikunLBEnabled

`func (o *KubernetesProfilesListDto) GetTaikunLBEnabled() bool`

GetTaikunLBEnabled returns the TaikunLBEnabled field if non-nil, zero value otherwise.

### GetTaikunLBEnabledOk

`func (o *KubernetesProfilesListDto) GetTaikunLBEnabledOk() (*bool, bool)`

GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBEnabled

`func (o *KubernetesProfilesListDto) SetTaikunLBEnabled(v bool)`

SetTaikunLBEnabled sets TaikunLBEnabled field to given value.

### HasTaikunLBEnabled

`func (o *KubernetesProfilesListDto) HasTaikunLBEnabled() bool`

HasTaikunLBEnabled returns a boolean if a field has been set.

### GetAllowSchedulingOnMaster

`func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMaster() bool`

GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field if non-nil, zero value otherwise.

### GetAllowSchedulingOnMasterOk

`func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMasterOk() (*bool, bool)`

GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSchedulingOnMaster

`func (o *KubernetesProfilesListDto) SetAllowSchedulingOnMaster(v bool)`

SetAllowSchedulingOnMaster sets AllowSchedulingOnMaster field to given value.

### HasAllowSchedulingOnMaster

`func (o *KubernetesProfilesListDto) HasAllowSchedulingOnMaster() bool`

HasAllowSchedulingOnMaster returns a boolean if a field has been set.

### GetUniqueClusterName

`func (o *KubernetesProfilesListDto) GetUniqueClusterName() bool`

GetUniqueClusterName returns the UniqueClusterName field if non-nil, zero value otherwise.

### GetUniqueClusterNameOk

`func (o *KubernetesProfilesListDto) GetUniqueClusterNameOk() (*bool, bool)`

GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClusterName

`func (o *KubernetesProfilesListDto) SetUniqueClusterName(v bool)`

SetUniqueClusterName sets UniqueClusterName field to given value.

### HasUniqueClusterName

`func (o *KubernetesProfilesListDto) HasUniqueClusterName() bool`

HasUniqueClusterName returns a boolean if a field has been set.

### GetProjects

`func (o *KubernetesProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *KubernetesProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *KubernetesProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *KubernetesProfilesListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *KubernetesProfilesListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *KubernetesProfilesListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetCreatedBy

`func (o *KubernetesProfilesListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KubernetesProfilesListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KubernetesProfilesListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KubernetesProfilesListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *KubernetesProfilesListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *KubernetesProfilesListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *KubernetesProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesProfilesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *KubernetesProfilesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *KubernetesProfilesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *KubernetesProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *KubernetesProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *KubernetesProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *KubernetesProfilesListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *KubernetesProfilesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *KubernetesProfilesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *KubernetesProfilesListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *KubernetesProfilesListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *KubernetesProfilesListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *KubernetesProfilesListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *KubernetesProfilesListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *KubernetesProfilesListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetProxmoxStorage

`func (o *KubernetesProfilesListDto) GetProxmoxStorage() string`

GetProxmoxStorage returns the ProxmoxStorage field if non-nil, zero value otherwise.

### GetProxmoxStorageOk

`func (o *KubernetesProfilesListDto) GetProxmoxStorageOk() (*string, bool)`

GetProxmoxStorageOk returns a tuple with the ProxmoxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxStorage

`func (o *KubernetesProfilesListDto) SetProxmoxStorage(v string)`

SetProxmoxStorage sets ProxmoxStorage field to given value.

### HasProxmoxStorage

`func (o *KubernetesProfilesListDto) HasProxmoxStorage() bool`

HasProxmoxStorage returns a boolean if a field has been set.

### SetProxmoxStorageNil

`func (o *KubernetesProfilesListDto) SetProxmoxStorageNil(b bool)`

 SetProxmoxStorageNil sets the value for ProxmoxStorage to be an explicit nil

### UnsetProxmoxStorage
`func (o *KubernetesProfilesListDto) UnsetProxmoxStorage()`

UnsetProxmoxStorage ensures that no value is present for ProxmoxStorage, not even an explicit nil
### GetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabled() bool`

GetNvidiaGpuOperatorEnabled returns the NvidiaGpuOperatorEnabled field if non-nil, zero value otherwise.

### GetNvidiaGpuOperatorEnabledOk

`func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabledOk() (*bool, bool)`

GetNvidiaGpuOperatorEnabledOk returns a tuple with the NvidiaGpuOperatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesListDto) SetNvidiaGpuOperatorEnabled(v bool)`

SetNvidiaGpuOperatorEnabled sets NvidiaGpuOperatorEnabled field to given value.

### HasNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesListDto) HasNvidiaGpuOperatorEnabled() bool`

HasNvidiaGpuOperatorEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


