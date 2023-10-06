# UpdateSubscriptionCommand

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
**TrialDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateSubscriptionCommand

`func NewUpdateSubscriptionCommand() *UpdateSubscriptionCommand`

NewUpdateSubscriptionCommand instantiates a new UpdateSubscriptionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionCommandWithDefaults

`func NewUpdateSubscriptionCommandWithDefaults() *UpdateSubscriptionCommand`

NewUpdateSubscriptionCommandWithDefaults instantiates a new UpdateSubscriptionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateSubscriptionCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSubscriptionCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSubscriptionCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSubscriptionCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateSubscriptionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSubscriptionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSubscriptionCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSubscriptionCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSubscriptionCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSubscriptionCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectLimit

`func (o *UpdateSubscriptionCommand) GetProjectLimit() int32`

GetProjectLimit returns the ProjectLimit field if non-nil, zero value otherwise.

### GetProjectLimitOk

`func (o *UpdateSubscriptionCommand) GetProjectLimitOk() (*int32, bool)`

GetProjectLimitOk returns a tuple with the ProjectLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLimit

`func (o *UpdateSubscriptionCommand) SetProjectLimit(v int32)`

SetProjectLimit sets ProjectLimit field to given value.

### HasProjectLimit

`func (o *UpdateSubscriptionCommand) HasProjectLimit() bool`

HasProjectLimit returns a boolean if a field has been set.

### GetServerLimit

`func (o *UpdateSubscriptionCommand) GetServerLimit() int32`

GetServerLimit returns the ServerLimit field if non-nil, zero value otherwise.

### GetServerLimitOk

`func (o *UpdateSubscriptionCommand) GetServerLimitOk() (*int32, bool)`

GetServerLimitOk returns a tuple with the ServerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLimit

`func (o *UpdateSubscriptionCommand) SetServerLimit(v int32)`

SetServerLimit sets ServerLimit field to given value.

### HasServerLimit

`func (o *UpdateSubscriptionCommand) HasServerLimit() bool`

HasServerLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *UpdateSubscriptionCommand) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UpdateSubscriptionCommand) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UpdateSubscriptionCommand) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UpdateSubscriptionCommand) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCloudCredentialLimit

`func (o *UpdateSubscriptionCommand) GetCloudCredentialLimit() int32`

GetCloudCredentialLimit returns the CloudCredentialLimit field if non-nil, zero value otherwise.

### GetCloudCredentialLimitOk

`func (o *UpdateSubscriptionCommand) GetCloudCredentialLimitOk() (*int32, bool)`

GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialLimit

`func (o *UpdateSubscriptionCommand) SetCloudCredentialLimit(v int32)`

SetCloudCredentialLimit sets CloudCredentialLimit field to given value.

### HasCloudCredentialLimit

`func (o *UpdateSubscriptionCommand) HasCloudCredentialLimit() bool`

HasCloudCredentialLimit returns a boolean if a field has been set.

### GetMonthlyPrice

`func (o *UpdateSubscriptionCommand) GetMonthlyPrice() float64`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *UpdateSubscriptionCommand) GetMonthlyPriceOk() (*float64, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *UpdateSubscriptionCommand) SetMonthlyPrice(v float64)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *UpdateSubscriptionCommand) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetYearlyPrice

`func (o *UpdateSubscriptionCommand) GetYearlyPrice() float64`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *UpdateSubscriptionCommand) GetYearlyPriceOk() (*float64, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *UpdateSubscriptionCommand) SetYearlyPrice(v float64)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *UpdateSubscriptionCommand) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.

### GetTcuPrice

`func (o *UpdateSubscriptionCommand) GetTcuPrice() float64`

GetTcuPrice returns the TcuPrice field if non-nil, zero value otherwise.

### GetTcuPriceOk

`func (o *UpdateSubscriptionCommand) GetTcuPriceOk() (*float64, bool)`

GetTcuPriceOk returns a tuple with the TcuPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcuPrice

`func (o *UpdateSubscriptionCommand) SetTcuPrice(v float64)`

SetTcuPrice sets TcuPrice field to given value.

### HasTcuPrice

`func (o *UpdateSubscriptionCommand) HasTcuPrice() bool`

HasTcuPrice returns a boolean if a field has been set.

### GetTrialDays

`func (o *UpdateSubscriptionCommand) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *UpdateSubscriptionCommand) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *UpdateSubscriptionCommand) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *UpdateSubscriptionCommand) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


