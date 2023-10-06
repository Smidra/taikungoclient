# DateInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 

## Methods

### NewDateInterval

`func NewDateInterval() *DateInterval`

NewDateInterval instantiates a new DateInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateIntervalWithDefaults

`func NewDateIntervalWithDefaults() *DateInterval`

NewDateIntervalWithDefaults instantiates a new DateInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *DateInterval) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DateInterval) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DateInterval) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *DateInterval) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *DateInterval) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DateInterval) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DateInterval) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DateInterval) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPrice

`func (o *DateInterval) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DateInterval) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DateInterval) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DateInterval) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


