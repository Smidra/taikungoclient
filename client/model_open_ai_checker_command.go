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

// checks if the OpenAiCheckerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAiCheckerCommand{}

// OpenAiCheckerCommand struct for OpenAiCheckerCommand
type OpenAiCheckerCommand struct {
	Token NullableString `json:"token,omitempty"`
}

// NewOpenAiCheckerCommand instantiates a new OpenAiCheckerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAiCheckerCommand() *OpenAiCheckerCommand {
	this := OpenAiCheckerCommand{}
	return &this
}

// NewOpenAiCheckerCommandWithDefaults instantiates a new OpenAiCheckerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAiCheckerCommandWithDefaults() *OpenAiCheckerCommand {
	this := OpenAiCheckerCommand{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenAiCheckerCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenAiCheckerCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *OpenAiCheckerCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *OpenAiCheckerCommand) SetToken(v string) {
	o.Token.Set(&v)
}

// SetTokenNil sets the value for Token to be an explicit nil
func (o *OpenAiCheckerCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *OpenAiCheckerCommand) UnsetToken() {
	o.Token.Unset()
}

func (o OpenAiCheckerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAiCheckerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableOpenAiCheckerCommand struct {
	value *OpenAiCheckerCommand
	isSet bool
}

func (v NullableOpenAiCheckerCommand) Get() *OpenAiCheckerCommand {
	return v.value
}

func (v *NullableOpenAiCheckerCommand) Set(val *OpenAiCheckerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAiCheckerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAiCheckerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAiCheckerCommand(val *OpenAiCheckerCommand) *NullableOpenAiCheckerCommand {
	return &NullableOpenAiCheckerCommand{value: val, isSet: true}
}

func (v NullableOpenAiCheckerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAiCheckerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
