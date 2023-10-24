# UpdateServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]ServerUpdateDto**](ServerUpdateDto.md) |  | [optional] 

## Methods

### NewUpdateServerCommand

`func NewUpdateServerCommand() *UpdateServerCommand`

NewUpdateServerCommand instantiates a new UpdateServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerCommandWithDefaults

`func NewUpdateServerCommandWithDefaults() *UpdateServerCommand`

NewUpdateServerCommandWithDefaults instantiates a new UpdateServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *UpdateServerCommand) GetServers() []ServerUpdateDto`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateServerCommand) GetServersOk() (*[]ServerUpdateDto, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateServerCommand) SetServers(v []ServerUpdateDto)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateServerCommand) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *UpdateServerCommand) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *UpdateServerCommand) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


