# AdminProjectsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**KubesprayCurrentVersion** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ServersCount** | Pointer to **int32** |  | [optional] 
**Tcu** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CloudType** | Pointer to [**CloudType**](CloudType.md) |  | [optional] 

## Methods

### NewAdminProjectsResponseData

`func NewAdminProjectsResponseData() *AdminProjectsResponseData`

NewAdminProjectsResponseData instantiates a new AdminProjectsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminProjectsResponseDataWithDefaults

`func NewAdminProjectsResponseDataWithDefaults() *AdminProjectsResponseData`

NewAdminProjectsResponseDataWithDefaults instantiates a new AdminProjectsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminProjectsResponseData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminProjectsResponseData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminProjectsResponseData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminProjectsResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminProjectsResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminProjectsResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminProjectsResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminProjectsResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminProjectsResponseData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminProjectsResponseData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationName

`func (o *AdminProjectsResponseData) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AdminProjectsResponseData) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AdminProjectsResponseData) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AdminProjectsResponseData) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AdminProjectsResponseData) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AdminProjectsResponseData) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetIsLocked

`func (o *AdminProjectsResponseData) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AdminProjectsResponseData) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AdminProjectsResponseData) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AdminProjectsResponseData) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetKubernetesCurrentVersion

`func (o *AdminProjectsResponseData) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *AdminProjectsResponseData) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *AdminProjectsResponseData) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *AdminProjectsResponseData) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *AdminProjectsResponseData) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *AdminProjectsResponseData) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil
### GetKubesprayCurrentVersion

`func (o *AdminProjectsResponseData) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *AdminProjectsResponseData) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *AdminProjectsResponseData) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.

### HasKubesprayCurrentVersion

`func (o *AdminProjectsResponseData) HasKubesprayCurrentVersion() bool`

HasKubesprayCurrentVersion returns a boolean if a field has been set.

### SetKubesprayCurrentVersionNil

`func (o *AdminProjectsResponseData) SetKubesprayCurrentVersionNil(b bool)`

 SetKubesprayCurrentVersionNil sets the value for KubesprayCurrentVersion to be an explicit nil

### UnsetKubesprayCurrentVersion
`func (o *AdminProjectsResponseData) UnsetKubesprayCurrentVersion()`

UnsetKubesprayCurrentVersion ensures that no value is present for KubesprayCurrentVersion, not even an explicit nil
### GetStatus

`func (o *AdminProjectsResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminProjectsResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminProjectsResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminProjectsResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AdminProjectsResponseData) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AdminProjectsResponseData) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetServersCount

`func (o *AdminProjectsResponseData) GetServersCount() int32`

GetServersCount returns the ServersCount field if non-nil, zero value otherwise.

### GetServersCountOk

`func (o *AdminProjectsResponseData) GetServersCountOk() (*int32, bool)`

GetServersCountOk returns a tuple with the ServersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersCount

`func (o *AdminProjectsResponseData) SetServersCount(v int32)`

SetServersCount sets ServersCount field to given value.

### HasServersCount

`func (o *AdminProjectsResponseData) HasServersCount() bool`

HasServersCount returns a boolean if a field has been set.

### GetTcu

`func (o *AdminProjectsResponseData) GetTcu() int32`

GetTcu returns the Tcu field if non-nil, zero value otherwise.

### GetTcuOk

`func (o *AdminProjectsResponseData) GetTcuOk() (*int32, bool)`

GetTcuOk returns a tuple with the Tcu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcu

`func (o *AdminProjectsResponseData) SetTcu(v int32)`

SetTcu sets Tcu field to given value.

### HasTcu

`func (o *AdminProjectsResponseData) HasTcu() bool`

HasTcu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminProjectsResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminProjectsResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminProjectsResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminProjectsResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AdminProjectsResponseData) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AdminProjectsResponseData) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCloudType

`func (o *AdminProjectsResponseData) GetCloudType() CloudType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AdminProjectsResponseData) GetCloudTypeOk() (*CloudType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AdminProjectsResponseData) SetCloudType(v CloudType)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *AdminProjectsResponseData) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


