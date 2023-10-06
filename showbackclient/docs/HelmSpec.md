# HelmSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **NullableString** |  | [optional] 
**ReleaseName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**TargetNamespace** | Pointer to **NullableString** |  | [optional] 
**StorageNamespace** | Pointer to **NullableString** |  | [optional] 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Values** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 

## Methods

### NewHelmSpec

`func NewHelmSpec() *HelmSpec`

NewHelmSpec instantiates a new HelmSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmSpecWithDefaults

`func NewHelmSpecWithDefaults() *HelmSpec`

NewHelmSpecWithDefaults instantiates a new HelmSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *HelmSpec) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HelmSpec) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HelmSpec) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HelmSpec) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *HelmSpec) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *HelmSpec) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetReleaseName

`func (o *HelmSpec) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *HelmSpec) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *HelmSpec) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *HelmSpec) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### SetReleaseNameNil

`func (o *HelmSpec) SetReleaseNameNil(b bool)`

 SetReleaseNameNil sets the value for ReleaseName to be an explicit nil

### UnsetReleaseName
`func (o *HelmSpec) UnsetReleaseName()`

UnsetReleaseName ensures that no value is present for ReleaseName, not even an explicit nil
### GetUrl

`func (o *HelmSpec) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HelmSpec) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HelmSpec) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HelmSpec) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *HelmSpec) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *HelmSpec) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTargetNamespace

`func (o *HelmSpec) GetTargetNamespace() string`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *HelmSpec) GetTargetNamespaceOk() (*string, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *HelmSpec) SetTargetNamespace(v string)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *HelmSpec) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

### SetTargetNamespaceNil

`func (o *HelmSpec) SetTargetNamespaceNil(b bool)`

 SetTargetNamespaceNil sets the value for TargetNamespace to be an explicit nil

### UnsetTargetNamespace
`func (o *HelmSpec) UnsetTargetNamespace()`

UnsetTargetNamespace ensures that no value is present for TargetNamespace, not even an explicit nil
### GetStorageNamespace

`func (o *HelmSpec) GetStorageNamespace() string`

GetStorageNamespace returns the StorageNamespace field if non-nil, zero value otherwise.

### GetStorageNamespaceOk

`func (o *HelmSpec) GetStorageNamespaceOk() (*string, bool)`

GetStorageNamespaceOk returns a tuple with the StorageNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNamespace

`func (o *HelmSpec) SetStorageNamespace(v string)`

SetStorageNamespace sets StorageNamespace field to given value.

### HasStorageNamespace

`func (o *HelmSpec) HasStorageNamespace() bool`

HasStorageNamespace returns a boolean if a field has been set.

### SetStorageNamespaceNil

`func (o *HelmSpec) SetStorageNamespaceNil(b bool)`

 SetStorageNamespaceNil sets the value for StorageNamespace to be an explicit nil

### UnsetStorageNamespace
`func (o *HelmSpec) UnsetStorageNamespace()`

UnsetStorageNamespace ensures that no value is present for StorageNamespace, not even an explicit nil
### GetChart

`func (o *HelmSpec) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmSpec) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmSpec) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *HelmSpec) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetValues

`func (o *HelmSpec) GetValues() JsonNode`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HelmSpec) GetValuesOk() (*JsonNode, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HelmSpec) SetValues(v JsonNode)`

SetValues sets Values field to given value.

### HasValues

`func (o *HelmSpec) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


