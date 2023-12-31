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

// checks if the UserTokenCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTokenCreateDto{}

// UserTokenCreateDto struct for UserTokenCreateDto
type UserTokenCreateDto struct {
	AccessKey NullableString `json:"accessKey,omitempty"`
	SecretKey NullableString `json:"secretKey,omitempty"`
}

// NewUserTokenCreateDto instantiates a new UserTokenCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTokenCreateDto() *UserTokenCreateDto {
	this := UserTokenCreateDto{}
	return &this
}

// NewUserTokenCreateDtoWithDefaults instantiates a new UserTokenCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokenCreateDtoWithDefaults() *UserTokenCreateDto {
	this := UserTokenCreateDto{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTokenCreateDto) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey.Get()) {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokenCreateDto) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *UserTokenCreateDto) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *UserTokenCreateDto) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}

// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *UserTokenCreateDto) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *UserTokenCreateDto) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTokenCreateDto) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey.Get()) {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokenCreateDto) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *UserTokenCreateDto) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *UserTokenCreateDto) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}

// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *UserTokenCreateDto) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *UserTokenCreateDto) UnsetSecretKey() {
	o.SecretKey.Unset()
}

func (o UserTokenCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTokenCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey.IsSet() {
		toSerialize["accessKey"] = o.AccessKey.Get()
	}
	if o.SecretKey.IsSet() {
		toSerialize["secretKey"] = o.SecretKey.Get()
	}
	return toSerialize, nil
}

type NullableUserTokenCreateDto struct {
	value *UserTokenCreateDto
	isSet bool
}

func (v NullableUserTokenCreateDto) Get() *UserTokenCreateDto {
	return v.value
}

func (v *NullableUserTokenCreateDto) Set(val *UserTokenCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTokenCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTokenCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTokenCreateDto(val *UserTokenCreateDto) *NullableUserTokenCreateDto {
	return &NullableUserTokenCreateDto{value: val, isSet: true}
}

func (v NullableUserTokenCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTokenCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
