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

// checks if the CreateProjectFromTemplateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectFromTemplateCommand{}

// CreateProjectFromTemplateCommand struct for CreateProjectFromTemplateCommand
type CreateProjectFromTemplateCommand struct {
	Id          *int32         `json:"id,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	CanCommit   *bool          `json:"canCommit,omitempty"`
}

// NewCreateProjectFromTemplateCommand instantiates a new CreateProjectFromTemplateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectFromTemplateCommand() *CreateProjectFromTemplateCommand {
	this := CreateProjectFromTemplateCommand{}
	return &this
}

// NewCreateProjectFromTemplateCommandWithDefaults instantiates a new CreateProjectFromTemplateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectFromTemplateCommandWithDefaults() *CreateProjectFromTemplateCommand {
	this := CreateProjectFromTemplateCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateProjectFromTemplateCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectFromTemplateCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateProjectFromTemplateCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateProjectFromTemplateCommand) SetId(v int32) {
	o.Id = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectFromTemplateCommand) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectFromTemplateCommand) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *CreateProjectFromTemplateCommand) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *CreateProjectFromTemplateCommand) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *CreateProjectFromTemplateCommand) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *CreateProjectFromTemplateCommand) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetCanCommit returns the CanCommit field value if set, zero value otherwise.
func (o *CreateProjectFromTemplateCommand) GetCanCommit() bool {
	if o == nil || IsNil(o.CanCommit) {
		var ret bool
		return ret
	}
	return *o.CanCommit
}

// GetCanCommitOk returns a tuple with the CanCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectFromTemplateCommand) GetCanCommitOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCommit) {
		return nil, false
	}
	return o.CanCommit, true
}

// HasCanCommit returns a boolean if a field has been set.
func (o *CreateProjectFromTemplateCommand) HasCanCommit() bool {
	if o != nil && !IsNil(o.CanCommit) {
		return true
	}

	return false
}

// SetCanCommit gets a reference to the given bool and assigns it to the CanCommit field.
func (o *CreateProjectFromTemplateCommand) SetCanCommit(v bool) {
	o.CanCommit = &v
}

func (o CreateProjectFromTemplateCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectFromTemplateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.CanCommit) {
		toSerialize["canCommit"] = o.CanCommit
	}
	return toSerialize, nil
}

type NullableCreateProjectFromTemplateCommand struct {
	value *CreateProjectFromTemplateCommand
	isSet bool
}

func (v NullableCreateProjectFromTemplateCommand) Get() *CreateProjectFromTemplateCommand {
	return v.value
}

func (v *NullableCreateProjectFromTemplateCommand) Set(val *CreateProjectFromTemplateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectFromTemplateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectFromTemplateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectFromTemplateCommand(val *CreateProjectFromTemplateCommand) *NullableCreateProjectFromTemplateCommand {
	return &NullableCreateProjectFromTemplateCommand{value: val, isSet: true}
}

func (v NullableCreateProjectFromTemplateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectFromTemplateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
