# SslNetworkQueryResponse

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
**QuietHours** | Pointer to [**[]SslQuietHourResponse**](SslQuietHourResponse.md) |  | [optional] 
**BlackoutStart** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**BlackoutEnd** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 

## Methods

### NewSslNetworkQueryResponse

`func NewSslNetworkQueryResponse() *SslNetworkQueryResponse`

NewSslNetworkQueryResponse instantiates a new SslNetworkQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslNetworkQueryResponseWithDefaults

`func NewSslNetworkQueryResponseWithDefaults() *SslNetworkQueryResponse`

NewSslNetworkQueryResponseWithDefaults instantiates a new SslNetworkQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *SslNetworkQueryResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *SslNetworkQueryResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *SslNetworkQueryResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *SslNetworkQueryResponse) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *SslNetworkQueryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslNetworkQueryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslNetworkQueryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SslNetworkQueryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SslNetworkQueryResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SslNetworkQueryResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAgentPoolName

`func (o *SslNetworkQueryResponse) GetAgentPoolName() string`

GetAgentPoolName returns the AgentPoolName field if non-nil, zero value otherwise.

### GetAgentPoolNameOk

`func (o *SslNetworkQueryResponse) GetAgentPoolNameOk() (*string, bool)`

GetAgentPoolNameOk returns a tuple with the AgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolName

`func (o *SslNetworkQueryResponse) SetAgentPoolName(v string)`

SetAgentPoolName sets AgentPoolName field to given value.

### HasAgentPoolName

`func (o *SslNetworkQueryResponse) HasAgentPoolName() bool`

HasAgentPoolName returns a boolean if a field has been set.

### SetAgentPoolNameNil

`func (o *SslNetworkQueryResponse) SetAgentPoolNameNil(b bool)`

 SetAgentPoolNameNil sets the value for AgentPoolName to be an explicit nil

### UnsetAgentPoolName
`func (o *SslNetworkQueryResponse) UnsetAgentPoolName()`

UnsetAgentPoolName ensures that no value is present for AgentPoolName, not even an explicit nil
### GetAgentPoolId

`func (o *SslNetworkQueryResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *SslNetworkQueryResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *SslNetworkQueryResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *SslNetworkQueryResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### SetAgentPoolIdNil

`func (o *SslNetworkQueryResponse) SetAgentPoolIdNil(b bool)`

 SetAgentPoolIdNil sets the value for AgentPoolId to be an explicit nil

### UnsetAgentPoolId
`func (o *SslNetworkQueryResponse) UnsetAgentPoolId()`

UnsetAgentPoolId ensures that no value is present for AgentPoolId, not even an explicit nil
### GetDescription

`func (o *SslNetworkQueryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SslNetworkQueryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SslNetworkQueryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SslNetworkQueryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SslNetworkQueryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SslNetworkQueryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *SslNetworkQueryResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SslNetworkQueryResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SslNetworkQueryResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SslNetworkQueryResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoverSchedule

`func (o *SslNetworkQueryResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetDiscoverSchedule returns the DiscoverSchedule field if non-nil, zero value otherwise.

### GetDiscoverScheduleOk

`func (o *SslNetworkQueryResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverSchedule

`func (o *SslNetworkQueryResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetDiscoverSchedule sets DiscoverSchedule field to given value.

### HasDiscoverSchedule

`func (o *SslNetworkQueryResponse) HasDiscoverSchedule() bool`

HasDiscoverSchedule returns a boolean if a field has been set.

### GetMonitorSchedule

`func (o *SslNetworkQueryResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetMonitorSchedule returns the MonitorSchedule field if non-nil, zero value otherwise.

### GetMonitorScheduleOk

`func (o *SslNetworkQueryResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetMonitorScheduleOk returns a tuple with the MonitorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSchedule

`func (o *SslNetworkQueryResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetMonitorSchedule sets MonitorSchedule field to given value.

### HasMonitorSchedule

`func (o *SslNetworkQueryResponse) HasMonitorSchedule() bool`

HasMonitorSchedule returns a boolean if a field has been set.

### GetDiscoverPercentComplete

`func (o *SslNetworkQueryResponse) GetDiscoverPercentComplete() float64`

GetDiscoverPercentComplete returns the DiscoverPercentComplete field if non-nil, zero value otherwise.

### GetDiscoverPercentCompleteOk

`func (o *SslNetworkQueryResponse) GetDiscoverPercentCompleteOk() (*float64, bool)`

GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPercentComplete

`func (o *SslNetworkQueryResponse) SetDiscoverPercentComplete(v float64)`

SetDiscoverPercentComplete sets DiscoverPercentComplete field to given value.

### HasDiscoverPercentComplete

`func (o *SslNetworkQueryResponse) HasDiscoverPercentComplete() bool`

HasDiscoverPercentComplete returns a boolean if a field has been set.

### GetMonitorPercentComplete

`func (o *SslNetworkQueryResponse) GetMonitorPercentComplete() float64`

GetMonitorPercentComplete returns the MonitorPercentComplete field if non-nil, zero value otherwise.

### GetMonitorPercentCompleteOk

`func (o *SslNetworkQueryResponse) GetMonitorPercentCompleteOk() (*float64, bool)`

GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPercentComplete

`func (o *SslNetworkQueryResponse) SetMonitorPercentComplete(v float64)`

SetMonitorPercentComplete sets MonitorPercentComplete field to given value.

### HasMonitorPercentComplete

`func (o *SslNetworkQueryResponse) HasMonitorPercentComplete() bool`

HasMonitorPercentComplete returns a boolean if a field has been set.

### GetDiscoverStatus

`func (o *SslNetworkQueryResponse) GetDiscoverStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetDiscoverStatus returns the DiscoverStatus field if non-nil, zero value otherwise.

### GetDiscoverStatusOk

`func (o *SslNetworkQueryResponse) GetDiscoverStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetDiscoverStatusOk returns a tuple with the DiscoverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverStatus

`func (o *SslNetworkQueryResponse) SetDiscoverStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetDiscoverStatus sets DiscoverStatus field to given value.

### HasDiscoverStatus

`func (o *SslNetworkQueryResponse) HasDiscoverStatus() bool`

HasDiscoverStatus returns a boolean if a field has been set.

### GetMonitorStatus

`func (o *SslNetworkQueryResponse) GetMonitorStatus() CSSCMSCoreEnumsSslNetworkJobStatus`

GetMonitorStatus returns the MonitorStatus field if non-nil, zero value otherwise.

### GetMonitorStatusOk

`func (o *SslNetworkQueryResponse) GetMonitorStatusOk() (*CSSCMSCoreEnumsSslNetworkJobStatus, bool)`

GetMonitorStatusOk returns a tuple with the MonitorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorStatus

`func (o *SslNetworkQueryResponse) SetMonitorStatus(v CSSCMSCoreEnumsSslNetworkJobStatus)`

SetMonitorStatus sets MonitorStatus field to given value.

### HasMonitorStatus

`func (o *SslNetworkQueryResponse) HasMonitorStatus() bool`

HasMonitorStatus returns a boolean if a field has been set.

### GetDiscoverLastScanned

`func (o *SslNetworkQueryResponse) GetDiscoverLastScanned() time.Time`

GetDiscoverLastScanned returns the DiscoverLastScanned field if non-nil, zero value otherwise.

### GetDiscoverLastScannedOk

`func (o *SslNetworkQueryResponse) GetDiscoverLastScannedOk() (*time.Time, bool)`

GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLastScanned

`func (o *SslNetworkQueryResponse) SetDiscoverLastScanned(v time.Time)`

SetDiscoverLastScanned sets DiscoverLastScanned field to given value.

### HasDiscoverLastScanned

`func (o *SslNetworkQueryResponse) HasDiscoverLastScanned() bool`

HasDiscoverLastScanned returns a boolean if a field has been set.

### SetDiscoverLastScannedNil

`func (o *SslNetworkQueryResponse) SetDiscoverLastScannedNil(b bool)`

 SetDiscoverLastScannedNil sets the value for DiscoverLastScanned to be an explicit nil

### UnsetDiscoverLastScanned
`func (o *SslNetworkQueryResponse) UnsetDiscoverLastScanned()`

UnsetDiscoverLastScanned ensures that no value is present for DiscoverLastScanned, not even an explicit nil
### GetMonitorLastScanned

`func (o *SslNetworkQueryResponse) GetMonitorLastScanned() time.Time`

GetMonitorLastScanned returns the MonitorLastScanned field if non-nil, zero value otherwise.

### GetMonitorLastScannedOk

`func (o *SslNetworkQueryResponse) GetMonitorLastScannedOk() (*time.Time, bool)`

GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorLastScanned

`func (o *SslNetworkQueryResponse) SetMonitorLastScanned(v time.Time)`

SetMonitorLastScanned sets MonitorLastScanned field to given value.

### HasMonitorLastScanned

`func (o *SslNetworkQueryResponse) HasMonitorLastScanned() bool`

HasMonitorLastScanned returns a boolean if a field has been set.

### SetMonitorLastScannedNil

`func (o *SslNetworkQueryResponse) SetMonitorLastScannedNil(b bool)`

 SetMonitorLastScannedNil sets the value for MonitorLastScanned to be an explicit nil

### UnsetMonitorLastScanned
`func (o *SslNetworkQueryResponse) UnsetMonitorLastScanned()`

UnsetMonitorLastScanned ensures that no value is present for MonitorLastScanned, not even an explicit nil
### GetSslAlertRecipients

`func (o *SslNetworkQueryResponse) GetSslAlertRecipients() []string`

GetSslAlertRecipients returns the SslAlertRecipients field if non-nil, zero value otherwise.

### GetSslAlertRecipientsOk

`func (o *SslNetworkQueryResponse) GetSslAlertRecipientsOk() (*[]string, bool)`

GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAlertRecipients

`func (o *SslNetworkQueryResponse) SetSslAlertRecipients(v []string)`

SetSslAlertRecipients sets SslAlertRecipients field to given value.

### HasSslAlertRecipients

`func (o *SslNetworkQueryResponse) HasSslAlertRecipients() bool`

HasSslAlertRecipients returns a boolean if a field has been set.

### SetSslAlertRecipientsNil

`func (o *SslNetworkQueryResponse) SetSslAlertRecipientsNil(b bool)`

 SetSslAlertRecipientsNil sets the value for SslAlertRecipients to be an explicit nil

### UnsetSslAlertRecipients
`func (o *SslNetworkQueryResponse) UnsetSslAlertRecipients()`

UnsetSslAlertRecipients ensures that no value is present for SslAlertRecipients, not even an explicit nil
### GetGetRobots

`func (o *SslNetworkQueryResponse) GetGetRobots() bool`

GetGetRobots returns the GetRobots field if non-nil, zero value otherwise.

### GetGetRobotsOk

`func (o *SslNetworkQueryResponse) GetGetRobotsOk() (*bool, bool)`

GetGetRobotsOk returns a tuple with the GetRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRobots

`func (o *SslNetworkQueryResponse) SetGetRobots(v bool)`

SetGetRobots sets GetRobots field to given value.

### HasGetRobots

`func (o *SslNetworkQueryResponse) HasGetRobots() bool`

HasGetRobots returns a boolean if a field has been set.

### GetDiscoverTimeoutMs

`func (o *SslNetworkQueryResponse) GetDiscoverTimeoutMs() float64`

GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field if non-nil, zero value otherwise.

### GetDiscoverTimeoutMsOk

`func (o *SslNetworkQueryResponse) GetDiscoverTimeoutMsOk() (*float64, bool)`

GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTimeoutMs

`func (o *SslNetworkQueryResponse) SetDiscoverTimeoutMs(v float64)`

SetDiscoverTimeoutMs sets DiscoverTimeoutMs field to given value.

### HasDiscoverTimeoutMs

`func (o *SslNetworkQueryResponse) HasDiscoverTimeoutMs() bool`

HasDiscoverTimeoutMs returns a boolean if a field has been set.

### GetMonitorTimeoutMs

`func (o *SslNetworkQueryResponse) GetMonitorTimeoutMs() float64`

GetMonitorTimeoutMs returns the MonitorTimeoutMs field if non-nil, zero value otherwise.

### GetMonitorTimeoutMsOk

`func (o *SslNetworkQueryResponse) GetMonitorTimeoutMsOk() (*float64, bool)`

GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeoutMs

`func (o *SslNetworkQueryResponse) SetMonitorTimeoutMs(v float64)`

SetMonitorTimeoutMs sets MonitorTimeoutMs field to given value.

### HasMonitorTimeoutMs

`func (o *SslNetworkQueryResponse) HasMonitorTimeoutMs() bool`

HasMonitorTimeoutMs returns a boolean if a field has been set.

### GetExpirationAlertDays

`func (o *SslNetworkQueryResponse) GetExpirationAlertDays() float64`

GetExpirationAlertDays returns the ExpirationAlertDays field if non-nil, zero value otherwise.

### GetExpirationAlertDaysOk

`func (o *SslNetworkQueryResponse) GetExpirationAlertDaysOk() (*float64, bool)`

GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAlertDays

`func (o *SslNetworkQueryResponse) SetExpirationAlertDays(v float64)`

SetExpirationAlertDays sets ExpirationAlertDays field to given value.

### HasExpirationAlertDays

`func (o *SslNetworkQueryResponse) HasExpirationAlertDays() bool`

HasExpirationAlertDays returns a boolean if a field has been set.

### GetDiscoverJobParts

`func (o *SslNetworkQueryResponse) GetDiscoverJobParts() int32`

GetDiscoverJobParts returns the DiscoverJobParts field if non-nil, zero value otherwise.

### GetDiscoverJobPartsOk

`func (o *SslNetworkQueryResponse) GetDiscoverJobPartsOk() (*int32, bool)`

GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverJobParts

`func (o *SslNetworkQueryResponse) SetDiscoverJobParts(v int32)`

SetDiscoverJobParts sets DiscoverJobParts field to given value.

### HasDiscoverJobParts

`func (o *SslNetworkQueryResponse) HasDiscoverJobParts() bool`

HasDiscoverJobParts returns a boolean if a field has been set.

### GetMonitorJobParts

`func (o *SslNetworkQueryResponse) GetMonitorJobParts() int32`

GetMonitorJobParts returns the MonitorJobParts field if non-nil, zero value otherwise.

### GetMonitorJobPartsOk

`func (o *SslNetworkQueryResponse) GetMonitorJobPartsOk() (*int32, bool)`

GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobParts

`func (o *SslNetworkQueryResponse) SetMonitorJobParts(v int32)`

SetMonitorJobParts sets MonitorJobParts field to given value.

### HasMonitorJobParts

`func (o *SslNetworkQueryResponse) HasMonitorJobParts() bool`

HasMonitorJobParts returns a boolean if a field has been set.

### GetQuietHours

`func (o *SslNetworkQueryResponse) GetQuietHours() []SslQuietHourResponse`

GetQuietHours returns the QuietHours field if non-nil, zero value otherwise.

### GetQuietHoursOk

`func (o *SslNetworkQueryResponse) GetQuietHoursOk() (*[]SslQuietHourResponse, bool)`

GetQuietHoursOk returns a tuple with the QuietHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHours

`func (o *SslNetworkQueryResponse) SetQuietHours(v []SslQuietHourResponse)`

SetQuietHours sets QuietHours field to given value.

### HasQuietHours

`func (o *SslNetworkQueryResponse) HasQuietHours() bool`

HasQuietHours returns a boolean if a field has been set.

### SetQuietHoursNil

`func (o *SslNetworkQueryResponse) SetQuietHoursNil(b bool)`

 SetQuietHoursNil sets the value for QuietHours to be an explicit nil

### UnsetQuietHours
`func (o *SslNetworkQueryResponse) UnsetQuietHours()`

UnsetQuietHours ensures that no value is present for QuietHours, not even an explicit nil
### GetBlackoutStart

`func (o *SslNetworkQueryResponse) GetBlackoutStart() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutStart returns the BlackoutStart field if non-nil, zero value otherwise.

### GetBlackoutStartOk

`func (o *SslNetworkQueryResponse) GetBlackoutStartOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutStartOk returns a tuple with the BlackoutStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStart

`func (o *SslNetworkQueryResponse) SetBlackoutStart(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutStart sets BlackoutStart field to given value.

### HasBlackoutStart

`func (o *SslNetworkQueryResponse) HasBlackoutStart() bool`

HasBlackoutStart returns a boolean if a field has been set.

### GetBlackoutEnd

`func (o *SslNetworkQueryResponse) GetBlackoutEnd() KeyfactorCommonSchedulingModelsWeeklyModel`

GetBlackoutEnd returns the BlackoutEnd field if non-nil, zero value otherwise.

### GetBlackoutEndOk

`func (o *SslNetworkQueryResponse) GetBlackoutEndOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetBlackoutEndOk returns a tuple with the BlackoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEnd

`func (o *SslNetworkQueryResponse) SetBlackoutEnd(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetBlackoutEnd sets BlackoutEnd field to given value.

### HasBlackoutEnd

`func (o *SslNetworkQueryResponse) HasBlackoutEnd() bool`

HasBlackoutEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


