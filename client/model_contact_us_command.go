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

// checks if the ContactUsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactUsCommand{}

// ContactUsCommand struct for ContactUsCommand
type ContactUsCommand struct {
	Name          NullableString `json:"name,omitempty"`
	BusinessEmail NullableString `json:"businessEmail,omitempty"`
	CompanyName   NullableString `json:"companyName,omitempty"`
	Comment       NullableString `json:"comment,omitempty"`
}

// NewContactUsCommand instantiates a new ContactUsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactUsCommand() *ContactUsCommand {
	this := ContactUsCommand{}
	return &this
}

// NewContactUsCommandWithDefaults instantiates a new ContactUsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactUsCommandWithDefaults() *ContactUsCommand {
	this := ContactUsCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactUsCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactUsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ContactUsCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ContactUsCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ContactUsCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ContactUsCommand) UnsetName() {
	o.Name.Unset()
}

// GetBusinessEmail returns the BusinessEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactUsCommand) GetBusinessEmail() string {
	if o == nil || IsNil(o.BusinessEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessEmail.Get()
}

// GetBusinessEmailOk returns a tuple with the BusinessEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactUsCommand) GetBusinessEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessEmail.Get(), o.BusinessEmail.IsSet()
}

// HasBusinessEmail returns a boolean if a field has been set.
func (o *ContactUsCommand) HasBusinessEmail() bool {
	if o != nil && o.BusinessEmail.IsSet() {
		return true
	}

	return false
}

// SetBusinessEmail gets a reference to the given NullableString and assigns it to the BusinessEmail field.
func (o *ContactUsCommand) SetBusinessEmail(v string) {
	o.BusinessEmail.Set(&v)
}

// SetBusinessEmailNil sets the value for BusinessEmail to be an explicit nil
func (o *ContactUsCommand) SetBusinessEmailNil() {
	o.BusinessEmail.Set(nil)
}

// UnsetBusinessEmail ensures that no value is present for BusinessEmail, not even an explicit nil
func (o *ContactUsCommand) UnsetBusinessEmail() {
	o.BusinessEmail.Unset()
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactUsCommand) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactUsCommand) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ContactUsCommand) HasCompanyName() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *ContactUsCommand) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}

// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *ContactUsCommand) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *ContactUsCommand) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactUsCommand) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactUsCommand) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ContactUsCommand) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ContactUsCommand) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ContactUsCommand) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ContactUsCommand) UnsetComment() {
	o.Comment.Unset()
}

func (o ContactUsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactUsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.BusinessEmail.IsSet() {
		toSerialize["businessEmail"] = o.BusinessEmail.Get()
	}
	if o.CompanyName.IsSet() {
		toSerialize["companyName"] = o.CompanyName.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableContactUsCommand struct {
	value *ContactUsCommand
	isSet bool
}

func (v NullableContactUsCommand) Get() *ContactUsCommand {
	return v.value
}

func (v *NullableContactUsCommand) Set(val *ContactUsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableContactUsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableContactUsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactUsCommand(val *ContactUsCommand) *NullableContactUsCommand {
	return &NullableContactUsCommand{value: val, isSet: true}
}

func (v NullableContactUsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactUsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
