# PodListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**RestartCount** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Node** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPodListDto

`func NewPodListDto() *PodListDto`

NewPodListDto instantiates a new PodListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodListDtoWithDefaults

`func NewPodListDtoWithDefaults() *PodListDto`

NewPodListDtoWithDefaults instantiates a new PodListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *PodListDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *PodListDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *PodListDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *PodListDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *PodListDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *PodListDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetRestartCount

`func (o *PodListDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *PodListDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *PodListDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *PodListDto) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetNamespace

`func (o *PodListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PodListDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *PodListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *PodListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetNode

`func (o *PodListDto) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PodListDto) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PodListDto) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *PodListDto) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *PodListDto) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *PodListDto) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetAge

`func (o *PodListDto) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PodListDto) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PodListDto) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *PodListDto) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *PodListDto) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *PodListDto) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetStatus

`func (o *PodListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PodListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PodListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PodListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PodListDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PodListDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPhase

`func (o *PodListDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PodListDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PodListDto) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PodListDto) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *PodListDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *PodListDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetContainer

`func (o *PodListDto) GetContainer() []string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PodListDto) GetContainerOk() (*[]string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PodListDto) SetContainer(v []string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PodListDto) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *PodListDto) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *PodListDto) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


