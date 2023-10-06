# OrganizationCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** |  | [optional] 
**DiscountRate** | Pointer to **NullableFloat64** |  | [optional] 
**IsEligibleUpdateSubscription** | Pointer to **bool** |  | [optional] 
**AdminCloudCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewOrganizationCreateCommand

`func NewOrganizationCreateCommand() *OrganizationCreateCommand`

NewOrganizationCreateCommand instantiates a new OrganizationCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateCommandWithDefaults

`func NewOrganizationCreateCommandWithDefaults() *OrganizationCreateCommand`

NewOrganizationCreateCommandWithDefaults instantiates a new OrganizationCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationCreateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationCreateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationCreateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *OrganizationCreateCommand) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationCreateCommand) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationCreateCommand) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OrganizationCreateCommand) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *OrganizationCreateCommand) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationCreateCommand) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPhone

`func (o *OrganizationCreateCommand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationCreateCommand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationCreateCommand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationCreateCommand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *OrganizationCreateCommand) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *OrganizationCreateCommand) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *OrganizationCreateCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationCreateCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationCreateCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationCreateCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationCreateCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationCreateCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetBillingEmail

`func (o *OrganizationCreateCommand) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrganizationCreateCommand) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrganizationCreateCommand) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrganizationCreateCommand) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *OrganizationCreateCommand) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *OrganizationCreateCommand) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetAddress

`func (o *OrganizationCreateCommand) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationCreateCommand) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationCreateCommand) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganizationCreateCommand) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *OrganizationCreateCommand) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *OrganizationCreateCommand) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCountry

`func (o *OrganizationCreateCommand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationCreateCommand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationCreateCommand) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationCreateCommand) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *OrganizationCreateCommand) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OrganizationCreateCommand) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *OrganizationCreateCommand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationCreateCommand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationCreateCommand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationCreateCommand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrganizationCreateCommand) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrganizationCreateCommand) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetVatNumber

`func (o *OrganizationCreateCommand) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *OrganizationCreateCommand) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *OrganizationCreateCommand) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *OrganizationCreateCommand) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *OrganizationCreateCommand) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *OrganizationCreateCommand) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetDiscountRate

`func (o *OrganizationCreateCommand) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *OrganizationCreateCommand) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *OrganizationCreateCommand) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *OrganizationCreateCommand) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### SetDiscountRateNil

`func (o *OrganizationCreateCommand) SetDiscountRateNil(b bool)`

 SetDiscountRateNil sets the value for DiscountRate to be an explicit nil

### UnsetDiscountRate
`func (o *OrganizationCreateCommand) UnsetDiscountRate()`

UnsetDiscountRate ensures that no value is present for DiscountRate, not even an explicit nil
### GetIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) GetIsEligibleUpdateSubscription() bool`

GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field if non-nil, zero value otherwise.

### GetIsEligibleUpdateSubscriptionOk

`func (o *OrganizationCreateCommand) GetIsEligibleUpdateSubscriptionOk() (*bool, bool)`

GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) SetIsEligibleUpdateSubscription(v bool)`

SetIsEligibleUpdateSubscription sets IsEligibleUpdateSubscription field to given value.

### HasIsEligibleUpdateSubscription

`func (o *OrganizationCreateCommand) HasIsEligibleUpdateSubscription() bool`

HasIsEligibleUpdateSubscription returns a boolean if a field has been set.

### GetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialId() int32`

GetAdminCloudCredentialId returns the AdminCloudCredentialId field if non-nil, zero value otherwise.

### GetAdminCloudCredentialIdOk

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialIdOk() (*int32, bool)`

GetAdminCloudCredentialIdOk returns a tuple with the AdminCloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) SetAdminCloudCredentialId(v int32)`

SetAdminCloudCredentialId sets AdminCloudCredentialId field to given value.

### HasAdminCloudCredentialId

`func (o *OrganizationCreateCommand) HasAdminCloudCredentialId() bool`

HasAdminCloudCredentialId returns a boolean if a field has been set.

### SetAdminCloudCredentialIdNil

`func (o *OrganizationCreateCommand) SetAdminCloudCredentialIdNil(b bool)`

 SetAdminCloudCredentialIdNil sets the value for AdminCloudCredentialId to be an explicit nil

### UnsetAdminCloudCredentialId
`func (o *OrganizationCreateCommand) UnsetAdminCloudCredentialId()`

UnsetAdminCloudCredentialId ensures that no value is present for AdminCloudCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


