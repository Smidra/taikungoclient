# UserForListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**HasCustomerId** | Pointer to **bool** |  | [optional] 
**HasPaymentMethod** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**IsEmailConfirmed** | Pointer to **bool** |  | [optional] 
**IsEmailNotificationEnabled** | Pointer to **bool** |  | [optional] 
**IsForcedToResetPassword** | Pointer to **bool** |  | [optional] 
**IsCsm** | Pointer to **bool** |  | [optional] 
**IsEligibleUpdateSubscription** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsApprovedByPartner** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**HasRepo** | Pointer to **bool** |  | [optional] 
**DemoModeEnabled** | Pointer to **bool** |  | [optional] 
**LastLoginAt** | Pointer to **NullableString** |  | [optional] 
**BoundProjects** | Pointer to [**[]ProjectDto**](ProjectDto.md) |  | [optional] 
**Partner** | Pointer to [**PartnerDetailsForUserDto**](PartnerDetailsForUserDto.md) |  | [optional] 

## Methods

### NewUserForListDto

`func NewUserForListDto() *UserForListDto`

NewUserForListDto instantiates a new UserForListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserForListDtoWithDefaults

`func NewUserForListDtoWithDefaults() *UserForListDto`

NewUserForListDtoWithDefaults instantiates a new UserForListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserForListDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserForListDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserForListDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserForListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserForListDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserForListDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUsername

`func (o *UserForListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserForListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserForListDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserForListDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserForListDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserForListDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetOrganizationName

`func (o *UserForListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *UserForListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *UserForListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *UserForListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *UserForListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *UserForListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetHasCustomerId

`func (o *UserForListDto) GetHasCustomerId() bool`

GetHasCustomerId returns the HasCustomerId field if non-nil, zero value otherwise.

### GetHasCustomerIdOk

`func (o *UserForListDto) GetHasCustomerIdOk() (*bool, bool)`

GetHasCustomerIdOk returns a tuple with the HasCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomerId

`func (o *UserForListDto) SetHasCustomerId(v bool)`

SetHasCustomerId sets HasCustomerId field to given value.

### HasHasCustomerId

`func (o *UserForListDto) HasHasCustomerId() bool`

HasHasCustomerId returns a boolean if a field has been set.

### GetHasPaymentMethod

`func (o *UserForListDto) GetHasPaymentMethod() bool`

GetHasPaymentMethod returns the HasPaymentMethod field if non-nil, zero value otherwise.

### GetHasPaymentMethodOk

`func (o *UserForListDto) GetHasPaymentMethodOk() (*bool, bool)`

GetHasPaymentMethodOk returns a tuple with the HasPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPaymentMethod

`func (o *UserForListDto) SetHasPaymentMethod(v bool)`

SetHasPaymentMethod sets HasPaymentMethod field to given value.

### HasHasPaymentMethod

`func (o *UserForListDto) HasHasPaymentMethod() bool`

HasHasPaymentMethod returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UserForListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserForListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserForListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UserForListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRole

`func (o *UserForListDto) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserForListDto) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserForListDto) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserForListDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmail

`func (o *UserForListDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserForListDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserForListDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserForListDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserForListDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserForListDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDisplayName

`func (o *UserForListDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserForListDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserForListDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserForListDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserForListDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserForListDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCreatedAt

`func (o *UserForListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserForListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserForListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserForListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *UserForListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UserForListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsEmailConfirmed

`func (o *UserForListDto) GetIsEmailConfirmed() bool`

GetIsEmailConfirmed returns the IsEmailConfirmed field if non-nil, zero value otherwise.

### GetIsEmailConfirmedOk

`func (o *UserForListDto) GetIsEmailConfirmedOk() (*bool, bool)`

GetIsEmailConfirmedOk returns a tuple with the IsEmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailConfirmed

`func (o *UserForListDto) SetIsEmailConfirmed(v bool)`

SetIsEmailConfirmed sets IsEmailConfirmed field to given value.

### HasIsEmailConfirmed

`func (o *UserForListDto) HasIsEmailConfirmed() bool`

HasIsEmailConfirmed returns a boolean if a field has been set.

### GetIsEmailNotificationEnabled

`func (o *UserForListDto) GetIsEmailNotificationEnabled() bool`

GetIsEmailNotificationEnabled returns the IsEmailNotificationEnabled field if non-nil, zero value otherwise.

### GetIsEmailNotificationEnabledOk

`func (o *UserForListDto) GetIsEmailNotificationEnabledOk() (*bool, bool)`

GetIsEmailNotificationEnabledOk returns a tuple with the IsEmailNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailNotificationEnabled

`func (o *UserForListDto) SetIsEmailNotificationEnabled(v bool)`

SetIsEmailNotificationEnabled sets IsEmailNotificationEnabled field to given value.

### HasIsEmailNotificationEnabled

`func (o *UserForListDto) HasIsEmailNotificationEnabled() bool`

HasIsEmailNotificationEnabled returns a boolean if a field has been set.

### GetIsForcedToResetPassword

`func (o *UserForListDto) GetIsForcedToResetPassword() bool`

GetIsForcedToResetPassword returns the IsForcedToResetPassword field if non-nil, zero value otherwise.

### GetIsForcedToResetPasswordOk

`func (o *UserForListDto) GetIsForcedToResetPasswordOk() (*bool, bool)`

GetIsForcedToResetPasswordOk returns a tuple with the IsForcedToResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcedToResetPassword

`func (o *UserForListDto) SetIsForcedToResetPassword(v bool)`

SetIsForcedToResetPassword sets IsForcedToResetPassword field to given value.

### HasIsForcedToResetPassword

`func (o *UserForListDto) HasIsForcedToResetPassword() bool`

HasIsForcedToResetPassword returns a boolean if a field has been set.

### GetIsCsm

`func (o *UserForListDto) GetIsCsm() bool`

GetIsCsm returns the IsCsm field if non-nil, zero value otherwise.

### GetIsCsmOk

`func (o *UserForListDto) GetIsCsmOk() (*bool, bool)`

GetIsCsmOk returns a tuple with the IsCsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCsm

`func (o *UserForListDto) SetIsCsm(v bool)`

SetIsCsm sets IsCsm field to given value.

### HasIsCsm

`func (o *UserForListDto) HasIsCsm() bool`

HasIsCsm returns a boolean if a field has been set.

### GetIsEligibleUpdateSubscription

`func (o *UserForListDto) GetIsEligibleUpdateSubscription() bool`

GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field if non-nil, zero value otherwise.

### GetIsEligibleUpdateSubscriptionOk

`func (o *UserForListDto) GetIsEligibleUpdateSubscriptionOk() (*bool, bool)`

GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleUpdateSubscription

`func (o *UserForListDto) SetIsEligibleUpdateSubscription(v bool)`

SetIsEligibleUpdateSubscription sets IsEligibleUpdateSubscription field to given value.

### HasIsEligibleUpdateSubscription

`func (o *UserForListDto) HasIsEligibleUpdateSubscription() bool`

HasIsEligibleUpdateSubscription returns a boolean if a field has been set.

### GetIsLocked

`func (o *UserForListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserForListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserForListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UserForListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsApprovedByPartner

`func (o *UserForListDto) GetIsApprovedByPartner() bool`

GetIsApprovedByPartner returns the IsApprovedByPartner field if non-nil, zero value otherwise.

### GetIsApprovedByPartnerOk

`func (o *UserForListDto) GetIsApprovedByPartnerOk() (*bool, bool)`

GetIsApprovedByPartnerOk returns a tuple with the IsApprovedByPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovedByPartner

`func (o *UserForListDto) SetIsApprovedByPartner(v bool)`

SetIsApprovedByPartner sets IsApprovedByPartner field to given value.

### HasIsApprovedByPartner

`func (o *UserForListDto) HasIsApprovedByPartner() bool`

HasIsApprovedByPartner returns a boolean if a field has been set.

### GetOwner

`func (o *UserForListDto) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserForListDto) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserForListDto) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UserForListDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *UserForListDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *UserForListDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *UserForListDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *UserForListDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetHasRepo

`func (o *UserForListDto) GetHasRepo() bool`

GetHasRepo returns the HasRepo field if non-nil, zero value otherwise.

### GetHasRepoOk

`func (o *UserForListDto) GetHasRepoOk() (*bool, bool)`

GetHasRepoOk returns a tuple with the HasRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRepo

`func (o *UserForListDto) SetHasRepo(v bool)`

SetHasRepo sets HasRepo field to given value.

### HasHasRepo

`func (o *UserForListDto) HasHasRepo() bool`

HasHasRepo returns a boolean if a field has been set.

### GetDemoModeEnabled

`func (o *UserForListDto) GetDemoModeEnabled() bool`

GetDemoModeEnabled returns the DemoModeEnabled field if non-nil, zero value otherwise.

### GetDemoModeEnabledOk

`func (o *UserForListDto) GetDemoModeEnabledOk() (*bool, bool)`

GetDemoModeEnabledOk returns a tuple with the DemoModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoModeEnabled

`func (o *UserForListDto) SetDemoModeEnabled(v bool)`

SetDemoModeEnabled sets DemoModeEnabled field to given value.

### HasDemoModeEnabled

`func (o *UserForListDto) HasDemoModeEnabled() bool`

HasDemoModeEnabled returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *UserForListDto) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UserForListDto) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UserForListDto) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UserForListDto) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *UserForListDto) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *UserForListDto) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetBoundProjects

`func (o *UserForListDto) GetBoundProjects() []ProjectDto`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *UserForListDto) GetBoundProjectsOk() (*[]ProjectDto, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *UserForListDto) SetBoundProjects(v []ProjectDto)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *UserForListDto) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### SetBoundProjectsNil

`func (o *UserForListDto) SetBoundProjectsNil(b bool)`

 SetBoundProjectsNil sets the value for BoundProjects to be an explicit nil

### UnsetBoundProjects
`func (o *UserForListDto) UnsetBoundProjects()`

UnsetBoundProjects ensures that no value is present for BoundProjects, not even an explicit nil
### GetPartner

`func (o *UserForListDto) GetPartner() PartnerDetailsForUserDto`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *UserForListDto) GetPartnerOk() (*PartnerDetailsForUserDto, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *UserForListDto) SetPartner(v PartnerDetailsForUserDto)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *UserForListDto) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


