# StandaloneProfileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**StandAloneProfileSecurityGroups** | Pointer to [**[]StandaloneProfileSecurityGroupListDto**](StandaloneProfileSecurityGroupListDto.md) |  | [optional] 

## Methods

### NewStandaloneProfileListDto

`func NewStandaloneProfileListDto() *StandaloneProfileListDto`

NewStandaloneProfileListDto instantiates a new StandaloneProfileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneProfileListDtoWithDefaults

`func NewStandaloneProfileListDtoWithDefaults() *StandaloneProfileListDto`

NewStandaloneProfileListDtoWithDefaults instantiates a new StandaloneProfileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneProfileListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneProfileListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneProfileListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandaloneProfileListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandaloneProfileListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneProfileListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneProfileListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandaloneProfileListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandaloneProfileListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandaloneProfileListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRevision

`func (o *StandaloneProfileListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StandaloneProfileListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StandaloneProfileListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StandaloneProfileListDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetIsLocked

`func (o *StandaloneProfileListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *StandaloneProfileListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *StandaloneProfileListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *StandaloneProfileListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetStandAloneProfileSecurityGroups

`func (o *StandaloneProfileListDto) GetStandAloneProfileSecurityGroups() []StandaloneProfileSecurityGroupListDto`

GetStandAloneProfileSecurityGroups returns the StandAloneProfileSecurityGroups field if non-nil, zero value otherwise.

### GetStandAloneProfileSecurityGroupsOk

`func (o *StandaloneProfileListDto) GetStandAloneProfileSecurityGroupsOk() (*[]StandaloneProfileSecurityGroupListDto, bool)`

GetStandAloneProfileSecurityGroupsOk returns a tuple with the StandAloneProfileSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfileSecurityGroups

`func (o *StandaloneProfileListDto) SetStandAloneProfileSecurityGroups(v []StandaloneProfileSecurityGroupListDto)`

SetStandAloneProfileSecurityGroups sets StandAloneProfileSecurityGroups field to given value.

### HasStandAloneProfileSecurityGroups

`func (o *StandaloneProfileListDto) HasStandAloneProfileSecurityGroups() bool`

HasStandAloneProfileSecurityGroups returns a boolean if a field has been set.

### SetStandAloneProfileSecurityGroupsNil

`func (o *StandaloneProfileListDto) SetStandAloneProfileSecurityGroupsNil(b bool)`

 SetStandAloneProfileSecurityGroupsNil sets the value for StandAloneProfileSecurityGroups to be an explicit nil

### UnsetStandAloneProfileSecurityGroups
`func (o *StandaloneProfileListDto) UnsetStandAloneProfileSecurityGroups()`

UnsetStandAloneProfileSecurityGroups ensures that no value is present for StandAloneProfileSecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


