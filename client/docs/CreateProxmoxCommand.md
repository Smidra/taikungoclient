# CreateProxmoxCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**TokenId** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**TokenSecret** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 
**VmTemplateName** | Pointer to **NullableString** |  | [optional] 
**Continent** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Hypervisors** | Pointer to **[]string** |  | [optional] 
**PublicNetwork** | Pointer to [**CreateProxmoxNetworkDto**](CreateProxmoxNetworkDto.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**CreateProxmoxNetworkDto**](CreateProxmoxNetworkDto.md) |  | [optional] 

## Methods

### NewCreateProxmoxCommand

`func NewCreateProxmoxCommand() *CreateProxmoxCommand`

NewCreateProxmoxCommand instantiates a new CreateProxmoxCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxmoxCommandWithDefaults

`func NewCreateProxmoxCommandWithDefaults() *CreateProxmoxCommand`

NewCreateProxmoxCommandWithDefaults instantiates a new CreateProxmoxCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProxmoxCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProxmoxCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProxmoxCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProxmoxCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateProxmoxCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateProxmoxCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTokenId

`func (o *CreateProxmoxCommand) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateProxmoxCommand) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateProxmoxCommand) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CreateProxmoxCommand) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *CreateProxmoxCommand) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *CreateProxmoxCommand) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetUrl

`func (o *CreateProxmoxCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateProxmoxCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateProxmoxCommand) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateProxmoxCommand) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreateProxmoxCommand) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateProxmoxCommand) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTokenSecret

`func (o *CreateProxmoxCommand) GetTokenSecret() string`

GetTokenSecret returns the TokenSecret field if non-nil, zero value otherwise.

### GetTokenSecretOk

`func (o *CreateProxmoxCommand) GetTokenSecretOk() (*string, bool)`

GetTokenSecretOk returns a tuple with the TokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSecret

`func (o *CreateProxmoxCommand) SetTokenSecret(v string)`

SetTokenSecret sets TokenSecret field to given value.

### HasTokenSecret

`func (o *CreateProxmoxCommand) HasTokenSecret() bool`

HasTokenSecret returns a boolean if a field has been set.

### SetTokenSecretNil

`func (o *CreateProxmoxCommand) SetTokenSecretNil(b bool)`

 SetTokenSecretNil sets the value for TokenSecret to be an explicit nil

### UnsetTokenSecret
`func (o *CreateProxmoxCommand) UnsetTokenSecret()`

UnsetTokenSecret ensures that no value is present for TokenSecret, not even an explicit nil
### GetStorage

`func (o *CreateProxmoxCommand) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateProxmoxCommand) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateProxmoxCommand) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateProxmoxCommand) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CreateProxmoxCommand) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CreateProxmoxCommand) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetVmTemplateName

`func (o *CreateProxmoxCommand) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *CreateProxmoxCommand) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *CreateProxmoxCommand) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.

### HasVmTemplateName

`func (o *CreateProxmoxCommand) HasVmTemplateName() bool`

HasVmTemplateName returns a boolean if a field has been set.

### SetVmTemplateNameNil

`func (o *CreateProxmoxCommand) SetVmTemplateNameNil(b bool)`

 SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil

### UnsetVmTemplateName
`func (o *CreateProxmoxCommand) UnsetVmTemplateName()`

UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
### GetContinent

`func (o *CreateProxmoxCommand) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *CreateProxmoxCommand) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *CreateProxmoxCommand) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *CreateProxmoxCommand) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### SetContinentNil

`func (o *CreateProxmoxCommand) SetContinentNil(b bool)`

 SetContinentNil sets the value for Continent to be an explicit nil

### UnsetContinent
`func (o *CreateProxmoxCommand) UnsetContinent()`

UnsetContinent ensures that no value is present for Continent, not even an explicit nil
### GetOrganizationId

`func (o *CreateProxmoxCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateProxmoxCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateProxmoxCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateProxmoxCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateProxmoxCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateProxmoxCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetHypervisors

`func (o *CreateProxmoxCommand) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *CreateProxmoxCommand) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *CreateProxmoxCommand) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *CreateProxmoxCommand) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.

### SetHypervisorsNil

`func (o *CreateProxmoxCommand) SetHypervisorsNil(b bool)`

 SetHypervisorsNil sets the value for Hypervisors to be an explicit nil

### UnsetHypervisors
`func (o *CreateProxmoxCommand) UnsetHypervisors()`

UnsetHypervisors ensures that no value is present for Hypervisors, not even an explicit nil
### GetPublicNetwork

`func (o *CreateProxmoxCommand) GetPublicNetwork() CreateProxmoxNetworkDto`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *CreateProxmoxCommand) GetPublicNetworkOk() (*CreateProxmoxNetworkDto, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *CreateProxmoxCommand) SetPublicNetwork(v CreateProxmoxNetworkDto)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *CreateProxmoxCommand) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *CreateProxmoxCommand) GetPrivateNetwork() CreateProxmoxNetworkDto`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *CreateProxmoxCommand) GetPrivateNetworkOk() (*CreateProxmoxNetworkDto, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *CreateProxmoxCommand) SetPrivateNetwork(v CreateProxmoxNetworkDto)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *CreateProxmoxCommand) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


