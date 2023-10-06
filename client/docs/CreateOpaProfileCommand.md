# CreateOpaProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ForbidNodePort** | Pointer to **bool** |  | [optional] 
**ForbidHttpIngress** | Pointer to **bool** |  | [optional] 
**RequireProbe** | Pointer to **bool** |  | [optional] 
**UniqueIngresses** | Pointer to **bool** |  | [optional] 
**UniqueServiceSelector** | Pointer to **bool** |  | [optional] 
**ForcePodResource** | Pointer to **bool** |  | [optional] 
**AllowedRepo** | Pointer to **[]string** |  | [optional] 
**ForbidSpecificTags** | Pointer to **[]string** |  | [optional] 
**IngressWhitelist** | Pointer to **[]string** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateOpaProfileCommand

`func NewCreateOpaProfileCommand() *CreateOpaProfileCommand`

NewCreateOpaProfileCommand instantiates a new CreateOpaProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpaProfileCommandWithDefaults

`func NewCreateOpaProfileCommandWithDefaults() *CreateOpaProfileCommand`

NewCreateOpaProfileCommandWithDefaults instantiates a new CreateOpaProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOpaProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOpaProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOpaProfileCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOpaProfileCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateOpaProfileCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOpaProfileCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForbidNodePort

`func (o *CreateOpaProfileCommand) GetForbidNodePort() bool`

GetForbidNodePort returns the ForbidNodePort field if non-nil, zero value otherwise.

### GetForbidNodePortOk

`func (o *CreateOpaProfileCommand) GetForbidNodePortOk() (*bool, bool)`

GetForbidNodePortOk returns a tuple with the ForbidNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidNodePort

`func (o *CreateOpaProfileCommand) SetForbidNodePort(v bool)`

SetForbidNodePort sets ForbidNodePort field to given value.

### HasForbidNodePort

`func (o *CreateOpaProfileCommand) HasForbidNodePort() bool`

HasForbidNodePort returns a boolean if a field has been set.

### GetForbidHttpIngress

`func (o *CreateOpaProfileCommand) GetForbidHttpIngress() bool`

GetForbidHttpIngress returns the ForbidHttpIngress field if non-nil, zero value otherwise.

### GetForbidHttpIngressOk

`func (o *CreateOpaProfileCommand) GetForbidHttpIngressOk() (*bool, bool)`

GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidHttpIngress

`func (o *CreateOpaProfileCommand) SetForbidHttpIngress(v bool)`

SetForbidHttpIngress sets ForbidHttpIngress field to given value.

### HasForbidHttpIngress

`func (o *CreateOpaProfileCommand) HasForbidHttpIngress() bool`

HasForbidHttpIngress returns a boolean if a field has been set.

### GetRequireProbe

`func (o *CreateOpaProfileCommand) GetRequireProbe() bool`

GetRequireProbe returns the RequireProbe field if non-nil, zero value otherwise.

### GetRequireProbeOk

`func (o *CreateOpaProfileCommand) GetRequireProbeOk() (*bool, bool)`

GetRequireProbeOk returns a tuple with the RequireProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProbe

`func (o *CreateOpaProfileCommand) SetRequireProbe(v bool)`

SetRequireProbe sets RequireProbe field to given value.

### HasRequireProbe

`func (o *CreateOpaProfileCommand) HasRequireProbe() bool`

HasRequireProbe returns a boolean if a field has been set.

### GetUniqueIngresses

`func (o *CreateOpaProfileCommand) GetUniqueIngresses() bool`

GetUniqueIngresses returns the UniqueIngresses field if non-nil, zero value otherwise.

### GetUniqueIngressesOk

`func (o *CreateOpaProfileCommand) GetUniqueIngressesOk() (*bool, bool)`

GetUniqueIngressesOk returns a tuple with the UniqueIngresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIngresses

`func (o *CreateOpaProfileCommand) SetUniqueIngresses(v bool)`

SetUniqueIngresses sets UniqueIngresses field to given value.

### HasUniqueIngresses

`func (o *CreateOpaProfileCommand) HasUniqueIngresses() bool`

HasUniqueIngresses returns a boolean if a field has been set.

### GetUniqueServiceSelector

`func (o *CreateOpaProfileCommand) GetUniqueServiceSelector() bool`

GetUniqueServiceSelector returns the UniqueServiceSelector field if non-nil, zero value otherwise.

### GetUniqueServiceSelectorOk

`func (o *CreateOpaProfileCommand) GetUniqueServiceSelectorOk() (*bool, bool)`

GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueServiceSelector

`func (o *CreateOpaProfileCommand) SetUniqueServiceSelector(v bool)`

SetUniqueServiceSelector sets UniqueServiceSelector field to given value.

### HasUniqueServiceSelector

`func (o *CreateOpaProfileCommand) HasUniqueServiceSelector() bool`

HasUniqueServiceSelector returns a boolean if a field has been set.

### GetForcePodResource

`func (o *CreateOpaProfileCommand) GetForcePodResource() bool`

GetForcePodResource returns the ForcePodResource field if non-nil, zero value otherwise.

### GetForcePodResourceOk

`func (o *CreateOpaProfileCommand) GetForcePodResourceOk() (*bool, bool)`

GetForcePodResourceOk returns a tuple with the ForcePodResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePodResource

`func (o *CreateOpaProfileCommand) SetForcePodResource(v bool)`

SetForcePodResource sets ForcePodResource field to given value.

### HasForcePodResource

`func (o *CreateOpaProfileCommand) HasForcePodResource() bool`

HasForcePodResource returns a boolean if a field has been set.

### GetAllowedRepo

`func (o *CreateOpaProfileCommand) GetAllowedRepo() []string`

GetAllowedRepo returns the AllowedRepo field if non-nil, zero value otherwise.

### GetAllowedRepoOk

`func (o *CreateOpaProfileCommand) GetAllowedRepoOk() (*[]string, bool)`

GetAllowedRepoOk returns a tuple with the AllowedRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRepo

`func (o *CreateOpaProfileCommand) SetAllowedRepo(v []string)`

SetAllowedRepo sets AllowedRepo field to given value.

### HasAllowedRepo

`func (o *CreateOpaProfileCommand) HasAllowedRepo() bool`

HasAllowedRepo returns a boolean if a field has been set.

### SetAllowedRepoNil

`func (o *CreateOpaProfileCommand) SetAllowedRepoNil(b bool)`

 SetAllowedRepoNil sets the value for AllowedRepo to be an explicit nil

### UnsetAllowedRepo
`func (o *CreateOpaProfileCommand) UnsetAllowedRepo()`

UnsetAllowedRepo ensures that no value is present for AllowedRepo, not even an explicit nil
### GetForbidSpecificTags

`func (o *CreateOpaProfileCommand) GetForbidSpecificTags() []string`

GetForbidSpecificTags returns the ForbidSpecificTags field if non-nil, zero value otherwise.

### GetForbidSpecificTagsOk

`func (o *CreateOpaProfileCommand) GetForbidSpecificTagsOk() (*[]string, bool)`

GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidSpecificTags

`func (o *CreateOpaProfileCommand) SetForbidSpecificTags(v []string)`

SetForbidSpecificTags sets ForbidSpecificTags field to given value.

### HasForbidSpecificTags

`func (o *CreateOpaProfileCommand) HasForbidSpecificTags() bool`

HasForbidSpecificTags returns a boolean if a field has been set.

### SetForbidSpecificTagsNil

`func (o *CreateOpaProfileCommand) SetForbidSpecificTagsNil(b bool)`

 SetForbidSpecificTagsNil sets the value for ForbidSpecificTags to be an explicit nil

### UnsetForbidSpecificTags
`func (o *CreateOpaProfileCommand) UnsetForbidSpecificTags()`

UnsetForbidSpecificTags ensures that no value is present for ForbidSpecificTags, not even an explicit nil
### GetIngressWhitelist

`func (o *CreateOpaProfileCommand) GetIngressWhitelist() []string`

GetIngressWhitelist returns the IngressWhitelist field if non-nil, zero value otherwise.

### GetIngressWhitelistOk

`func (o *CreateOpaProfileCommand) GetIngressWhitelistOk() (*[]string, bool)`

GetIngressWhitelistOk returns a tuple with the IngressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressWhitelist

`func (o *CreateOpaProfileCommand) SetIngressWhitelist(v []string)`

SetIngressWhitelist sets IngressWhitelist field to given value.

### HasIngressWhitelist

`func (o *CreateOpaProfileCommand) HasIngressWhitelist() bool`

HasIngressWhitelist returns a boolean if a field has been set.

### SetIngressWhitelistNil

`func (o *CreateOpaProfileCommand) SetIngressWhitelistNil(b bool)`

 SetIngressWhitelistNil sets the value for IngressWhitelist to be an explicit nil

### UnsetIngressWhitelist
`func (o *CreateOpaProfileCommand) UnsetIngressWhitelist()`

UnsetIngressWhitelist ensures that no value is present for IngressWhitelist, not even an explicit nil
### GetOrganizationId

`func (o *CreateOpaProfileCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOpaProfileCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOpaProfileCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOpaProfileCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateOpaProfileCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateOpaProfileCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


