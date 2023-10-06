# CardInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationMonth** | Pointer to **NullableString** |  | [optional] 
**ExpirationYear** | Pointer to **NullableString** |  | [optional] 
**Last4** | Pointer to **NullableString** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**HolderName** | Pointer to **NullableString** |  | [optional] 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewCardInformationDto

`func NewCardInformationDto() *CardInformationDto`

NewCardInformationDto instantiates a new CardInformationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInformationDtoWithDefaults

`func NewCardInformationDtoWithDefaults() *CardInformationDto`

NewCardInformationDtoWithDefaults instantiates a new CardInformationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationMonth

`func (o *CardInformationDto) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *CardInformationDto) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *CardInformationDto) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *CardInformationDto) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### SetExpirationMonthNil

`func (o *CardInformationDto) SetExpirationMonthNil(b bool)`

 SetExpirationMonthNil sets the value for ExpirationMonth to be an explicit nil

### UnsetExpirationMonth
`func (o *CardInformationDto) UnsetExpirationMonth()`

UnsetExpirationMonth ensures that no value is present for ExpirationMonth, not even an explicit nil
### GetExpirationYear

`func (o *CardInformationDto) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *CardInformationDto) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *CardInformationDto) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *CardInformationDto) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### SetExpirationYearNil

`func (o *CardInformationDto) SetExpirationYearNil(b bool)`

 SetExpirationYearNil sets the value for ExpirationYear to be an explicit nil

### UnsetExpirationYear
`func (o *CardInformationDto) UnsetExpirationYear()`

UnsetExpirationYear ensures that no value is present for ExpirationYear, not even an explicit nil
### GetLast4

`func (o *CardInformationDto) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *CardInformationDto) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *CardInformationDto) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *CardInformationDto) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *CardInformationDto) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *CardInformationDto) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetBrand

`func (o *CardInformationDto) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CardInformationDto) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CardInformationDto) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CardInformationDto) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CardInformationDto) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CardInformationDto) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetHolderName

`func (o *CardInformationDto) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *CardInformationDto) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *CardInformationDto) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *CardInformationDto) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### SetHolderNameNil

`func (o *CardInformationDto) SetHolderNameNil(b bool)`

 SetHolderNameNil sets the value for HolderName to be an explicit nil

### UnsetHolderName
`func (o *CardInformationDto) UnsetHolderName()`

UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
### GetBalance

`func (o *CardInformationDto) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CardInformationDto) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CardInformationDto) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CardInformationDto) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


