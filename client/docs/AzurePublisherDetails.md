# AzurePublisherDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to [**AzureCommonImages**](AzureCommonImages.md) |  | [optional] 

## Methods

### NewAzurePublisherDetails

`func NewAzurePublisherDetails() *AzurePublisherDetails`

NewAzurePublisherDetails instantiates a new AzurePublisherDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePublisherDetailsWithDefaults

`func NewAzurePublisherDetailsWithDefaults() *AzurePublisherDetails`

NewAzurePublisherDetailsWithDefaults instantiates a new AzurePublisherDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *AzurePublisherDetails) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AzurePublisherDetails) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AzurePublisherDetails) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AzurePublisherDetails) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *AzurePublisherDetails) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *AzurePublisherDetails) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetImage

`func (o *AzurePublisherDetails) GetImage() AzureCommonImages`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AzurePublisherDetails) GetImageOk() (*AzureCommonImages, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AzurePublisherDetails) SetImage(v AzureCommonImages)`

SetImage sets Image field to given value.

### HasImage

`func (o *AzurePublisherDetails) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


