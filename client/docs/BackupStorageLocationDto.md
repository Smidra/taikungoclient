# BackupStorageLocationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**LastValidated** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**AccessMode** | Pointer to **NullableString** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 
**BackupCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupStorageLocationDto

`func NewBackupStorageLocationDto() *BackupStorageLocationDto`

NewBackupStorageLocationDto instantiates a new BackupStorageLocationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageLocationDtoWithDefaults

`func NewBackupStorageLocationDtoWithDefaults() *BackupStorageLocationDto`

NewBackupStorageLocationDtoWithDefaults instantiates a new BackupStorageLocationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *BackupStorageLocationDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *BackupStorageLocationDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *BackupStorageLocationDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *BackupStorageLocationDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *BackupStorageLocationDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *BackupStorageLocationDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetProvider

`func (o *BackupStorageLocationDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BackupStorageLocationDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BackupStorageLocationDto) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BackupStorageLocationDto) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *BackupStorageLocationDto) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *BackupStorageLocationDto) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetNamespace

`func (o *BackupStorageLocationDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BackupStorageLocationDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BackupStorageLocationDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BackupStorageLocationDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *BackupStorageLocationDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *BackupStorageLocationDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetLastValidated

`func (o *BackupStorageLocationDto) GetLastValidated() time.Time`

GetLastValidated returns the LastValidated field if non-nil, zero value otherwise.

### GetLastValidatedOk

`func (o *BackupStorageLocationDto) GetLastValidatedOk() (*time.Time, bool)`

GetLastValidatedOk returns a tuple with the LastValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidated

`func (o *BackupStorageLocationDto) SetLastValidated(v time.Time)`

SetLastValidated sets LastValidated field to given value.

### HasLastValidated

`func (o *BackupStorageLocationDto) HasLastValidated() bool`

HasLastValidated returns a boolean if a field has been set.

### SetLastValidatedNil

`func (o *BackupStorageLocationDto) SetLastValidatedNil(b bool)`

 SetLastValidatedNil sets the value for LastValidated to be an explicit nil

### UnsetLastValidated
`func (o *BackupStorageLocationDto) UnsetLastValidated()`

UnsetLastValidated ensures that no value is present for LastValidated, not even an explicit nil
### GetCreatedAt

`func (o *BackupStorageLocationDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupStorageLocationDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupStorageLocationDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BackupStorageLocationDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BackupStorageLocationDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BackupStorageLocationDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetAccessMode

`func (o *BackupStorageLocationDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *BackupStorageLocationDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *BackupStorageLocationDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *BackupStorageLocationDto) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### SetAccessModeNil

`func (o *BackupStorageLocationDto) SetAccessModeNil(b bool)`

 SetAccessModeNil sets the value for AccessMode to be an explicit nil

### UnsetAccessMode
`func (o *BackupStorageLocationDto) UnsetAccessMode()`

UnsetAccessMode ensures that no value is present for AccessMode, not even an explicit nil
### GetPhase

`func (o *BackupStorageLocationDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *BackupStorageLocationDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *BackupStorageLocationDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *BackupStorageLocationDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *BackupStorageLocationDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *BackupStorageLocationDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetBackupCredentialId

`func (o *BackupStorageLocationDto) GetBackupCredentialId() int32`

GetBackupCredentialId returns the BackupCredentialId field if non-nil, zero value otherwise.

### GetBackupCredentialIdOk

`func (o *BackupStorageLocationDto) GetBackupCredentialIdOk() (*int32, bool)`

GetBackupCredentialIdOk returns a tuple with the BackupCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCredentialId

`func (o *BackupStorageLocationDto) SetBackupCredentialId(v int32)`

SetBackupCredentialId sets BackupCredentialId field to given value.

### HasBackupCredentialId

`func (o *BackupStorageLocationDto) HasBackupCredentialId() bool`

HasBackupCredentialId returns a boolean if a field has been set.

### SetBackupCredentialIdNil

`func (o *BackupStorageLocationDto) SetBackupCredentialIdNil(b bool)`

 SetBackupCredentialIdNil sets the value for BackupCredentialId to be an explicit nil

### UnsetBackupCredentialId
`func (o *BackupStorageLocationDto) UnsetBackupCredentialId()`

UnsetBackupCredentialId ensures that no value is present for BackupCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


