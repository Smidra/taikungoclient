# StandAloneProfileFullDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**StandAloneProfileSecurityGroups** | Pointer to [**[]StandAloneProfileSecurityGroupFullDto**](StandAloneProfileSecurityGroupFullDto.md) |  | [optional] 

## Methods

### NewStandAloneProfileFullDto

`func NewStandAloneProfileFullDto() *StandAloneProfileFullDto`

NewStandAloneProfileFullDto instantiates a new StandAloneProfileFullDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfileFullDtoWithDefaults

`func NewStandAloneProfileFullDtoWithDefaults() *StandAloneProfileFullDto`

NewStandAloneProfileFullDtoWithDefaults instantiates a new StandAloneProfileFullDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneProfileFullDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneProfileFullDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneProfileFullDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandAloneProfileFullDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandAloneProfileFullDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfileFullDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfileFullDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandAloneProfileFullDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StandAloneProfileFullDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneProfileFullDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublicKey

`func (o *StandAloneProfileFullDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *StandAloneProfileFullDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *StandAloneProfileFullDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *StandAloneProfileFullDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *StandAloneProfileFullDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *StandAloneProfileFullDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetRevision

`func (o *StandAloneProfileFullDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StandAloneProfileFullDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StandAloneProfileFullDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StandAloneProfileFullDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStandAloneProfileSecurityGroups

`func (o *StandAloneProfileFullDto) GetStandAloneProfileSecurityGroups() []StandAloneProfileSecurityGroupFullDto`

GetStandAloneProfileSecurityGroups returns the StandAloneProfileSecurityGroups field if non-nil, zero value otherwise.

### GetStandAloneProfileSecurityGroupsOk

`func (o *StandAloneProfileFullDto) GetStandAloneProfileSecurityGroupsOk() (*[]StandAloneProfileSecurityGroupFullDto, bool)`

GetStandAloneProfileSecurityGroupsOk returns a tuple with the StandAloneProfileSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfileSecurityGroups

`func (o *StandAloneProfileFullDto) SetStandAloneProfileSecurityGroups(v []StandAloneProfileSecurityGroupFullDto)`

SetStandAloneProfileSecurityGroups sets StandAloneProfileSecurityGroups field to given value.

### HasStandAloneProfileSecurityGroups

`func (o *StandAloneProfileFullDto) HasStandAloneProfileSecurityGroups() bool`

HasStandAloneProfileSecurityGroups returns a boolean if a field has been set.

### SetStandAloneProfileSecurityGroupsNil

`func (o *StandAloneProfileFullDto) SetStandAloneProfileSecurityGroupsNil(b bool)`

 SetStandAloneProfileSecurityGroupsNil sets the value for StandAloneProfileSecurityGroups to be an explicit nil

### UnsetStandAloneProfileSecurityGroups
`func (o *StandAloneProfileFullDto) UnsetStandAloneProfileSecurityGroups()`

UnsetStandAloneProfileSecurityGroups ensures that no value is present for StandAloneProfileSecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


