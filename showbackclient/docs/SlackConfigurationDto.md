# SlackConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Channel** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**SlackType** | Pointer to [**SlackType**](SlackType.md) |  | [optional] 

## Methods

### NewSlackConfigurationDto

`func NewSlackConfigurationDto() *SlackConfigurationDto`

NewSlackConfigurationDto instantiates a new SlackConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackConfigurationDtoWithDefaults

`func NewSlackConfigurationDtoWithDefaults() *SlackConfigurationDto`

NewSlackConfigurationDtoWithDefaults instantiates a new SlackConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlackConfigurationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlackConfigurationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlackConfigurationDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SlackConfigurationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SlackConfigurationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackConfigurationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackConfigurationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlackConfigurationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SlackConfigurationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SlackConfigurationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *SlackConfigurationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SlackConfigurationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SlackConfigurationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SlackConfigurationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SlackConfigurationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SlackConfigurationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetChannel

`func (o *SlackConfigurationDto) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SlackConfigurationDto) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SlackConfigurationDto) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SlackConfigurationDto) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *SlackConfigurationDto) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *SlackConfigurationDto) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetOrganizationName

`func (o *SlackConfigurationDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *SlackConfigurationDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *SlackConfigurationDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *SlackConfigurationDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *SlackConfigurationDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *SlackConfigurationDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *SlackConfigurationDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SlackConfigurationDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SlackConfigurationDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SlackConfigurationDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *SlackConfigurationDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *SlackConfigurationDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetSlackType

`func (o *SlackConfigurationDto) GetSlackType() SlackType`

GetSlackType returns the SlackType field if non-nil, zero value otherwise.

### GetSlackTypeOk

`func (o *SlackConfigurationDto) GetSlackTypeOk() (*SlackType, bool)`

GetSlackTypeOk returns a tuple with the SlackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackType

`func (o *SlackConfigurationDto) SetSlackType(v SlackType)`

SetSlackType sets SlackType field to given value.

### HasSlackType

`func (o *SlackConfigurationDto) HasSlackType() bool`

HasSlackType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


