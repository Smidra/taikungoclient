# InfraBillingSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfraProductId** | Pointer to **int32** |  | [optional] 
**InfraProductName** | Pointer to **NullableString** |  | [optional] 
**Intervals** | Pointer to [**[]DateInterval**](DateInterval.md) |  | [optional] 
**TotalPrice** | Pointer to **float64** |  | [optional] 

## Methods

### NewInfraBillingSummaryDto

`func NewInfraBillingSummaryDto() *InfraBillingSummaryDto`

NewInfraBillingSummaryDto instantiates a new InfraBillingSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraBillingSummaryDtoWithDefaults

`func NewInfraBillingSummaryDtoWithDefaults() *InfraBillingSummaryDto`

NewInfraBillingSummaryDtoWithDefaults instantiates a new InfraBillingSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfraProductId

`func (o *InfraBillingSummaryDto) GetInfraProductId() int32`

GetInfraProductId returns the InfraProductId field if non-nil, zero value otherwise.

### GetInfraProductIdOk

`func (o *InfraBillingSummaryDto) GetInfraProductIdOk() (*int32, bool)`

GetInfraProductIdOk returns a tuple with the InfraProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraProductId

`func (o *InfraBillingSummaryDto) SetInfraProductId(v int32)`

SetInfraProductId sets InfraProductId field to given value.

### HasInfraProductId

`func (o *InfraBillingSummaryDto) HasInfraProductId() bool`

HasInfraProductId returns a boolean if a field has been set.

### GetInfraProductName

`func (o *InfraBillingSummaryDto) GetInfraProductName() string`

GetInfraProductName returns the InfraProductName field if non-nil, zero value otherwise.

### GetInfraProductNameOk

`func (o *InfraBillingSummaryDto) GetInfraProductNameOk() (*string, bool)`

GetInfraProductNameOk returns a tuple with the InfraProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraProductName

`func (o *InfraBillingSummaryDto) SetInfraProductName(v string)`

SetInfraProductName sets InfraProductName field to given value.

### HasInfraProductName

`func (o *InfraBillingSummaryDto) HasInfraProductName() bool`

HasInfraProductName returns a boolean if a field has been set.

### SetInfraProductNameNil

`func (o *InfraBillingSummaryDto) SetInfraProductNameNil(b bool)`

 SetInfraProductNameNil sets the value for InfraProductName to be an explicit nil

### UnsetInfraProductName
`func (o *InfraBillingSummaryDto) UnsetInfraProductName()`

UnsetInfraProductName ensures that no value is present for InfraProductName, not even an explicit nil
### GetIntervals

`func (o *InfraBillingSummaryDto) GetIntervals() []DateInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InfraBillingSummaryDto) GetIntervalsOk() (*[]DateInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InfraBillingSummaryDto) SetIntervals(v []DateInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InfraBillingSummaryDto) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### SetIntervalsNil

`func (o *InfraBillingSummaryDto) SetIntervalsNil(b bool)`

 SetIntervalsNil sets the value for Intervals to be an explicit nil

### UnsetIntervals
`func (o *InfraBillingSummaryDto) UnsetIntervals()`

UnsetIntervals ensures that no value is present for Intervals, not even an explicit nil
### GetTotalPrice

`func (o *InfraBillingSummaryDto) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InfraBillingSummaryDto) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InfraBillingSummaryDto) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *InfraBillingSummaryDto) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


