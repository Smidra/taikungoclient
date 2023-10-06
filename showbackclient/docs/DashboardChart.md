# DashboardChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**OrganizationEntityForDashboard**](OrganizationEntityForDashboard.md) |  | [optional] 
**Projects** | Pointer to [**ProjectChartDto**](ProjectChartDto.md) |  | [optional] 
**CloudCredentials** | Pointer to [**CredentialChartDto**](CredentialChartDto.md) |  | [optional] 
**Servers** | Pointer to [**ServerChartDto**](ServerChartDto.md) |  | [optional] 
**StandAloneVms** | Pointer to [**ServerChartDto**](ServerChartDto.md) |  | [optional] 

## Methods

### NewDashboardChart

`func NewDashboardChart() *DashboardChart`

NewDashboardChart instantiates a new DashboardChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardChartWithDefaults

`func NewDashboardChartWithDefaults() *DashboardChart`

NewDashboardChartWithDefaults instantiates a new DashboardChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *DashboardChart) GetOrganization() OrganizationEntityForDashboard`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DashboardChart) GetOrganizationOk() (*OrganizationEntityForDashboard, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DashboardChart) SetOrganization(v OrganizationEntityForDashboard)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DashboardChart) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjects

`func (o *DashboardChart) GetProjects() ProjectChartDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *DashboardChart) GetProjectsOk() (*ProjectChartDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *DashboardChart) SetProjects(v ProjectChartDto)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *DashboardChart) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCloudCredentials

`func (o *DashboardChart) GetCloudCredentials() CredentialChartDto`

GetCloudCredentials returns the CloudCredentials field if non-nil, zero value otherwise.

### GetCloudCredentialsOk

`func (o *DashboardChart) GetCloudCredentialsOk() (*CredentialChartDto, bool)`

GetCloudCredentialsOk returns a tuple with the CloudCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentials

`func (o *DashboardChart) SetCloudCredentials(v CredentialChartDto)`

SetCloudCredentials sets CloudCredentials field to given value.

### HasCloudCredentials

`func (o *DashboardChart) HasCloudCredentials() bool`

HasCloudCredentials returns a boolean if a field has been set.

### GetServers

`func (o *DashboardChart) GetServers() ServerChartDto`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DashboardChart) GetServersOk() (*ServerChartDto, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DashboardChart) SetServers(v ServerChartDto)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DashboardChart) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetStandAloneVms

`func (o *DashboardChart) GetStandAloneVms() ServerChartDto`

GetStandAloneVms returns the StandAloneVms field if non-nil, zero value otherwise.

### GetStandAloneVmsOk

`func (o *DashboardChart) GetStandAloneVmsOk() (*ServerChartDto, bool)`

GetStandAloneVmsOk returns a tuple with the StandAloneVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneVms

`func (o *DashboardChart) SetStandAloneVms(v ServerChartDto)`

SetStandAloneVms sets StandAloneVms field to given value.

### HasStandAloneVms

`func (o *DashboardChart) HasStandAloneVms() bool`

HasStandAloneVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


