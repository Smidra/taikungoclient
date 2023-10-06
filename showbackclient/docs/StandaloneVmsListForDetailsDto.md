# StandaloneVmsListForDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**CloudInit** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**VolumeSize** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**SshPublicKey** | Pointer to **NullableString** |  | [optional] 
**CurrentFlavor** | Pointer to **NullableString** |  | [optional] 
**TargetFlavor** | Pointer to **NullableString** |  | [optional] 
**PublicIpEnabled** | Pointer to **bool** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**SpotPrice** | Pointer to **NullableString** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**ActionButtons** | Pointer to [**StandaloneVisibilityDto**](StandaloneVisibilityDto.md) |  | [optional] 
**IsWindows** | Pointer to **bool** |  | [optional] 
**Disks** | Pointer to [**[]StandAloneVmDiskForDetailsDto**](StandAloneVmDiskForDetailsDto.md) |  | [optional] 
**StandAloneMetaDatas** | Pointer to [**[]StandAloneMetaDataDtoForVm**](StandAloneMetaDataDtoForVm.md) |  | [optional] 
**Profile** | Pointer to [**StandAloneProfileForDetailsDto**](StandAloneProfileForDetailsDto.md) |  | [optional] 

## Methods

### NewStandaloneVmsListForDetailsDto

`func NewStandaloneVmsListForDetailsDto() *StandaloneVmsListForDetailsDto`

NewStandaloneVmsListForDetailsDto instantiates a new StandaloneVmsListForDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneVmsListForDetailsDtoWithDefaults

`func NewStandaloneVmsListForDetailsDtoWithDefaults() *StandaloneVmsListForDetailsDto`

NewStandaloneVmsListForDetailsDtoWithDefaults instantiates a new StandaloneVmsListForDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneVmsListForDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneVmsListForDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneVmsListForDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandaloneVmsListForDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandaloneVmsListForDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneVmsListForDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneVmsListForDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandaloneVmsListForDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandaloneVmsListForDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandaloneVmsListForDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageName

`func (o *StandaloneVmsListForDetailsDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *StandaloneVmsListForDetailsDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *StandaloneVmsListForDetailsDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *StandaloneVmsListForDetailsDto) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *StandaloneVmsListForDetailsDto) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *StandaloneVmsListForDetailsDto) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetImageId

`func (o *StandaloneVmsListForDetailsDto) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *StandaloneVmsListForDetailsDto) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *StandaloneVmsListForDetailsDto) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *StandaloneVmsListForDetailsDto) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *StandaloneVmsListForDetailsDto) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *StandaloneVmsListForDetailsDto) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetStatus

`func (o *StandaloneVmsListForDetailsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandaloneVmsListForDetailsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandaloneVmsListForDetailsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StandaloneVmsListForDetailsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StandaloneVmsListForDetailsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StandaloneVmsListForDetailsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCloudInit

`func (o *StandaloneVmsListForDetailsDto) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *StandaloneVmsListForDetailsDto) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *StandaloneVmsListForDetailsDto) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *StandaloneVmsListForDetailsDto) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### SetCloudInitNil

`func (o *StandaloneVmsListForDetailsDto) SetCloudInitNil(b bool)`

 SetCloudInitNil sets the value for CloudInit to be an explicit nil

### UnsetCloudInit
`func (o *StandaloneVmsListForDetailsDto) UnsetCloudInit()`

UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
### GetVolumeType

`func (o *StandaloneVmsListForDetailsDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandaloneVmsListForDetailsDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandaloneVmsListForDetailsDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StandaloneVmsListForDetailsDto) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *StandaloneVmsListForDetailsDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *StandaloneVmsListForDetailsDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetVolumeSize

`func (o *StandaloneVmsListForDetailsDto) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *StandaloneVmsListForDetailsDto) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *StandaloneVmsListForDetailsDto) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *StandaloneVmsListForDetailsDto) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StandaloneVmsListForDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StandaloneVmsListForDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StandaloneVmsListForDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StandaloneVmsListForDetailsDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *StandaloneVmsListForDetailsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StandaloneVmsListForDetailsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *StandaloneVmsListForDetailsDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StandaloneVmsListForDetailsDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StandaloneVmsListForDetailsDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *StandaloneVmsListForDetailsDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *StandaloneVmsListForDetailsDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *StandaloneVmsListForDetailsDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *StandaloneVmsListForDetailsDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StandaloneVmsListForDetailsDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StandaloneVmsListForDetailsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *StandaloneVmsListForDetailsDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *StandaloneVmsListForDetailsDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *StandaloneVmsListForDetailsDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *StandaloneVmsListForDetailsDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *StandaloneVmsListForDetailsDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *StandaloneVmsListForDetailsDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetSshPublicKey

`func (o *StandaloneVmsListForDetailsDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *StandaloneVmsListForDetailsDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *StandaloneVmsListForDetailsDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *StandaloneVmsListForDetailsDto) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### SetSshPublicKeyNil

`func (o *StandaloneVmsListForDetailsDto) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *StandaloneVmsListForDetailsDto) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
### GetCurrentFlavor

`func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavor() string`

GetCurrentFlavor returns the CurrentFlavor field if non-nil, zero value otherwise.

### GetCurrentFlavorOk

`func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavorOk() (*string, bool)`

GetCurrentFlavorOk returns a tuple with the CurrentFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFlavor

`func (o *StandaloneVmsListForDetailsDto) SetCurrentFlavor(v string)`

SetCurrentFlavor sets CurrentFlavor field to given value.

### HasCurrentFlavor

`func (o *StandaloneVmsListForDetailsDto) HasCurrentFlavor() bool`

HasCurrentFlavor returns a boolean if a field has been set.

### SetCurrentFlavorNil

`func (o *StandaloneVmsListForDetailsDto) SetCurrentFlavorNil(b bool)`

 SetCurrentFlavorNil sets the value for CurrentFlavor to be an explicit nil

### UnsetCurrentFlavor
`func (o *StandaloneVmsListForDetailsDto) UnsetCurrentFlavor()`

UnsetCurrentFlavor ensures that no value is present for CurrentFlavor, not even an explicit nil
### GetTargetFlavor

`func (o *StandaloneVmsListForDetailsDto) GetTargetFlavor() string`

GetTargetFlavor returns the TargetFlavor field if non-nil, zero value otherwise.

### GetTargetFlavorOk

`func (o *StandaloneVmsListForDetailsDto) GetTargetFlavorOk() (*string, bool)`

GetTargetFlavorOk returns a tuple with the TargetFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFlavor

`func (o *StandaloneVmsListForDetailsDto) SetTargetFlavor(v string)`

SetTargetFlavor sets TargetFlavor field to given value.

### HasTargetFlavor

`func (o *StandaloneVmsListForDetailsDto) HasTargetFlavor() bool`

HasTargetFlavor returns a boolean if a field has been set.

### SetTargetFlavorNil

`func (o *StandaloneVmsListForDetailsDto) SetTargetFlavorNil(b bool)`

 SetTargetFlavorNil sets the value for TargetFlavor to be an explicit nil

### UnsetTargetFlavor
`func (o *StandaloneVmsListForDetailsDto) UnsetTargetFlavor()`

UnsetTargetFlavor ensures that no value is present for TargetFlavor, not even an explicit nil
### GetPublicIpEnabled

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabled() bool`

GetPublicIpEnabled returns the PublicIpEnabled field if non-nil, zero value otherwise.

### GetPublicIpEnabledOk

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabledOk() (*bool, bool)`

GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpEnabled

`func (o *StandaloneVmsListForDetailsDto) SetPublicIpEnabled(v bool)`

SetPublicIpEnabled sets PublicIpEnabled field to given value.

### HasPublicIpEnabled

`func (o *StandaloneVmsListForDetailsDto) HasPublicIpEnabled() bool`

HasPublicIpEnabled returns a boolean if a field has been set.

### GetPublicIp

`func (o *StandaloneVmsListForDetailsDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *StandaloneVmsListForDetailsDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *StandaloneVmsListForDetailsDto) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *StandaloneVmsListForDetailsDto) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *StandaloneVmsListForDetailsDto) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetIpAddress

`func (o *StandaloneVmsListForDetailsDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StandaloneVmsListForDetailsDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StandaloneVmsListForDetailsDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StandaloneVmsListForDetailsDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *StandaloneVmsListForDetailsDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *StandaloneVmsListForDetailsDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetSpotPrice

`func (o *StandaloneVmsListForDetailsDto) GetSpotPrice() string`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *StandaloneVmsListForDetailsDto) GetSpotPriceOk() (*string, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *StandaloneVmsListForDetailsDto) SetSpotPrice(v string)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *StandaloneVmsListForDetailsDto) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### SetSpotPriceNil

`func (o *StandaloneVmsListForDetailsDto) SetSpotPriceNil(b bool)`

 SetSpotPriceNil sets the value for SpotPrice to be an explicit nil

### UnsetSpotPrice
`func (o *StandaloneVmsListForDetailsDto) UnsetSpotPrice()`

UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
### GetSpotInstance

`func (o *StandaloneVmsListForDetailsDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *StandaloneVmsListForDetailsDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *StandaloneVmsListForDetailsDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *StandaloneVmsListForDetailsDto) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetActionButtons

`func (o *StandaloneVmsListForDetailsDto) GetActionButtons() StandaloneVisibilityDto`

GetActionButtons returns the ActionButtons field if non-nil, zero value otherwise.

### GetActionButtonsOk

`func (o *StandaloneVmsListForDetailsDto) GetActionButtonsOk() (*StandaloneVisibilityDto, bool)`

GetActionButtonsOk returns a tuple with the ActionButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionButtons

`func (o *StandaloneVmsListForDetailsDto) SetActionButtons(v StandaloneVisibilityDto)`

SetActionButtons sets ActionButtons field to given value.

### HasActionButtons

`func (o *StandaloneVmsListForDetailsDto) HasActionButtons() bool`

HasActionButtons returns a boolean if a field has been set.

### GetIsWindows

`func (o *StandaloneVmsListForDetailsDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *StandaloneVmsListForDetailsDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *StandaloneVmsListForDetailsDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *StandaloneVmsListForDetailsDto) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### GetDisks

`func (o *StandaloneVmsListForDetailsDto) GetDisks() []StandAloneVmDiskForDetailsDto`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *StandaloneVmsListForDetailsDto) GetDisksOk() (*[]StandAloneVmDiskForDetailsDto, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *StandaloneVmsListForDetailsDto) SetDisks(v []StandAloneVmDiskForDetailsDto)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *StandaloneVmsListForDetailsDto) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *StandaloneVmsListForDetailsDto) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *StandaloneVmsListForDetailsDto) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetStandAloneMetaDatas

`func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatas() []StandAloneMetaDataDtoForVm`

GetStandAloneMetaDatas returns the StandAloneMetaDatas field if non-nil, zero value otherwise.

### GetStandAloneMetaDatasOk

`func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatasOk() (*[]StandAloneMetaDataDtoForVm, bool)`

GetStandAloneMetaDatasOk returns a tuple with the StandAloneMetaDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneMetaDatas

`func (o *StandaloneVmsListForDetailsDto) SetStandAloneMetaDatas(v []StandAloneMetaDataDtoForVm)`

SetStandAloneMetaDatas sets StandAloneMetaDatas field to given value.

### HasStandAloneMetaDatas

`func (o *StandaloneVmsListForDetailsDto) HasStandAloneMetaDatas() bool`

HasStandAloneMetaDatas returns a boolean if a field has been set.

### SetStandAloneMetaDatasNil

`func (o *StandaloneVmsListForDetailsDto) SetStandAloneMetaDatasNil(b bool)`

 SetStandAloneMetaDatasNil sets the value for StandAloneMetaDatas to be an explicit nil

### UnsetStandAloneMetaDatas
`func (o *StandaloneVmsListForDetailsDto) UnsetStandAloneMetaDatas()`

UnsetStandAloneMetaDatas ensures that no value is present for StandAloneMetaDatas, not even an explicit nil
### GetProfile

`func (o *StandaloneVmsListForDetailsDto) GetProfile() StandAloneProfileForDetailsDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *StandaloneVmsListForDetailsDto) GetProfileOk() (*StandAloneProfileForDetailsDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *StandaloneVmsListForDetailsDto) SetProfile(v StandAloneProfileForDetailsDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *StandaloneVmsListForDetailsDto) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


