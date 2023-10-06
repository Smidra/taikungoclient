# KubernetesAlertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**StartsAt** | Pointer to **NullableString** |  | [optional] 
**EndAt** | Pointer to **NullableString** |  | [optional] 
**IsSolved** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**IsSilenced** | Pointer to **bool** |  | [optional] 
**SilenceReason** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKubernetesAlertDto

`func NewKubernetesAlertDto() *KubernetesAlertDto`

NewKubernetesAlertDto instantiates a new KubernetesAlertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAlertDtoWithDefaults

`func NewKubernetesAlertDtoWithDefaults() *KubernetesAlertDto`

NewKubernetesAlertDtoWithDefaults instantiates a new KubernetesAlertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesAlertDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesAlertDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesAlertDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesAlertDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesAlertDto) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesAlertDto) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesAlertDto) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesAlertDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesAlertDto) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesAlertDto) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *KubernetesAlertDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesAlertDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesAlertDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesAlertDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KubernetesAlertDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KubernetesAlertDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *KubernetesAlertDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KubernetesAlertDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KubernetesAlertDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KubernetesAlertDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *KubernetesAlertDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *KubernetesAlertDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSeverity

`func (o *KubernetesAlertDto) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *KubernetesAlertDto) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *KubernetesAlertDto) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *KubernetesAlertDto) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *KubernetesAlertDto) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *KubernetesAlertDto) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetFingerprint

`func (o *KubernetesAlertDto) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KubernetesAlertDto) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KubernetesAlertDto) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KubernetesAlertDto) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *KubernetesAlertDto) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *KubernetesAlertDto) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetStatus

`func (o *KubernetesAlertDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesAlertDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesAlertDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesAlertDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KubernetesAlertDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KubernetesAlertDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartsAt

`func (o *KubernetesAlertDto) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *KubernetesAlertDto) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *KubernetesAlertDto) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *KubernetesAlertDto) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *KubernetesAlertDto) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *KubernetesAlertDto) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetEndAt

`func (o *KubernetesAlertDto) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *KubernetesAlertDto) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *KubernetesAlertDto) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *KubernetesAlertDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *KubernetesAlertDto) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *KubernetesAlertDto) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetIsSolved

`func (o *KubernetesAlertDto) GetIsSolved() bool`

GetIsSolved returns the IsSolved field if non-nil, zero value otherwise.

### GetIsSolvedOk

`func (o *KubernetesAlertDto) GetIsSolvedOk() (*bool, bool)`

GetIsSolvedOk returns a tuple with the IsSolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSolved

`func (o *KubernetesAlertDto) SetIsSolved(v bool)`

SetIsSolved sets IsSolved field to given value.

### HasIsSolved

`func (o *KubernetesAlertDto) HasIsSolved() bool`

HasIsSolved returns a boolean if a field has been set.

### GetProjectId

`func (o *KubernetesAlertDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubernetesAlertDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubernetesAlertDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KubernetesAlertDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *KubernetesAlertDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KubernetesAlertDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KubernetesAlertDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *KubernetesAlertDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *KubernetesAlertDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *KubernetesAlertDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsSilenced

`func (o *KubernetesAlertDto) GetIsSilenced() bool`

GetIsSilenced returns the IsSilenced field if non-nil, zero value otherwise.

### GetIsSilencedOk

`func (o *KubernetesAlertDto) GetIsSilencedOk() (*bool, bool)`

GetIsSilencedOk returns a tuple with the IsSilenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSilenced

`func (o *KubernetesAlertDto) SetIsSilenced(v bool)`

SetIsSilenced sets IsSilenced field to given value.

### HasIsSilenced

`func (o *KubernetesAlertDto) HasIsSilenced() bool`

HasIsSilenced returns a boolean if a field has been set.

### GetSilenceReason

`func (o *KubernetesAlertDto) GetSilenceReason() string`

GetSilenceReason returns the SilenceReason field if non-nil, zero value otherwise.

### GetSilenceReasonOk

`func (o *KubernetesAlertDto) GetSilenceReasonOk() (*string, bool)`

GetSilenceReasonOk returns a tuple with the SilenceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilenceReason

`func (o *KubernetesAlertDto) SetSilenceReason(v string)`

SetSilenceReason sets SilenceReason field to given value.

### HasSilenceReason

`func (o *KubernetesAlertDto) HasSilenceReason() bool`

HasSilenceReason returns a boolean if a field has been set.

### SetSilenceReasonNil

`func (o *KubernetesAlertDto) SetSilenceReasonNil(b bool)`

 SetSilenceReasonNil sets the value for SilenceReason to be an explicit nil

### UnsetSilenceReason
`func (o *KubernetesAlertDto) UnsetSilenceReason()`

UnsetSilenceReason ensures that no value is present for SilenceReason, not even an explicit nil
### GetLastModifiedBy

`func (o *KubernetesAlertDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *KubernetesAlertDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *KubernetesAlertDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *KubernetesAlertDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *KubernetesAlertDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *KubernetesAlertDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


