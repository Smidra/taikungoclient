# ProjectDetailsForVmsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectStatus** | Pointer to [**ProjectStatus**](ProjectStatus.md) |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**CloudType** | Pointer to [**CloudType**](CloudType.md) |  | [optional] 
**CloudName** | Pointer to **NullableString** |  | [optional] 
**CloudId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**HasSelectedFlavors** | Pointer to **NullableBool** |  | [optional] 
**IsMaintenanceModeEnabled** | Pointer to **bool** |  | [optional] 
**ProjectCloudRevision** | Pointer to **NullableInt32** |  | [optional] 
**CloudCredentialRevision** | Pointer to **NullableInt32** |  | [optional] 
**AllowFullSpotKubernetes** | Pointer to **bool** |  | [optional] 
**AllowSpotWorkers** | Pointer to **bool** |  | [optional] 
**AllowSpotVMs** | Pointer to **bool** |  | [optional] 
**MaxSpotPrice** | Pointer to **NullableFloat64** |  | [optional] 
**TotalHourlyCost** | Pointer to **float64** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**Hypervisors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectDetailsForVmsDto

`func NewProjectDetailsForVmsDto() *ProjectDetailsForVmsDto`

NewProjectDetailsForVmsDto instantiates a new ProjectDetailsForVmsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsForVmsDtoWithDefaults

`func NewProjectDetailsForVmsDtoWithDefaults() *ProjectDetailsForVmsDto`

NewProjectDetailsForVmsDtoWithDefaults instantiates a new ProjectDetailsForVmsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectStatus

`func (o *ProjectDetailsForVmsDto) GetProjectStatus() ProjectStatus`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *ProjectDetailsForVmsDto) GetProjectStatusOk() (*ProjectStatus, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *ProjectDetailsForVmsDto) SetProjectStatus(v ProjectStatus)`

SetProjectStatus sets ProjectStatus field to given value.

### HasProjectStatus

`func (o *ProjectDetailsForVmsDto) HasProjectStatus() bool`

HasProjectStatus returns a boolean if a field has been set.

### GetProjectName

`func (o *ProjectDetailsForVmsDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectDetailsForVmsDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectDetailsForVmsDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectDetailsForVmsDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ProjectDetailsForVmsDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectDetailsForVmsDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetProjectId

`func (o *ProjectDetailsForVmsDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectDetailsForVmsDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectDetailsForVmsDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectDetailsForVmsDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCloudType

`func (o *ProjectDetailsForVmsDto) GetCloudType() CloudType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectDetailsForVmsDto) GetCloudTypeOk() (*CloudType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectDetailsForVmsDto) SetCloudType(v CloudType)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ProjectDetailsForVmsDto) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudName

`func (o *ProjectDetailsForVmsDto) GetCloudName() string`

GetCloudName returns the CloudName field if non-nil, zero value otherwise.

### GetCloudNameOk

`func (o *ProjectDetailsForVmsDto) GetCloudNameOk() (*string, bool)`

GetCloudNameOk returns a tuple with the CloudName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudName

`func (o *ProjectDetailsForVmsDto) SetCloudName(v string)`

SetCloudName sets CloudName field to given value.

### HasCloudName

`func (o *ProjectDetailsForVmsDto) HasCloudName() bool`

HasCloudName returns a boolean if a field has been set.

### SetCloudNameNil

`func (o *ProjectDetailsForVmsDto) SetCloudNameNil(b bool)`

 SetCloudNameNil sets the value for CloudName to be an explicit nil

### UnsetCloudName
`func (o *ProjectDetailsForVmsDto) UnsetCloudName()`

UnsetCloudName ensures that no value is present for CloudName, not even an explicit nil
### GetCloudId

`func (o *ProjectDetailsForVmsDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ProjectDetailsForVmsDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ProjectDetailsForVmsDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *ProjectDetailsForVmsDto) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### SetCloudIdNil

`func (o *ProjectDetailsForVmsDto) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *ProjectDetailsForVmsDto) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetOrganizationName

`func (o *ProjectDetailsForVmsDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectDetailsForVmsDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectDetailsForVmsDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ProjectDetailsForVmsDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ProjectDetailsForVmsDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectDetailsForVmsDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ProjectDetailsForVmsDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectDetailsForVmsDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectDetailsForVmsDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectDetailsForVmsDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsLocked

`func (o *ProjectDetailsForVmsDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectDetailsForVmsDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectDetailsForVmsDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectDetailsForVmsDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetHasSelectedFlavors

`func (o *ProjectDetailsForVmsDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectDetailsForVmsDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectDetailsForVmsDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.

### HasHasSelectedFlavors

`func (o *ProjectDetailsForVmsDto) HasHasSelectedFlavors() bool`

HasHasSelectedFlavors returns a boolean if a field has been set.

### SetHasSelectedFlavorsNil

`func (o *ProjectDetailsForVmsDto) SetHasSelectedFlavorsNil(b bool)`

 SetHasSelectedFlavorsNil sets the value for HasSelectedFlavors to be an explicit nil

### UnsetHasSelectedFlavors
`func (o *ProjectDetailsForVmsDto) UnsetHasSelectedFlavors()`

UnsetHasSelectedFlavors ensures that no value is present for HasSelectedFlavors, not even an explicit nil
### GetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ProjectDetailsForVmsDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.

### HasIsMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) HasIsMaintenanceModeEnabled() bool`

HasIsMaintenanceModeEnabled returns a boolean if a field has been set.

### GetProjectCloudRevision

`func (o *ProjectDetailsForVmsDto) GetProjectCloudRevision() int32`

GetProjectCloudRevision returns the ProjectCloudRevision field if non-nil, zero value otherwise.

### GetProjectCloudRevisionOk

`func (o *ProjectDetailsForVmsDto) GetProjectCloudRevisionOk() (*int32, bool)`

GetProjectCloudRevisionOk returns a tuple with the ProjectCloudRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCloudRevision

`func (o *ProjectDetailsForVmsDto) SetProjectCloudRevision(v int32)`

SetProjectCloudRevision sets ProjectCloudRevision field to given value.

### HasProjectCloudRevision

`func (o *ProjectDetailsForVmsDto) HasProjectCloudRevision() bool`

HasProjectCloudRevision returns a boolean if a field has been set.

### SetProjectCloudRevisionNil

`func (o *ProjectDetailsForVmsDto) SetProjectCloudRevisionNil(b bool)`

 SetProjectCloudRevisionNil sets the value for ProjectCloudRevision to be an explicit nil

### UnsetProjectCloudRevision
`func (o *ProjectDetailsForVmsDto) UnsetProjectCloudRevision()`

UnsetProjectCloudRevision ensures that no value is present for ProjectCloudRevision, not even an explicit nil
### GetCloudCredentialRevision

`func (o *ProjectDetailsForVmsDto) GetCloudCredentialRevision() int32`

GetCloudCredentialRevision returns the CloudCredentialRevision field if non-nil, zero value otherwise.

### GetCloudCredentialRevisionOk

`func (o *ProjectDetailsForVmsDto) GetCloudCredentialRevisionOk() (*int32, bool)`

GetCloudCredentialRevisionOk returns a tuple with the CloudCredentialRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialRevision

`func (o *ProjectDetailsForVmsDto) SetCloudCredentialRevision(v int32)`

SetCloudCredentialRevision sets CloudCredentialRevision field to given value.

### HasCloudCredentialRevision

`func (o *ProjectDetailsForVmsDto) HasCloudCredentialRevision() bool`

HasCloudCredentialRevision returns a boolean if a field has been set.

### SetCloudCredentialRevisionNil

`func (o *ProjectDetailsForVmsDto) SetCloudCredentialRevisionNil(b bool)`

 SetCloudCredentialRevisionNil sets the value for CloudCredentialRevision to be an explicit nil

### UnsetCloudCredentialRevision
`func (o *ProjectDetailsForVmsDto) UnsetCloudCredentialRevision()`

UnsetCloudCredentialRevision ensures that no value is present for CloudCredentialRevision, not even an explicit nil
### GetAllowFullSpotKubernetes

`func (o *ProjectDetailsForVmsDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectDetailsForVmsDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectDetailsForVmsDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.

### HasAllowFullSpotKubernetes

`func (o *ProjectDetailsForVmsDto) HasAllowFullSpotKubernetes() bool`

HasAllowFullSpotKubernetes returns a boolean if a field has been set.

### GetAllowSpotWorkers

`func (o *ProjectDetailsForVmsDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectDetailsForVmsDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectDetailsForVmsDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.

### HasAllowSpotWorkers

`func (o *ProjectDetailsForVmsDto) HasAllowSpotWorkers() bool`

HasAllowSpotWorkers returns a boolean if a field has been set.

### GetAllowSpotVMs

`func (o *ProjectDetailsForVmsDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectDetailsForVmsDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectDetailsForVmsDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.

### HasAllowSpotVMs

`func (o *ProjectDetailsForVmsDto) HasAllowSpotVMs() bool`

HasAllowSpotVMs returns a boolean if a field has been set.

### GetMaxSpotPrice

`func (o *ProjectDetailsForVmsDto) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *ProjectDetailsForVmsDto) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *ProjectDetailsForVmsDto) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.

### HasMaxSpotPrice

`func (o *ProjectDetailsForVmsDto) HasMaxSpotPrice() bool`

HasMaxSpotPrice returns a boolean if a field has been set.

### SetMaxSpotPriceNil

`func (o *ProjectDetailsForVmsDto) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *ProjectDetailsForVmsDto) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetTotalHourlyCost

`func (o *ProjectDetailsForVmsDto) GetTotalHourlyCost() float64`

GetTotalHourlyCost returns the TotalHourlyCost field if non-nil, zero value otherwise.

### GetTotalHourlyCostOk

`func (o *ProjectDetailsForVmsDto) GetTotalHourlyCostOk() (*float64, bool)`

GetTotalHourlyCostOk returns a tuple with the TotalHourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHourlyCost

`func (o *ProjectDetailsForVmsDto) SetTotalHourlyCost(v float64)`

SetTotalHourlyCost sets TotalHourlyCost field to given value.

### HasTotalHourlyCost

`func (o *ProjectDetailsForVmsDto) HasTotalHourlyCost() bool`

HasTotalHourlyCost returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *ProjectDetailsForVmsDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectDetailsForVmsDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectDetailsForVmsDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *ProjectDetailsForVmsDto) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *ProjectDetailsForVmsDto) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *ProjectDetailsForVmsDto) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetHypervisors

`func (o *ProjectDetailsForVmsDto) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *ProjectDetailsForVmsDto) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *ProjectDetailsForVmsDto) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *ProjectDetailsForVmsDto) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.

### SetHypervisorsNil

`func (o *ProjectDetailsForVmsDto) SetHypervisorsNil(b bool)`

 SetHypervisorsNil sets the value for Hypervisors to be an explicit nil

### UnsetHypervisors
`func (o *ProjectDetailsForVmsDto) UnsetHypervisors()`

UnsetHypervisors ensures that no value is present for Hypervisors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


