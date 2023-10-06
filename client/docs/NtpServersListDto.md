# NtpServersListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**AccessProfileName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNtpServersListDto

`func NewNtpServersListDto() *NtpServersListDto`

NewNtpServersListDto instantiates a new NtpServersListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpServersListDtoWithDefaults

`func NewNtpServersListDtoWithDefaults() *NtpServersListDto`

NewNtpServersListDtoWithDefaults instantiates a new NtpServersListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NtpServersListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NtpServersListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NtpServersListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NtpServersListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *NtpServersListDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NtpServersListDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NtpServersListDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NtpServersListDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *NtpServersListDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NtpServersListDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAccessProfileName

`func (o *NtpServersListDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *NtpServersListDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *NtpServersListDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.

### HasAccessProfileName

`func (o *NtpServersListDto) HasAccessProfileName() bool`

HasAccessProfileName returns a boolean if a field has been set.

### SetAccessProfileNameNil

`func (o *NtpServersListDto) SetAccessProfileNameNil(b bool)`

 SetAccessProfileNameNil sets the value for AccessProfileName to be an explicit nil

### UnsetAccessProfileName
`func (o *NtpServersListDto) UnsetAccessProfileName()`

UnsetAccessProfileName ensures that no value is present for AccessProfileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


