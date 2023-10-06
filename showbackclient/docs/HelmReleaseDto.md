# HelmReleaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to [**HelmMetadata**](HelmMetadata.md) |  | [optional] 
**Spec** | Pointer to [**HelmSpec**](HelmSpec.md) |  | [optional] 
**Status** | Pointer to [**HelmStatus**](HelmStatus.md) |  | [optional] 

## Methods

### NewHelmReleaseDto

`func NewHelmReleaseDto() *HelmReleaseDto`

NewHelmReleaseDto instantiates a new HelmReleaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmReleaseDtoWithDefaults

`func NewHelmReleaseDtoWithDefaults() *HelmReleaseDto`

NewHelmReleaseDtoWithDefaults instantiates a new HelmReleaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HelmReleaseDto) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HelmReleaseDto) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HelmReleaseDto) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HelmReleaseDto) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *HelmReleaseDto) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *HelmReleaseDto) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetKind

`func (o *HelmReleaseDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HelmReleaseDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HelmReleaseDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HelmReleaseDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *HelmReleaseDto) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *HelmReleaseDto) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetMetadata

`func (o *HelmReleaseDto) GetMetadata() HelmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HelmReleaseDto) GetMetadataOk() (*HelmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HelmReleaseDto) SetMetadata(v HelmMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HelmReleaseDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HelmReleaseDto) GetSpec() HelmSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HelmReleaseDto) GetSpecOk() (*HelmSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HelmReleaseDto) SetSpec(v HelmSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *HelmReleaseDto) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *HelmReleaseDto) GetStatus() HelmStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HelmReleaseDto) GetStatusOk() (*HelmStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HelmReleaseDto) SetStatus(v HelmStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HelmReleaseDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


