# InvoiceListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**RequiredPaymentAction** | Pointer to **bool** |  | [optional] 
**IsPaid** | Pointer to **bool** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionType** | Pointer to **NullableString** |  | [optional] 
**SubscriptionName** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**InvoiceNumber** | Pointer to **NullableString** |  | [optional] 
**OrganizationSubscriptionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewInvoiceListDto

`func NewInvoiceListDto() *InvoiceListDto`

NewInvoiceListDto instantiates a new InvoiceListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListDtoWithDefaults

`func NewInvoiceListDtoWithDefaults() *InvoiceListDto`

NewInvoiceListDtoWithDefaults instantiates a new InvoiceListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InvoiceListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InvoiceListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InvoiceListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDate

`func (o *InvoiceListDto) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceListDto) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceListDto) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceListDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *InvoiceListDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *InvoiceListDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *InvoiceListDto) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceListDto) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceListDto) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceListDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *InvoiceListDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *InvoiceListDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetRequiredPaymentAction

`func (o *InvoiceListDto) GetRequiredPaymentAction() bool`

GetRequiredPaymentAction returns the RequiredPaymentAction field if non-nil, zero value otherwise.

### GetRequiredPaymentActionOk

`func (o *InvoiceListDto) GetRequiredPaymentActionOk() (*bool, bool)`

GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPaymentAction

`func (o *InvoiceListDto) SetRequiredPaymentAction(v bool)`

SetRequiredPaymentAction sets RequiredPaymentAction field to given value.

### HasRequiredPaymentAction

`func (o *InvoiceListDto) HasRequiredPaymentAction() bool`

HasRequiredPaymentAction returns a boolean if a field has been set.

### GetIsPaid

`func (o *InvoiceListDto) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *InvoiceListDto) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *InvoiceListDto) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *InvoiceListDto) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetInvoiceId

`func (o *InvoiceListDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceListDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceListDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceListDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *InvoiceListDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceListDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionType

`func (o *InvoiceListDto) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *InvoiceListDto) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *InvoiceListDto) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *InvoiceListDto) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *InvoiceListDto) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *InvoiceListDto) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetSubscriptionName

`func (o *InvoiceListDto) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *InvoiceListDto) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *InvoiceListDto) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *InvoiceListDto) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### SetSubscriptionNameNil

`func (o *InvoiceListDto) SetSubscriptionNameNil(b bool)`

 SetSubscriptionNameNil sets the value for SubscriptionName to be an explicit nil

### UnsetSubscriptionName
`func (o *InvoiceListDto) UnsetSubscriptionName()`

UnsetSubscriptionName ensures that no value is present for SubscriptionName, not even an explicit nil
### GetPrice

`func (o *InvoiceListDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceListDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceListDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceListDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InvoiceListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InvoiceListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InvoiceListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InvoiceListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *InvoiceListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InvoiceListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InvoiceListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *InvoiceListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *InvoiceListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *InvoiceListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetInvoiceNumber

`func (o *InvoiceListDto) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceListDto) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceListDto) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoiceListDto) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### SetInvoiceNumberNil

`func (o *InvoiceListDto) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *InvoiceListDto) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetOrganizationSubscriptionId

`func (o *InvoiceListDto) GetOrganizationSubscriptionId() int32`

GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field if non-nil, zero value otherwise.

### GetOrganizationSubscriptionIdOk

`func (o *InvoiceListDto) GetOrganizationSubscriptionIdOk() (*int32, bool)`

GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSubscriptionId

`func (o *InvoiceListDto) SetOrganizationSubscriptionId(v int32)`

SetOrganizationSubscriptionId sets OrganizationSubscriptionId field to given value.

### HasOrganizationSubscriptionId

`func (o *InvoiceListDto) HasOrganizationSubscriptionId() bool`

HasOrganizationSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


