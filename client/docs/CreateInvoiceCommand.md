# CreateInvoiceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationSubscriptionId** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**IsPaid** | Pointer to **bool** |  | [optional] 
**RequiredPaymentAction** | Pointer to **bool** |  | [optional] 
**StripeInvoiceId** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 

## Methods

### NewCreateInvoiceCommand

`func NewCreateInvoiceCommand() *CreateInvoiceCommand`

NewCreateInvoiceCommand instantiates a new CreateInvoiceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceCommandWithDefaults

`func NewCreateInvoiceCommandWithDefaults() *CreateInvoiceCommand`

NewCreateInvoiceCommandWithDefaults instantiates a new CreateInvoiceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInvoiceCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInvoiceCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInvoiceCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInvoiceCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateInvoiceCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateInvoiceCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationSubscriptionId

`func (o *CreateInvoiceCommand) GetOrganizationSubscriptionId() int32`

GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field if non-nil, zero value otherwise.

### GetOrganizationSubscriptionIdOk

`func (o *CreateInvoiceCommand) GetOrganizationSubscriptionIdOk() (*int32, bool)`

GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSubscriptionId

`func (o *CreateInvoiceCommand) SetOrganizationSubscriptionId(v int32)`

SetOrganizationSubscriptionId sets OrganizationSubscriptionId field to given value.

### HasOrganizationSubscriptionId

`func (o *CreateInvoiceCommand) HasOrganizationSubscriptionId() bool`

HasOrganizationSubscriptionId returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateInvoiceCommand) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateInvoiceCommand) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateInvoiceCommand) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateInvoiceCommand) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateInvoiceCommand) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateInvoiceCommand) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateInvoiceCommand) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateInvoiceCommand) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDueDate

`func (o *CreateInvoiceCommand) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreateInvoiceCommand) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreateInvoiceCommand) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CreateInvoiceCommand) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetIsPaid

`func (o *CreateInvoiceCommand) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *CreateInvoiceCommand) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *CreateInvoiceCommand) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *CreateInvoiceCommand) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetRequiredPaymentAction

`func (o *CreateInvoiceCommand) GetRequiredPaymentAction() bool`

GetRequiredPaymentAction returns the RequiredPaymentAction field if non-nil, zero value otherwise.

### GetRequiredPaymentActionOk

`func (o *CreateInvoiceCommand) GetRequiredPaymentActionOk() (*bool, bool)`

GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPaymentAction

`func (o *CreateInvoiceCommand) SetRequiredPaymentAction(v bool)`

SetRequiredPaymentAction sets RequiredPaymentAction field to given value.

### HasRequiredPaymentAction

`func (o *CreateInvoiceCommand) HasRequiredPaymentAction() bool`

HasRequiredPaymentAction returns a boolean if a field has been set.

### GetStripeInvoiceId

`func (o *CreateInvoiceCommand) GetStripeInvoiceId() string`

GetStripeInvoiceId returns the StripeInvoiceId field if non-nil, zero value otherwise.

### GetStripeInvoiceIdOk

`func (o *CreateInvoiceCommand) GetStripeInvoiceIdOk() (*string, bool)`

GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoiceId

`func (o *CreateInvoiceCommand) SetStripeInvoiceId(v string)`

SetStripeInvoiceId sets StripeInvoiceId field to given value.

### HasStripeInvoiceId

`func (o *CreateInvoiceCommand) HasStripeInvoiceId() bool`

HasStripeInvoiceId returns a boolean if a field has been set.

### SetStripeInvoiceIdNil

`func (o *CreateInvoiceCommand) SetStripeInvoiceIdNil(b bool)`

 SetStripeInvoiceIdNil sets the value for StripeInvoiceId to be an explicit nil

### UnsetStripeInvoiceId
`func (o *CreateInvoiceCommand) UnsetStripeInvoiceId()`

UnsetStripeInvoiceId ensures that no value is present for StripeInvoiceId, not even an explicit nil
### GetPrice

`func (o *CreateInvoiceCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateInvoiceCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateInvoiceCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateInvoiceCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


