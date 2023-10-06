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

// checks if the ProjectTemplateListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTemplateListDto{}

// ProjectTemplateListDto struct for ProjectTemplateListDto
type ProjectTemplateListDto struct {
	Id                      *int32         `json:"id,omitempty"`
	Name                    NullableString `json:"name,omitempty"`
	MonitoringEnabled       *bool          `json:"monitoringEnabled,omitempty"`
	BackupEnabled           *bool          `json:"backupEnabled,omitempty"`
	AllowFullSpotKubernetes *bool          `json:"allowFullSpotKubernetes,omitempty"`
	AllowSpotVms            *bool          `json:"allowSpotVms,omitempty"`
	AllowSpotWorkers        *bool          `json:"allowSpotWorkers,omitempty"`
	KubernetesVersion       NullableString `json:"kubernetesVersion,omitempty"`
	OrganizationName        NullableString `json:"organizationName,omitempty"`
	OrganizationId          NullableInt32  `json:"organizationId,omitempty"`
}

// NewProjectTemplateListDto instantiates a new ProjectTemplateListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTemplateListDto() *ProjectTemplateListDto {
	this := ProjectTemplateListDto{}
	return &this
}

// NewProjectTemplateListDtoWithDefaults instantiates a new ProjectTemplateListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTemplateListDtoWithDefaults() *ProjectTemplateListDto {
	this := ProjectTemplateListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectTemplateListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTemplateListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTemplateListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectTemplateListDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectTemplateListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectTemplateListDto) UnsetName() {
	o.Name.Unset()
}

// GetMonitoringEnabled returns the MonitoringEnabled field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetMonitoringEnabled() bool {
	if o == nil || IsNil(o.MonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.MonitoringEnabled
}

// GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonitoringEnabled) {
		return nil, false
	}
	return o.MonitoringEnabled, true
}

// HasMonitoringEnabled returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasMonitoringEnabled() bool {
	if o != nil && !IsNil(o.MonitoringEnabled) {
		return true
	}

	return false
}

// SetMonitoringEnabled gets a reference to the given bool and assigns it to the MonitoringEnabled field.
func (o *ProjectTemplateListDto) SetMonitoringEnabled(v bool) {
	o.MonitoringEnabled = &v
}

// GetBackupEnabled returns the BackupEnabled field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetBackupEnabled() bool {
	if o == nil || IsNil(o.BackupEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupEnabled
}

// GetBackupEnabledOk returns a tuple with the BackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupEnabled) {
		return nil, false
	}
	return o.BackupEnabled, true
}

// HasBackupEnabled returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasBackupEnabled() bool {
	if o != nil && !IsNil(o.BackupEnabled) {
		return true
	}

	return false
}

// SetBackupEnabled gets a reference to the given bool and assigns it to the BackupEnabled field.
func (o *ProjectTemplateListDto) SetBackupEnabled(v bool) {
	o.BackupEnabled = &v
}

// GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetAllowFullSpotKubernetes() bool {
	if o == nil || IsNil(o.AllowFullSpotKubernetes) {
		var ret bool
		return ret
	}
	return *o.AllowFullSpotKubernetes
}

// GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetAllowFullSpotKubernetesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowFullSpotKubernetes) {
		return nil, false
	}
	return o.AllowFullSpotKubernetes, true
}

// HasAllowFullSpotKubernetes returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasAllowFullSpotKubernetes() bool {
	if o != nil && !IsNil(o.AllowFullSpotKubernetes) {
		return true
	}

	return false
}

// SetAllowFullSpotKubernetes gets a reference to the given bool and assigns it to the AllowFullSpotKubernetes field.
func (o *ProjectTemplateListDto) SetAllowFullSpotKubernetes(v bool) {
	o.AllowFullSpotKubernetes = &v
}

// GetAllowSpotVms returns the AllowSpotVms field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetAllowSpotVms() bool {
	if o == nil || IsNil(o.AllowSpotVms) {
		var ret bool
		return ret
	}
	return *o.AllowSpotVms
}

// GetAllowSpotVmsOk returns a tuple with the AllowSpotVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetAllowSpotVmsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSpotVms) {
		return nil, false
	}
	return o.AllowSpotVms, true
}

// HasAllowSpotVms returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasAllowSpotVms() bool {
	if o != nil && !IsNil(o.AllowSpotVms) {
		return true
	}

	return false
}

// SetAllowSpotVms gets a reference to the given bool and assigns it to the AllowSpotVms field.
func (o *ProjectTemplateListDto) SetAllowSpotVms(v bool) {
	o.AllowSpotVms = &v
}

// GetAllowSpotWorkers returns the AllowSpotWorkers field value if set, zero value otherwise.
func (o *ProjectTemplateListDto) GetAllowSpotWorkers() bool {
	if o == nil || IsNil(o.AllowSpotWorkers) {
		var ret bool
		return ret
	}
	return *o.AllowSpotWorkers
}

// GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTemplateListDto) GetAllowSpotWorkersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSpotWorkers) {
		return nil, false
	}
	return o.AllowSpotWorkers, true
}

// HasAllowSpotWorkers returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasAllowSpotWorkers() bool {
	if o != nil && !IsNil(o.AllowSpotWorkers) {
		return true
	}

	return false
}

// SetAllowSpotWorkers gets a reference to the given bool and assigns it to the AllowSpotWorkers field.
func (o *ProjectTemplateListDto) SetAllowSpotWorkers(v bool) {
	o.AllowSpotWorkers = &v
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTemplateListDto) GetKubernetesVersion() string {
	if o == nil || IsNil(o.KubernetesVersion.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesVersion.Get()
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTemplateListDto) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesVersion.Get(), o.KubernetesVersion.IsSet()
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion.IsSet() {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given NullableString and assigns it to the KubernetesVersion field.
func (o *ProjectTemplateListDto) SetKubernetesVersion(v string) {
	o.KubernetesVersion.Set(&v)
}

// SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil
func (o *ProjectTemplateListDto) SetKubernetesVersionNil() {
	o.KubernetesVersion.Set(nil)
}

// UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
func (o *ProjectTemplateListDto) UnsetKubernetesVersion() {
	o.KubernetesVersion.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTemplateListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTemplateListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *ProjectTemplateListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *ProjectTemplateListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *ProjectTemplateListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTemplateListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTemplateListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ProjectTemplateListDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *ProjectTemplateListDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *ProjectTemplateListDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *ProjectTemplateListDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o ProjectTemplateListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTemplateListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.MonitoringEnabled) {
		toSerialize["monitoringEnabled"] = o.MonitoringEnabled
	}
	if !IsNil(o.BackupEnabled) {
		toSerialize["backupEnabled"] = o.BackupEnabled
	}
	if !IsNil(o.AllowFullSpotKubernetes) {
		toSerialize["allowFullSpotKubernetes"] = o.AllowFullSpotKubernetes
	}
	if !IsNil(o.AllowSpotVms) {
		toSerialize["allowSpotVms"] = o.AllowSpotVms
	}
	if !IsNil(o.AllowSpotWorkers) {
		toSerialize["allowSpotWorkers"] = o.AllowSpotWorkers
	}
	if o.KubernetesVersion.IsSet() {
		toSerialize["kubernetesVersion"] = o.KubernetesVersion.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableProjectTemplateListDto struct {
	value *ProjectTemplateListDto
	isSet bool
}

func (v NullableProjectTemplateListDto) Get() *ProjectTemplateListDto {
	return v.value
}

func (v *NullableProjectTemplateListDto) Set(val *ProjectTemplateListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTemplateListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTemplateListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTemplateListDto(val *ProjectTemplateListDto) *NullableProjectTemplateListDto {
	return &NullableProjectTemplateListDto{value: val, isSet: true}
}

func (v NullableProjectTemplateListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTemplateListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}