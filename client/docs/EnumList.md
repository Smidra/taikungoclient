# EnumList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ProjectStatuses** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ServerRoles** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ServerStatuses** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**UserRoles** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**SecurityGroupRules** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**PrometheusTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**AuditLogs** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**RebootOptions** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**Availability** | Pointer to [**[]CommonAvailabilityDto**](CommonAvailabilityDto.md) |  | [optional] 
**SlackTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**RequestLogs** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**AzureQuotas** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ShowbackKinds** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**AlertTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ReminderTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**AwsPlatforms** | Pointer to [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | [optional] 
**CronPeriods** | Pointer to [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | [optional] 
**ValidityPeriods** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**AlertingIntegrationTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**GoogleImageTypes** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**StandaloneVmStatuses** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**OpenstackContinents** | Pointer to [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | [optional] 
**RetentionPeriods** | Pointer to [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | [optional] 
**TicketPriorities** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**ProxmoxRoles** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 

## Methods

### NewEnumList

`func NewEnumList() *EnumList`

NewEnumList instantiates a new EnumList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumListWithDefaults

`func NewEnumListWithDefaults() *EnumList`

NewEnumListWithDefaults instantiates a new EnumList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudTypes

`func (o *EnumList) GetCloudTypes() []CommonDropdownDto`

GetCloudTypes returns the CloudTypes field if non-nil, zero value otherwise.

### GetCloudTypesOk

`func (o *EnumList) GetCloudTypesOk() (*[]CommonDropdownDto, bool)`

GetCloudTypesOk returns a tuple with the CloudTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTypes

`func (o *EnumList) SetCloudTypes(v []CommonDropdownDto)`

SetCloudTypes sets CloudTypes field to given value.

### HasCloudTypes

`func (o *EnumList) HasCloudTypes() bool`

HasCloudTypes returns a boolean if a field has been set.

### SetCloudTypesNil

`func (o *EnumList) SetCloudTypesNil(b bool)`

 SetCloudTypesNil sets the value for CloudTypes to be an explicit nil

### UnsetCloudTypes
`func (o *EnumList) UnsetCloudTypes()`

UnsetCloudTypes ensures that no value is present for CloudTypes, not even an explicit nil
### GetProjectStatuses

`func (o *EnumList) GetProjectStatuses() []CommonDropdownDto`

GetProjectStatuses returns the ProjectStatuses field if non-nil, zero value otherwise.

### GetProjectStatusesOk

`func (o *EnumList) GetProjectStatusesOk() (*[]CommonDropdownDto, bool)`

GetProjectStatusesOk returns a tuple with the ProjectStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatuses

`func (o *EnumList) SetProjectStatuses(v []CommonDropdownDto)`

SetProjectStatuses sets ProjectStatuses field to given value.

### HasProjectStatuses

`func (o *EnumList) HasProjectStatuses() bool`

HasProjectStatuses returns a boolean if a field has been set.

### SetProjectStatusesNil

`func (o *EnumList) SetProjectStatusesNil(b bool)`

 SetProjectStatusesNil sets the value for ProjectStatuses to be an explicit nil

### UnsetProjectStatuses
`func (o *EnumList) UnsetProjectStatuses()`

UnsetProjectStatuses ensures that no value is present for ProjectStatuses, not even an explicit nil
### GetServerRoles

`func (o *EnumList) GetServerRoles() []CommonDropdownDto`

GetServerRoles returns the ServerRoles field if non-nil, zero value otherwise.

### GetServerRolesOk

`func (o *EnumList) GetServerRolesOk() (*[]CommonDropdownDto, bool)`

GetServerRolesOk returns a tuple with the ServerRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoles

`func (o *EnumList) SetServerRoles(v []CommonDropdownDto)`

SetServerRoles sets ServerRoles field to given value.

### HasServerRoles

`func (o *EnumList) HasServerRoles() bool`

HasServerRoles returns a boolean if a field has been set.

### SetServerRolesNil

`func (o *EnumList) SetServerRolesNil(b bool)`

 SetServerRolesNil sets the value for ServerRoles to be an explicit nil

### UnsetServerRoles
`func (o *EnumList) UnsetServerRoles()`

UnsetServerRoles ensures that no value is present for ServerRoles, not even an explicit nil
### GetServerStatuses

`func (o *EnumList) GetServerStatuses() []CommonDropdownDto`

GetServerStatuses returns the ServerStatuses field if non-nil, zero value otherwise.

### GetServerStatusesOk

`func (o *EnumList) GetServerStatusesOk() (*[]CommonDropdownDto, bool)`

GetServerStatusesOk returns a tuple with the ServerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatuses

`func (o *EnumList) SetServerStatuses(v []CommonDropdownDto)`

SetServerStatuses sets ServerStatuses field to given value.

### HasServerStatuses

`func (o *EnumList) HasServerStatuses() bool`

HasServerStatuses returns a boolean if a field has been set.

### SetServerStatusesNil

`func (o *EnumList) SetServerStatusesNil(b bool)`

 SetServerStatusesNil sets the value for ServerStatuses to be an explicit nil

### UnsetServerStatuses
`func (o *EnumList) UnsetServerStatuses()`

UnsetServerStatuses ensures that no value is present for ServerStatuses, not even an explicit nil
### GetUserRoles

`func (o *EnumList) GetUserRoles() []CommonDropdownDto`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *EnumList) GetUserRolesOk() (*[]CommonDropdownDto, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *EnumList) SetUserRoles(v []CommonDropdownDto)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *EnumList) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.

### SetUserRolesNil

`func (o *EnumList) SetUserRolesNil(b bool)`

 SetUserRolesNil sets the value for UserRoles to be an explicit nil

### UnsetUserRoles
`func (o *EnumList) UnsetUserRoles()`

UnsetUserRoles ensures that no value is present for UserRoles, not even an explicit nil
### GetSecurityGroupRules

`func (o *EnumList) GetSecurityGroupRules() []CommonDropdownDto`

GetSecurityGroupRules returns the SecurityGroupRules field if non-nil, zero value otherwise.

### GetSecurityGroupRulesOk

`func (o *EnumList) GetSecurityGroupRulesOk() (*[]CommonDropdownDto, bool)`

GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRules

`func (o *EnumList) SetSecurityGroupRules(v []CommonDropdownDto)`

SetSecurityGroupRules sets SecurityGroupRules field to given value.

### HasSecurityGroupRules

`func (o *EnumList) HasSecurityGroupRules() bool`

HasSecurityGroupRules returns a boolean if a field has been set.

### SetSecurityGroupRulesNil

`func (o *EnumList) SetSecurityGroupRulesNil(b bool)`

 SetSecurityGroupRulesNil sets the value for SecurityGroupRules to be an explicit nil

### UnsetSecurityGroupRules
`func (o *EnumList) UnsetSecurityGroupRules()`

UnsetSecurityGroupRules ensures that no value is present for SecurityGroupRules, not even an explicit nil
### GetPrometheusTypes

`func (o *EnumList) GetPrometheusTypes() []CommonDropdownDto`

GetPrometheusTypes returns the PrometheusTypes field if non-nil, zero value otherwise.

### GetPrometheusTypesOk

`func (o *EnumList) GetPrometheusTypesOk() (*[]CommonDropdownDto, bool)`

GetPrometheusTypesOk returns a tuple with the PrometheusTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusTypes

`func (o *EnumList) SetPrometheusTypes(v []CommonDropdownDto)`

SetPrometheusTypes sets PrometheusTypes field to given value.

### HasPrometheusTypes

`func (o *EnumList) HasPrometheusTypes() bool`

HasPrometheusTypes returns a boolean if a field has been set.

### SetPrometheusTypesNil

`func (o *EnumList) SetPrometheusTypesNil(b bool)`

 SetPrometheusTypesNil sets the value for PrometheusTypes to be an explicit nil

### UnsetPrometheusTypes
`func (o *EnumList) UnsetPrometheusTypes()`

UnsetPrometheusTypes ensures that no value is present for PrometheusTypes, not even an explicit nil
### GetAuditLogs

`func (o *EnumList) GetAuditLogs() []CommonDropdownDto`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *EnumList) GetAuditLogsOk() (*[]CommonDropdownDto, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *EnumList) SetAuditLogs(v []CommonDropdownDto)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *EnumList) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### SetAuditLogsNil

`func (o *EnumList) SetAuditLogsNil(b bool)`

 SetAuditLogsNil sets the value for AuditLogs to be an explicit nil

### UnsetAuditLogs
`func (o *EnumList) UnsetAuditLogs()`

UnsetAuditLogs ensures that no value is present for AuditLogs, not even an explicit nil
### GetRebootOptions

`func (o *EnumList) GetRebootOptions() []CommonDropdownDto`

GetRebootOptions returns the RebootOptions field if non-nil, zero value otherwise.

### GetRebootOptionsOk

`func (o *EnumList) GetRebootOptionsOk() (*[]CommonDropdownDto, bool)`

GetRebootOptionsOk returns a tuple with the RebootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootOptions

`func (o *EnumList) SetRebootOptions(v []CommonDropdownDto)`

SetRebootOptions sets RebootOptions field to given value.

### HasRebootOptions

`func (o *EnumList) HasRebootOptions() bool`

HasRebootOptions returns a boolean if a field has been set.

### SetRebootOptionsNil

`func (o *EnumList) SetRebootOptionsNil(b bool)`

 SetRebootOptionsNil sets the value for RebootOptions to be an explicit nil

### UnsetRebootOptions
`func (o *EnumList) UnsetRebootOptions()`

UnsetRebootOptions ensures that no value is present for RebootOptions, not even an explicit nil
### GetAvailability

`func (o *EnumList) GetAvailability() []CommonAvailabilityDto`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *EnumList) GetAvailabilityOk() (*[]CommonAvailabilityDto, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *EnumList) SetAvailability(v []CommonAvailabilityDto)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *EnumList) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *EnumList) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *EnumList) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetSlackTypes

`func (o *EnumList) GetSlackTypes() []CommonDropdownDto`

GetSlackTypes returns the SlackTypes field if non-nil, zero value otherwise.

### GetSlackTypesOk

`func (o *EnumList) GetSlackTypesOk() (*[]CommonDropdownDto, bool)`

GetSlackTypesOk returns a tuple with the SlackTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackTypes

`func (o *EnumList) SetSlackTypes(v []CommonDropdownDto)`

SetSlackTypes sets SlackTypes field to given value.

### HasSlackTypes

`func (o *EnumList) HasSlackTypes() bool`

HasSlackTypes returns a boolean if a field has been set.

### SetSlackTypesNil

`func (o *EnumList) SetSlackTypesNil(b bool)`

 SetSlackTypesNil sets the value for SlackTypes to be an explicit nil

### UnsetSlackTypes
`func (o *EnumList) UnsetSlackTypes()`

UnsetSlackTypes ensures that no value is present for SlackTypes, not even an explicit nil
### GetRequestLogs

`func (o *EnumList) GetRequestLogs() []CommonDropdownDto`

GetRequestLogs returns the RequestLogs field if non-nil, zero value otherwise.

### GetRequestLogsOk

`func (o *EnumList) GetRequestLogsOk() (*[]CommonDropdownDto, bool)`

GetRequestLogsOk returns a tuple with the RequestLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLogs

`func (o *EnumList) SetRequestLogs(v []CommonDropdownDto)`

SetRequestLogs sets RequestLogs field to given value.

### HasRequestLogs

`func (o *EnumList) HasRequestLogs() bool`

HasRequestLogs returns a boolean if a field has been set.

### SetRequestLogsNil

`func (o *EnumList) SetRequestLogsNil(b bool)`

 SetRequestLogsNil sets the value for RequestLogs to be an explicit nil

### UnsetRequestLogs
`func (o *EnumList) UnsetRequestLogs()`

UnsetRequestLogs ensures that no value is present for RequestLogs, not even an explicit nil
### GetAzureQuotas

`func (o *EnumList) GetAzureQuotas() []CommonDropdownDto`

GetAzureQuotas returns the AzureQuotas field if non-nil, zero value otherwise.

### GetAzureQuotasOk

`func (o *EnumList) GetAzureQuotasOk() (*[]CommonDropdownDto, bool)`

GetAzureQuotasOk returns a tuple with the AzureQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureQuotas

`func (o *EnumList) SetAzureQuotas(v []CommonDropdownDto)`

SetAzureQuotas sets AzureQuotas field to given value.

### HasAzureQuotas

`func (o *EnumList) HasAzureQuotas() bool`

HasAzureQuotas returns a boolean if a field has been set.

### SetAzureQuotasNil

`func (o *EnumList) SetAzureQuotasNil(b bool)`

 SetAzureQuotasNil sets the value for AzureQuotas to be an explicit nil

### UnsetAzureQuotas
`func (o *EnumList) UnsetAzureQuotas()`

UnsetAzureQuotas ensures that no value is present for AzureQuotas, not even an explicit nil
### GetShowbackKinds

`func (o *EnumList) GetShowbackKinds() []CommonDropdownDto`

GetShowbackKinds returns the ShowbackKinds field if non-nil, zero value otherwise.

### GetShowbackKindsOk

`func (o *EnumList) GetShowbackKindsOk() (*[]CommonDropdownDto, bool)`

GetShowbackKindsOk returns a tuple with the ShowbackKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackKinds

`func (o *EnumList) SetShowbackKinds(v []CommonDropdownDto)`

SetShowbackKinds sets ShowbackKinds field to given value.

### HasShowbackKinds

`func (o *EnumList) HasShowbackKinds() bool`

HasShowbackKinds returns a boolean if a field has been set.

### SetShowbackKindsNil

`func (o *EnumList) SetShowbackKindsNil(b bool)`

 SetShowbackKindsNil sets the value for ShowbackKinds to be an explicit nil

### UnsetShowbackKinds
`func (o *EnumList) UnsetShowbackKinds()`

UnsetShowbackKinds ensures that no value is present for ShowbackKinds, not even an explicit nil
### GetAlertTypes

`func (o *EnumList) GetAlertTypes() []CommonDropdownDto`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *EnumList) GetAlertTypesOk() (*[]CommonDropdownDto, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *EnumList) SetAlertTypes(v []CommonDropdownDto)`

SetAlertTypes sets AlertTypes field to given value.

### HasAlertTypes

`func (o *EnumList) HasAlertTypes() bool`

HasAlertTypes returns a boolean if a field has been set.

### SetAlertTypesNil

`func (o *EnumList) SetAlertTypesNil(b bool)`

 SetAlertTypesNil sets the value for AlertTypes to be an explicit nil

### UnsetAlertTypes
`func (o *EnumList) UnsetAlertTypes()`

UnsetAlertTypes ensures that no value is present for AlertTypes, not even an explicit nil
### GetReminderTypes

`func (o *EnumList) GetReminderTypes() []CommonDropdownDto`

GetReminderTypes returns the ReminderTypes field if non-nil, zero value otherwise.

### GetReminderTypesOk

`func (o *EnumList) GetReminderTypesOk() (*[]CommonDropdownDto, bool)`

GetReminderTypesOk returns a tuple with the ReminderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTypes

`func (o *EnumList) SetReminderTypes(v []CommonDropdownDto)`

SetReminderTypes sets ReminderTypes field to given value.

### HasReminderTypes

`func (o *EnumList) HasReminderTypes() bool`

HasReminderTypes returns a boolean if a field has been set.

### SetReminderTypesNil

`func (o *EnumList) SetReminderTypesNil(b bool)`

 SetReminderTypesNil sets the value for ReminderTypes to be an explicit nil

### UnsetReminderTypes
`func (o *EnumList) UnsetReminderTypes()`

UnsetReminderTypes ensures that no value is present for ReminderTypes, not even an explicit nil
### GetAwsPlatforms

`func (o *EnumList) GetAwsPlatforms() []CommonStringBasedDropdownDto`

GetAwsPlatforms returns the AwsPlatforms field if non-nil, zero value otherwise.

### GetAwsPlatformsOk

`func (o *EnumList) GetAwsPlatformsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetAwsPlatformsOk returns a tuple with the AwsPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPlatforms

`func (o *EnumList) SetAwsPlatforms(v []CommonStringBasedDropdownDto)`

SetAwsPlatforms sets AwsPlatforms field to given value.

### HasAwsPlatforms

`func (o *EnumList) HasAwsPlatforms() bool`

HasAwsPlatforms returns a boolean if a field has been set.

### SetAwsPlatformsNil

`func (o *EnumList) SetAwsPlatformsNil(b bool)`

 SetAwsPlatformsNil sets the value for AwsPlatforms to be an explicit nil

### UnsetAwsPlatforms
`func (o *EnumList) UnsetAwsPlatforms()`

UnsetAwsPlatforms ensures that no value is present for AwsPlatforms, not even an explicit nil
### GetCronPeriods

`func (o *EnumList) GetCronPeriods() []CommonStringBasedDropdownDto`

GetCronPeriods returns the CronPeriods field if non-nil, zero value otherwise.

### GetCronPeriodsOk

`func (o *EnumList) GetCronPeriodsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetCronPeriodsOk returns a tuple with the CronPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronPeriods

`func (o *EnumList) SetCronPeriods(v []CommonStringBasedDropdownDto)`

SetCronPeriods sets CronPeriods field to given value.

### HasCronPeriods

`func (o *EnumList) HasCronPeriods() bool`

HasCronPeriods returns a boolean if a field has been set.

### SetCronPeriodsNil

`func (o *EnumList) SetCronPeriodsNil(b bool)`

 SetCronPeriodsNil sets the value for CronPeriods to be an explicit nil

### UnsetCronPeriods
`func (o *EnumList) UnsetCronPeriods()`

UnsetCronPeriods ensures that no value is present for CronPeriods, not even an explicit nil
### GetValidityPeriods

`func (o *EnumList) GetValidityPeriods() []CommonDropdownDto`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *EnumList) GetValidityPeriodsOk() (*[]CommonDropdownDto, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *EnumList) SetValidityPeriods(v []CommonDropdownDto)`

SetValidityPeriods sets ValidityPeriods field to given value.

### HasValidityPeriods

`func (o *EnumList) HasValidityPeriods() bool`

HasValidityPeriods returns a boolean if a field has been set.

### SetValidityPeriodsNil

`func (o *EnumList) SetValidityPeriodsNil(b bool)`

 SetValidityPeriodsNil sets the value for ValidityPeriods to be an explicit nil

### UnsetValidityPeriods
`func (o *EnumList) UnsetValidityPeriods()`

UnsetValidityPeriods ensures that no value is present for ValidityPeriods, not even an explicit nil
### GetAlertingIntegrationTypes

`func (o *EnumList) GetAlertingIntegrationTypes() []CommonDropdownDto`

GetAlertingIntegrationTypes returns the AlertingIntegrationTypes field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypesOk

`func (o *EnumList) GetAlertingIntegrationTypesOk() (*[]CommonDropdownDto, bool)`

GetAlertingIntegrationTypesOk returns a tuple with the AlertingIntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationTypes

`func (o *EnumList) SetAlertingIntegrationTypes(v []CommonDropdownDto)`

SetAlertingIntegrationTypes sets AlertingIntegrationTypes field to given value.

### HasAlertingIntegrationTypes

`func (o *EnumList) HasAlertingIntegrationTypes() bool`

HasAlertingIntegrationTypes returns a boolean if a field has been set.

### SetAlertingIntegrationTypesNil

`func (o *EnumList) SetAlertingIntegrationTypesNil(b bool)`

 SetAlertingIntegrationTypesNil sets the value for AlertingIntegrationTypes to be an explicit nil

### UnsetAlertingIntegrationTypes
`func (o *EnumList) UnsetAlertingIntegrationTypes()`

UnsetAlertingIntegrationTypes ensures that no value is present for AlertingIntegrationTypes, not even an explicit nil
### GetGoogleImageTypes

`func (o *EnumList) GetGoogleImageTypes() []CommonDropdownDto`

GetGoogleImageTypes returns the GoogleImageTypes field if non-nil, zero value otherwise.

### GetGoogleImageTypesOk

`func (o *EnumList) GetGoogleImageTypesOk() (*[]CommonDropdownDto, bool)`

GetGoogleImageTypesOk returns a tuple with the GoogleImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleImageTypes

`func (o *EnumList) SetGoogleImageTypes(v []CommonDropdownDto)`

SetGoogleImageTypes sets GoogleImageTypes field to given value.

### HasGoogleImageTypes

`func (o *EnumList) HasGoogleImageTypes() bool`

HasGoogleImageTypes returns a boolean if a field has been set.

### SetGoogleImageTypesNil

`func (o *EnumList) SetGoogleImageTypesNil(b bool)`

 SetGoogleImageTypesNil sets the value for GoogleImageTypes to be an explicit nil

### UnsetGoogleImageTypes
`func (o *EnumList) UnsetGoogleImageTypes()`

UnsetGoogleImageTypes ensures that no value is present for GoogleImageTypes, not even an explicit nil
### GetStandaloneVmStatuses

`func (o *EnumList) GetStandaloneVmStatuses() []CommonDropdownDto`

GetStandaloneVmStatuses returns the StandaloneVmStatuses field if non-nil, zero value otherwise.

### GetStandaloneVmStatusesOk

`func (o *EnumList) GetStandaloneVmStatusesOk() (*[]CommonDropdownDto, bool)`

GetStandaloneVmStatusesOk returns a tuple with the StandaloneVmStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmStatuses

`func (o *EnumList) SetStandaloneVmStatuses(v []CommonDropdownDto)`

SetStandaloneVmStatuses sets StandaloneVmStatuses field to given value.

### HasStandaloneVmStatuses

`func (o *EnumList) HasStandaloneVmStatuses() bool`

HasStandaloneVmStatuses returns a boolean if a field has been set.

### SetStandaloneVmStatusesNil

`func (o *EnumList) SetStandaloneVmStatusesNil(b bool)`

 SetStandaloneVmStatusesNil sets the value for StandaloneVmStatuses to be an explicit nil

### UnsetStandaloneVmStatuses
`func (o *EnumList) UnsetStandaloneVmStatuses()`

UnsetStandaloneVmStatuses ensures that no value is present for StandaloneVmStatuses, not even an explicit nil
### GetOpenstackContinents

`func (o *EnumList) GetOpenstackContinents() []CommonStringBasedDropdownDto`

GetOpenstackContinents returns the OpenstackContinents field if non-nil, zero value otherwise.

### GetOpenstackContinentsOk

`func (o *EnumList) GetOpenstackContinentsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetOpenstackContinentsOk returns a tuple with the OpenstackContinents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackContinents

`func (o *EnumList) SetOpenstackContinents(v []CommonStringBasedDropdownDto)`

SetOpenstackContinents sets OpenstackContinents field to given value.

### HasOpenstackContinents

`func (o *EnumList) HasOpenstackContinents() bool`

HasOpenstackContinents returns a boolean if a field has been set.

### SetOpenstackContinentsNil

`func (o *EnumList) SetOpenstackContinentsNil(b bool)`

 SetOpenstackContinentsNil sets the value for OpenstackContinents to be an explicit nil

### UnsetOpenstackContinents
`func (o *EnumList) UnsetOpenstackContinents()`

UnsetOpenstackContinents ensures that no value is present for OpenstackContinents, not even an explicit nil
### GetRetentionPeriods

`func (o *EnumList) GetRetentionPeriods() []CommonStringBasedDropdownDto`

GetRetentionPeriods returns the RetentionPeriods field if non-nil, zero value otherwise.

### GetRetentionPeriodsOk

`func (o *EnumList) GetRetentionPeriodsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetRetentionPeriodsOk returns a tuple with the RetentionPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriods

`func (o *EnumList) SetRetentionPeriods(v []CommonStringBasedDropdownDto)`

SetRetentionPeriods sets RetentionPeriods field to given value.

### HasRetentionPeriods

`func (o *EnumList) HasRetentionPeriods() bool`

HasRetentionPeriods returns a boolean if a field has been set.

### SetRetentionPeriodsNil

`func (o *EnumList) SetRetentionPeriodsNil(b bool)`

 SetRetentionPeriodsNil sets the value for RetentionPeriods to be an explicit nil

### UnsetRetentionPeriods
`func (o *EnumList) UnsetRetentionPeriods()`

UnsetRetentionPeriods ensures that no value is present for RetentionPeriods, not even an explicit nil
### GetTicketPriorities

`func (o *EnumList) GetTicketPriorities() []CommonDropdownDto`

GetTicketPriorities returns the TicketPriorities field if non-nil, zero value otherwise.

### GetTicketPrioritiesOk

`func (o *EnumList) GetTicketPrioritiesOk() (*[]CommonDropdownDto, bool)`

GetTicketPrioritiesOk returns a tuple with the TicketPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketPriorities

`func (o *EnumList) SetTicketPriorities(v []CommonDropdownDto)`

SetTicketPriorities sets TicketPriorities field to given value.

### HasTicketPriorities

`func (o *EnumList) HasTicketPriorities() bool`

HasTicketPriorities returns a boolean if a field has been set.

### SetTicketPrioritiesNil

`func (o *EnumList) SetTicketPrioritiesNil(b bool)`

 SetTicketPrioritiesNil sets the value for TicketPriorities to be an explicit nil

### UnsetTicketPriorities
`func (o *EnumList) UnsetTicketPriorities()`

UnsetTicketPriorities ensures that no value is present for TicketPriorities, not even an explicit nil
### GetProxmoxRoles

`func (o *EnumList) GetProxmoxRoles() []CommonDropdownDto`

GetProxmoxRoles returns the ProxmoxRoles field if non-nil, zero value otherwise.

### GetProxmoxRolesOk

`func (o *EnumList) GetProxmoxRolesOk() (*[]CommonDropdownDto, bool)`

GetProxmoxRolesOk returns a tuple with the ProxmoxRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxRoles

`func (o *EnumList) SetProxmoxRoles(v []CommonDropdownDto)`

SetProxmoxRoles sets ProxmoxRoles field to given value.

### HasProxmoxRoles

`func (o *EnumList) HasProxmoxRoles() bool`

HasProxmoxRoles returns a boolean if a field has been set.

### SetProxmoxRolesNil

`func (o *EnumList) SetProxmoxRolesNil(b bool)`

 SetProxmoxRolesNil sets the value for ProxmoxRoles to be an explicit nil

### UnsetProxmoxRoles
`func (o *EnumList) UnsetProxmoxRoles()`

UnsetProxmoxRoles ensures that no value is present for ProxmoxRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


