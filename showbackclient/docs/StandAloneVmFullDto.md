# StandAloneVmFullDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | Pointer to **NullableString** |  | [optional] 
**VolumeSize** | Pointer to **int64** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**SshPublicKey** | Pointer to **NullableString** |  | [optional] 
**PublicIpEnabled** | Pointer to **bool** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**ShutOff** | Pointer to **bool** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**CloudInit** | Pointer to **NullableString** |  | [optional] 
**IsWindows** | Pointer to **bool** |  | [optional] 
**StandAloneProfile** | Pointer to [**StandAloneProfileFullDto**](StandAloneProfileFullDto.md) |  | [optional] 
**StandAloneVmDisks** | Pointer to [**[]StandAloneVmDiskFullDto**](StandAloneVmDiskFullDto.md) |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SpotPrice** | Pointer to **NullableString** |  | [optional] 
**Hypervisor** | Pointer to **NullableString** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStandAloneVmFullDto

`func NewStandAloneVmFullDto() *StandAloneVmFullDto`

NewStandAloneVmFullDto instantiates a new StandAloneVmFullDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneVmFullDtoWithDefaults

`func NewStandAloneVmFullDtoWithDefaults() *StandAloneVmFullDto`

NewStandAloneVmFullDtoWithDefaults instantiates a new StandAloneVmFullDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneVmFullDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneVmFullDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneVmFullDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandAloneVmFullDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandAloneVmFullDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneVmFullDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneVmFullDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandAloneVmFullDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandAloneVmFullDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneVmFullDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFlavorId

`func (o *StandAloneVmFullDto) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *StandAloneVmFullDto) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *StandAloneVmFullDto) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *StandAloneVmFullDto) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *StandAloneVmFullDto) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *StandAloneVmFullDto) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetVolumeSize

`func (o *StandAloneVmFullDto) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *StandAloneVmFullDto) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *StandAloneVmFullDto) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *StandAloneVmFullDto) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetRam

`func (o *StandAloneVmFullDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *StandAloneVmFullDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *StandAloneVmFullDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *StandAloneVmFullDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCpu

`func (o *StandAloneVmFullDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *StandAloneVmFullDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *StandAloneVmFullDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *StandAloneVmFullDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetVolumeType

`func (o *StandAloneVmFullDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandAloneVmFullDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandAloneVmFullDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StandAloneVmFullDto) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *StandAloneVmFullDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *StandAloneVmFullDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetSshPublicKey

`func (o *StandAloneVmFullDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *StandAloneVmFullDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *StandAloneVmFullDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *StandAloneVmFullDto) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### SetSshPublicKeyNil

`func (o *StandAloneVmFullDto) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *StandAloneVmFullDto) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
### GetPublicIpEnabled

`func (o *StandAloneVmFullDto) GetPublicIpEnabled() bool`

GetPublicIpEnabled returns the PublicIpEnabled field if non-nil, zero value otherwise.

### GetPublicIpEnabledOk

`func (o *StandAloneVmFullDto) GetPublicIpEnabledOk() (*bool, bool)`

GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpEnabled

`func (o *StandAloneVmFullDto) SetPublicIpEnabled(v bool)`

SetPublicIpEnabled sets PublicIpEnabled field to given value.

### HasPublicIpEnabled

`func (o *StandAloneVmFullDto) HasPublicIpEnabled() bool`

HasPublicIpEnabled returns a boolean if a field has been set.

### GetImageId

`func (o *StandAloneVmFullDto) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *StandAloneVmFullDto) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *StandAloneVmFullDto) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *StandAloneVmFullDto) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *StandAloneVmFullDto) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *StandAloneVmFullDto) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetImageName

`func (o *StandAloneVmFullDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *StandAloneVmFullDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *StandAloneVmFullDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *StandAloneVmFullDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *StandAloneVmFullDto) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *StandAloneVmFullDto) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetStatus

`func (o *StandAloneVmFullDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandAloneVmFullDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandAloneVmFullDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StandAloneVmFullDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StandAloneVmFullDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StandAloneVmFullDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRevision

`func (o *StandAloneVmFullDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StandAloneVmFullDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StandAloneVmFullDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StandAloneVmFullDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetShutOff

`func (o *StandAloneVmFullDto) GetShutOff() bool`

GetShutOff returns the ShutOff field if non-nil, zero value otherwise.

### GetShutOffOk

`func (o *StandAloneVmFullDto) GetShutOffOk() (*bool, bool)`

GetShutOffOk returns a tuple with the ShutOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutOff

`func (o *StandAloneVmFullDto) SetShutOff(v bool)`

SetShutOff sets ShutOff field to given value.

### HasShutOff

`func (o *StandAloneVmFullDto) HasShutOff() bool`

HasShutOff returns a boolean if a field has been set.

### GetPublicIp

`func (o *StandAloneVmFullDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *StandAloneVmFullDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *StandAloneVmFullDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *StandAloneVmFullDto) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *StandAloneVmFullDto) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *StandAloneVmFullDto) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetIpAddress

`func (o *StandAloneVmFullDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StandAloneVmFullDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StandAloneVmFullDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StandAloneVmFullDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *StandAloneVmFullDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *StandAloneVmFullDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetCloudInit

`func (o *StandAloneVmFullDto) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *StandAloneVmFullDto) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *StandAloneVmFullDto) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *StandAloneVmFullDto) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### SetCloudInitNil

`func (o *StandAloneVmFullDto) SetCloudInitNil(b bool)`

 SetCloudInitNil sets the value for CloudInit to be an explicit nil

### UnsetCloudInit
`func (o *StandAloneVmFullDto) UnsetCloudInit()`

UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
### GetIsWindows

`func (o *StandAloneVmFullDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *StandAloneVmFullDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *StandAloneVmFullDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *StandAloneVmFullDto) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### GetStandAloneProfile

`func (o *StandAloneVmFullDto) GetStandAloneProfile() StandAloneProfileFullDto`

GetStandAloneProfile returns the StandAloneProfile field if non-nil, zero value otherwise.

### GetStandAloneProfileOk

`func (o *StandAloneVmFullDto) GetStandAloneProfileOk() (*StandAloneProfileFullDto, bool)`

GetStandAloneProfileOk returns a tuple with the StandAloneProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfile

`func (o *StandAloneVmFullDto) SetStandAloneProfile(v StandAloneProfileFullDto)`

SetStandAloneProfile sets StandAloneProfile field to given value.

### HasStandAloneProfile

`func (o *StandAloneVmFullDto) HasStandAloneProfile() bool`

HasStandAloneProfile returns a boolean if a field has been set.

### GetStandAloneVmDisks

`func (o *StandAloneVmFullDto) GetStandAloneVmDisks() []StandAloneVmDiskFullDto`

GetStandAloneVmDisks returns the StandAloneVmDisks field if non-nil, zero value otherwise.

### GetStandAloneVmDisksOk

`func (o *StandAloneVmFullDto) GetStandAloneVmDisksOk() (*[]StandAloneVmDiskFullDto, bool)`

GetStandAloneVmDisksOk returns a tuple with the StandAloneVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneVmDisks

`func (o *StandAloneVmFullDto) SetStandAloneVmDisks(v []StandAloneVmDiskFullDto)`

SetStandAloneVmDisks sets StandAloneVmDisks field to given value.

### HasStandAloneVmDisks

`func (o *StandAloneVmFullDto) HasStandAloneVmDisks() bool`

HasStandAloneVmDisks returns a boolean if a field has been set.

### SetStandAloneVmDisksNil

`func (o *StandAloneVmFullDto) SetStandAloneVmDisksNil(b bool)`

 SetStandAloneVmDisksNil sets the value for StandAloneVmDisks to be an explicit nil

### UnsetStandAloneVmDisks
`func (o *StandAloneVmFullDto) UnsetStandAloneVmDisks()`

UnsetStandAloneVmDisks ensures that no value is present for StandAloneVmDisks, not even an explicit nil
### GetUsername

`func (o *StandAloneVmFullDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StandAloneVmFullDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StandAloneVmFullDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StandAloneVmFullDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *StandAloneVmFullDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *StandAloneVmFullDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *StandAloneVmFullDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StandAloneVmFullDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StandAloneVmFullDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StandAloneVmFullDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StandAloneVmFullDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StandAloneVmFullDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSpotPrice

`func (o *StandAloneVmFullDto) GetSpotPrice() string`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *StandAloneVmFullDto) GetSpotPriceOk() (*string, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *StandAloneVmFullDto) SetSpotPrice(v string)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *StandAloneVmFullDto) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### SetSpotPriceNil

`func (o *StandAloneVmFullDto) SetSpotPriceNil(b bool)`

 SetSpotPriceNil sets the value for SpotPrice to be an explicit nil

### UnsetSpotPrice
`func (o *StandAloneVmFullDto) UnsetSpotPrice()`

UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
### GetHypervisor

`func (o *StandAloneVmFullDto) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *StandAloneVmFullDto) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *StandAloneVmFullDto) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *StandAloneVmFullDto) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### SetHypervisorNil

`func (o *StandAloneVmFullDto) SetHypervisorNil(b bool)`

 SetHypervisorNil sets the value for Hypervisor to be an explicit nil

### UnsetHypervisor
`func (o *StandAloneVmFullDto) UnsetHypervisor()`

UnsetHypervisor ensures that no value is present for Hypervisor, not even an explicit nil
### GetSpotInstance

`func (o *StandAloneVmFullDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *StandAloneVmFullDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *StandAloneVmFullDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *StandAloneVmFullDto) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *StandAloneVmFullDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *StandAloneVmFullDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *StandAloneVmFullDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *StandAloneVmFullDto) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *StandAloneVmFullDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *StandAloneVmFullDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


