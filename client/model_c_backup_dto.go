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
	"time"
)

// checks if the CBackupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CBackupDto{}

// CBackupDto struct for CBackupDto
type CBackupDto struct {
	MetadataName NullableString `json:"metadataName,omitempty"`
	CreatedAt    NullableTime   `json:"createdAt,omitempty"`
	Expiration   NullableTime   `json:"expiration,omitempty"`
	ScheduleName NullableString `json:"scheduleName,omitempty"`
	Namespace    NullableString `json:"namespace,omitempty"`
	Location     NullableString `json:"location,omitempty"`
	Phase        NullableString `json:"phase,omitempty"`
}

// NewCBackupDto instantiates a new CBackupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCBackupDto() *CBackupDto {
	this := CBackupDto{}
	return &this
}

// NewCBackupDtoWithDefaults instantiates a new CBackupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCBackupDtoWithDefaults() *CBackupDto {
	this := CBackupDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *CBackupDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *CBackupDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}

// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *CBackupDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *CBackupDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CBackupDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *CBackupDto) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CBackupDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CBackupDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *CBackupDto) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *CBackupDto) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *CBackupDto) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *CBackupDto) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetScheduleName() string {
	if o == nil || IsNil(o.ScheduleName.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduleName.Get()
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetScheduleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleName.Get(), o.ScheduleName.IsSet()
}

// HasScheduleName returns a boolean if a field has been set.
func (o *CBackupDto) HasScheduleName() bool {
	if o != nil && o.ScheduleName.IsSet() {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given NullableString and assigns it to the ScheduleName field.
func (o *CBackupDto) SetScheduleName(v string) {
	o.ScheduleName.Set(&v)
}

// SetScheduleNameNil sets the value for ScheduleName to be an explicit nil
func (o *CBackupDto) SetScheduleNameNil() {
	o.ScheduleName.Set(nil)
}

// UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
func (o *CBackupDto) UnsetScheduleName() {
	o.ScheduleName.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *CBackupDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *CBackupDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *CBackupDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *CBackupDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *CBackupDto) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *CBackupDto) SetLocation(v string) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *CBackupDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *CBackupDto) UnsetLocation() {
	o.Location.Unset()
}

// GetPhase returns the Phase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CBackupDto) GetPhase() string {
	if o == nil || IsNil(o.Phase.Get()) {
		var ret string
		return ret
	}
	return *o.Phase.Get()
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CBackupDto) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase.Get(), o.Phase.IsSet()
}

// HasPhase returns a boolean if a field has been set.
func (o *CBackupDto) HasPhase() bool {
	if o != nil && o.Phase.IsSet() {
		return true
	}

	return false
}

// SetPhase gets a reference to the given NullableString and assigns it to the Phase field.
func (o *CBackupDto) SetPhase(v string) {
	o.Phase.Set(&v)
}

// SetPhaseNil sets the value for Phase to be an explicit nil
func (o *CBackupDto) SetPhaseNil() {
	o.Phase.Set(nil)
}

// UnsetPhase ensures that no value is present for Phase, not even an explicit nil
func (o *CBackupDto) UnsetPhase() {
	o.Phase.Unset()
}

func (o CBackupDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CBackupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.ScheduleName.IsSet() {
		toSerialize["scheduleName"] = o.ScheduleName.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Phase.IsSet() {
		toSerialize["phase"] = o.Phase.Get()
	}
	return toSerialize, nil
}

type NullableCBackupDto struct {
	value *CBackupDto
	isSet bool
}

func (v NullableCBackupDto) Get() *CBackupDto {
	return v.value
}

func (v *NullableCBackupDto) Set(val *CBackupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCBackupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCBackupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCBackupDto(val *CBackupDto) *NullableCBackupDto {
	return &NullableCBackupDto{value: val, isSet: true}
}

func (v NullableCBackupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCBackupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
