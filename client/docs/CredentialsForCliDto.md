# CredentialsForCliDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CloudCredentialName** | Pointer to **NullableString** |  | [optional] 
**CloudType** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**PartnerName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCredentialsForCliDto

`func NewCredentialsForCliDto() *CredentialsForCliDto`

NewCredentialsForCliDto instantiates a new CredentialsForCliDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsForCliDtoWithDefaults

`func NewCredentialsForCliDtoWithDefaults() *CredentialsForCliDto`

NewCredentialsForCliDtoWithDefaults instantiates a new CredentialsForCliDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialsForCliDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialsForCliDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialsForCliDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialsForCliDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *CredentialsForCliDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *CredentialsForCliDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *CredentialsForCliDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *CredentialsForCliDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### SetCloudCredentialNameNil

`func (o *CredentialsForCliDto) SetCloudCredentialNameNil(b bool)`

 SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil

### UnsetCloudCredentialName
`func (o *CredentialsForCliDto) UnsetCloudCredentialName()`

UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
### GetCloudType

`func (o *CredentialsForCliDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CredentialsForCliDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CredentialsForCliDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *CredentialsForCliDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### SetCloudTypeNil

`func (o *CredentialsForCliDto) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *CredentialsForCliDto) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetOrganizationId

`func (o *CredentialsForCliDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CredentialsForCliDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CredentialsForCliDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CredentialsForCliDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *CredentialsForCliDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CredentialsForCliDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CredentialsForCliDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CredentialsForCliDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *CredentialsForCliDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *CredentialsForCliDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetPartnerName

`func (o *CredentialsForCliDto) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *CredentialsForCliDto) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *CredentialsForCliDto) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *CredentialsForCliDto) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *CredentialsForCliDto) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *CredentialsForCliDto) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


