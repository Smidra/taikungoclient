# BindOrganizationsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]OrganizationDto**](OrganizationDto.md) |  | [optional] 
**PartnerId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBindOrganizationsCommand

`func NewBindOrganizationsCommand() *BindOrganizationsCommand`

NewBindOrganizationsCommand instantiates a new BindOrganizationsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindOrganizationsCommandWithDefaults

`func NewBindOrganizationsCommandWithDefaults() *BindOrganizationsCommand`

NewBindOrganizationsCommandWithDefaults instantiates a new BindOrganizationsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *BindOrganizationsCommand) GetOrganizations() []OrganizationDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *BindOrganizationsCommand) GetOrganizationsOk() (*[]OrganizationDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *BindOrganizationsCommand) SetOrganizations(v []OrganizationDto)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *BindOrganizationsCommand) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### SetOrganizationsNil

`func (o *BindOrganizationsCommand) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *BindOrganizationsCommand) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil
### GetPartnerId

`func (o *BindOrganizationsCommand) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BindOrganizationsCommand) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BindOrganizationsCommand) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *BindOrganizationsCommand) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *BindOrganizationsCommand) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *BindOrganizationsCommand) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


