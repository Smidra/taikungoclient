# MonitoringCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PrometheusUrl** | Pointer to **NullableString** |  | [optional] 
**LokiUrl** | Pointer to **NullableString** |  | [optional] 
**AlertManagerUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMonitoringCredentialsListDto

`func NewMonitoringCredentialsListDto() *MonitoringCredentialsListDto`

NewMonitoringCredentialsListDto instantiates a new MonitoringCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringCredentialsListDtoWithDefaults

`func NewMonitoringCredentialsListDtoWithDefaults() *MonitoringCredentialsListDto`

NewMonitoringCredentialsListDtoWithDefaults instantiates a new MonitoringCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitoringCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MonitoringCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *MonitoringCredentialsListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MonitoringCredentialsListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MonitoringCredentialsListDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MonitoringCredentialsListDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *MonitoringCredentialsListDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MonitoringCredentialsListDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *MonitoringCredentialsListDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MonitoringCredentialsListDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MonitoringCredentialsListDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MonitoringCredentialsListDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MonitoringCredentialsListDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MonitoringCredentialsListDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPrometheusUrl

`func (o *MonitoringCredentialsListDto) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *MonitoringCredentialsListDto) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *MonitoringCredentialsListDto) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.

### HasPrometheusUrl

`func (o *MonitoringCredentialsListDto) HasPrometheusUrl() bool`

HasPrometheusUrl returns a boolean if a field has been set.

### SetPrometheusUrlNil

`func (o *MonitoringCredentialsListDto) SetPrometheusUrlNil(b bool)`

 SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil

### UnsetPrometheusUrl
`func (o *MonitoringCredentialsListDto) UnsetPrometheusUrl()`

UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
### GetLokiUrl

`func (o *MonitoringCredentialsListDto) GetLokiUrl() string`

GetLokiUrl returns the LokiUrl field if non-nil, zero value otherwise.

### GetLokiUrlOk

`func (o *MonitoringCredentialsListDto) GetLokiUrlOk() (*string, bool)`

GetLokiUrlOk returns a tuple with the LokiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiUrl

`func (o *MonitoringCredentialsListDto) SetLokiUrl(v string)`

SetLokiUrl sets LokiUrl field to given value.

### HasLokiUrl

`func (o *MonitoringCredentialsListDto) HasLokiUrl() bool`

HasLokiUrl returns a boolean if a field has been set.

### SetLokiUrlNil

`func (o *MonitoringCredentialsListDto) SetLokiUrlNil(b bool)`

 SetLokiUrlNil sets the value for LokiUrl to be an explicit nil

### UnsetLokiUrl
`func (o *MonitoringCredentialsListDto) UnsetLokiUrl()`

UnsetLokiUrl ensures that no value is present for LokiUrl, not even an explicit nil
### GetAlertManagerUrl

`func (o *MonitoringCredentialsListDto) GetAlertManagerUrl() string`

GetAlertManagerUrl returns the AlertManagerUrl field if non-nil, zero value otherwise.

### GetAlertManagerUrlOk

`func (o *MonitoringCredentialsListDto) GetAlertManagerUrlOk() (*string, bool)`

GetAlertManagerUrlOk returns a tuple with the AlertManagerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertManagerUrl

`func (o *MonitoringCredentialsListDto) SetAlertManagerUrl(v string)`

SetAlertManagerUrl sets AlertManagerUrl field to given value.

### HasAlertManagerUrl

`func (o *MonitoringCredentialsListDto) HasAlertManagerUrl() bool`

HasAlertManagerUrl returns a boolean if a field has been set.

### SetAlertManagerUrlNil

`func (o *MonitoringCredentialsListDto) SetAlertManagerUrlNil(b bool)`

 SetAlertManagerUrlNil sets the value for AlertManagerUrl to be an explicit nil

### UnsetAlertManagerUrl
`func (o *MonitoringCredentialsListDto) UnsetAlertManagerUrl()`

UnsetAlertManagerUrl ensures that no value is present for AlertManagerUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


