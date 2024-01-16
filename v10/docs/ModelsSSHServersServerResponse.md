# ModelsSSHServersServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**ServerGroupId** | Pointer to **string** |  | [optional] 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**ModelsSSHUsersSshUserResponse**](ModelsSSHUsersSshUserResponse.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Orchestrator** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelsSSHServersServerResponse

`func NewModelsSSHServersServerResponse() *ModelsSSHServersServerResponse`

NewModelsSSHServersServerResponse instantiates a new ModelsSSHServersServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServersServerResponseWithDefaults

`func NewModelsSSHServersServerResponseWithDefaults() *ModelsSSHServersServerResponse`

NewModelsSSHServersServerResponseWithDefaults instantiates a new ModelsSSHServersServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHServersServerResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHServersServerResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHServersServerResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHServersServerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgentId

`func (o *ModelsSSHServersServerResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsSSHServersServerResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsSSHServersServerResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsSSHServersServerResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetHostname

`func (o *ModelsSSHServersServerResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ModelsSSHServersServerResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ModelsSSHServersServerResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ModelsSSHServersServerResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerGroupId

`func (o *ModelsSSHServersServerResponse) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ModelsSSHServersServerResponse) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ModelsSSHServersServerResponse) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ModelsSSHServersServerResponse) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### GetSyncSchedule

`func (o *ModelsSSHServersServerResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *ModelsSSHServersServerResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *ModelsSSHServersServerResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *ModelsSSHServersServerResponse) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *ModelsSSHServersServerResponse) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *ModelsSSHServersServerResponse) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *ModelsSSHServersServerResponse) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *ModelsSSHServersServerResponse) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### GetOwner

`func (o *ModelsSSHServersServerResponse) GetOwner() ModelsSSHUsersSshUserResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsSSHServersServerResponse) GetOwnerOk() (*ModelsSSHUsersSshUserResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsSSHServersServerResponse) SetOwner(v ModelsSSHUsersSshUserResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelsSSHServersServerResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroupName

`func (o *ModelsSSHServersServerResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelsSSHServersServerResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelsSSHServersServerResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ModelsSSHServersServerResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetOrchestrator

`func (o *ModelsSSHServersServerResponse) GetOrchestrator() string`

GetOrchestrator returns the Orchestrator field if non-nil, zero value otherwise.

### GetOrchestratorOk

`func (o *ModelsSSHServersServerResponse) GetOrchestratorOk() (*string, bool)`

GetOrchestratorOk returns a tuple with the Orchestrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrator

`func (o *ModelsSSHServersServerResponse) SetOrchestrator(v string)`

SetOrchestrator sets Orchestrator field to given value.

### HasOrchestrator

`func (o *ModelsSSHServersServerResponse) HasOrchestrator() bool`

HasOrchestrator returns a boolean if a field has been set.

### GetPort

`func (o *ModelsSSHServersServerResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelsSSHServersServerResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelsSSHServersServerResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelsSSHServersServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


