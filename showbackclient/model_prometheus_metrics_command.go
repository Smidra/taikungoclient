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
	"time"
)

// checks if the PrometheusMetricsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusMetricsCommand{}

// PrometheusMetricsCommand struct for PrometheusMetricsCommand
type PrometheusMetricsCommand struct {
	ProjectId      *int32         `json:"projectId,omitempty"`
	Parameters     NullableString `json:"parameters,omitempty"`
	Time           NullableTime   `json:"time,omitempty"`
	Start          NullableTime   `json:"start,omitempty"`
	End            NullableTime   `json:"end,omitempty"`
	IsGraphEnabled *bool          `json:"isGraphEnabled,omitempty"`
	IsAutoComplete *bool          `json:"isAutoComplete,omitempty"`
	Step           NullableString `json:"step,omitempty"`
}

// NewPrometheusMetricsCommand instantiates a new PrometheusMetricsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusMetricsCommand() *PrometheusMetricsCommand {
	this := PrometheusMetricsCommand{}
	return &this
}

// NewPrometheusMetricsCommandWithDefaults instantiates a new PrometheusMetricsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusMetricsCommandWithDefaults() *PrometheusMetricsCommand {
	this := PrometheusMetricsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PrometheusMetricsCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMetricsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PrometheusMetricsCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricsCommand) GetParameters() string {
	if o == nil || IsNil(o.Parameters.Get()) {
		var ret string
		return ret
	}
	return *o.Parameters.Get()
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricsCommand) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters.Get(), o.Parameters.IsSet()
}

// HasParameters returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasParameters() bool {
	if o != nil && o.Parameters.IsSet() {
		return true
	}

	return false
}

// SetParameters gets a reference to the given NullableString and assigns it to the Parameters field.
func (o *PrometheusMetricsCommand) SetParameters(v string) {
	o.Parameters.Set(&v)
}

// SetParametersNil sets the value for Parameters to be an explicit nil
func (o *PrometheusMetricsCommand) SetParametersNil() {
	o.Parameters.Set(nil)
}

// UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
func (o *PrometheusMetricsCommand) UnsetParameters() {
	o.Parameters.Unset()
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricsCommand) GetTime() time.Time {
	if o == nil || IsNil(o.Time.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricsCommand) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableTime and assigns it to the Time field.
func (o *PrometheusMetricsCommand) SetTime(v time.Time) {
	o.Time.Set(&v)
}

// SetTimeNil sets the value for Time to be an explicit nil
func (o *PrometheusMetricsCommand) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *PrometheusMetricsCommand) UnsetTime() {
	o.Time.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricsCommand) GetStart() time.Time {
	if o == nil || IsNil(o.Start.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricsCommand) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableTime and assigns it to the Start field.
func (o *PrometheusMetricsCommand) SetStart(v time.Time) {
	o.Start.Set(&v)
}

// SetStartNil sets the value for Start to be an explicit nil
func (o *PrometheusMetricsCommand) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *PrometheusMetricsCommand) UnsetStart() {
	o.Start.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricsCommand) GetEnd() time.Time {
	if o == nil || IsNil(o.End.Get()) {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricsCommand) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableTime and assigns it to the End field.
func (o *PrometheusMetricsCommand) SetEnd(v time.Time) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil
func (o *PrometheusMetricsCommand) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *PrometheusMetricsCommand) UnsetEnd() {
	o.End.Unset()
}

// GetIsGraphEnabled returns the IsGraphEnabled field value if set, zero value otherwise.
func (o *PrometheusMetricsCommand) GetIsGraphEnabled() bool {
	if o == nil || IsNil(o.IsGraphEnabled) {
		var ret bool
		return ret
	}
	return *o.IsGraphEnabled
}

// GetIsGraphEnabledOk returns a tuple with the IsGraphEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMetricsCommand) GetIsGraphEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGraphEnabled) {
		return nil, false
	}
	return o.IsGraphEnabled, true
}

// HasIsGraphEnabled returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasIsGraphEnabled() bool {
	if o != nil && !IsNil(o.IsGraphEnabled) {
		return true
	}

	return false
}

// SetIsGraphEnabled gets a reference to the given bool and assigns it to the IsGraphEnabled field.
func (o *PrometheusMetricsCommand) SetIsGraphEnabled(v bool) {
	o.IsGraphEnabled = &v
}

// GetIsAutoComplete returns the IsAutoComplete field value if set, zero value otherwise.
func (o *PrometheusMetricsCommand) GetIsAutoComplete() bool {
	if o == nil || IsNil(o.IsAutoComplete) {
		var ret bool
		return ret
	}
	return *o.IsAutoComplete
}

// GetIsAutoCompleteOk returns a tuple with the IsAutoComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMetricsCommand) GetIsAutoCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoComplete) {
		return nil, false
	}
	return o.IsAutoComplete, true
}

// HasIsAutoComplete returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasIsAutoComplete() bool {
	if o != nil && !IsNil(o.IsAutoComplete) {
		return true
	}

	return false
}

// SetIsAutoComplete gets a reference to the given bool and assigns it to the IsAutoComplete field.
func (o *PrometheusMetricsCommand) SetIsAutoComplete(v bool) {
	o.IsAutoComplete = &v
}

// GetStep returns the Step field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricsCommand) GetStep() string {
	if o == nil || IsNil(o.Step.Get()) {
		var ret string
		return ret
	}
	return *o.Step.Get()
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricsCommand) GetStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Step.Get(), o.Step.IsSet()
}

// HasStep returns a boolean if a field has been set.
func (o *PrometheusMetricsCommand) HasStep() bool {
	if o != nil && o.Step.IsSet() {
		return true
	}

	return false
}

// SetStep gets a reference to the given NullableString and assigns it to the Step field.
func (o *PrometheusMetricsCommand) SetStep(v string) {
	o.Step.Set(&v)
}

// SetStepNil sets the value for Step to be an explicit nil
func (o *PrometheusMetricsCommand) SetStepNil() {
	o.Step.Set(nil)
}

// UnsetStep ensures that no value is present for Step, not even an explicit nil
func (o *PrometheusMetricsCommand) UnsetStep() {
	o.Step.Unset()
}

func (o PrometheusMetricsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusMetricsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Parameters.IsSet() {
		toSerialize["parameters"] = o.Parameters.Get()
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if !IsNil(o.IsGraphEnabled) {
		toSerialize["isGraphEnabled"] = o.IsGraphEnabled
	}
	if !IsNil(o.IsAutoComplete) {
		toSerialize["isAutoComplete"] = o.IsAutoComplete
	}
	if o.Step.IsSet() {
		toSerialize["step"] = o.Step.Get()
	}
	return toSerialize, nil
}

type NullablePrometheusMetricsCommand struct {
	value *PrometheusMetricsCommand
	isSet bool
}

func (v NullablePrometheusMetricsCommand) Get() *PrometheusMetricsCommand {
	return v.value
}

func (v *NullablePrometheusMetricsCommand) Set(val *PrometheusMetricsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusMetricsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusMetricsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusMetricsCommand(val *PrometheusMetricsCommand) *NullablePrometheusMetricsCommand {
	return &NullablePrometheusMetricsCommand{value: val, isSet: true}
}

func (v NullablePrometheusMetricsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusMetricsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}