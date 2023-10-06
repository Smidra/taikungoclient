# CreateProxmoxNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**NetMask** | Pointer to **int32** |  | [optional] 
**BeginAllocationRange** | Pointer to **NullableString** |  | [optional] 
**EndAllocationRange** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateProxmoxNetworkDto

`func NewCreateProxmoxNetworkDto() *CreateProxmoxNetworkDto`

NewCreateProxmoxNetworkDto instantiates a new CreateProxmoxNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxmoxNetworkDtoWithDefaults

`func NewCreateProxmoxNetworkDtoWithDefaults() *CreateProxmoxNetworkDto`

NewCreateProxmoxNetworkDtoWithDefaults instantiates a new CreateProxmoxNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *CreateProxmoxNetworkDto) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *CreateProxmoxNetworkDto) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *CreateProxmoxNetworkDto) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *CreateProxmoxNetworkDto) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *CreateProxmoxNetworkDto) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *CreateProxmoxNetworkDto) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetGateway

`func (o *CreateProxmoxNetworkDto) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateProxmoxNetworkDto) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateProxmoxNetworkDto) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateProxmoxNetworkDto) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *CreateProxmoxNetworkDto) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *CreateProxmoxNetworkDto) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetIpAddress

`func (o *CreateProxmoxNetworkDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateProxmoxNetworkDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateProxmoxNetworkDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateProxmoxNetworkDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *CreateProxmoxNetworkDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *CreateProxmoxNetworkDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetMask

`func (o *CreateProxmoxNetworkDto) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *CreateProxmoxNetworkDto) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *CreateProxmoxNetworkDto) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *CreateProxmoxNetworkDto) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### GetBeginAllocationRange

`func (o *CreateProxmoxNetworkDto) GetBeginAllocationRange() string`

GetBeginAllocationRange returns the BeginAllocationRange field if non-nil, zero value otherwise.

### GetBeginAllocationRangeOk

`func (o *CreateProxmoxNetworkDto) GetBeginAllocationRangeOk() (*string, bool)`

GetBeginAllocationRangeOk returns a tuple with the BeginAllocationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginAllocationRange

`func (o *CreateProxmoxNetworkDto) SetBeginAllocationRange(v string)`

SetBeginAllocationRange sets BeginAllocationRange field to given value.

### HasBeginAllocationRange

`func (o *CreateProxmoxNetworkDto) HasBeginAllocationRange() bool`

HasBeginAllocationRange returns a boolean if a field has been set.

### SetBeginAllocationRangeNil

`func (o *CreateProxmoxNetworkDto) SetBeginAllocationRangeNil(b bool)`

 SetBeginAllocationRangeNil sets the value for BeginAllocationRange to be an explicit nil

### UnsetBeginAllocationRange
`func (o *CreateProxmoxNetworkDto) UnsetBeginAllocationRange()`

UnsetBeginAllocationRange ensures that no value is present for BeginAllocationRange, not even an explicit nil
### GetEndAllocationRange

`func (o *CreateProxmoxNetworkDto) GetEndAllocationRange() string`

GetEndAllocationRange returns the EndAllocationRange field if non-nil, zero value otherwise.

### GetEndAllocationRangeOk

`func (o *CreateProxmoxNetworkDto) GetEndAllocationRangeOk() (*string, bool)`

GetEndAllocationRangeOk returns a tuple with the EndAllocationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAllocationRange

`func (o *CreateProxmoxNetworkDto) SetEndAllocationRange(v string)`

SetEndAllocationRange sets EndAllocationRange field to given value.

### HasEndAllocationRange

`func (o *CreateProxmoxNetworkDto) HasEndAllocationRange() bool`

HasEndAllocationRange returns a boolean if a field has been set.

### SetEndAllocationRangeNil

`func (o *CreateProxmoxNetworkDto) SetEndAllocationRangeNil(b bool)`

 SetEndAllocationRangeNil sets the value for EndAllocationRange to be an explicit nil

### UnsetEndAllocationRange
`func (o *CreateProxmoxNetworkDto) UnsetEndAllocationRange()`

UnsetEndAllocationRange ensures that no value is present for EndAllocationRange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


