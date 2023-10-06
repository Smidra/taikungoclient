# ProxmoxNetworkListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**NetMask** | Pointer to **int32** |  | [optional] 
**BeginAllocationRange** | Pointer to **NullableString** |  | [optional] 
**EndAllocationRange** | Pointer to **NullableString** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 

## Methods

### NewProxmoxNetworkListDto

`func NewProxmoxNetworkListDto() *ProxmoxNetworkListDto`

NewProxmoxNetworkListDto instantiates a new ProxmoxNetworkListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxNetworkListDtoWithDefaults

`func NewProxmoxNetworkListDtoWithDefaults() *ProxmoxNetworkListDto`

NewProxmoxNetworkListDtoWithDefaults instantiates a new ProxmoxNetworkListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *ProxmoxNetworkListDto) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *ProxmoxNetworkListDto) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *ProxmoxNetworkListDto) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *ProxmoxNetworkListDto) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *ProxmoxNetworkListDto) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *ProxmoxNetworkListDto) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetGateway

`func (o *ProxmoxNetworkListDto) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ProxmoxNetworkListDto) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ProxmoxNetworkListDto) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ProxmoxNetworkListDto) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *ProxmoxNetworkListDto) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ProxmoxNetworkListDto) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetIpAddress

`func (o *ProxmoxNetworkListDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ProxmoxNetworkListDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ProxmoxNetworkListDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ProxmoxNetworkListDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ProxmoxNetworkListDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ProxmoxNetworkListDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetMask

`func (o *ProxmoxNetworkListDto) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *ProxmoxNetworkListDto) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *ProxmoxNetworkListDto) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *ProxmoxNetworkListDto) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### GetBeginAllocationRange

`func (o *ProxmoxNetworkListDto) GetBeginAllocationRange() string`

GetBeginAllocationRange returns the BeginAllocationRange field if non-nil, zero value otherwise.

### GetBeginAllocationRangeOk

`func (o *ProxmoxNetworkListDto) GetBeginAllocationRangeOk() (*string, bool)`

GetBeginAllocationRangeOk returns a tuple with the BeginAllocationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginAllocationRange

`func (o *ProxmoxNetworkListDto) SetBeginAllocationRange(v string)`

SetBeginAllocationRange sets BeginAllocationRange field to given value.

### HasBeginAllocationRange

`func (o *ProxmoxNetworkListDto) HasBeginAllocationRange() bool`

HasBeginAllocationRange returns a boolean if a field has been set.

### SetBeginAllocationRangeNil

`func (o *ProxmoxNetworkListDto) SetBeginAllocationRangeNil(b bool)`

 SetBeginAllocationRangeNil sets the value for BeginAllocationRange to be an explicit nil

### UnsetBeginAllocationRange
`func (o *ProxmoxNetworkListDto) UnsetBeginAllocationRange()`

UnsetBeginAllocationRange ensures that no value is present for BeginAllocationRange, not even an explicit nil
### GetEndAllocationRange

`func (o *ProxmoxNetworkListDto) GetEndAllocationRange() string`

GetEndAllocationRange returns the EndAllocationRange field if non-nil, zero value otherwise.

### GetEndAllocationRangeOk

`func (o *ProxmoxNetworkListDto) GetEndAllocationRangeOk() (*string, bool)`

GetEndAllocationRangeOk returns a tuple with the EndAllocationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAllocationRange

`func (o *ProxmoxNetworkListDto) SetEndAllocationRange(v string)`

SetEndAllocationRange sets EndAllocationRange field to given value.

### HasEndAllocationRange

`func (o *ProxmoxNetworkListDto) HasEndAllocationRange() bool`

HasEndAllocationRange returns a boolean if a field has been set.

### SetEndAllocationRangeNil

`func (o *ProxmoxNetworkListDto) SetEndAllocationRangeNil(b bool)`

 SetEndAllocationRangeNil sets the value for EndAllocationRange to be an explicit nil

### UnsetEndAllocationRange
`func (o *ProxmoxNetworkListDto) UnsetEndAllocationRange()`

UnsetEndAllocationRange ensures that no value is present for EndAllocationRange, not even an explicit nil
### GetIsPrivate

`func (o *ProxmoxNetworkListDto) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ProxmoxNetworkListDto) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ProxmoxNetworkListDto) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ProxmoxNetworkListDto) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


