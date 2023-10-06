# InfraBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]InfraBillingSummaryDto**](InfraBillingSummaryDto.md) |  | [optional] 
**TotalPrice** | Pointer to **float64** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewInfraBillingInfo

`func NewInfraBillingInfo() *InfraBillingInfo`

NewInfraBillingInfo instantiates a new InfraBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraBillingInfoWithDefaults

`func NewInfraBillingInfoWithDefaults() *InfraBillingInfo`

NewInfraBillingInfoWithDefaults instantiates a new InfraBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InfraBillingInfo) GetData() []InfraBillingSummaryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InfraBillingInfo) GetDataOk() (*[]InfraBillingSummaryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InfraBillingInfo) SetData(v []InfraBillingSummaryDto)`

SetData sets Data field to given value.

### HasData

`func (o *InfraBillingInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *InfraBillingInfo) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *InfraBillingInfo) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalPrice

`func (o *InfraBillingInfo) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InfraBillingInfo) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InfraBillingInfo) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *InfraBillingInfo) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTotalCount

`func (o *InfraBillingInfo) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InfraBillingInfo) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InfraBillingInfo) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *InfraBillingInfo) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


