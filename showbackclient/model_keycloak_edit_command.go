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
)

// checks if the KeycloakEditCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeycloakEditCommand{}

// KeycloakEditCommand struct for KeycloakEditCommand
type KeycloakEditCommand struct {
	Id           *int32         `json:"id,omitempty"`
	Name         NullableString `json:"name,omitempty"`
	Url          NullableString `json:"url,omitempty"`
	RealmsName   NullableString `json:"realmsName,omitempty"`
	ClientId     NullableString `json:"clientId,omitempty"`
	ClientSecret NullableString `json:"clientSecret,omitempty"`
}

// NewKeycloakEditCommand instantiates a new KeycloakEditCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeycloakEditCommand() *KeycloakEditCommand {
	this := KeycloakEditCommand{}
	return &this
}

// NewKeycloakEditCommandWithDefaults instantiates a new KeycloakEditCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeycloakEditCommandWithDefaults() *KeycloakEditCommand {
	this := KeycloakEditCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeycloakEditCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeycloakEditCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeycloakEditCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakEditCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakEditCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KeycloakEditCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *KeycloakEditCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KeycloakEditCommand) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakEditCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakEditCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *KeycloakEditCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *KeycloakEditCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *KeycloakEditCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetRealmsName returns the RealmsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakEditCommand) GetRealmsName() string {
	if o == nil || IsNil(o.RealmsName.Get()) {
		var ret string
		return ret
	}
	return *o.RealmsName.Get()
}

// GetRealmsNameOk returns a tuple with the RealmsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakEditCommand) GetRealmsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RealmsName.Get(), o.RealmsName.IsSet()
}

// HasRealmsName returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasRealmsName() bool {
	if o != nil && o.RealmsName.IsSet() {
		return true
	}

	return false
}

// SetRealmsName gets a reference to the given NullableString and assigns it to the RealmsName field.
func (o *KeycloakEditCommand) SetRealmsName(v string) {
	o.RealmsName.Set(&v)
}

// SetRealmsNameNil sets the value for RealmsName to be an explicit nil
func (o *KeycloakEditCommand) SetRealmsNameNil() {
	o.RealmsName.Set(nil)
}

// UnsetRealmsName ensures that no value is present for RealmsName, not even an explicit nil
func (o *KeycloakEditCommand) UnsetRealmsName() {
	o.RealmsName.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakEditCommand) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakEditCommand) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *KeycloakEditCommand) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *KeycloakEditCommand) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *KeycloakEditCommand) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakEditCommand) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakEditCommand) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *KeycloakEditCommand) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *KeycloakEditCommand) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *KeycloakEditCommand) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *KeycloakEditCommand) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

func (o KeycloakEditCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeycloakEditCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.RealmsName.IsSet() {
		toSerialize["realmsName"] = o.RealmsName.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	return toSerialize, nil
}

type NullableKeycloakEditCommand struct {
	value *KeycloakEditCommand
	isSet bool
}

func (v NullableKeycloakEditCommand) Get() *KeycloakEditCommand {
	return v.value
}

func (v *NullableKeycloakEditCommand) Set(val *KeycloakEditCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKeycloakEditCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKeycloakEditCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeycloakEditCommand(val *KeycloakEditCommand) *NullableKeycloakEditCommand {
	return &NullableKeycloakEditCommand{value: val, isSet: true}
}

func (v NullableKeycloakEditCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeycloakEditCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
