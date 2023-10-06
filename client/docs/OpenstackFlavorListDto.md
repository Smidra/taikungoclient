# OpenstackFlavorListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ram** | Pointer to **int64** |  | [optional] 
**Cpu** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenstackFlavorListDto

`func NewOpenstackFlavorListDto() *OpenstackFlavorListDto`

NewOpenstackFlavorListDto instantiates a new OpenstackFlavorListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackFlavorListDtoWithDefaults

`func NewOpenstackFlavorListDtoWithDefaults() *OpenstackFlavorListDto`

NewOpenstackFlavorListDtoWithDefaults instantiates a new OpenstackFlavorListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRam

`func (o *OpenstackFlavorListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *OpenstackFlavorListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *OpenstackFlavorListDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *OpenstackFlavorListDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCpu

`func (o *OpenstackFlavorListDto) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OpenstackFlavorListDto) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OpenstackFlavorListDto) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OpenstackFlavorListDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetName

`func (o *OpenstackFlavorListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenstackFlavorListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenstackFlavorListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenstackFlavorListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenstackFlavorListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenstackFlavorListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *OpenstackFlavorListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenstackFlavorListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenstackFlavorListDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenstackFlavorListDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OpenstackFlavorListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenstackFlavorListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


