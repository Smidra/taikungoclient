# CreateStandAloneDiskCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandaloneVmId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**DeviceName** | Pointer to **NullableString** |  | [optional] 
**LunId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateStandAloneDiskCommand

`func NewCreateStandAloneDiskCommand() *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommand instantiates a new CreateStandAloneDiskCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandAloneDiskCommandWithDefaults

`func NewCreateStandAloneDiskCommandWithDefaults() *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommandWithDefaults instantiates a new CreateStandAloneDiskCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmId() int32`

GetStandaloneVmId returns the StandaloneVmId field if non-nil, zero value otherwise.

### GetStandaloneVmIdOk

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmIdOk() (*int32, bool)`

GetStandaloneVmIdOk returns a tuple with the StandaloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) SetStandaloneVmId(v int32)`

SetStandaloneVmId sets StandaloneVmId field to given value.

### HasStandaloneVmId

`func (o *CreateStandAloneDiskCommand) HasStandaloneVmId() bool`

HasStandaloneVmId returns a boolean if a field has been set.

### GetName

`func (o *CreateStandAloneDiskCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandAloneDiskCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandAloneDiskCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStandAloneDiskCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateStandAloneDiskCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateStandAloneDiskCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *CreateStandAloneDiskCommand) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStandAloneDiskCommand) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStandAloneDiskCommand) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateStandAloneDiskCommand) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *CreateStandAloneDiskCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateStandAloneDiskCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateStandAloneDiskCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreateStandAloneDiskCommand) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CreateStandAloneDiskCommand) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CreateStandAloneDiskCommand) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetDeviceName

`func (o *CreateStandAloneDiskCommand) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CreateStandAloneDiskCommand) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CreateStandAloneDiskCommand) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *CreateStandAloneDiskCommand) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *CreateStandAloneDiskCommand) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *CreateStandAloneDiskCommand) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetLunId

`func (o *CreateStandAloneDiskCommand) GetLunId() int32`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *CreateStandAloneDiskCommand) GetLunIdOk() (*int32, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *CreateStandAloneDiskCommand) SetLunId(v int32)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *CreateStandAloneDiskCommand) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### SetLunIdNil

`func (o *CreateStandAloneDiskCommand) SetLunIdNil(b bool)`

 SetLunIdNil sets the value for LunId to be an explicit nil

### UnsetLunId
`func (o *CreateStandAloneDiskCommand) UnsetLunId()`

UnsetLunId ensures that no value is present for LunId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


