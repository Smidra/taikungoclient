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

// checks if the CommitStandAloneVmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitStandAloneVmCommand{}

// CommitStandAloneVmCommand struct for CommitStandAloneVmCommand
type CommitStandAloneVmCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewCommitStandAloneVmCommand instantiates a new CommitStandAloneVmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitStandAloneVmCommand() *CommitStandAloneVmCommand {
	this := CommitStandAloneVmCommand{}
	return &this
}

// NewCommitStandAloneVmCommandWithDefaults instantiates a new CommitStandAloneVmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitStandAloneVmCommandWithDefaults() *CommitStandAloneVmCommand {
	this := CommitStandAloneVmCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CommitStandAloneVmCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitStandAloneVmCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CommitStandAloneVmCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CommitStandAloneVmCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o CommitStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitStandAloneVmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableCommitStandAloneVmCommand struct {
	value *CommitStandAloneVmCommand
	isSet bool
}

func (v NullableCommitStandAloneVmCommand) Get() *CommitStandAloneVmCommand {
	return v.value
}

func (v *NullableCommitStandAloneVmCommand) Set(val *CommitStandAloneVmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitStandAloneVmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitStandAloneVmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitStandAloneVmCommand(val *CommitStandAloneVmCommand) *NullableCommitStandAloneVmCommand {
	return &NullableCommitStandAloneVmCommand{value: val, isSet: true}
}

func (v NullableCommitStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitStandAloneVmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
