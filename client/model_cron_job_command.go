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

// checks if the CronJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronJobCommand{}

// CronJobCommand struct for CronJobCommand
type CronJobCommand struct {
	CronPeriod NullableString `json:"cronPeriod,omitempty"`
}

// NewCronJobCommand instantiates a new CronJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronJobCommand() *CronJobCommand {
	this := CronJobCommand{}
	return &this
}

// NewCronJobCommandWithDefaults instantiates a new CronJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronJobCommandWithDefaults() *CronJobCommand {
	this := CronJobCommand{}
	return &this
}

// GetCronPeriod returns the CronPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CronJobCommand) GetCronPeriod() string {
	if o == nil || IsNil(o.CronPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.CronPeriod.Get()
}

// GetCronPeriodOk returns a tuple with the CronPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CronJobCommand) GetCronPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CronPeriod.Get(), o.CronPeriod.IsSet()
}

// HasCronPeriod returns a boolean if a field has been set.
func (o *CronJobCommand) HasCronPeriod() bool {
	if o != nil && o.CronPeriod.IsSet() {
		return true
	}

	return false
}

// SetCronPeriod gets a reference to the given NullableString and assigns it to the CronPeriod field.
func (o *CronJobCommand) SetCronPeriod(v string) {
	o.CronPeriod.Set(&v)
}

// SetCronPeriodNil sets the value for CronPeriod to be an explicit nil
func (o *CronJobCommand) SetCronPeriodNil() {
	o.CronPeriod.Set(nil)
}

// UnsetCronPeriod ensures that no value is present for CronPeriod, not even an explicit nil
func (o *CronJobCommand) UnsetCronPeriod() {
	o.CronPeriod.Unset()
}

func (o CronJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronJobCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CronPeriod.IsSet() {
		toSerialize["cronPeriod"] = o.CronPeriod.Get()
	}
	return toSerialize, nil
}

type NullableCronJobCommand struct {
	value *CronJobCommand
	isSet bool
}

func (v NullableCronJobCommand) Get() *CronJobCommand {
	return v.value
}

func (v *NullableCronJobCommand) Set(val *CronJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCronJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCronJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronJobCommand(val *CronJobCommand) *NullableCronJobCommand {
	return &NullableCronJobCommand{value: val, isSet: true}
}

func (v NullableCronJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
