# DeleteStandAloneVmDiskCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandaloneVmId** | Pointer to **int32** |  | [optional] 
**VmDiskIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewDeleteStandAloneVmDiskCommand

`func NewDeleteStandAloneVmDiskCommand() *DeleteStandAloneVmDiskCommand`

NewDeleteStandAloneVmDiskCommand instantiates a new DeleteStandAloneVmDiskCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteStandAloneVmDiskCommandWithDefaults

`func NewDeleteStandAloneVmDiskCommandWithDefaults() *DeleteStandAloneVmDiskCommand`

NewDeleteStandAloneVmDiskCommandWithDefaults instantiates a new DeleteStandAloneVmDiskCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandaloneVmId

`func (o *DeleteStandAloneVmDiskCommand) GetStandaloneVmId() int32`

GetStandaloneVmId returns the StandaloneVmId field if non-nil, zero value otherwise.

### GetStandaloneVmIdOk

`func (o *DeleteStandAloneVmDiskCommand) GetStandaloneVmIdOk() (*int32, bool)`

GetStandaloneVmIdOk returns a tuple with the StandaloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmId

`func (o *DeleteStandAloneVmDiskCommand) SetStandaloneVmId(v int32)`

SetStandaloneVmId sets StandaloneVmId field to given value.

### HasStandaloneVmId

`func (o *DeleteStandAloneVmDiskCommand) HasStandaloneVmId() bool`

HasStandaloneVmId returns a boolean if a field has been set.

### GetVmDiskIds

`func (o *DeleteStandAloneVmDiskCommand) GetVmDiskIds() []int32`

GetVmDiskIds returns the VmDiskIds field if non-nil, zero value otherwise.

### GetVmDiskIdsOk

`func (o *DeleteStandAloneVmDiskCommand) GetVmDiskIdsOk() (*[]int32, bool)`

GetVmDiskIdsOk returns a tuple with the VmDiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskIds

`func (o *DeleteStandAloneVmDiskCommand) SetVmDiskIds(v []int32)`

SetVmDiskIds sets VmDiskIds field to given value.

### HasVmDiskIds

`func (o *DeleteStandAloneVmDiskCommand) HasVmDiskIds() bool`

HasVmDiskIds returns a boolean if a field has been set.

### SetVmDiskIdsNil

`func (o *DeleteStandAloneVmDiskCommand) SetVmDiskIdsNil(b bool)`

 SetVmDiskIdsNil sets the value for VmDiskIds to be an explicit nil

### UnsetVmDiskIds
`func (o *DeleteStandAloneVmDiskCommand) UnsetVmDiskIds()`

UnsetVmDiskIds ensures that no value is present for VmDiskIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


