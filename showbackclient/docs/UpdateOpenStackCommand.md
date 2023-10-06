# UpdateOpenStackCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OpenStackUser** | Pointer to **NullableString** |  | [optional] 
**OpenStackPassword** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateOpenStackCommand

`func NewUpdateOpenStackCommand() *UpdateOpenStackCommand`

NewUpdateOpenStackCommand instantiates a new UpdateOpenStackCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOpenStackCommandWithDefaults

`func NewUpdateOpenStackCommandWithDefaults() *UpdateOpenStackCommand`

NewUpdateOpenStackCommandWithDefaults instantiates a new UpdateOpenStackCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOpenStackCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOpenStackCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOpenStackCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOpenStackCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOpenStackCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOpenStackCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOpenStackCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOpenStackCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateOpenStackCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateOpenStackCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOpenStackUser

`func (o *UpdateOpenStackCommand) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *UpdateOpenStackCommand) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *UpdateOpenStackCommand) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.

### HasOpenStackUser

`func (o *UpdateOpenStackCommand) HasOpenStackUser() bool`

HasOpenStackUser returns a boolean if a field has been set.

### SetOpenStackUserNil

`func (o *UpdateOpenStackCommand) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *UpdateOpenStackCommand) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *UpdateOpenStackCommand) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *UpdateOpenStackCommand) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *UpdateOpenStackCommand) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.

### HasOpenStackPassword

`func (o *UpdateOpenStackCommand) HasOpenStackPassword() bool`

HasOpenStackPassword returns a boolean if a field has been set.

### SetOpenStackPasswordNil

`func (o *UpdateOpenStackCommand) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *UpdateOpenStackCommand) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


