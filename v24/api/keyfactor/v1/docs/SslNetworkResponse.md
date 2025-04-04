# SslNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoMonitor** | Pointer to **bool** |  | [optional] 
**NetworkId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AgentPoolName** | Pointer to **NullableString** |  | [optional] 
**AgentPoolId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DiscoverSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**MonitorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**DiscoverPercentComplete** | Pointer to **float64** |  | [optional] 
**MonitorPercentComplete** | Pointer to **float64** |  | [optional] 
**DiscoverStatus** | Pointer to [**CSSCMSCoreEnumsSslNetworkJobStatus**](CSSCMSCoreEnumsSslNetworkJobStatus.md) |  | [optional] 
**MonitorStatus** | Pointer to [**CSSCMSCoreEnumsSslNetworkJobStatus**](CSSCMSCoreEnumsSslNetworkJobStatus.md) |  | [optional] 
**DiscoverLastScanned** | Pointer to **NullableTime** |  | [optional] 
**MonitorLastScanned** | Pointer to **NullableTime** |  | [optional] 
**SslAlertRecipients** | Pointer to **[]string** |  | [optional] 
**GetRobots** | Pointer to **bool** |  | [optional] 
**DiscoverTimeoutMs** | Pointer to **float64** |  | [optional] 
**MonitorTimeoutMs** | Pointer to **float64** |  | [optional] 
**ExpirationAlertDays** | Pointer to **float64** |  | [optional] 
**DiscoverJobParts** | Pointer to **int32** |  | [optional] 
**MonitorJobParts** | Pointer to **int32** |  | [optional] 
**QuietHours** | Pointer to [**[]SslQuietHourResponse**](SslQuietHourResponse.md) |  | [optional] 
**BlackoutStart** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**BlackoutEnd** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 

## Methods

### NewSslNetworkResponse

`func NewSslNetworkResponse() *SslNetworkResponse`

NewSslNetworkResponse instantiates a new SslNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslNetworkResponseWithDefaults

`func NewSslNetworkResponseWithDefaults() *SslNetworkResponse`

NewSslNetworkResponseWithDefaults instantiates a new SslNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoMonitor

`func (o *SslNetworkResponse) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *SslNetworkResponse) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *SslNetworkResponse) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *SslNetworkResponse) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.

### GetNetworkId

`func (o *SslNetworkResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *SslNetworkResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *SslNetworkResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *SslNetworkResponse) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *SslNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslNetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SslNetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SslNetworkResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SslNetworkResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAgentPoolName

`func (o *SslNetworkResponse) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *SslNetworkResponse) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *SslNetworkResponse) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *SslNetworkResponse) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### SetAgentPoolNameNil

`func (o *SslNetworkResponse) SetAgentPoolNameNil(b bool)`

 SetAgentPoolNameNil sets the value for AgentPoolName to be an explicit nil

### UnsetAgentPoolName
`func (o *SslNetworkResponse) UnsetAgentPoolName()`

UnsetAgentPoolName ensures that no value is present for AgentPoolName, not even an explicit nil
### GetAgentPoolId

`func (o *SslNetworkResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *SslNetworkResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *SslNetworkResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *SslNetworkResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### SetAgentPoolIdNil

`func (o *SslNetworkResponse) SetAgentPoolIdNil(b bool)`

 SetAgentPoolIdNil sets the value for AgentPoolId to be an explicit nil

### UnsetAgentPoolId
`func (o *SslNetworkResponse) UnsetAgentPoolId()`

UnsetAgentPoolId ensures that no value is present for AgentPoolId, not even an explicit nil
### GetDescription

`func (o *SslNetworkResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SslNetworkResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SslNetworkResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SslNetworkResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SslNetworkResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SslNetworkResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *SslNetworkResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SslNetworkResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SslNetworkResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SslNetworkResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *SslNetworkResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *SslNetworkResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *SslNetworkResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *SslNetworkResponse) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *SslNetworkResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *SslNetworkResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *SslNetworkResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *SslNetworkResponse) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetDiscoverPercentComplete

`func (o *SslNetworkResponse) GetDiscoverPercentComplete() float64`

GetDiscoverPercentComplete returns the DiscoverPercentComplete field if non-nil, zero value otherwise.

### GetDiscoverPercentCompleteOk

`func (o *SslNetworkResponse) GetDiscoverPercentCompleteOk() (*float64, bool)`

GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPercentComplete

`func (o *SslNetworkResponse) SetDiscoverPercentComplete(v float64)`

SetDiscoverPercentComplete sets DiscoverPercentComplete field to given value.

### HasDiscoverPercentComplete

`func (o *SslNetworkResponse) HasDiscoverPercentComplete() bool`

HasDiscoverPercentComplete returns a boolean if a field has been set.

### GetMonitorPercentComplete

`func (o *SslNetworkResponse) GetMonitorPercentComplete() float64`

GetMonitorPercentComplete returns the MonitorPercentComplete field if non-nil, zero value otherwise.

### GetMonitorPercentCompleteOk

`func (o *SslNetworkResponse) GetMonitorPercentCompleteOk() (*float64, bool)`

GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPercentComplete

`func (o *SslNetworkResponse) SetMonitorPercentComplete(v float64)`

SetMonitorPercentComplete sets MonitorPercentComplete field to given value.

### HasMonitorPercentComplete

`func (o *SslNetworkResponse) HasMonitorPercentComplete() bool`

HasMonitorPercentComplete returns a boolean if a field has been set.

### GetDiscoverStatus

`func (o *SslNetworkResponse) GetDiscoverStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetDiscoverStatus returns the DiscoverStatus field if non-nil, zero value otherwise.

### GetDiscoverStatusOk

`func (o *SslNetworkResponse) GetDiscoverStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetDiscoverStatusOk returns a tuple with the DiscoverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverStatus

`func (o *SslNetworkResponse) SetDiscoverStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetDiscoverStatus sets DiscoverStatus field to given value.

### HasDiscoverStatus

`func (o *SslNetworkResponse) HasDiscoverStatus() bool`

HasDiscoverStatus returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *SslNetworkResponse) GetMonitorStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *SslNetworkResponse) GetMonitorStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *SslNetworkResponse) SetMonitorStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *SslNetworkResponse) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetDiscoverLastScanned

`func (o *SslNetworkResponse) GetDiscoverLastScanned() time.Time`

GetDiscoverLastScanned returns the DiscoverLastScanned field if non-nil, zero value otherwise.

### GetDiscoverLastScannedOk

`func (o *SslNetworkResponse) GetDiscoverLastScannedOk() (*time.Time, bool)`

GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLastScanned

`func (o *SslNetworkResponse) SetDiscoverLastScanned(v time.Time)`

SetDiscoverLastScanned sets DiscoverLastScanned field to given value.

### HasDiscoverLastScanned

`func (o *SslNetworkResponse) HasDiscoverLastScanned() bool`

HasDiscoverLastScanned returns a boolean if a field has been set.

### SetDiscoverLastScannedNil

`func (o *SslNetworkResponse) SetDiscoverLastScannedNil(b bool)`

 SetDiscoverLastScannedNil sets the value for DiscoverLastScanned to be an explicit nil

### UnsetDiscoverLastScanned
`func (o *SslNetworkResponse) UnsetDiscoverLastScanned()`

UnsetDiscoverLastScanned ensures that no value is present for DiscoverLastScanned, not even an explicit nil
### GetMonitorLastScanned

`func (o *SslNetworkResponse) GetMonitorLastScanned() time.Time`

GetMonitorLastScanned returns the MonitorLastScanned field if non-nil, zero value otherwise.

### GetMonitorLastScannedOk

`func (o *SslNetworkResponse) GetMonitorLastScannedOk() (*time.Time, bool)`

GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorLastScanned

`func (o *SslNetworkResponse) SetMonitorLastScanned(v time.Time)`

SetMonitorLastScanned sets MonitorLastScanned field to given value.

### HasMonitorLastScanned

`func (o *SslNetworkResponse) HasMonitorLastScanned() bool`

HasMonitorLastScanned returns a boolean if a field has been set.

### SetMonitorLastScannedNil

`func (o *SslNetworkResponse) SetMonitorLastScannedNil(b bool)`

 SetMonitorLastScannedNil sets the value for MonitorLastScanned to be an explicit nil

### UnsetMonitorLastScanned
`func (o *SslNetworkResponse) UnsetMonitorLastScanned()`

UnsetMonitorLastScanned ensures that no value is present for MonitorLastScanned, not even an explicit nil
### GetSslAlertRecipients

`func (o *SslNetworkResponse) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *SslNetworkResponse) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *SslNetworkResponse) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *SslNetworkResponse) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *SslNetworkResponse) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *SslNetworkResponse) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetGetRobots

`func (o *SslNetworkResponse) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *SslNetworkResponse) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *SslNetworkResponse) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *SslNetworkResponse) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *SslNetworkResponse) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *SslNetworkResponse) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *SslNetworkResponse) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *SslNetworkResponse) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *SslNetworkResponse) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *SslNetworkResponse) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *SslNetworkResponse) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *SslNetworkResponse) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *SslNetworkResponse) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *SslNetworkResponse) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *SslNetworkResponse) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *SslNetworkResponse) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetDiscoverJobParts

`func (o *SslNetworkResponse) GetDiscoverJobParts() int32`

GetDiscoverJobParts returns the DiscoverJobParts field if non-nil, zero value otherwise.

### GetDiscoverJobPartsOk

`func (o *SslNetworkResponse) GetDiscoverJobPartsOk() (*int32, bool)`

GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverJobParts

`func (o *SslNetworkResponse) SetDiscoverJobParts(v int32)`

SetDiscoverJobParts sets DiscoverJobParts field to given value.

### HasDiscoverJobParts

`func (o *SslNetworkResponse) HasDiscoverJobParts() bool`

HasDiscoverJobParts returns a boolean if a field has been set.

### GetMonitorJobParts

`func (o *SslNetworkResponse) GetMonitorJobParts() int32`

GetMonitorJobParts returns the MonitorJobParts field if non-nil, zero value otherwise.

### GetMonitorJobPartsOk

`func (o *SslNetworkResponse) GetMonitorJobPartsOk() (*int32, bool)`

GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobParts

`func (o *SslNetworkResponse) SetMonitorJobParts(v int32)`

SetMonitorJobParts sets MonitorJobParts field to given value.

### HasMonitorJobParts

`func (o *SslNetworkResponse) HasMonitorJobParts() bool`

HasMonitorJobParts returns a boolean if a field has been set.

### GetQuietHours

`func (o *SslNetworkResponse) GetQuietHours() []SslQuietHourResponse`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *SslNetworkResponse) GetQuietHoursOk() (*[]SslQuietHourResponse, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *SslNetworkResponse) SetQuietHours(v []SslQuietHourResponse)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *SslNetworkResponse) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *SslNetworkResponse) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *SslNetworkResponse) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *SslNetworkResponse) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *SslNetworkResponse) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *SslNetworkResponse) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *SslNetworkResponse) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *SslNetworkResponse) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *SslNetworkResponse) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *SslNetworkResponse) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *SslNetworkResponse) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


