# InfraBillingSummariesCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**BeginApply** | Pointer to **NullableTime** |  | [optional] 
**EndApply** | Pointer to **NullableTime** |  | [optional] 
**ProductId** | Pointer to **int32** |  | [optional] 

## Methods

### NewInfraBillingSummariesCreateCommand

`func NewInfraBillingSummariesCreateCommand() *InfraBillingSummariesCreateCommand`

NewInfraBillingSummariesCreateCommand instantiates a new InfraBillingSummariesCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraBillingSummariesCreateCommandWithDefaults

`func NewInfraBillingSummariesCreateCommandWithDefaults() *InfraBillingSummariesCreateCommand`

NewInfraBillingSummariesCreateCommandWithDefaults instantiates a new InfraBillingSummariesCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *InfraBillingSummariesCreateCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InfraBillingSummariesCreateCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InfraBillingSummariesCreateCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InfraBillingSummariesCreateCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InfraBillingSummariesCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InfraBillingSummariesCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InfraBillingSummariesCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InfraBillingSummariesCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetBeginApply

`func (o *InfraBillingSummariesCreateCommand) GetBeginApply() time.Time`

GetBeginApply returns the BeginApply field if non-nil, zero value otherwise.

### GetBeginApplyOk

`func (o *InfraBillingSummariesCreateCommand) GetBeginApplyOk() (*time.Time, bool)`

GetBeginApplyOk returns a tuple with the BeginApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginApply

`func (o *InfraBillingSummariesCreateCommand) SetBeginApply(v time.Time)`

SetBeginApply sets BeginApply field to given value.

### HasBeginApply

`func (o *InfraBillingSummariesCreateCommand) HasBeginApply() bool`

HasBeginApply returns a boolean if a field has been set.

### SetBeginApplyNil

`func (o *InfraBillingSummariesCreateCommand) SetBeginApplyNil(b bool)`

 SetBeginApplyNil sets the value for BeginApply to be an explicit nil

### UnsetBeginApply
`func (o *InfraBillingSummariesCreateCommand) UnsetBeginApply()`

UnsetBeginApply ensures that no value is present for BeginApply, not even an explicit nil
### GetEndApply

`func (o *InfraBillingSummariesCreateCommand) GetEndApply() time.Time`

GetEndApply returns the EndApply field if non-nil, zero value otherwise.

### GetEndApplyOk

`func (o *InfraBillingSummariesCreateCommand) GetEndApplyOk() (*time.Time, bool)`

GetEndApplyOk returns a tuple with the EndApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndApply

`func (o *InfraBillingSummariesCreateCommand) SetEndApply(v time.Time)`

SetEndApply sets EndApply field to given value.

### HasEndApply

`func (o *InfraBillingSummariesCreateCommand) HasEndApply() bool`

HasEndApply returns a boolean if a field has been set.

### SetEndApplyNil

`func (o *InfraBillingSummariesCreateCommand) SetEndApplyNil(b bool)`

 SetEndApplyNil sets the value for EndApply to be an explicit nil

### UnsetEndApply
`func (o *InfraBillingSummariesCreateCommand) UnsetEndApply()`

UnsetEndApply ensures that no value is present for EndApply, not even an explicit nil
### GetProductId

`func (o *InfraBillingSummariesCreateCommand) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *InfraBillingSummariesCreateCommand) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *InfraBillingSummariesCreateCommand) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *InfraBillingSummariesCreateCommand) HasProductId() bool`

HasProductId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


