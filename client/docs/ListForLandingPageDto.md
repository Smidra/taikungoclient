# ListForLandingPageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectLimit** | Pointer to **int32** |  | [optional] 
**ServerLimit** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**CloudCredentialLimit** | Pointer to **int32** |  | [optional] 
**MonthlyPrice** | Pointer to **float64** |  | [optional] 
**YearlyPrice** | Pointer to **float64** |  | [optional] 
**TcuPrice** | Pointer to **float64** |  | [optional] 
**IsDeprecated** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**PartnerId** | Pointer to **NullableInt32** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsFree** | Pointer to **bool** |  | [optional] 
**IsEnterprise** | Pointer to **bool** |  | [optional] 

## Methods

### NewListForLandingPageDto

`func NewListForLandingPageDto() *ListForLandingPageDto`

NewListForLandingPageDto instantiates a new ListForLandingPageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListForLandingPageDtoWithDefaults

`func NewListForLandingPageDtoWithDefaults() *ListForLandingPageDto`

NewListForLandingPageDtoWithDefaults instantiates a new ListForLandingPageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListForLandingPageDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListForLandingPageDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListForLandingPageDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListForLandingPageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListForLandingPageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListForLandingPageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListForLandingPageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListForLandingPageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListForLandingPageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListForLandingPageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectLimit

`func (o *ListForLandingPageDto) GetProjectLimit() int32`

GetProjectLimit returns the ProjectLimit field if non-nil, zero value otherwise.

### GetProjectLimitOk

`func (o *ListForLandingPageDto) GetProjectLimitOk() (*int32, bool)`

GetProjectLimitOk returns a tuple with the ProjectLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLimit

`func (o *ListForLandingPageDto) SetProjectLimit(v int32)`

SetProjectLimit sets ProjectLimit field to given value.

### HasProjectLimit

`func (o *ListForLandingPageDto) HasProjectLimit() bool`

HasProjectLimit returns a boolean if a field has been set.

### GetServerLimit

`func (o *ListForLandingPageDto) GetServerLimit() int32`

GetServerLimit returns the ServerLimit field if non-nil, zero value otherwise.

### GetServerLimitOk

`func (o *ListForLandingPageDto) GetServerLimitOk() (*int32, bool)`

GetServerLimitOk returns a tuple with the ServerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLimit

`func (o *ListForLandingPageDto) SetServerLimit(v int32)`

SetServerLimit sets ServerLimit field to given value.

### HasServerLimit

`func (o *ListForLandingPageDto) HasServerLimit() bool`

HasServerLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *ListForLandingPageDto) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *ListForLandingPageDto) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *ListForLandingPageDto) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *ListForLandingPageDto) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCloudCredentialLimit

`func (o *ListForLandingPageDto) GetCloudCredentialLimit() int32`

GetCloudCredentialLimit returns the CloudCredentialLimit field if non-nil, zero value otherwise.

### GetCloudCredentialLimitOk

`func (o *ListForLandingPageDto) GetCloudCredentialLimitOk() (*int32, bool)`

GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialLimit

`func (o *ListForLandingPageDto) SetCloudCredentialLimit(v int32)`

SetCloudCredentialLimit sets CloudCredentialLimit field to given value.

### HasCloudCredentialLimit

`func (o *ListForLandingPageDto) HasCloudCredentialLimit() bool`

HasCloudCredentialLimit returns a boolean if a field has been set.

### GetMonthlyPrice

`func (o *ListForLandingPageDto) GetMonthlyPrice() float64`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *ListForLandingPageDto) GetMonthlyPriceOk() (*float64, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *ListForLandingPageDto) SetMonthlyPrice(v float64)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *ListForLandingPageDto) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetYearlyPrice

`func (o *ListForLandingPageDto) GetYearlyPrice() float64`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *ListForLandingPageDto) GetYearlyPriceOk() (*float64, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *ListForLandingPageDto) SetYearlyPrice(v float64)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *ListForLandingPageDto) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.

### GetTcuPrice

`func (o *ListForLandingPageDto) GetTcuPrice() float64`

GetTcuPrice returns the TcuPrice field if non-nil, zero value otherwise.

### GetTcuPriceOk

`func (o *ListForLandingPageDto) GetTcuPriceOk() (*float64, bool)`

GetTcuPriceOk returns a tuple with the TcuPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcuPrice

`func (o *ListForLandingPageDto) SetTcuPrice(v float64)`

SetTcuPrice sets TcuPrice field to given value.

### HasTcuPrice

`func (o *ListForLandingPageDto) HasTcuPrice() bool`

HasTcuPrice returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *ListForLandingPageDto) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ListForLandingPageDto) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ListForLandingPageDto) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *ListForLandingPageDto) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetCurrency

`func (o *ListForLandingPageDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListForLandingPageDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListForLandingPageDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListForLandingPageDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ListForLandingPageDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ListForLandingPageDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPartnerId

`func (o *ListForLandingPageDto) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ListForLandingPageDto) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ListForLandingPageDto) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ListForLandingPageDto) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *ListForLandingPageDto) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *ListForLandingPageDto) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetTrialDays

`func (o *ListForLandingPageDto) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *ListForLandingPageDto) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *ListForLandingPageDto) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *ListForLandingPageDto) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetDescription

`func (o *ListForLandingPageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListForLandingPageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListForLandingPageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListForLandingPageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListForLandingPageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListForLandingPageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFree

`func (o *ListForLandingPageDto) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *ListForLandingPageDto) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *ListForLandingPageDto) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *ListForLandingPageDto) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetIsEnterprise

`func (o *ListForLandingPageDto) GetIsEnterprise() bool`

GetIsEnterprise returns the IsEnterprise field if non-nil, zero value otherwise.

### GetIsEnterpriseOk

`func (o *ListForLandingPageDto) GetIsEnterpriseOk() (*bool, bool)`

GetIsEnterpriseOk returns a tuple with the IsEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterprise

`func (o *ListForLandingPageDto) SetIsEnterprise(v bool)`

SetIsEnterprise sets IsEnterprise field to given value.

### HasIsEnterprise

`func (o *ListForLandingPageDto) HasIsEnterprise() bool`

HasIsEnterprise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


