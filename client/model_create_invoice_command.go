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
	"time"
)

// checks if the CreateInvoiceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInvoiceCommand{}

// CreateInvoiceCommand struct for CreateInvoiceCommand
type CreateInvoiceCommand struct {
	Name                       NullableString `json:"name,omitempty"`
	OrganizationSubscriptionId *int32         `json:"organizationSubscriptionId,omitempty"`
	StartDate                  *time.Time     `json:"startDate,omitempty"`
	EndDate                    *time.Time     `json:"endDate,omitempty"`
	DueDate                    *time.Time     `json:"dueDate,omitempty"`
	IsPaid                     *bool          `json:"isPaid,omitempty"`
	RequiredPaymentAction      *bool          `json:"requiredPaymentAction,omitempty"`
	StripeInvoiceId            NullableString `json:"stripeInvoiceId,omitempty"`
	Price                      *float64       `json:"price,omitempty"`
}

// NewCreateInvoiceCommand instantiates a new CreateInvoiceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvoiceCommand() *CreateInvoiceCommand {
	this := CreateInvoiceCommand{}
	return &this
}

// NewCreateInvoiceCommandWithDefaults instantiates a new CreateInvoiceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvoiceCommandWithDefaults() *CreateInvoiceCommand {
	this := CreateInvoiceCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateInvoiceCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateInvoiceCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateInvoiceCommand) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetOrganizationSubscriptionId() int32 {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		var ret int32
		return ret
	}
	return *o.OrganizationSubscriptionId
}

// GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetOrganizationSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationSubscriptionId) {
		return nil, false
	}
	return o.OrganizationSubscriptionId, true
}

// HasOrganizationSubscriptionId returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasOrganizationSubscriptionId() bool {
	if o != nil && !IsNil(o.OrganizationSubscriptionId) {
		return true
	}

	return false
}

// SetOrganizationSubscriptionId gets a reference to the given int32 and assigns it to the OrganizationSubscriptionId field.
func (o *CreateInvoiceCommand) SetOrganizationSubscriptionId(v int32) {
	o.OrganizationSubscriptionId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CreateInvoiceCommand) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CreateInvoiceCommand) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *CreateInvoiceCommand) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetIsPaid returns the IsPaid field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetIsPaid() bool {
	if o == nil || IsNil(o.IsPaid) {
		var ret bool
		return ret
	}
	return *o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetIsPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaid) {
		return nil, false
	}
	return o.IsPaid, true
}

// HasIsPaid returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasIsPaid() bool {
	if o != nil && !IsNil(o.IsPaid) {
		return true
	}

	return false
}

// SetIsPaid gets a reference to the given bool and assigns it to the IsPaid field.
func (o *CreateInvoiceCommand) SetIsPaid(v bool) {
	o.IsPaid = &v
}

// GetRequiredPaymentAction returns the RequiredPaymentAction field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetRequiredPaymentAction() bool {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		var ret bool
		return ret
	}
	return *o.RequiredPaymentAction
}

// GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetRequiredPaymentActionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredPaymentAction) {
		return nil, false
	}
	return o.RequiredPaymentAction, true
}

// HasRequiredPaymentAction returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasRequiredPaymentAction() bool {
	if o != nil && !IsNil(o.RequiredPaymentAction) {
		return true
	}

	return false
}

// SetRequiredPaymentAction gets a reference to the given bool and assigns it to the RequiredPaymentAction field.
func (o *CreateInvoiceCommand) SetRequiredPaymentAction(v bool) {
	o.RequiredPaymentAction = &v
}

// GetStripeInvoiceId returns the StripeInvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceCommand) GetStripeInvoiceId() string {
	if o == nil || IsNil(o.StripeInvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.StripeInvoiceId.Get()
}

// GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceCommand) GetStripeInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StripeInvoiceId.Get(), o.StripeInvoiceId.IsSet()
}

// HasStripeInvoiceId returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasStripeInvoiceId() bool {
	if o != nil && o.StripeInvoiceId.IsSet() {
		return true
	}

	return false
}

// SetStripeInvoiceId gets a reference to the given NullableString and assigns it to the StripeInvoiceId field.
func (o *CreateInvoiceCommand) SetStripeInvoiceId(v string) {
	o.StripeInvoiceId.Set(&v)
}

// SetStripeInvoiceIdNil sets the value for StripeInvoiceId to be an explicit nil
func (o *CreateInvoiceCommand) SetStripeInvoiceIdNil() {
	o.StripeInvoiceId.Set(nil)
}

// UnsetStripeInvoiceId ensures that no value is present for StripeInvoiceId, not even an explicit nil
func (o *CreateInvoiceCommand) UnsetStripeInvoiceId() {
	o.StripeInvoiceId.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateInvoiceCommand) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceCommand) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateInvoiceCommand) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *CreateInvoiceCommand) SetPrice(v float64) {
	o.Price = &v
}

func (o CreateInvoiceCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInvoiceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OrganizationSubscriptionId) {
		toSerialize["organizationSubscriptionId"] = o.OrganizationSubscriptionId
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.IsPaid) {
		toSerialize["isPaid"] = o.IsPaid
	}
	if !IsNil(o.RequiredPaymentAction) {
		toSerialize["requiredPaymentAction"] = o.RequiredPaymentAction
	}
	if o.StripeInvoiceId.IsSet() {
		toSerialize["stripeInvoiceId"] = o.StripeInvoiceId.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableCreateInvoiceCommand struct {
	value *CreateInvoiceCommand
	isSet bool
}

func (v NullableCreateInvoiceCommand) Get() *CreateInvoiceCommand {
	return v.value
}

func (v *NullableCreateInvoiceCommand) Set(val *CreateInvoiceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceCommand(val *CreateInvoiceCommand) *NullableCreateInvoiceCommand {
	return &NullableCreateInvoiceCommand{value: val, isSet: true}
}

func (v NullableCreateInvoiceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
