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

// checks if the LeaveTaikunCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LeaveTaikunCommand{}

// LeaveTaikunCommand struct for LeaveTaikunCommand
type LeaveTaikunCommand struct {
	Reason  NullableString `json:"reason,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewLeaveTaikunCommand instantiates a new LeaveTaikunCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaveTaikunCommand() *LeaveTaikunCommand {
	this := LeaveTaikunCommand{}
	return &this
}

// NewLeaveTaikunCommandWithDefaults instantiates a new LeaveTaikunCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaveTaikunCommandWithDefaults() *LeaveTaikunCommand {
	this := LeaveTaikunCommand{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeaveTaikunCommand) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeaveTaikunCommand) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *LeaveTaikunCommand) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *LeaveTaikunCommand) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *LeaveTaikunCommand) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *LeaveTaikunCommand) UnsetReason() {
	o.Reason.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeaveTaikunCommand) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeaveTaikunCommand) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *LeaveTaikunCommand) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *LeaveTaikunCommand) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *LeaveTaikunCommand) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *LeaveTaikunCommand) UnsetMessage() {
	o.Message.Unset()
}

func (o LeaveTaikunCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LeaveTaikunCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableLeaveTaikunCommand struct {
	value *LeaveTaikunCommand
	isSet bool
}

func (v NullableLeaveTaikunCommand) Get() *LeaveTaikunCommand {
	return v.value
}

func (v *NullableLeaveTaikunCommand) Set(val *LeaveTaikunCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaveTaikunCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaveTaikunCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaveTaikunCommand(val *LeaveTaikunCommand) *NullableLeaveTaikunCommand {
	return &NullableLeaveTaikunCommand{value: val, isSet: true}
}

func (v NullableLeaveTaikunCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaveTaikunCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
