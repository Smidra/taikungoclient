# PodDisruptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PodDisruptionDto**](PodDisruptionDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPodDisruptions

`func NewPodDisruptions() *PodDisruptions`

NewPodDisruptions instantiates a new PodDisruptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodDisruptionsWithDefaults

`func NewPodDisruptionsWithDefaults() *PodDisruptions`

NewPodDisruptionsWithDefaults instantiates a new PodDisruptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PodDisruptions) GetData() []PodDisruptionDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PodDisruptions) GetDataOk() (*[]PodDisruptionDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PodDisruptions) SetData(v []PodDisruptionDto)`

SetData sets Data field to given value.

### HasData

`func (o *PodDisruptions) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PodDisruptions) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PodDisruptions) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *PodDisruptions) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PodDisruptions) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PodDisruptions) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PodDisruptions) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


