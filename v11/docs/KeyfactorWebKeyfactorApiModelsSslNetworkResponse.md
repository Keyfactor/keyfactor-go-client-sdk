# KeyfactorWebKeyfactorApiModelsSslNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**DiscoverStatus** | Pointer to **int32** |  | [optional] 
**MonitorStatus** | Pointer to **int32** |  | [optional] 
**DiscoverLastScanned** | Pointer to **NullableTime** |  | [optional] 
**MonitorLastScanned** | Pointer to **NullableTime** |  | [optional] 
**SslAlertRecipients** | Pointer to **[]string** |  | [optional] 
**GetRobots** | Pointer to **bool** |  | [optional] 
**DiscoverTimeoutMs** | Pointer to **float64** |  | [optional] 
**MonitorTimeoutMs** | Pointer to **float64** |  | [optional] 
**ExpirationAlertDays** | Pointer to **float64** |  | [optional] 
**DiscoverJobParts** | Pointer to **int32** |  | [optional] 
**MonitorJobParts** | Pointer to **int32** |  | [optional] 
**QuietHours** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSslQuietHourResponse**](KeyfactorWebKeyfactorApiModelsSslQuietHourResponse.md) |  | [optional] 
**BlackoutStart** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**BlackoutEnd** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**AutoMonitor** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSslNetworkResponse

`func NewKeyfactorWebKeyfactorApiModelsSslNetworkResponse() *KeyfactorWebKeyfactorApiModelsSslNetworkResponse`

NewKeyfactorWebKeyfactorApiModelsSslNetworkResponse instantiates a new KeyfactorWebKeyfactorApiModelsSslNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSslNetworkResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSslNetworkResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSslNetworkResponse`

NewKeyfactorWebKeyfactorApiModelsSslNetworkResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSslNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### SetAgentPoolNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetAgentPoolNameNil(b bool)`

 SetAgentPoolNameNil sets the value for AgentPoolName to be an explicit nil

### UnsetAgentPoolName
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetAgentPoolName()`

UnsetAgentPoolName ensures that no value is present for AgentPoolName, not even an explicit nil
### GetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### SetAgentPoolIdNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetAgentPoolIdNil(b bool)`

 SetAgentPoolIdNil sets the value for AgentPoolId to be an explicit nil

### UnsetAgentPoolId
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetAgentPoolId()`

UnsetAgentPoolId ensures that no value is present for AgentPoolId, not even an explicit nil
### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentComplete() float64`

GetDiscoverPercentComplete returns the DiscoverPercentComplete field if non-nil, zero value otherwise.

### GetDiscoverPercentCompleteOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentCompleteOk() (*float64, bool)`

GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverPercentComplete(v float64)`

SetDiscoverPercentComplete sets DiscoverPercentComplete field to given value.

### HasDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverPercentComplete() bool`

HasDiscoverPercentComplete returns a boolean if a field has been set.

### GetMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorPercentComplete() float64`

GetMonitorPercentComplete returns the MonitorPercentComplete field if non-nil, zero value otherwise.

### GetMonitorPercentCompleteOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorPercentCompleteOk() (*float64, bool)`

GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorPercentComplete(v float64)`

SetMonitorPercentComplete sets MonitorPercentComplete field to given value.

### HasMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorPercentComplete() bool`

HasMonitorPercentComplete returns a boolean if a field has been set.

### GetDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverStatus() int32`

GetDiscoverStatus returns the DiscoverStatus field if non-nil, zero value otherwise.

### GetDiscoverStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverStatusOk() (*int32, bool)`

GetDiscoverStatusOk returns a tuple with the DiscoverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverStatus(v int32)`

SetDiscoverStatus sets DiscoverStatus field to given value.

### HasDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverStatus() bool`

HasDiscoverStatus returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorStatus() int32`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorStatusOk() (*int32, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorStatus(v int32)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScanned() time.Time`

GetDiscoverLastScanned returns the DiscoverLastScanned field if non-nil, zero value otherwise.

### GetDiscoverLastScannedOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScannedOk() (*time.Time, bool)`

GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverLastScanned(v time.Time)`

SetDiscoverLastScanned sets DiscoverLastScanned field to given value.

### HasDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverLastScanned() bool`

HasDiscoverLastScanned returns a boolean if a field has been set.

### SetDiscoverLastScannedNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverLastScannedNil(b bool)`

 SetDiscoverLastScannedNil sets the value for DiscoverLastScanned to be an explicit nil

### UnsetDiscoverLastScanned
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetDiscoverLastScanned()`

UnsetDiscoverLastScanned ensures that no value is present for DiscoverLastScanned, not even an explicit nil
### GetMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorLastScanned() time.Time`

GetMonitorLastScanned returns the MonitorLastScanned field if non-nil, zero value otherwise.

### GetMonitorLastScannedOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorLastScannedOk() (*time.Time, bool)`

GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorLastScanned(v time.Time)`

SetMonitorLastScanned sets MonitorLastScanned field to given value.

### HasMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorLastScanned() bool`

HasMonitorLastScanned returns a boolean if a field has been set.

### SetMonitorLastScannedNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorLastScannedNil(b bool)`

 SetMonitorLastScannedNil sets the value for MonitorLastScanned to be an explicit nil

### UnsetMonitorLastScanned
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetMonitorLastScanned()`

UnsetMonitorLastScanned ensures that no value is present for MonitorLastScanned, not even an explicit nil
### GetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverJobParts() int32`

GetDiscoverJobParts returns the DiscoverJobParts field if non-nil, zero value otherwise.

### GetDiscoverJobPartsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetDiscoverJobPartsOk() (*int32, bool)`

GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetDiscoverJobParts(v int32)`

SetDiscoverJobParts sets DiscoverJobParts field to given value.

### HasDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasDiscoverJobParts() bool`

HasDiscoverJobParts returns a boolean if a field has been set.

### GetMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorJobParts() int32`

GetMonitorJobParts returns the MonitorJobParts field if non-nil, zero value otherwise.

### GetMonitorJobPartsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetMonitorJobPartsOk() (*int32, bool)`

GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetMonitorJobParts(v int32)`

SetMonitorJobParts sets MonitorJobParts field to given value.

### HasMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasMonitorJobParts() bool`

HasMonitorJobParts returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetQuietHours() []KeyfactorWebKeyfactorApiModelsSslQuietHourResponse`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetQuietHoursOk() (*[]KeyfactorWebKeyfactorApiModelsSslQuietHourResponse, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetQuietHours(v []KeyfactorWebKeyfactorApiModelsSslQuietHourResponse)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.

### GetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAutoMonitor() bool`

GetAutoMonitor returns the AutoMonitor field if non-nil, zero value otherwise.

### GetAutoMonitorOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) GetAutoMonitorOk() (*bool, bool)`

GetAutoMonitorOk returns a tuple with the AutoMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) SetAutoMonitor(v bool)`

SetAutoMonitor sets AutoMonitor field to given value.

### HasAutoMonitor

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkResponse) HasAutoMonitor() bool`

HasAutoMonitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


