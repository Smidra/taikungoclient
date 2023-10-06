# BindRulesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]BindRulesToOrganizationDto**](BindRulesToOrganizationDto.md) |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBindRulesCommand

`func NewBindRulesCommand() *BindRulesCommand`

NewBindRulesCommand instantiates a new BindRulesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindRulesCommandWithDefaults

`func NewBindRulesCommandWithDefaults() *BindRulesCommand`

NewBindRulesCommandWithDefaults instantiates a new BindRulesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *BindRulesCommand) GetRules() []BindRulesToOrganizationDto`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BindRulesCommand) GetRulesOk() (*[]BindRulesToOrganizationDto, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BindRulesCommand) SetRules(v []BindRulesToOrganizationDto)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BindRulesCommand) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *BindRulesCommand) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BindRulesCommand) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetOrganizationId

`func (o *BindRulesCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BindRulesCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BindRulesCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BindRulesCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


