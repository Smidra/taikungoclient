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

// checks if the AzureSubscriptionListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureSubscriptionListCommand{}

// AzureSubscriptionListCommand struct for AzureSubscriptionListCommand
type AzureSubscriptionListCommand struct {
	ClientId     NullableString `json:"clientId,omitempty"`
	ClientSecret NullableString `json:"clientSecret,omitempty"`
	TenantId     NullableString `json:"tenantId,omitempty"`
}

// NewAzureSubscriptionListCommand instantiates a new AzureSubscriptionListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureSubscriptionListCommand() *AzureSubscriptionListCommand {
	this := AzureSubscriptionListCommand{}
	return &this
}

// NewAzureSubscriptionListCommandWithDefaults instantiates a new AzureSubscriptionListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureSubscriptionListCommandWithDefaults() *AzureSubscriptionListCommand {
	this := AzureSubscriptionListCommand{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSubscriptionListCommand) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSubscriptionListCommand) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *AzureSubscriptionListCommand) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *AzureSubscriptionListCommand) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *AzureSubscriptionListCommand) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *AzureSubscriptionListCommand) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSubscriptionListCommand) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSubscriptionListCommand) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AzureSubscriptionListCommand) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *AzureSubscriptionListCommand) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *AzureSubscriptionListCommand) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *AzureSubscriptionListCommand) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSubscriptionListCommand) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSubscriptionListCommand) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *AzureSubscriptionListCommand) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *AzureSubscriptionListCommand) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *AzureSubscriptionListCommand) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *AzureSubscriptionListCommand) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o AzureSubscriptionListCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureSubscriptionListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableAzureSubscriptionListCommand struct {
	value *AzureSubscriptionListCommand
	isSet bool
}

func (v NullableAzureSubscriptionListCommand) Get() *AzureSubscriptionListCommand {
	return v.value
}

func (v *NullableAzureSubscriptionListCommand) Set(val *AzureSubscriptionListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureSubscriptionListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureSubscriptionListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureSubscriptionListCommand(val *AzureSubscriptionListCommand) *NullableAzureSubscriptionListCommand {
	return &NullableAzureSubscriptionListCommand{value: val, isSet: true}
}

func (v NullableAzureSubscriptionListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureSubscriptionListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
