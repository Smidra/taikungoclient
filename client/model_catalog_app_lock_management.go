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

// checks if the CatalogAppLockManagement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogAppLockManagement{}

// CatalogAppLockManagement struct for CatalogAppLockManagement
type CatalogAppLockManagement struct {
	Id   *int32         `json:"id,omitempty"`
	Mode NullableString `json:"mode,omitempty"`
}

// NewCatalogAppLockManagement instantiates a new CatalogAppLockManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogAppLockManagement() *CatalogAppLockManagement {
	this := CatalogAppLockManagement{}
	return &this
}

// NewCatalogAppLockManagementWithDefaults instantiates a new CatalogAppLockManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogAppLockManagementWithDefaults() *CatalogAppLockManagement {
	this := CatalogAppLockManagement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogAppLockManagement) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppLockManagement) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogAppLockManagement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CatalogAppLockManagement) SetId(v int32) {
	o.Id = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppLockManagement) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppLockManagement) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *CatalogAppLockManagement) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *CatalogAppLockManagement) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *CatalogAppLockManagement) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *CatalogAppLockManagement) UnsetMode() {
	o.Mode.Unset()
}

func (o CatalogAppLockManagement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogAppLockManagement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableCatalogAppLockManagement struct {
	value *CatalogAppLockManagement
	isSet bool
}

func (v NullableCatalogAppLockManagement) Get() *CatalogAppLockManagement {
	return v.value
}

func (v *NullableCatalogAppLockManagement) Set(val *CatalogAppLockManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogAppLockManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogAppLockManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogAppLockManagement(val *CatalogAppLockManagement) *NullableCatalogAppLockManagement {
	return &NullableCatalogAppLockManagement{value: val, isSet: true}
}

func (v NullableCatalogAppLockManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogAppLockManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
