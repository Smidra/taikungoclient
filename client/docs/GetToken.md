# GetToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** |  | [optional] 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**RefreshTokenExpireTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetToken

`func NewGetToken() *GetToken`

NewGetToken instantiates a new GetToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenWithDefaults

`func NewGetTokenWithDefaults() *GetToken`

NewGetTokenWithDefaults instantiates a new GetToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GetToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GetToken) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GetToken) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetRefreshToken

`func (o *GetToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GetToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GetToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GetToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *GetToken) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *GetToken) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetRefreshTokenExpireTime

`func (o *GetToken) GetRefreshTokenExpireTime() time.Time`

GetRefreshTokenExpireTime returns the RefreshTokenExpireTime field if non-nil, zero value otherwise.

### GetRefreshTokenExpireTimeOk

`func (o *GetToken) GetRefreshTokenExpireTimeOk() (*time.Time, bool)`

GetRefreshTokenExpireTimeOk returns a tuple with the RefreshTokenExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpireTime

`func (o *GetToken) SetRefreshTokenExpireTime(v time.Time)`

SetRefreshTokenExpireTime sets RefreshTokenExpireTime field to given value.

### HasRefreshTokenExpireTime

`func (o *GetToken) HasRefreshTokenExpireTime() bool`

HasRefreshTokenExpireTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


