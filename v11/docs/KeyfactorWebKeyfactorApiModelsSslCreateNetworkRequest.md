# KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AgentPoolName** | **string** |  | 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**DiscoverSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**MonitorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**SslAlertRecipients** | Pointer to **[]string** |  | [optional] 
**AutoMonitor** | Pointer to **bool** |  | [optional] 
**GetRobots** | Pointer to **bool** |  | [optional] 
**DiscoverTimeoutMs** | Pointer to **float64** |  | [optional] 
**MonitorTimeoutMs** | Pointer to **float64** |  | [optional] 
**ExpirationAlertDays** | Pointer to **float64** |  | [optional] 
**QuietHours** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSslQuietHourRequest**](KeyfactorWebKeyfactorApiModelsSslQuietHourRequest.md) |  | [optional] 
**BlackoutStart** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**BlackoutEnd** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest

`func NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest(name string, agentPoolName string, description string, ) *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest`

NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest instantiates a new KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest`

NewKeyfactorWebKeyfactorApiModelsSslCreateNetworkRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.

### GetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetQuietHours() []KeyfactorWebKeyfactorApiModelsSslQuietHourRequest`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetQuietHoursOk() (*[]KeyfactorWebKeyfactorApiModelsSslQuietHourRequest, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetQuietHours(v []KeyfactorWebKeyfactorApiModelsSslQuietHourRequest)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslCreateNetworkRequest) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


