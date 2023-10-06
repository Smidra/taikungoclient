# CostComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**HourlyQuantity** | Pointer to **NullableString** |  | [optional] 
**MonthlyQuantity** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableString** |  | [optional] 
**HourlyCost** | Pointer to **NullableString** |  | [optional] 
**MonthlyCost** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCostComponent

`func NewCostComponent() *CostComponent`

NewCostComponent instantiates a new CostComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostComponentWithDefaults

`func NewCostComponentWithDefaults() *CostComponent`

NewCostComponentWithDefaults instantiates a new CostComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CostComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CostComponent) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CostComponent) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnit

`func (o *CostComponent) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CostComponent) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CostComponent) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CostComponent) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *CostComponent) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *CostComponent) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetHourlyQuantity

`func (o *CostComponent) GetHourlyQuantity() string`

GetHourlyQuantity returns the HourlyQuantity field if non-nil, zero value otherwise.

### GetHourlyQuantityOk

`func (o *CostComponent) GetHourlyQuantityOk() (*string, bool)`

GetHourlyQuantityOk returns a tuple with the HourlyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyQuantity

`func (o *CostComponent) SetHourlyQuantity(v string)`

SetHourlyQuantity sets HourlyQuantity field to given value.

### HasHourlyQuantity

`func (o *CostComponent) HasHourlyQuantity() bool`

HasHourlyQuantity returns a boolean if a field has been set.

### SetHourlyQuantityNil

`func (o *CostComponent) SetHourlyQuantityNil(b bool)`

 SetHourlyQuantityNil sets the value for HourlyQuantity to be an explicit nil

### UnsetHourlyQuantity
`func (o *CostComponent) UnsetHourlyQuantity()`

UnsetHourlyQuantity ensures that no value is present for HourlyQuantity, not even an explicit nil
### GetMonthlyQuantity

`func (o *CostComponent) GetMonthlyQuantity() string`

GetMonthlyQuantity returns the MonthlyQuantity field if non-nil, zero value otherwise.

### GetMonthlyQuantityOk

`func (o *CostComponent) GetMonthlyQuantityOk() (*string, bool)`

GetMonthlyQuantityOk returns a tuple with the MonthlyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyQuantity

`func (o *CostComponent) SetMonthlyQuantity(v string)`

SetMonthlyQuantity sets MonthlyQuantity field to given value.

### HasMonthlyQuantity

`func (o *CostComponent) HasMonthlyQuantity() bool`

HasMonthlyQuantity returns a boolean if a field has been set.

### SetMonthlyQuantityNil

`func (o *CostComponent) SetMonthlyQuantityNil(b bool)`

 SetMonthlyQuantityNil sets the value for MonthlyQuantity to be an explicit nil

### UnsetMonthlyQuantity
`func (o *CostComponent) UnsetMonthlyQuantity()`

UnsetMonthlyQuantity ensures that no value is present for MonthlyQuantity, not even an explicit nil
### GetPrice

`func (o *CostComponent) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CostComponent) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CostComponent) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CostComponent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CostComponent) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CostComponent) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetHourlyCost

`func (o *CostComponent) GetHourlyCost() string`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *CostComponent) GetHourlyCostOk() (*string, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *CostComponent) SetHourlyCost(v string)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *CostComponent) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### SetHourlyCostNil

`func (o *CostComponent) SetHourlyCostNil(b bool)`

 SetHourlyCostNil sets the value for HourlyCost to be an explicit nil

### UnsetHourlyCost
`func (o *CostComponent) UnsetHourlyCost()`

UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
### GetMonthlyCost

`func (o *CostComponent) GetMonthlyCost() string`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *CostComponent) GetMonthlyCostOk() (*string, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *CostComponent) SetMonthlyCost(v string)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *CostComponent) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### SetMonthlyCostNil

`func (o *CostComponent) SetMonthlyCostNil(b bool)`

 SetMonthlyCostNil sets the value for MonthlyCost to be an explicit nil

### UnsetMonthlyCost
`func (o *CostComponent) UnsetMonthlyCost()`

UnsetMonthlyCost ensures that no value is present for MonthlyCost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


