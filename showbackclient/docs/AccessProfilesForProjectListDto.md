# AccessProfilesForProjectListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**HttpProxy** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**SshUsers** | Pointer to [**[]SshUserListDto**](SshUserListDto.md) |  | [optional] 
**DnsServers** | Pointer to [**[]DnsServerListDto**](DnsServerListDto.md) |  | [optional] 
**NtpServers** | Pointer to [**[]NtpServerListDto**](NtpServerListDto.md) |  | [optional] 
**AllowedHosts** | Pointer to [**[]AllowedHostListDto**](AllowedHostListDto.md) |  | [optional] 

## Methods

### NewAccessProfilesForProjectListDto

`func NewAccessProfilesForProjectListDto() *AccessProfilesForProjectListDto`

NewAccessProfilesForProjectListDto instantiates a new AccessProfilesForProjectListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfilesForProjectListDtoWithDefaults

`func NewAccessProfilesForProjectListDtoWithDefaults() *AccessProfilesForProjectListDto`

NewAccessProfilesForProjectListDtoWithDefaults instantiates a new AccessProfilesForProjectListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessProfilesForProjectListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessProfilesForProjectListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessProfilesForProjectListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessProfilesForProjectListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessProfilesForProjectListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfilesForProjectListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfilesForProjectListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessProfilesForProjectListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccessProfilesForProjectListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccessProfilesForProjectListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHttpProxy

`func (o *AccessProfilesForProjectListDto) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *AccessProfilesForProjectListDto) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *AccessProfilesForProjectListDto) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *AccessProfilesForProjectListDto) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *AccessProfilesForProjectListDto) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *AccessProfilesForProjectListDto) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetOrganizationId

`func (o *AccessProfilesForProjectListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccessProfilesForProjectListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccessProfilesForProjectListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AccessProfilesForProjectListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *AccessProfilesForProjectListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AccessProfilesForProjectListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *AccessProfilesForProjectListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AccessProfilesForProjectListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AccessProfilesForProjectListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AccessProfilesForProjectListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *AccessProfilesForProjectListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AccessProfilesForProjectListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetRevision

`func (o *AccessProfilesForProjectListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AccessProfilesForProjectListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AccessProfilesForProjectListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AccessProfilesForProjectListDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSshUsers

`func (o *AccessProfilesForProjectListDto) GetSshUsers() []SshUserListDto`

GetSshUsers returns the SshUsers field if non-nil, zero value otherwise.

### GetSshUsersOk

`func (o *AccessProfilesForProjectListDto) GetSshUsersOk() (*[]SshUserListDto, bool)`

GetSshUsersOk returns a tuple with the SshUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsers

`func (o *AccessProfilesForProjectListDto) SetSshUsers(v []SshUserListDto)`

SetSshUsers sets SshUsers field to given value.

### HasSshUsers

`func (o *AccessProfilesForProjectListDto) HasSshUsers() bool`

HasSshUsers returns a boolean if a field has been set.

### SetSshUsersNil

`func (o *AccessProfilesForProjectListDto) SetSshUsersNil(b bool)`

 SetSshUsersNil sets the value for SshUsers to be an explicit nil

### UnsetSshUsers
`func (o *AccessProfilesForProjectListDto) UnsetSshUsers()`

UnsetSshUsers ensures that no value is present for SshUsers, not even an explicit nil
### GetDnsServers

`func (o *AccessProfilesForProjectListDto) GetDnsServers() []DnsServerListDto`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AccessProfilesForProjectListDto) GetDnsServersOk() (*[]DnsServerListDto, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AccessProfilesForProjectListDto) SetDnsServers(v []DnsServerListDto)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AccessProfilesForProjectListDto) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *AccessProfilesForProjectListDto) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *AccessProfilesForProjectListDto) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetNtpServers

`func (o *AccessProfilesForProjectListDto) GetNtpServers() []NtpServerListDto`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *AccessProfilesForProjectListDto) GetNtpServersOk() (*[]NtpServerListDto, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *AccessProfilesForProjectListDto) SetNtpServers(v []NtpServerListDto)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *AccessProfilesForProjectListDto) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *AccessProfilesForProjectListDto) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *AccessProfilesForProjectListDto) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetAllowedHosts

`func (o *AccessProfilesForProjectListDto) GetAllowedHosts() []AllowedHostListDto`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *AccessProfilesForProjectListDto) GetAllowedHostsOk() (*[]AllowedHostListDto, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *AccessProfilesForProjectListDto) SetAllowedHosts(v []AllowedHostListDto)`

SetAllowedHosts sets AllowedHosts field to given value.

### HasAllowedHosts

`func (o *AccessProfilesForProjectListDto) HasAllowedHosts() bool`

HasAllowedHosts returns a boolean if a field has been set.

### SetAllowedHostsNil

`func (o *AccessProfilesForProjectListDto) SetAllowedHostsNil(b bool)`

 SetAllowedHostsNil sets the value for AllowedHosts to be an explicit nil

### UnsetAllowedHosts
`func (o *AccessProfilesForProjectListDto) UnsetAllowedHosts()`

UnsetAllowedHosts ensures that no value is present for AllowedHosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


