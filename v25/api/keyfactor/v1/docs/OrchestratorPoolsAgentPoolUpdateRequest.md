# OrchestratorPoolsAgentPoolUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPoolId** | **string** |  | 
**Name** | **string** |  | 
**Agents** | Pointer to [**[]OrchestratorPoolsAgentPoolAgentCreationRequest**](OrchestratorPoolsAgentPoolAgentCreationRequest.md) |  | [optional] 

## Methods

### NewOrchestratorPoolsAgentPoolUpdateRequest

`func NewOrchestratorPoolsAgentPoolUpdateRequest(agentPoolId string, name string, ) *OrchestratorPoolsAgentPoolUpdateRequest`

NewOrchestratorPoolsAgentPoolUpdateRequest instantiates a new OrchestratorPoolsAgentPoolUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorPoolsAgentPoolUpdateRequestWithDefaults

`func NewOrchestratorPoolsAgentPoolUpdateRequestWithDefaults() *OrchestratorPoolsAgentPoolUpdateRequest`

NewOrchestratorPoolsAgentPoolUpdateRequestWithDefaults instantiates a new OrchestratorPoolsAgentPoolUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPoolId

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.


### GetName

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgents

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetAgents() []OrchestratorPoolsAgentPoolAgentCreationRequest`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) GetAgentsOk() (*[]OrchestratorPoolsAgentPoolAgentCreationRequest, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) SetAgents(v []OrchestratorPoolsAgentPoolAgentCreationRequest)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *OrchestratorPoolsAgentPoolUpdateRequest) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *OrchestratorPoolsAgentPoolUpdateRequest) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


