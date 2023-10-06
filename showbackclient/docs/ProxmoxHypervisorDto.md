# ProxmoxHypervisorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**IsBound** | Pointer to **bool** |  | [optional] 
**UsedByServer** | Pointer to **bool** |  | [optional] 

## Methods

### NewProxmoxHypervisorDto

`func NewProxmoxHypervisorDto() *ProxmoxHypervisorDto`

NewProxmoxHypervisorDto instantiates a new ProxmoxHypervisorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxHypervisorDtoWithDefaults

`func NewProxmoxHypervisorDtoWithDefaults() *ProxmoxHypervisorDto`

NewProxmoxHypervisorDtoWithDefaults instantiates a new ProxmoxHypervisorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProxmoxHypervisorDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxmoxHypervisorDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxmoxHypervisorDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProxmoxHypervisorDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProxmoxHypervisorDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProxmoxHypervisorDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsBound

`func (o *ProxmoxHypervisorDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *ProxmoxHypervisorDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *ProxmoxHypervisorDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *ProxmoxHypervisorDto) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.

### GetUsedByServer

`func (o *ProxmoxHypervisorDto) GetUsedByServer() bool`

GetUsedByServer returns the UsedByServer field if non-nil, zero value otherwise.

### GetUsedByServerOk

`func (o *ProxmoxHypervisorDto) GetUsedByServerOk() (*bool, bool)`

GetUsedByServerOk returns a tuple with the UsedByServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByServer

`func (o *ProxmoxHypervisorDto) SetUsedByServer(v bool)`

SetUsedByServer sets UsedByServer field to given value.

### HasUsedByServer

`func (o *ProxmoxHypervisorDto) HasUsedByServer() bool`

HasUsedByServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


