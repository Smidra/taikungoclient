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

// checks if the ForgotPasswordCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordCommand{}

// ForgotPasswordCommand struct for ForgotPasswordCommand
type ForgotPasswordCommand struct {
	Email NullableString `json:"email,omitempty"`
}

// NewForgotPasswordCommand instantiates a new ForgotPasswordCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordCommand() *ForgotPasswordCommand {
	this := ForgotPasswordCommand{}
	return &this
}

// NewForgotPasswordCommandWithDefaults instantiates a new ForgotPasswordCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordCommandWithDefaults() *ForgotPasswordCommand {
	this := ForgotPasswordCommand{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForgotPasswordCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForgotPasswordCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ForgotPasswordCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ForgotPasswordCommand) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *ForgotPasswordCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ForgotPasswordCommand) UnsetEmail() {
	o.Email.Unset()
}

func (o ForgotPasswordCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableForgotPasswordCommand struct {
	value *ForgotPasswordCommand
	isSet bool
}

func (v NullableForgotPasswordCommand) Get() *ForgotPasswordCommand {
	return v.value
}

func (v *NullableForgotPasswordCommand) Set(val *ForgotPasswordCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordCommand(val *ForgotPasswordCommand) *NullableForgotPasswordCommand {
	return &NullableForgotPasswordCommand{value: val, isSet: true}
}

func (v NullableForgotPasswordCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
