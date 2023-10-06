# PrometheusBillingCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int32** |  | [optional] 
**PrometheusRuleId** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 

## Methods

### NewPrometheusBillingCreateCommand

`func NewPrometheusBillingCreateCommand() *PrometheusBillingCreateCommand`

NewPrometheusBillingCreateCommand instantiates a new PrometheusBillingCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusBillingCreateCommandWithDefaults

`func NewPrometheusBillingCreateCommandWithDefaults() *PrometheusBillingCreateCommand`

NewPrometheusBillingCreateCommandWithDefaults instantiates a new PrometheusBillingCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *PrometheusBillingCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PrometheusBillingCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PrometheusBillingCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PrometheusBillingCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPrometheusRuleId

`func (o *PrometheusBillingCreateCommand) GetPrometheusRuleId() int32`

GetPrometheusRuleId returns the PrometheusRuleId field if non-nil, zero value otherwise.

### GetPrometheusRuleIdOk

`func (o *PrometheusBillingCreateCommand) GetPrometheusRuleIdOk() (*int32, bool)`

GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusRuleId

`func (o *PrometheusBillingCreateCommand) SetPrometheusRuleId(v int32)`

SetPrometheusRuleId sets PrometheusRuleId field to given value.

### HasPrometheusRuleId

`func (o *PrometheusBillingCreateCommand) HasPrometheusRuleId() bool`

HasPrometheusRuleId returns a boolean if a field has been set.

### GetStartDate

`func (o *PrometheusBillingCreateCommand) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PrometheusBillingCreateCommand) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PrometheusBillingCreateCommand) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PrometheusBillingCreateCommand) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetPrice

`func (o *PrometheusBillingCreateCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PrometheusBillingCreateCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PrometheusBillingCreateCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PrometheusBillingCreateCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


