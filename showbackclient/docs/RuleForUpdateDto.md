# RuleForUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**MetricName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**PrometheusType**](PrometheusType.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**LabelsToAdd** | Pointer to [**[]PrometheusLabelListDto**](PrometheusLabelListDto.md) |  | [optional] 
**LabelsToDelete** | Pointer to [**[]PrometheusLabelDeleteDto**](PrometheusLabelDeleteDto.md) |  | [optional] 
**LabelsToUpdate** | Pointer to [**[]PrometheusLabelUpdateDto**](PrometheusLabelUpdateDto.md) |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**RuleDiscountRate** | Pointer to **NullableInt32** |  | [optional] 
**OperationCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRuleForUpdateDto

`func NewRuleForUpdateDto() *RuleForUpdateDto`

NewRuleForUpdateDto instantiates a new RuleForUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleForUpdateDtoWithDefaults

`func NewRuleForUpdateDtoWithDefaults() *RuleForUpdateDto`

NewRuleForUpdateDtoWithDefaults instantiates a new RuleForUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleForUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleForUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleForUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleForUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RuleForUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RuleForUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetricName

`func (o *RuleForUpdateDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RuleForUpdateDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RuleForUpdateDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RuleForUpdateDto) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *RuleForUpdateDto) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *RuleForUpdateDto) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetType

`func (o *RuleForUpdateDto) GetType() PrometheusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleForUpdateDto) GetTypeOk() (*PrometheusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleForUpdateDto) SetType(v PrometheusType)`

SetType sets Type field to given value.

### HasType

`func (o *RuleForUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *RuleForUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RuleForUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RuleForUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RuleForUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *RuleForUpdateDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *RuleForUpdateDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetLabelsToAdd

`func (o *RuleForUpdateDto) GetLabelsToAdd() []PrometheusLabelListDto`

GetLabelsToAdd returns the LabelsToAdd field if non-nil, zero value otherwise.

### GetLabelsToAddOk

`func (o *RuleForUpdateDto) GetLabelsToAddOk() (*[]PrometheusLabelListDto, bool)`

GetLabelsToAddOk returns a tuple with the LabelsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsToAdd

`func (o *RuleForUpdateDto) SetLabelsToAdd(v []PrometheusLabelListDto)`

SetLabelsToAdd sets LabelsToAdd field to given value.

### HasLabelsToAdd

`func (o *RuleForUpdateDto) HasLabelsToAdd() bool`

HasLabelsToAdd returns a boolean if a field has been set.

### SetLabelsToAddNil

`func (o *RuleForUpdateDto) SetLabelsToAddNil(b bool)`

 SetLabelsToAddNil sets the value for LabelsToAdd to be an explicit nil

### UnsetLabelsToAdd
`func (o *RuleForUpdateDto) UnsetLabelsToAdd()`

UnsetLabelsToAdd ensures that no value is present for LabelsToAdd, not even an explicit nil
### GetLabelsToDelete

`func (o *RuleForUpdateDto) GetLabelsToDelete() []PrometheusLabelDeleteDto`

GetLabelsToDelete returns the LabelsToDelete field if non-nil, zero value otherwise.

### GetLabelsToDeleteOk

`func (o *RuleForUpdateDto) GetLabelsToDeleteOk() (*[]PrometheusLabelDeleteDto, bool)`

GetLabelsToDeleteOk returns a tuple with the LabelsToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsToDelete

`func (o *RuleForUpdateDto) SetLabelsToDelete(v []PrometheusLabelDeleteDto)`

SetLabelsToDelete sets LabelsToDelete field to given value.

### HasLabelsToDelete

`func (o *RuleForUpdateDto) HasLabelsToDelete() bool`

HasLabelsToDelete returns a boolean if a field has been set.

### SetLabelsToDeleteNil

`func (o *RuleForUpdateDto) SetLabelsToDeleteNil(b bool)`

 SetLabelsToDeleteNil sets the value for LabelsToDelete to be an explicit nil

### UnsetLabelsToDelete
`func (o *RuleForUpdateDto) UnsetLabelsToDelete()`

UnsetLabelsToDelete ensures that no value is present for LabelsToDelete, not even an explicit nil
### GetLabelsToUpdate

`func (o *RuleForUpdateDto) GetLabelsToUpdate() []PrometheusLabelUpdateDto`

GetLabelsToUpdate returns the LabelsToUpdate field if non-nil, zero value otherwise.

### GetLabelsToUpdateOk

`func (o *RuleForUpdateDto) GetLabelsToUpdateOk() (*[]PrometheusLabelUpdateDto, bool)`

GetLabelsToUpdateOk returns a tuple with the LabelsToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsToUpdate

`func (o *RuleForUpdateDto) SetLabelsToUpdate(v []PrometheusLabelUpdateDto)`

SetLabelsToUpdate sets LabelsToUpdate field to given value.

### HasLabelsToUpdate

`func (o *RuleForUpdateDto) HasLabelsToUpdate() bool`

HasLabelsToUpdate returns a boolean if a field has been set.

### SetLabelsToUpdateNil

`func (o *RuleForUpdateDto) SetLabelsToUpdateNil(b bool)`

 SetLabelsToUpdateNil sets the value for LabelsToUpdate to be an explicit nil

### UnsetLabelsToUpdate
`func (o *RuleForUpdateDto) UnsetLabelsToUpdate()`

UnsetLabelsToUpdate ensures that no value is present for LabelsToUpdate, not even an explicit nil
### GetOrganizationId

`func (o *RuleForUpdateDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RuleForUpdateDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RuleForUpdateDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RuleForUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *RuleForUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *RuleForUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetRuleDiscountRate

`func (o *RuleForUpdateDto) GetRuleDiscountRate() int32`

GetRuleDiscountRate returns the RuleDiscountRate field if non-nil, zero value otherwise.

### GetRuleDiscountRateOk

`func (o *RuleForUpdateDto) GetRuleDiscountRateOk() (*int32, bool)`

GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDiscountRate

`func (o *RuleForUpdateDto) SetRuleDiscountRate(v int32)`

SetRuleDiscountRate sets RuleDiscountRate field to given value.

### HasRuleDiscountRate

`func (o *RuleForUpdateDto) HasRuleDiscountRate() bool`

HasRuleDiscountRate returns a boolean if a field has been set.

### SetRuleDiscountRateNil

`func (o *RuleForUpdateDto) SetRuleDiscountRateNil(b bool)`

 SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil

### UnsetRuleDiscountRate
`func (o *RuleForUpdateDto) UnsetRuleDiscountRate()`

UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil
### GetOperationCredentialId

`func (o *RuleForUpdateDto) GetOperationCredentialId() int32`

GetOperationCredentialId returns the OperationCredentialId field if non-nil, zero value otherwise.

### GetOperationCredentialIdOk

`func (o *RuleForUpdateDto) GetOperationCredentialIdOk() (*int32, bool)`

GetOperationCredentialIdOk returns a tuple with the OperationCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCredentialId

`func (o *RuleForUpdateDto) SetOperationCredentialId(v int32)`

SetOperationCredentialId sets OperationCredentialId field to given value.

### HasOperationCredentialId

`func (o *RuleForUpdateDto) HasOperationCredentialId() bool`

HasOperationCredentialId returns a boolean if a field has been set.

### SetOperationCredentialIdNil

`func (o *RuleForUpdateDto) SetOperationCredentialIdNil(b bool)`

 SetOperationCredentialIdNil sets the value for OperationCredentialId to be an explicit nil

### UnsetOperationCredentialId
`func (o *RuleForUpdateDto) UnsetOperationCredentialId()`

UnsetOperationCredentialId ensures that no value is present for OperationCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


