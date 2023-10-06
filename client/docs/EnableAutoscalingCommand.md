# EnableAutoscalingCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AutoscalingGroupName** | Pointer to **NullableString** |  | [optional] 
**MinSize** | Pointer to **int32** |  | [optional] 
**MaxSize** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **float64** |  | [optional] 
**Flavor** | Pointer to **NullableString** |  | [optional] 
**SpotEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewEnableAutoscalingCommand

`func NewEnableAutoscalingCommand() *EnableAutoscalingCommand`

NewEnableAutoscalingCommand instantiates a new EnableAutoscalingCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableAutoscalingCommandWithDefaults

`func NewEnableAutoscalingCommandWithDefaults() *EnableAutoscalingCommand`

NewEnableAutoscalingCommandWithDefaults instantiates a new EnableAutoscalingCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnableAutoscalingCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnableAutoscalingCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnableAutoscalingCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnableAutoscalingCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoscalingGroupName

`func (o *EnableAutoscalingCommand) GetAutoscalingGroupName() string`

GetAutoscalingGroupName returns the AutoscalingGroupName field if non-nil, zero value otherwise.

### GetAutoscalingGroupNameOk

`func (o *EnableAutoscalingCommand) GetAutoscalingGroupNameOk() (*string, bool)`

GetAutoscalingGroupNameOk returns a tuple with the AutoscalingGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroupName

`func (o *EnableAutoscalingCommand) SetAutoscalingGroupName(v string)`

SetAutoscalingGroupName sets AutoscalingGroupName field to given value.

### HasAutoscalingGroupName

`func (o *EnableAutoscalingCommand) HasAutoscalingGroupName() bool`

HasAutoscalingGroupName returns a boolean if a field has been set.

### SetAutoscalingGroupNameNil

`func (o *EnableAutoscalingCommand) SetAutoscalingGroupNameNil(b bool)`

 SetAutoscalingGroupNameNil sets the value for AutoscalingGroupName to be an explicit nil

### UnsetAutoscalingGroupName
`func (o *EnableAutoscalingCommand) UnsetAutoscalingGroupName()`

UnsetAutoscalingGroupName ensures that no value is present for AutoscalingGroupName, not even an explicit nil
### GetMinSize

`func (o *EnableAutoscalingCommand) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *EnableAutoscalingCommand) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *EnableAutoscalingCommand) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *EnableAutoscalingCommand) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetMaxSize

`func (o *EnableAutoscalingCommand) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *EnableAutoscalingCommand) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *EnableAutoscalingCommand) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *EnableAutoscalingCommand) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetDiskSize

`func (o *EnableAutoscalingCommand) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *EnableAutoscalingCommand) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *EnableAutoscalingCommand) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *EnableAutoscalingCommand) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetFlavor

`func (o *EnableAutoscalingCommand) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *EnableAutoscalingCommand) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *EnableAutoscalingCommand) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *EnableAutoscalingCommand) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *EnableAutoscalingCommand) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *EnableAutoscalingCommand) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetSpotEnabled

`func (o *EnableAutoscalingCommand) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *EnableAutoscalingCommand) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *EnableAutoscalingCommand) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.

### HasSpotEnabled

`func (o *EnableAutoscalingCommand) HasSpotEnabled() bool`

HasSpotEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


