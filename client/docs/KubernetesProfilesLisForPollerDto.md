# KubernetesProfilesLisForPollerDto

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
**ProxmoxStorage** | Pointer to **NullableString** |  | [optional] 
**NvidiaGpuOperatorEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubernetesProfilesLisForPollerDto

`func NewKubernetesProfilesLisForPollerDto() *KubernetesProfilesLisForPollerDto`

NewKubernetesProfilesLisForPollerDto instantiates a new KubernetesProfilesLisForPollerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProfilesLisForPollerDtoWithDefaults

`func NewKubernetesProfilesLisForPollerDtoWithDefaults() *KubernetesProfilesLisForPollerDto`

NewKubernetesProfilesLisForPollerDtoWithDefaults instantiates a new KubernetesProfilesLisForPollerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesProfilesLisForPollerDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesProfilesLisForPollerDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesProfilesLisForPollerDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesProfilesLisForPollerDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesProfilesLisForPollerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesProfilesLisForPollerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesProfilesLisForPollerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesProfilesLisForPollerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubernetesProfilesLisForPollerDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesProfilesLisForPollerDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *KubernetesProfilesLisForPollerDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KubernetesProfilesLisForPollerDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KubernetesProfilesLisForPollerDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *KubernetesProfilesLisForPollerDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *KubernetesProfilesLisForPollerDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *KubernetesProfilesLisForPollerDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *KubernetesProfilesLisForPollerDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *KubernetesProfilesLisForPollerDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *KubernetesProfilesLisForPollerDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *KubernetesProfilesLisForPollerDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *KubernetesProfilesLisForPollerDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *KubernetesProfilesLisForPollerDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCni

`func (o *KubernetesProfilesLisForPollerDto) GetCni() string`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *KubernetesProfilesLisForPollerDto) GetCniOk() (*string, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *KubernetesProfilesLisForPollerDto) SetCni(v string)`

SetCni sets Cni field to given value.

### HasCni

`func (o *KubernetesProfilesLisForPollerDto) HasCni() bool`

HasCni returns a boolean if a field has been set.

### SetCniNil

`func (o *KubernetesProfilesLisForPollerDto) SetCniNil(b bool)`

 SetCniNil sets the value for Cni to be an explicit nil

### UnsetCni
`func (o *KubernetesProfilesLisForPollerDto) UnsetCni()`

UnsetCni ensures that no value is present for Cni, not even an explicit nil
### GetOctaviaEnabled

`func (o *KubernetesProfilesLisForPollerDto) GetOctaviaEnabled() bool`

GetOctaviaEnabled returns the OctaviaEnabled field if non-nil, zero value otherwise.

### GetOctaviaEnabledOk

`func (o *KubernetesProfilesLisForPollerDto) GetOctaviaEnabledOk() (*bool, bool)`

GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctaviaEnabled

`func (o *KubernetesProfilesLisForPollerDto) SetOctaviaEnabled(v bool)`

SetOctaviaEnabled sets OctaviaEnabled field to given value.

### HasOctaviaEnabled

`func (o *KubernetesProfilesLisForPollerDto) HasOctaviaEnabled() bool`

HasOctaviaEnabled returns a boolean if a field has been set.

### GetExposeNodePortOnBastion

`func (o *KubernetesProfilesLisForPollerDto) GetExposeNodePortOnBastion() bool`

GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field if non-nil, zero value otherwise.

### GetExposeNodePortOnBastionOk

`func (o *KubernetesProfilesLisForPollerDto) GetExposeNodePortOnBastionOk() (*bool, bool)`

GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeNodePortOnBastion

`func (o *KubernetesProfilesLisForPollerDto) SetExposeNodePortOnBastion(v bool)`

SetExposeNodePortOnBastion sets ExposeNodePortOnBastion field to given value.

### HasExposeNodePortOnBastion

`func (o *KubernetesProfilesLisForPollerDto) HasExposeNodePortOnBastion() bool`

HasExposeNodePortOnBastion returns a boolean if a field has been set.

### GetIsLocked

`func (o *KubernetesProfilesLisForPollerDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *KubernetesProfilesLisForPollerDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *KubernetesProfilesLisForPollerDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *KubernetesProfilesLisForPollerDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetTaikunLBEnabled

`func (o *KubernetesProfilesLisForPollerDto) GetTaikunLBEnabled() bool`

GetTaikunLBEnabled returns the TaikunLBEnabled field if non-nil, zero value otherwise.

### GetTaikunLBEnabledOk

`func (o *KubernetesProfilesLisForPollerDto) GetTaikunLBEnabledOk() (*bool, bool)`

GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBEnabled

`func (o *KubernetesProfilesLisForPollerDto) SetTaikunLBEnabled(v bool)`

SetTaikunLBEnabled sets TaikunLBEnabled field to given value.

### HasTaikunLBEnabled

`func (o *KubernetesProfilesLisForPollerDto) HasTaikunLBEnabled() bool`

HasTaikunLBEnabled returns a boolean if a field has been set.

### GetAllowSchedulingOnMaster

`func (o *KubernetesProfilesLisForPollerDto) GetAllowSchedulingOnMaster() bool`

GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field if non-nil, zero value otherwise.

### GetAllowSchedulingOnMasterOk

`func (o *KubernetesProfilesLisForPollerDto) GetAllowSchedulingOnMasterOk() (*bool, bool)`

GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSchedulingOnMaster

`func (o *KubernetesProfilesLisForPollerDto) SetAllowSchedulingOnMaster(v bool)`

SetAllowSchedulingOnMaster sets AllowSchedulingOnMaster field to given value.

### HasAllowSchedulingOnMaster

`func (o *KubernetesProfilesLisForPollerDto) HasAllowSchedulingOnMaster() bool`

HasAllowSchedulingOnMaster returns a boolean if a field has been set.

### GetUniqueClusterName

`func (o *KubernetesProfilesLisForPollerDto) GetUniqueClusterName() bool`

GetUniqueClusterName returns the UniqueClusterName field if non-nil, zero value otherwise.

### GetUniqueClusterNameOk

`func (o *KubernetesProfilesLisForPollerDto) GetUniqueClusterNameOk() (*bool, bool)`

GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClusterName

`func (o *KubernetesProfilesLisForPollerDto) SetUniqueClusterName(v bool)`

SetUniqueClusterName sets UniqueClusterName field to given value.

### HasUniqueClusterName

`func (o *KubernetesProfilesLisForPollerDto) HasUniqueClusterName() bool`

HasUniqueClusterName returns a boolean if a field has been set.

### GetProxmoxStorage

`func (o *KubernetesProfilesLisForPollerDto) GetProxmoxStorage() string`

GetProxmoxStorage returns the ProxmoxStorage field if non-nil, zero value otherwise.

### GetProxmoxStorageOk

`func (o *KubernetesProfilesLisForPollerDto) GetProxmoxStorageOk() (*string, bool)`

GetProxmoxStorageOk returns a tuple with the ProxmoxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxStorage

`func (o *KubernetesProfilesLisForPollerDto) SetProxmoxStorage(v string)`

SetProxmoxStorage sets ProxmoxStorage field to given value.

### HasProxmoxStorage

`func (o *KubernetesProfilesLisForPollerDto) HasProxmoxStorage() bool`

HasProxmoxStorage returns a boolean if a field has been set.

### SetProxmoxStorageNil

`func (o *KubernetesProfilesLisForPollerDto) SetProxmoxStorageNil(b bool)`

 SetProxmoxStorageNil sets the value for ProxmoxStorage to be an explicit nil

### UnsetProxmoxStorage
`func (o *KubernetesProfilesLisForPollerDto) UnsetProxmoxStorage()`

UnsetProxmoxStorage ensures that no value is present for ProxmoxStorage, not even an explicit nil
### GetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesLisForPollerDto) GetNvidiaGpuOperatorEnabled() bool`

GetNvidiaGpuOperatorEnabled returns the NvidiaGpuOperatorEnabled field if non-nil, zero value otherwise.

### GetNvidiaGpuOperatorEnabledOk

`func (o *KubernetesProfilesLisForPollerDto) GetNvidiaGpuOperatorEnabledOk() (*bool, bool)`

GetNvidiaGpuOperatorEnabledOk returns a tuple with the NvidiaGpuOperatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesLisForPollerDto) SetNvidiaGpuOperatorEnabled(v bool)`

SetNvidiaGpuOperatorEnabled sets NvidiaGpuOperatorEnabled field to given value.

### HasNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesLisForPollerDto) HasNvidiaGpuOperatorEnabled() bool`

HasNvidiaGpuOperatorEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


