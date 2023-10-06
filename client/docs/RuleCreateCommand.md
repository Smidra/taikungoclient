# RuleCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**MetricName** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to [**[]PrometheusLabelListDto**](PrometheusLabelListDto.md) |  | [optional] 
**Type** | Pointer to [**PrometheusType**](PrometheusType.md) |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**PartnerId** | Pointer to **NullableInt32** |  | [optional] 
**OperationCredentialId** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **[]int32** |  | [optional] 
**RuleDiscountRate** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRuleCreateCommand

`func NewRuleCreateCommand() *RuleCreateCommand`

NewRuleCreateCommand instantiates a new RuleCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCreateCommandWithDefaults

`func NewRuleCreateCommandWithDefaults() *RuleCreateCommand`

NewRuleCreateCommandWithDefaults instantiates a new RuleCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCreateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleCreateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RuleCreateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RuleCreateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetricName

`func (o *RuleCreateCommand) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RuleCreateCommand) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RuleCreateCommand) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RuleCreateCommand) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *RuleCreateCommand) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *RuleCreateCommand) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetLabels

`func (o *RuleCreateCommand) GetLabels() []PrometheusLabelListDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RuleCreateCommand) GetLabelsOk() (*[]PrometheusLabelListDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RuleCreateCommand) SetLabels(v []PrometheusLabelListDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RuleCreateCommand) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *RuleCreateCommand) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RuleCreateCommand) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *RuleCreateCommand) GetType() PrometheusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleCreateCommand) GetTypeOk() (*PrometheusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleCreateCommand) SetType(v PrometheusType)`

SetType sets Type field to given value.

### HasType

`func (o *RuleCreateCommand) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *RuleCreateCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RuleCreateCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RuleCreateCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RuleCreateCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPartnerId

`func (o *RuleCreateCommand) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *RuleCreateCommand) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *RuleCreateCommand) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *RuleCreateCommand) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *RuleCreateCommand) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *RuleCreateCommand) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetOperationCredentialId

`func (o *RuleCreateCommand) GetOperationCredentialId() int32`

GetOperationCredentialId returns the OperationCredentialId field if non-nil, zero value otherwise.

### GetOperationCredentialIdOk

`func (o *RuleCreateCommand) GetOperationCredentialIdOk() (*int32, bool)`

GetOperationCredentialIdOk returns a tuple with the OperationCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCredentialId

`func (o *RuleCreateCommand) SetOperationCredentialId(v int32)`

SetOperationCredentialId sets OperationCredentialId field to given value.

### HasOperationCredentialId

`func (o *RuleCreateCommand) HasOperationCredentialId() bool`

HasOperationCredentialId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *RuleCreateCommand) GetOrganizationId() []int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RuleCreateCommand) GetOrganizationIdOk() (*[]int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RuleCreateCommand) SetOrganizationId(v []int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RuleCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *RuleCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *RuleCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetRuleDiscountRate

`func (o *RuleCreateCommand) GetRuleDiscountRate() int32`

GetRuleDiscountRate returns the RuleDiscountRate field if non-nil, zero value otherwise.

### GetRuleDiscountRateOk

`func (o *RuleCreateCommand) GetRuleDiscountRateOk() (*int32, bool)`

GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDiscountRate

`func (o *RuleCreateCommand) SetRuleDiscountRate(v int32)`

SetRuleDiscountRate sets RuleDiscountRate field to given value.

### HasRuleDiscountRate

`func (o *RuleCreateCommand) HasRuleDiscountRate() bool`

HasRuleDiscountRate returns a boolean if a field has been set.

### SetRuleDiscountRateNil

`func (o *RuleCreateCommand) SetRuleDiscountRateNil(b bool)`

 SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil

### UnsetRuleDiscountRate
`func (o *RuleCreateCommand) UnsetRuleDiscountRate()`

UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


