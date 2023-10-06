# AwsImagesPostListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**SortBy** | Pointer to **NullableString** |  | [optional] 
**SortDirection** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Latest** | Pointer to **bool** |  | [optional] 
**Owners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAwsImagesPostListCommand

`func NewAwsImagesPostListCommand() *AwsImagesPostListCommand`

NewAwsImagesPostListCommand instantiates a new AwsImagesPostListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsImagesPostListCommandWithDefaults

`func NewAwsImagesPostListCommandWithDefaults() *AwsImagesPostListCommand`

NewAwsImagesPostListCommandWithDefaults instantiates a new AwsImagesPostListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *AwsImagesPostListCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AwsImagesPostListCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AwsImagesPostListCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *AwsImagesPostListCommand) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetLimit

`func (o *AwsImagesPostListCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AwsImagesPostListCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AwsImagesPostListCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AwsImagesPostListCommand) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *AwsImagesPostListCommand) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AwsImagesPostListCommand) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *AwsImagesPostListCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsImagesPostListCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsImagesPostListCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AwsImagesPostListCommand) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *AwsImagesPostListCommand) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *AwsImagesPostListCommand) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSortBy

`func (o *AwsImagesPostListCommand) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *AwsImagesPostListCommand) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *AwsImagesPostListCommand) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *AwsImagesPostListCommand) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *AwsImagesPostListCommand) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *AwsImagesPostListCommand) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortDirection

`func (o *AwsImagesPostListCommand) GetSortDirection() string`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *AwsImagesPostListCommand) GetSortDirectionOk() (*string, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *AwsImagesPostListCommand) SetSortDirection(v string)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *AwsImagesPostListCommand) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### SetSortDirectionNil

`func (o *AwsImagesPostListCommand) SetSortDirectionNil(b bool)`

 SetSortDirectionNil sets the value for SortDirection to be an explicit nil

### UnsetSortDirection
`func (o *AwsImagesPostListCommand) UnsetSortDirection()`

UnsetSortDirection ensures that no value is present for SortDirection, not even an explicit nil
### GetSearch

`func (o *AwsImagesPostListCommand) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AwsImagesPostListCommand) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AwsImagesPostListCommand) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AwsImagesPostListCommand) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *AwsImagesPostListCommand) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *AwsImagesPostListCommand) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetLatest

`func (o *AwsImagesPostListCommand) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *AwsImagesPostListCommand) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *AwsImagesPostListCommand) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *AwsImagesPostListCommand) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetOwners

`func (o *AwsImagesPostListCommand) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AwsImagesPostListCommand) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AwsImagesPostListCommand) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AwsImagesPostListCommand) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwnersNil

`func (o *AwsImagesPostListCommand) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *AwsImagesPostListCommand) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


