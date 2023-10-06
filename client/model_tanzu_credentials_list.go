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

// checks if the TanzuCredentialsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuCredentialsList{}

// TanzuCredentialsList struct for TanzuCredentialsList
type TanzuCredentialsList struct {
	Data       []TanzuCredentialsListDto `json:"data,omitempty"`
	TotalCount *int32                    `json:"totalCount,omitempty"`
}

// NewTanzuCredentialsList instantiates a new TanzuCredentialsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuCredentialsList() *TanzuCredentialsList {
	this := TanzuCredentialsList{}
	return &this
}

// NewTanzuCredentialsListWithDefaults instantiates a new TanzuCredentialsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuCredentialsListWithDefaults() *TanzuCredentialsList {
	this := TanzuCredentialsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsList) GetData() []TanzuCredentialsListDto {
	if o == nil {
		var ret []TanzuCredentialsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsList) GetDataOk() ([]TanzuCredentialsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TanzuCredentialsList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TanzuCredentialsListDto and assigns it to the Data field.
func (o *TanzuCredentialsList) SetData(v []TanzuCredentialsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *TanzuCredentialsList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *TanzuCredentialsList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *TanzuCredentialsList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o TanzuCredentialsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuCredentialsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableTanzuCredentialsList struct {
	value *TanzuCredentialsList
	isSet bool
}

func (v NullableTanzuCredentialsList) Get() *TanzuCredentialsList {
	return v.value
}

func (v *NullableTanzuCredentialsList) Set(val *TanzuCredentialsList) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuCredentialsList) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuCredentialsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuCredentialsList(val *TanzuCredentialsList) *NullableTanzuCredentialsList {
	return &NullableTanzuCredentialsList{value: val, isSet: true}
}

func (v NullableTanzuCredentialsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuCredentialsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
