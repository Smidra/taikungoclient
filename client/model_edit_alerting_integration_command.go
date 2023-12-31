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

// checks if the EditAlertingIntegrationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAlertingIntegrationCommand{}

// EditAlertingIntegrationCommand struct for EditAlertingIntegrationCommand
type EditAlertingIntegrationCommand struct {
	Id                      *int32                   `json:"id,omitempty"`
	Url                     NullableString           `json:"url,omitempty"`
	Token                   NullableString           `json:"token,omitempty"`
	AlertingIntegrationType *AlertingIntegrationType `json:"alertingIntegrationType,omitempty"`
	AlertingProfileId       *int32                   `json:"alertingProfileId,omitempty"`
}

// NewEditAlertingIntegrationCommand instantiates a new EditAlertingIntegrationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAlertingIntegrationCommand() *EditAlertingIntegrationCommand {
	this := EditAlertingIntegrationCommand{}
	return &this
}

// NewEditAlertingIntegrationCommandWithDefaults instantiates a new EditAlertingIntegrationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAlertingIntegrationCommandWithDefaults() *EditAlertingIntegrationCommand {
	this := EditAlertingIntegrationCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EditAlertingIntegrationCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAlertingIntegrationCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EditAlertingIntegrationCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EditAlertingIntegrationCommand) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAlertingIntegrationCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAlertingIntegrationCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *EditAlertingIntegrationCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *EditAlertingIntegrationCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *EditAlertingIntegrationCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *EditAlertingIntegrationCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAlertingIntegrationCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAlertingIntegrationCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *EditAlertingIntegrationCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *EditAlertingIntegrationCommand) SetToken(v string) {
	o.Token.Set(&v)
}

// SetTokenNil sets the value for Token to be an explicit nil
func (o *EditAlertingIntegrationCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *EditAlertingIntegrationCommand) UnsetToken() {
	o.Token.Unset()
}

// GetAlertingIntegrationType returns the AlertingIntegrationType field value if set, zero value otherwise.
func (o *EditAlertingIntegrationCommand) GetAlertingIntegrationType() AlertingIntegrationType {
	if o == nil || IsNil(o.AlertingIntegrationType) {
		var ret AlertingIntegrationType
		return ret
	}
	return *o.AlertingIntegrationType
}

// GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAlertingIntegrationCommand) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool) {
	if o == nil || IsNil(o.AlertingIntegrationType) {
		return nil, false
	}
	return o.AlertingIntegrationType, true
}

// HasAlertingIntegrationType returns a boolean if a field has been set.
func (o *EditAlertingIntegrationCommand) HasAlertingIntegrationType() bool {
	if o != nil && !IsNil(o.AlertingIntegrationType) {
		return true
	}

	return false
}

// SetAlertingIntegrationType gets a reference to the given AlertingIntegrationType and assigns it to the AlertingIntegrationType field.
func (o *EditAlertingIntegrationCommand) SetAlertingIntegrationType(v AlertingIntegrationType) {
	o.AlertingIntegrationType = &v
}

// GetAlertingProfileId returns the AlertingProfileId field value if set, zero value otherwise.
func (o *EditAlertingIntegrationCommand) GetAlertingProfileId() int32 {
	if o == nil || IsNil(o.AlertingProfileId) {
		var ret int32
		return ret
	}
	return *o.AlertingProfileId
}

// GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAlertingIntegrationCommand) GetAlertingProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AlertingProfileId) {
		return nil, false
	}
	return o.AlertingProfileId, true
}

// HasAlertingProfileId returns a boolean if a field has been set.
func (o *EditAlertingIntegrationCommand) HasAlertingProfileId() bool {
	if o != nil && !IsNil(o.AlertingProfileId) {
		return true
	}

	return false
}

// SetAlertingProfileId gets a reference to the given int32 and assigns it to the AlertingProfileId field.
func (o *EditAlertingIntegrationCommand) SetAlertingProfileId(v int32) {
	o.AlertingProfileId = &v
}

func (o EditAlertingIntegrationCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAlertingIntegrationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if !IsNil(o.AlertingIntegrationType) {
		toSerialize["alertingIntegrationType"] = o.AlertingIntegrationType
	}
	if !IsNil(o.AlertingProfileId) {
		toSerialize["alertingProfileId"] = o.AlertingProfileId
	}
	return toSerialize, nil
}

type NullableEditAlertingIntegrationCommand struct {
	value *EditAlertingIntegrationCommand
	isSet bool
}

func (v NullableEditAlertingIntegrationCommand) Get() *EditAlertingIntegrationCommand {
	return v.value
}

func (v *NullableEditAlertingIntegrationCommand) Set(val *EditAlertingIntegrationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAlertingIntegrationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAlertingIntegrationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAlertingIntegrationCommand(val *EditAlertingIntegrationCommand) *NullableEditAlertingIntegrationCommand {
	return &NullableEditAlertingIntegrationCommand{value: val, isSet: true}
}

func (v NullableEditAlertingIntegrationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAlertingIntegrationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
