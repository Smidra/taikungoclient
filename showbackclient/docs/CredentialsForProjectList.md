# CredentialsForProjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | Pointer to [**CloudType**](CloudType.md) |  | [optional] 
**CloudCredentialRevision** | Pointer to **int32** |  | [optional] 
**BillingEnabled** | Pointer to **bool** |  | [optional] 
**ContinentName** | Pointer to **NullableString** |  | [optional] 
**RequiresVPN** | Pointer to **bool** |  | [optional] 
**Openstack** | Pointer to [**OpenstackCredentialsForProjectDto**](OpenstackCredentialsForProjectDto.md) |  | [optional] 
**Azure** | Pointer to [**AzureCredentialsForProjectDto**](AzureCredentialsForProjectDto.md) |  | [optional] 
**Aws** | Pointer to [**AwsCredentialsForProjectDto**](AwsCredentialsForProjectDto.md) |  | [optional] 
**Google** | Pointer to [**GoogleCredentialForProjectDto**](GoogleCredentialForProjectDto.md) |  | [optional] 
**Tanzu** | Pointer to [**TanzuCredentialsForProjectDto**](TanzuCredentialsForProjectDto.md) |  | [optional] 
**Proxmox** | Pointer to [**ProxmoxCredentialsForProjectDto**](ProxmoxCredentialsForProjectDto.md) |  | [optional] 
**Openshift** | Pointer to [**OpenshiftCredentialForProjectDto**](OpenshiftCredentialForProjectDto.md) |  | [optional] 

## Methods

### NewCredentialsForProjectList

`func NewCredentialsForProjectList() *CredentialsForProjectList`

NewCredentialsForProjectList instantiates a new CredentialsForProjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsForProjectListWithDefaults

`func NewCredentialsForProjectListWithDefaults() *CredentialsForProjectList`

NewCredentialsForProjectListWithDefaults instantiates a new CredentialsForProjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *CredentialsForProjectList) GetCloudType() CloudType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CredentialsForProjectList) GetCloudTypeOk() (*CloudType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CredentialsForProjectList) SetCloudType(v CloudType)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *CredentialsForProjectList) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudCredentialRevision

`func (o *CredentialsForProjectList) GetCloudCredentialRevision() int32`

GetCloudCredentialRevision returns the CloudCredentialRevision field if non-nil, zero value otherwise.

### GetCloudCredentialRevisionOk

`func (o *CredentialsForProjectList) GetCloudCredentialRevisionOk() (*int32, bool)`

GetCloudCredentialRevisionOk returns a tuple with the CloudCredentialRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialRevision

`func (o *CredentialsForProjectList) SetCloudCredentialRevision(v int32)`

SetCloudCredentialRevision sets CloudCredentialRevision field to given value.

### HasCloudCredentialRevision

`func (o *CredentialsForProjectList) HasCloudCredentialRevision() bool`

HasCloudCredentialRevision returns a boolean if a field has been set.

### GetBillingEnabled

`func (o *CredentialsForProjectList) GetBillingEnabled() bool`

GetBillingEnabled returns the BillingEnabled field if non-nil, zero value otherwise.

### GetBillingEnabledOk

`func (o *CredentialsForProjectList) GetBillingEnabledOk() (*bool, bool)`

GetBillingEnabledOk returns a tuple with the BillingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEnabled

`func (o *CredentialsForProjectList) SetBillingEnabled(v bool)`

SetBillingEnabled sets BillingEnabled field to given value.

### HasBillingEnabled

`func (o *CredentialsForProjectList) HasBillingEnabled() bool`

HasBillingEnabled returns a boolean if a field has been set.

### GetContinentName

`func (o *CredentialsForProjectList) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *CredentialsForProjectList) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *CredentialsForProjectList) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *CredentialsForProjectList) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.

### SetContinentNameNil

`func (o *CredentialsForProjectList) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *CredentialsForProjectList) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
### GetRequiresVPN

`func (o *CredentialsForProjectList) GetRequiresVPN() bool`

GetRequiresVPN returns the RequiresVPN field if non-nil, zero value otherwise.

### GetRequiresVPNOk

`func (o *CredentialsForProjectList) GetRequiresVPNOk() (*bool, bool)`

GetRequiresVPNOk returns a tuple with the RequiresVPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVPN

`func (o *CredentialsForProjectList) SetRequiresVPN(v bool)`

SetRequiresVPN sets RequiresVPN field to given value.

### HasRequiresVPN

`func (o *CredentialsForProjectList) HasRequiresVPN() bool`

HasRequiresVPN returns a boolean if a field has been set.

### GetOpenstack

`func (o *CredentialsForProjectList) GetOpenstack() OpenstackCredentialsForProjectDto`

GetOpenstack returns the Openstack field if non-nil, zero value otherwise.

### GetOpenstackOk

`func (o *CredentialsForProjectList) GetOpenstackOk() (*OpenstackCredentialsForProjectDto, bool)`

GetOpenstackOk returns a tuple with the Openstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstack

`func (o *CredentialsForProjectList) SetOpenstack(v OpenstackCredentialsForProjectDto)`

SetOpenstack sets Openstack field to given value.

### HasOpenstack

`func (o *CredentialsForProjectList) HasOpenstack() bool`

HasOpenstack returns a boolean if a field has been set.

### GetAzure

`func (o *CredentialsForProjectList) GetAzure() AzureCredentialsForProjectDto`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *CredentialsForProjectList) GetAzureOk() (*AzureCredentialsForProjectDto, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *CredentialsForProjectList) SetAzure(v AzureCredentialsForProjectDto)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *CredentialsForProjectList) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetAws

`func (o *CredentialsForProjectList) GetAws() AwsCredentialsForProjectDto`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *CredentialsForProjectList) GetAwsOk() (*AwsCredentialsForProjectDto, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *CredentialsForProjectList) SetAws(v AwsCredentialsForProjectDto)`

SetAws sets Aws field to given value.

### HasAws

`func (o *CredentialsForProjectList) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetGoogle

`func (o *CredentialsForProjectList) GetGoogle() GoogleCredentialForProjectDto`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *CredentialsForProjectList) GetGoogleOk() (*GoogleCredentialForProjectDto, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *CredentialsForProjectList) SetGoogle(v GoogleCredentialForProjectDto)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *CredentialsForProjectList) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetTanzu

`func (o *CredentialsForProjectList) GetTanzu() TanzuCredentialsForProjectDto`

GetTanzu returns the Tanzu field if non-nil, zero value otherwise.

### GetTanzuOk

`func (o *CredentialsForProjectList) GetTanzuOk() (*TanzuCredentialsForProjectDto, bool)`

GetTanzuOk returns a tuple with the Tanzu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzu

`func (o *CredentialsForProjectList) SetTanzu(v TanzuCredentialsForProjectDto)`

SetTanzu sets Tanzu field to given value.

### HasTanzu

`func (o *CredentialsForProjectList) HasTanzu() bool`

HasTanzu returns a boolean if a field has been set.

### GetProxmox

`func (o *CredentialsForProjectList) GetProxmox() ProxmoxCredentialsForProjectDto`

GetProxmox returns the Proxmox field if non-nil, zero value otherwise.

### GetProxmoxOk

`func (o *CredentialsForProjectList) GetProxmoxOk() (*ProxmoxCredentialsForProjectDto, bool)`

GetProxmoxOk returns a tuple with the Proxmox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmox

`func (o *CredentialsForProjectList) SetProxmox(v ProxmoxCredentialsForProjectDto)`

SetProxmox sets Proxmox field to given value.

### HasProxmox

`func (o *CredentialsForProjectList) HasProxmox() bool`

HasProxmox returns a boolean if a field has been set.

### GetOpenshift

`func (o *CredentialsForProjectList) GetOpenshift() OpenshiftCredentialForProjectDto`

GetOpenshift returns the Openshift field if non-nil, zero value otherwise.

### GetOpenshiftOk

`func (o *CredentialsForProjectList) GetOpenshiftOk() (*OpenshiftCredentialForProjectDto, bool)`

GetOpenshiftOk returns a tuple with the Openshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenshift

`func (o *CredentialsForProjectList) SetOpenshift(v OpenshiftCredentialForProjectDto)`

SetOpenshift sets Openshift field to given value.

### HasOpenshift

`func (o *CredentialsForProjectList) HasOpenshift() bool`

HasOpenshift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


