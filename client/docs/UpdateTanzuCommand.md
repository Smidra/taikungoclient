# UpdateTanzuCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateTanzuCommand

`func NewUpdateTanzuCommand() *UpdateTanzuCommand`

NewUpdateTanzuCommand instantiates a new UpdateTanzuCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTanzuCommandWithDefaults

`func NewUpdateTanzuCommandWithDefaults() *UpdateTanzuCommand`

NewUpdateTanzuCommandWithDefaults instantiates a new UpdateTanzuCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTanzuCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTanzuCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTanzuCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTanzuCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateTanzuCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTanzuCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTanzuCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTanzuCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTanzuCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTanzuCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPassword

`func (o *UpdateTanzuCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateTanzuCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateTanzuCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateTanzuCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateTanzuCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateTanzuCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


