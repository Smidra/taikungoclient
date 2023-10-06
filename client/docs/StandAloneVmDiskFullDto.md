# StandAloneVmDiskFullDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**DeviceName** | Pointer to **NullableString** |  | [optional] 
**LunId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStandAloneVmDiskFullDto

`func NewStandAloneVmDiskFullDto() *StandAloneVmDiskFullDto`

NewStandAloneVmDiskFullDto instantiates a new StandAloneVmDiskFullDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneVmDiskFullDtoWithDefaults

`func NewStandAloneVmDiskFullDtoWithDefaults() *StandAloneVmDiskFullDto`

NewStandAloneVmDiskFullDtoWithDefaults instantiates a new StandAloneVmDiskFullDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneVmDiskFullDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneVmDiskFullDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneVmDiskFullDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandAloneVmDiskFullDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandAloneVmDiskFullDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneVmDiskFullDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneVmDiskFullDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandAloneVmDiskFullDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandAloneVmDiskFullDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneVmDiskFullDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *StandAloneVmDiskFullDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StandAloneVmDiskFullDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StandAloneVmDiskFullDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StandAloneVmDiskFullDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *StandAloneVmDiskFullDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandAloneVmDiskFullDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandAloneVmDiskFullDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StandAloneVmDiskFullDto) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *StandAloneVmDiskFullDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *StandAloneVmDiskFullDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetDeviceName

`func (o *StandAloneVmDiskFullDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StandAloneVmDiskFullDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StandAloneVmDiskFullDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StandAloneVmDiskFullDto) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *StandAloneVmDiskFullDto) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *StandAloneVmDiskFullDto) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetLunId

`func (o *StandAloneVmDiskFullDto) GetLunId() string`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *StandAloneVmDiskFullDto) GetLunIdOk() (*string, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *StandAloneVmDiskFullDto) SetLunId(v string)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *StandAloneVmDiskFullDto) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### SetLunIdNil

`func (o *StandAloneVmDiskFullDto) SetLunIdNil(b bool)`

 SetLunIdNil sets the value for LunId to be an explicit nil

### UnsetLunId
`func (o *StandAloneVmDiskFullDto) UnsetLunId()`

UnsetLunId ensures that no value is present for LunId, not even an explicit nil
### GetStatus

`func (o *StandAloneVmDiskFullDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandAloneVmDiskFullDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandAloneVmDiskFullDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StandAloneVmDiskFullDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StandAloneVmDiskFullDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StandAloneVmDiskFullDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


