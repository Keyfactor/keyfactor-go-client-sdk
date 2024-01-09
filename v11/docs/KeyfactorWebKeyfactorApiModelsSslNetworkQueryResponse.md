# KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse

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
**QuietHours** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSslQuietHourResponse**](KeyfactorWebKeyfactorApiModelsSslQuietHourResponse.md) |  | [optional] 
**BlackoutStart** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**BlackoutEnd** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse

`func NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse() *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse`

NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse instantiates a new KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse`

NewKeyfactorWebKeyfactorApiModelsSslNetworkQueryResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### SetAgentPoolNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetAgentPoolNameNil(b bool)`

 SetAgentPoolNameNil sets the value for AgentPoolName to be an explicit nil

### UnsetAgentPoolName
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetAgentPoolName()`

UnsetAgentPoolName ensures that no value is present for AgentPoolName, not even an explicit nil
### GetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### SetAgentPoolIdNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetAgentPoolIdNil(b bool)`

 SetAgentPoolIdNil sets the value for AgentPoolId to be an explicit nil

### UnsetAgentPoolId
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetAgentPoolId()`

UnsetAgentPoolId ensures that no value is present for AgentPoolId, not even an explicit nil
### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverPercentComplete() float64`

GetDiscoverPercentComplete returns the DiscoverPercentComplete field if non-nil, zero value otherwise.

### GetDiscoverPercentCompleteOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverPercentCompleteOk() (*float64, bool)`

GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverPercentComplete(v float64)`

SetDiscoverPercentComplete sets DiscoverPercentComplete field to given value.

### HasDiscoverPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverPercentComplete() bool`

HasDiscoverPercentComplete returns a boolean if a field has been set.

### GetMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorPercentComplete() float64`

GetMonitorPercentComplete returns the MonitorPercentComplete field if non-nil, zero value otherwise.

### GetMonitorPercentCompleteOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorPercentCompleteOk() (*float64, bool)`

GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorPercentComplete(v float64)`

SetMonitorPercentComplete sets MonitorPercentComplete field to given value.

### HasMonitorPercentComplete

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorPercentComplete() bool`

HasMonitorPercentComplete returns a boolean if a field has been set.

### GetDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetDiscoverStatus returns the DiscoverStatus field if non-nil, zero value otherwise.

### GetDiscoverStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetDiscoverStatusOk returns a tuple with the DiscoverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetDiscoverStatus sets DiscoverStatus field to given value.

### HasDiscoverStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverStatus() bool`

HasDiscoverStatus returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverLastScanned() time.Time`

GetDiscoverLastScanned returns the DiscoverLastScanned field if non-nil, zero value otherwise.

### GetDiscoverLastScannedOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverLastScannedOk() (*time.Time, bool)`

GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverLastScanned(v time.Time)`

SetDiscoverLastScanned sets DiscoverLastScanned field to given value.

### HasDiscoverLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverLastScanned() bool`

HasDiscoverLastScanned returns a boolean if a field has been set.

### SetDiscoverLastScannedNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverLastScannedNil(b bool)`

 SetDiscoverLastScannedNil sets the value for DiscoverLastScanned to be an explicit nil

### UnsetDiscoverLastScanned
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetDiscoverLastScanned()`

UnsetDiscoverLastScanned ensures that no value is present for DiscoverLastScanned, not even an explicit nil
### GetMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorLastScanned() time.Time`

GetMonitorLastScanned returns the MonitorLastScanned field if non-nil, zero value otherwise.

### GetMonitorLastScannedOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorLastScannedOk() (*time.Time, bool)`

GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorLastScanned(v time.Time)`

SetMonitorLastScanned sets MonitorLastScanned field to given value.

### HasMonitorLastScanned

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorLastScanned() bool`

HasMonitorLastScanned returns a boolean if a field has been set.

### SetMonitorLastScannedNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorLastScannedNil(b bool)`

 SetMonitorLastScannedNil sets the value for MonitorLastScanned to be an explicit nil

### UnsetMonitorLastScanned
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetMonitorLastScanned()`

UnsetMonitorLastScanned ensures that no value is present for MonitorLastScanned, not even an explicit nil
### GetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverJobParts() int32`

GetDiscoverJobParts returns the DiscoverJobParts field if non-nil, zero value otherwise.

### GetDiscoverJobPartsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetDiscoverJobPartsOk() (*int32, bool)`

GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetDiscoverJobParts(v int32)`

SetDiscoverJobParts sets DiscoverJobParts field to given value.

### HasDiscoverJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasDiscoverJobParts() bool`

HasDiscoverJobParts returns a boolean if a field has been set.

### GetMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorJobParts() int32`

GetMonitorJobParts returns the MonitorJobParts field if non-nil, zero value otherwise.

### GetMonitorJobPartsOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetMonitorJobPartsOk() (*int32, bool)`

GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetMonitorJobParts(v int32)`

SetMonitorJobParts sets MonitorJobParts field to given value.

### HasMonitorJobParts

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasMonitorJobParts() bool`

HasMonitorJobParts returns a boolean if a field has been set.

### GetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetQuietHours() []KeyfactorWebKeyfactorApiModelsSslQuietHourResponse`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetQuietHoursOk() (*[]KeyfactorWebKeyfactorApiModelsSslQuietHourResponse, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetQuietHours(v []KeyfactorWebKeyfactorApiModelsSslQuietHourResponse)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *KeyfactorWebKeyfactorApiModelsSslNetworkQueryResponse) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


