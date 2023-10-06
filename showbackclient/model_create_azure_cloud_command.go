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

// checks if the CreateAzureCloudCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzureCloudCommand{}

// CreateAzureCloudCommand struct for CreateAzureCloudCommand
type CreateAzureCloudCommand struct {
	Name                NullableString `json:"name,omitempty"`
	AzureSubscriptionId NullableString `json:"azureSubscriptionId,omitempty"`
	AzureClientId       NullableString `json:"azureClientId,omitempty"`
	AzureClientSecret   NullableString `json:"azureClientSecret,omitempty"`
	AzureTenantId       NullableString `json:"azureTenantId,omitempty"`
	AzureLocation       NullableString `json:"azureLocation,omitempty"`
	AzCount             *int32         `json:"azCount,omitempty"`
	OrganizationId      NullableInt32  `json:"organizationId,omitempty"`
}

// NewCreateAzureCloudCommand instantiates a new CreateAzureCloudCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzureCloudCommand() *CreateAzureCloudCommand {
	this := CreateAzureCloudCommand{}
	return &this
}

// NewCreateAzureCloudCommandWithDefaults instantiates a new CreateAzureCloudCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzureCloudCommandWithDefaults() *CreateAzureCloudCommand {
	this := CreateAzureCloudCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateAzureCloudCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateAzureCloudCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetName() {
	o.Name.Unset()
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetAzureSubscriptionId() string {
	if o == nil || IsNil(o.AzureSubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId.Get()
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureSubscriptionId.Get(), o.AzureSubscriptionId.IsSet()
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzureSubscriptionId() bool {
	if o != nil && o.AzureSubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given NullableString and assigns it to the AzureSubscriptionId field.
func (o *CreateAzureCloudCommand) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId.Set(&v)
}

// SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil
func (o *CreateAzureCloudCommand) SetAzureSubscriptionIdNil() {
	o.AzureSubscriptionId.Set(nil)
}

// UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetAzureSubscriptionId() {
	o.AzureSubscriptionId.Unset()
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetAzureClientId() string {
	if o == nil || IsNil(o.AzureClientId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientId.Get()
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientId.Get(), o.AzureClientId.IsSet()
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzureClientId() bool {
	if o != nil && o.AzureClientId.IsSet() {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given NullableString and assigns it to the AzureClientId field.
func (o *CreateAzureCloudCommand) SetAzureClientId(v string) {
	o.AzureClientId.Set(&v)
}

// SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil
func (o *CreateAzureCloudCommand) SetAzureClientIdNil() {
	o.AzureClientId.Set(nil)
}

// UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetAzureClientId() {
	o.AzureClientId.Unset()
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetAzureClientSecret() string {
	if o == nil || IsNil(o.AzureClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientSecret.Get()
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetAzureClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientSecret.Get(), o.AzureClientSecret.IsSet()
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret.IsSet() {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given NullableString and assigns it to the AzureClientSecret field.
func (o *CreateAzureCloudCommand) SetAzureClientSecret(v string) {
	o.AzureClientSecret.Set(&v)
}

// SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil
func (o *CreateAzureCloudCommand) SetAzureClientSecretNil() {
	o.AzureClientSecret.Set(nil)
}

// UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetAzureClientSecret() {
	o.AzureClientSecret.Unset()
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetAzureTenantId() string {
	if o == nil || IsNil(o.AzureTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureTenantId.Get()
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureTenantId.Get(), o.AzureTenantId.IsSet()
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId.IsSet() {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given NullableString and assigns it to the AzureTenantId field.
func (o *CreateAzureCloudCommand) SetAzureTenantId(v string) {
	o.AzureTenantId.Set(&v)
}

// SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil
func (o *CreateAzureCloudCommand) SetAzureTenantIdNil() {
	o.AzureTenantId.Set(nil)
}

// UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetAzureTenantId() {
	o.AzureTenantId.Unset()
}

// GetAzureLocation returns the AzureLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetAzureLocation() string {
	if o == nil || IsNil(o.AzureLocation.Get()) {
		var ret string
		return ret
	}
	return *o.AzureLocation.Get()
}

// GetAzureLocationOk returns a tuple with the AzureLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetAzureLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureLocation.Get(), o.AzureLocation.IsSet()
}

// HasAzureLocation returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzureLocation() bool {
	if o != nil && o.AzureLocation.IsSet() {
		return true
	}

	return false
}

// SetAzureLocation gets a reference to the given NullableString and assigns it to the AzureLocation field.
func (o *CreateAzureCloudCommand) SetAzureLocation(v string) {
	o.AzureLocation.Set(&v)
}

// SetAzureLocationNil sets the value for AzureLocation to be an explicit nil
func (o *CreateAzureCloudCommand) SetAzureLocationNil() {
	o.AzureLocation.Set(nil)
}

// UnsetAzureLocation ensures that no value is present for AzureLocation, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetAzureLocation() {
	o.AzureLocation.Unset()
}

// GetAzCount returns the AzCount field value if set, zero value otherwise.
func (o *CreateAzureCloudCommand) GetAzCount() int32 {
	if o == nil || IsNil(o.AzCount) {
		var ret int32
		return ret
	}
	return *o.AzCount
}

// GetAzCountOk returns a tuple with the AzCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureCloudCommand) GetAzCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AzCount) {
		return nil, false
	}
	return o.AzCount, true
}

// HasAzCount returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasAzCount() bool {
	if o != nil && !IsNil(o.AzCount) {
		return true
	}

	return false
}

// SetAzCount gets a reference to the given int32 and assigns it to the AzCount field.
func (o *CreateAzureCloudCommand) SetAzCount(v int32) {
	o.AzCount = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAzureCloudCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAzureCloudCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateAzureCloudCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateAzureCloudCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateAzureCloudCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateAzureCloudCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o CreateAzureCloudCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAzureCloudCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AzureSubscriptionId.IsSet() {
		toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId.Get()
	}
	if o.AzureClientId.IsSet() {
		toSerialize["azureClientId"] = o.AzureClientId.Get()
	}
	if o.AzureClientSecret.IsSet() {
		toSerialize["azureClientSecret"] = o.AzureClientSecret.Get()
	}
	if o.AzureTenantId.IsSet() {
		toSerialize["azureTenantId"] = o.AzureTenantId.Get()
	}
	if o.AzureLocation.IsSet() {
		toSerialize["azureLocation"] = o.AzureLocation.Get()
	}
	if !IsNil(o.AzCount) {
		toSerialize["azCount"] = o.AzCount
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableCreateAzureCloudCommand struct {
	value *CreateAzureCloudCommand
	isSet bool
}

func (v NullableCreateAzureCloudCommand) Get() *CreateAzureCloudCommand {
	return v.value
}

func (v *NullableCreateAzureCloudCommand) Set(val *CreateAzureCloudCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAzureCloudCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAzureCloudCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAzureCloudCommand(val *CreateAzureCloudCommand) *NullableCreateAzureCloudCommand {
	return &NullableCreateAzureCloudCommand{value: val, isSet: true}
}

func (v NullableCreateAzureCloudCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAzureCloudCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}