# ProxmoxFlavorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 

## Methods

### NewProxmoxFlavorData

`func NewProxmoxFlavorData() *ProxmoxFlavorData`

NewProxmoxFlavorData instantiates a new ProxmoxFlavorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxFlavorDataWithDefaults

`func NewProxmoxFlavorDataWithDefaults() *ProxmoxFlavorData`

NewProxmoxFlavorDataWithDefaults instantiates a new ProxmoxFlavorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProxmoxFlavorData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxmoxFlavorData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxmoxFlavorData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProxmoxFlavorData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProxmoxFlavorData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProxmoxFlavorData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *ProxmoxFlavorData) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProxmoxFlavorData) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProxmoxFlavorData) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProxmoxFlavorData) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *ProxmoxFlavorData) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ProxmoxFlavorData) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ProxmoxFlavorData) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ProxmoxFlavorData) HasRam() bool`

HasRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


