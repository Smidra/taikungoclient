# AvailableEndpointsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EndpointElements**](EndpointElements.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAvailableEndpointsList

`func NewAvailableEndpointsList() *AvailableEndpointsList`

NewAvailableEndpointsList instantiates a new AvailableEndpointsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableEndpointsListWithDefaults

`func NewAvailableEndpointsListWithDefaults() *AvailableEndpointsList`

NewAvailableEndpointsListWithDefaults instantiates a new AvailableEndpointsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AvailableEndpointsList) GetData() []EndpointElements`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AvailableEndpointsList) GetDataOk() (*[]EndpointElements, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AvailableEndpointsList) SetData(v []EndpointElements)`

SetData sets Data field to given value.

### HasData

`func (o *AvailableEndpointsList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AvailableEndpointsList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AvailableEndpointsList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AvailableEndpointsList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AvailableEndpointsList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AvailableEndpointsList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AvailableEndpointsList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


