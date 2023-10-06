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

// checks if the OrganizationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDto{}

// OrganizationDto struct for OrganizationDto
type OrganizationDto struct {
	OrganizationId *int32 `json:"organizationId,omitempty"`
	IsBound        *bool  `json:"isBound,omitempty"`
}

// NewOrganizationDto instantiates a new OrganizationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDto() *OrganizationDto {
	this := OrganizationDto{}
	return &this
}

// NewOrganizationDtoWithDefaults instantiates a new OrganizationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDtoWithDefaults() *OrganizationDto {
	this := OrganizationDto{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *OrganizationDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *OrganizationDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *OrganizationDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *OrganizationDto) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o OrganizationDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableOrganizationDto struct {
	value *OrganizationDto
	isSet bool
}

func (v NullableOrganizationDto) Get() *OrganizationDto {
	return v.value
}

func (v *NullableOrganizationDto) Set(val *OrganizationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDto(val *OrganizationDto) *NullableOrganizationDto {
	return &NullableOrganizationDto{value: val, isSet: true}
}

func (v NullableOrganizationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
