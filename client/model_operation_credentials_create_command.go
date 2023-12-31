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

// checks if the OperationCredentialsCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationCredentialsCreateCommand{}

// OperationCredentialsCreateCommand struct for OperationCredentialsCreateCommand
type OperationCredentialsCreateCommand struct {
	Name               NullableString `json:"name,omitempty"`
	PrometheusUsername NullableString `json:"prometheusUsername,omitempty"`
	PrometheusPassword NullableString `json:"prometheusPassword,omitempty"`
	PrometheusUrl      NullableString `json:"prometheusUrl,omitempty"`
	OrganizationId     NullableInt32  `json:"organizationId,omitempty"`
}

// NewOperationCredentialsCreateCommand instantiates a new OperationCredentialsCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationCredentialsCreateCommand() *OperationCredentialsCreateCommand {
	this := OperationCredentialsCreateCommand{}
	return &this
}

// NewOperationCredentialsCreateCommandWithDefaults instantiates a new OperationCredentialsCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationCredentialsCreateCommandWithDefaults() *OperationCredentialsCreateCommand {
	this := OperationCredentialsCreateCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsCreateCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OperationCredentialsCreateCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OperationCredentialsCreateCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *OperationCredentialsCreateCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OperationCredentialsCreateCommand) UnsetName() {
	o.Name.Unset()
}

// GetPrometheusUsername returns the PrometheusUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsCreateCommand) GetPrometheusUsername() string {
	if o == nil || IsNil(o.PrometheusUsername.Get()) {
		var ret string
		return ret
	}
	return *o.PrometheusUsername.Get()
}

// GetPrometheusUsernameOk returns a tuple with the PrometheusUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsCreateCommand) GetPrometheusUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrometheusUsername.Get(), o.PrometheusUsername.IsSet()
}

// HasPrometheusUsername returns a boolean if a field has been set.
func (o *OperationCredentialsCreateCommand) HasPrometheusUsername() bool {
	if o != nil && o.PrometheusUsername.IsSet() {
		return true
	}

	return false
}

// SetPrometheusUsername gets a reference to the given NullableString and assigns it to the PrometheusUsername field.
func (o *OperationCredentialsCreateCommand) SetPrometheusUsername(v string) {
	o.PrometheusUsername.Set(&v)
}

// SetPrometheusUsernameNil sets the value for PrometheusUsername to be an explicit nil
func (o *OperationCredentialsCreateCommand) SetPrometheusUsernameNil() {
	o.PrometheusUsername.Set(nil)
}

// UnsetPrometheusUsername ensures that no value is present for PrometheusUsername, not even an explicit nil
func (o *OperationCredentialsCreateCommand) UnsetPrometheusUsername() {
	o.PrometheusUsername.Unset()
}

// GetPrometheusPassword returns the PrometheusPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsCreateCommand) GetPrometheusPassword() string {
	if o == nil || IsNil(o.PrometheusPassword.Get()) {
		var ret string
		return ret
	}
	return *o.PrometheusPassword.Get()
}

// GetPrometheusPasswordOk returns a tuple with the PrometheusPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsCreateCommand) GetPrometheusPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrometheusPassword.Get(), o.PrometheusPassword.IsSet()
}

// HasPrometheusPassword returns a boolean if a field has been set.
func (o *OperationCredentialsCreateCommand) HasPrometheusPassword() bool {
	if o != nil && o.PrometheusPassword.IsSet() {
		return true
	}

	return false
}

// SetPrometheusPassword gets a reference to the given NullableString and assigns it to the PrometheusPassword field.
func (o *OperationCredentialsCreateCommand) SetPrometheusPassword(v string) {
	o.PrometheusPassword.Set(&v)
}

// SetPrometheusPasswordNil sets the value for PrometheusPassword to be an explicit nil
func (o *OperationCredentialsCreateCommand) SetPrometheusPasswordNil() {
	o.PrometheusPassword.Set(nil)
}

// UnsetPrometheusPassword ensures that no value is present for PrometheusPassword, not even an explicit nil
func (o *OperationCredentialsCreateCommand) UnsetPrometheusPassword() {
	o.PrometheusPassword.Unset()
}

// GetPrometheusUrl returns the PrometheusUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsCreateCommand) GetPrometheusUrl() string {
	if o == nil || IsNil(o.PrometheusUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrometheusUrl.Get()
}

// GetPrometheusUrlOk returns a tuple with the PrometheusUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsCreateCommand) GetPrometheusUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrometheusUrl.Get(), o.PrometheusUrl.IsSet()
}

// HasPrometheusUrl returns a boolean if a field has been set.
func (o *OperationCredentialsCreateCommand) HasPrometheusUrl() bool {
	if o != nil && o.PrometheusUrl.IsSet() {
		return true
	}

	return false
}

// SetPrometheusUrl gets a reference to the given NullableString and assigns it to the PrometheusUrl field.
func (o *OperationCredentialsCreateCommand) SetPrometheusUrl(v string) {
	o.PrometheusUrl.Set(&v)
}

// SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil
func (o *OperationCredentialsCreateCommand) SetPrometheusUrlNil() {
	o.PrometheusUrl.Set(nil)
}

// UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
func (o *OperationCredentialsCreateCommand) UnsetPrometheusUrl() {
	o.PrometheusUrl.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsCreateCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsCreateCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OperationCredentialsCreateCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *OperationCredentialsCreateCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *OperationCredentialsCreateCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *OperationCredentialsCreateCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o OperationCredentialsCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationCredentialsCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PrometheusUsername.IsSet() {
		toSerialize["prometheusUsername"] = o.PrometheusUsername.Get()
	}
	if o.PrometheusPassword.IsSet() {
		toSerialize["prometheusPassword"] = o.PrometheusPassword.Get()
	}
	if o.PrometheusUrl.IsSet() {
		toSerialize["prometheusUrl"] = o.PrometheusUrl.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableOperationCredentialsCreateCommand struct {
	value *OperationCredentialsCreateCommand
	isSet bool
}

func (v NullableOperationCredentialsCreateCommand) Get() *OperationCredentialsCreateCommand {
	return v.value
}

func (v *NullableOperationCredentialsCreateCommand) Set(val *OperationCredentialsCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationCredentialsCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationCredentialsCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationCredentialsCreateCommand(val *OperationCredentialsCreateCommand) *NullableOperationCredentialsCreateCommand {
	return &NullableOperationCredentialsCreateCommand{value: val, isSet: true}
}

func (v NullableOperationCredentialsCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationCredentialsCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
