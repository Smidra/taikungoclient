# HelmStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**Failures** | Pointer to **int64** |  | [optional] 
**HelmChart** | Pointer to **NullableString** |  | [optional] 
**ObservedGeneration** | Pointer to **int64** |  | [optional] 

## Methods

### NewHelmStatus

`func NewHelmStatus() *HelmStatus`

NewHelmStatus instantiates a new HelmStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmStatusWithDefaults

`func NewHelmStatusWithDefaults() *HelmStatus`

NewHelmStatusWithDefaults instantiates a new HelmStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *HelmStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HelmStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HelmStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HelmStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *HelmStatus) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *HelmStatus) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetFailures

`func (o *HelmStatus) GetFailures() int64`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *HelmStatus) GetFailuresOk() (*int64, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *HelmStatus) SetFailures(v int64)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *HelmStatus) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetHelmChart

`func (o *HelmStatus) GetHelmChart() string`

GetHelmChart returns the HelmChart field if non-nil, zero value otherwise.

### GetHelmChartOk

`func (o *HelmStatus) GetHelmChartOk() (*string, bool)`

GetHelmChartOk returns a tuple with the HelmChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChart

`func (o *HelmStatus) SetHelmChart(v string)`

SetHelmChart sets HelmChart field to given value.

### HasHelmChart

`func (o *HelmStatus) HasHelmChart() bool`

HasHelmChart returns a boolean if a field has been set.

### SetHelmChartNil

`func (o *HelmStatus) SetHelmChartNil(b bool)`

 SetHelmChartNil sets the value for HelmChart to be an explicit nil

### UnsetHelmChart
`func (o *HelmStatus) UnsetHelmChart()`

UnsetHelmChart ensures that no value is present for HelmChart, not even an explicit nil
### GetObservedGeneration

`func (o *HelmStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *HelmStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *HelmStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *HelmStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


