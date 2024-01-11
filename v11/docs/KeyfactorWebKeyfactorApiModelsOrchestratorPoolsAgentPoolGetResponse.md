# KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPoolId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DiscoverAgentsCount** | Pointer to **int32** |  | [optional] 
**MonitorAgentsCount** | Pointer to **int32** |  | [optional] 
**Agents** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse() *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDiscoverAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetDiscoverAgentsCount() int32`

GetDiscoverAgentsCount returns the DiscoverAgentsCount field if non-nil, zero value otherwise.

### GetDiscoverAgentsCountOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetDiscoverAgentsCountOk() (*int32, bool)`

GetDiscoverAgentsCountOk returns a tuple with the DiscoverAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetDiscoverAgentsCount(v int32)`

SetDiscoverAgentsCount sets DiscoverAgentsCount field to given value.

### HasDiscoverAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) HasDiscoverAgentsCount() bool`

HasDiscoverAgentsCount returns a boolean if a field has been set.

### GetMonitorAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetMonitorAgentsCount() int32`

GetMonitorAgentsCount returns the MonitorAgentsCount field if non-nil, zero value otherwise.

### GetMonitorAgentsCountOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetMonitorAgentsCountOk() (*int32, bool)`

GetMonitorAgentsCountOk returns a tuple with the MonitorAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetMonitorAgentsCount(v int32)`

SetMonitorAgentsCount sets MonitorAgentsCount field to given value.

### HasMonitorAgentsCount

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) HasMonitorAgentsCount() bool`

HasMonitorAgentsCount returns a boolean if a field has been set.

### GetAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetAgents() []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) GetAgentsOk() (*[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetAgents(v []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


