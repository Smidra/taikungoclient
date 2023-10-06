# AlertingProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**SlackConfigurationId** | Pointer to **NullableInt32** |  | [optional] 
**SlackConfigurationName** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Emails** | Pointer to [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | [optional] 
**Webhooks** | Pointer to [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**Reminder** | Pointer to [**AlertingReminder**](AlertingReminder.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertingProfilesListDto

`func NewAlertingProfilesListDto() *AlertingProfilesListDto`

NewAlertingProfilesListDto instantiates a new AlertingProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfilesListDtoWithDefaults

`func NewAlertingProfilesListDtoWithDefaults() *AlertingProfilesListDto`

NewAlertingProfilesListDtoWithDefaults instantiates a new AlertingProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertingProfilesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertingProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertingProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertingProfilesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertingProfilesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AlertingProfilesListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AlertingProfilesListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *AlertingProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertingProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertingProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AlertingProfilesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *AlertingProfilesListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AlertingProfilesListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *AlertingProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AlertingProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AlertingProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AlertingProfilesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AlertingProfilesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AlertingProfilesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetSlackConfigurationId

`func (o *AlertingProfilesListDto) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *AlertingProfilesListDto) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.

### HasSlackConfigurationId

`func (o *AlertingProfilesListDto) HasSlackConfigurationId() bool`

HasSlackConfigurationId returns a boolean if a field has been set.

### SetSlackConfigurationIdNil

`func (o *AlertingProfilesListDto) SetSlackConfigurationIdNil(b bool)`

 SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil

### UnsetSlackConfigurationId
`func (o *AlertingProfilesListDto) UnsetSlackConfigurationId()`

UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
### GetSlackConfigurationName

`func (o *AlertingProfilesListDto) GetSlackConfigurationName() string`

GetSlackConfigurationName returns the SlackConfigurationName field if non-nil, zero value otherwise.

### GetSlackConfigurationNameOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationNameOk() (*string, bool)`

GetSlackConfigurationNameOk returns a tuple with the SlackConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationName

`func (o *AlertingProfilesListDto) SetSlackConfigurationName(v string)`

SetSlackConfigurationName sets SlackConfigurationName field to given value.

### HasSlackConfigurationName

`func (o *AlertingProfilesListDto) HasSlackConfigurationName() bool`

HasSlackConfigurationName returns a boolean if a field has been set.

### SetSlackConfigurationNameNil

`func (o *AlertingProfilesListDto) SetSlackConfigurationNameNil(b bool)`

 SetSlackConfigurationNameNil sets the value for SlackConfigurationName to be an explicit nil

### UnsetSlackConfigurationName
`func (o *AlertingProfilesListDto) UnsetSlackConfigurationName()`

UnsetSlackConfigurationName ensures that no value is present for SlackConfigurationName, not even an explicit nil
### GetIsLocked

`func (o *AlertingProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AlertingProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AlertingProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AlertingProfilesListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetEmails

`func (o *AlertingProfilesListDto) GetEmails() []AlertingEmailDto`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AlertingProfilesListDto) GetEmailsOk() (*[]AlertingEmailDto, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AlertingProfilesListDto) SetEmails(v []AlertingEmailDto)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AlertingProfilesListDto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *AlertingProfilesListDto) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *AlertingProfilesListDto) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetWebhooks

`func (o *AlertingProfilesListDto) GetWebhooks() []AlertingWebhookDto`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AlertingProfilesListDto) GetWebhooksOk() (*[]AlertingWebhookDto, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AlertingProfilesListDto) SetWebhooks(v []AlertingWebhookDto)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *AlertingProfilesListDto) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooksNil

`func (o *AlertingProfilesListDto) SetWebhooksNil(b bool)`

 SetWebhooksNil sets the value for Webhooks to be an explicit nil

### UnsetWebhooks
`func (o *AlertingProfilesListDto) UnsetWebhooks()`

UnsetWebhooks ensures that no value is present for Webhooks, not even an explicit nil
### GetProjects

`func (o *AlertingProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AlertingProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AlertingProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *AlertingProfilesListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *AlertingProfilesListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *AlertingProfilesListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetCreatedBy

`func (o *AlertingProfilesListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertingProfilesListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertingProfilesListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AlertingProfilesListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AlertingProfilesListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AlertingProfilesListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *AlertingProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AlertingProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AlertingProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AlertingProfilesListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *AlertingProfilesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AlertingProfilesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *AlertingProfilesListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AlertingProfilesListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AlertingProfilesListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *AlertingProfilesListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *AlertingProfilesListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *AlertingProfilesListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetReminder

`func (o *AlertingProfilesListDto) GetReminder() AlertingReminder`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *AlertingProfilesListDto) GetReminderOk() (*AlertingReminder, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *AlertingProfilesListDto) SetReminder(v AlertingReminder)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *AlertingProfilesListDto) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertingProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertingProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertingProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertingProfilesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AlertingProfilesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AlertingProfilesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


