# AzResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAzResult

`func NewAzResult() *AzResult`

NewAzResult instantiates a new AzResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzResultWithDefaults

`func NewAzResultWithDefaults() *AzResult`

NewAzResultWithDefaults instantiates a new AzResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *AzResult) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *AzResult) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *AzResult) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *AzResult) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *AzResult) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *AzResult) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetTotalCount

`func (o *AzResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AzResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AzResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AzResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


