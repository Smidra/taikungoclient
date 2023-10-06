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

// checks if the DeleteSubscriptionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSubscriptionCommand{}

// DeleteSubscriptionCommand struct for DeleteSubscriptionCommand
type DeleteSubscriptionCommand struct {
	Id *int32 `json:"id,omitempty"`
}

// NewDeleteSubscriptionCommand instantiates a new DeleteSubscriptionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSubscriptionCommand() *DeleteSubscriptionCommand {
	this := DeleteSubscriptionCommand{}
	return &this
}

// NewDeleteSubscriptionCommandWithDefaults instantiates a new DeleteSubscriptionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSubscriptionCommandWithDefaults() *DeleteSubscriptionCommand {
	this := DeleteSubscriptionCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteSubscriptionCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubscriptionCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteSubscriptionCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeleteSubscriptionCommand) SetId(v int32) {
	o.Id = &v
}

func (o DeleteSubscriptionCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSubscriptionCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteSubscriptionCommand struct {
	value *DeleteSubscriptionCommand
	isSet bool
}

func (v NullableDeleteSubscriptionCommand) Get() *DeleteSubscriptionCommand {
	return v.value
}

func (v *NullableDeleteSubscriptionCommand) Set(val *DeleteSubscriptionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSubscriptionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSubscriptionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSubscriptionCommand(val *DeleteSubscriptionCommand) *NullableDeleteSubscriptionCommand {
	return &NullableDeleteSubscriptionCommand{value: val, isSet: true}
}

func (v NullableDeleteSubscriptionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSubscriptionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}