# Chart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**ChartSpec**](ChartSpec.md) |  | [optional] 

## Methods

### NewChart

`func NewChart() *Chart`

NewChart instantiates a new Chart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartWithDefaults

`func NewChartWithDefaults() *Chart`

NewChartWithDefaults instantiates a new Chart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *Chart) GetSpec() ChartSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Chart) GetSpecOk() (*ChartSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Chart) SetSpec(v ChartSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Chart) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


