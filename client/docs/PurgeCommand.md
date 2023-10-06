# PurgeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ServerIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPurgeCommand

`func NewPurgeCommand() *PurgeCommand`

NewPurgeCommand instantiates a new PurgeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeCommandWithDefaults

`func NewPurgeCommandWithDefaults() *PurgeCommand`

NewPurgeCommandWithDefaults instantiates a new PurgeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PurgeCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PurgeCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PurgeCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PurgeCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServerIds

`func (o *PurgeCommand) GetServerIds() []int32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *PurgeCommand) GetServerIdsOk() (*[]int32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *PurgeCommand) SetServerIds(v []int32)`

SetServerIds sets ServerIds field to given value.

### HasServerIds

`func (o *PurgeCommand) HasServerIds() bool`

HasServerIds returns a boolean if a field has been set.

### SetServerIdsNil

`func (o *PurgeCommand) SetServerIdsNil(b bool)`

 SetServerIdsNil sets the value for ServerIds to be an explicit nil

### UnsetServerIds
`func (o *PurgeCommand) UnsetServerIds()`

UnsetServerIds ensures that no value is present for ServerIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


