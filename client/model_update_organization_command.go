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

// checks if the UpdateOrganizationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationCommand{}

// UpdateOrganizationCommand struct for UpdateOrganizationCommand
type UpdateOrganizationCommand struct {
	Id                           *int32          `json:"id,omitempty"`
	Name                         NullableString  `json:"name,omitempty"`
	FullName                     NullableString  `json:"fullName,omitempty"`
	Phone                        NullableString  `json:"phone,omitempty"`
	Email                        NullableString  `json:"email,omitempty"`
	Address                      NullableString  `json:"address,omitempty"`
	Country                      NullableString  `json:"country,omitempty"`
	City                         NullableString  `json:"city,omitempty"`
	VatNumber                    NullableString  `json:"vatNumber,omitempty"`
	IsLocked                     *bool           `json:"isLocked,omitempty"`
	BillingEmail                 NullableString  `json:"billingEmail,omitempty"`
	DiscountRate                 NullableFloat64 `json:"discountRate,omitempty"`
	IsEligibleUpdateSubscription *bool           `json:"isEligibleUpdateSubscription,omitempty"`
}

// NewUpdateOrganizationCommand instantiates a new UpdateOrganizationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationCommand() *UpdateOrganizationCommand {
	this := UpdateOrganizationCommand{}
	return &this
}

// NewUpdateOrganizationCommandWithDefaults instantiates a new UpdateOrganizationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationCommandWithDefaults() *UpdateOrganizationCommand {
	this := UpdateOrganizationCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateOrganizationCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateOrganizationCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateOrganizationCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateOrganizationCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetName() {
	o.Name.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *UpdateOrganizationCommand) SetFullName(v string) {
	o.FullName.Set(&v)
}

// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *UpdateOrganizationCommand) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetFullName() {
	o.FullName.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *UpdateOrganizationCommand) SetPhone(v string) {
	o.Phone.Set(&v)
}

// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *UpdateOrganizationCommand) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetPhone() {
	o.Phone.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateOrganizationCommand) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateOrganizationCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetEmail() {
	o.Email.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *UpdateOrganizationCommand) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *UpdateOrganizationCommand) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetAddress() {
	o.Address.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *UpdateOrganizationCommand) SetCountry(v string) {
	o.Country.Set(&v)
}

// SetCountryNil sets the value for Country to be an explicit nil
func (o *UpdateOrganizationCommand) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetCountry() {
	o.Country.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *UpdateOrganizationCommand) SetCity(v string) {
	o.City.Set(&v)
}

// SetCityNil sets the value for City to be an explicit nil
func (o *UpdateOrganizationCommand) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetCity() {
	o.City.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *UpdateOrganizationCommand) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}

// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *UpdateOrganizationCommand) SetVatNumberNil() {
	o.VatNumber.Set(nil)
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetVatNumber() {
	o.VatNumber.Unset()
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *UpdateOrganizationCommand) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCommand) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *UpdateOrganizationCommand) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasBillingEmail() bool {
	if o != nil && o.BillingEmail.IsSet() {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given NullableString and assigns it to the BillingEmail field.
func (o *UpdateOrganizationCommand) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}

// SetBillingEmailNil sets the value for BillingEmail to be an explicit nil
func (o *UpdateOrganizationCommand) SetBillingEmailNil() {
	o.BillingEmail.Set(nil)
}

// UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetBillingEmail() {
	o.BillingEmail.Unset()
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationCommand) GetDiscountRate() float64 {
	if o == nil || IsNil(o.DiscountRate.Get()) {
		var ret float64
		return ret
	}
	return *o.DiscountRate.Get()
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationCommand) GetDiscountRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountRate.Get(), o.DiscountRate.IsSet()
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasDiscountRate() bool {
	if o != nil && o.DiscountRate.IsSet() {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given NullableFloat64 and assigns it to the DiscountRate field.
func (o *UpdateOrganizationCommand) SetDiscountRate(v float64) {
	o.DiscountRate.Set(&v)
}

// SetDiscountRateNil sets the value for DiscountRate to be an explicit nil
func (o *UpdateOrganizationCommand) SetDiscountRateNil() {
	o.DiscountRate.Set(nil)
}

// UnsetDiscountRate ensures that no value is present for DiscountRate, not even an explicit nil
func (o *UpdateOrganizationCommand) UnsetDiscountRate() {
	o.DiscountRate.Unset()
}

// GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field value if set, zero value otherwise.
func (o *UpdateOrganizationCommand) GetIsEligibleUpdateSubscription() bool {
	if o == nil || IsNil(o.IsEligibleUpdateSubscription) {
		var ret bool
		return ret
	}
	return *o.IsEligibleUpdateSubscription
}

// GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCommand) GetIsEligibleUpdateSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligibleUpdateSubscription) {
		return nil, false
	}
	return o.IsEligibleUpdateSubscription, true
}

// HasIsEligibleUpdateSubscription returns a boolean if a field has been set.
func (o *UpdateOrganizationCommand) HasIsEligibleUpdateSubscription() bool {
	if o != nil && !IsNil(o.IsEligibleUpdateSubscription) {
		return true
	}

	return false
}

// SetIsEligibleUpdateSubscription gets a reference to the given bool and assigns it to the IsEligibleUpdateSubscription field.
func (o *UpdateOrganizationCommand) SetIsEligibleUpdateSubscription(v bool) {
	o.IsEligibleUpdateSubscription = &v
}

func (o UpdateOrganizationCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vatNumber"] = o.VatNumber.Get()
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.BillingEmail.IsSet() {
		toSerialize["billingEmail"] = o.BillingEmail.Get()
	}
	if o.DiscountRate.IsSet() {
		toSerialize["discountRate"] = o.DiscountRate.Get()
	}
	if !IsNil(o.IsEligibleUpdateSubscription) {
		toSerialize["isEligibleUpdateSubscription"] = o.IsEligibleUpdateSubscription
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationCommand struct {
	value *UpdateOrganizationCommand
	isSet bool
}

func (v NullableUpdateOrganizationCommand) Get() *UpdateOrganizationCommand {
	return v.value
}

func (v *NullableUpdateOrganizationCommand) Set(val *UpdateOrganizationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationCommand(val *UpdateOrganizationCommand) *NullableUpdateOrganizationCommand {
	return &NullableUpdateOrganizationCommand{value: val, isSet: true}
}

func (v NullableUpdateOrganizationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
