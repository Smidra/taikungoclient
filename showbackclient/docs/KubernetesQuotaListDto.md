# KubernetesQuotaListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SumOfCpu** | Pointer to **float64** |  | [optional] 
**SumOfRamInGb** | Pointer to **float64** |  | [optional] 
**SumOfCpuUsage** | Pointer to **float64** |  | [optional] 
**SumOfRamUsage** | Pointer to **float64** |  | [optional] 
**PodsCapacity** | Pointer to **int32** |  | [optional] 
**PodsTotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewKubernetesQuotaListDto

`func NewKubernetesQuotaListDto() *KubernetesQuotaListDto`

NewKubernetesQuotaListDto instantiates a new KubernetesQuotaListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesQuotaListDtoWithDefaults

`func NewKubernetesQuotaListDtoWithDefaults() *KubernetesQuotaListDto`

NewKubernetesQuotaListDtoWithDefaults instantiates a new KubernetesQuotaListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSumOfCpu

`func (o *KubernetesQuotaListDto) GetSumOfCpu() float64`

GetSumOfCpu returns the SumOfCpu field if non-nil, zero value otherwise.

### GetSumOfCpuOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuOk() (*float64, bool)`

GetSumOfCpuOk returns a tuple with the SumOfCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpu

`func (o *KubernetesQuotaListDto) SetSumOfCpu(v float64)`

SetSumOfCpu sets SumOfCpu field to given value.

### HasSumOfCpu

`func (o *KubernetesQuotaListDto) HasSumOfCpu() bool`

HasSumOfCpu returns a boolean if a field has been set.

### GetSumOfRamInGb

`func (o *KubernetesQuotaListDto) GetSumOfRamInGb() float64`

GetSumOfRamInGb returns the SumOfRamInGb field if non-nil, zero value otherwise.

### GetSumOfRamInGbOk

`func (o *KubernetesQuotaListDto) GetSumOfRamInGbOk() (*float64, bool)`

GetSumOfRamInGbOk returns a tuple with the SumOfRamInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamInGb

`func (o *KubernetesQuotaListDto) SetSumOfRamInGb(v float64)`

SetSumOfRamInGb sets SumOfRamInGb field to given value.

### HasSumOfRamInGb

`func (o *KubernetesQuotaListDto) HasSumOfRamInGb() bool`

HasSumOfRamInGb returns a boolean if a field has been set.

### GetSumOfCpuUsage

`func (o *KubernetesQuotaListDto) GetSumOfCpuUsage() float64`

GetSumOfCpuUsage returns the SumOfCpuUsage field if non-nil, zero value otherwise.

### GetSumOfCpuUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuUsageOk() (*float64, bool)`

GetSumOfCpuUsageOk returns a tuple with the SumOfCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpuUsage

`func (o *KubernetesQuotaListDto) SetSumOfCpuUsage(v float64)`

SetSumOfCpuUsage sets SumOfCpuUsage field to given value.

### HasSumOfCpuUsage

`func (o *KubernetesQuotaListDto) HasSumOfCpuUsage() bool`

HasSumOfCpuUsage returns a boolean if a field has been set.

### GetSumOfRamUsage

`func (o *KubernetesQuotaListDto) GetSumOfRamUsage() float64`

GetSumOfRamUsage returns the SumOfRamUsage field if non-nil, zero value otherwise.

### GetSumOfRamUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfRamUsageOk() (*float64, bool)`

GetSumOfRamUsageOk returns a tuple with the SumOfRamUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamUsage

`func (o *KubernetesQuotaListDto) SetSumOfRamUsage(v float64)`

SetSumOfRamUsage sets SumOfRamUsage field to given value.

### HasSumOfRamUsage

`func (o *KubernetesQuotaListDto) HasSumOfRamUsage() bool`

HasSumOfRamUsage returns a boolean if a field has been set.

### GetPodsCapacity

`func (o *KubernetesQuotaListDto) GetPodsCapacity() int32`

GetPodsCapacity returns the PodsCapacity field if non-nil, zero value otherwise.

### GetPodsCapacityOk

`func (o *KubernetesQuotaListDto) GetPodsCapacityOk() (*int32, bool)`

GetPodsCapacityOk returns a tuple with the PodsCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsCapacity

`func (o *KubernetesQuotaListDto) SetPodsCapacity(v int32)`

SetPodsCapacity sets PodsCapacity field to given value.

### HasPodsCapacity

`func (o *KubernetesQuotaListDto) HasPodsCapacity() bool`

HasPodsCapacity returns a boolean if a field has been set.

### GetPodsTotalCount

`func (o *KubernetesQuotaListDto) GetPodsTotalCount() int32`

GetPodsTotalCount returns the PodsTotalCount field if non-nil, zero value otherwise.

### GetPodsTotalCountOk

`func (o *KubernetesQuotaListDto) GetPodsTotalCountOk() (*int32, bool)`

GetPodsTotalCountOk returns a tuple with the PodsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsTotalCount

`func (o *KubernetesQuotaListDto) SetPodsTotalCount(v int32)`

SetPodsTotalCount sets PodsTotalCount field to given value.

### HasPodsTotalCount

`func (o *KubernetesQuotaListDto) HasPodsTotalCount() bool`

HasPodsTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


