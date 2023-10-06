# EnableBackupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**S3CredentialId** | Pointer to **int32** |  | [optional] 

## Methods

### NewEnableBackupCommand

`func NewEnableBackupCommand() *EnableBackupCommand`

NewEnableBackupCommand instantiates a new EnableBackupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableBackupCommandWithDefaults

`func NewEnableBackupCommandWithDefaults() *EnableBackupCommand`

NewEnableBackupCommandWithDefaults instantiates a new EnableBackupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *EnableBackupCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnableBackupCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnableBackupCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *EnableBackupCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetS3CredentialId

`func (o *EnableBackupCommand) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *EnableBackupCommand) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *EnableBackupCommand) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.

### HasS3CredentialId

`func (o *EnableBackupCommand) HasS3CredentialId() bool`

HasS3CredentialId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


