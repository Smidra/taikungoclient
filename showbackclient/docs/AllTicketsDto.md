# AllTicketsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CurrentStatusDate** | Pointer to **NullableString** |  | [optional] 
**LastModifier** | Pointer to **NullableString** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**PartnerLogo** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PartnerName** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAllTicketsDto

`func NewAllTicketsDto() *AllTicketsDto`

NewAllTicketsDto instantiates a new AllTicketsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllTicketsDtoWithDefaults

`func NewAllTicketsDtoWithDefaults() *AllTicketsDto`

NewAllTicketsDtoWithDefaults instantiates a new AllTicketsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllTicketsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllTicketsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllTicketsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllTicketsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AllTicketsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AllTicketsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AllTicketsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllTicketsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllTicketsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllTicketsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AllTicketsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AllTicketsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *AllTicketsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AllTicketsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AllTicketsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AllTicketsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AllTicketsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AllTicketsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCreatedAt

`func (o *AllTicketsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllTicketsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllTicketsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllTicketsDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AllTicketsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AllTicketsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetOrganizationId

`func (o *AllTicketsDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AllTicketsDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AllTicketsDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AllTicketsDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AllTicketsDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AllTicketsDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AllTicketsDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AllTicketsDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AllTicketsDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AllTicketsDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCreatedBy

`func (o *AllTicketsDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AllTicketsDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AllTicketsDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AllTicketsDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AllTicketsDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AllTicketsDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCurrentStatusDate

`func (o *AllTicketsDto) GetCurrentStatusDate() string`

GetCurrentStatusDate returns the CurrentStatusDate field if non-nil, zero value otherwise.

### GetCurrentStatusDateOk

`func (o *AllTicketsDto) GetCurrentStatusDateOk() (*string, bool)`

GetCurrentStatusDateOk returns a tuple with the CurrentStatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatusDate

`func (o *AllTicketsDto) SetCurrentStatusDate(v string)`

SetCurrentStatusDate sets CurrentStatusDate field to given value.

### HasCurrentStatusDate

`func (o *AllTicketsDto) HasCurrentStatusDate() bool`

HasCurrentStatusDate returns a boolean if a field has been set.

### SetCurrentStatusDateNil

`func (o *AllTicketsDto) SetCurrentStatusDateNil(b bool)`

 SetCurrentStatusDateNil sets the value for CurrentStatusDate to be an explicit nil

### UnsetCurrentStatusDate
`func (o *AllTicketsDto) UnsetCurrentStatusDate()`

UnsetCurrentStatusDate ensures that no value is present for CurrentStatusDate, not even an explicit nil
### GetLastModifier

`func (o *AllTicketsDto) GetLastModifier() string`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *AllTicketsDto) GetLastModifierOk() (*string, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *AllTicketsDto) SetLastModifier(v string)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *AllTicketsDto) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### SetLastModifierNil

`func (o *AllTicketsDto) SetLastModifierNil(b bool)`

 SetLastModifierNil sets the value for LastModifier to be an explicit nil

### UnsetLastModifier
`func (o *AllTicketsDto) UnsetLastModifier()`

UnsetLastModifier ensures that no value is present for LastModifier, not even an explicit nil
### GetNumber

`func (o *AllTicketsDto) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AllTicketsDto) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AllTicketsDto) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AllTicketsDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPartnerLogo

`func (o *AllTicketsDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *AllTicketsDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *AllTicketsDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.

### HasPartnerLogo

`func (o *AllTicketsDto) HasPartnerLogo() bool`

HasPartnerLogo returns a boolean if a field has been set.

### SetPartnerLogoNil

`func (o *AllTicketsDto) SetPartnerLogoNil(b bool)`

 SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil

### UnsetPartnerLogo
`func (o *AllTicketsDto) UnsetPartnerLogo()`

UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
### GetDescription

`func (o *AllTicketsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllTicketsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllTicketsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllTicketsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AllTicketsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AllTicketsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPartnerName

`func (o *AllTicketsDto) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *AllTicketsDto) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *AllTicketsDto) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *AllTicketsDto) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *AllTicketsDto) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *AllTicketsDto) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetUserId

`func (o *AllTicketsDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AllTicketsDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AllTicketsDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AllTicketsDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AllTicketsDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AllTicketsDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


