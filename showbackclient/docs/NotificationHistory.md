# NotificationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]NotificationListDto**](NotificationListDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewNotificationHistory

`func NewNotificationHistory() *NotificationHistory`

NewNotificationHistory instantiates a new NotificationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationHistoryWithDefaults

`func NewNotificationHistoryWithDefaults() *NotificationHistory`

NewNotificationHistoryWithDefaults instantiates a new NotificationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NotificationHistory) GetData() []NotificationListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationHistory) GetDataOk() (*[]NotificationListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationHistory) SetData(v []NotificationListDto)`

SetData sets Data field to given value.

### HasData

`func (o *NotificationHistory) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *NotificationHistory) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *NotificationHistory) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *NotificationHistory) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NotificationHistory) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NotificationHistory) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *NotificationHistory) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


