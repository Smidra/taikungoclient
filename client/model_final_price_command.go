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

// checks if the FinalPriceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinalPriceCommand{}

// FinalPriceCommand struct for FinalPriceCommand
type FinalPriceCommand struct {
	SubscriptionId *int32 `json:"subscriptionId,omitempty"`
}

// NewFinalPriceCommand instantiates a new FinalPriceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinalPriceCommand() *FinalPriceCommand {
	this := FinalPriceCommand{}
	return &this
}

// NewFinalPriceCommandWithDefaults instantiates a new FinalPriceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinalPriceCommandWithDefaults() *FinalPriceCommand {
	this := FinalPriceCommand{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *FinalPriceCommand) GetSubscriptionId() int32 {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret int32
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalPriceCommand) GetSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *FinalPriceCommand) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given int32 and assigns it to the SubscriptionId field.
func (o *FinalPriceCommand) SetSubscriptionId(v int32) {
	o.SubscriptionId = &v
}

func (o FinalPriceCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinalPriceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableFinalPriceCommand struct {
	value *FinalPriceCommand
	isSet bool
}

func (v NullableFinalPriceCommand) Get() *FinalPriceCommand {
	return v.value
}

func (v *NullableFinalPriceCommand) Set(val *FinalPriceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableFinalPriceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableFinalPriceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinalPriceCommand(val *FinalPriceCommand) *NullableFinalPriceCommand {
	return &NullableFinalPriceCommand{value: val, isSet: true}
}

func (v NullableFinalPriceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinalPriceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
