# ResetProjectStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**ProjectStatus**](ProjectStatus.md) |  | [optional] 

## Methods

### NewResetProjectStatusCommand

`func NewResetProjectStatusCommand() *ResetProjectStatusCommand`

NewResetProjectStatusCommand instantiates a new ResetProjectStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetProjectStatusCommandWithDefaults

`func NewResetProjectStatusCommandWithDefaults() *ResetProjectStatusCommand`

NewResetProjectStatusCommandWithDefaults instantiates a new ResetProjectStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ResetProjectStatusCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ResetProjectStatusCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ResetProjectStatusCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ResetProjectStatusCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *ResetProjectStatusCommand) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResetProjectStatusCommand) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResetProjectStatusCommand) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResetProjectStatusCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


