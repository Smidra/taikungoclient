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

// checks if the PrometheusBillingCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusBillingCreateCommand{}

// PrometheusBillingCreateCommand struct for PrometheusBillingCreateCommand
type PrometheusBillingCreateCommand struct {
	OrganizationId   *int32     `json:"organizationId,omitempty"`
	PrometheusRuleId *int32     `json:"prometheusRuleId,omitempty"`
	StartDate        *time.Time `json:"startDate,omitempty"`
	Price            *float64   `json:"price,omitempty"`
}

// NewPrometheusBillingCreateCommand instantiates a new PrometheusBillingCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusBillingCreateCommand() *PrometheusBillingCreateCommand {
	this := PrometheusBillingCreateCommand{}
	return &this
}

// NewPrometheusBillingCreateCommandWithDefaults instantiates a new PrometheusBillingCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusBillingCreateCommandWithDefaults() *PrometheusBillingCreateCommand {
	this := PrometheusBillingCreateCommand{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PrometheusBillingCreateCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingCreateCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PrometheusBillingCreateCommand) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *PrometheusBillingCreateCommand) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetPrometheusRuleId returns the PrometheusRuleId field value if set, zero value otherwise.
func (o *PrometheusBillingCreateCommand) GetPrometheusRuleId() int32 {
	if o == nil || IsNil(o.PrometheusRuleId) {
		var ret int32
		return ret
	}
	return *o.PrometheusRuleId
}

// GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingCreateCommand) GetPrometheusRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PrometheusRuleId) {
		return nil, false
	}
	return o.PrometheusRuleId, true
}

// HasPrometheusRuleId returns a boolean if a field has been set.
func (o *PrometheusBillingCreateCommand) HasPrometheusRuleId() bool {
	if o != nil && !IsNil(o.PrometheusRuleId) {
		return true
	}

	return false
}

// SetPrometheusRuleId gets a reference to the given int32 and assigns it to the PrometheusRuleId field.
func (o *PrometheusBillingCreateCommand) SetPrometheusRuleId(v int32) {
	o.PrometheusRuleId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PrometheusBillingCreateCommand) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingCreateCommand) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PrometheusBillingCreateCommand) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *PrometheusBillingCreateCommand) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PrometheusBillingCreateCommand) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingCreateCommand) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PrometheusBillingCreateCommand) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PrometheusBillingCreateCommand) SetPrice(v float64) {
	o.Price = &v
}

func (o PrometheusBillingCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusBillingCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.PrometheusRuleId) {
		toSerialize["prometheusRuleId"] = o.PrometheusRuleId
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullablePrometheusBillingCreateCommand struct {
	value *PrometheusBillingCreateCommand
	isSet bool
}

func (v NullablePrometheusBillingCreateCommand) Get() *PrometheusBillingCreateCommand {
	return v.value
}

func (v *NullablePrometheusBillingCreateCommand) Set(val *PrometheusBillingCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusBillingCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusBillingCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusBillingCreateCommand(val *PrometheusBillingCreateCommand) *NullablePrometheusBillingCreateCommand {
	return &NullablePrometheusBillingCreateCommand{value: val, isSet: true}
}

func (v NullablePrometheusBillingCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusBillingCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
