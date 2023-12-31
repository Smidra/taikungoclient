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

// checks if the AdminOrganizationsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminOrganizationsListDto{}

// AdminOrganizationsListDto struct for AdminOrganizationsListDto
type AdminOrganizationsListDto struct {
	Id          *int32         `json:"id,omitempty"`
	Name        NullableString `json:"name,omitempty"`
	CustomerId  NullableString `json:"customerId,omitempty"`
	PartnerName NullableString `json:"partnerName,omitempty"`
	PartnerLogo NullableString `json:"partnerLogo,omitempty"`
}

// NewAdminOrganizationsListDto instantiates a new AdminOrganizationsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminOrganizationsListDto() *AdminOrganizationsListDto {
	this := AdminOrganizationsListDto{}
	return &this
}

// NewAdminOrganizationsListDtoWithDefaults instantiates a new AdminOrganizationsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminOrganizationsListDtoWithDefaults() *AdminOrganizationsListDto {
	this := AdminOrganizationsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminOrganizationsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminOrganizationsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminOrganizationsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AdminOrganizationsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminOrganizationsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminOrganizationsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AdminOrganizationsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AdminOrganizationsListDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AdminOrganizationsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AdminOrganizationsListDto) UnsetName() {
	o.Name.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminOrganizationsListDto) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminOrganizationsListDto) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AdminOrganizationsListDto) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *AdminOrganizationsListDto) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *AdminOrganizationsListDto) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *AdminOrganizationsListDto) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminOrganizationsListDto) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerName.Get()
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminOrganizationsListDto) GetPartnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerName.Get(), o.PartnerName.IsSet()
}

// HasPartnerName returns a boolean if a field has been set.
func (o *AdminOrganizationsListDto) HasPartnerName() bool {
	if o != nil && o.PartnerName.IsSet() {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given NullableString and assigns it to the PartnerName field.
func (o *AdminOrganizationsListDto) SetPartnerName(v string) {
	o.PartnerName.Set(&v)
}

// SetPartnerNameNil sets the value for PartnerName to be an explicit nil
func (o *AdminOrganizationsListDto) SetPartnerNameNil() {
	o.PartnerName.Set(nil)
}

// UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
func (o *AdminOrganizationsListDto) UnsetPartnerName() {
	o.PartnerName.Unset()
}

// GetPartnerLogo returns the PartnerLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminOrganizationsListDto) GetPartnerLogo() string {
	if o == nil || IsNil(o.PartnerLogo.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerLogo.Get()
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminOrganizationsListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerLogo.Get(), o.PartnerLogo.IsSet()
}

// HasPartnerLogo returns a boolean if a field has been set.
func (o *AdminOrganizationsListDto) HasPartnerLogo() bool {
	if o != nil && o.PartnerLogo.IsSet() {
		return true
	}

	return false
}

// SetPartnerLogo gets a reference to the given NullableString and assigns it to the PartnerLogo field.
func (o *AdminOrganizationsListDto) SetPartnerLogo(v string) {
	o.PartnerLogo.Set(&v)
}

// SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil
func (o *AdminOrganizationsListDto) SetPartnerLogoNil() {
	o.PartnerLogo.Set(nil)
}

// UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
func (o *AdminOrganizationsListDto) UnsetPartnerLogo() {
	o.PartnerLogo.Unset()
}

func (o AdminOrganizationsListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminOrganizationsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if o.PartnerName.IsSet() {
		toSerialize["partnerName"] = o.PartnerName.Get()
	}
	if o.PartnerLogo.IsSet() {
		toSerialize["partnerLogo"] = o.PartnerLogo.Get()
	}
	return toSerialize, nil
}

type NullableAdminOrganizationsListDto struct {
	value *AdminOrganizationsListDto
	isSet bool
}

func (v NullableAdminOrganizationsListDto) Get() *AdminOrganizationsListDto {
	return v.value
}

func (v *NullableAdminOrganizationsListDto) Set(val *AdminOrganizationsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminOrganizationsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminOrganizationsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminOrganizationsListDto(val *AdminOrganizationsListDto) *NullableAdminOrganizationsListDto {
	return &NullableAdminOrganizationsListDto{value: val, isSet: true}
}

func (v NullableAdminOrganizationsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminOrganizationsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
