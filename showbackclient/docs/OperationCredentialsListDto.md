# OperationCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PrometheusUsername** | Pointer to **NullableString** |  | [optional] 
**PrometheusPassword** | Pointer to **NullableString** |  | [optional] 
**PrometheusUrl** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]SimplePrometheusEntity**](SimplePrometheusEntity.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperationCredentialsListDto

`func NewOperationCredentialsListDto() *OperationCredentialsListDto`

NewOperationCredentialsListDto instantiates a new OperationCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCredentialsListDtoWithDefaults

`func NewOperationCredentialsListDtoWithDefaults() *OperationCredentialsListDto`

NewOperationCredentialsListDtoWithDefaults instantiates a new OperationCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperationCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperationCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperationCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OperationCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OperationCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperationCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OperationCredentialsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OperationCredentialsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrometheusUsername

`func (o *OperationCredentialsListDto) GetPrometheusUsername() string`

GetPrometheusUsername returns the PrometheusUsername field if non-nil, zero value otherwise.

### GetPrometheusUsernameOk

`func (o *OperationCredentialsListDto) GetPrometheusUsernameOk() (*string, bool)`

GetPrometheusUsernameOk returns a tuple with the PrometheusUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUsername

`func (o *OperationCredentialsListDto) SetPrometheusUsername(v string)`

SetPrometheusUsername sets PrometheusUsername field to given value.

### HasPrometheusUsername

`func (o *OperationCredentialsListDto) HasPrometheusUsername() bool`

HasPrometheusUsername returns a boolean if a field has been set.

### SetPrometheusUsernameNil

`func (o *OperationCredentialsListDto) SetPrometheusUsernameNil(b bool)`

 SetPrometheusUsernameNil sets the value for PrometheusUsername to be an explicit nil

### UnsetPrometheusUsername
`func (o *OperationCredentialsListDto) UnsetPrometheusUsername()`

UnsetPrometheusUsername ensures that no value is present for PrometheusUsername, not even an explicit nil
### GetPrometheusPassword

`func (o *OperationCredentialsListDto) GetPrometheusPassword() string`

GetPrometheusPassword returns the PrometheusPassword field if non-nil, zero value otherwise.

### GetPrometheusPasswordOk

`func (o *OperationCredentialsListDto) GetPrometheusPasswordOk() (*string, bool)`

GetPrometheusPasswordOk returns a tuple with the PrometheusPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusPassword

`func (o *OperationCredentialsListDto) SetPrometheusPassword(v string)`

SetPrometheusPassword sets PrometheusPassword field to given value.

### HasPrometheusPassword

`func (o *OperationCredentialsListDto) HasPrometheusPassword() bool`

HasPrometheusPassword returns a boolean if a field has been set.

### SetPrometheusPasswordNil

`func (o *OperationCredentialsListDto) SetPrometheusPasswordNil(b bool)`

 SetPrometheusPasswordNil sets the value for PrometheusPassword to be an explicit nil

### UnsetPrometheusPassword
`func (o *OperationCredentialsListDto) UnsetPrometheusPassword()`

UnsetPrometheusPassword ensures that no value is present for PrometheusPassword, not even an explicit nil
### GetPrometheusUrl

`func (o *OperationCredentialsListDto) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *OperationCredentialsListDto) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *OperationCredentialsListDto) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.

### HasPrometheusUrl

`func (o *OperationCredentialsListDto) HasPrometheusUrl() bool`

HasPrometheusUrl returns a boolean if a field has been set.

### SetPrometheusUrlNil

`func (o *OperationCredentialsListDto) SetPrometheusUrlNil(b bool)`

 SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil

### UnsetPrometheusUrl
`func (o *OperationCredentialsListDto) UnsetPrometheusUrl()`

UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
### GetOrganizationId

`func (o *OperationCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OperationCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OperationCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OperationCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *OperationCredentialsListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *OperationCredentialsListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *OperationCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OperationCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OperationCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OperationCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OperationCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OperationCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetIsLocked

`func (o *OperationCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OperationCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OperationCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *OperationCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetRules

`func (o *OperationCredentialsListDto) GetRules() []SimplePrometheusEntity`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OperationCredentialsListDto) GetRulesOk() (*[]SimplePrometheusEntity, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OperationCredentialsListDto) SetRules(v []SimplePrometheusEntity)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OperationCredentialsListDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *OperationCredentialsListDto) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *OperationCredentialsListDto) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetCreatedBy

`func (o *OperationCredentialsListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OperationCredentialsListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OperationCredentialsListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OperationCredentialsListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OperationCredentialsListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OperationCredentialsListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *OperationCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *OperationCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *OperationCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *OperationCredentialsListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *OperationCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *OperationCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *OperationCredentialsListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *OperationCredentialsListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *OperationCredentialsListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *OperationCredentialsListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *OperationCredentialsListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *OperationCredentialsListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetIsDefault

`func (o *OperationCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OperationCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OperationCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OperationCredentialsListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


