# StripeInvoiceListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**InvoiceStatus** | Pointer to **NullableString** |  | [optional] 
**ChargeStatus** | Pointer to **NullableString** |  | [optional] 
**ChargeReason** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStripeInvoiceListDto

`func NewStripeInvoiceListDto() *StripeInvoiceListDto`

NewStripeInvoiceListDto instantiates a new StripeInvoiceListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeInvoiceListDtoWithDefaults

`func NewStripeInvoiceListDtoWithDefaults() *StripeInvoiceListDto`

NewStripeInvoiceListDtoWithDefaults instantiates a new StripeInvoiceListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeInvoiceListDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeInvoiceListDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeInvoiceListDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeInvoiceListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StripeInvoiceListDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StripeInvoiceListDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInvoiceStatus

`func (o *StripeInvoiceListDto) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *StripeInvoiceListDto) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *StripeInvoiceListDto) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *StripeInvoiceListDto) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### SetInvoiceStatusNil

`func (o *StripeInvoiceListDto) SetInvoiceStatusNil(b bool)`

 SetInvoiceStatusNil sets the value for InvoiceStatus to be an explicit nil

### UnsetInvoiceStatus
`func (o *StripeInvoiceListDto) UnsetInvoiceStatus()`

UnsetInvoiceStatus ensures that no value is present for InvoiceStatus, not even an explicit nil
### GetChargeStatus

`func (o *StripeInvoiceListDto) GetChargeStatus() string`

GetChargeStatus returns the ChargeStatus field if non-nil, zero value otherwise.

### GetChargeStatusOk

`func (o *StripeInvoiceListDto) GetChargeStatusOk() (*string, bool)`

GetChargeStatusOk returns a tuple with the ChargeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeStatus

`func (o *StripeInvoiceListDto) SetChargeStatus(v string)`

SetChargeStatus sets ChargeStatus field to given value.

### HasChargeStatus

`func (o *StripeInvoiceListDto) HasChargeStatus() bool`

HasChargeStatus returns a boolean if a field has been set.

### SetChargeStatusNil

`func (o *StripeInvoiceListDto) SetChargeStatusNil(b bool)`

 SetChargeStatusNil sets the value for ChargeStatus to be an explicit nil

### UnsetChargeStatus
`func (o *StripeInvoiceListDto) UnsetChargeStatus()`

UnsetChargeStatus ensures that no value is present for ChargeStatus, not even an explicit nil
### GetChargeReason

`func (o *StripeInvoiceListDto) GetChargeReason() string`

GetChargeReason returns the ChargeReason field if non-nil, zero value otherwise.

### GetChargeReasonOk

`func (o *StripeInvoiceListDto) GetChargeReasonOk() (*string, bool)`

GetChargeReasonOk returns a tuple with the ChargeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeReason

`func (o *StripeInvoiceListDto) SetChargeReason(v string)`

SetChargeReason sets ChargeReason field to given value.

### HasChargeReason

`func (o *StripeInvoiceListDto) HasChargeReason() bool`

HasChargeReason returns a boolean if a field has been set.

### SetChargeReasonNil

`func (o *StripeInvoiceListDto) SetChargeReasonNil(b bool)`

 SetChargeReasonNil sets the value for ChargeReason to be an explicit nil

### UnsetChargeReason
`func (o *StripeInvoiceListDto) UnsetChargeReason()`

UnsetChargeReason ensures that no value is present for ChargeReason, not even an explicit nil
### GetPrice

`func (o *StripeInvoiceListDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StripeInvoiceListDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *StripeInvoiceListDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *StripeInvoiceListDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStartDate

`func (o *StripeInvoiceListDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StripeInvoiceListDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StripeInvoiceListDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StripeInvoiceListDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *StripeInvoiceListDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StripeInvoiceListDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StripeInvoiceListDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StripeInvoiceListDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


