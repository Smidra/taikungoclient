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

// checks if the PrometheusRulesSearchResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusRulesSearchResponseData{}

// PrometheusRulesSearchResponseData struct for PrometheusRulesSearchResponseData
type PrometheusRulesSearchResponseData struct {
	Id      *int32         `json:"id,omitempty"`
	Name    NullableString `json:"name,omitempty"`
	Partner *int32         `json:"partner,omitempty"`
}

// NewPrometheusRulesSearchResponseData instantiates a new PrometheusRulesSearchResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusRulesSearchResponseData() *PrometheusRulesSearchResponseData {
	this := PrometheusRulesSearchResponseData{}
	return &this
}

// NewPrometheusRulesSearchResponseDataWithDefaults instantiates a new PrometheusRulesSearchResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusRulesSearchResponseDataWithDefaults() *PrometheusRulesSearchResponseData {
	this := PrometheusRulesSearchResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrometheusRulesSearchResponseData) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesSearchResponseData) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrometheusRulesSearchResponseData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PrometheusRulesSearchResponseData) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRulesSearchResponseData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRulesSearchResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrometheusRulesSearchResponseData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrometheusRulesSearchResponseData) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PrometheusRulesSearchResponseData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrometheusRulesSearchResponseData) UnsetName() {
	o.Name.Unset()
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *PrometheusRulesSearchResponseData) GetPartner() int32 {
	if o == nil || IsNil(o.Partner) {
		var ret int32
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesSearchResponseData) GetPartnerOk() (*int32, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *PrometheusRulesSearchResponseData) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given int32 and assigns it to the Partner field.
func (o *PrometheusRulesSearchResponseData) SetPartner(v int32) {
	o.Partner = &v
}

func (o PrometheusRulesSearchResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusRulesSearchResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	return toSerialize, nil
}

type NullablePrometheusRulesSearchResponseData struct {
	value *PrometheusRulesSearchResponseData
	isSet bool
}

func (v NullablePrometheusRulesSearchResponseData) Get() *PrometheusRulesSearchResponseData {
	return v.value
}

func (v *NullablePrometheusRulesSearchResponseData) Set(val *PrometheusRulesSearchResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusRulesSearchResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusRulesSearchResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusRulesSearchResponseData(val *PrometheusRulesSearchResponseData) *NullablePrometheusRulesSearchResponseData {
	return &NullablePrometheusRulesSearchResponseData{value: val, isSet: true}
}

func (v NullablePrometheusRulesSearchResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusRulesSearchResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
