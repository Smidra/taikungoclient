# AiCredentialDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AiType**](AiType.md) |  | [optional] 

## Methods

### NewAiCredentialDto

`func NewAiCredentialDto() *AiCredentialDto`

NewAiCredentialDto instantiates a new AiCredentialDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiCredentialDtoWithDefaults

`func NewAiCredentialDtoWithDefaults() *AiCredentialDto`

NewAiCredentialDtoWithDefaults instantiates a new AiCredentialDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AiCredentialDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiCredentialDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiCredentialDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiCredentialDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AiCredentialDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AiCredentialDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetName

`func (o *AiCredentialDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiCredentialDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiCredentialDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiCredentialDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AiCredentialDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AiCredentialDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetApiKey

`func (o *AiCredentialDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AiCredentialDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AiCredentialDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AiCredentialDto) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *AiCredentialDto) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *AiCredentialDto) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetType

`func (o *AiCredentialDto) GetType() AiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiCredentialDto) GetTypeOk() (*AiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiCredentialDto) SetType(v AiType)`

SetType sets Type field to given value.

### HasType

`func (o *AiCredentialDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


