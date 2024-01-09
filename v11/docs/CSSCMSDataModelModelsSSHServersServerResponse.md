# CSSCMSDataModelModelsSSHServersServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md) |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**Orchestrator** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHServersServerResponse

`func NewCSSCMSDataModelModelsSSHServersServerResponse() *CSSCMSDataModelModelsSSHServersServerResponse`

NewCSSCMSDataModelModelsSSHServersServerResponse instantiates a new CSSCMSDataModelModelsSSHServersServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServersServerResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHServersServerResponseWithDefaults() *CSSCMSDataModelModelsSSHServersServerResponse`

NewCSSCMSDataModelModelsSSHServersServerResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHServersServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAgentId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetHostname

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### GetOwner

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetOwner() CSSCMSDataModelModelsSSHUsersSshUserResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetOwnerOk() (*CSSCMSDataModelModelsSSHUsersSshUserResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetOwner(v CSSCMSDataModelModelsSSHUsersSshUserResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroupName

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetOrchestrator

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetOrchestrator() string`

GetOrchestrator returns the Orchestrator field if non-nil, zero value otherwise.

### GetOrchestratorOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetOrchestratorOk() (*string, bool)`

GetOrchestratorOk returns a tuple with the Orchestrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrator

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetOrchestrator(v string)`

SetOrchestrator sets Orchestrator field to given value.

### HasOrchestrator

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasOrchestrator() bool`

HasOrchestrator returns a boolean if a field has been set.

### SetOrchestratorNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetOrchestratorNil(b bool)`

 SetOrchestratorNil sets the value for Orchestrator to be an explicit nil

### UnsetOrchestrator
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetOrchestrator()`

UnsetOrchestrator ensures that no value is present for Orchestrator, not even an explicit nil
### GetPort

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CSSCMSDataModelModelsSSHServersServerResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CSSCMSDataModelModelsSSHServersServerResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


