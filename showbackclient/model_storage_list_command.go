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

// checks if the StorageListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageListCommand{}

// StorageListCommand struct for StorageListCommand
type StorageListCommand struct {
	Url         NullableString `json:"url,omitempty"`
	TokenId     NullableString `json:"tokenId,omitempty"`
	TokenSecret NullableString `json:"tokenSecret,omitempty"`
}

// NewStorageListCommand instantiates a new StorageListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageListCommand() *StorageListCommand {
	this := StorageListCommand{}
	return &this
}

// NewStorageListCommandWithDefaults instantiates a new StorageListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageListCommandWithDefaults() *StorageListCommand {
	this := StorageListCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageListCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *StorageListCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *StorageListCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *StorageListCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *StorageListCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageListCommand) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageListCommand) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *StorageListCommand) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *StorageListCommand) SetTokenId(v string) {
	o.TokenId.Set(&v)
}

// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *StorageListCommand) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *StorageListCommand) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetTokenSecret returns the TokenSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageListCommand) GetTokenSecret() string {
	if o == nil || IsNil(o.TokenSecret.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSecret.Get()
}

// GetTokenSecretOk returns a tuple with the TokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageListCommand) GetTokenSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSecret.Get(), o.TokenSecret.IsSet()
}

// HasTokenSecret returns a boolean if a field has been set.
func (o *StorageListCommand) HasTokenSecret() bool {
	if o != nil && o.TokenSecret.IsSet() {
		return true
	}

	return false
}

// SetTokenSecret gets a reference to the given NullableString and assigns it to the TokenSecret field.
func (o *StorageListCommand) SetTokenSecret(v string) {
	o.TokenSecret.Set(&v)
}

// SetTokenSecretNil sets the value for TokenSecret to be an explicit nil
func (o *StorageListCommand) SetTokenSecretNil() {
	o.TokenSecret.Set(nil)
}

// UnsetTokenSecret ensures that no value is present for TokenSecret, not even an explicit nil
func (o *StorageListCommand) UnsetTokenSecret() {
	o.TokenSecret.Unset()
}

func (o StorageListCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["tokenId"] = o.TokenId.Get()
	}
	if o.TokenSecret.IsSet() {
		toSerialize["tokenSecret"] = o.TokenSecret.Get()
	}
	return toSerialize, nil
}

type NullableStorageListCommand struct {
	value *StorageListCommand
	isSet bool
}

func (v NullableStorageListCommand) Get() *StorageListCommand {
	return v.value
}

func (v *NullableStorageListCommand) Set(val *StorageListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageListCommand(val *StorageListCommand) *NullableStorageListCommand {
	return &NullableStorageListCommand{value: val, isSet: true}
}

func (v NullableStorageListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
