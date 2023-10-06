# CreateSubscriptionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectLimit** | Pointer to **int32** |  | [optional] 
**ServerLimit** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**CloudCredentialLimit** | Pointer to **int32** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**MonthlyPrice** | Pointer to **NullableFloat64** |  | [optional] 
**TcuPrice** | Pointer to **NullableFloat64** |  | [optional] 
**YearlyPrice** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewCreateSubscriptionCommand

`func NewCreateSubscriptionCommand() *CreateSubscriptionCommand`

NewCreateSubscriptionCommand instantiates a new CreateSubscriptionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionCommandWithDefaults

`func NewCreateSubscriptionCommandWithDefaults() *CreateSubscriptionCommand`

NewCreateSubscriptionCommandWithDefaults instantiates a new CreateSubscriptionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSubscriptionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubscriptionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubscriptionCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSubscriptionCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateSubscriptionCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateSubscriptionCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectLimit

`func (o *CreateSubscriptionCommand) GetProjectLimit() int32`

GetProjectLimit returns the ProjectLimit field if non-nil, zero value otherwise.

### GetProjectLimitOk

`func (o *CreateSubscriptionCommand) GetProjectLimitOk() (*int32, bool)`

GetProjectLimitOk returns a tuple with the ProjectLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLimit

`func (o *CreateSubscriptionCommand) SetProjectLimit(v int32)`

SetProjectLimit sets ProjectLimit field to given value.

### HasProjectLimit

`func (o *CreateSubscriptionCommand) HasProjectLimit() bool`

HasProjectLimit returns a boolean if a field has been set.

### GetServerLimit

`func (o *CreateSubscriptionCommand) GetServerLimit() int32`

GetServerLimit returns the ServerLimit field if non-nil, zero value otherwise.

### GetServerLimitOk

`func (o *CreateSubscriptionCommand) GetServerLimitOk() (*int32, bool)`

GetServerLimitOk returns a tuple with the ServerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLimit

`func (o *CreateSubscriptionCommand) SetServerLimit(v int32)`

SetServerLimit sets ServerLimit field to given value.

### HasServerLimit

`func (o *CreateSubscriptionCommand) HasServerLimit() bool`

HasServerLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *CreateSubscriptionCommand) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *CreateSubscriptionCommand) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *CreateSubscriptionCommand) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *CreateSubscriptionCommand) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCloudCredentialLimit

`func (o *CreateSubscriptionCommand) GetCloudCredentialLimit() int32`

GetCloudCredentialLimit returns the CloudCredentialLimit field if non-nil, zero value otherwise.

### GetCloudCredentialLimitOk

`func (o *CreateSubscriptionCommand) GetCloudCredentialLimitOk() (*int32, bool)`

GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialLimit

`func (o *CreateSubscriptionCommand) SetCloudCredentialLimit(v int32)`

SetCloudCredentialLimit sets CloudCredentialLimit field to given value.

### HasCloudCredentialLimit

`func (o *CreateSubscriptionCommand) HasCloudCredentialLimit() bool`

HasCloudCredentialLimit returns a boolean if a field has been set.

### GetTrialDays

`func (o *CreateSubscriptionCommand) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *CreateSubscriptionCommand) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *CreateSubscriptionCommand) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *CreateSubscriptionCommand) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetMonthlyPrice

`func (o *CreateSubscriptionCommand) GetMonthlyPrice() float64`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *CreateSubscriptionCommand) GetMonthlyPriceOk() (*float64, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *CreateSubscriptionCommand) SetMonthlyPrice(v float64)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *CreateSubscriptionCommand) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### SetMonthlyPriceNil

`func (o *CreateSubscriptionCommand) SetMonthlyPriceNil(b bool)`

 SetMonthlyPriceNil sets the value for MonthlyPrice to be an explicit nil

### UnsetMonthlyPrice
`func (o *CreateSubscriptionCommand) UnsetMonthlyPrice()`

UnsetMonthlyPrice ensures that no value is present for MonthlyPrice, not even an explicit nil
### GetTcuPrice

`func (o *CreateSubscriptionCommand) GetTcuPrice() float64`

GetTcuPrice returns the TcuPrice field if non-nil, zero value otherwise.

### GetTcuPriceOk

`func (o *CreateSubscriptionCommand) GetTcuPriceOk() (*float64, bool)`

GetTcuPriceOk returns a tuple with the TcuPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcuPrice

`func (o *CreateSubscriptionCommand) SetTcuPrice(v float64)`

SetTcuPrice sets TcuPrice field to given value.

### HasTcuPrice

`func (o *CreateSubscriptionCommand) HasTcuPrice() bool`

HasTcuPrice returns a boolean if a field has been set.

### SetTcuPriceNil

`func (o *CreateSubscriptionCommand) SetTcuPriceNil(b bool)`

 SetTcuPriceNil sets the value for TcuPrice to be an explicit nil

### UnsetTcuPrice
`func (o *CreateSubscriptionCommand) UnsetTcuPrice()`

UnsetTcuPrice ensures that no value is present for TcuPrice, not even an explicit nil
### GetYearlyPrice

`func (o *CreateSubscriptionCommand) GetYearlyPrice() float64`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *CreateSubscriptionCommand) GetYearlyPriceOk() (*float64, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *CreateSubscriptionCommand) SetYearlyPrice(v float64)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *CreateSubscriptionCommand) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.

### SetYearlyPriceNil

`func (o *CreateSubscriptionCommand) SetYearlyPriceNil(b bool)`

 SetYearlyPriceNil sets the value for YearlyPrice to be an explicit nil

### UnsetYearlyPrice
`func (o *CreateSubscriptionCommand) UnsetYearlyPrice()`

UnsetYearlyPrice ensures that no value is present for YearlyPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


