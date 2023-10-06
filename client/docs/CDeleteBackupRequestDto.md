# CDeleteBackupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**BackupName** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCDeleteBackupRequestDto

`func NewCDeleteBackupRequestDto() *CDeleteBackupRequestDto`

NewCDeleteBackupRequestDto instantiates a new CDeleteBackupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDeleteBackupRequestDtoWithDefaults

`func NewCDeleteBackupRequestDtoWithDefaults() *CDeleteBackupRequestDto`

NewCDeleteBackupRequestDtoWithDefaults instantiates a new CDeleteBackupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *CDeleteBackupRequestDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CDeleteBackupRequestDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CDeleteBackupRequestDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *CDeleteBackupRequestDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *CDeleteBackupRequestDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *CDeleteBackupRequestDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetCreatedAt

`func (o *CDeleteBackupRequestDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CDeleteBackupRequestDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CDeleteBackupRequestDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CDeleteBackupRequestDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CDeleteBackupRequestDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CDeleteBackupRequestDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetBackupName

`func (o *CDeleteBackupRequestDto) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *CDeleteBackupRequestDto) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *CDeleteBackupRequestDto) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *CDeleteBackupRequestDto) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *CDeleteBackupRequestDto) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *CDeleteBackupRequestDto) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetNamespace

`func (o *CDeleteBackupRequestDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CDeleteBackupRequestDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CDeleteBackupRequestDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CDeleteBackupRequestDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *CDeleteBackupRequestDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *CDeleteBackupRequestDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetPhase

`func (o *CDeleteBackupRequestDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CDeleteBackupRequestDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CDeleteBackupRequestDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CDeleteBackupRequestDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *CDeleteBackupRequestDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CDeleteBackupRequestDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


