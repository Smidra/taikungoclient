# ResetStandAloneVmDiskStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandAloneVmId** | Pointer to **int32** |  | [optional] 
**DiskIds** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to [**StandAloneVmDiskStatus**](StandAloneVmDiskStatus.md) |  | [optional] 

## Methods

### NewResetStandAloneVmDiskStatusCommand

`func NewResetStandAloneVmDiskStatusCommand() *ResetStandAloneVmDiskStatusCommand`

NewResetStandAloneVmDiskStatusCommand instantiates a new ResetStandAloneVmDiskStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetStandAloneVmDiskStatusCommandWithDefaults

`func NewResetStandAloneVmDiskStatusCommandWithDefaults() *ResetStandAloneVmDiskStatusCommand`

NewResetStandAloneVmDiskStatusCommandWithDefaults instantiates a new ResetStandAloneVmDiskStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandAloneVmId

`func (o *ResetStandAloneVmDiskStatusCommand) GetStandAloneVmId() int32`

GetStandAloneVmId returns the StandAloneVmId field if non-nil, zero value otherwise.

### GetStandAloneVmIdOk

`func (o *ResetStandAloneVmDiskStatusCommand) GetStandAloneVmIdOk() (*int32, bool)`

GetStandAloneVmIdOk returns a tuple with the StandAloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneVmId

`func (o *ResetStandAloneVmDiskStatusCommand) SetStandAloneVmId(v int32)`

SetStandAloneVmId sets StandAloneVmId field to given value.

### HasStandAloneVmId

`func (o *ResetStandAloneVmDiskStatusCommand) HasStandAloneVmId() bool`

HasStandAloneVmId returns a boolean if a field has been set.

### GetDiskIds

`func (o *ResetStandAloneVmDiskStatusCommand) GetDiskIds() []int32`

GetDiskIds returns the DiskIds field if non-nil, zero value otherwise.

### GetDiskIdsOk

`func (o *ResetStandAloneVmDiskStatusCommand) GetDiskIdsOk() (*[]int32, bool)`

GetDiskIdsOk returns a tuple with the DiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIds

`func (o *ResetStandAloneVmDiskStatusCommand) SetDiskIds(v []int32)`

SetDiskIds sets DiskIds field to given value.

### HasDiskIds

`func (o *ResetStandAloneVmDiskStatusCommand) HasDiskIds() bool`

HasDiskIds returns a boolean if a field has been set.

### SetDiskIdsNil

`func (o *ResetStandAloneVmDiskStatusCommand) SetDiskIdsNil(b bool)`

 SetDiskIdsNil sets the value for DiskIds to be an explicit nil

### UnsetDiskIds
`func (o *ResetStandAloneVmDiskStatusCommand) UnsetDiskIds()`

UnsetDiskIds ensures that no value is present for DiskIds, not even an explicit nil
### GetStatus

`func (o *ResetStandAloneVmDiskStatusCommand) GetStatus() StandAloneVmDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResetStandAloneVmDiskStatusCommand) GetStatusOk() (*StandAloneVmDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResetStandAloneVmDiskStatusCommand) SetStatus(v StandAloneVmDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResetStandAloneVmDiskStatusCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


