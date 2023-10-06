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

// checks if the UpdateAzureCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAzureCommand{}

// UpdateAzureCommand struct for UpdateAzureCommand
type UpdateAzureCommand struct {
	Id                *int32         `json:"id,omitempty"`
	Name              NullableString `json:"name,omitempty"`
	AzureClientSecret NullableString `json:"azureClientSecret,omitempty"`
	AzureClientId     NullableString `json:"azureClientId,omitempty"`
}

// NewUpdateAzureCommand instantiates a new UpdateAzureCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAzureCommand() *UpdateAzureCommand {
	this := UpdateAzureCommand{}
	return &this
}

// NewUpdateAzureCommandWithDefaults instantiates a new UpdateAzureCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAzureCommandWithDefaults() *UpdateAzureCommand {
	this := UpdateAzureCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAzureCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAzureCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAzureCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateAzureCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAzureCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAzureCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAzureCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateAzureCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateAzureCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateAzureCommand) UnsetName() {
	o.Name.Unset()
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAzureCommand) GetAzureClientSecret() string {
	if o == nil || IsNil(o.AzureClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientSecret.Get()
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAzureCommand) GetAzureClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientSecret.Get(), o.AzureClientSecret.IsSet()
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *UpdateAzureCommand) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret.IsSet() {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given NullableString and assigns it to the AzureClientSecret field.
func (o *UpdateAzureCommand) SetAzureClientSecret(v string) {
	o.AzureClientSecret.Set(&v)
}

// SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil
func (o *UpdateAzureCommand) SetAzureClientSecretNil() {
	o.AzureClientSecret.Set(nil)
}

// UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
func (o *UpdateAzureCommand) UnsetAzureClientSecret() {
	o.AzureClientSecret.Unset()
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAzureCommand) GetAzureClientId() string {
	if o == nil || IsNil(o.AzureClientId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientId.Get()
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAzureCommand) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientId.Get(), o.AzureClientId.IsSet()
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *UpdateAzureCommand) HasAzureClientId() bool {
	if o != nil && o.AzureClientId.IsSet() {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given NullableString and assigns it to the AzureClientId field.
func (o *UpdateAzureCommand) SetAzureClientId(v string) {
	o.AzureClientId.Set(&v)
}

// SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil
func (o *UpdateAzureCommand) SetAzureClientIdNil() {
	o.AzureClientId.Set(nil)
}

// UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
func (o *UpdateAzureCommand) UnsetAzureClientId() {
	o.AzureClientId.Unset()
}

func (o UpdateAzureCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAzureCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AzureClientSecret.IsSet() {
		toSerialize["azureClientSecret"] = o.AzureClientSecret.Get()
	}
	if o.AzureClientId.IsSet() {
		toSerialize["azureClientId"] = o.AzureClientId.Get()
	}
	return toSerialize, nil
}

type NullableUpdateAzureCommand struct {
	value *UpdateAzureCommand
	isSet bool
}

func (v NullableUpdateAzureCommand) Get() *UpdateAzureCommand {
	return v.value
}

func (v *NullableUpdateAzureCommand) Set(val *UpdateAzureCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAzureCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAzureCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAzureCommand(val *UpdateAzureCommand) *NullableUpdateAzureCommand {
	return &NullableUpdateAzureCommand{value: val, isSet: true}
}

func (v NullableUpdateAzureCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAzureCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
