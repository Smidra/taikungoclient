# StorageClassDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableString** |  | [optional] 
**Provisioner** | Pointer to **NullableString** |  | [optional] 
**ReclaimPolicy** | Pointer to **NullableString** |  | [optional] 
**VolumeBindingMode** | Pointer to **NullableString** |  | [optional] 
**AllowVolumeExpansion** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewStorageClassDto

`func NewStorageClassDto() *StorageClassDto`

NewStorageClassDto instantiates a new StorageClassDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassDtoWithDefaults

`func NewStorageClassDtoWithDefaults() *StorageClassDto`

NewStorageClassDtoWithDefaults instantiates a new StorageClassDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *StorageClassDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *StorageClassDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *StorageClassDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *StorageClassDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *StorageClassDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *StorageClassDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetAge

`func (o *StorageClassDto) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *StorageClassDto) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *StorageClassDto) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *StorageClassDto) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *StorageClassDto) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *StorageClassDto) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetProvisioner

`func (o *StorageClassDto) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *StorageClassDto) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *StorageClassDto) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.

### HasProvisioner

`func (o *StorageClassDto) HasProvisioner() bool`

HasProvisioner returns a boolean if a field has been set.

### SetProvisionerNil

`func (o *StorageClassDto) SetProvisionerNil(b bool)`

 SetProvisionerNil sets the value for Provisioner to be an explicit nil

### UnsetProvisioner
`func (o *StorageClassDto) UnsetProvisioner()`

UnsetProvisioner ensures that no value is present for Provisioner, not even an explicit nil
### GetReclaimPolicy

`func (o *StorageClassDto) GetReclaimPolicy() string`

GetReclaimPolicy returns the ReclaimPolicy field if non-nil, zero value otherwise.

### GetReclaimPolicyOk

`func (o *StorageClassDto) GetReclaimPolicyOk() (*string, bool)`

GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimPolicy

`func (o *StorageClassDto) SetReclaimPolicy(v string)`

SetReclaimPolicy sets ReclaimPolicy field to given value.

### HasReclaimPolicy

`func (o *StorageClassDto) HasReclaimPolicy() bool`

HasReclaimPolicy returns a boolean if a field has been set.

### SetReclaimPolicyNil

`func (o *StorageClassDto) SetReclaimPolicyNil(b bool)`

 SetReclaimPolicyNil sets the value for ReclaimPolicy to be an explicit nil

### UnsetReclaimPolicy
`func (o *StorageClassDto) UnsetReclaimPolicy()`

UnsetReclaimPolicy ensures that no value is present for ReclaimPolicy, not even an explicit nil
### GetVolumeBindingMode

`func (o *StorageClassDto) GetVolumeBindingMode() string`

GetVolumeBindingMode returns the VolumeBindingMode field if non-nil, zero value otherwise.

### GetVolumeBindingModeOk

`func (o *StorageClassDto) GetVolumeBindingModeOk() (*string, bool)`

GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBindingMode

`func (o *StorageClassDto) SetVolumeBindingMode(v string)`

SetVolumeBindingMode sets VolumeBindingMode field to given value.

### HasVolumeBindingMode

`func (o *StorageClassDto) HasVolumeBindingMode() bool`

HasVolumeBindingMode returns a boolean if a field has been set.

### SetVolumeBindingModeNil

`func (o *StorageClassDto) SetVolumeBindingModeNil(b bool)`

 SetVolumeBindingModeNil sets the value for VolumeBindingMode to be an explicit nil

### UnsetVolumeBindingMode
`func (o *StorageClassDto) UnsetVolumeBindingMode()`

UnsetVolumeBindingMode ensures that no value is present for VolumeBindingMode, not even an explicit nil
### GetAllowVolumeExpansion

`func (o *StorageClassDto) GetAllowVolumeExpansion() bool`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *StorageClassDto) GetAllowVolumeExpansionOk() (*bool, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *StorageClassDto) SetAllowVolumeExpansion(v bool)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.

### HasAllowVolumeExpansion

`func (o *StorageClassDto) HasAllowVolumeExpansion() bool`

HasAllowVolumeExpansion returns a boolean if a field has been set.

### SetAllowVolumeExpansionNil

`func (o *StorageClassDto) SetAllowVolumeExpansionNil(b bool)`

 SetAllowVolumeExpansionNil sets the value for AllowVolumeExpansion to be an explicit nil

### UnsetAllowVolumeExpansion
`func (o *StorageClassDto) UnsetAllowVolumeExpansion()`

UnsetAllowVolumeExpansion ensures that no value is present for AllowVolumeExpansion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


