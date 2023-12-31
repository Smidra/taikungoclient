/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the GroupedShowbackSummariesByProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedShowbackSummariesByProjectDto{}

// GroupedShowbackSummariesByProjectDto struct for GroupedShowbackSummariesByProjectDto
type GroupedShowbackSummariesByProjectDto struct {
	ProjectName NullableString  `json:"projectName,omitempty"`
	TotalPrice  NullableFloat64 `json:"totalPrice,omitempty"`
}

// NewGroupedShowbackSummariesByProjectDto instantiates a new GroupedShowbackSummariesByProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedShowbackSummariesByProjectDto() *GroupedShowbackSummariesByProjectDto {
	this := GroupedShowbackSummariesByProjectDto{}
	return &this
}

// NewGroupedShowbackSummariesByProjectDtoWithDefaults instantiates a new GroupedShowbackSummariesByProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedShowbackSummariesByProjectDtoWithDefaults() *GroupedShowbackSummariesByProjectDto {
	this := GroupedShowbackSummariesByProjectDto{}
	return &this
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedShowbackSummariesByProjectDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedShowbackSummariesByProjectDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *GroupedShowbackSummariesByProjectDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *GroupedShowbackSummariesByProjectDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *GroupedShowbackSummariesByProjectDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *GroupedShowbackSummariesByProjectDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedShowbackSummariesByProjectDto) GetTotalPrice() float64 {
	if o == nil || IsNil(o.TotalPrice.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalPrice.Get()
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedShowbackSummariesByProjectDto) GetTotalPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPrice.Get(), o.TotalPrice.IsSet()
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *GroupedShowbackSummariesByProjectDto) HasTotalPrice() bool {
	if o != nil && o.TotalPrice.IsSet() {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given NullableFloat64 and assigns it to the TotalPrice field.
func (o *GroupedShowbackSummariesByProjectDto) SetTotalPrice(v float64) {
	o.TotalPrice.Set(&v)
}

// SetTotalPriceNil sets the value for TotalPrice to be an explicit nil
func (o *GroupedShowbackSummariesByProjectDto) SetTotalPriceNil() {
	o.TotalPrice.Set(nil)
}

// UnsetTotalPrice ensures that no value is present for TotalPrice, not even an explicit nil
func (o *GroupedShowbackSummariesByProjectDto) UnsetTotalPrice() {
	o.TotalPrice.Unset()
}

func (o GroupedShowbackSummariesByProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedShowbackSummariesByProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if o.TotalPrice.IsSet() {
		toSerialize["totalPrice"] = o.TotalPrice.Get()
	}
	return toSerialize, nil
}

type NullableGroupedShowbackSummariesByProjectDto struct {
	value *GroupedShowbackSummariesByProjectDto
	isSet bool
}

func (v NullableGroupedShowbackSummariesByProjectDto) Get() *GroupedShowbackSummariesByProjectDto {
	return v.value
}

func (v *NullableGroupedShowbackSummariesByProjectDto) Set(val *GroupedShowbackSummariesByProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedShowbackSummariesByProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedShowbackSummariesByProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedShowbackSummariesByProjectDto(val *GroupedShowbackSummariesByProjectDto) *NullableGroupedShowbackSummariesByProjectDto {
	return &NullableGroupedShowbackSummariesByProjectDto{value: val, isSet: true}
}

func (v NullableGroupedShowbackSummariesByProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedShowbackSummariesByProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
