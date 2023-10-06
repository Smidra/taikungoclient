# EditArticleCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEditArticleCommand

`func NewEditArticleCommand() *EditArticleCommand`

NewEditArticleCommand instantiates a new EditArticleCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditArticleCommandWithDefaults

`func NewEditArticleCommandWithDefaults() *EditArticleCommand`

NewEditArticleCommandWithDefaults instantiates a new EditArticleCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *EditArticleCommand) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EditArticleCommand) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EditArticleCommand) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EditArticleCommand) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### SetMessageIdNil

`func (o *EditArticleCommand) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *EditArticleCommand) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetBody

`func (o *EditArticleCommand) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EditArticleCommand) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EditArticleCommand) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EditArticleCommand) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *EditArticleCommand) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *EditArticleCommand) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


