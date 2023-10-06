# GoogleFlavorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Ram** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **interface{}** |  | [optional] 
**LinuxPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsPrice** | Pointer to **NullableString** |  | [optional] 
**LinuxSpotPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsSpotPrice** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoogleFlavorDto

`func NewGoogleFlavorDto() *GoogleFlavorDto`

NewGoogleFlavorDto instantiates a new GoogleFlavorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleFlavorDtoWithDefaults

`func NewGoogleFlavorDtoWithDefaults() *GoogleFlavorDto`

NewGoogleFlavorDtoWithDefaults instantiates a new GoogleFlavorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GoogleFlavorDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleFlavorDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleFlavorDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoogleFlavorDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GoogleFlavorDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GoogleFlavorDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *GoogleFlavorDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GoogleFlavorDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GoogleFlavorDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GoogleFlavorDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *GoogleFlavorDto) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *GoogleFlavorDto) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetRam

`func (o *GoogleFlavorDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *GoogleFlavorDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *GoogleFlavorDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *GoogleFlavorDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### SetRamNil

`func (o *GoogleFlavorDto) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *GoogleFlavorDto) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetDescription

`func (o *GoogleFlavorDto) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoogleFlavorDto) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoogleFlavorDto) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoogleFlavorDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GoogleFlavorDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GoogleFlavorDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinuxPrice

`func (o *GoogleFlavorDto) GetLinuxPrice() string`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *GoogleFlavorDto) GetLinuxPriceOk() (*string, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *GoogleFlavorDto) SetLinuxPrice(v string)`

SetLinuxPrice sets LinuxPrice field to given value.

### HasLinuxPrice

`func (o *GoogleFlavorDto) HasLinuxPrice() bool`

HasLinuxPrice returns a boolean if a field has been set.

### SetLinuxPriceNil

`func (o *GoogleFlavorDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *GoogleFlavorDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsPrice

`func (o *GoogleFlavorDto) GetWindowsPrice() string`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *GoogleFlavorDto) GetWindowsPriceOk() (*string, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *GoogleFlavorDto) SetWindowsPrice(v string)`

SetWindowsPrice sets WindowsPrice field to given value.

### HasWindowsPrice

`func (o *GoogleFlavorDto) HasWindowsPrice() bool`

HasWindowsPrice returns a boolean if a field has been set.

### SetWindowsPriceNil

`func (o *GoogleFlavorDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *GoogleFlavorDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
### GetLinuxSpotPrice

`func (o *GoogleFlavorDto) GetLinuxSpotPrice() string`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *GoogleFlavorDto) GetLinuxSpotPriceOk() (*string, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *GoogleFlavorDto) SetLinuxSpotPrice(v string)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.

### HasLinuxSpotPrice

`func (o *GoogleFlavorDto) HasLinuxSpotPrice() bool`

HasLinuxSpotPrice returns a boolean if a field has been set.

### SetLinuxSpotPriceNil

`func (o *GoogleFlavorDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *GoogleFlavorDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *GoogleFlavorDto) GetWindowsSpotPrice() string`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *GoogleFlavorDto) GetWindowsSpotPriceOk() (*string, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *GoogleFlavorDto) SetWindowsSpotPrice(v string)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.

### HasWindowsSpotPrice

`func (o *GoogleFlavorDto) HasWindowsSpotPrice() bool`

HasWindowsSpotPrice returns a boolean if a field has been set.

### SetWindowsSpotPriceNil

`func (o *GoogleFlavorDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *GoogleFlavorDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


