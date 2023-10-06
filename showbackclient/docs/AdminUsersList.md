# AdminUsersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AdminUsersResponseData**](AdminUsersResponseData.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminUsersList

`func NewAdminUsersList() *AdminUsersList`

NewAdminUsersList instantiates a new AdminUsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsersListWithDefaults

`func NewAdminUsersListWithDefaults() *AdminUsersList`

NewAdminUsersListWithDefaults instantiates a new AdminUsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AdminUsersList) GetData() []AdminUsersResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminUsersList) GetDataOk() (*[]AdminUsersResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminUsersList) SetData(v []AdminUsersResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *AdminUsersList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AdminUsersList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AdminUsersList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AdminUsersList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AdminUsersList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AdminUsersList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AdminUsersList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


