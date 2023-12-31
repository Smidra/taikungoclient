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

// checks if the CheckAzureCpuQuotaCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAzureCpuQuotaCommand{}

// CheckAzureCpuQuotaCommand struct for CheckAzureCpuQuotaCommand
type CheckAzureCpuQuotaCommand struct {
	CloudId *int32 `json:"cloudId,omitempty"`
}

// NewCheckAzureCpuQuotaCommand instantiates a new CheckAzureCpuQuotaCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAzureCpuQuotaCommand() *CheckAzureCpuQuotaCommand {
	this := CheckAzureCpuQuotaCommand{}
	return &this
}

// NewCheckAzureCpuQuotaCommandWithDefaults instantiates a new CheckAzureCpuQuotaCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAzureCpuQuotaCommandWithDefaults() *CheckAzureCpuQuotaCommand {
	this := CheckAzureCpuQuotaCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *CheckAzureCpuQuotaCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAzureCpuQuotaCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *CheckAzureCpuQuotaCommand) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *CheckAzureCpuQuotaCommand) SetCloudId(v int32) {
	o.CloudId = &v
}

func (o CheckAzureCpuQuotaCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAzureCpuQuotaCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	return toSerialize, nil
}

type NullableCheckAzureCpuQuotaCommand struct {
	value *CheckAzureCpuQuotaCommand
	isSet bool
}

func (v NullableCheckAzureCpuQuotaCommand) Get() *CheckAzureCpuQuotaCommand {
	return v.value
}

func (v *NullableCheckAzureCpuQuotaCommand) Set(val *CheckAzureCpuQuotaCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAzureCpuQuotaCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAzureCpuQuotaCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAzureCpuQuotaCommand(val *CheckAzureCpuQuotaCommand) *NullableCheckAzureCpuQuotaCommand {
	return &NullableCheckAzureCpuQuotaCommand{value: val, isSet: true}
}

func (v NullableCheckAzureCpuQuotaCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAzureCpuQuotaCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
