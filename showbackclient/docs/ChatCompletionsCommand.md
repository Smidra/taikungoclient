# ChatCompletionsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **NullableString** |  | [optional] 
**Messages** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChatCompletionsCommand

`func NewChatCompletionsCommand() *ChatCompletionsCommand`

NewChatCompletionsCommand instantiates a new ChatCompletionsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionsCommandWithDefaults

`func NewChatCompletionsCommandWithDefaults() *ChatCompletionsCommand`

NewChatCompletionsCommandWithDefaults instantiates a new ChatCompletionsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatCompletionsCommand) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionsCommand) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionsCommand) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatCompletionsCommand) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ChatCompletionsCommand) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ChatCompletionsCommand) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetMessages

`func (o *ChatCompletionsCommand) GetMessages() interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionsCommand) GetMessagesOk() (*interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionsCommand) SetMessages(v interface{})`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChatCompletionsCommand) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ChatCompletionsCommand) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ChatCompletionsCommand) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


