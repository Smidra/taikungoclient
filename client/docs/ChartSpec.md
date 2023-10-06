# ChartSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**SourceRef** | Pointer to [**SourceRef**](SourceRef.md) |  | [optional] 
**Interval** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChartSpec

`func NewChartSpec() *ChartSpec`

NewChartSpec instantiates a new ChartSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartSpecWithDefaults

`func NewChartSpecWithDefaults() *ChartSpec`

NewChartSpecWithDefaults instantiates a new ChartSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *ChartSpec) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *ChartSpec) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *ChartSpec) SetChart(v string)`

SetChart sets Chart field to given value.

### HasChart

`func (o *ChartSpec) HasChart() bool`

HasChart returns a boolean if a field has been set.

### SetChartNil

`func (o *ChartSpec) SetChartNil(b bool)`

 SetChartNil sets the value for Chart to be an explicit nil

### UnsetChart
`func (o *ChartSpec) UnsetChart()`

UnsetChart ensures that no value is present for Chart, not even an explicit nil
### GetVersion

`func (o *ChartSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChartSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChartSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ChartSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ChartSpec) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ChartSpec) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSourceRef

`func (o *ChartSpec) GetSourceRef() SourceRef`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ChartSpec) GetSourceRefOk() (*SourceRef, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ChartSpec) SetSourceRef(v SourceRef)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *ChartSpec) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetInterval

`func (o *ChartSpec) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ChartSpec) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ChartSpec) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ChartSpec) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *ChartSpec) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *ChartSpec) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


