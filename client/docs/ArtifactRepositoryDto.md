# ArtifactRepositoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**VerifiedPublisher** | Pointer to **bool** |  | [optional] 
**Official** | Pointer to **bool** |  | [optional] 
**IsBound** | Pointer to **bool** |  | [optional] 
**HasCatalogApp** | Pointer to **bool** |  | [optional] 
**TrueUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewArtifactRepositoryDto

`func NewArtifactRepositoryDto() *ArtifactRepositoryDto`

NewArtifactRepositoryDto instantiates a new ArtifactRepositoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryDtoWithDefaults

`func NewArtifactRepositoryDtoWithDefaults() *ArtifactRepositoryDto`

NewArtifactRepositoryDtoWithDefaults instantiates a new ArtifactRepositoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryId

`func (o *ArtifactRepositoryDto) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ArtifactRepositoryDto) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ArtifactRepositoryDto) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ArtifactRepositoryDto) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### SetRepositoryIdNil

`func (o *ArtifactRepositoryDto) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *ArtifactRepositoryDto) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetName

`func (o *ArtifactRepositoryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactRepositoryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactRepositoryDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactRepositoryDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ArtifactRepositoryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ArtifactRepositoryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *ArtifactRepositoryDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArtifactRepositoryDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArtifactRepositoryDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ArtifactRepositoryDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ArtifactRepositoryDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ArtifactRepositoryDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUrl

`func (o *ArtifactRepositoryDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArtifactRepositoryDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArtifactRepositoryDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ArtifactRepositoryDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ArtifactRepositoryDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ArtifactRepositoryDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetOrganizationName

`func (o *ArtifactRepositoryDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ArtifactRepositoryDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ArtifactRepositoryDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ArtifactRepositoryDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ArtifactRepositoryDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ArtifactRepositoryDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetDisabled

`func (o *ArtifactRepositoryDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ArtifactRepositoryDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ArtifactRepositoryDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ArtifactRepositoryDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVerifiedPublisher

`func (o *ArtifactRepositoryDto) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *ArtifactRepositoryDto) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *ArtifactRepositoryDto) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *ArtifactRepositoryDto) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### GetOfficial

`func (o *ArtifactRepositoryDto) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *ArtifactRepositoryDto) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *ArtifactRepositoryDto) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *ArtifactRepositoryDto) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetIsBound

`func (o *ArtifactRepositoryDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *ArtifactRepositoryDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *ArtifactRepositoryDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *ArtifactRepositoryDto) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.

### GetHasCatalogApp

`func (o *ArtifactRepositoryDto) GetHasCatalogApp() bool`

GetHasCatalogApp returns the HasCatalogApp field if non-nil, zero value otherwise.

### GetHasCatalogAppOk

`func (o *ArtifactRepositoryDto) GetHasCatalogAppOk() (*bool, bool)`

GetHasCatalogAppOk returns a tuple with the HasCatalogApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCatalogApp

`func (o *ArtifactRepositoryDto) SetHasCatalogApp(v bool)`

SetHasCatalogApp sets HasCatalogApp field to given value.

### HasHasCatalogApp

`func (o *ArtifactRepositoryDto) HasHasCatalogApp() bool`

HasHasCatalogApp returns a boolean if a field has been set.

### GetTrueUrl

`func (o *ArtifactRepositoryDto) GetTrueUrl() string`

GetTrueUrl returns the TrueUrl field if non-nil, zero value otherwise.

### GetTrueUrlOk

`func (o *ArtifactRepositoryDto) GetTrueUrlOk() (*string, bool)`

GetTrueUrlOk returns a tuple with the TrueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueUrl

`func (o *ArtifactRepositoryDto) SetTrueUrl(v string)`

SetTrueUrl sets TrueUrl field to given value.

### HasTrueUrl

`func (o *ArtifactRepositoryDto) HasTrueUrl() bool`

HasTrueUrl returns a boolean if a field has been set.

### SetTrueUrlNil

`func (o *ArtifactRepositoryDto) SetTrueUrlNil(b bool)`

 SetTrueUrlNil sets the value for TrueUrl to be an explicit nil

### UnsetTrueUrl
`func (o *ArtifactRepositoryDto) UnsetTrueUrl()`

UnsetTrueUrl ensures that no value is present for TrueUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


