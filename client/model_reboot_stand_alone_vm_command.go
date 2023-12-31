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

// checks if the RebootStandAloneVmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RebootStandAloneVmCommand{}

// RebootStandAloneVmCommand struct for RebootStandAloneVmCommand
type RebootStandAloneVmCommand struct {
	Id   *int32         `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
}

// NewRebootStandAloneVmCommand instantiates a new RebootStandAloneVmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootStandAloneVmCommand() *RebootStandAloneVmCommand {
	this := RebootStandAloneVmCommand{}
	return &this
}

// NewRebootStandAloneVmCommandWithDefaults instantiates a new RebootStandAloneVmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootStandAloneVmCommandWithDefaults() *RebootStandAloneVmCommand {
	this := RebootStandAloneVmCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RebootStandAloneVmCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootStandAloneVmCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RebootStandAloneVmCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RebootStandAloneVmCommand) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RebootStandAloneVmCommand) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RebootStandAloneVmCommand) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *RebootStandAloneVmCommand) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *RebootStandAloneVmCommand) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *RebootStandAloneVmCommand) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *RebootStandAloneVmCommand) UnsetType() {
	o.Type.Unset()
}

func (o RebootStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RebootStandAloneVmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableRebootStandAloneVmCommand struct {
	value *RebootStandAloneVmCommand
	isSet bool
}

func (v NullableRebootStandAloneVmCommand) Get() *RebootStandAloneVmCommand {
	return v.value
}

func (v *NullableRebootStandAloneVmCommand) Set(val *RebootStandAloneVmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootStandAloneVmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootStandAloneVmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootStandAloneVmCommand(val *RebootStandAloneVmCommand) *NullableRebootStandAloneVmCommand {
	return &NullableRebootStandAloneVmCommand{value: val, isSet: true}
}

func (v NullableRebootStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootStandAloneVmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
