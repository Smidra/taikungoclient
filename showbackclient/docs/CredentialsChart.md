# CredentialsChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amazon** | Pointer to [**[]AmazonCredentialsListDto**](AmazonCredentialsListDto.md) |  | [optional] 
**Openstack** | Pointer to [**[]OpenstackCredentialsListDto**](OpenstackCredentialsListDto.md) |  | [optional] 
**Azure** | Pointer to [**[]AzureCredentialsListDto**](AzureCredentialsListDto.md) |  | [optional] 
**Google** | Pointer to [**[]GoogleCredentialsListDto**](GoogleCredentialsListDto.md) |  | [optional] 
**Tanzu** | Pointer to [**[]TanzuCredentialsListDto**](TanzuCredentialsListDto.md) |  | [optional] 
**TotalCountOpenstack** | Pointer to **int32** |  | [optional] 
**TotalCountAws** | Pointer to **int32** |  | [optional] 
**TotalCountAzure** | Pointer to **int32** |  | [optional] 
**TotalCountGoogle** | Pointer to **int32** |  | [optional] 
**TotalCountTanzu** | Pointer to **int32** |  | [optional] 

## Methods

### NewCredentialsChart

`func NewCredentialsChart() *CredentialsChart`

NewCredentialsChart instantiates a new CredentialsChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsChartWithDefaults

`func NewCredentialsChartWithDefaults() *CredentialsChart`

NewCredentialsChartWithDefaults instantiates a new CredentialsChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazon

`func (o *CredentialsChart) GetAmazon() []AmazonCredentialsListDto`

GetAmazon returns the Amazon field if non-nil, zero value otherwise.

### GetAmazonOk

`func (o *CredentialsChart) GetAmazonOk() (*[]AmazonCredentialsListDto, bool)`

GetAmazonOk returns a tuple with the Amazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazon

`func (o *CredentialsChart) SetAmazon(v []AmazonCredentialsListDto)`

SetAmazon sets Amazon field to given value.

### HasAmazon

`func (o *CredentialsChart) HasAmazon() bool`

HasAmazon returns a boolean if a field has been set.

### SetAmazonNil

`func (o *CredentialsChart) SetAmazonNil(b bool)`

 SetAmazonNil sets the value for Amazon to be an explicit nil

### UnsetAmazon
`func (o *CredentialsChart) UnsetAmazon()`

UnsetAmazon ensures that no value is present for Amazon, not even an explicit nil
### GetOpenstack

`func (o *CredentialsChart) GetOpenstack() []OpenstackCredentialsListDto`

GetOpenstack returns the Openstack field if non-nil, zero value otherwise.

### GetOpenstackOk

`func (o *CredentialsChart) GetOpenstackOk() (*[]OpenstackCredentialsListDto, bool)`

GetOpenstackOk returns a tuple with the Openstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstack

`func (o *CredentialsChart) SetOpenstack(v []OpenstackCredentialsListDto)`

SetOpenstack sets Openstack field to given value.

### HasOpenstack

`func (o *CredentialsChart) HasOpenstack() bool`

HasOpenstack returns a boolean if a field has been set.

### SetOpenstackNil

`func (o *CredentialsChart) SetOpenstackNil(b bool)`

 SetOpenstackNil sets the value for Openstack to be an explicit nil

### UnsetOpenstack
`func (o *CredentialsChart) UnsetOpenstack()`

UnsetOpenstack ensures that no value is present for Openstack, not even an explicit nil
### GetAzure

`func (o *CredentialsChart) GetAzure() []AzureCredentialsListDto`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *CredentialsChart) GetAzureOk() (*[]AzureCredentialsListDto, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *CredentialsChart) SetAzure(v []AzureCredentialsListDto)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *CredentialsChart) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *CredentialsChart) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *CredentialsChart) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil
### GetGoogle

`func (o *CredentialsChart) GetGoogle() []GoogleCredentialsListDto`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *CredentialsChart) GetGoogleOk() (*[]GoogleCredentialsListDto, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *CredentialsChart) SetGoogle(v []GoogleCredentialsListDto)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *CredentialsChart) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### SetGoogleNil

`func (o *CredentialsChart) SetGoogleNil(b bool)`

 SetGoogleNil sets the value for Google to be an explicit nil

### UnsetGoogle
`func (o *CredentialsChart) UnsetGoogle()`

UnsetGoogle ensures that no value is present for Google, not even an explicit nil
### GetTanzu

`func (o *CredentialsChart) GetTanzu() []TanzuCredentialsListDto`

GetTanzu returns the Tanzu field if non-nil, zero value otherwise.

### GetTanzuOk

`func (o *CredentialsChart) GetTanzuOk() (*[]TanzuCredentialsListDto, bool)`

GetTanzuOk returns a tuple with the Tanzu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanzu

`func (o *CredentialsChart) SetTanzu(v []TanzuCredentialsListDto)`

SetTanzu sets Tanzu field to given value.

### HasTanzu

`func (o *CredentialsChart) HasTanzu() bool`

HasTanzu returns a boolean if a field has been set.

### SetTanzuNil

`func (o *CredentialsChart) SetTanzuNil(b bool)`

 SetTanzuNil sets the value for Tanzu to be an explicit nil

### UnsetTanzu
`func (o *CredentialsChart) UnsetTanzu()`

UnsetTanzu ensures that no value is present for Tanzu, not even an explicit nil
### GetTotalCountOpenstack

`func (o *CredentialsChart) GetTotalCountOpenstack() int32`

GetTotalCountOpenstack returns the TotalCountOpenstack field if non-nil, zero value otherwise.

### GetTotalCountOpenstackOk

`func (o *CredentialsChart) GetTotalCountOpenstackOk() (*int32, bool)`

GetTotalCountOpenstackOk returns a tuple with the TotalCountOpenstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountOpenstack

`func (o *CredentialsChart) SetTotalCountOpenstack(v int32)`

SetTotalCountOpenstack sets TotalCountOpenstack field to given value.

### HasTotalCountOpenstack

`func (o *CredentialsChart) HasTotalCountOpenstack() bool`

HasTotalCountOpenstack returns a boolean if a field has been set.

### GetTotalCountAws

`func (o *CredentialsChart) GetTotalCountAws() int32`

GetTotalCountAws returns the TotalCountAws field if non-nil, zero value otherwise.

### GetTotalCountAwsOk

`func (o *CredentialsChart) GetTotalCountAwsOk() (*int32, bool)`

GetTotalCountAwsOk returns a tuple with the TotalCountAws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountAws

`func (o *CredentialsChart) SetTotalCountAws(v int32)`

SetTotalCountAws sets TotalCountAws field to given value.

### HasTotalCountAws

`func (o *CredentialsChart) HasTotalCountAws() bool`

HasTotalCountAws returns a boolean if a field has been set.

### GetTotalCountAzure

`func (o *CredentialsChart) GetTotalCountAzure() int32`

GetTotalCountAzure returns the TotalCountAzure field if non-nil, zero value otherwise.

### GetTotalCountAzureOk

`func (o *CredentialsChart) GetTotalCountAzureOk() (*int32, bool)`

GetTotalCountAzureOk returns a tuple with the TotalCountAzure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountAzure

`func (o *CredentialsChart) SetTotalCountAzure(v int32)`

SetTotalCountAzure sets TotalCountAzure field to given value.

### HasTotalCountAzure

`func (o *CredentialsChart) HasTotalCountAzure() bool`

HasTotalCountAzure returns a boolean if a field has been set.

### GetTotalCountGoogle

`func (o *CredentialsChart) GetTotalCountGoogle() int32`

GetTotalCountGoogle returns the TotalCountGoogle field if non-nil, zero value otherwise.

### GetTotalCountGoogleOk

`func (o *CredentialsChart) GetTotalCountGoogleOk() (*int32, bool)`

GetTotalCountGoogleOk returns a tuple with the TotalCountGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountGoogle

`func (o *CredentialsChart) SetTotalCountGoogle(v int32)`

SetTotalCountGoogle sets TotalCountGoogle field to given value.

### HasTotalCountGoogle

`func (o *CredentialsChart) HasTotalCountGoogle() bool`

HasTotalCountGoogle returns a boolean if a field has been set.

### GetTotalCountTanzu

`func (o *CredentialsChart) GetTotalCountTanzu() int32`

GetTotalCountTanzu returns the TotalCountTanzu field if non-nil, zero value otherwise.

### GetTotalCountTanzuOk

`func (o *CredentialsChart) GetTotalCountTanzuOk() (*int32, bool)`

GetTotalCountTanzuOk returns a tuple with the TotalCountTanzu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountTanzu

`func (o *CredentialsChart) SetTotalCountTanzu(v int32)`

SetTotalCountTanzu sets TotalCountTanzu field to given value.

### HasTotalCountTanzu

`func (o *CredentialsChart) HasTotalCountTanzu() bool`

HasTotalCountTanzu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


