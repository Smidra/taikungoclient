# UserGroupEntityListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserGroupEntityListDto

`func NewUserGroupEntityListDto() *UserGroupEntityListDto`

NewUserGroupEntityListDto instantiates a new UserGroupEntityListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupEntityListDtoWithDefaults

`func NewUserGroupEntityListDtoWithDefaults() *UserGroupEntityListDto`

NewUserGroupEntityListDtoWithDefaults instantiates a new UserGroupEntityListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupEntityListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupEntityListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupEntityListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupEntityListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserGroupEntityListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupEntityListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupEntityListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroupEntityListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserGroupEntityListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserGroupEntityListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


