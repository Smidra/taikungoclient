# UpdateStandAloneVmCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **NullableInt32** |  | [optional] 
**Disks** | Pointer to [**[]UpdateStandAloneVmDiskDto**](UpdateStandAloneVmDiskDto.md) |  | [optional] 

## Methods

### NewUpdateStandAloneVmCommand

`func NewUpdateStandAloneVmCommand() *UpdateStandAloneVmCommand`

NewUpdateStandAloneVmCommand instantiates a new UpdateStandAloneVmCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStandAloneVmCommandWithDefaults

`func NewUpdateStandAloneVmCommandWithDefaults() *UpdateStandAloneVmCommand`

NewUpdateStandAloneVmCommandWithDefaults instantiates a new UpdateStandAloneVmCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStandAloneVmCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStandAloneVmCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStandAloneVmCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStandAloneVmCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *UpdateStandAloneVmCommand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateStandAloneVmCommand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateStandAloneVmCommand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateStandAloneVmCommand) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *UpdateStandAloneVmCommand) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *UpdateStandAloneVmCommand) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetPublicIp

`func (o *UpdateStandAloneVmCommand) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *UpdateStandAloneVmCommand) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *UpdateStandAloneVmCommand) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *UpdateStandAloneVmCommand) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *UpdateStandAloneVmCommand) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *UpdateStandAloneVmCommand) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetInstanceId

`func (o *UpdateStandAloneVmCommand) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *UpdateStandAloneVmCommand) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *UpdateStandAloneVmCommand) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *UpdateStandAloneVmCommand) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *UpdateStandAloneVmCommand) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *UpdateStandAloneVmCommand) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetFlavorId

`func (o *UpdateStandAloneVmCommand) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *UpdateStandAloneVmCommand) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *UpdateStandAloneVmCommand) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *UpdateStandAloneVmCommand) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *UpdateStandAloneVmCommand) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *UpdateStandAloneVmCommand) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetRevision

`func (o *UpdateStandAloneVmCommand) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UpdateStandAloneVmCommand) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UpdateStandAloneVmCommand) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *UpdateStandAloneVmCommand) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *UpdateStandAloneVmCommand) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *UpdateStandAloneVmCommand) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetDisks

`func (o *UpdateStandAloneVmCommand) GetDisks() []UpdateStandAloneVmDiskDto`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *UpdateStandAloneVmCommand) GetDisksOk() (*[]UpdateStandAloneVmDiskDto, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *UpdateStandAloneVmCommand) SetDisks(v []UpdateStandAloneVmDiskDto)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *UpdateStandAloneVmCommand) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *UpdateStandAloneVmCommand) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *UpdateStandAloneVmCommand) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


