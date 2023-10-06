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

// checks if the UpdateStandAloneVmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStandAloneVmCommand{}

// UpdateStandAloneVmCommand struct for UpdateStandAloneVmCommand
type UpdateStandAloneVmCommand struct {
	Id         *int32                      `json:"id,omitempty"`
	IpAddress  NullableString              `json:"ipAddress,omitempty"`
	PublicIp   NullableString              `json:"publicIp,omitempty"`
	InstanceId NullableString              `json:"instanceId,omitempty"`
	FlavorId   NullableString              `json:"flavorId,omitempty"`
	Revision   NullableInt32               `json:"revision,omitempty"`
	Disks      []UpdateStandAloneVmDiskDto `json:"disks,omitempty"`
}

// NewUpdateStandAloneVmCommand instantiates a new UpdateStandAloneVmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStandAloneVmCommand() *UpdateStandAloneVmCommand {
	this := UpdateStandAloneVmCommand{}
	return &this
}

// NewUpdateStandAloneVmCommandWithDefaults instantiates a new UpdateStandAloneVmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStandAloneVmCommandWithDefaults() *UpdateStandAloneVmCommand {
	this := UpdateStandAloneVmCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateStandAloneVmCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStandAloneVmCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateStandAloneVmCommand) SetId(v int32) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *UpdateStandAloneVmCommand) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *UpdateStandAloneVmCommand) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *UpdateStandAloneVmCommand) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp.Get()) {
		var ret string
		return ret
	}
	return *o.PublicIp.Get()
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIp.Get(), o.PublicIp.IsSet()
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasPublicIp() bool {
	if o != nil && o.PublicIp.IsSet() {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given NullableString and assigns it to the PublicIp field.
func (o *UpdateStandAloneVmCommand) SetPublicIp(v string) {
	o.PublicIp.Set(&v)
}

// SetPublicIpNil sets the value for PublicIp to be an explicit nil
func (o *UpdateStandAloneVmCommand) SetPublicIpNil() {
	o.PublicIp.Set(nil)
}

// UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
func (o *UpdateStandAloneVmCommand) UnsetPublicIp() {
	o.PublicIp.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *UpdateStandAloneVmCommand) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *UpdateStandAloneVmCommand) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *UpdateStandAloneVmCommand) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetFlavorId returns the FlavorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetFlavorId() string {
	if o == nil || IsNil(o.FlavorId.Get()) {
		var ret string
		return ret
	}
	return *o.FlavorId.Get()
}

// GetFlavorIdOk returns a tuple with the FlavorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetFlavorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlavorId.Get(), o.FlavorId.IsSet()
}

// HasFlavorId returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasFlavorId() bool {
	if o != nil && o.FlavorId.IsSet() {
		return true
	}

	return false
}

// SetFlavorId gets a reference to the given NullableString and assigns it to the FlavorId field.
func (o *UpdateStandAloneVmCommand) SetFlavorId(v string) {
	o.FlavorId.Set(&v)
}

// SetFlavorIdNil sets the value for FlavorId to be an explicit nil
func (o *UpdateStandAloneVmCommand) SetFlavorIdNil() {
	o.FlavorId.Set(nil)
}

// UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
func (o *UpdateStandAloneVmCommand) UnsetFlavorId() {
	o.FlavorId.Unset()
}

// GetRevision returns the Revision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetRevision() int32 {
	if o == nil || IsNil(o.Revision.Get()) {
		var ret int32
		return ret
	}
	return *o.Revision.Get()
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetRevisionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revision.Get(), o.Revision.IsSet()
}

// HasRevision returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasRevision() bool {
	if o != nil && o.Revision.IsSet() {
		return true
	}

	return false
}

// SetRevision gets a reference to the given NullableInt32 and assigns it to the Revision field.
func (o *UpdateStandAloneVmCommand) SetRevision(v int32) {
	o.Revision.Set(&v)
}

// SetRevisionNil sets the value for Revision to be an explicit nil
func (o *UpdateStandAloneVmCommand) SetRevisionNil() {
	o.Revision.Set(nil)
}

// UnsetRevision ensures that no value is present for Revision, not even an explicit nil
func (o *UpdateStandAloneVmCommand) UnsetRevision() {
	o.Revision.Unset()
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetDisks() []UpdateStandAloneVmDiskDto {
	if o == nil {
		var ret []UpdateStandAloneVmDiskDto
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetDisksOk() ([]UpdateStandAloneVmDiskDto, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasDisks() bool {
	if o != nil && IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []UpdateStandAloneVmDiskDto and assigns it to the Disks field.
func (o *UpdateStandAloneVmCommand) SetDisks(v []UpdateStandAloneVmDiskDto) {
	o.Disks = v
}

func (o UpdateStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStandAloneVmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if o.PublicIp.IsSet() {
		toSerialize["publicIp"] = o.PublicIp.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["instanceId"] = o.InstanceId.Get()
	}
	if o.FlavorId.IsSet() {
		toSerialize["flavorId"] = o.FlavorId.Get()
	}
	if o.Revision.IsSet() {
		toSerialize["revision"] = o.Revision.Get()
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	return toSerialize, nil
}

type NullableUpdateStandAloneVmCommand struct {
	value *UpdateStandAloneVmCommand
	isSet bool
}

func (v NullableUpdateStandAloneVmCommand) Get() *UpdateStandAloneVmCommand {
	return v.value
}

func (v *NullableUpdateStandAloneVmCommand) Set(val *UpdateStandAloneVmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStandAloneVmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStandAloneVmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStandAloneVmCommand(val *UpdateStandAloneVmCommand) *NullableUpdateStandAloneVmCommand {
	return &NullableUpdateStandAloneVmCommand{value: val, isSet: true}
}

func (v NullableUpdateStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStandAloneVmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
