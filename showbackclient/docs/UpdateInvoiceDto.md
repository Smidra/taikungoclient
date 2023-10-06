# UpdateInvoiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationSubscriptionId** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**IsPaid** | Pointer to **NullableBool** |  | [optional] 
**RequiredPaymentAction** | Pointer to **NullableBool** |  | [optional] 
**StripeInvoiceId** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewUpdateInvoiceDto

`func NewUpdateInvoiceDto() *UpdateInvoiceDto`

NewUpdateInvoiceDto instantiates a new UpdateInvoiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInvoiceDtoWithDefaults

`func NewUpdateInvoiceDtoWithDefaults() *UpdateInvoiceDto`

NewUpdateInvoiceDtoWithDefaults instantiates a new UpdateInvoiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateInvoiceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInvoiceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInvoiceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInvoiceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateInvoiceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateInvoiceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationSubscriptionId

`func (o *UpdateInvoiceDto) GetOrganizationSubscriptionId() int32`

GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field if non-nil, zero value otherwise.

### GetOrganizationSubscriptionIdOk

`func (o *UpdateInvoiceDto) GetOrganizationSubscriptionIdOk() (*int32, bool)`

GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSubscriptionId

`func (o *UpdateInvoiceDto) SetOrganizationSubscriptionId(v int32)`

SetOrganizationSubscriptionId sets OrganizationSubscriptionId field to given value.

### HasOrganizationSubscriptionId

`func (o *UpdateInvoiceDto) HasOrganizationSubscriptionId() bool`

HasOrganizationSubscriptionId returns a boolean if a field has been set.

### SetOrganizationSubscriptionIdNil

`func (o *UpdateInvoiceDto) SetOrganizationSubscriptionIdNil(b bool)`

 SetOrganizationSubscriptionIdNil sets the value for OrganizationSubscriptionId to be an explicit nil

### UnsetOrganizationSubscriptionId
`func (o *UpdateInvoiceDto) UnsetOrganizationSubscriptionId()`

UnsetOrganizationSubscriptionId ensures that no value is present for OrganizationSubscriptionId, not even an explicit nil
### GetStartDate

`func (o *UpdateInvoiceDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateInvoiceDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateInvoiceDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateInvoiceDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UpdateInvoiceDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UpdateInvoiceDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *UpdateInvoiceDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateInvoiceDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateInvoiceDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateInvoiceDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateInvoiceDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateInvoiceDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDueDate

`func (o *UpdateInvoiceDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UpdateInvoiceDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UpdateInvoiceDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UpdateInvoiceDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *UpdateInvoiceDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *UpdateInvoiceDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetIsPaid

`func (o *UpdateInvoiceDto) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UpdateInvoiceDto) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UpdateInvoiceDto) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UpdateInvoiceDto) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### SetIsPaidNil

`func (o *UpdateInvoiceDto) SetIsPaidNil(b bool)`

 SetIsPaidNil sets the value for IsPaid to be an explicit nil

### UnsetIsPaid
`func (o *UpdateInvoiceDto) UnsetIsPaid()`

UnsetIsPaid ensures that no value is present for IsPaid, not even an explicit nil
### GetRequiredPaymentAction

`func (o *UpdateInvoiceDto) GetRequiredPaymentAction() bool`

GetRequiredPaymentAction returns the RequiredPaymentAction field if non-nil, zero value otherwise.

### GetRequiredPaymentActionOk

`func (o *UpdateInvoiceDto) GetRequiredPaymentActionOk() (*bool, bool)`

GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPaymentAction

`func (o *UpdateInvoiceDto) SetRequiredPaymentAction(v bool)`

SetRequiredPaymentAction sets RequiredPaymentAction field to given value.

### HasRequiredPaymentAction

`func (o *UpdateInvoiceDto) HasRequiredPaymentAction() bool`

HasRequiredPaymentAction returns a boolean if a field has been set.

### SetRequiredPaymentActionNil

`func (o *UpdateInvoiceDto) SetRequiredPaymentActionNil(b bool)`

 SetRequiredPaymentActionNil sets the value for RequiredPaymentAction to be an explicit nil

### UnsetRequiredPaymentAction
`func (o *UpdateInvoiceDto) UnsetRequiredPaymentAction()`

UnsetRequiredPaymentAction ensures that no value is present for RequiredPaymentAction, not even an explicit nil
### GetStripeInvoiceId

`func (o *UpdateInvoiceDto) GetStripeInvoiceId() string`

GetStripeInvoiceId returns the StripeInvoiceId field if non-nil, zero value otherwise.

### GetStripeInvoiceIdOk

`func (o *UpdateInvoiceDto) GetStripeInvoiceIdOk() (*string, bool)`

GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoiceId

`func (o *UpdateInvoiceDto) SetStripeInvoiceId(v string)`

SetStripeInvoiceId sets StripeInvoiceId field to given value.

### HasStripeInvoiceId

`func (o *UpdateInvoiceDto) HasStripeInvoiceId() bool`

HasStripeInvoiceId returns a boolean if a field has been set.

### SetStripeInvoiceIdNil

`func (o *UpdateInvoiceDto) SetStripeInvoiceIdNil(b bool)`

 SetStripeInvoiceIdNil sets the value for StripeInvoiceId to be an explicit nil

### UnsetStripeInvoiceId
`func (o *UpdateInvoiceDto) UnsetStripeInvoiceId()`

UnsetStripeInvoiceId ensures that no value is present for StripeInvoiceId, not even an explicit nil
### GetPrice

`func (o *UpdateInvoiceDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateInvoiceDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateInvoiceDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateInvoiceDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *UpdateInvoiceDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *UpdateInvoiceDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


