# CRestoreDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**BackupName** | Pointer to **NullableString** |  | [optional] 
**ScheduleName** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**ExcludeNamespaces** | Pointer to **[]string** |  | [optional] 
**IncludeNamespaces** | Pointer to **[]string** |  | [optional] 
**CompletionDateTime** | Pointer to **time.Time** |  | [optional] 
**StartTimeStamp** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**Warnings** | Pointer to **int64** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCRestoreDto

`func NewCRestoreDto() *CRestoreDto`

NewCRestoreDto instantiates a new CRestoreDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCRestoreDtoWithDefaults

`func NewCRestoreDtoWithDefaults() *CRestoreDto`

NewCRestoreDtoWithDefaults instantiates a new CRestoreDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *CRestoreDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CRestoreDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CRestoreDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *CRestoreDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *CRestoreDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *CRestoreDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetBackupName

`func (o *CRestoreDto) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *CRestoreDto) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *CRestoreDto) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *CRestoreDto) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *CRestoreDto) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *CRestoreDto) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetScheduleName

`func (o *CRestoreDto) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *CRestoreDto) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *CRestoreDto) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *CRestoreDto) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### SetScheduleNameNil

`func (o *CRestoreDto) SetScheduleNameNil(b bool)`

 SetScheduleNameNil sets the value for ScheduleName to be an explicit nil

### UnsetScheduleName
`func (o *CRestoreDto) UnsetScheduleName()`

UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
### GetNamespace

`func (o *CRestoreDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CRestoreDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CRestoreDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CRestoreDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *CRestoreDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *CRestoreDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetExcludeNamespaces

`func (o *CRestoreDto) GetExcludeNamespaces() []string`

GetExcludeNamespaces returns the ExcludeNamespaces field if non-nil, zero value otherwise.

### GetExcludeNamespacesOk

`func (o *CRestoreDto) GetExcludeNamespacesOk() (*[]string, bool)`

GetExcludeNamespacesOk returns a tuple with the ExcludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeNamespaces

`func (o *CRestoreDto) SetExcludeNamespaces(v []string)`

SetExcludeNamespaces sets ExcludeNamespaces field to given value.

### HasExcludeNamespaces

`func (o *CRestoreDto) HasExcludeNamespaces() bool`

HasExcludeNamespaces returns a boolean if a field has been set.

### SetExcludeNamespacesNil

`func (o *CRestoreDto) SetExcludeNamespacesNil(b bool)`

 SetExcludeNamespacesNil sets the value for ExcludeNamespaces to be an explicit nil

### UnsetExcludeNamespaces
`func (o *CRestoreDto) UnsetExcludeNamespaces()`

UnsetExcludeNamespaces ensures that no value is present for ExcludeNamespaces, not even an explicit nil
### GetIncludeNamespaces

`func (o *CRestoreDto) GetIncludeNamespaces() []string`

GetIncludeNamespaces returns the IncludeNamespaces field if non-nil, zero value otherwise.

### GetIncludeNamespacesOk

`func (o *CRestoreDto) GetIncludeNamespacesOk() (*[]string, bool)`

GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNamespaces

`func (o *CRestoreDto) SetIncludeNamespaces(v []string)`

SetIncludeNamespaces sets IncludeNamespaces field to given value.

### HasIncludeNamespaces

`func (o *CRestoreDto) HasIncludeNamespaces() bool`

HasIncludeNamespaces returns a boolean if a field has been set.

### SetIncludeNamespacesNil

`func (o *CRestoreDto) SetIncludeNamespacesNil(b bool)`

 SetIncludeNamespacesNil sets the value for IncludeNamespaces to be an explicit nil

### UnsetIncludeNamespaces
`func (o *CRestoreDto) UnsetIncludeNamespaces()`

UnsetIncludeNamespaces ensures that no value is present for IncludeNamespaces, not even an explicit nil
### GetCompletionDateTime

`func (o *CRestoreDto) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *CRestoreDto) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *CRestoreDto) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *CRestoreDto) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetStartTimeStamp

`func (o *CRestoreDto) GetStartTimeStamp() time.Time`

GetStartTimeStamp returns the StartTimeStamp field if non-nil, zero value otherwise.

### GetStartTimeStampOk

`func (o *CRestoreDto) GetStartTimeStampOk() (*time.Time, bool)`

GetStartTimeStampOk returns a tuple with the StartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeStamp

`func (o *CRestoreDto) SetStartTimeStamp(v time.Time)`

SetStartTimeStamp sets StartTimeStamp field to given value.

### HasStartTimeStamp

`func (o *CRestoreDto) HasStartTimeStamp() bool`

HasStartTimeStamp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CRestoreDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CRestoreDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CRestoreDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CRestoreDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CRestoreDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CRestoreDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetWarnings

`func (o *CRestoreDto) GetWarnings() int64`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CRestoreDto) GetWarningsOk() (*int64, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CRestoreDto) SetWarnings(v int64)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CRestoreDto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetPhase

`func (o *CRestoreDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CRestoreDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CRestoreDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CRestoreDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *CRestoreDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CRestoreDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


