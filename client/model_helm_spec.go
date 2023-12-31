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

// checks if the HelmSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmSpec{}

// HelmSpec struct for HelmSpec
type HelmSpec struct {
	Interval         NullableString `json:"interval,omitempty"`
	ReleaseName      NullableString `json:"releaseName,omitempty"`
	Url              NullableString `json:"url,omitempty"`
	TargetNamespace  NullableString `json:"targetNamespace,omitempty"`
	StorageNamespace NullableString `json:"storageNamespace,omitempty"`
	Chart            *Chart         `json:"chart,omitempty"`
	Values           *JsonNode      `json:"values,omitempty"`
}

// NewHelmSpec instantiates a new HelmSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmSpec() *HelmSpec {
	this := HelmSpec{}
	return &this
}

// NewHelmSpecWithDefaults instantiates a new HelmSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmSpecWithDefaults() *HelmSpec {
	this := HelmSpec{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmSpec) GetInterval() string {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret string
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmSpec) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *HelmSpec) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableString and assigns it to the Interval field.
func (o *HelmSpec) SetInterval(v string) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *HelmSpec) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *HelmSpec) UnsetInterval() {
	o.Interval.Unset()
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmSpec) GetReleaseName() string {
	if o == nil || IsNil(o.ReleaseName.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseName.Get()
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmSpec) GetReleaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseName.Get(), o.ReleaseName.IsSet()
}

// HasReleaseName returns a boolean if a field has been set.
func (o *HelmSpec) HasReleaseName() bool {
	if o != nil && o.ReleaseName.IsSet() {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given NullableString and assigns it to the ReleaseName field.
func (o *HelmSpec) SetReleaseName(v string) {
	o.ReleaseName.Set(&v)
}

// SetReleaseNameNil sets the value for ReleaseName to be an explicit nil
func (o *HelmSpec) SetReleaseNameNil() {
	o.ReleaseName.Set(nil)
}

// UnsetReleaseName ensures that no value is present for ReleaseName, not even an explicit nil
func (o *HelmSpec) UnsetReleaseName() {
	o.ReleaseName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmSpec) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmSpec) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *HelmSpec) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *HelmSpec) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *HelmSpec) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *HelmSpec) UnsetUrl() {
	o.Url.Unset()
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmSpec) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace.Get()) {
		var ret string
		return ret
	}
	return *o.TargetNamespace.Get()
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmSpec) GetTargetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetNamespace.Get(), o.TargetNamespace.IsSet()
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *HelmSpec) HasTargetNamespace() bool {
	if o != nil && o.TargetNamespace.IsSet() {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given NullableString and assigns it to the TargetNamespace field.
func (o *HelmSpec) SetTargetNamespace(v string) {
	o.TargetNamespace.Set(&v)
}

// SetTargetNamespaceNil sets the value for TargetNamespace to be an explicit nil
func (o *HelmSpec) SetTargetNamespaceNil() {
	o.TargetNamespace.Set(nil)
}

// UnsetTargetNamespace ensures that no value is present for TargetNamespace, not even an explicit nil
func (o *HelmSpec) UnsetTargetNamespace() {
	o.TargetNamespace.Unset()
}

// GetStorageNamespace returns the StorageNamespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmSpec) GetStorageNamespace() string {
	if o == nil || IsNil(o.StorageNamespace.Get()) {
		var ret string
		return ret
	}
	return *o.StorageNamespace.Get()
}

// GetStorageNamespaceOk returns a tuple with the StorageNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmSpec) GetStorageNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageNamespace.Get(), o.StorageNamespace.IsSet()
}

// HasStorageNamespace returns a boolean if a field has been set.
func (o *HelmSpec) HasStorageNamespace() bool {
	if o != nil && o.StorageNamespace.IsSet() {
		return true
	}

	return false
}

// SetStorageNamespace gets a reference to the given NullableString and assigns it to the StorageNamespace field.
func (o *HelmSpec) SetStorageNamespace(v string) {
	o.StorageNamespace.Set(&v)
}

// SetStorageNamespaceNil sets the value for StorageNamespace to be an explicit nil
func (o *HelmSpec) SetStorageNamespaceNil() {
	o.StorageNamespace.Set(nil)
}

// UnsetStorageNamespace ensures that no value is present for StorageNamespace, not even an explicit nil
func (o *HelmSpec) UnsetStorageNamespace() {
	o.StorageNamespace.Unset()
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *HelmSpec) GetChart() Chart {
	if o == nil || IsNil(o.Chart) {
		var ret Chart
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmSpec) GetChartOk() (*Chart, bool) {
	if o == nil || IsNil(o.Chart) {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *HelmSpec) HasChart() bool {
	if o != nil && !IsNil(o.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given Chart and assigns it to the Chart field.
func (o *HelmSpec) SetChart(v Chart) {
	o.Chart = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *HelmSpec) GetValues() JsonNode {
	if o == nil || IsNil(o.Values) {
		var ret JsonNode
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmSpec) GetValuesOk() (*JsonNode, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HelmSpec) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given JsonNode and assigns it to the Values field.
func (o *HelmSpec) SetValues(v JsonNode) {
	o.Values = &v
}

func (o HelmSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.ReleaseName.IsSet() {
		toSerialize["releaseName"] = o.ReleaseName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.TargetNamespace.IsSet() {
		toSerialize["targetNamespace"] = o.TargetNamespace.Get()
	}
	if o.StorageNamespace.IsSet() {
		toSerialize["storageNamespace"] = o.StorageNamespace.Get()
	}
	if !IsNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableHelmSpec struct {
	value *HelmSpec
	isSet bool
}

func (v NullableHelmSpec) Get() *HelmSpec {
	return v.value
}

func (v *NullableHelmSpec) Set(val *HelmSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmSpec(val *HelmSpec) *NullableHelmSpec {
	return &NullableHelmSpec{value: val, isSet: true}
}

func (v NullableHelmSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
