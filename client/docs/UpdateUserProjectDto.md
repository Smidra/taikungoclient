# UpdateUserProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsBound** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateUserProjectDto

`func NewUpdateUserProjectDto() *UpdateUserProjectDto`

NewUpdateUserProjectDto instantiates a new UpdateUserProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserProjectDtoWithDefaults

`func NewUpdateUserProjectDtoWithDefaults() *UpdateUserProjectDto`

NewUpdateUserProjectDtoWithDefaults instantiates a new UpdateUserProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateUserProjectDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserProjectDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserProjectDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUserProjectDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserProjectDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserProjectDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserProjectDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserProjectDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateUserProjectDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateUserProjectDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsBound

`func (o *UpdateUserProjectDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *UpdateUserProjectDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *UpdateUserProjectDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *UpdateUserProjectDto) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


