# OpaProfileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ForbidNodePort** | Pointer to **bool** |  | [optional] 
**ForbidHttpIngress** | Pointer to **bool** |  | [optional] 
**RequireProbe** | Pointer to **bool** |  | [optional] 
**UniqueIngresses** | Pointer to **bool** |  | [optional] 
**UniqueServiceSelector** | Pointer to **bool** |  | [optional] 
**AllowedRepo** | Pointer to **[]string** |  | [optional] 
**ForbidSpecificTags** | Pointer to **[]string** |  | [optional] 
**IngressWhitelist** | Pointer to **[]string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Projects** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 

## Methods

### NewOpaProfileListDto

`func NewOpaProfileListDto() *OpaProfileListDto`

NewOpaProfileListDto instantiates a new OpaProfileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaProfileListDtoWithDefaults

`func NewOpaProfileListDtoWithDefaults() *OpaProfileListDto`

NewOpaProfileListDtoWithDefaults instantiates a new OpaProfileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpaProfileListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaProfileListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaProfileListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpaProfileListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OpaProfileListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpaProfileListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpaProfileListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpaProfileListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpaProfileListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpaProfileListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForbidNodePort

`func (o *OpaProfileListDto) GetForbidNodePort() bool`

GetForbidNodePort returns the ForbidNodePort field if non-nil, zero value otherwise.

### GetForbidNodePortOk

`func (o *OpaProfileListDto) GetForbidNodePortOk() (*bool, bool)`

GetForbidNodePortOk returns a tuple with the ForbidNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidNodePort

`func (o *OpaProfileListDto) SetForbidNodePort(v bool)`

SetForbidNodePort sets ForbidNodePort field to given value.

### HasForbidNodePort

`func (o *OpaProfileListDto) HasForbidNodePort() bool`

HasForbidNodePort returns a boolean if a field has been set.

### GetForbidHttpIngress

`func (o *OpaProfileListDto) GetForbidHttpIngress() bool`

GetForbidHttpIngress returns the ForbidHttpIngress field if non-nil, zero value otherwise.

### GetForbidHttpIngressOk

`func (o *OpaProfileListDto) GetForbidHttpIngressOk() (*bool, bool)`

GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidHttpIngress

`func (o *OpaProfileListDto) SetForbidHttpIngress(v bool)`

SetForbidHttpIngress sets ForbidHttpIngress field to given value.

### HasForbidHttpIngress

`func (o *OpaProfileListDto) HasForbidHttpIngress() bool`

HasForbidHttpIngress returns a boolean if a field has been set.

### GetRequireProbe

`func (o *OpaProfileListDto) GetRequireProbe() bool`

GetRequireProbe returns the RequireProbe field if non-nil, zero value otherwise.

### GetRequireProbeOk

`func (o *OpaProfileListDto) GetRequireProbeOk() (*bool, bool)`

GetRequireProbeOk returns a tuple with the RequireProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProbe

`func (o *OpaProfileListDto) SetRequireProbe(v bool)`

SetRequireProbe sets RequireProbe field to given value.

### HasRequireProbe

`func (o *OpaProfileListDto) HasRequireProbe() bool`

HasRequireProbe returns a boolean if a field has been set.

### GetUniqueIngresses

`func (o *OpaProfileListDto) GetUniqueIngresses() bool`

GetUniqueIngresses returns the UniqueIngresses field if non-nil, zero value otherwise.

### GetUniqueIngressesOk

`func (o *OpaProfileListDto) GetUniqueIngressesOk() (*bool, bool)`

GetUniqueIngressesOk returns a tuple with the UniqueIngresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIngresses

`func (o *OpaProfileListDto) SetUniqueIngresses(v bool)`

SetUniqueIngresses sets UniqueIngresses field to given value.

### HasUniqueIngresses

`func (o *OpaProfileListDto) HasUniqueIngresses() bool`

HasUniqueIngresses returns a boolean if a field has been set.

### GetUniqueServiceSelector

`func (o *OpaProfileListDto) GetUniqueServiceSelector() bool`

GetUniqueServiceSelector returns the UniqueServiceSelector field if non-nil, zero value otherwise.

### GetUniqueServiceSelectorOk

`func (o *OpaProfileListDto) GetUniqueServiceSelectorOk() (*bool, bool)`

GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueServiceSelector

`func (o *OpaProfileListDto) SetUniqueServiceSelector(v bool)`

SetUniqueServiceSelector sets UniqueServiceSelector field to given value.

### HasUniqueServiceSelector

`func (o *OpaProfileListDto) HasUniqueServiceSelector() bool`

HasUniqueServiceSelector returns a boolean if a field has been set.

### GetAllowedRepo

`func (o *OpaProfileListDto) GetAllowedRepo() []string`

GetAllowedRepo returns the AllowedRepo field if non-nil, zero value otherwise.

### GetAllowedRepoOk

`func (o *OpaProfileListDto) GetAllowedRepoOk() (*[]string, bool)`

GetAllowedRepoOk returns a tuple with the AllowedRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRepo

`func (o *OpaProfileListDto) SetAllowedRepo(v []string)`

SetAllowedRepo sets AllowedRepo field to given value.

### HasAllowedRepo

`func (o *OpaProfileListDto) HasAllowedRepo() bool`

HasAllowedRepo returns a boolean if a field has been set.

### SetAllowedRepoNil

`func (o *OpaProfileListDto) SetAllowedRepoNil(b bool)`

 SetAllowedRepoNil sets the value for AllowedRepo to be an explicit nil

### UnsetAllowedRepo
`func (o *OpaProfileListDto) UnsetAllowedRepo()`

UnsetAllowedRepo ensures that no value is present for AllowedRepo, not even an explicit nil
### GetForbidSpecificTags

`func (o *OpaProfileListDto) GetForbidSpecificTags() []string`

GetForbidSpecificTags returns the ForbidSpecificTags field if non-nil, zero value otherwise.

### GetForbidSpecificTagsOk

`func (o *OpaProfileListDto) GetForbidSpecificTagsOk() (*[]string, bool)`

GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidSpecificTags

`func (o *OpaProfileListDto) SetForbidSpecificTags(v []string)`

SetForbidSpecificTags sets ForbidSpecificTags field to given value.

### HasForbidSpecificTags

`func (o *OpaProfileListDto) HasForbidSpecificTags() bool`

HasForbidSpecificTags returns a boolean if a field has been set.

### SetForbidSpecificTagsNil

`func (o *OpaProfileListDto) SetForbidSpecificTagsNil(b bool)`

 SetForbidSpecificTagsNil sets the value for ForbidSpecificTags to be an explicit nil

### UnsetForbidSpecificTags
`func (o *OpaProfileListDto) UnsetForbidSpecificTags()`

UnsetForbidSpecificTags ensures that no value is present for ForbidSpecificTags, not even an explicit nil
### GetIngressWhitelist

`func (o *OpaProfileListDto) GetIngressWhitelist() []string`

GetIngressWhitelist returns the IngressWhitelist field if non-nil, zero value otherwise.

### GetIngressWhitelistOk

`func (o *OpaProfileListDto) GetIngressWhitelistOk() (*[]string, bool)`

GetIngressWhitelistOk returns a tuple with the IngressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressWhitelist

`func (o *OpaProfileListDto) SetIngressWhitelist(v []string)`

SetIngressWhitelist sets IngressWhitelist field to given value.

### HasIngressWhitelist

`func (o *OpaProfileListDto) HasIngressWhitelist() bool`

HasIngressWhitelist returns a boolean if a field has been set.

### SetIngressWhitelistNil

`func (o *OpaProfileListDto) SetIngressWhitelistNil(b bool)`

 SetIngressWhitelistNil sets the value for IngressWhitelist to be an explicit nil

### UnsetIngressWhitelist
`func (o *OpaProfileListDto) UnsetIngressWhitelist()`

UnsetIngressWhitelist ensures that no value is present for IngressWhitelist, not even an explicit nil
### GetIsLocked

`func (o *OpaProfileListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OpaProfileListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OpaProfileListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *OpaProfileListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetRevision

`func (o *OpaProfileListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OpaProfileListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OpaProfileListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *OpaProfileListDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OpaProfileListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpaProfileListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpaProfileListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OpaProfileListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OpaProfileListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OpaProfileListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OpaProfileListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OpaProfileListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OpaProfileListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OpaProfileListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCreatedAt

`func (o *OpaProfileListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpaProfileListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpaProfileListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpaProfileListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *OpaProfileListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OpaProfileListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsDefault

`func (o *OpaProfileListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OpaProfileListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OpaProfileListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OpaProfileListDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProjects

`func (o *OpaProfileListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OpaProfileListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OpaProfileListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OpaProfileListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *OpaProfileListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *OpaProfileListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


