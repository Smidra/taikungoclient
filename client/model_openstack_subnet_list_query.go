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

// checks if the OpenstackSubnetListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackSubnetListQuery{}

// OpenstackSubnetListQuery struct for OpenstackSubnetListQuery
type OpenstackSubnetListQuery struct {
	OpenStackUser          NullableString `json:"openStackUser,omitempty"`
	OpenStackPassword      NullableString `json:"openStackPassword,omitempty"`
	OpenStackUrl           NullableString `json:"openStackUrl,omitempty"`
	OpenStackProject       NullableString `json:"openStackProject,omitempty"`
	OpenStackProjectId     NullableString `json:"openStackProjectId,omitempty"`
	OpenStackDomain        NullableString `json:"openStackDomain,omitempty"`
	OpenStackRegion        NullableString `json:"openStackRegion,omitempty"`
	ApplicationCredEnabled *bool          `json:"applicationCredEnabled,omitempty"`
}

// NewOpenstackSubnetListQuery instantiates a new OpenstackSubnetListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackSubnetListQuery() *OpenstackSubnetListQuery {
	this := OpenstackSubnetListQuery{}
	return &this
}

// NewOpenstackSubnetListQueryWithDefaults instantiates a new OpenstackSubnetListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackSubnetListQueryWithDefaults() *OpenstackSubnetListQuery {
	this := OpenstackSubnetListQuery{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *OpenstackSubnetListQuery) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}

// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *OpenstackSubnetListQuery) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}

// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl.Get()
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUrl.Get(), o.OpenStackUrl.IsSet()
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackUrl() bool {
	if o != nil && o.OpenStackUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given NullableString and assigns it to the OpenStackUrl field.
func (o *OpenstackSubnetListQuery) SetOpenStackUrl(v string) {
	o.OpenStackUrl.Set(&v)
}

// SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackUrlNil() {
	o.OpenStackUrl.Set(nil)
}

// UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackUrl() {
	o.OpenStackUrl.Unset()
}

// GetOpenStackProject returns the OpenStackProject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackProject() string {
	if o == nil || IsNil(o.OpenStackProject.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackProject.Get()
}

// GetOpenStackProjectOk returns a tuple with the OpenStackProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackProject.Get(), o.OpenStackProject.IsSet()
}

// HasOpenStackProject returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackProject() bool {
	if o != nil && o.OpenStackProject.IsSet() {
		return true
	}

	return false
}

// SetOpenStackProject gets a reference to the given NullableString and assigns it to the OpenStackProject field.
func (o *OpenstackSubnetListQuery) SetOpenStackProject(v string) {
	o.OpenStackProject.Set(&v)
}

// SetOpenStackProjectNil sets the value for OpenStackProject to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackProjectNil() {
	o.OpenStackProject.Set(nil)
}

// UnsetOpenStackProject ensures that no value is present for OpenStackProject, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackProject() {
	o.OpenStackProject.Unset()
}

// GetOpenStackProjectId returns the OpenStackProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackProjectId() string {
	if o == nil || IsNil(o.OpenStackProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackProjectId.Get()
}

// GetOpenStackProjectIdOk returns a tuple with the OpenStackProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackProjectId.Get(), o.OpenStackProjectId.IsSet()
}

// HasOpenStackProjectId returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackProjectId() bool {
	if o != nil && o.OpenStackProjectId.IsSet() {
		return true
	}

	return false
}

// SetOpenStackProjectId gets a reference to the given NullableString and assigns it to the OpenStackProjectId field.
func (o *OpenstackSubnetListQuery) SetOpenStackProjectId(v string) {
	o.OpenStackProjectId.Set(&v)
}

// SetOpenStackProjectIdNil sets the value for OpenStackProjectId to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackProjectIdNil() {
	o.OpenStackProjectId.Set(nil)
}

// UnsetOpenStackProjectId ensures that no value is present for OpenStackProjectId, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackProjectId() {
	o.OpenStackProjectId.Unset()
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain.Get()
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackDomain.Get(), o.OpenStackDomain.IsSet()
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackDomain() bool {
	if o != nil && o.OpenStackDomain.IsSet() {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given NullableString and assigns it to the OpenStackDomain field.
func (o *OpenstackSubnetListQuery) SetOpenStackDomain(v string) {
	o.OpenStackDomain.Set(&v)
}

// SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackDomainNil() {
	o.OpenStackDomain.Set(nil)
}

// UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackDomain() {
	o.OpenStackDomain.Unset()
}

// GetOpenStackRegion returns the OpenStackRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackSubnetListQuery) GetOpenStackRegion() string {
	if o == nil || IsNil(o.OpenStackRegion.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackRegion.Get()
}

// GetOpenStackRegionOk returns a tuple with the OpenStackRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackSubnetListQuery) GetOpenStackRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackRegion.Get(), o.OpenStackRegion.IsSet()
}

// HasOpenStackRegion returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasOpenStackRegion() bool {
	if o != nil && o.OpenStackRegion.IsSet() {
		return true
	}

	return false
}

// SetOpenStackRegion gets a reference to the given NullableString and assigns it to the OpenStackRegion field.
func (o *OpenstackSubnetListQuery) SetOpenStackRegion(v string) {
	o.OpenStackRegion.Set(&v)
}

// SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil
func (o *OpenstackSubnetListQuery) SetOpenStackRegionNil() {
	o.OpenStackRegion.Set(nil)
}

// UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
func (o *OpenstackSubnetListQuery) UnsetOpenStackRegion() {
	o.OpenStackRegion.Unset()
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenstackSubnetListQuery) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackSubnetListQuery) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenstackSubnetListQuery) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenstackSubnetListQuery) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

func (o OpenstackSubnetListQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackSubnetListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	if o.OpenStackUrl.IsSet() {
		toSerialize["openStackUrl"] = o.OpenStackUrl.Get()
	}
	if o.OpenStackProject.IsSet() {
		toSerialize["openStackProject"] = o.OpenStackProject.Get()
	}
	if o.OpenStackProjectId.IsSet() {
		toSerialize["openStackProjectId"] = o.OpenStackProjectId.Get()
	}
	if o.OpenStackDomain.IsSet() {
		toSerialize["openStackDomain"] = o.OpenStackDomain.Get()
	}
	if o.OpenStackRegion.IsSet() {
		toSerialize["openStackRegion"] = o.OpenStackRegion.Get()
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	return toSerialize, nil
}

type NullableOpenstackSubnetListQuery struct {
	value *OpenstackSubnetListQuery
	isSet bool
}

func (v NullableOpenstackSubnetListQuery) Get() *OpenstackSubnetListQuery {
	return v.value
}

func (v *NullableOpenstackSubnetListQuery) Set(val *OpenstackSubnetListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackSubnetListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackSubnetListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackSubnetListQuery(val *OpenstackSubnetListQuery) *NullableOpenstackSubnetListQuery {
	return &NullableOpenstackSubnetListQuery{value: val, isSet: true}
}

func (v NullableOpenstackSubnetListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackSubnetListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
