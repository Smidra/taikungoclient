# ProjectActionVisibilityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**Repair** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**Upgrade** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**Monitoring** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableBackup** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableBackup** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableOpa** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableOpa** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableAutoscaler** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableAutoscaler** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**VmRepair** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**VmCommit** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**Lock** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**Unlock** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableSpotWorker** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableSpotWorker** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableFullSpot** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableFullSpot** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableSpotVm** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableSpotVm** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**AttachAlertingProfile** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DetachAlertingProfile** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**EnableAi** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**DisableAi** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 
**AiAssistant** | Pointer to [**ProjectButtonStatusDto**](ProjectButtonStatusDto.md) |  | [optional] 

## Methods

### NewProjectActionVisibilityDto

`func NewProjectActionVisibilityDto() *ProjectActionVisibilityDto`

NewProjectActionVisibilityDto instantiates a new ProjectActionVisibilityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectActionVisibilityDtoWithDefaults

`func NewProjectActionVisibilityDtoWithDefaults() *ProjectActionVisibilityDto`

NewProjectActionVisibilityDtoWithDefaults instantiates a new ProjectActionVisibilityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *ProjectActionVisibilityDto) GetCommit() ProjectButtonStatusDto`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectActionVisibilityDto) GetCommitOk() (*ProjectButtonStatusDto, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectActionVisibilityDto) SetCommit(v ProjectButtonStatusDto)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectActionVisibilityDto) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetRepair

`func (o *ProjectActionVisibilityDto) GetRepair() ProjectButtonStatusDto`

GetRepair returns the Repair field if non-nil, zero value otherwise.

### GetRepairOk

`func (o *ProjectActionVisibilityDto) GetRepairOk() (*ProjectButtonStatusDto, bool)`

GetRepairOk returns a tuple with the Repair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepair

`func (o *ProjectActionVisibilityDto) SetRepair(v ProjectButtonStatusDto)`

SetRepair sets Repair field to given value.

### HasRepair

`func (o *ProjectActionVisibilityDto) HasRepair() bool`

HasRepair returns a boolean if a field has been set.

### GetUpgrade

`func (o *ProjectActionVisibilityDto) GetUpgrade() ProjectButtonStatusDto`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *ProjectActionVisibilityDto) GetUpgradeOk() (*ProjectButtonStatusDto, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *ProjectActionVisibilityDto) SetUpgrade(v ProjectButtonStatusDto)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *ProjectActionVisibilityDto) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetMonitoring

`func (o *ProjectActionVisibilityDto) GetMonitoring() ProjectButtonStatusDto`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *ProjectActionVisibilityDto) GetMonitoringOk() (*ProjectButtonStatusDto, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *ProjectActionVisibilityDto) SetMonitoring(v ProjectButtonStatusDto)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *ProjectActionVisibilityDto) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.

### GetEnableBackup

`func (o *ProjectActionVisibilityDto) GetEnableBackup() ProjectButtonStatusDto`

GetEnableBackup returns the EnableBackup field if non-nil, zero value otherwise.

### GetEnableBackupOk

`func (o *ProjectActionVisibilityDto) GetEnableBackupOk() (*ProjectButtonStatusDto, bool)`

GetEnableBackupOk returns a tuple with the EnableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBackup

`func (o *ProjectActionVisibilityDto) SetEnableBackup(v ProjectButtonStatusDto)`

SetEnableBackup sets EnableBackup field to given value.

### HasEnableBackup

`func (o *ProjectActionVisibilityDto) HasEnableBackup() bool`

HasEnableBackup returns a boolean if a field has been set.

### GetDisableBackup

`func (o *ProjectActionVisibilityDto) GetDisableBackup() ProjectButtonStatusDto`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *ProjectActionVisibilityDto) GetDisableBackupOk() (*ProjectButtonStatusDto, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *ProjectActionVisibilityDto) SetDisableBackup(v ProjectButtonStatusDto)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *ProjectActionVisibilityDto) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetEnableOpa

`func (o *ProjectActionVisibilityDto) GetEnableOpa() ProjectButtonStatusDto`

GetEnableOpa returns the EnableOpa field if non-nil, zero value otherwise.

### GetEnableOpaOk

`func (o *ProjectActionVisibilityDto) GetEnableOpaOk() (*ProjectButtonStatusDto, bool)`

GetEnableOpaOk returns a tuple with the EnableOpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpa

`func (o *ProjectActionVisibilityDto) SetEnableOpa(v ProjectButtonStatusDto)`

SetEnableOpa sets EnableOpa field to given value.

### HasEnableOpa

`func (o *ProjectActionVisibilityDto) HasEnableOpa() bool`

HasEnableOpa returns a boolean if a field has been set.

### GetDisableOpa

`func (o *ProjectActionVisibilityDto) GetDisableOpa() ProjectButtonStatusDto`

GetDisableOpa returns the DisableOpa field if non-nil, zero value otherwise.

### GetDisableOpaOk

`func (o *ProjectActionVisibilityDto) GetDisableOpaOk() (*ProjectButtonStatusDto, bool)`

GetDisableOpaOk returns a tuple with the DisableOpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOpa

`func (o *ProjectActionVisibilityDto) SetDisableOpa(v ProjectButtonStatusDto)`

SetDisableOpa sets DisableOpa field to given value.

### HasDisableOpa

`func (o *ProjectActionVisibilityDto) HasDisableOpa() bool`

HasDisableOpa returns a boolean if a field has been set.

### GetEnableAutoscaler

`func (o *ProjectActionVisibilityDto) GetEnableAutoscaler() ProjectButtonStatusDto`

GetEnableAutoscaler returns the EnableAutoscaler field if non-nil, zero value otherwise.

### GetEnableAutoscalerOk

`func (o *ProjectActionVisibilityDto) GetEnableAutoscalerOk() (*ProjectButtonStatusDto, bool)`

GetEnableAutoscalerOk returns a tuple with the EnableAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoscaler

`func (o *ProjectActionVisibilityDto) SetEnableAutoscaler(v ProjectButtonStatusDto)`

SetEnableAutoscaler sets EnableAutoscaler field to given value.

### HasEnableAutoscaler

`func (o *ProjectActionVisibilityDto) HasEnableAutoscaler() bool`

HasEnableAutoscaler returns a boolean if a field has been set.

### GetDisableAutoscaler

`func (o *ProjectActionVisibilityDto) GetDisableAutoscaler() ProjectButtonStatusDto`

GetDisableAutoscaler returns the DisableAutoscaler field if non-nil, zero value otherwise.

### GetDisableAutoscalerOk

`func (o *ProjectActionVisibilityDto) GetDisableAutoscalerOk() (*ProjectButtonStatusDto, bool)`

GetDisableAutoscalerOk returns a tuple with the DisableAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoscaler

`func (o *ProjectActionVisibilityDto) SetDisableAutoscaler(v ProjectButtonStatusDto)`

SetDisableAutoscaler sets DisableAutoscaler field to given value.

### HasDisableAutoscaler

`func (o *ProjectActionVisibilityDto) HasDisableAutoscaler() bool`

HasDisableAutoscaler returns a boolean if a field has been set.

### GetVmRepair

`func (o *ProjectActionVisibilityDto) GetVmRepair() ProjectButtonStatusDto`

GetVmRepair returns the VmRepair field if non-nil, zero value otherwise.

### GetVmRepairOk

`func (o *ProjectActionVisibilityDto) GetVmRepairOk() (*ProjectButtonStatusDto, bool)`

GetVmRepairOk returns a tuple with the VmRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRepair

`func (o *ProjectActionVisibilityDto) SetVmRepair(v ProjectButtonStatusDto)`

SetVmRepair sets VmRepair field to given value.

### HasVmRepair

`func (o *ProjectActionVisibilityDto) HasVmRepair() bool`

HasVmRepair returns a boolean if a field has been set.

### GetVmCommit

`func (o *ProjectActionVisibilityDto) GetVmCommit() ProjectButtonStatusDto`

GetVmCommit returns the VmCommit field if non-nil, zero value otherwise.

### GetVmCommitOk

`func (o *ProjectActionVisibilityDto) GetVmCommitOk() (*ProjectButtonStatusDto, bool)`

GetVmCommitOk returns a tuple with the VmCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCommit

`func (o *ProjectActionVisibilityDto) SetVmCommit(v ProjectButtonStatusDto)`

SetVmCommit sets VmCommit field to given value.

### HasVmCommit

`func (o *ProjectActionVisibilityDto) HasVmCommit() bool`

HasVmCommit returns a boolean if a field has been set.

### GetLock

`func (o *ProjectActionVisibilityDto) GetLock() ProjectButtonStatusDto`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *ProjectActionVisibilityDto) GetLockOk() (*ProjectButtonStatusDto, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *ProjectActionVisibilityDto) SetLock(v ProjectButtonStatusDto)`

SetLock sets Lock field to given value.

### HasLock

`func (o *ProjectActionVisibilityDto) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetUnlock

`func (o *ProjectActionVisibilityDto) GetUnlock() ProjectButtonStatusDto`

GetUnlock returns the Unlock field if non-nil, zero value otherwise.

### GetUnlockOk

`func (o *ProjectActionVisibilityDto) GetUnlockOk() (*ProjectButtonStatusDto, bool)`

GetUnlockOk returns a tuple with the Unlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlock

`func (o *ProjectActionVisibilityDto) SetUnlock(v ProjectButtonStatusDto)`

SetUnlock sets Unlock field to given value.

### HasUnlock

`func (o *ProjectActionVisibilityDto) HasUnlock() bool`

HasUnlock returns a boolean if a field has been set.

### GetEnableSpotWorker

`func (o *ProjectActionVisibilityDto) GetEnableSpotWorker() ProjectButtonStatusDto`

GetEnableSpotWorker returns the EnableSpotWorker field if non-nil, zero value otherwise.

### GetEnableSpotWorkerOk

`func (o *ProjectActionVisibilityDto) GetEnableSpotWorkerOk() (*ProjectButtonStatusDto, bool)`

GetEnableSpotWorkerOk returns a tuple with the EnableSpotWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotWorker

`func (o *ProjectActionVisibilityDto) SetEnableSpotWorker(v ProjectButtonStatusDto)`

SetEnableSpotWorker sets EnableSpotWorker field to given value.

### HasEnableSpotWorker

`func (o *ProjectActionVisibilityDto) HasEnableSpotWorker() bool`

HasEnableSpotWorker returns a boolean if a field has been set.

### GetDisableSpotWorker

`func (o *ProjectActionVisibilityDto) GetDisableSpotWorker() ProjectButtonStatusDto`

GetDisableSpotWorker returns the DisableSpotWorker field if non-nil, zero value otherwise.

### GetDisableSpotWorkerOk

`func (o *ProjectActionVisibilityDto) GetDisableSpotWorkerOk() (*ProjectButtonStatusDto, bool)`

GetDisableSpotWorkerOk returns a tuple with the DisableSpotWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpotWorker

`func (o *ProjectActionVisibilityDto) SetDisableSpotWorker(v ProjectButtonStatusDto)`

SetDisableSpotWorker sets DisableSpotWorker field to given value.

### HasDisableSpotWorker

`func (o *ProjectActionVisibilityDto) HasDisableSpotWorker() bool`

HasDisableSpotWorker returns a boolean if a field has been set.

### GetEnableFullSpot

`func (o *ProjectActionVisibilityDto) GetEnableFullSpot() ProjectButtonStatusDto`

GetEnableFullSpot returns the EnableFullSpot field if non-nil, zero value otherwise.

### GetEnableFullSpotOk

`func (o *ProjectActionVisibilityDto) GetEnableFullSpotOk() (*ProjectButtonStatusDto, bool)`

GetEnableFullSpotOk returns a tuple with the EnableFullSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullSpot

`func (o *ProjectActionVisibilityDto) SetEnableFullSpot(v ProjectButtonStatusDto)`

SetEnableFullSpot sets EnableFullSpot field to given value.

### HasEnableFullSpot

`func (o *ProjectActionVisibilityDto) HasEnableFullSpot() bool`

HasEnableFullSpot returns a boolean if a field has been set.

### GetDisableFullSpot

`func (o *ProjectActionVisibilityDto) GetDisableFullSpot() ProjectButtonStatusDto`

GetDisableFullSpot returns the DisableFullSpot field if non-nil, zero value otherwise.

### GetDisableFullSpotOk

`func (o *ProjectActionVisibilityDto) GetDisableFullSpotOk() (*ProjectButtonStatusDto, bool)`

GetDisableFullSpotOk returns a tuple with the DisableFullSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFullSpot

`func (o *ProjectActionVisibilityDto) SetDisableFullSpot(v ProjectButtonStatusDto)`

SetDisableFullSpot sets DisableFullSpot field to given value.

### HasDisableFullSpot

`func (o *ProjectActionVisibilityDto) HasDisableFullSpot() bool`

HasDisableFullSpot returns a boolean if a field has been set.

### GetEnableSpotVm

`func (o *ProjectActionVisibilityDto) GetEnableSpotVm() ProjectButtonStatusDto`

GetEnableSpotVm returns the EnableSpotVm field if non-nil, zero value otherwise.

### GetEnableSpotVmOk

`func (o *ProjectActionVisibilityDto) GetEnableSpotVmOk() (*ProjectButtonStatusDto, bool)`

GetEnableSpotVmOk returns a tuple with the EnableSpotVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotVm

`func (o *ProjectActionVisibilityDto) SetEnableSpotVm(v ProjectButtonStatusDto)`

SetEnableSpotVm sets EnableSpotVm field to given value.

### HasEnableSpotVm

`func (o *ProjectActionVisibilityDto) HasEnableSpotVm() bool`

HasEnableSpotVm returns a boolean if a field has been set.

### GetDisableSpotVm

`func (o *ProjectActionVisibilityDto) GetDisableSpotVm() ProjectButtonStatusDto`

GetDisableSpotVm returns the DisableSpotVm field if non-nil, zero value otherwise.

### GetDisableSpotVmOk

`func (o *ProjectActionVisibilityDto) GetDisableSpotVmOk() (*ProjectButtonStatusDto, bool)`

GetDisableSpotVmOk returns a tuple with the DisableSpotVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpotVm

`func (o *ProjectActionVisibilityDto) SetDisableSpotVm(v ProjectButtonStatusDto)`

SetDisableSpotVm sets DisableSpotVm field to given value.

### HasDisableSpotVm

`func (o *ProjectActionVisibilityDto) HasDisableSpotVm() bool`

HasDisableSpotVm returns a boolean if a field has been set.

### GetAttachAlertingProfile

`func (o *ProjectActionVisibilityDto) GetAttachAlertingProfile() ProjectButtonStatusDto`

GetAttachAlertingProfile returns the AttachAlertingProfile field if non-nil, zero value otherwise.

### GetAttachAlertingProfileOk

`func (o *ProjectActionVisibilityDto) GetAttachAlertingProfileOk() (*ProjectButtonStatusDto, bool)`

GetAttachAlertingProfileOk returns a tuple with the AttachAlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlertingProfile

`func (o *ProjectActionVisibilityDto) SetAttachAlertingProfile(v ProjectButtonStatusDto)`

SetAttachAlertingProfile sets AttachAlertingProfile field to given value.

### HasAttachAlertingProfile

`func (o *ProjectActionVisibilityDto) HasAttachAlertingProfile() bool`

HasAttachAlertingProfile returns a boolean if a field has been set.

### GetDetachAlertingProfile

`func (o *ProjectActionVisibilityDto) GetDetachAlertingProfile() ProjectButtonStatusDto`

GetDetachAlertingProfile returns the DetachAlertingProfile field if non-nil, zero value otherwise.

### GetDetachAlertingProfileOk

`func (o *ProjectActionVisibilityDto) GetDetachAlertingProfileOk() (*ProjectButtonStatusDto, bool)`

GetDetachAlertingProfileOk returns a tuple with the DetachAlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachAlertingProfile

`func (o *ProjectActionVisibilityDto) SetDetachAlertingProfile(v ProjectButtonStatusDto)`

SetDetachAlertingProfile sets DetachAlertingProfile field to given value.

### HasDetachAlertingProfile

`func (o *ProjectActionVisibilityDto) HasDetachAlertingProfile() bool`

HasDetachAlertingProfile returns a boolean if a field has been set.

### GetEnableAi

`func (o *ProjectActionVisibilityDto) GetEnableAi() ProjectButtonStatusDto`

GetEnableAi returns the EnableAi field if non-nil, zero value otherwise.

### GetEnableAiOk

`func (o *ProjectActionVisibilityDto) GetEnableAiOk() (*ProjectButtonStatusDto, bool)`

GetEnableAiOk returns a tuple with the EnableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAi

`func (o *ProjectActionVisibilityDto) SetEnableAi(v ProjectButtonStatusDto)`

SetEnableAi sets EnableAi field to given value.

### HasEnableAi

`func (o *ProjectActionVisibilityDto) HasEnableAi() bool`

HasEnableAi returns a boolean if a field has been set.

### GetDisableAi

`func (o *ProjectActionVisibilityDto) GetDisableAi() ProjectButtonStatusDto`

GetDisableAi returns the DisableAi field if non-nil, zero value otherwise.

### GetDisableAiOk

`func (o *ProjectActionVisibilityDto) GetDisableAiOk() (*ProjectButtonStatusDto, bool)`

GetDisableAiOk returns a tuple with the DisableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAi

`func (o *ProjectActionVisibilityDto) SetDisableAi(v ProjectButtonStatusDto)`

SetDisableAi sets DisableAi field to given value.

### HasDisableAi

`func (o *ProjectActionVisibilityDto) HasDisableAi() bool`

HasDisableAi returns a boolean if a field has been set.

### GetAiAssistant

`func (o *ProjectActionVisibilityDto) GetAiAssistant() ProjectButtonStatusDto`

GetAiAssistant returns the AiAssistant field if non-nil, zero value otherwise.

### GetAiAssistantOk

`func (o *ProjectActionVisibilityDto) GetAiAssistantOk() (*ProjectButtonStatusDto, bool)`

GetAiAssistantOk returns a tuple with the AiAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAssistant

`func (o *ProjectActionVisibilityDto) SetAiAssistant(v ProjectButtonStatusDto)`

SetAiAssistant sets AiAssistant field to given value.

### HasAiAssistant

`func (o *ProjectActionVisibilityDto) HasAiAssistant() bool`

HasAiAssistant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


