# PrometheusDashboardDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ExpressionDecoded** | Pointer to **NullableString** |  | [optional] 
**ExpressionEncoded** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsReadonly** | Pointer to **bool** |  | [optional] 

## Methods

### NewPrometheusDashboardDto

`func NewPrometheusDashboardDto() *PrometheusDashboardDto`

NewPrometheusDashboardDto instantiates a new PrometheusDashboardDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusDashboardDtoWithDefaults

`func NewPrometheusDashboardDtoWithDefaults() *PrometheusDashboardDto`

NewPrometheusDashboardDtoWithDefaults instantiates a new PrometheusDashboardDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusDashboardDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusDashboardDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusDashboardDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PrometheusDashboardDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrometheusDashboardDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrometheusDashboardDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrometheusDashboardDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrometheusDashboardDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PrometheusDashboardDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PrometheusDashboardDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpressionDecoded

`func (o *PrometheusDashboardDto) GetExpressionDecoded() string`

GetExpressionDecoded returns the ExpressionDecoded field if non-nil, zero value otherwise.

### GetExpressionDecodedOk

`func (o *PrometheusDashboardDto) GetExpressionDecodedOk() (*string, bool)`

GetExpressionDecodedOk returns a tuple with the ExpressionDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionDecoded

`func (o *PrometheusDashboardDto) SetExpressionDecoded(v string)`

SetExpressionDecoded sets ExpressionDecoded field to given value.

### HasExpressionDecoded

`func (o *PrometheusDashboardDto) HasExpressionDecoded() bool`

HasExpressionDecoded returns a boolean if a field has been set.

### SetExpressionDecodedNil

`func (o *PrometheusDashboardDto) SetExpressionDecodedNil(b bool)`

 SetExpressionDecodedNil sets the value for ExpressionDecoded to be an explicit nil

### UnsetExpressionDecoded
`func (o *PrometheusDashboardDto) UnsetExpressionDecoded()`

UnsetExpressionDecoded ensures that no value is present for ExpressionDecoded, not even an explicit nil
### GetExpressionEncoded

`func (o *PrometheusDashboardDto) GetExpressionEncoded() string`

GetExpressionEncoded returns the ExpressionEncoded field if non-nil, zero value otherwise.

### GetExpressionEncodedOk

`func (o *PrometheusDashboardDto) GetExpressionEncodedOk() (*string, bool)`

GetExpressionEncodedOk returns a tuple with the ExpressionEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionEncoded

`func (o *PrometheusDashboardDto) SetExpressionEncoded(v string)`

SetExpressionEncoded sets ExpressionEncoded field to given value.

### HasExpressionEncoded

`func (o *PrometheusDashboardDto) HasExpressionEncoded() bool`

HasExpressionEncoded returns a boolean if a field has been set.

### SetExpressionEncodedNil

`func (o *PrometheusDashboardDto) SetExpressionEncodedNil(b bool)`

 SetExpressionEncodedNil sets the value for ExpressionEncoded to be an explicit nil

### UnsetExpressionEncoded
`func (o *PrometheusDashboardDto) UnsetExpressionEncoded()`

UnsetExpressionEncoded ensures that no value is present for ExpressionEncoded, not even an explicit nil
### GetDescription

`func (o *PrometheusDashboardDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrometheusDashboardDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrometheusDashboardDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrometheusDashboardDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrometheusDashboardDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrometheusDashboardDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsReadonly

`func (o *PrometheusDashboardDto) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *PrometheusDashboardDto) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *PrometheusDashboardDto) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.

### HasIsReadonly

`func (o *PrometheusDashboardDto) HasIsReadonly() bool`

HasIsReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


