# AwsOwnerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**OwnerName** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to [**AwsCommonImages**](AwsCommonImages.md) |  | [optional] 

## Methods

### NewAwsOwnerDetails

`func NewAwsOwnerDetails() *AwsOwnerDetails`

NewAwsOwnerDetails instantiates a new AwsOwnerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsOwnerDetailsWithDefaults

`func NewAwsOwnerDetailsWithDefaults() *AwsOwnerDetails`

NewAwsOwnerDetailsWithDefaults instantiates a new AwsOwnerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *AwsOwnerDetails) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AwsOwnerDetails) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AwsOwnerDetails) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AwsOwnerDetails) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *AwsOwnerDetails) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AwsOwnerDetails) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwnerName

`func (o *AwsOwnerDetails) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *AwsOwnerDetails) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *AwsOwnerDetails) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *AwsOwnerDetails) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *AwsOwnerDetails) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *AwsOwnerDetails) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetImage

`func (o *AwsOwnerDetails) GetImage() AwsCommonImages`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AwsOwnerDetails) GetImageOk() (*AwsCommonImages, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AwsOwnerDetails) SetImage(v AwsCommonImages)`

SetImage sets Image field to given value.

### HasImage

`func (o *AwsOwnerDetails) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


