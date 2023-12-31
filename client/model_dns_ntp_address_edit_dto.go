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

// checks if the DnsNtpAddressEditDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsNtpAddressEditDto{}

// DnsNtpAddressEditDto struct for DnsNtpAddressEditDto
type DnsNtpAddressEditDto struct {
	Address NullableString `json:"address,omitempty"`
}

// NewDnsNtpAddressEditDto instantiates a new DnsNtpAddressEditDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsNtpAddressEditDto() *DnsNtpAddressEditDto {
	this := DnsNtpAddressEditDto{}
	return &this
}

// NewDnsNtpAddressEditDtoWithDefaults instantiates a new DnsNtpAddressEditDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsNtpAddressEditDtoWithDefaults() *DnsNtpAddressEditDto {
	this := DnsNtpAddressEditDto{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsNtpAddressEditDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsNtpAddressEditDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *DnsNtpAddressEditDto) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *DnsNtpAddressEditDto) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *DnsNtpAddressEditDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *DnsNtpAddressEditDto) UnsetAddress() {
	o.Address.Unset()
}

func (o DnsNtpAddressEditDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsNtpAddressEditDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return toSerialize, nil
}

type NullableDnsNtpAddressEditDto struct {
	value *DnsNtpAddressEditDto
	isSet bool
}

func (v NullableDnsNtpAddressEditDto) Get() *DnsNtpAddressEditDto {
	return v.value
}

func (v *NullableDnsNtpAddressEditDto) Set(val *DnsNtpAddressEditDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsNtpAddressEditDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsNtpAddressEditDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsNtpAddressEditDto(val *DnsNtpAddressEditDto) *NullableDnsNtpAddressEditDto {
	return &NullableDnsNtpAddressEditDto{value: val, isSet: true}
}

func (v NullableDnsNtpAddressEditDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsNtpAddressEditDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
