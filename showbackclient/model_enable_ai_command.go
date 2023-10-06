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

// checks if the EnableAiCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableAiCommand{}

// EnableAiCommand struct for EnableAiCommand
type EnableAiCommand struct {
	ProjectId      *int32 `json:"projectId,omitempty"`
	AiCredentialId *int32 `json:"aiCredentialId,omitempty"`
}

// NewEnableAiCommand instantiates a new EnableAiCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableAiCommand() *EnableAiCommand {
	this := EnableAiCommand{}
	return &this
}

// NewEnableAiCommandWithDefaults instantiates a new EnableAiCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableAiCommandWithDefaults() *EnableAiCommand {
	this := EnableAiCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *EnableAiCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAiCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *EnableAiCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *EnableAiCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetAiCredentialId returns the AiCredentialId field value if set, zero value otherwise.
func (o *EnableAiCommand) GetAiCredentialId() int32 {
	if o == nil || IsNil(o.AiCredentialId) {
		var ret int32
		return ret
	}
	return *o.AiCredentialId
}

// GetAiCredentialIdOk returns a tuple with the AiCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableAiCommand) GetAiCredentialIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AiCredentialId) {
		return nil, false
	}
	return o.AiCredentialId, true
}

// HasAiCredentialId returns a boolean if a field has been set.
func (o *EnableAiCommand) HasAiCredentialId() bool {
	if o != nil && !IsNil(o.AiCredentialId) {
		return true
	}

	return false
}

// SetAiCredentialId gets a reference to the given int32 and assigns it to the AiCredentialId field.
func (o *EnableAiCommand) SetAiCredentialId(v int32) {
	o.AiCredentialId = &v
}

func (o EnableAiCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableAiCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.AiCredentialId) {
		toSerialize["aiCredentialId"] = o.AiCredentialId
	}
	return toSerialize, nil
}

type NullableEnableAiCommand struct {
	value *EnableAiCommand
	isSet bool
}

func (v NullableEnableAiCommand) Get() *EnableAiCommand {
	return v.value
}

func (v *NullableEnableAiCommand) Set(val *EnableAiCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableAiCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableAiCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableAiCommand(val *EnableAiCommand) *NullableEnableAiCommand {
	return &NullableEnableAiCommand{value: val, isSet: true}
}

func (v NullableEnableAiCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableAiCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
