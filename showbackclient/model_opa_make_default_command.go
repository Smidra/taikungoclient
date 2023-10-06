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

// checks if the OpaMakeDefaultCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpaMakeDefaultCommand{}

// OpaMakeDefaultCommand struct for OpaMakeDefaultCommand
type OpaMakeDefaultCommand struct {
	Id *int32 `json:"id,omitempty"`
}

// NewOpaMakeDefaultCommand instantiates a new OpaMakeDefaultCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaMakeDefaultCommand() *OpaMakeDefaultCommand {
	this := OpaMakeDefaultCommand{}
	return &this
}

// NewOpaMakeDefaultCommandWithDefaults instantiates a new OpaMakeDefaultCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaMakeDefaultCommandWithDefaults() *OpaMakeDefaultCommand {
	this := OpaMakeDefaultCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpaMakeDefaultCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaMakeDefaultCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpaMakeDefaultCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpaMakeDefaultCommand) SetId(v int32) {
	o.Id = &v
}

func (o OpaMakeDefaultCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpaMakeDefaultCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOpaMakeDefaultCommand struct {
	value *OpaMakeDefaultCommand
	isSet bool
}

func (v NullableOpaMakeDefaultCommand) Get() *OpaMakeDefaultCommand {
	return v.value
}

func (v *NullableOpaMakeDefaultCommand) Set(val *OpaMakeDefaultCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaMakeDefaultCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaMakeDefaultCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaMakeDefaultCommand(val *OpaMakeDefaultCommand) *NullableOpaMakeDefaultCommand {
	return &NullableOpaMakeDefaultCommand{value: val, isSet: true}
}

func (v NullableOpaMakeDefaultCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaMakeDefaultCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}