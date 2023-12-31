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

// checks if the BindRulesToOrganizationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindRulesToOrganizationDto{}

// BindRulesToOrganizationDto struct for BindRulesToOrganizationDto
type BindRulesToOrganizationDto struct {
	Id           *int32          `json:"id,omitempty"`
	Name         NullableString  `json:"name,omitempty"`
	DiscountRate NullableFloat64 `json:"discountRate,omitempty"`
	IsBound      *bool           `json:"isBound,omitempty"`
}

// NewBindRulesToOrganizationDto instantiates a new BindRulesToOrganizationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindRulesToOrganizationDto() *BindRulesToOrganizationDto {
	this := BindRulesToOrganizationDto{}
	return &this
}

// NewBindRulesToOrganizationDtoWithDefaults instantiates a new BindRulesToOrganizationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindRulesToOrganizationDtoWithDefaults() *BindRulesToOrganizationDto {
	this := BindRulesToOrganizationDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BindRulesToOrganizationDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindRulesToOrganizationDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BindRulesToOrganizationDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BindRulesToOrganizationDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindRulesToOrganizationDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindRulesToOrganizationDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BindRulesToOrganizationDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BindRulesToOrganizationDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BindRulesToOrganizationDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BindRulesToOrganizationDto) UnsetName() {
	o.Name.Unset()
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindRulesToOrganizationDto) GetDiscountRate() float64 {
	if o == nil || IsNil(o.DiscountRate.Get()) {
		var ret float64
		return ret
	}
	return *o.DiscountRate.Get()
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindRulesToOrganizationDto) GetDiscountRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountRate.Get(), o.DiscountRate.IsSet()
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *BindRulesToOrganizationDto) HasDiscountRate() bool {
	if o != nil && o.DiscountRate.IsSet() {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given NullableFloat64 and assigns it to the DiscountRate field.
func (o *BindRulesToOrganizationDto) SetDiscountRate(v float64) {
	o.DiscountRate.Set(&v)
}

// SetDiscountRateNil sets the value for DiscountRate to be an explicit nil
func (o *BindRulesToOrganizationDto) SetDiscountRateNil() {
	o.DiscountRate.Set(nil)
}

// UnsetDiscountRate ensures that no value is present for DiscountRate, not even an explicit nil
func (o *BindRulesToOrganizationDto) UnsetDiscountRate() {
	o.DiscountRate.Unset()
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *BindRulesToOrganizationDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindRulesToOrganizationDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *BindRulesToOrganizationDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *BindRulesToOrganizationDto) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o BindRulesToOrganizationDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindRulesToOrganizationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DiscountRate.IsSet() {
		toSerialize["discountRate"] = o.DiscountRate.Get()
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableBindRulesToOrganizationDto struct {
	value *BindRulesToOrganizationDto
	isSet bool
}

func (v NullableBindRulesToOrganizationDto) Get() *BindRulesToOrganizationDto {
	return v.value
}

func (v *NullableBindRulesToOrganizationDto) Set(val *BindRulesToOrganizationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBindRulesToOrganizationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBindRulesToOrganizationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindRulesToOrganizationDto(val *BindRulesToOrganizationDto) *NullableBindRulesToOrganizationDto {
	return &NullableBindRulesToOrganizationDto{value: val, isSet: true}
}

func (v NullableBindRulesToOrganizationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindRulesToOrganizationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
