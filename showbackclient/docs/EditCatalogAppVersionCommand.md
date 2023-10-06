# EditCatalogAppVersionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogAppId** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEditCatalogAppVersionCommand

`func NewEditCatalogAppVersionCommand() *EditCatalogAppVersionCommand`

NewEditCatalogAppVersionCommand instantiates a new EditCatalogAppVersionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditCatalogAppVersionCommandWithDefaults

`func NewEditCatalogAppVersionCommandWithDefaults() *EditCatalogAppVersionCommand`

NewEditCatalogAppVersionCommandWithDefaults instantiates a new EditCatalogAppVersionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogAppId

`func (o *EditCatalogAppVersionCommand) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *EditCatalogAppVersionCommand) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *EditCatalogAppVersionCommand) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.

### HasCatalogAppId

`func (o *EditCatalogAppVersionCommand) HasCatalogAppId() bool`

HasCatalogAppId returns a boolean if a field has been set.

### GetVersion

`func (o *EditCatalogAppVersionCommand) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EditCatalogAppVersionCommand) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EditCatalogAppVersionCommand) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EditCatalogAppVersionCommand) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EditCatalogAppVersionCommand) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EditCatalogAppVersionCommand) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


