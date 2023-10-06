# ReplyTicketCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReplyTicketCommand

`func NewReplyTicketCommand() *ReplyTicketCommand`

NewReplyTicketCommand instantiates a new ReplyTicketCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyTicketCommandWithDefaults

`func NewReplyTicketCommandWithDefaults() *ReplyTicketCommand`

NewReplyTicketCommandWithDefaults instantiates a new ReplyTicketCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *ReplyTicketCommand) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *ReplyTicketCommand) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *ReplyTicketCommand) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *ReplyTicketCommand) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### SetTicketIdNil

`func (o *ReplyTicketCommand) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *ReplyTicketCommand) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetBody

`func (o *ReplyTicketCommand) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReplyTicketCommand) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReplyTicketCommand) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReplyTicketCommand) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ReplyTicketCommand) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ReplyTicketCommand) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


