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

// checks if the Parameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parameter{}

// Parameter struct for Parameter
type Parameter struct {
	Label    NullableString `json:"label,omitempty"`
	Operator NullableString `json:"operator,omitempty"`
	Value    NullableString `json:"value,omitempty"`
}

// NewParameter instantiates a new Parameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameter() *Parameter {
	this := Parameter{}
	return &this
}

// NewParameterWithDefaults instantiates a new Parameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterWithDefaults() *Parameter {
	this := Parameter{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Parameter) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Parameter) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *Parameter) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *Parameter) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil
func (o *Parameter) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *Parameter) UnsetLabel() {
	o.Label.Unset()
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Parameter) GetOperator() string {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Parameter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *Parameter) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *Parameter) SetOperator(v string) {
	o.Operator.Set(&v)
}

// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *Parameter) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *Parameter) UnsetOperator() {
	o.Operator.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Parameter) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Parameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Parameter) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *Parameter) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *Parameter) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Parameter) UnsetValue() {
	o.Value.Unset()
}

func (o Parameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableParameter struct {
	value *Parameter
	isSet bool
}

func (v NullableParameter) Get() *Parameter {
	return v.value
}

func (v *NullableParameter) Set(val *Parameter) {
	v.value = val
	v.isSet = true
}

func (v NullableParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameter(val *Parameter) *NullableParameter {
	return &NullableParameter{value: val, isSet: true}
}

func (v NullableParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
