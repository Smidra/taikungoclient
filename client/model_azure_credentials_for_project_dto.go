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

// checks if the AzureCredentialsForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCredentialsForProjectDto{}

// AzureCredentialsForProjectDto struct for AzureCredentialsForProjectDto
type AzureCredentialsForProjectDto struct {
	AzureClientId       NullableString `json:"azureClientId,omitempty"`
	AzureClientSecret   NullableString `json:"azureClientSecret,omitempty"`
	AzureSubscriptionId NullableString `json:"azureSubscriptionId,omitempty"`
	AzureTenantId       NullableString `json:"azureTenantId,omitempty"`
	AzureLocation       NullableString `json:"azureLocation,omitempty"`
}

// NewAzureCredentialsForProjectDto instantiates a new AzureCredentialsForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCredentialsForProjectDto() *AzureCredentialsForProjectDto {
	this := AzureCredentialsForProjectDto{}
	return &this
}

// NewAzureCredentialsForProjectDtoWithDefaults instantiates a new AzureCredentialsForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCredentialsForProjectDtoWithDefaults() *AzureCredentialsForProjectDto {
	this := AzureCredentialsForProjectDto{}
	return &this
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsForProjectDto) GetAzureClientId() string {
	if o == nil || IsNil(o.AzureClientId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientId.Get()
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsForProjectDto) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientId.Get(), o.AzureClientId.IsSet()
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *AzureCredentialsForProjectDto) HasAzureClientId() bool {
	if o != nil && o.AzureClientId.IsSet() {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given NullableString and assigns it to the AzureClientId field.
func (o *AzureCredentialsForProjectDto) SetAzureClientId(v string) {
	o.AzureClientId.Set(&v)
}

// SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil
func (o *AzureCredentialsForProjectDto) SetAzureClientIdNil() {
	o.AzureClientId.Set(nil)
}

// UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
func (o *AzureCredentialsForProjectDto) UnsetAzureClientId() {
	o.AzureClientId.Unset()
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsForProjectDto) GetAzureClientSecret() string {
	if o == nil || IsNil(o.AzureClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientSecret.Get()
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsForProjectDto) GetAzureClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientSecret.Get(), o.AzureClientSecret.IsSet()
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *AzureCredentialsForProjectDto) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret.IsSet() {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given NullableString and assigns it to the AzureClientSecret field.
func (o *AzureCredentialsForProjectDto) SetAzureClientSecret(v string) {
	o.AzureClientSecret.Set(&v)
}

// SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil
func (o *AzureCredentialsForProjectDto) SetAzureClientSecretNil() {
	o.AzureClientSecret.Set(nil)
}

// UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
func (o *AzureCredentialsForProjectDto) UnsetAzureClientSecret() {
	o.AzureClientSecret.Unset()
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsForProjectDto) GetAzureSubscriptionId() string {
	if o == nil || IsNil(o.AzureSubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId.Get()
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsForProjectDto) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureSubscriptionId.Get(), o.AzureSubscriptionId.IsSet()
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *AzureCredentialsForProjectDto) HasAzureSubscriptionId() bool {
	if o != nil && o.AzureSubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given NullableString and assigns it to the AzureSubscriptionId field.
func (o *AzureCredentialsForProjectDto) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId.Set(&v)
}

// SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil
func (o *AzureCredentialsForProjectDto) SetAzureSubscriptionIdNil() {
	o.AzureSubscriptionId.Set(nil)
}

// UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
func (o *AzureCredentialsForProjectDto) UnsetAzureSubscriptionId() {
	o.AzureSubscriptionId.Unset()
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsForProjectDto) GetAzureTenantId() string {
	if o == nil || IsNil(o.AzureTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureTenantId.Get()
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsForProjectDto) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureTenantId.Get(), o.AzureTenantId.IsSet()
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *AzureCredentialsForProjectDto) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId.IsSet() {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given NullableString and assigns it to the AzureTenantId field.
func (o *AzureCredentialsForProjectDto) SetAzureTenantId(v string) {
	o.AzureTenantId.Set(&v)
}

// SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil
func (o *AzureCredentialsForProjectDto) SetAzureTenantIdNil() {
	o.AzureTenantId.Set(nil)
}

// UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
func (o *AzureCredentialsForProjectDto) UnsetAzureTenantId() {
	o.AzureTenantId.Unset()
}

// GetAzureLocation returns the AzureLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsForProjectDto) GetAzureLocation() string {
	if o == nil || IsNil(o.AzureLocation.Get()) {
		var ret string
		return ret
	}
	return *o.AzureLocation.Get()
}

// GetAzureLocationOk returns a tuple with the AzureLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsForProjectDto) GetAzureLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureLocation.Get(), o.AzureLocation.IsSet()
}

// HasAzureLocation returns a boolean if a field has been set.
func (o *AzureCredentialsForProjectDto) HasAzureLocation() bool {
	if o != nil && o.AzureLocation.IsSet() {
		return true
	}

	return false
}

// SetAzureLocation gets a reference to the given NullableString and assigns it to the AzureLocation field.
func (o *AzureCredentialsForProjectDto) SetAzureLocation(v string) {
	o.AzureLocation.Set(&v)
}

// SetAzureLocationNil sets the value for AzureLocation to be an explicit nil
func (o *AzureCredentialsForProjectDto) SetAzureLocationNil() {
	o.AzureLocation.Set(nil)
}

// UnsetAzureLocation ensures that no value is present for AzureLocation, not even an explicit nil
func (o *AzureCredentialsForProjectDto) UnsetAzureLocation() {
	o.AzureLocation.Unset()
}

func (o AzureCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCredentialsForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureClientId.IsSet() {
		toSerialize["azureClientId"] = o.AzureClientId.Get()
	}
	if o.AzureClientSecret.IsSet() {
		toSerialize["azureClientSecret"] = o.AzureClientSecret.Get()
	}
	if o.AzureSubscriptionId.IsSet() {
		toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId.Get()
	}
	if o.AzureTenantId.IsSet() {
		toSerialize["azureTenantId"] = o.AzureTenantId.Get()
	}
	if o.AzureLocation.IsSet() {
		toSerialize["azureLocation"] = o.AzureLocation.Get()
	}
	return toSerialize, nil
}

type NullableAzureCredentialsForProjectDto struct {
	value *AzureCredentialsForProjectDto
	isSet bool
}

func (v NullableAzureCredentialsForProjectDto) Get() *AzureCredentialsForProjectDto {
	return v.value
}

func (v *NullableAzureCredentialsForProjectDto) Set(val *AzureCredentialsForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCredentialsForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCredentialsForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCredentialsForProjectDto(val *AzureCredentialsForProjectDto) *NullableAzureCredentialsForProjectDto {
	return &NullableAzureCredentialsForProjectDto{value: val, isSet: true}
}

func (v NullableAzureCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCredentialsForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
