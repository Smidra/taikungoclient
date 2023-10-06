# UserDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserForListDto**](UserForListDto.md) |  | [optional] 
**IsMaintenanceModeEnabled** | Pointer to **bool** |  | [optional] 
**DemoOrganization** | Pointer to **NullableInt32** |  | [optional] 
**TrialDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUserDetails

`func NewUserDetails() *UserDetails`

NewUserDetails instantiates a new UserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailsWithDefaults

`func NewUserDetailsWithDefaults() *UserDetails`

NewUserDetailsWithDefaults instantiates a new UserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserDetails) GetData() UserForListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserDetails) GetDataOk() (*UserForListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserDetails) SetData(v UserForListDto)`

SetData sets Data field to given value.

### HasData

`func (o *UserDetails) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIsMaintenanceModeEnabled

`func (o *UserDetails) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *UserDetails) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *UserDetails) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.

### HasIsMaintenanceModeEnabled

`func (o *UserDetails) HasIsMaintenanceModeEnabled() bool`

HasIsMaintenanceModeEnabled returns a boolean if a field has been set.

### GetDemoOrganization

`func (o *UserDetails) GetDemoOrganization() int32`

GetDemoOrganization returns the DemoOrganization field if non-nil, zero value otherwise.

### GetDemoOrganizationOk

`func (o *UserDetails) GetDemoOrganizationOk() (*int32, bool)`

GetDemoOrganizationOk returns a tuple with the DemoOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoOrganization

`func (o *UserDetails) SetDemoOrganization(v int32)`

SetDemoOrganization sets DemoOrganization field to given value.

### HasDemoOrganization

`func (o *UserDetails) HasDemoOrganization() bool`

HasDemoOrganization returns a boolean if a field has been set.

### SetDemoOrganizationNil

`func (o *UserDetails) SetDemoOrganizationNil(b bool)`

 SetDemoOrganizationNil sets the value for DemoOrganization to be an explicit nil

### UnsetDemoOrganization
`func (o *UserDetails) UnsetDemoOrganization()`

UnsetDemoOrganization ensures that no value is present for DemoOrganization, not even an explicit nil
### GetTrialDays

`func (o *UserDetails) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *UserDetails) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *UserDetails) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *UserDetails) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### SetTrialDaysNil

`func (o *UserDetails) SetTrialDaysNil(b bool)`

 SetTrialDaysNil sets the value for TrialDays to be an explicit nil

### UnsetTrialDays
`func (o *UserDetails) UnsetTrialDays()`

UnsetTrialDays ensures that no value is present for TrialDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


