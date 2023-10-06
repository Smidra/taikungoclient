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

// checks if the OpenstackCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackCredentialsListDto{}

// OpenstackCredentialsListDto struct for OpenstackCredentialsListDto
type OpenstackCredentialsListDto struct {
	Id                     *int32              `json:"id,omitempty"`
	ProjectCount           *int32              `json:"projectCount,omitempty"`
	IsLocked               *bool               `json:"isLocked,omitempty"`
	Name                   NullableString      `json:"name,omitempty"`
	User                   NullableString      `json:"user,omitempty"`
	Url                    NullableString      `json:"url,omitempty"`
	Project                NullableString      `json:"project,omitempty"`
	Domain                 NullableString      `json:"domain,omitempty"`
	Region                 NullableString      `json:"region,omitempty"`
	PublicNetwork          NullableString      `json:"publicNetwork,omitempty"`
	ImportNetwork          *bool               `json:"importNetwork,omitempty"`
	TenantId               NullableString      `json:"tenantId,omitempty"`
	AvailabilityZone       NullableString      `json:"availabilityZone,omitempty"`
	VolumeType             NullableString      `json:"volumeType,omitempty"`
	InternalSubnetId       NullableString      `json:"internalSubnetId,omitempty"`
	Projects               []CommonDropdownDto `json:"projects,omitempty"`
	CreatedBy              NullableString      `json:"createdBy,omitempty"`
	LastModified           NullableString      `json:"lastModified,omitempty"`
	LastModifiedBy         NullableString      `json:"lastModifiedBy,omitempty"`
	IsDefault              *bool               `json:"isDefault,omitempty"`
	OrganizationId         *int32              `json:"organizationId,omitempty"`
	OrganizationName       NullableString      `json:"organizationName,omitempty"`
	CreatedAt              NullableString      `json:"createdAt,omitempty"`
	ContinentName          NullableString      `json:"continentName,omitempty"`
	IsAdmin                *bool               `json:"isAdmin,omitempty"`
	IsInfra                *bool               `json:"isInfra,omitempty"`
	ApplicationCredEnabled *bool               `json:"applicationCredEnabled,omitempty"`
}

// NewOpenstackCredentialsListDto instantiates a new OpenstackCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackCredentialsListDto() *OpenstackCredentialsListDto {
	this := OpenstackCredentialsListDto{}
	return &this
}

// NewOpenstackCredentialsListDtoWithDefaults instantiates a new OpenstackCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackCredentialsListDtoWithDefaults() *OpenstackCredentialsListDto {
	this := OpenstackCredentialsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpenstackCredentialsListDto) SetId(v int32) {
	o.Id = &v
}

// GetProjectCount returns the ProjectCount field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetProjectCount() int32 {
	if o == nil || IsNil(o.ProjectCount) {
		var ret int32
		return ret
	}
	return *o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectCount) {
		return nil, false
	}
	return o.ProjectCount, true
}

// HasProjectCount returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasProjectCount() bool {
	if o != nil && !IsNil(o.ProjectCount) {
		return true
	}

	return false
}

// SetProjectCount gets a reference to the given int32 and assigns it to the ProjectCount field.
func (o *OpenstackCredentialsListDto) SetProjectCount(v int32) {
	o.ProjectCount = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *OpenstackCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpenstackCredentialsListDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *OpenstackCredentialsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetName() {
	o.Name.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *OpenstackCredentialsListDto) SetUser(v string) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *OpenstackCredentialsListDto) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetUser() {
	o.User.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *OpenstackCredentialsListDto) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *OpenstackCredentialsListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *OpenstackCredentialsListDto) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *OpenstackCredentialsListDto) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetProject() {
	o.Project.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *OpenstackCredentialsListDto) SetDomain(v string) {
	o.Domain.Set(&v)
}

// SetDomainNil sets the value for Domain to be an explicit nil
func (o *OpenstackCredentialsListDto) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetDomain() {
	o.Domain.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *OpenstackCredentialsListDto) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *OpenstackCredentialsListDto) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetRegion() {
	o.Region.Unset()
}

// GetPublicNetwork returns the PublicNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetPublicNetwork() string {
	if o == nil || IsNil(o.PublicNetwork.Get()) {
		var ret string
		return ret
	}
	return *o.PublicNetwork.Get()
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetPublicNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicNetwork.Get(), o.PublicNetwork.IsSet()
}

// HasPublicNetwork returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasPublicNetwork() bool {
	if o != nil && o.PublicNetwork.IsSet() {
		return true
	}

	return false
}

// SetPublicNetwork gets a reference to the given NullableString and assigns it to the PublicNetwork field.
func (o *OpenstackCredentialsListDto) SetPublicNetwork(v string) {
	o.PublicNetwork.Set(&v)
}

// SetPublicNetworkNil sets the value for PublicNetwork to be an explicit nil
func (o *OpenstackCredentialsListDto) SetPublicNetworkNil() {
	o.PublicNetwork.Set(nil)
}

// UnsetPublicNetwork ensures that no value is present for PublicNetwork, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetPublicNetwork() {
	o.PublicNetwork.Unset()
}

// GetImportNetwork returns the ImportNetwork field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetImportNetwork() bool {
	if o == nil || IsNil(o.ImportNetwork) {
		var ret bool
		return ret
	}
	return *o.ImportNetwork
}

// GetImportNetworkOk returns a tuple with the ImportNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetImportNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportNetwork) {
		return nil, false
	}
	return o.ImportNetwork, true
}

// HasImportNetwork returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasImportNetwork() bool {
	if o != nil && !IsNil(o.ImportNetwork) {
		return true
	}

	return false
}

// SetImportNetwork gets a reference to the given bool and assigns it to the ImportNetwork field.
func (o *OpenstackCredentialsListDto) SetImportNetwork(v bool) {
	o.ImportNetwork = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *OpenstackCredentialsListDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *OpenstackCredentialsListDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone.Get()) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone.Get()
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZone.Get(), o.AvailabilityZone.IsSet()
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given NullableString and assigns it to the AvailabilityZone field.
func (o *OpenstackCredentialsListDto) SetAvailabilityZone(v string) {
	o.AvailabilityZone.Set(&v)
}

// SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil
func (o *OpenstackCredentialsListDto) SetAvailabilityZoneNil() {
	o.AvailabilityZone.Set(nil)
}

// UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetAvailabilityZone() {
	o.AvailabilityZone.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType.Get()) {
		var ret string
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullableString and assigns it to the VolumeType field.
func (o *OpenstackCredentialsListDto) SetVolumeType(v string) {
	o.VolumeType.Set(&v)
}

// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *OpenstackCredentialsListDto) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetInternalSubnetId returns the InternalSubnetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetInternalSubnetId() string {
	if o == nil || IsNil(o.InternalSubnetId.Get()) {
		var ret string
		return ret
	}
	return *o.InternalSubnetId.Get()
}

// GetInternalSubnetIdOk returns a tuple with the InternalSubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetInternalSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalSubnetId.Get(), o.InternalSubnetId.IsSet()
}

// HasInternalSubnetId returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasInternalSubnetId() bool {
	if o != nil && o.InternalSubnetId.IsSet() {
		return true
	}

	return false
}

// SetInternalSubnetId gets a reference to the given NullableString and assigns it to the InternalSubnetId field.
func (o *OpenstackCredentialsListDto) SetInternalSubnetId(v string) {
	o.InternalSubnetId.Set(&v)
}

// SetInternalSubnetIdNil sets the value for InternalSubnetId to be an explicit nil
func (o *OpenstackCredentialsListDto) SetInternalSubnetIdNil() {
	o.InternalSubnetId.Set(nil)
}

// UnsetInternalSubnetId ensures that no value is present for InternalSubnetId, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetInternalSubnetId() {
	o.InternalSubnetId.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *OpenstackCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *OpenstackCredentialsListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *OpenstackCredentialsListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *OpenstackCredentialsListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *OpenstackCredentialsListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *OpenstackCredentialsListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *OpenstackCredentialsListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *OpenstackCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *OpenstackCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *OpenstackCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *OpenstackCredentialsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *OpenstackCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *OpenstackCredentialsListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *OpenstackCredentialsListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}

// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *OpenstackCredentialsListDto) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *OpenstackCredentialsListDto) UnsetContinentName() {
	o.ContinentName.Unset()
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *OpenstackCredentialsListDto) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsInfra returns the IsInfra field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetIsInfra() bool {
	if o == nil || IsNil(o.IsInfra) {
		var ret bool
		return ret
	}
	return *o.IsInfra
}

// GetIsInfraOk returns a tuple with the IsInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetIsInfraOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInfra) {
		return nil, false
	}
	return o.IsInfra, true
}

// HasIsInfra returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasIsInfra() bool {
	if o != nil && !IsNil(o.IsInfra) {
		return true
	}

	return false
}

// SetIsInfra gets a reference to the given bool and assigns it to the IsInfra field.
func (o *OpenstackCredentialsListDto) SetIsInfra(v bool) {
	o.IsInfra = &v
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenstackCredentialsListDto) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsListDto) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenstackCredentialsListDto) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenstackCredentialsListDto) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

func (o OpenstackCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectCount) {
		toSerialize["projectCount"] = o.ProjectCount
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.PublicNetwork.IsSet() {
		toSerialize["publicNetwork"] = o.PublicNetwork.Get()
	}
	if !IsNil(o.ImportNetwork) {
		toSerialize["importNetwork"] = o.ImportNetwork
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.AvailabilityZone.IsSet() {
		toSerialize["availabilityZone"] = o.AvailabilityZone.Get()
	}
	if o.VolumeType.IsSet() {
		toSerialize["volumeType"] = o.VolumeType.Get()
	}
	if o.InternalSubnetId.IsSet() {
		toSerialize["internalSubnetId"] = o.InternalSubnetId.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.ContinentName.IsSet() {
		toSerialize["continentName"] = o.ContinentName.Get()
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !IsNil(o.IsInfra) {
		toSerialize["isInfra"] = o.IsInfra
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	return toSerialize, nil
}

type NullableOpenstackCredentialsListDto struct {
	value *OpenstackCredentialsListDto
	isSet bool
}

func (v NullableOpenstackCredentialsListDto) Get() *OpenstackCredentialsListDto {
	return v.value
}

func (v *NullableOpenstackCredentialsListDto) Set(val *OpenstackCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackCredentialsListDto(val *OpenstackCredentialsListDto) *NullableOpenstackCredentialsListDto {
	return &NullableOpenstackCredentialsListDto{value: val, isSet: true}
}

func (v NullableOpenstackCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}