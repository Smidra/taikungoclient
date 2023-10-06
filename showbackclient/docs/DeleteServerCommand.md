# DeleteServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ServerIds** | Pointer to **[]int32** |  | [optional] 
**DeleteAutoscalingServers** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeleteServerCommand

`func NewDeleteServerCommand() *DeleteServerCommand`

NewDeleteServerCommand instantiates a new DeleteServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServerCommandWithDefaults

`func NewDeleteServerCommandWithDefaults() *DeleteServerCommand`

NewDeleteServerCommandWithDefaults instantiates a new DeleteServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteServerCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteServerCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteServerCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DeleteServerCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServerIds

`func (o *DeleteServerCommand) GetServerIds() []int32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *DeleteServerCommand) GetServerIdsOk() (*[]int32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *DeleteServerCommand) SetServerIds(v []int32)`

SetServerIds sets ServerIds field to given value.

### HasServerIds

`func (o *DeleteServerCommand) HasServerIds() bool`

HasServerIds returns a boolean if a field has been set.

### SetServerIdsNil

`func (o *DeleteServerCommand) SetServerIdsNil(b bool)`

 SetServerIdsNil sets the value for ServerIds to be an explicit nil

### UnsetServerIds
`func (o *DeleteServerCommand) UnsetServerIds()`

UnsetServerIds ensures that no value is present for ServerIds, not even an explicit nil
### GetDeleteAutoscalingServers

`func (o *DeleteServerCommand) GetDeleteAutoscalingServers() bool`

GetDeleteAutoscalingServers returns the DeleteAutoscalingServers field if non-nil, zero value otherwise.

### GetDeleteAutoscalingServersOk

`func (o *DeleteServerCommand) GetDeleteAutoscalingServersOk() (*bool, bool)`

GetDeleteAutoscalingServersOk returns a tuple with the DeleteAutoscalingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAutoscalingServers

`func (o *DeleteServerCommand) SetDeleteAutoscalingServers(v bool)`

SetDeleteAutoscalingServers sets DeleteAutoscalingServers field to given value.

### HasDeleteAutoscalingServers

`func (o *DeleteServerCommand) HasDeleteAutoscalingServers() bool`

HasDeleteAutoscalingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


