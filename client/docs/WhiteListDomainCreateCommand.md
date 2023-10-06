# WhiteListDomainCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhiteListDomains** | Pointer to [**[]WhiteListDomainCreateDto**](WhiteListDomainCreateDto.md) |  | [optional] 
**PartnerId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewWhiteListDomainCreateCommand

`func NewWhiteListDomainCreateCommand() *WhiteListDomainCreateCommand`

NewWhiteListDomainCreateCommand instantiates a new WhiteListDomainCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteListDomainCreateCommandWithDefaults

`func NewWhiteListDomainCreateCommandWithDefaults() *WhiteListDomainCreateCommand`

NewWhiteListDomainCreateCommandWithDefaults instantiates a new WhiteListDomainCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhiteListDomains

`func (o *WhiteListDomainCreateCommand) GetWhiteListDomains() []WhiteListDomainCreateDto`

GetWhiteListDomains returns the WhiteListDomains field if non-nil, zero value otherwise.

### GetWhiteListDomainsOk

`func (o *WhiteListDomainCreateCommand) GetWhiteListDomainsOk() (*[]WhiteListDomainCreateDto, bool)`

GetWhiteListDomainsOk returns a tuple with the WhiteListDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListDomains

`func (o *WhiteListDomainCreateCommand) SetWhiteListDomains(v []WhiteListDomainCreateDto)`

SetWhiteListDomains sets WhiteListDomains field to given value.

### HasWhiteListDomains

`func (o *WhiteListDomainCreateCommand) HasWhiteListDomains() bool`

HasWhiteListDomains returns a boolean if a field has been set.

### SetWhiteListDomainsNil

`func (o *WhiteListDomainCreateCommand) SetWhiteListDomainsNil(b bool)`

 SetWhiteListDomainsNil sets the value for WhiteListDomains to be an explicit nil

### UnsetWhiteListDomains
`func (o *WhiteListDomainCreateCommand) UnsetWhiteListDomains()`

UnsetWhiteListDomains ensures that no value is present for WhiteListDomains, not even an explicit nil
### GetPartnerId

`func (o *WhiteListDomainCreateCommand) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *WhiteListDomainCreateCommand) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *WhiteListDomainCreateCommand) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *WhiteListDomainCreateCommand) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *WhiteListDomainCreateCommand) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *WhiteListDomainCreateCommand) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


