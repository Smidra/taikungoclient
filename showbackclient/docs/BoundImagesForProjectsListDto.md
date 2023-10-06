# BoundImagesForProjectsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**IsAzure** | Pointer to **bool** |  | [optional] 
**IsAws** | Pointer to **bool** |  | [optional] 
**IsOpenstack** | Pointer to **bool** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableFloat64** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**CloudId** | Pointer to **NullableInt32** |  | [optional] 
**IsWindows** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoundImagesForProjectsListDto

`func NewBoundImagesForProjectsListDto() *BoundImagesForProjectsListDto`

NewBoundImagesForProjectsListDto instantiates a new BoundImagesForProjectsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundImagesForProjectsListDtoWithDefaults

`func NewBoundImagesForProjectsListDtoWithDefaults() *BoundImagesForProjectsListDto`

NewBoundImagesForProjectsListDtoWithDefaults instantiates a new BoundImagesForProjectsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoundImagesForProjectsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoundImagesForProjectsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoundImagesForProjectsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoundImagesForProjectsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoundImagesForProjectsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundImagesForProjectsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundImagesForProjectsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoundImagesForProjectsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BoundImagesForProjectsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BoundImagesForProjectsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *BoundImagesForProjectsListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BoundImagesForProjectsListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BoundImagesForProjectsListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BoundImagesForProjectsListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BoundImagesForProjectsListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BoundImagesForProjectsListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIsAzure

`func (o *BoundImagesForProjectsListDto) GetIsAzure() bool`

GetIsAzure returns the IsAzure field if non-nil, zero value otherwise.

### GetIsAzureOk

`func (o *BoundImagesForProjectsListDto) GetIsAzureOk() (*bool, bool)`

GetIsAzureOk returns a tuple with the IsAzure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzure

`func (o *BoundImagesForProjectsListDto) SetIsAzure(v bool)`

SetIsAzure sets IsAzure field to given value.

### HasIsAzure

`func (o *BoundImagesForProjectsListDto) HasIsAzure() bool`

HasIsAzure returns a boolean if a field has been set.

### GetIsAws

`func (o *BoundImagesForProjectsListDto) GetIsAws() bool`

GetIsAws returns the IsAws field if non-nil, zero value otherwise.

### GetIsAwsOk

`func (o *BoundImagesForProjectsListDto) GetIsAwsOk() (*bool, bool)`

GetIsAwsOk returns a tuple with the IsAws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAws

`func (o *BoundImagesForProjectsListDto) SetIsAws(v bool)`

SetIsAws sets IsAws field to given value.

### HasIsAws

`func (o *BoundImagesForProjectsListDto) HasIsAws() bool`

HasIsAws returns a boolean if a field has been set.

### GetIsOpenstack

`func (o *BoundImagesForProjectsListDto) GetIsOpenstack() bool`

GetIsOpenstack returns the IsOpenstack field if non-nil, zero value otherwise.

### GetIsOpenstackOk

`func (o *BoundImagesForProjectsListDto) GetIsOpenstackOk() (*bool, bool)`

GetIsOpenstackOk returns a tuple with the IsOpenstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenstack

`func (o *BoundImagesForProjectsListDto) SetIsOpenstack(v bool)`

SetIsOpenstack sets IsOpenstack field to given value.

### HasIsOpenstack

`func (o *BoundImagesForProjectsListDto) HasIsOpenstack() bool`

HasIsOpenstack returns a boolean if a field has been set.

### GetProjectName

`func (o *BoundImagesForProjectsListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BoundImagesForProjectsListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BoundImagesForProjectsListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BoundImagesForProjectsListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BoundImagesForProjectsListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BoundImagesForProjectsListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetSize

`func (o *BoundImagesForProjectsListDto) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BoundImagesForProjectsListDto) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BoundImagesForProjectsListDto) SetSize(v float64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BoundImagesForProjectsListDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BoundImagesForProjectsListDto) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BoundImagesForProjectsListDto) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetImageId

`func (o *BoundImagesForProjectsListDto) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *BoundImagesForProjectsListDto) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *BoundImagesForProjectsListDto) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *BoundImagesForProjectsListDto) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *BoundImagesForProjectsListDto) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *BoundImagesForProjectsListDto) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetCloudId

`func (o *BoundImagesForProjectsListDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *BoundImagesForProjectsListDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *BoundImagesForProjectsListDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *BoundImagesForProjectsListDto) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### SetCloudIdNil

`func (o *BoundImagesForProjectsListDto) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *BoundImagesForProjectsListDto) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetIsWindows

`func (o *BoundImagesForProjectsListDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *BoundImagesForProjectsListDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *BoundImagesForProjectsListDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *BoundImagesForProjectsListDto) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


