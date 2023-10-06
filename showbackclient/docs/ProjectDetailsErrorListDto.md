# ProjectDetailsErrorListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProjectDetailsErrorType**](ProjectDetailsErrorType.md) |  | [optional] 
**Message** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectDetailsErrorListDto

`func NewProjectDetailsErrorListDto() *ProjectDetailsErrorListDto`

NewProjectDetailsErrorListDto instantiates a new ProjectDetailsErrorListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsErrorListDtoWithDefaults

`func NewProjectDetailsErrorListDtoWithDefaults() *ProjectDetailsErrorListDto`

NewProjectDetailsErrorListDtoWithDefaults instantiates a new ProjectDetailsErrorListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProjectDetailsErrorListDto) GetType() ProjectDetailsErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectDetailsErrorListDto) GetTypeOk() (*ProjectDetailsErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectDetailsErrorListDto) SetType(v ProjectDetailsErrorType)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectDetailsErrorListDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectDetailsErrorListDto) GetMessage() []string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectDetailsErrorListDto) GetMessageOk() (*[]string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectDetailsErrorListDto) SetMessage(v []string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectDetailsErrorListDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProjectDetailsErrorListDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProjectDetailsErrorListDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


