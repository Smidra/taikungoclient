# BindProjectsToCatalogCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]UpdateCatalogDto**](UpdateCatalogDto.md) |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBindProjectsToCatalogCommand

`func NewBindProjectsToCatalogCommand() *BindProjectsToCatalogCommand`

NewBindProjectsToCatalogCommand instantiates a new BindProjectsToCatalogCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindProjectsToCatalogCommandWithDefaults

`func NewBindProjectsToCatalogCommandWithDefaults() *BindProjectsToCatalogCommand`

NewBindProjectsToCatalogCommandWithDefaults instantiates a new BindProjectsToCatalogCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *BindProjectsToCatalogCommand) GetProjects() []UpdateCatalogDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BindProjectsToCatalogCommand) GetProjectsOk() (*[]UpdateCatalogDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BindProjectsToCatalogCommand) SetProjects(v []UpdateCatalogDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *BindProjectsToCatalogCommand) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *BindProjectsToCatalogCommand) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *BindProjectsToCatalogCommand) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetCatalogId

`func (o *BindProjectsToCatalogCommand) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *BindProjectsToCatalogCommand) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *BindProjectsToCatalogCommand) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *BindProjectsToCatalogCommand) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


