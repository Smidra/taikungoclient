# CreateTanzuCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**TanzuContinent** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateTanzuCommand

`func NewCreateTanzuCommand() *CreateTanzuCommand`

NewCreateTanzuCommand instantiates a new CreateTanzuCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTanzuCommandWithDefaults

`func NewCreateTanzuCommandWithDefaults() *CreateTanzuCommand`

NewCreateTanzuCommandWithDefaults instantiates a new CreateTanzuCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *CreateTanzuCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateTanzuCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateTanzuCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateTanzuCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateTanzuCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateTanzuCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetName

`func (o *CreateTanzuCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTanzuCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTanzuCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTanzuCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTanzuCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTanzuCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsername

`func (o *CreateTanzuCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateTanzuCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateTanzuCommand) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateTanzuCommand) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateTanzuCommand) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateTanzuCommand) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetUrl

`func (o *CreateTanzuCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateTanzuCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateTanzuCommand) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateTanzuCommand) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreateTanzuCommand) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateTanzuCommand) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPassword

`func (o *CreateTanzuCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateTanzuCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateTanzuCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateTanzuCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateTanzuCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateTanzuCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetVolumeType

`func (o *CreateTanzuCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateTanzuCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateTanzuCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreateTanzuCommand) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CreateTanzuCommand) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CreateTanzuCommand) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetNamespace

`func (o *CreateTanzuCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateTanzuCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateTanzuCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateTanzuCommand) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *CreateTanzuCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *CreateTanzuCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTanzuContinent

`func (o *CreateTanzuCommand) GetTanzuContinent() string`

GetTanzuContinent returns the TanzuContinent field if non-nil, zero value otherwise.

### GetTanzuContinentOk

`func (o *CreateTanzuCommand) GetTanzuContinentOk() (*string, bool)`

GetTanzuContinentOk returns a tuple with the TanzuContinent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzuContinent

`func (o *CreateTanzuCommand) SetTanzuContinent(v string)`

SetTanzuContinent sets TanzuContinent field to given value.

### HasTanzuContinent

`func (o *CreateTanzuCommand) HasTanzuContinent() bool`

HasTanzuContinent returns a boolean if a field has been set.

### SetTanzuContinentNil

`func (o *CreateTanzuCommand) SetTanzuContinentNil(b bool)`

 SetTanzuContinentNil sets the value for TanzuContinent to be an explicit nil

### UnsetTanzuContinent
`func (o *CreateTanzuCommand) UnsetTanzuContinent()`

UnsetTanzuContinent ensures that no value is present for TanzuContinent, not even an explicit nil
### GetPort

`func (o *CreateTanzuCommand) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateTanzuCommand) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateTanzuCommand) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateTanzuCommand) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CreateTanzuCommand) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CreateTanzuCommand) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


