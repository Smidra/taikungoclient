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

// checks if the CommonStringBasedDropdownDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonStringBasedDropdownDto{}

// CommonStringBasedDropdownDto struct for CommonStringBasedDropdownDto
type CommonStringBasedDropdownDto struct {
	Id   NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewCommonStringBasedDropdownDto instantiates a new CommonStringBasedDropdownDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonStringBasedDropdownDto() *CommonStringBasedDropdownDto {
	this := CommonStringBasedDropdownDto{}
	return &this
}

// NewCommonStringBasedDropdownDtoWithDefaults instantiates a new CommonStringBasedDropdownDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonStringBasedDropdownDtoWithDefaults() *CommonStringBasedDropdownDto {
	this := CommonStringBasedDropdownDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonStringBasedDropdownDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonStringBasedDropdownDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CommonStringBasedDropdownDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CommonStringBasedDropdownDto) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CommonStringBasedDropdownDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CommonStringBasedDropdownDto) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonStringBasedDropdownDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonStringBasedDropdownDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CommonStringBasedDropdownDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CommonStringBasedDropdownDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CommonStringBasedDropdownDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CommonStringBasedDropdownDto) UnsetName() {
	o.Name.Unset()
}

func (o CommonStringBasedDropdownDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonStringBasedDropdownDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableCommonStringBasedDropdownDto struct {
	value *CommonStringBasedDropdownDto
	isSet bool
}

func (v NullableCommonStringBasedDropdownDto) Get() *CommonStringBasedDropdownDto {
	return v.value
}

func (v *NullableCommonStringBasedDropdownDto) Set(val *CommonStringBasedDropdownDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonStringBasedDropdownDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonStringBasedDropdownDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonStringBasedDropdownDto(val *CommonStringBasedDropdownDto) *NullableCommonStringBasedDropdownDto {
	return &NullableCommonStringBasedDropdownDto{value: val, isSet: true}
}

func (v NullableCommonStringBasedDropdownDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonStringBasedDropdownDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
