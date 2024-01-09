# KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest

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
**NetworkId** | **string** |  | 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest

`func NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest(name string, agentPoolName string, description string, networkId string, ) *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest`

NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest instantiates a new KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest`

NewKeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.

### GetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHours() []KeyfactorWebKeyfactorApiModelsSslQuietHourRequest`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHoursOk() (*[]KeyfactorWebKeyfactorApiModelsSslQuietHourRequest, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetQuietHours(v []KeyfactorWebKeyfactorApiModelsSslQuietHourRequest)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.

### GetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslUpdateNetworkRequest) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


