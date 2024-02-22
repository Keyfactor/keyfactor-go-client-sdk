# KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPoolId** | **string** |  | 
**Name** | **string** |  | 
**Agents** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentCreationRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentCreationRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest(agentPoolId string, name string, ) *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.


### GetName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetAgents() []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentCreationRequest`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) GetAgentsOk() (*[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentCreationRequest, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) SetAgents(v []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentCreationRequest)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


