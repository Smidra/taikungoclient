# KubeConfigRoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKubeConfigRoleDto

`func NewKubeConfigRoleDto() *KubeConfigRoleDto`

NewKubeConfigRoleDto instantiates a new KubeConfigRoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigRoleDtoWithDefaults

`func NewKubeConfigRoleDtoWithDefaults() *KubeConfigRoleDto`

NewKubeConfigRoleDtoWithDefaults instantiates a new KubeConfigRoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubeConfigRoleDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubeConfigRoleDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubeConfigRoleDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubeConfigRoleDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubeConfigRoleDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubeConfigRoleDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubeConfigRoleDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubeConfigRoleDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubeConfigRoleDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubeConfigRoleDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


