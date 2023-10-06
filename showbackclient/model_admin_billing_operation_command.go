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

// checks if the AdminBillingOperationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminBillingOperationCommand{}

// AdminBillingOperationCommand struct for AdminBillingOperationCommand
type AdminBillingOperationCommand struct {
	CloudCredentialId *int32 `json:"cloudCredentialId,omitempty"`
}

// NewAdminBillingOperationCommand instantiates a new AdminBillingOperationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminBillingOperationCommand() *AdminBillingOperationCommand {
	this := AdminBillingOperationCommand{}
	return &this
}

// NewAdminBillingOperationCommandWithDefaults instantiates a new AdminBillingOperationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminBillingOperationCommandWithDefaults() *AdminBillingOperationCommand {
	this := AdminBillingOperationCommand{}
	return &this
}

// GetCloudCredentialId returns the CloudCredentialId field value if set, zero value otherwise.
func (o *AdminBillingOperationCommand) GetCloudCredentialId() int32 {
	if o == nil || IsNil(o.CloudCredentialId) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialId
}

// GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminBillingOperationCommand) GetCloudCredentialIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudCredentialId) {
		return nil, false
	}
	return o.CloudCredentialId, true
}

// HasCloudCredentialId returns a boolean if a field has been set.
func (o *AdminBillingOperationCommand) HasCloudCredentialId() bool {
	if o != nil && !IsNil(o.CloudCredentialId) {
		return true
	}

	return false
}

// SetCloudCredentialId gets a reference to the given int32 and assigns it to the CloudCredentialId field.
func (o *AdminBillingOperationCommand) SetCloudCredentialId(v int32) {
	o.CloudCredentialId = &v
}

func (o AdminBillingOperationCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminBillingOperationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudCredentialId) {
		toSerialize["cloudCredentialId"] = o.CloudCredentialId
	}
	return toSerialize, nil
}

type NullableAdminBillingOperationCommand struct {
	value *AdminBillingOperationCommand
	isSet bool
}

func (v NullableAdminBillingOperationCommand) Get() *AdminBillingOperationCommand {
	return v.value
}

func (v *NullableAdminBillingOperationCommand) Set(val *AdminBillingOperationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminBillingOperationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminBillingOperationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminBillingOperationCommand(val *AdminBillingOperationCommand) *NullableAdminBillingOperationCommand {
	return &NullableAdminBillingOperationCommand{value: val, isSet: true}
}

func (v NullableAdminBillingOperationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminBillingOperationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}