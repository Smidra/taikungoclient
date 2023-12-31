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

// checks if the JsonNodeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonNodeOptions{}

// JsonNodeOptions struct for JsonNodeOptions
type JsonNodeOptions struct {
	PropertyNameCaseInsensitive *bool `json:"propertyNameCaseInsensitive,omitempty"`
}

// NewJsonNodeOptions instantiates a new JsonNodeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonNodeOptions() *JsonNodeOptions {
	this := JsonNodeOptions{}
	return &this
}

// NewJsonNodeOptionsWithDefaults instantiates a new JsonNodeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonNodeOptionsWithDefaults() *JsonNodeOptions {
	this := JsonNodeOptions{}
	return &this
}

// GetPropertyNameCaseInsensitive returns the PropertyNameCaseInsensitive field value if set, zero value otherwise.
func (o *JsonNodeOptions) GetPropertyNameCaseInsensitive() bool {
	if o == nil || IsNil(o.PropertyNameCaseInsensitive) {
		var ret bool
		return ret
	}
	return *o.PropertyNameCaseInsensitive
}

// GetPropertyNameCaseInsensitiveOk returns a tuple with the PropertyNameCaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeOptions) GetPropertyNameCaseInsensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.PropertyNameCaseInsensitive) {
		return nil, false
	}
	return o.PropertyNameCaseInsensitive, true
}

// HasPropertyNameCaseInsensitive returns a boolean if a field has been set.
func (o *JsonNodeOptions) HasPropertyNameCaseInsensitive() bool {
	if o != nil && !IsNil(o.PropertyNameCaseInsensitive) {
		return true
	}

	return false
}

// SetPropertyNameCaseInsensitive gets a reference to the given bool and assigns it to the PropertyNameCaseInsensitive field.
func (o *JsonNodeOptions) SetPropertyNameCaseInsensitive(v bool) {
	o.PropertyNameCaseInsensitive = &v
}

func (o JsonNodeOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonNodeOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyNameCaseInsensitive) {
		toSerialize["propertyNameCaseInsensitive"] = o.PropertyNameCaseInsensitive
	}
	return toSerialize, nil
}

type NullableJsonNodeOptions struct {
	value *JsonNodeOptions
	isSet bool
}

func (v NullableJsonNodeOptions) Get() *JsonNodeOptions {
	return v.value
}

func (v *NullableJsonNodeOptions) Set(val *JsonNodeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonNodeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonNodeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonNodeOptions(val *JsonNodeOptions) *NullableJsonNodeOptions {
	return &NullableJsonNodeOptions{value: val, isSet: true}
}

func (v NullableJsonNodeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonNodeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
