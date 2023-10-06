# ServersListForDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServerListDto**](ServerListDto.md) |  | [optional] 
**Project** | Pointer to [**ProjectDetailsForServersDto**](ProjectDetailsForServersDto.md) |  | [optional] 

## Methods

### NewServersListForDetails

`func NewServersListForDetails() *ServersListForDetails`

NewServersListForDetails instantiates a new ServersListForDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersListForDetailsWithDefaults

`func NewServersListForDetailsWithDefaults() *ServersListForDetails`

NewServersListForDetailsWithDefaults instantiates a new ServersListForDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServersListForDetails) GetData() []ServerListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServersListForDetails) GetDataOk() (*[]ServerListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServersListForDetails) SetData(v []ServerListDto)`

SetData sets Data field to given value.

### HasData

`func (o *ServersListForDetails) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ServersListForDetails) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ServersListForDetails) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetProject

`func (o *ServersListForDetails) GetProject() ProjectDetailsForServersDto`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServersListForDetails) GetProjectOk() (*ProjectDetailsForServersDto, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServersListForDetails) SetProject(v ProjectDetailsForServersDto)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServersListForDetails) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


