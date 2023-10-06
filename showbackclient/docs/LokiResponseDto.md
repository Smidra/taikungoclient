# LokiResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]Parameter**](Parameter.md) |  | [optional] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Direction** | Pointer to **NullableString** |  | [optional] 
**CanDownload** | Pointer to **bool** |  | [optional] 

## Methods

### NewLokiResponseDto

`func NewLokiResponseDto() *LokiResponseDto`

NewLokiResponseDto instantiates a new LokiResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLokiResponseDtoWithDefaults

`func NewLokiResponseDtoWithDefaults() *LokiResponseDto`

NewLokiResponseDtoWithDefaults instantiates a new LokiResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *LokiResponseDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LokiResponseDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LokiResponseDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LokiResponseDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParameters

`func (o *LokiResponseDto) GetParameters() []Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LokiResponseDto) GetParametersOk() (*[]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LokiResponseDto) SetParameters(v []Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LokiResponseDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *LokiResponseDto) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *LokiResponseDto) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetFilters

`func (o *LokiResponseDto) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LokiResponseDto) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LokiResponseDto) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *LokiResponseDto) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *LokiResponseDto) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *LokiResponseDto) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetStart

`func (o *LokiResponseDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LokiResponseDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LokiResponseDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *LokiResponseDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *LokiResponseDto) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LokiResponseDto) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *LokiResponseDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LokiResponseDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LokiResponseDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LokiResponseDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *LokiResponseDto) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *LokiResponseDto) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetLimit

`func (o *LokiResponseDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LokiResponseDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LokiResponseDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LokiResponseDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *LokiResponseDto) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *LokiResponseDto) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetDirection

`func (o *LokiResponseDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LokiResponseDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LokiResponseDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LokiResponseDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *LokiResponseDto) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *LokiResponseDto) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetCanDownload

`func (o *LokiResponseDto) GetCanDownload() bool`

GetCanDownload returns the CanDownload field if non-nil, zero value otherwise.

### GetCanDownloadOk

`func (o *LokiResponseDto) GetCanDownloadOk() (*bool, bool)`

GetCanDownloadOk returns a tuple with the CanDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDownload

`func (o *LokiResponseDto) SetCanDownload(v bool)`

SetCanDownload sets CanDownload field to given value.

### HasCanDownload

`func (o *LokiResponseDto) HasCanDownload() bool`

HasCanDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


