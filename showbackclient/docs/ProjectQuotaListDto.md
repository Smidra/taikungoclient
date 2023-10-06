# ProjectQuotaListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerCpu** | Pointer to **int64** |  | [optional] 
**ServerRam** | Pointer to **int64** |  | [optional] 
**ServerDiskSize** | Pointer to **int64** |  | [optional] 
**VmCpu** | Pointer to **int64** |  | [optional] 
**VmRam** | Pointer to **int64** |  | [optional] 
**VmVolumeSize** | Pointer to **int64** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectQuotaListDto

`func NewProjectQuotaListDto() *ProjectQuotaListDto`

NewProjectQuotaListDto instantiates a new ProjectQuotaListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectQuotaListDtoWithDefaults

`func NewProjectQuotaListDtoWithDefaults() *ProjectQuotaListDto`

NewProjectQuotaListDtoWithDefaults instantiates a new ProjectQuotaListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerCpu

`func (o *ProjectQuotaListDto) GetServerCpu() int64`

GetServerCpu returns the ServerCpu field if non-nil, zero value otherwise.

### GetServerCpuOk

`func (o *ProjectQuotaListDto) GetServerCpuOk() (*int64, bool)`

GetServerCpuOk returns a tuple with the ServerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCpu

`func (o *ProjectQuotaListDto) SetServerCpu(v int64)`

SetServerCpu sets ServerCpu field to given value.

### HasServerCpu

`func (o *ProjectQuotaListDto) HasServerCpu() bool`

HasServerCpu returns a boolean if a field has been set.

### GetServerRam

`func (o *ProjectQuotaListDto) GetServerRam() int64`

GetServerRam returns the ServerRam field if non-nil, zero value otherwise.

### GetServerRamOk

`func (o *ProjectQuotaListDto) GetServerRamOk() (*int64, bool)`

GetServerRamOk returns a tuple with the ServerRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRam

`func (o *ProjectQuotaListDto) SetServerRam(v int64)`

SetServerRam sets ServerRam field to given value.

### HasServerRam

`func (o *ProjectQuotaListDto) HasServerRam() bool`

HasServerRam returns a boolean if a field has been set.

### GetServerDiskSize

`func (o *ProjectQuotaListDto) GetServerDiskSize() int64`

GetServerDiskSize returns the ServerDiskSize field if non-nil, zero value otherwise.

### GetServerDiskSizeOk

`func (o *ProjectQuotaListDto) GetServerDiskSizeOk() (*int64, bool)`

GetServerDiskSizeOk returns a tuple with the ServerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDiskSize

`func (o *ProjectQuotaListDto) SetServerDiskSize(v int64)`

SetServerDiskSize sets ServerDiskSize field to given value.

### HasServerDiskSize

`func (o *ProjectQuotaListDto) HasServerDiskSize() bool`

HasServerDiskSize returns a boolean if a field has been set.

### GetVmCpu

`func (o *ProjectQuotaListDto) GetVmCpu() int64`

GetVmCpu returns the VmCpu field if non-nil, zero value otherwise.

### GetVmCpuOk

`func (o *ProjectQuotaListDto) GetVmCpuOk() (*int64, bool)`

GetVmCpuOk returns a tuple with the VmCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpu

`func (o *ProjectQuotaListDto) SetVmCpu(v int64)`

SetVmCpu sets VmCpu field to given value.

### HasVmCpu

`func (o *ProjectQuotaListDto) HasVmCpu() bool`

HasVmCpu returns a boolean if a field has been set.

### GetVmRam

`func (o *ProjectQuotaListDto) GetVmRam() int64`

GetVmRam returns the VmRam field if non-nil, zero value otherwise.

### GetVmRamOk

`func (o *ProjectQuotaListDto) GetVmRamOk() (*int64, bool)`

GetVmRamOk returns a tuple with the VmRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRam

`func (o *ProjectQuotaListDto) SetVmRam(v int64)`

SetVmRam sets VmRam field to given value.

### HasVmRam

`func (o *ProjectQuotaListDto) HasVmRam() bool`

HasVmRam returns a boolean if a field has been set.

### GetVmVolumeSize

`func (o *ProjectQuotaListDto) GetVmVolumeSize() int64`

GetVmVolumeSize returns the VmVolumeSize field if non-nil, zero value otherwise.

### GetVmVolumeSizeOk

`func (o *ProjectQuotaListDto) GetVmVolumeSizeOk() (*int64, bool)`

GetVmVolumeSizeOk returns a tuple with the VmVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVolumeSize

`func (o *ProjectQuotaListDto) SetVmVolumeSize(v int64)`

SetVmVolumeSize sets VmVolumeSize field to given value.

### HasVmVolumeSize

`func (o *ProjectQuotaListDto) HasVmVolumeSize() bool`

HasVmVolumeSize returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectQuotaListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectQuotaListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectQuotaListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectQuotaListDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *ProjectQuotaListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectQuotaListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectQuotaListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectQuotaListDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ProjectQuotaListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectQuotaListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


