# BindPrometheusOrganizationsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]BindOrganizationsToRuleDto**](BindOrganizationsToRuleDto.md) |  | [optional] 
**PrometheusRuleId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBindPrometheusOrganizationsCommand

`func NewBindPrometheusOrganizationsCommand() *BindPrometheusOrganizationsCommand`

NewBindPrometheusOrganizationsCommand instantiates a new BindPrometheusOrganizationsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindPrometheusOrganizationsCommandWithDefaults

`func NewBindPrometheusOrganizationsCommandWithDefaults() *BindPrometheusOrganizationsCommand`

NewBindPrometheusOrganizationsCommandWithDefaults instantiates a new BindPrometheusOrganizationsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *BindPrometheusOrganizationsCommand) GetOrganizations() []BindOrganizationsToRuleDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *BindPrometheusOrganizationsCommand) GetOrganizationsOk() (*[]BindOrganizationsToRuleDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *BindPrometheusOrganizationsCommand) SetOrganizations(v []BindOrganizationsToRuleDto)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *BindPrometheusOrganizationsCommand) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### SetOrganizationsNil

`func (o *BindPrometheusOrganizationsCommand) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *BindPrometheusOrganizationsCommand) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil
### GetPrometheusRuleId

`func (o *BindPrometheusOrganizationsCommand) GetPrometheusRuleId() int32`

GetPrometheusRuleId returns the PrometheusRuleId field if non-nil, zero value otherwise.

### GetPrometheusRuleIdOk

`func (o *BindPrometheusOrganizationsCommand) GetPrometheusRuleIdOk() (*int32, bool)`

GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusRuleId

`func (o *BindPrometheusOrganizationsCommand) SetPrometheusRuleId(v int32)`

SetPrometheusRuleId sets PrometheusRuleId field to given value.

### HasPrometheusRuleId

`func (o *BindPrometheusOrganizationsCommand) HasPrometheusRuleId() bool`

HasPrometheusRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


