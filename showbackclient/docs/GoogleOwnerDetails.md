# GoogleOwnerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to [**GoogleCommonImages**](GoogleCommonImages.md) |  | [optional] 

## Methods

### NewGoogleOwnerDetails

`func NewGoogleOwnerDetails() *GoogleOwnerDetails`

NewGoogleOwnerDetails instantiates a new GoogleOwnerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleOwnerDetailsWithDefaults

`func NewGoogleOwnerDetailsWithDefaults() *GoogleOwnerDetails`

NewGoogleOwnerDetailsWithDefaults instantiates a new GoogleOwnerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *GoogleOwnerDetails) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GoogleOwnerDetails) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GoogleOwnerDetails) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GoogleOwnerDetails) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *GoogleOwnerDetails) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *GoogleOwnerDetails) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetImage

`func (o *GoogleOwnerDetails) GetImage() GoogleCommonImages`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GoogleOwnerDetails) GetImageOk() (*GoogleCommonImages, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GoogleOwnerDetails) SetImage(v GoogleCommonImages)`

SetImage sets Image field to given value.

### HasImage

`func (o *GoogleOwnerDetails) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


