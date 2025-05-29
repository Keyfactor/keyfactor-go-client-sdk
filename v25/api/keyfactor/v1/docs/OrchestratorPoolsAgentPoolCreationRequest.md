# OrchestratorPoolsAgentPoolCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Agents** | Pointer to [**[]OrchestratorPoolsAgentPoolAgentCreationRequest**](OrchestratorPoolsAgentPoolAgentCreationRequest.md) |  | [optional] 

## Methods

### NewOrchestratorPoolsAgentPoolCreationRequest

`func NewOrchestratorPoolsAgentPoolCreationRequest(name string, ) *OrchestratorPoolsAgentPoolCreationRequest`

NewOrchestratorPoolsAgentPoolCreationRequest instantiates a new OrchestratorPoolsAgentPoolCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorPoolsAgentPoolCreationRequestWithDefaults

`func NewOrchestratorPoolsAgentPoolCreationRequestWithDefaults() *OrchestratorPoolsAgentPoolCreationRequest`

NewOrchestratorPoolsAgentPoolCreationRequestWithDefaults instantiates a new OrchestratorPoolsAgentPoolCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrchestratorPoolsAgentPoolCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrchestratorPoolsAgentPoolCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrchestratorPoolsAgentPoolCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAgents

`func (o *OrchestratorPoolsAgentPoolCreationRequest) GetAgents() []OrchestratorPoolsAgentPoolAgentCreationRequest`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *OrchestratorPoolsAgentPoolCreationRequest) GetAgentsOk() (*[]OrchestratorPoolsAgentPoolAgentCreationRequest, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *OrchestratorPoolsAgentPoolCreationRequest) SetAgents(v []OrchestratorPoolsAgentPoolAgentCreationRequest)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *OrchestratorPoolsAgentPoolCreationRequest) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *OrchestratorPoolsAgentPoolCreationRequest) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *OrchestratorPoolsAgentPoolCreationRequest) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


