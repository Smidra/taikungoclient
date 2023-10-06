# CloudCredentialsDropdownRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CloudType** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to [**[]ProjectWithFlavorsAndImagesDto**](ProjectWithFlavorsAndImagesDto.md) |  | [optional] 

## Methods

### NewCloudCredentialsDropdownRecordDto

`func NewCloudCredentialsDropdownRecordDto() *CloudCredentialsDropdownRecordDto`

NewCloudCredentialsDropdownRecordDto instantiates a new CloudCredentialsDropdownRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsDropdownRecordDtoWithDefaults

`func NewCloudCredentialsDropdownRecordDtoWithDefaults() *CloudCredentialsDropdownRecordDto`

NewCloudCredentialsDropdownRecordDtoWithDefaults instantiates a new CloudCredentialsDropdownRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudCredentialsDropdownRecordDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialsDropdownRecordDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialsDropdownRecordDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCredentialsDropdownRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudCredentialsDropdownRecordDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentialsDropdownRecordDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentialsDropdownRecordDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCredentialsDropdownRecordDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CloudCredentialsDropdownRecordDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloudCredentialsDropdownRecordDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCloudType

`func (o *CloudCredentialsDropdownRecordDto) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CloudCredentialsDropdownRecordDto) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CloudCredentialsDropdownRecordDto) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *CloudCredentialsDropdownRecordDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### SetCloudTypeNil

`func (o *CloudCredentialsDropdownRecordDto) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *CloudCredentialsDropdownRecordDto) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetProjects

`func (o *CloudCredentialsDropdownRecordDto) GetProjects() []ProjectWithFlavorsAndImagesDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CloudCredentialsDropdownRecordDto) GetProjectsOk() (*[]ProjectWithFlavorsAndImagesDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CloudCredentialsDropdownRecordDto) SetProjects(v []ProjectWithFlavorsAndImagesDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CloudCredentialsDropdownRecordDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *CloudCredentialsDropdownRecordDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *CloudCredentialsDropdownRecordDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


