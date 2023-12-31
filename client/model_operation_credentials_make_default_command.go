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

// checks if the OperationCredentialsMakeDefaultCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationCredentialsMakeDefaultCommand{}

// OperationCredentialsMakeDefaultCommand struct for OperationCredentialsMakeDefaultCommand
type OperationCredentialsMakeDefaultCommand struct {
	Id *int32 `json:"id,omitempty"`
}

// NewOperationCredentialsMakeDefaultCommand instantiates a new OperationCredentialsMakeDefaultCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationCredentialsMakeDefaultCommand() *OperationCredentialsMakeDefaultCommand {
	this := OperationCredentialsMakeDefaultCommand{}
	return &this
}

// NewOperationCredentialsMakeDefaultCommandWithDefaults instantiates a new OperationCredentialsMakeDefaultCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationCredentialsMakeDefaultCommandWithDefaults() *OperationCredentialsMakeDefaultCommand {
	this := OperationCredentialsMakeDefaultCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperationCredentialsMakeDefaultCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCredentialsMakeDefaultCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperationCredentialsMakeDefaultCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OperationCredentialsMakeDefaultCommand) SetId(v int32) {
	o.Id = &v
}

func (o OperationCredentialsMakeDefaultCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationCredentialsMakeDefaultCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOperationCredentialsMakeDefaultCommand struct {
	value *OperationCredentialsMakeDefaultCommand
	isSet bool
}

func (v NullableOperationCredentialsMakeDefaultCommand) Get() *OperationCredentialsMakeDefaultCommand {
	return v.value
}

func (v *NullableOperationCredentialsMakeDefaultCommand) Set(val *OperationCredentialsMakeDefaultCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationCredentialsMakeDefaultCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationCredentialsMakeDefaultCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationCredentialsMakeDefaultCommand(val *OperationCredentialsMakeDefaultCommand) *NullableOperationCredentialsMakeDefaultCommand {
	return &NullableOperationCredentialsMakeDefaultCommand{value: val, isSet: true}
}

func (v NullableOperationCredentialsMakeDefaultCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationCredentialsMakeDefaultCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
