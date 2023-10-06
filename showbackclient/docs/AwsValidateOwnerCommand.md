# AwsValidateOwnerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | Pointer to **int32** |  | [optional] 
**Owners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAwsValidateOwnerCommand

`func NewAwsValidateOwnerCommand() *AwsValidateOwnerCommand`

NewAwsValidateOwnerCommand instantiates a new AwsValidateOwnerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsValidateOwnerCommandWithDefaults

`func NewAwsValidateOwnerCommandWithDefaults() *AwsValidateOwnerCommand`

NewAwsValidateOwnerCommandWithDefaults instantiates a new AwsValidateOwnerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *AwsValidateOwnerCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AwsValidateOwnerCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AwsValidateOwnerCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *AwsValidateOwnerCommand) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetOwners

`func (o *AwsValidateOwnerCommand) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AwsValidateOwnerCommand) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AwsValidateOwnerCommand) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AwsValidateOwnerCommand) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwnersNil

`func (o *AwsValidateOwnerCommand) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *AwsValidateOwnerCommand) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


