# KeycloakListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**RealmsName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**PartnerLogo** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeycloakListDto

`func NewKeycloakListDto() *KeycloakListDto`

NewKeycloakListDto instantiates a new KeycloakListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeycloakListDtoWithDefaults

`func NewKeycloakListDtoWithDefaults() *KeycloakListDto`

NewKeycloakListDtoWithDefaults instantiates a new KeycloakListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeycloakListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeycloakListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeycloakListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeycloakListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeycloakListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeycloakListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeycloakListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeycloakListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeycloakListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeycloakListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *KeycloakListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KeycloakListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KeycloakListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KeycloakListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *KeycloakListDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *KeycloakListDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetClientId

`func (o *KeycloakListDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KeycloakListDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KeycloakListDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KeycloakListDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *KeycloakListDto) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *KeycloakListDto) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *KeycloakListDto) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *KeycloakListDto) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *KeycloakListDto) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *KeycloakListDto) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *KeycloakListDto) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *KeycloakListDto) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetRealmsName

`func (o *KeycloakListDto) GetRealmsName() string`

GetRealmsName returns the RealmsName field if non-nil, zero value otherwise.

### GetRealmsNameOk

`func (o *KeycloakListDto) GetRealmsNameOk() (*string, bool)`

GetRealmsNameOk returns a tuple with the RealmsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmsName

`func (o *KeycloakListDto) SetRealmsName(v string)`

SetRealmsName sets RealmsName field to given value.

### HasRealmsName

`func (o *KeycloakListDto) HasRealmsName() bool`

HasRealmsName returns a boolean if a field has been set.

### SetRealmsNameNil

`func (o *KeycloakListDto) SetRealmsNameNil(b bool)`

 SetRealmsNameNil sets the value for RealmsName to be an explicit nil

### UnsetRealmsName
`func (o *KeycloakListDto) UnsetRealmsName()`

UnsetRealmsName ensures that no value is present for RealmsName, not even an explicit nil
### GetOrganizationId

`func (o *KeycloakListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KeycloakListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KeycloakListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *KeycloakListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *KeycloakListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *KeycloakListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *KeycloakListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *KeycloakListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *KeycloakListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *KeycloakListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetPartnerLogo

`func (o *KeycloakListDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *KeycloakListDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *KeycloakListDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.

### HasPartnerLogo

`func (o *KeycloakListDto) HasPartnerLogo() bool`

HasPartnerLogo returns a boolean if a field has been set.

### SetPartnerLogoNil

`func (o *KeycloakListDto) SetPartnerLogoNil(b bool)`

 SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil

### UnsetPartnerLogo
`func (o *KeycloakListDto) UnsetPartnerLogo()`

UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
### GetEnabled

`func (o *KeycloakListDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeycloakListDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeycloakListDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeycloakListDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


