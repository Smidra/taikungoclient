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

// checks if the DnsServerCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsServerCreateDto{}

// DnsServerCreateDto struct for DnsServerCreateDto
type DnsServerCreateDto struct {
	Address NullableString `json:"address,omitempty"`
}

// NewDnsServerCreateDto instantiates a new DnsServerCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsServerCreateDto() *DnsServerCreateDto {
	this := DnsServerCreateDto{}
	return &this
}

// NewDnsServerCreateDtoWithDefaults instantiates a new DnsServerCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsServerCreateDtoWithDefaults() *DnsServerCreateDto {
	this := DnsServerCreateDto{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsServerCreateDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsServerCreateDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *DnsServerCreateDto) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *DnsServerCreateDto) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *DnsServerCreateDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *DnsServerCreateDto) UnsetAddress() {
	o.Address.Unset()
}

func (o DnsServerCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsServerCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return toSerialize, nil
}

type NullableDnsServerCreateDto struct {
	value *DnsServerCreateDto
	isSet bool
}

func (v NullableDnsServerCreateDto) Get() *DnsServerCreateDto {
	return v.value
}

func (v *NullableDnsServerCreateDto) Set(val *DnsServerCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsServerCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsServerCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsServerCreateDto(val *DnsServerCreateDto) *NullableDnsServerCreateDto {
	return &NullableDnsServerCreateDto{value: val, isSet: true}
}

func (v NullableDnsServerCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsServerCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
