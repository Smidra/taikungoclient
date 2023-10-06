# InvoiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**OrganizationSubscriptionId** | Pointer to **int32** |  | [optional] 
**IsPaid** | Pointer to **bool** |  | [optional] 
**RequiredPaymentAction** | Pointer to **bool** |  | [optional] 
**StripeInvoiceId** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoiceDto

`func NewInvoiceDto() *InvoiceDto`

NewInvoiceDto instantiates a new InvoiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDtoWithDefaults

`func NewInvoiceDtoWithDefaults() *InvoiceDto`

NewInvoiceDtoWithDefaults instantiates a new InvoiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InvoiceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InvoiceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InvoiceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDocumentNumber

`func (o *InvoiceDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *InvoiceDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *InvoiceDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *InvoiceDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *InvoiceDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *InvoiceDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetOrganizationSubscriptionId

`func (o *InvoiceDto) GetOrganizationSubscriptionId() int32`

GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field if non-nil, zero value otherwise.

### GetOrganizationSubscriptionIdOk

`func (o *InvoiceDto) GetOrganizationSubscriptionIdOk() (*int32, bool)`

GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSubscriptionId

`func (o *InvoiceDto) SetOrganizationSubscriptionId(v int32)`

SetOrganizationSubscriptionId sets OrganizationSubscriptionId field to given value.

### HasOrganizationSubscriptionId

`func (o *InvoiceDto) HasOrganizationSubscriptionId() bool`

HasOrganizationSubscriptionId returns a boolean if a field has been set.

### GetIsPaid

`func (o *InvoiceDto) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *InvoiceDto) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *InvoiceDto) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *InvoiceDto) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetRequiredPaymentAction

`func (o *InvoiceDto) GetRequiredPaymentAction() bool`

GetRequiredPaymentAction returns the RequiredPaymentAction field if non-nil, zero value otherwise.

### GetRequiredPaymentActionOk

`func (o *InvoiceDto) GetRequiredPaymentActionOk() (*bool, bool)`

GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPaymentAction

`func (o *InvoiceDto) SetRequiredPaymentAction(v bool)`

SetRequiredPaymentAction sets RequiredPaymentAction field to given value.

### HasRequiredPaymentAction

`func (o *InvoiceDto) HasRequiredPaymentAction() bool`

HasRequiredPaymentAction returns a boolean if a field has been set.

### GetStripeInvoiceId

`func (o *InvoiceDto) GetStripeInvoiceId() string`

GetStripeInvoiceId returns the StripeInvoiceId field if non-nil, zero value otherwise.

### GetStripeInvoiceIdOk

`func (o *InvoiceDto) GetStripeInvoiceIdOk() (*string, bool)`

GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoiceId

`func (o *InvoiceDto) SetStripeInvoiceId(v string)`

SetStripeInvoiceId sets StripeInvoiceId field to given value.

### HasStripeInvoiceId

`func (o *InvoiceDto) HasStripeInvoiceId() bool`

HasStripeInvoiceId returns a boolean if a field has been set.

### SetStripeInvoiceIdNil

`func (o *InvoiceDto) SetStripeInvoiceIdNil(b bool)`

 SetStripeInvoiceIdNil sets the value for StripeInvoiceId to be an explicit nil

### UnsetStripeInvoiceId
`func (o *InvoiceDto) UnsetStripeInvoiceId()`

UnsetStripeInvoiceId ensures that no value is present for StripeInvoiceId, not even an explicit nil
### GetPrice

`func (o *InvoiceDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStartDate

`func (o *InvoiceDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InvoiceDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


