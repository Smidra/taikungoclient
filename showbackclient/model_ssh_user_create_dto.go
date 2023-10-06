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

// checks if the SshUserCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshUserCreateDto{}

// SshUserCreateDto struct for SshUserCreateDto
type SshUserCreateDto struct {
	Name         NullableString `json:"name,omitempty"`
	SshPublicKey NullableString `json:"sshPublicKey,omitempty"`
}

// NewSshUserCreateDto instantiates a new SshUserCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshUserCreateDto() *SshUserCreateDto {
	this := SshUserCreateDto{}
	return &this
}

// NewSshUserCreateDtoWithDefaults instantiates a new SshUserCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshUserCreateDtoWithDefaults() *SshUserCreateDto {
	this := SshUserCreateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshUserCreateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshUserCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SshUserCreateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SshUserCreateDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *SshUserCreateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SshUserCreateDto) UnsetName() {
	o.Name.Unset()
}

// GetSshPublicKey returns the SshPublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshUserCreateDto) GetSshPublicKey() string {
	if o == nil || IsNil(o.SshPublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPublicKey.Get()
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshUserCreateDto) GetSshPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPublicKey.Get(), o.SshPublicKey.IsSet()
}

// HasSshPublicKey returns a boolean if a field has been set.
func (o *SshUserCreateDto) HasSshPublicKey() bool {
	if o != nil && o.SshPublicKey.IsSet() {
		return true
	}

	return false
}

// SetSshPublicKey gets a reference to the given NullableString and assigns it to the SshPublicKey field.
func (o *SshUserCreateDto) SetSshPublicKey(v string) {
	o.SshPublicKey.Set(&v)
}

// SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil
func (o *SshUserCreateDto) SetSshPublicKeyNil() {
	o.SshPublicKey.Set(nil)
}

// UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
func (o *SshUserCreateDto) UnsetSshPublicKey() {
	o.SshPublicKey.Unset()
}

func (o SshUserCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshUserCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SshPublicKey.IsSet() {
		toSerialize["sshPublicKey"] = o.SshPublicKey.Get()
	}
	return toSerialize, nil
}

type NullableSshUserCreateDto struct {
	value *SshUserCreateDto
	isSet bool
}

func (v NullableSshUserCreateDto) Get() *SshUserCreateDto {
	return v.value
}

func (v *NullableSshUserCreateDto) Set(val *SshUserCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSshUserCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSshUserCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshUserCreateDto(val *SshUserCreateDto) *NullableSshUserCreateDto {
	return &NullableSshUserCreateDto{value: val, isSet: true}
}

func (v NullableSshUserCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshUserCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}