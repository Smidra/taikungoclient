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

// checks if the UpdateHealthStatusCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHealthStatusCommand{}

// UpdateHealthStatusCommand struct for UpdateHealthStatusCommand
type UpdateHealthStatusCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Health    *ProjectHealth `json:"health,omitempty"`
}

// NewUpdateHealthStatusCommand instantiates a new UpdateHealthStatusCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHealthStatusCommand() *UpdateHealthStatusCommand {
	this := UpdateHealthStatusCommand{}
	return &this
}

// NewUpdateHealthStatusCommandWithDefaults instantiates a new UpdateHealthStatusCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHealthStatusCommandWithDefaults() *UpdateHealthStatusCommand {
	this := UpdateHealthStatusCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *UpdateHealthStatusCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHealthStatusCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *UpdateHealthStatusCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *UpdateHealthStatusCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *UpdateHealthStatusCommand) GetHealth() ProjectHealth {
	if o == nil || IsNil(o.Health) {
		var ret ProjectHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHealthStatusCommand) GetHealthOk() (*ProjectHealth, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *UpdateHealthStatusCommand) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given ProjectHealth and assigns it to the Health field.
func (o *UpdateHealthStatusCommand) SetHealth(v ProjectHealth) {
	o.Health = &v
}

func (o UpdateHealthStatusCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHealthStatusCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	return toSerialize, nil
}

type NullableUpdateHealthStatusCommand struct {
	value *UpdateHealthStatusCommand
	isSet bool
}

func (v NullableUpdateHealthStatusCommand) Get() *UpdateHealthStatusCommand {
	return v.value
}

func (v *NullableUpdateHealthStatusCommand) Set(val *UpdateHealthStatusCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHealthStatusCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHealthStatusCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHealthStatusCommand(val *UpdateHealthStatusCommand) *NullableUpdateHealthStatusCommand {
	return &NullableUpdateHealthStatusCommand{value: val, isSet: true}
}

func (v NullableUpdateHealthStatusCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHealthStatusCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
