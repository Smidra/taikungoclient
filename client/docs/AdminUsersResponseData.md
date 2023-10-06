# AdminUsersResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **bool** |  | [optional] 
**Csm** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdminUsersResponseData

`func NewAdminUsersResponseData() *AdminUsersResponseData`

NewAdminUsersResponseData instantiates a new AdminUsersResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsersResponseDataWithDefaults

`func NewAdminUsersResponseDataWithDefaults() *AdminUsersResponseData`

NewAdminUsersResponseDataWithDefaults instantiates a new AdminUsersResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminUsersResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminUsersResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminUsersResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminUsersResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AdminUsersResponseData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdminUsersResponseData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AdminUsersResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUsersResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUsersResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminUsersResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminUsersResponseData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminUsersResponseData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *AdminUsersResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUsersResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUsersResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdminUsersResponseData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AdminUsersResponseData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminUsersResponseData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRole

`func (o *AdminUsersResponseData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminUsersResponseData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminUsersResponseData) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminUsersResponseData) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AdminUsersResponseData) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AdminUsersResponseData) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetOrganizationName

`func (o *AdminUsersResponseData) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AdminUsersResponseData) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AdminUsersResponseData) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AdminUsersResponseData) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AdminUsersResponseData) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AdminUsersResponseData) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOwner

`func (o *AdminUsersResponseData) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AdminUsersResponseData) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AdminUsersResponseData) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AdminUsersResponseData) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCsm

`func (o *AdminUsersResponseData) GetCsm() bool`

GetCsm returns the Csm field if non-nil, zero value otherwise.

### GetCsmOk

`func (o *AdminUsersResponseData) GetCsmOk() (*bool, bool)`

GetCsmOk returns a tuple with the Csm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsm

`func (o *AdminUsersResponseData) SetCsm(v bool)`

SetCsm sets Csm field to given value.

### HasCsm

`func (o *AdminUsersResponseData) HasCsm() bool`

HasCsm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


