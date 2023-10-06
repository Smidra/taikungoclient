# KubernetesCronJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | Pointer to **NullableString** |  | [optional] 
**LastSchedule** | Pointer to **NullableString** |  | [optional] 
**Suspend** | Pointer to **NullableBool** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKubernetesCronJobDto

`func NewKubernetesCronJobDto() *KubernetesCronJobDto`

NewKubernetesCronJobDto instantiates a new KubernetesCronJobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCronJobDtoWithDefaults

`func NewKubernetesCronJobDtoWithDefaults() *KubernetesCronJobDto`

NewKubernetesCronJobDtoWithDefaults instantiates a new KubernetesCronJobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *KubernetesCronJobDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *KubernetesCronJobDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *KubernetesCronJobDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *KubernetesCronJobDto) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### SetMetadataNameNil

`func (o *KubernetesCronJobDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *KubernetesCronJobDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetLastSchedule

`func (o *KubernetesCronJobDto) GetLastSchedule() string`

GetLastSchedule returns the LastSchedule field if non-nil, zero value otherwise.

### GetLastScheduleOk

`func (o *KubernetesCronJobDto) GetLastScheduleOk() (*string, bool)`

GetLastScheduleOk returns a tuple with the LastSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedule

`func (o *KubernetesCronJobDto) SetLastSchedule(v string)`

SetLastSchedule sets LastSchedule field to given value.

### HasLastSchedule

`func (o *KubernetesCronJobDto) HasLastSchedule() bool`

HasLastSchedule returns a boolean if a field has been set.

### SetLastScheduleNil

`func (o *KubernetesCronJobDto) SetLastScheduleNil(b bool)`

 SetLastScheduleNil sets the value for LastSchedule to be an explicit nil

### UnsetLastSchedule
`func (o *KubernetesCronJobDto) UnsetLastSchedule()`

UnsetLastSchedule ensures that no value is present for LastSchedule, not even an explicit nil
### GetSuspend

`func (o *KubernetesCronJobDto) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *KubernetesCronJobDto) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *KubernetesCronJobDto) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *KubernetesCronJobDto) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### SetSuspendNil

`func (o *KubernetesCronJobDto) SetSuspendNil(b bool)`

 SetSuspendNil sets the value for Suspend to be an explicit nil

### UnsetSuspend
`func (o *KubernetesCronJobDto) UnsetSuspend()`

UnsetSuspend ensures that no value is present for Suspend, not even an explicit nil
### GetSchedule

`func (o *KubernetesCronJobDto) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KubernetesCronJobDto) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KubernetesCronJobDto) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KubernetesCronJobDto) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *KubernetesCronJobDto) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *KubernetesCronJobDto) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetNamespace

`func (o *KubernetesCronJobDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesCronJobDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesCronJobDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesCronJobDto) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *KubernetesCronJobDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *KubernetesCronJobDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetAge

`func (o *KubernetesCronJobDto) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *KubernetesCronJobDto) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *KubernetesCronJobDto) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *KubernetesCronJobDto) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *KubernetesCronJobDto) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *KubernetesCronJobDto) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


