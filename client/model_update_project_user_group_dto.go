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

// checks if the UpdateProjectUserGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectUserGroupDto{}

// UpdateProjectUserGroupDto struct for UpdateProjectUserGroupDto
type UpdateProjectUserGroupDto struct {
	UserGroupId *int32 `json:"userGroupId,omitempty"`
	IsBound     *bool  `json:"isBound,omitempty"`
}

// NewUpdateProjectUserGroupDto instantiates a new UpdateProjectUserGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectUserGroupDto() *UpdateProjectUserGroupDto {
	this := UpdateProjectUserGroupDto{}
	return &this
}

// NewUpdateProjectUserGroupDtoWithDefaults instantiates a new UpdateProjectUserGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectUserGroupDtoWithDefaults() *UpdateProjectUserGroupDto {
	this := UpdateProjectUserGroupDto{}
	return &this
}

// GetUserGroupId returns the UserGroupId field value if set, zero value otherwise.
func (o *UpdateProjectUserGroupDto) GetUserGroupId() int32 {
	if o == nil || IsNil(o.UserGroupId) {
		var ret int32
		return ret
	}
	return *o.UserGroupId
}

// GetUserGroupIdOk returns a tuple with the UserGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectUserGroupDto) GetUserGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserGroupId) {
		return nil, false
	}
	return o.UserGroupId, true
}

// HasUserGroupId returns a boolean if a field has been set.
func (o *UpdateProjectUserGroupDto) HasUserGroupId() bool {
	if o != nil && !IsNil(o.UserGroupId) {
		return true
	}

	return false
}

// SetUserGroupId gets a reference to the given int32 and assigns it to the UserGroupId field.
func (o *UpdateProjectUserGroupDto) SetUserGroupId(v int32) {
	o.UserGroupId = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *UpdateProjectUserGroupDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectUserGroupDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *UpdateProjectUserGroupDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *UpdateProjectUserGroupDto) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o UpdateProjectUserGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectUserGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserGroupId) {
		toSerialize["userGroupId"] = o.UserGroupId
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableUpdateProjectUserGroupDto struct {
	value *UpdateProjectUserGroupDto
	isSet bool
}

func (v NullableUpdateProjectUserGroupDto) Get() *UpdateProjectUserGroupDto {
	return v.value
}

func (v *NullableUpdateProjectUserGroupDto) Set(val *UpdateProjectUserGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectUserGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectUserGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectUserGroupDto(val *UpdateProjectUserGroupDto) *NullableUpdateProjectUserGroupDto {
	return &NullableUpdateProjectUserGroupDto{value: val, isSet: true}
}

func (v NullableUpdateProjectUserGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectUserGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
