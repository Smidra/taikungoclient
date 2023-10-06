# FromTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTrigger** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFromTemplateDto

`func NewFromTemplateDto() *FromTemplateDto`

NewFromTemplateDto instantiates a new FromTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFromTemplateDtoWithDefaults

`func NewFromTemplateDtoWithDefaults() *FromTemplateDto`

NewFromTemplateDtoWithDefaults instantiates a new FromTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTrigger

`func (o *FromTemplateDto) GetAutoTrigger() bool`

GetAutoTrigger returns the AutoTrigger field if non-nil, zero value otherwise.

### GetAutoTriggerOk

`func (o *FromTemplateDto) GetAutoTriggerOk() (*bool, bool)`

GetAutoTriggerOk returns a tuple with the AutoTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTrigger

`func (o *FromTemplateDto) SetAutoTrigger(v bool)`

SetAutoTrigger sets AutoTrigger field to given value.

### HasAutoTrigger

`func (o *FromTemplateDto) HasAutoTrigger() bool`

HasAutoTrigger returns a boolean if a field has been set.

### GetUserId

`func (o *FromTemplateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FromTemplateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FromTemplateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FromTemplateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *FromTemplateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *FromTemplateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *FromTemplateDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FromTemplateDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FromTemplateDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FromTemplateDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *FromTemplateDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FromTemplateDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


