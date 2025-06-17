# OrchestratorPoolsAgentPoolGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPoolId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DiscoverAgentsCount** | Pointer to **int32** |  | [optional] 
**MonitorAgentsCount** | Pointer to **int32** |  | [optional] 
**Agents** | Pointer to [**[]OrchestratorPoolsAgentPoolAgentGetResponse**](OrchestratorPoolsAgentPoolAgentGetResponse.md) |  | [optional] 

## Methods

### NewOrchestratorPoolsAgentPoolGetResponse

`func NewOrchestratorPoolsAgentPoolGetResponse() *OrchestratorPoolsAgentPoolGetResponse`

NewOrchestratorPoolsAgentPoolGetResponse instantiates a new OrchestratorPoolsAgentPoolGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorPoolsAgentPoolGetResponseWithDefaults

`func NewOrchestratorPoolsAgentPoolGetResponseWithDefaults() *OrchestratorPoolsAgentPoolGetResponse`

NewOrchestratorPoolsAgentPoolGetResponseWithDefaults instantiates a new OrchestratorPoolsAgentPoolGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPoolId

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *OrchestratorPoolsAgentPoolGetResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetName

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrchestratorPoolsAgentPoolGetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrchestratorPoolsAgentPoolGetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDiscoverAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetDiscoverAgentsCount() int32`

GetDiscoverAgentsCount returns the DiscoverAgentsCount field if non-nil, zero value otherwise.

### GetDiscoverAgentsCountOk

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetDiscoverAgentsCountOk() (*int32, bool)`

GetDiscoverAgentsCountOk returns a tuple with the DiscoverAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetDiscoverAgentsCount(v int32)`

SetDiscoverAgentsCount sets DiscoverAgentsCount field to given value.

### HasDiscoverAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) HasDiscoverAgentsCount() bool`

HasDiscoverAgentsCount returns a boolean if a field has been set.

### GetMonitorAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetMonitorAgentsCount() int32`

GetMonitorAgentsCount returns the MonitorAgentsCount field if non-nil, zero value otherwise.

### GetMonitorAgentsCountOk

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetMonitorAgentsCountOk() (*int32, bool)`

GetMonitorAgentsCountOk returns a tuple with the MonitorAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetMonitorAgentsCount(v int32)`

SetMonitorAgentsCount sets MonitorAgentsCount field to given value.

### HasMonitorAgentsCount

`func (o *OrchestratorPoolsAgentPoolGetResponse) HasMonitorAgentsCount() bool`

HasMonitorAgentsCount returns a boolean if a field has been set.

### GetAgents

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetAgents() []OrchestratorPoolsAgentPoolAgentGetResponse`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *OrchestratorPoolsAgentPoolGetResponse) GetAgentsOk() (*[]OrchestratorPoolsAgentPoolAgentGetResponse, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetAgents(v []OrchestratorPoolsAgentPoolAgentGetResponse)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *OrchestratorPoolsAgentPoolGetResponse) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *OrchestratorPoolsAgentPoolGetResponse) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *OrchestratorPoolsAgentPoolGetResponse) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


