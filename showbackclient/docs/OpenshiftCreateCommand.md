# OpenshiftCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**KubeConfig** | Pointer to **NullableString** |  | [optional] 
**PullSecret** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**BaseDomain** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewOpenshiftCreateCommand

`func NewOpenshiftCreateCommand() *OpenshiftCreateCommand`

NewOpenshiftCreateCommand instantiates a new OpenshiftCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenshiftCreateCommandWithDefaults

`func NewOpenshiftCreateCommandWithDefaults() *OpenshiftCreateCommand`

NewOpenshiftCreateCommandWithDefaults instantiates a new OpenshiftCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenshiftCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenshiftCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenshiftCreateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenshiftCreateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenshiftCreateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenshiftCreateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKubeConfig

`func (o *OpenshiftCreateCommand) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *OpenshiftCreateCommand) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *OpenshiftCreateCommand) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *OpenshiftCreateCommand) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *OpenshiftCreateCommand) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *OpenshiftCreateCommand) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetPullSecret

`func (o *OpenshiftCreateCommand) GetPullSecret() string`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *OpenshiftCreateCommand) GetPullSecretOk() (*string, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *OpenshiftCreateCommand) SetPullSecret(v string)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *OpenshiftCreateCommand) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.

### SetPullSecretNil

`func (o *OpenshiftCreateCommand) SetPullSecretNil(b bool)`

 SetPullSecretNil sets the value for PullSecret to be an explicit nil

### UnsetPullSecret
`func (o *OpenshiftCreateCommand) UnsetPullSecret()`

UnsetPullSecret ensures that no value is present for PullSecret, not even an explicit nil
### GetStorageClass

`func (o *OpenshiftCreateCommand) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *OpenshiftCreateCommand) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *OpenshiftCreateCommand) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *OpenshiftCreateCommand) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *OpenshiftCreateCommand) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *OpenshiftCreateCommand) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetBaseDomain

`func (o *OpenshiftCreateCommand) GetBaseDomain() string`

GetBaseDomain returns the BaseDomain field if non-nil, zero value otherwise.

### GetBaseDomainOk

`func (o *OpenshiftCreateCommand) GetBaseDomainOk() (*string, bool)`

GetBaseDomainOk returns a tuple with the BaseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDomain

`func (o *OpenshiftCreateCommand) SetBaseDomain(v string)`

SetBaseDomain sets BaseDomain field to given value.

### HasBaseDomain

`func (o *OpenshiftCreateCommand) HasBaseDomain() bool`

HasBaseDomain returns a boolean if a field has been set.

### SetBaseDomainNil

`func (o *OpenshiftCreateCommand) SetBaseDomainNil(b bool)`

 SetBaseDomainNil sets the value for BaseDomain to be an explicit nil

### UnsetBaseDomain
`func (o *OpenshiftCreateCommand) UnsetBaseDomain()`

UnsetBaseDomain ensures that no value is present for BaseDomain, not even an explicit nil
### GetOrganizationId

`func (o *OpenshiftCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpenshiftCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpenshiftCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OpenshiftCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *OpenshiftCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *OpenshiftCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


