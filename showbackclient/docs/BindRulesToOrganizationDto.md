# BindRulesToOrganizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DiscountRate** | Pointer to **NullableFloat64** |  | [optional] 
**IsBound** | Pointer to **bool** |  | [optional] 

## Methods

### NewBindRulesToOrganizationDto

`func NewBindRulesToOrganizationDto() *BindRulesToOrganizationDto`

NewBindRulesToOrganizationDto instantiates a new BindRulesToOrganizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindRulesToOrganizationDtoWithDefaults

`func NewBindRulesToOrganizationDtoWithDefaults() *BindRulesToOrganizationDto`

NewBindRulesToOrganizationDtoWithDefaults instantiates a new BindRulesToOrganizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BindRulesToOrganizationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BindRulesToOrganizationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BindRulesToOrganizationDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BindRulesToOrganizationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BindRulesToOrganizationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BindRulesToOrganizationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BindRulesToOrganizationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BindRulesToOrganizationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BindRulesToOrganizationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BindRulesToOrganizationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDiscountRate

`func (o *BindRulesToOrganizationDto) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *BindRulesToOrganizationDto) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *BindRulesToOrganizationDto) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *BindRulesToOrganizationDto) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### SetDiscountRateNil

`func (o *BindRulesToOrganizationDto) SetDiscountRateNil(b bool)`

 SetDiscountRateNil sets the value for DiscountRate to be an explicit nil

### UnsetDiscountRate
`func (o *BindRulesToOrganizationDto) UnsetDiscountRate()`

UnsetDiscountRate ensures that no value is present for DiscountRate, not even an explicit nil
### GetIsBound

`func (o *BindRulesToOrganizationDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *BindRulesToOrganizationDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *BindRulesToOrganizationDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *BindRulesToOrganizationDto) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


