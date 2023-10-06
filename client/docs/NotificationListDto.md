# NotificationListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**ActionMessage** | Pointer to **NullableString** |  | [optional] 
**ActionStatus** | Pointer to [**ActionStatus**](ActionStatus.md) |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotificationListDto

`func NewNotificationListDto() *NotificationListDto`

NewNotificationListDto instantiates a new NotificationListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListDtoWithDefaults

`func NewNotificationListDtoWithDefaults() *NotificationListDto`

NewNotificationListDtoWithDefaults instantiates a new NotificationListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NotificationListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *NotificationListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NotificationListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetActionMessage

`func (o *NotificationListDto) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *NotificationListDto) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *NotificationListDto) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *NotificationListDto) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### SetActionMessageNil

`func (o *NotificationListDto) SetActionMessageNil(b bool)`

 SetActionMessageNil sets the value for ActionMessage to be an explicit nil

### UnsetActionMessage
`func (o *NotificationListDto) UnsetActionMessage()`

UnsetActionMessage ensures that no value is present for ActionMessage, not even an explicit nil
### GetActionStatus

`func (o *NotificationListDto) GetActionStatus() ActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NotificationListDto) GetActionStatusOk() (*ActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NotificationListDto) SetActionStatus(v ActionStatus)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NotificationListDto) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationListDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationListDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *NotificationListDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *NotificationListDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetCategory

`func (o *NotificationListDto) GetCategory() ActionType`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NotificationListDto) GetCategoryOk() (*ActionType, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NotificationListDto) SetCategory(v ActionType)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NotificationListDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetProjectName

`func (o *NotificationListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *NotificationListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *NotificationListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *NotificationListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *NotificationListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *NotificationListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetProjectId

`func (o *NotificationListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NotificationListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NotificationListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NotificationListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *NotificationListDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *NotificationListDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *NotificationListDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *NotificationListDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


