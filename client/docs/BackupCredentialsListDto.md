# BackupCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**S3Name** | Pointer to **NullableString** |  | [optional] 
**S3AccessKeyId** | Pointer to **NullableString** |  | [optional] 
**S3Endpoint** | Pointer to **NullableString** |  | [optional] 
**S3Region** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewBackupCredentialsListDto

`func NewBackupCredentialsListDto() *BackupCredentialsListDto`

NewBackupCredentialsListDto instantiates a new BackupCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCredentialsListDtoWithDefaults

`func NewBackupCredentialsListDtoWithDefaults() *BackupCredentialsListDto`

NewBackupCredentialsListDtoWithDefaults instantiates a new BackupCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BackupCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetS3Name

`func (o *BackupCredentialsListDto) GetS3Name() string`

GetS3Name returns the S3Name field if non-nil, zero value otherwise.

### GetS3NameOk

`func (o *BackupCredentialsListDto) GetS3NameOk() (*string, bool)`

GetS3NameOk returns a tuple with the S3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Name

`func (o *BackupCredentialsListDto) SetS3Name(v string)`

SetS3Name sets S3Name field to given value.

### HasS3Name

`func (o *BackupCredentialsListDto) HasS3Name() bool`

HasS3Name returns a boolean if a field has been set.

### SetS3NameNil

`func (o *BackupCredentialsListDto) SetS3NameNil(b bool)`

 SetS3NameNil sets the value for S3Name to be an explicit nil

### UnsetS3Name
`func (o *BackupCredentialsListDto) UnsetS3Name()`

UnsetS3Name ensures that no value is present for S3Name, not even an explicit nil
### GetS3AccessKeyId

`func (o *BackupCredentialsListDto) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *BackupCredentialsListDto) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *BackupCredentialsListDto) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.

### HasS3AccessKeyId

`func (o *BackupCredentialsListDto) HasS3AccessKeyId() bool`

HasS3AccessKeyId returns a boolean if a field has been set.

### SetS3AccessKeyIdNil

`func (o *BackupCredentialsListDto) SetS3AccessKeyIdNil(b bool)`

 SetS3AccessKeyIdNil sets the value for S3AccessKeyId to be an explicit nil

### UnsetS3AccessKeyId
`func (o *BackupCredentialsListDto) UnsetS3AccessKeyId()`

UnsetS3AccessKeyId ensures that no value is present for S3AccessKeyId, not even an explicit nil
### GetS3Endpoint

`func (o *BackupCredentialsListDto) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *BackupCredentialsListDto) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *BackupCredentialsListDto) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.

### HasS3Endpoint

`func (o *BackupCredentialsListDto) HasS3Endpoint() bool`

HasS3Endpoint returns a boolean if a field has been set.

### SetS3EndpointNil

`func (o *BackupCredentialsListDto) SetS3EndpointNil(b bool)`

 SetS3EndpointNil sets the value for S3Endpoint to be an explicit nil

### UnsetS3Endpoint
`func (o *BackupCredentialsListDto) UnsetS3Endpoint()`

UnsetS3Endpoint ensures that no value is present for S3Endpoint, not even an explicit nil
### GetS3Region

`func (o *BackupCredentialsListDto) GetS3Region() string`

GetS3Region returns the S3Region field if non-nil, zero value otherwise.

### GetS3RegionOk

`func (o *BackupCredentialsListDto) GetS3RegionOk() (*string, bool)`

GetS3RegionOk returns a tuple with the S3Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Region

`func (o *BackupCredentialsListDto) SetS3Region(v string)`

SetS3Region sets S3Region field to given value.

### HasS3Region

`func (o *BackupCredentialsListDto) HasS3Region() bool`

HasS3Region returns a boolean if a field has been set.

### SetS3RegionNil

`func (o *BackupCredentialsListDto) SetS3RegionNil(b bool)`

 SetS3RegionNil sets the value for S3Region to be an explicit nil

### UnsetS3Region
`func (o *BackupCredentialsListDto) UnsetS3Region()`

UnsetS3Region ensures that no value is present for S3Region, not even an explicit nil
### GetOrganizationId

`func (o *BackupCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BackupCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BackupCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BackupCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *BackupCredentialsListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *BackupCredentialsListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *BackupCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *BackupCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *BackupCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *BackupCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *BackupCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *BackupCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetProjects

`func (o *BackupCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BackupCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BackupCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *BackupCredentialsListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *BackupCredentialsListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *BackupCredentialsListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetIsLocked

`func (o *BackupCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *BackupCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *BackupCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *BackupCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BackupCredentialsListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupCredentialsListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupCredentialsListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BackupCredentialsListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *BackupCredentialsListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *BackupCredentialsListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *BackupCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BackupCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BackupCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *BackupCredentialsListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *BackupCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *BackupCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *BackupCredentialsListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BackupCredentialsListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BackupCredentialsListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BackupCredentialsListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *BackupCredentialsListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *BackupCredentialsListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetCreatedAt

`func (o *BackupCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BackupCredentialsListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BackupCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BackupCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsDefault

`func (o *BackupCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BackupCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BackupCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BackupCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


