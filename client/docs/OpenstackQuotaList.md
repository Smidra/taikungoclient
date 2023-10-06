# OpenstackQuotaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**OpenstackComputeQuotaDto**](OpenstackComputeQuotaDto.md) |  | [optional] 
**Volume** | Pointer to [**OpenstackVolumeQuotaDto**](OpenstackVolumeQuotaDto.md) |  | [optional] 
**Network** | Pointer to [**OpenstackNetworkDto**](OpenstackNetworkDto.md) |  | [optional] 

## Methods

### NewOpenstackQuotaList

`func NewOpenstackQuotaList() *OpenstackQuotaList`

NewOpenstackQuotaList instantiates a new OpenstackQuotaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackQuotaListWithDefaults

`func NewOpenstackQuotaListWithDefaults() *OpenstackQuotaList`

NewOpenstackQuotaListWithDefaults instantiates a new OpenstackQuotaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *OpenstackQuotaList) GetCompute() OpenstackComputeQuotaDto`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *OpenstackQuotaList) GetComputeOk() (*OpenstackComputeQuotaDto, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *OpenstackQuotaList) SetCompute(v OpenstackComputeQuotaDto)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *OpenstackQuotaList) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetVolume

`func (o *OpenstackQuotaList) GetVolume() OpenstackVolumeQuotaDto`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *OpenstackQuotaList) GetVolumeOk() (*OpenstackVolumeQuotaDto, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *OpenstackQuotaList) SetVolume(v OpenstackVolumeQuotaDto)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *OpenstackQuotaList) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetNetwork

`func (o *OpenstackQuotaList) GetNetwork() OpenstackNetworkDto`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OpenstackQuotaList) GetNetworkOk() (*OpenstackNetworkDto, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OpenstackQuotaList) SetNetwork(v OpenstackNetworkDto)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OpenstackQuotaList) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


