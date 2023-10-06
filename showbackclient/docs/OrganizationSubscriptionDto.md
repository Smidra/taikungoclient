# OrganizationSubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **int32** |  | [optional] 
**StripeSubscriptionId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionType** | Pointer to **NullableString** |  | [optional] 
**SubscriptionName** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Invoices** | Pointer to [**[]InvoiceDto**](InvoiceDto.md) |  | [optional] 

## Methods

### NewOrganizationSubscriptionDto

`func NewOrganizationSubscriptionDto() *OrganizationSubscriptionDto`

NewOrganizationSubscriptionDto instantiates a new OrganizationSubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionDtoWithDefaults

`func NewOrganizationSubscriptionDtoWithDefaults() *OrganizationSubscriptionDto`

NewOrganizationSubscriptionDtoWithDefaults instantiates a new OrganizationSubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationSubscriptionDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSubscriptionDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSubscriptionDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationSubscriptionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationSubscriptionDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationSubscriptionDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationSubscriptionDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationSubscriptionDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OrganizationSubscriptionDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationSubscriptionDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationSubscriptionDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OrganizationSubscriptionDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OrganizationSubscriptionDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OrganizationSubscriptionDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetSubscriptionId

`func (o *OrganizationSubscriptionDto) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *OrganizationSubscriptionDto) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *OrganizationSubscriptionDto) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *OrganizationSubscriptionDto) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *OrganizationSubscriptionDto) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *OrganizationSubscriptionDto) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *OrganizationSubscriptionDto) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *OrganizationSubscriptionDto) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### SetStripeSubscriptionIdNil

`func (o *OrganizationSubscriptionDto) SetStripeSubscriptionIdNil(b bool)`

 SetStripeSubscriptionIdNil sets the value for StripeSubscriptionId to be an explicit nil

### UnsetStripeSubscriptionId
`func (o *OrganizationSubscriptionDto) UnsetStripeSubscriptionId()`

UnsetStripeSubscriptionId ensures that no value is present for StripeSubscriptionId, not even an explicit nil
### GetSubscriptionType

`func (o *OrganizationSubscriptionDto) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *OrganizationSubscriptionDto) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *OrganizationSubscriptionDto) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *OrganizationSubscriptionDto) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *OrganizationSubscriptionDto) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *OrganizationSubscriptionDto) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetSubscriptionName

`func (o *OrganizationSubscriptionDto) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *OrganizationSubscriptionDto) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *OrganizationSubscriptionDto) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *OrganizationSubscriptionDto) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### SetSubscriptionNameNil

`func (o *OrganizationSubscriptionDto) SetSubscriptionNameNil(b bool)`

 SetSubscriptionNameNil sets the value for SubscriptionName to be an explicit nil

### UnsetSubscriptionName
`func (o *OrganizationSubscriptionDto) UnsetSubscriptionName()`

UnsetSubscriptionName ensures that no value is present for SubscriptionName, not even an explicit nil
### GetStartDate

`func (o *OrganizationSubscriptionDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OrganizationSubscriptionDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OrganizationSubscriptionDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OrganizationSubscriptionDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *OrganizationSubscriptionDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OrganizationSubscriptionDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OrganizationSubscriptionDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OrganizationSubscriptionDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *OrganizationSubscriptionDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *OrganizationSubscriptionDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetInvoices

`func (o *OrganizationSubscriptionDto) GetInvoices() []InvoiceDto`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *OrganizationSubscriptionDto) GetInvoicesOk() (*[]InvoiceDto, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *OrganizationSubscriptionDto) SetInvoices(v []InvoiceDto)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *OrganizationSubscriptionDto) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### SetInvoicesNil

`func (o *OrganizationSubscriptionDto) SetInvoicesNil(b bool)`

 SetInvoicesNil sets the value for Invoices to be an explicit nil

### UnsetInvoices
`func (o *OrganizationSubscriptionDto) UnsetInvoices()`

UnsetInvoices ensures that no value is present for Invoices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


