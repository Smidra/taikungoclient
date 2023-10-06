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

// checks if the CidrCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CidrCommand{}

// CidrCommand struct for CidrCommand
type CidrCommand struct {
	Cidr NullableString `json:"cidr,omitempty"`
}

// NewCidrCommand instantiates a new CidrCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCidrCommand() *CidrCommand {
	this := CidrCommand{}
	return &this
}

// NewCidrCommandWithDefaults instantiates a new CidrCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCidrCommandWithDefaults() *CidrCommand {
	this := CidrCommand{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CidrCommand) GetCidr() string {
	if o == nil || IsNil(o.Cidr.Get()) {
		var ret string
		return ret
	}
	return *o.Cidr.Get()
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CidrCommand) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cidr.Get(), o.Cidr.IsSet()
}

// HasCidr returns a boolean if a field has been set.
func (o *CidrCommand) HasCidr() bool {
	if o != nil && o.Cidr.IsSet() {
		return true
	}

	return false
}

// SetCidr gets a reference to the given NullableString and assigns it to the Cidr field.
func (o *CidrCommand) SetCidr(v string) {
	o.Cidr.Set(&v)
}

// SetCidrNil sets the value for Cidr to be an explicit nil
func (o *CidrCommand) SetCidrNil() {
	o.Cidr.Set(nil)
}

// UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
func (o *CidrCommand) UnsetCidr() {
	o.Cidr.Unset()
}

func (o CidrCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CidrCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cidr.IsSet() {
		toSerialize["cidr"] = o.Cidr.Get()
	}
	return toSerialize, nil
}

type NullableCidrCommand struct {
	value *CidrCommand
	isSet bool
}

func (v NullableCidrCommand) Get() *CidrCommand {
	return v.value
}

func (v *NullableCidrCommand) Set(val *CidrCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCidrCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCidrCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCidrCommand(val *CidrCommand) *NullableCidrCommand {
	return &NullableCidrCommand{value: val, isSet: true}
}

func (v NullableCidrCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCidrCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
