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

// checks if the ProjectButtonStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectButtonStatusDto{}

// ProjectButtonStatusDto struct for ProjectButtonStatusDto
type ProjectButtonStatusDto struct {
	Enable  *bool    `json:"enable,omitempty"`
	Reasons []string `json:"reasons,omitempty"`
}

// NewProjectButtonStatusDto instantiates a new ProjectButtonStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectButtonStatusDto() *ProjectButtonStatusDto {
	this := ProjectButtonStatusDto{}
	return &this
}

// NewProjectButtonStatusDtoWithDefaults instantiates a new ProjectButtonStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectButtonStatusDtoWithDefaults() *ProjectButtonStatusDto {
	this := ProjectButtonStatusDto{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ProjectButtonStatusDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectButtonStatusDto) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ProjectButtonStatusDto) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ProjectButtonStatusDto) SetEnable(v bool) {
	o.Enable = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectButtonStatusDto) GetReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectButtonStatusDto) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *ProjectButtonStatusDto) HasReasons() bool {
	if o != nil && IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *ProjectButtonStatusDto) SetReasons(v []string) {
	o.Reasons = v
}

func (o ProjectButtonStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectButtonStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableProjectButtonStatusDto struct {
	value *ProjectButtonStatusDto
	isSet bool
}

func (v NullableProjectButtonStatusDto) Get() *ProjectButtonStatusDto {
	return v.value
}

func (v *NullableProjectButtonStatusDto) Set(val *ProjectButtonStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectButtonStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectButtonStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectButtonStatusDto(val *ProjectButtonStatusDto) *NullableProjectButtonStatusDto {
	return &NullableProjectButtonStatusDto{value: val, isSet: true}
}

func (v NullableProjectButtonStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectButtonStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
