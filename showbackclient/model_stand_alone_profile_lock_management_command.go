/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the StandAloneProfileLockManagementCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileLockManagementCommand{}

// StandAloneProfileLockManagementCommand struct for StandAloneProfileLockManagementCommand
type StandAloneProfileLockManagementCommand struct {
	Id   *int32         `json:"id,omitempty"`
	Mode NullableString `json:"mode,omitempty"`
}

// NewStandAloneProfileLockManagementCommand instantiates a new StandAloneProfileLockManagementCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileLockManagementCommand() *StandAloneProfileLockManagementCommand {
	this := StandAloneProfileLockManagementCommand{}
	return &this
}

// NewStandAloneProfileLockManagementCommandWithDefaults instantiates a new StandAloneProfileLockManagementCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileLockManagementCommandWithDefaults() *StandAloneProfileLockManagementCommand {
	this := StandAloneProfileLockManagementCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneProfileLockManagementCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileLockManagementCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneProfileLockManagementCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneProfileLockManagementCommand) SetId(v int32) {
	o.Id = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileLockManagementCommand) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileLockManagementCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *StandAloneProfileLockManagementCommand) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *StandAloneProfileLockManagementCommand) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *StandAloneProfileLockManagementCommand) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *StandAloneProfileLockManagementCommand) UnsetMode() {
	o.Mode.Unset()
}

func (o StandAloneProfileLockManagementCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileLockManagementCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneProfileLockManagementCommand struct {
	value *StandAloneProfileLockManagementCommand
	isSet bool
}

func (v NullableStandAloneProfileLockManagementCommand) Get() *StandAloneProfileLockManagementCommand {
	return v.value
}

func (v *NullableStandAloneProfileLockManagementCommand) Set(val *StandAloneProfileLockManagementCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileLockManagementCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileLockManagementCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileLockManagementCommand(val *StandAloneProfileLockManagementCommand) *NullableStandAloneProfileLockManagementCommand {
	return &NullableStandAloneProfileLockManagementCommand{value: val, isSet: true}
}

func (v NullableStandAloneProfileLockManagementCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileLockManagementCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
