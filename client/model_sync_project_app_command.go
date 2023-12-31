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

// checks if the SyncProjectAppCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncProjectAppCommand{}

// SyncProjectAppCommand struct for SyncProjectAppCommand
type SyncProjectAppCommand struct {
	ProjectAppId *int32 `json:"projectAppId,omitempty"`
}

// NewSyncProjectAppCommand instantiates a new SyncProjectAppCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncProjectAppCommand() *SyncProjectAppCommand {
	this := SyncProjectAppCommand{}
	return &this
}

// NewSyncProjectAppCommandWithDefaults instantiates a new SyncProjectAppCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncProjectAppCommandWithDefaults() *SyncProjectAppCommand {
	this := SyncProjectAppCommand{}
	return &this
}

// GetProjectAppId returns the ProjectAppId field value if set, zero value otherwise.
func (o *SyncProjectAppCommand) GetProjectAppId() int32 {
	if o == nil || IsNil(o.ProjectAppId) {
		var ret int32
		return ret
	}
	return *o.ProjectAppId
}

// GetProjectAppIdOk returns a tuple with the ProjectAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncProjectAppCommand) GetProjectAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectAppId) {
		return nil, false
	}
	return o.ProjectAppId, true
}

// HasProjectAppId returns a boolean if a field has been set.
func (o *SyncProjectAppCommand) HasProjectAppId() bool {
	if o != nil && !IsNil(o.ProjectAppId) {
		return true
	}

	return false
}

// SetProjectAppId gets a reference to the given int32 and assigns it to the ProjectAppId field.
func (o *SyncProjectAppCommand) SetProjectAppId(v int32) {
	o.ProjectAppId = &v
}

func (o SyncProjectAppCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncProjectAppCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectAppId) {
		toSerialize["projectAppId"] = o.ProjectAppId
	}
	return toSerialize, nil
}

type NullableSyncProjectAppCommand struct {
	value *SyncProjectAppCommand
	isSet bool
}

func (v NullableSyncProjectAppCommand) Get() *SyncProjectAppCommand {
	return v.value
}

func (v *NullableSyncProjectAppCommand) Set(val *SyncProjectAppCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncProjectAppCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncProjectAppCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncProjectAppCommand(val *SyncProjectAppCommand) *NullableSyncProjectAppCommand {
	return &NullableSyncProjectAppCommand{value: val, isSet: true}
}

func (v NullableSyncProjectAppCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncProjectAppCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
