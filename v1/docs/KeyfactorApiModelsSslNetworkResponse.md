# KeyfactorApiModelsSslNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AgentPoolName** | Pointer to **string** |  | [optional] 
**AgentPoolId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DiscoverSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**MonitorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**DiscoverPercentComplete** | Pointer to **float64** |  | [optional] 
**MonitorPercentComplete** | Pointer to **float64** |  | [optional] 
**DiscoverStatus** | Pointer to **int32** |  | [optional] 
**MonitorStatus** | Pointer to **int32** |  | [optional] 
**DiscoverLastScanned** | Pointer to **time.Time** |  | [optional] 
**MonitorLastScanned** | Pointer to **time.Time** |  | [optional] 
**SslAlertRecipients** | Pointer to **[]string** |  | [optional] 
**AutoMonitor** | Pointer to **bool** |  | [optional] 
**GetRobots** | Pointer to **bool** |  | [optional] 
**DiscoverTimeoutMs** | Pointer to **float64** |  | [optional] 
**MonitorTimeoutMs** | Pointer to **float64** |  | [optional] 
**ExpirationAlertDays** | Pointer to **float64** |  | [optional] 
**DiscoverJobParts** | Pointer to **int32** |  | [optional] 
**MonitorJobParts** | Pointer to **int32** |  | [optional] 
**QuietHours** | Pointer to [**[]KeyfactorApiModelsSslQuietHourResponse**](KeyfactorApiModelsSslQuietHourResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsSslNetworkResponse

`func NewKeyfactorApiModelsSslNetworkResponse() *KeyfactorApiModelsSslNetworkResponse`

NewKeyfactorApiModelsSslNetworkResponse instantiates a new KeyfactorApiModelsSslNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsSslNetworkResponseWithDefaults

`func NewKeyfactorApiModelsSslNetworkResponseWithDefaults() *KeyfactorApiModelsSslNetworkResponse`

NewKeyfactorApiModelsSslNetworkResponseWithDefaults instantiates a new KeyfactorApiModelsSslNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *KeyfactorApiModelsSslNetworkResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KeyfactorApiModelsSslNetworkResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *KeyfactorApiModelsSslNetworkResponse) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorApiModelsSslNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsSslNetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorApiModelsSslNetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAgentPoolName

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorApiModelsSslNetworkResponse) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *KeyfactorApiModelsSslNetworkResponse) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### GetAgentPoolId

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *KeyfactorApiModelsSslNetworkResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *KeyfactorApiModelsSslNetworkResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetDescription

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *KeyfactorApiModelsSslNetworkResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorApiModelsSslNetworkResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorApiModelsSslNetworkResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetDiscoverPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentComplete() float64`

GetDiscoverPercentComplete returns the DiscoverPercentComplete field if non-nil, zero value otherwise.

### GetDiscoverPercentCompleteOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentCompleteOk() (*float64, bool)`

GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverPercentComplete(v float64)`

SetDiscoverPercentComplete sets DiscoverPercentComplete field to given value.

### HasDiscoverPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverPercentComplete() bool`

HasDiscoverPercentComplete returns a boolean if a field has been set.

### GetMonitorPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorPercentComplete() float64`

GetMonitorPercentComplete returns the MonitorPercentComplete field if non-nil, zero value otherwise.

### GetMonitorPercentCompleteOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorPercentCompleteOk() (*float64, bool)`

GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorPercentComplete(v float64)`

SetMonitorPercentComplete sets MonitorPercentComplete field to given value.

### HasMonitorPercentComplete

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorPercentComplete() bool`

HasMonitorPercentComplete returns a boolean if a field has been set.

### GetDiscoverStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverStatus() int32`

GetDiscoverStatus returns the DiscoverStatus field if non-nil, zero value otherwise.

### GetDiscoverStatusOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverStatusOk() (*int32, bool)`

GetDiscoverStatusOk returns a tuple with the DiscoverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverStatus(v int32)`

SetDiscoverStatus sets DiscoverStatus field to given value.

### HasDiscoverStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverStatus() bool`

HasDiscoverStatus returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorStatus() int32`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorStatusOk() (*int32, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorStatus(v int32)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetDiscoverLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScanned() time.Time`

GetDiscoverLastScanned returns the DiscoverLastScanned field if non-nil, zero value otherwise.

### GetDiscoverLastScannedOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScannedOk() (*time.Time, bool)`

GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverLastScanned(v time.Time)`

SetDiscoverLastScanned sets DiscoverLastScanned field to given value.

### HasDiscoverLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverLastScanned() bool`

HasDiscoverLastScanned returns a boolean if a field has been set.

### GetMonitorLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorLastScanned() time.Time`

GetMonitorLastScanned returns the MonitorLastScanned field if non-nil, zero value otherwise.

### GetMonitorLastScannedOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorLastScannedOk() (*time.Time, bool)`

GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorLastScanned(v time.Time)`

SetMonitorLastScanned sets MonitorLastScanned field to given value.

### HasMonitorLastScanned

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorLastScanned() bool`

HasMonitorLastScanned returns a boolean if a field has been set.

### GetSslAlertRecipients

`func (o *KeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorApiModelsSslNetworkResponse) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorApiModelsSslNetworkResponse) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### GetAutoMonitor

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *KeyfactorApiModelsSslNetworkResponse) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *KeyfactorApiModelsSslNetworkResponse) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.

### GetGetRobots

`func (o *KeyfactorApiModelsSslNetworkResponse) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorApiModelsSslNetworkResponse) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorApiModelsSslNetworkResponse) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorApiModelsSslNetworkResponse) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorApiModelsSslNetworkResponse) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetDiscoverJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverJobParts() int32`

GetDiscoverJobParts returns the DiscoverJobParts field if non-nil, zero value otherwise.

### GetDiscoverJobPartsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverJobPartsOk() (*int32, bool)`

GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverJobParts(v int32)`

SetDiscoverJobParts sets DiscoverJobParts field to given value.

### HasDiscoverJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverJobParts() bool`

HasDiscoverJobParts returns a boolean if a field has been set.

### GetMonitorJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorJobParts() int32`

GetMonitorJobParts returns the MonitorJobParts field if non-nil, zero value otherwise.

### GetMonitorJobPartsOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorJobPartsOk() (*int32, bool)`

GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorJobParts(v int32)`

SetMonitorJobParts sets MonitorJobParts field to given value.

### HasMonitorJobParts

`func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorJobParts() bool`

HasMonitorJobParts returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorApiModelsSslNetworkResponse) GetQuietHours() []KeyfactorApiModelsSslQuietHourResponse`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorApiModelsSslNetworkResponse) GetQuietHoursOk() (*[]KeyfactorApiModelsSslQuietHourResponse, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorApiModelsSslNetworkResponse) SetQuietHours(v []KeyfactorApiModelsSslQuietHourResponse)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorApiModelsSslNetworkResponse) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


