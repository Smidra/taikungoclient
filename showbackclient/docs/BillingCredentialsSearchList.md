# BillingCredentialsSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CommonSearchResponseData**](CommonSearchResponseData.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBillingCredentialsSearchList

`func NewBillingCredentialsSearchList() *BillingCredentialsSearchList`

NewBillingCredentialsSearchList instantiates a new BillingCredentialsSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCredentialsSearchListWithDefaults

`func NewBillingCredentialsSearchListWithDefaults() *BillingCredentialsSearchList`

NewBillingCredentialsSearchListWithDefaults instantiates a new BillingCredentialsSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BillingCredentialsSearchList) GetData() []CommonSearchResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingCredentialsSearchList) GetDataOk() (*[]CommonSearchResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingCredentialsSearchList) SetData(v []CommonSearchResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *BillingCredentialsSearchList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BillingCredentialsSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BillingCredentialsSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *BillingCredentialsSearchList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BillingCredentialsSearchList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BillingCredentialsSearchList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BillingCredentialsSearchList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


