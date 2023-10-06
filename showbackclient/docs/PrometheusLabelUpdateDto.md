# PrometheusLabelUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrometheusLabelUpdateDto

`func NewPrometheusLabelUpdateDto() *PrometheusLabelUpdateDto`

NewPrometheusLabelUpdateDto instantiates a new PrometheusLabelUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusLabelUpdateDtoWithDefaults

`func NewPrometheusLabelUpdateDtoWithDefaults() *PrometheusLabelUpdateDto`

NewPrometheusLabelUpdateDtoWithDefaults instantiates a new PrometheusLabelUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusLabelUpdateDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusLabelUpdateDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusLabelUpdateDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PrometheusLabelUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PrometheusLabelUpdateDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PrometheusLabelUpdateDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PrometheusLabelUpdateDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PrometheusLabelUpdateDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PrometheusLabelUpdateDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PrometheusLabelUpdateDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetValue

`func (o *PrometheusLabelUpdateDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrometheusLabelUpdateDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrometheusLabelUpdateDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PrometheusLabelUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PrometheusLabelUpdateDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PrometheusLabelUpdateDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


