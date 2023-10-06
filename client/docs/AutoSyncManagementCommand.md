# AutoSyncManagementCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoSyncManagementCommand

`func NewAutoSyncManagementCommand() *AutoSyncManagementCommand`

NewAutoSyncManagementCommand instantiates a new AutoSyncManagementCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSyncManagementCommandWithDefaults

`func NewAutoSyncManagementCommandWithDefaults() *AutoSyncManagementCommand`

NewAutoSyncManagementCommandWithDefaults instantiates a new AutoSyncManagementCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoSyncManagementCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoSyncManagementCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoSyncManagementCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutoSyncManagementCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *AutoSyncManagementCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AutoSyncManagementCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AutoSyncManagementCommand) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AutoSyncManagementCommand) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *AutoSyncManagementCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AutoSyncManagementCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


