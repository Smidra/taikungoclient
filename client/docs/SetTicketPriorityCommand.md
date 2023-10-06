# SetTicketPriorityCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**TicketPriority**](TicketPriority.md) |  | [optional] 

## Methods

### NewSetTicketPriorityCommand

`func NewSetTicketPriorityCommand() *SetTicketPriorityCommand`

NewSetTicketPriorityCommand instantiates a new SetTicketPriorityCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetTicketPriorityCommandWithDefaults

`func NewSetTicketPriorityCommandWithDefaults() *SetTicketPriorityCommand`

NewSetTicketPriorityCommandWithDefaults instantiates a new SetTicketPriorityCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetTicketPriorityCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetTicketPriorityCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetTicketPriorityCommand) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SetTicketPriorityCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SetTicketPriorityCommand) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SetTicketPriorityCommand) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPriority

`func (o *SetTicketPriorityCommand) GetPriority() TicketPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SetTicketPriorityCommand) GetPriorityOk() (*TicketPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SetTicketPriorityCommand) SetPriority(v TicketPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SetTicketPriorityCommand) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


