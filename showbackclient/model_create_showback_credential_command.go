/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the CreateShowbackCredentialCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShowbackCredentialCommand{}

// CreateShowbackCredentialCommand struct for CreateShowbackCredentialCommand
type CreateShowbackCredentialCommand struct {
	Name           NullableString `json:"name,omitempty"`
	Url            NullableString `json:"url,omitempty"`
	Username       NullableString `json:"username,omitempty"`
	Password       NullableString `json:"password,omitempty"`
	OrganizationId NullableInt32  `json:"organizationId,omitempty"`
}

// NewCreateShowbackCredentialCommand instantiates a new CreateShowbackCredentialCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShowbackCredentialCommand() *CreateShowbackCredentialCommand {
	this := CreateShowbackCredentialCommand{}
	return &this
}

// NewCreateShowbackCredentialCommandWithDefaults instantiates a new CreateShowbackCredentialCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShowbackCredentialCommandWithDefaults() *CreateShowbackCredentialCommand {
	this := CreateShowbackCredentialCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateShowbackCredentialCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateShowbackCredentialCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateShowbackCredentialCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateShowbackCredentialCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateShowbackCredentialCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateShowbackCredentialCommand) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateShowbackCredentialCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateShowbackCredentialCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateShowbackCredentialCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CreateShowbackCredentialCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *CreateShowbackCredentialCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CreateShowbackCredentialCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateShowbackCredentialCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateShowbackCredentialCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateShowbackCredentialCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *CreateShowbackCredentialCommand) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *CreateShowbackCredentialCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *CreateShowbackCredentialCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateShowbackCredentialCommand) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateShowbackCredentialCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateShowbackCredentialCommand) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CreateShowbackCredentialCommand) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateShowbackCredentialCommand) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateShowbackCredentialCommand) UnsetPassword() {
	o.Password.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateShowbackCredentialCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateShowbackCredentialCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateShowbackCredentialCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateShowbackCredentialCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateShowbackCredentialCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateShowbackCredentialCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o CreateShowbackCredentialCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShowbackCredentialCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableCreateShowbackCredentialCommand struct {
	value *CreateShowbackCredentialCommand
	isSet bool
}

func (v NullableCreateShowbackCredentialCommand) Get() *CreateShowbackCredentialCommand {
	return v.value
}

func (v *NullableCreateShowbackCredentialCommand) Set(val *CreateShowbackCredentialCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShowbackCredentialCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShowbackCredentialCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShowbackCredentialCommand(val *CreateShowbackCredentialCommand) *NullableCreateShowbackCredentialCommand {
	return &NullableCreateShowbackCredentialCommand{value: val, isSet: true}
}

func (v NullableCreateShowbackCredentialCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShowbackCredentialCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}