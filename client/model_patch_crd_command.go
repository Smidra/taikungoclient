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

// checks if the PatchCrdCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchCrdCommand{}

// PatchCrdCommand struct for PatchCrdCommand
type PatchCrdCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Yaml      NullableString `json:"yaml,omitempty"`
	Name      NullableString `json:"name,omitempty"`
}

// NewPatchCrdCommand instantiates a new PatchCrdCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchCrdCommand() *PatchCrdCommand {
	this := PatchCrdCommand{}
	return &this
}

// NewPatchCrdCommandWithDefaults instantiates a new PatchCrdCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchCrdCommandWithDefaults() *PatchCrdCommand {
	this := PatchCrdCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PatchCrdCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchCrdCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PatchCrdCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PatchCrdCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCrdCommand) GetYaml() string {
	if o == nil || IsNil(o.Yaml.Get()) {
		var ret string
		return ret
	}
	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCrdCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// HasYaml returns a boolean if a field has been set.
func (o *PatchCrdCommand) HasYaml() bool {
	if o != nil && o.Yaml.IsSet() {
		return true
	}

	return false
}

// SetYaml gets a reference to the given NullableString and assigns it to the Yaml field.
func (o *PatchCrdCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}

// SetYamlNil sets the value for Yaml to be an explicit nil
func (o *PatchCrdCommand) SetYamlNil() {
	o.Yaml.Set(nil)
}

// UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
func (o *PatchCrdCommand) UnsetYaml() {
	o.Yaml.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCrdCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCrdCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchCrdCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchCrdCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchCrdCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchCrdCommand) UnsetName() {
	o.Name.Unset()
}

func (o PatchCrdCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCrdCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Yaml.IsSet() {
		toSerialize["yaml"] = o.Yaml.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullablePatchCrdCommand struct {
	value *PatchCrdCommand
	isSet bool
}

func (v NullablePatchCrdCommand) Get() *PatchCrdCommand {
	return v.value
}

func (v *NullablePatchCrdCommand) Set(val *PatchCrdCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCrdCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCrdCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCrdCommand(val *PatchCrdCommand) *NullablePatchCrdCommand {
	return &NullablePatchCrdCommand{value: val, isSet: true}
}

func (v NullablePatchCrdCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCrdCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
