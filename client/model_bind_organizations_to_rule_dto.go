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

// checks if the BindOrganizationsToRuleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindOrganizationsToRuleDto{}

// BindOrganizationsToRuleDto struct for BindOrganizationsToRuleDto
type BindOrganizationsToRuleDto struct {
	OrganizationId   *int32          `json:"organizationId,omitempty"`
	OrganizationName NullableString  `json:"organizationName,omitempty"`
	RuleDiscountRate NullableFloat64 `json:"ruleDiscountRate,omitempty"`
	IsBound          *bool           `json:"isBound,omitempty"`
}

// NewBindOrganizationsToRuleDto instantiates a new BindOrganizationsToRuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindOrganizationsToRuleDto() *BindOrganizationsToRuleDto {
	this := BindOrganizationsToRuleDto{}
	return &this
}

// NewBindOrganizationsToRuleDtoWithDefaults instantiates a new BindOrganizationsToRuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindOrganizationsToRuleDtoWithDefaults() *BindOrganizationsToRuleDto {
	this := BindOrganizationsToRuleDto{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *BindOrganizationsToRuleDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindOrganizationsToRuleDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *BindOrganizationsToRuleDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *BindOrganizationsToRuleDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindOrganizationsToRuleDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindOrganizationsToRuleDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *BindOrganizationsToRuleDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *BindOrganizationsToRuleDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *BindOrganizationsToRuleDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *BindOrganizationsToRuleDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetRuleDiscountRate returns the RuleDiscountRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindOrganizationsToRuleDto) GetRuleDiscountRate() float64 {
	if o == nil || IsNil(o.RuleDiscountRate.Get()) {
		var ret float64
		return ret
	}
	return *o.RuleDiscountRate.Get()
}

// GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindOrganizationsToRuleDto) GetRuleDiscountRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleDiscountRate.Get(), o.RuleDiscountRate.IsSet()
}

// HasRuleDiscountRate returns a boolean if a field has been set.
func (o *BindOrganizationsToRuleDto) HasRuleDiscountRate() bool {
	if o != nil && o.RuleDiscountRate.IsSet() {
		return true
	}

	return false
}

// SetRuleDiscountRate gets a reference to the given NullableFloat64 and assigns it to the RuleDiscountRate field.
func (o *BindOrganizationsToRuleDto) SetRuleDiscountRate(v float64) {
	o.RuleDiscountRate.Set(&v)
}

// SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil
func (o *BindOrganizationsToRuleDto) SetRuleDiscountRateNil() {
	o.RuleDiscountRate.Set(nil)
}

// UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil
func (o *BindOrganizationsToRuleDto) UnsetRuleDiscountRate() {
	o.RuleDiscountRate.Unset()
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *BindOrganizationsToRuleDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindOrganizationsToRuleDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *BindOrganizationsToRuleDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *BindOrganizationsToRuleDto) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o BindOrganizationsToRuleDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindOrganizationsToRuleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.RuleDiscountRate.IsSet() {
		toSerialize["ruleDiscountRate"] = o.RuleDiscountRate.Get()
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableBindOrganizationsToRuleDto struct {
	value *BindOrganizationsToRuleDto
	isSet bool
}

func (v NullableBindOrganizationsToRuleDto) Get() *BindOrganizationsToRuleDto {
	return v.value
}

func (v *NullableBindOrganizationsToRuleDto) Set(val *BindOrganizationsToRuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBindOrganizationsToRuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBindOrganizationsToRuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindOrganizationsToRuleDto(val *BindOrganizationsToRuleDto) *NullableBindOrganizationsToRuleDto {
	return &NullableBindOrganizationsToRuleDto{value: val, isSet: true}
}

func (v NullableBindOrganizationsToRuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindOrganizationsToRuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
