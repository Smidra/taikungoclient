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

// checks if the BindOrganizationsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindOrganizationsCommand{}

// BindOrganizationsCommand struct for BindOrganizationsCommand
type BindOrganizationsCommand struct {
	Organizations []OrganizationDto `json:"organizations,omitempty"`
	PartnerId     NullableInt32     `json:"partnerId,omitempty"`
}

// NewBindOrganizationsCommand instantiates a new BindOrganizationsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindOrganizationsCommand() *BindOrganizationsCommand {
	this := BindOrganizationsCommand{}
	return &this
}

// NewBindOrganizationsCommandWithDefaults instantiates a new BindOrganizationsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindOrganizationsCommandWithDefaults() *BindOrganizationsCommand {
	this := BindOrganizationsCommand{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindOrganizationsCommand) GetOrganizations() []OrganizationDto {
	if o == nil {
		var ret []OrganizationDto
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindOrganizationsCommand) GetOrganizationsOk() ([]OrganizationDto, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *BindOrganizationsCommand) HasOrganizations() bool {
	if o != nil && IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationDto and assigns it to the Organizations field.
func (o *BindOrganizationsCommand) SetOrganizations(v []OrganizationDto) {
	o.Organizations = v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindOrganizationsCommand) GetPartnerId() int32 {
	if o == nil || IsNil(o.PartnerId.Get()) {
		var ret int32
		return ret
	}
	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindOrganizationsCommand) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// HasPartnerId returns a boolean if a field has been set.
func (o *BindOrganizationsCommand) HasPartnerId() bool {
	if o != nil && o.PartnerId.IsSet() {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given NullableInt32 and assigns it to the PartnerId field.
func (o *BindOrganizationsCommand) SetPartnerId(v int32) {
	o.PartnerId.Set(&v)
}

// SetPartnerIdNil sets the value for PartnerId to be an explicit nil
func (o *BindOrganizationsCommand) SetPartnerIdNil() {
	o.PartnerId.Set(nil)
}

// UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
func (o *BindOrganizationsCommand) UnsetPartnerId() {
	o.PartnerId.Unset()
}

func (o BindOrganizationsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindOrganizationsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if o.PartnerId.IsSet() {
		toSerialize["partnerId"] = o.PartnerId.Get()
	}
	return toSerialize, nil
}

type NullableBindOrganizationsCommand struct {
	value *BindOrganizationsCommand
	isSet bool
}

func (v NullableBindOrganizationsCommand) Get() *BindOrganizationsCommand {
	return v.value
}

func (v *NullableBindOrganizationsCommand) Set(val *BindOrganizationsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindOrganizationsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindOrganizationsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindOrganizationsCommand(val *BindOrganizationsCommand) *NullableBindOrganizationsCommand {
	return &NullableBindOrganizationsCommand{value: val, isSet: true}
}

func (v NullableBindOrganizationsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindOrganizationsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}