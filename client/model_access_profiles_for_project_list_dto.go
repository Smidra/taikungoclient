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

// checks if the AccessProfilesForProjectListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfilesForProjectListDto{}

// AccessProfilesForProjectListDto struct for AccessProfilesForProjectListDto
type AccessProfilesForProjectListDto struct {
	Id               *int32               `json:"id,omitempty"`
	Name             NullableString       `json:"name,omitempty"`
	HttpProxy        NullableString       `json:"httpProxy,omitempty"`
	OrganizationId   NullableInt32        `json:"organizationId,omitempty"`
	OrganizationName NullableString       `json:"organizationName,omitempty"`
	Revision         *int32               `json:"revision,omitempty"`
	SshUsers         []SshUserListDto     `json:"sshUsers,omitempty"`
	DnsServers       []DnsServerListDto   `json:"dnsServers,omitempty"`
	NtpServers       []NtpServerListDto   `json:"ntpServers,omitempty"`
	AllowedHosts     []AllowedHostListDto `json:"allowedHosts,omitempty"`
}

// NewAccessProfilesForProjectListDto instantiates a new AccessProfilesForProjectListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfilesForProjectListDto() *AccessProfilesForProjectListDto {
	this := AccessProfilesForProjectListDto{}
	return &this
}

// NewAccessProfilesForProjectListDtoWithDefaults instantiates a new AccessProfilesForProjectListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfilesForProjectListDtoWithDefaults() *AccessProfilesForProjectListDto {
	this := AccessProfilesForProjectListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessProfilesForProjectListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfilesForProjectListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AccessProfilesForProjectListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AccessProfilesForProjectListDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AccessProfilesForProjectListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AccessProfilesForProjectListDto) UnsetName() {
	o.Name.Unset()
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetHttpProxy() string {
	if o == nil || IsNil(o.HttpProxy.Get()) {
		var ret string
		return ret
	}
	return *o.HttpProxy.Get()
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetHttpProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpProxy.Get(), o.HttpProxy.IsSet()
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasHttpProxy() bool {
	if o != nil && o.HttpProxy.IsSet() {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given NullableString and assigns it to the HttpProxy field.
func (o *AccessProfilesForProjectListDto) SetHttpProxy(v string) {
	o.HttpProxy.Set(&v)
}

// SetHttpProxyNil sets the value for HttpProxy to be an explicit nil
func (o *AccessProfilesForProjectListDto) SetHttpProxyNil() {
	o.HttpProxy.Set(nil)
}

// UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
func (o *AccessProfilesForProjectListDto) UnsetHttpProxy() {
	o.HttpProxy.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *AccessProfilesForProjectListDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *AccessProfilesForProjectListDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *AccessProfilesForProjectListDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *AccessProfilesForProjectListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *AccessProfilesForProjectListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *AccessProfilesForProjectListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AccessProfilesForProjectListDto) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfilesForProjectListDto) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *AccessProfilesForProjectListDto) SetRevision(v int32) {
	o.Revision = &v
}

// GetSshUsers returns the SshUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetSshUsers() []SshUserListDto {
	if o == nil {
		var ret []SshUserListDto
		return ret
	}
	return o.SshUsers
}

// GetSshUsersOk returns a tuple with the SshUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetSshUsersOk() ([]SshUserListDto, bool) {
	if o == nil || IsNil(o.SshUsers) {
		return nil, false
	}
	return o.SshUsers, true
}

// HasSshUsers returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasSshUsers() bool {
	if o != nil && IsNil(o.SshUsers) {
		return true
	}

	return false
}

// SetSshUsers gets a reference to the given []SshUserListDto and assigns it to the SshUsers field.
func (o *AccessProfilesForProjectListDto) SetSshUsers(v []SshUserListDto) {
	o.SshUsers = v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetDnsServers() []DnsServerListDto {
	if o == nil {
		var ret []DnsServerListDto
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetDnsServersOk() ([]DnsServerListDto, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasDnsServers() bool {
	if o != nil && IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []DnsServerListDto and assigns it to the DnsServers field.
func (o *AccessProfilesForProjectListDto) SetDnsServers(v []DnsServerListDto) {
	o.DnsServers = v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetNtpServers() []NtpServerListDto {
	if o == nil {
		var ret []NtpServerListDto
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetNtpServersOk() ([]NtpServerListDto, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasNtpServers() bool {
	if o != nil && IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []NtpServerListDto and assigns it to the NtpServers field.
func (o *AccessProfilesForProjectListDto) SetNtpServers(v []NtpServerListDto) {
	o.NtpServers = v
}

// GetAllowedHosts returns the AllowedHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfilesForProjectListDto) GetAllowedHosts() []AllowedHostListDto {
	if o == nil {
		var ret []AllowedHostListDto
		return ret
	}
	return o.AllowedHosts
}

// GetAllowedHostsOk returns a tuple with the AllowedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfilesForProjectListDto) GetAllowedHostsOk() ([]AllowedHostListDto, bool) {
	if o == nil || IsNil(o.AllowedHosts) {
		return nil, false
	}
	return o.AllowedHosts, true
}

// HasAllowedHosts returns a boolean if a field has been set.
func (o *AccessProfilesForProjectListDto) HasAllowedHosts() bool {
	if o != nil && IsNil(o.AllowedHosts) {
		return true
	}

	return false
}

// SetAllowedHosts gets a reference to the given []AllowedHostListDto and assigns it to the AllowedHosts field.
func (o *AccessProfilesForProjectListDto) SetAllowedHosts(v []AllowedHostListDto) {
	o.AllowedHosts = v
}

func (o AccessProfilesForProjectListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfilesForProjectListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.HttpProxy.IsSet() {
		toSerialize["httpProxy"] = o.HttpProxy.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if o.SshUsers != nil {
		toSerialize["sshUsers"] = o.SshUsers
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.NtpServers != nil {
		toSerialize["ntpServers"] = o.NtpServers
	}
	if o.AllowedHosts != nil {
		toSerialize["allowedHosts"] = o.AllowedHosts
	}
	return toSerialize, nil
}

type NullableAccessProfilesForProjectListDto struct {
	value *AccessProfilesForProjectListDto
	isSet bool
}

func (v NullableAccessProfilesForProjectListDto) Get() *AccessProfilesForProjectListDto {
	return v.value
}

func (v *NullableAccessProfilesForProjectListDto) Set(val *AccessProfilesForProjectListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfilesForProjectListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfilesForProjectListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfilesForProjectListDto(val *AccessProfilesForProjectListDto) *NullableAccessProfilesForProjectListDto {
	return &NullableAccessProfilesForProjectListDto{value: val, isSet: true}
}

func (v NullableAccessProfilesForProjectListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfilesForProjectListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
