# KeyfactorApiModelsSslUpdateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** |  | 
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
**QuietHours** | Pointer to [**[]KeyfactorApiModelsSslQuietHourRequest**](KeyfactorApiModelsSslQuietHourRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsSslUpdateNetworkRequest

`func NewKeyfactorApiModelsSslUpdateNetworkRequest(networkId string, name string, agentPoolName string, description string, ) *KeyfactorApiModelsSslUpdateNetworkRequest`

NewKeyfactorApiModelsSslUpdateNetworkRequest instantiates a new KeyfactorApiModelsSslUpdateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults

`func NewKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults() *KeyfactorApiModelsSslUpdateNetworkRequest`

NewKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults instantiates a new KeyfactorApiModelsSslUpdateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgentPoolName

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.


### GetDescription

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetSslAlertRecipients

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### GetAutoMonitor

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.

### GetGetRobots

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHours() []KeyfactorApiModelsSslQuietHourRequest`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHoursOk() (*[]KeyfactorApiModelsSslQuietHourRequest, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetQuietHours(v []KeyfactorApiModelsSslQuietHourRequest)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


