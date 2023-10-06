# OpenshiftCredentialForProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | Pointer to **NullableString** |  | [optional] 
**PullSecret** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**BaseDomain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenshiftCredentialForProjectDto

`func NewOpenshiftCredentialForProjectDto() *OpenshiftCredentialForProjectDto`

NewOpenshiftCredentialForProjectDto instantiates a new OpenshiftCredentialForProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenshiftCredentialForProjectDtoWithDefaults

`func NewOpenshiftCredentialForProjectDtoWithDefaults() *OpenshiftCredentialForProjectDto`

NewOpenshiftCredentialForProjectDtoWithDefaults instantiates a new OpenshiftCredentialForProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *OpenshiftCredentialForProjectDto) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *OpenshiftCredentialForProjectDto) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *OpenshiftCredentialForProjectDto) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *OpenshiftCredentialForProjectDto) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *OpenshiftCredentialForProjectDto) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *OpenshiftCredentialForProjectDto) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetPullSecret

`func (o *OpenshiftCredentialForProjectDto) GetPullSecret() string`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *OpenshiftCredentialForProjectDto) GetPullSecretOk() (*string, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *OpenshiftCredentialForProjectDto) SetPullSecret(v string)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *OpenshiftCredentialForProjectDto) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.

### SetPullSecretNil

`func (o *OpenshiftCredentialForProjectDto) SetPullSecretNil(b bool)`

 SetPullSecretNil sets the value for PullSecret to be an explicit nil

### UnsetPullSecret
`func (o *OpenshiftCredentialForProjectDto) UnsetPullSecret()`

UnsetPullSecret ensures that no value is present for PullSecret, not even an explicit nil
### GetStorageClass

`func (o *OpenshiftCredentialForProjectDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *OpenshiftCredentialForProjectDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *OpenshiftCredentialForProjectDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *OpenshiftCredentialForProjectDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *OpenshiftCredentialForProjectDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *OpenshiftCredentialForProjectDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetBaseDomain

`func (o *OpenshiftCredentialForProjectDto) GetBaseDomain() string`

GetBaseDomain returns the BaseDomain field if non-nil, zero value otherwise.

### GetBaseDomainOk

`func (o *OpenshiftCredentialForProjectDto) GetBaseDomainOk() (*string, bool)`

GetBaseDomainOk returns a tuple with the BaseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDomain

`func (o *OpenshiftCredentialForProjectDto) SetBaseDomain(v string)`

SetBaseDomain sets BaseDomain field to given value.

### HasBaseDomain

`func (o *OpenshiftCredentialForProjectDto) HasBaseDomain() bool`

HasBaseDomain returns a boolean if a field has been set.

### SetBaseDomainNil

`func (o *OpenshiftCredentialForProjectDto) SetBaseDomainNil(b bool)`

 SetBaseDomainNil sets the value for BaseDomain to be an explicit nil

### UnsetBaseDomain
`func (o *OpenshiftCredentialForProjectDto) UnsetBaseDomain()`

UnsetBaseDomain ensures that no value is present for BaseDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


