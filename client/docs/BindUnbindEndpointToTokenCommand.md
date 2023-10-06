# BindUnbindEndpointToTokenCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **NullableString** |  | [optional] 
**Endpoints** | Pointer to [**[]AvailableEndpointData**](AvailableEndpointData.md) |  | [optional] 
**BindAll** | Pointer to **bool** |  | [optional] 

## Methods

### NewBindUnbindEndpointToTokenCommand

`func NewBindUnbindEndpointToTokenCommand() *BindUnbindEndpointToTokenCommand`

NewBindUnbindEndpointToTokenCommand instantiates a new BindUnbindEndpointToTokenCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindUnbindEndpointToTokenCommandWithDefaults

`func NewBindUnbindEndpointToTokenCommandWithDefaults() *BindUnbindEndpointToTokenCommand`

NewBindUnbindEndpointToTokenCommandWithDefaults instantiates a new BindUnbindEndpointToTokenCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *BindUnbindEndpointToTokenCommand) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BindUnbindEndpointToTokenCommand) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BindUnbindEndpointToTokenCommand) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *BindUnbindEndpointToTokenCommand) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *BindUnbindEndpointToTokenCommand) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *BindUnbindEndpointToTokenCommand) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetEndpoints

`func (o *BindUnbindEndpointToTokenCommand) GetEndpoints() []AvailableEndpointData`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *BindUnbindEndpointToTokenCommand) GetEndpointsOk() (*[]AvailableEndpointData, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *BindUnbindEndpointToTokenCommand) SetEndpoints(v []AvailableEndpointData)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *BindUnbindEndpointToTokenCommand) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *BindUnbindEndpointToTokenCommand) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *BindUnbindEndpointToTokenCommand) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetBindAll

`func (o *BindUnbindEndpointToTokenCommand) GetBindAll() bool`

GetBindAll returns the BindAll field if non-nil, zero value otherwise.

### GetBindAllOk

`func (o *BindUnbindEndpointToTokenCommand) GetBindAllOk() (*bool, bool)`

GetBindAllOk returns a tuple with the BindAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAll

`func (o *BindUnbindEndpointToTokenCommand) SetBindAll(v bool)`

SetBindAll sets BindAll field to given value.

### HasBindAll

`func (o *BindUnbindEndpointToTokenCommand) HasBindAll() bool`

HasBindAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


