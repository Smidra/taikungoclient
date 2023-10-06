# NotificationSendCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ActionType** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**ActionStatus** | Pointer to [**ActionStatus**](ActionStatus.md) |  | [optional] 
**ProjectType** | Pointer to [**ProjectType**](ProjectType.md) |  | [optional] 

## Methods

### NewNotificationSendCommand

`func NewNotificationSendCommand() *NotificationSendCommand`

NewNotificationSendCommand instantiates a new NotificationSendCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSendCommandWithDefaults

`func NewNotificationSendCommandWithDefaults() *NotificationSendCommand`

NewNotificationSendCommandWithDefaults instantiates a new NotificationSendCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *NotificationSendCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NotificationSendCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NotificationSendCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NotificationSendCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetActionType

`func (o *NotificationSendCommand) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *NotificationSendCommand) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *NotificationSendCommand) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *NotificationSendCommand) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionStatus

`func (o *NotificationSendCommand) GetActionStatus() ActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NotificationSendCommand) GetActionStatusOk() (*ActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NotificationSendCommand) SetActionStatus(v ActionStatus)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NotificationSendCommand) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetProjectType

`func (o *NotificationSendCommand) GetProjectType() ProjectType`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *NotificationSendCommand) GetProjectTypeOk() (*ProjectType, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *NotificationSendCommand) SetProjectType(v ProjectType)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *NotificationSendCommand) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


