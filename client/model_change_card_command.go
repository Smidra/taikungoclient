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

// checks if the ChangeCardCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeCardCommand{}

// ChangeCardCommand struct for ChangeCardCommand
type ChangeCardCommand struct {
	PaymentMethodId NullableString `json:"paymentMethodId,omitempty"`
}

// NewChangeCardCommand instantiates a new ChangeCardCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeCardCommand() *ChangeCardCommand {
	this := ChangeCardCommand{}
	return &this
}

// NewChangeCardCommandWithDefaults instantiates a new ChangeCardCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeCardCommandWithDefaults() *ChangeCardCommand {
	this := ChangeCardCommand{}
	return &this
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeCardCommand) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeCardCommand) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *ChangeCardCommand) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given NullableString and assigns it to the PaymentMethodId field.
func (o *ChangeCardCommand) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}

// SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil
func (o *ChangeCardCommand) SetPaymentMethodIdNil() {
	o.PaymentMethodId.Set(nil)
}

// UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
func (o *ChangeCardCommand) UnsetPaymentMethodId() {
	o.PaymentMethodId.Unset()
}

func (o ChangeCardCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeCardCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethodId.IsSet() {
		toSerialize["paymentMethodId"] = o.PaymentMethodId.Get()
	}
	return toSerialize, nil
}

type NullableChangeCardCommand struct {
	value *ChangeCardCommand
	isSet bool
}

func (v NullableChangeCardCommand) Get() *ChangeCardCommand {
	return v.value
}

func (v *NullableChangeCardCommand) Set(val *ChangeCardCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeCardCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeCardCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeCardCommand(val *ChangeCardCommand) *NullableChangeCardCommand {
	return &NullableChangeCardCommand{value: val, isSet: true}
}

func (v NullableChangeCardCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeCardCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}