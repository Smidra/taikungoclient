/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the EnableAutoscalingCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableAutoscalingCommand{}

// EnableAutoscalingCommand struct for EnableAutoscalingCommand
type EnableAutoscalingCommand struct {
	Id                   *int32         `json:"id,omitempty"`
	AutoscalingGroupName NullableString `json:"autoscalingGroupName,omitempty"`
	MinSize              *int32         `json:"minSize,omitempty"`
	MaxSize              *int32         `json:"maxSize,omitempty"`
	DiskSize             *float64       `json:"diskSize,omitempty"`
	Flavor               NullableString `json:"flavor,omitempty"`
	SpotEnabled          *bool          `json:"spotEnabled,omitempty"`
}

// NewEnableAutoscalingCommand instantiates a new EnableAutoscalingCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableAutoscalingCommand() *EnableAutoscalingCommand {
	this := EnableAutoscalingCommand{}
	return &this
}

// NewEnableAutoscalingCommandWithDefaults instantiates a new EnableAutoscalingCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableAutoscalingCommandWithDefaults() *EnableAutoscalingCommand {
	this := EnableAutoscalingCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnableAutoscalingCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAutoscalingCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnableAutoscalingCommand) SetId(v int32) {
	o.Id = &v
}

// GetAutoscalingGroupName returns the AutoscalingGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnableAutoscalingCommand) GetAutoscalingGroupName() string {
	if o == nil || IsNil(o.AutoscalingGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.AutoscalingGroupName.Get()
}

// GetAutoscalingGroupNameOk returns a tuple with the AutoscalingGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnableAutoscalingCommand) GetAutoscalingGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoscalingGroupName.Get(), o.AutoscalingGroupName.IsSet()
}

// HasAutoscalingGroupName returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasAutoscalingGroupName() bool {
	if o != nil && o.AutoscalingGroupName.IsSet() {
		return true
	}

	return false
}

// SetAutoscalingGroupName gets a reference to the given NullableString and assigns it to the AutoscalingGroupName field.
func (o *EnableAutoscalingCommand) SetAutoscalingGroupName(v string) {
	o.AutoscalingGroupName.Set(&v)
}

// SetAutoscalingGroupNameNil sets the value for AutoscalingGroupName to be an explicit nil
func (o *EnableAutoscalingCommand) SetAutoscalingGroupNameNil() {
	o.AutoscalingGroupName.Set(nil)
}

// UnsetAutoscalingGroupName ensures that no value is present for AutoscalingGroupName, not even an explicit nil
func (o *EnableAutoscalingCommand) UnsetAutoscalingGroupName() {
	o.AutoscalingGroupName.Unset()
}

// GetMinSize returns the MinSize field value if set, zero value otherwise.
func (o *EnableAutoscalingCommand) GetMinSize() int32 {
	if o == nil || IsNil(o.MinSize) {
		var ret int32
		return ret
	}
	return *o.MinSize
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAutoscalingCommand) GetMinSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSize) {
		return nil, false
	}
	return o.MinSize, true
}

// HasMinSize returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasMinSize() bool {
	if o != nil && !IsNil(o.MinSize) {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given int32 and assigns it to the MinSize field.
func (o *EnableAutoscalingCommand) SetMinSize(v int32) {
	o.MinSize = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *EnableAutoscalingCommand) GetMaxSize() int32 {
	if o == nil || IsNil(o.MaxSize) {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAutoscalingCommand) GetMaxSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSize) {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasMaxSize() bool {
	if o != nil && !IsNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *EnableAutoscalingCommand) SetMaxSize(v int32) {
	o.MaxSize = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *EnableAutoscalingCommand) GetDiskSize() float64 {
	if o == nil || IsNil(o.DiskSize) {
		var ret float64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAutoscalingCommand) GetDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given float64 and assigns it to the DiskSize field.
func (o *EnableAutoscalingCommand) SetDiskSize(v float64) {
	o.DiskSize = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnableAutoscalingCommand) GetFlavor() string {
	if o == nil || IsNil(o.Flavor.Get()) {
		var ret string
		return ret
	}
	return *o.Flavor.Get()
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnableAutoscalingCommand) GetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flavor.Get(), o.Flavor.IsSet()
}

// HasFlavor returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasFlavor() bool {
	if o != nil && o.Flavor.IsSet() {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given NullableString and assigns it to the Flavor field.
func (o *EnableAutoscalingCommand) SetFlavor(v string) {
	o.Flavor.Set(&v)
}

// SetFlavorNil sets the value for Flavor to be an explicit nil
func (o *EnableAutoscalingCommand) SetFlavorNil() {
	o.Flavor.Set(nil)
}

// UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
func (o *EnableAutoscalingCommand) UnsetFlavor() {
	o.Flavor.Unset()
}

// GetSpotEnabled returns the SpotEnabled field value if set, zero value otherwise.
func (o *EnableAutoscalingCommand) GetSpotEnabled() bool {
	if o == nil || IsNil(o.SpotEnabled) {
		var ret bool
		return ret
	}
	return *o.SpotEnabled
}

// GetSpotEnabledOk returns a tuple with the SpotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAutoscalingCommand) GetSpotEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SpotEnabled) {
		return nil, false
	}
	return o.SpotEnabled, true
}

// HasSpotEnabled returns a boolean if a field has been set.
func (o *EnableAutoscalingCommand) HasSpotEnabled() bool {
	if o != nil && !IsNil(o.SpotEnabled) {
		return true
	}

	return false
}

// SetSpotEnabled gets a reference to the given bool and assigns it to the SpotEnabled field.
func (o *EnableAutoscalingCommand) SetSpotEnabled(v bool) {
	o.SpotEnabled = &v
}

func (o EnableAutoscalingCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableAutoscalingCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.AutoscalingGroupName.IsSet() {
		toSerialize["autoscalingGroupName"] = o.AutoscalingGroupName.Get()
	}
	if !IsNil(o.MinSize) {
		toSerialize["minSize"] = o.MinSize
	}
	if !IsNil(o.MaxSize) {
		toSerialize["maxSize"] = o.MaxSize
	}
	if !IsNil(o.DiskSize) {
		toSerialize["diskSize"] = o.DiskSize
	}
	if o.Flavor.IsSet() {
		toSerialize["flavor"] = o.Flavor.Get()
	}
	if !IsNil(o.SpotEnabled) {
		toSerialize["spotEnabled"] = o.SpotEnabled
	}
	return toSerialize, nil
}

type NullableEnableAutoscalingCommand struct {
	value *EnableAutoscalingCommand
	isSet bool
}

func (v NullableEnableAutoscalingCommand) Get() *EnableAutoscalingCommand {
	return v.value
}

func (v *NullableEnableAutoscalingCommand) Set(val *EnableAutoscalingCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableAutoscalingCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableAutoscalingCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableAutoscalingCommand(val *EnableAutoscalingCommand) *NullableEnableAutoscalingCommand {
	return &NullableEnableAutoscalingCommand{value: val, isSet: true}
}

func (v NullableEnableAutoscalingCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableAutoscalingCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
