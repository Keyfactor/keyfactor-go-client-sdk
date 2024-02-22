# ModelsAgentsAgentPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPoolId** | Pointer to **string** | GUID identifier of the agent pool | [optional] 
**Name** | **string** | Name of the agent pool | 
**DiscoverAgentsCount** | Pointer to **int64** | Number of agents that can perform discovery jobs | [optional] 
**MonitorAgentsCount** | Pointer to **int64** | Number of agents that can perform monitoring jobs | [optional] 
**Agents** | Pointer to [**[]ModelsAgentsAgentPoolAgent**](ModelsAgentsAgentPoolAgent.md) | List of the agents assigned to the pool | [optional] 

## Methods

### NewModelsAgentsAgentPool

`func NewModelsAgentsAgentPool(name string, ) *ModelsAgentsAgentPool`

NewModelsAgentsAgentPool instantiates a new ModelsAgentsAgentPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAgentsAgentPoolWithDefaults

`func NewModelsAgentsAgentPoolWithDefaults() *ModelsAgentsAgentPool`

NewModelsAgentsAgentPoolWithDefaults instantiates a new ModelsAgentsAgentPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPoolId

`func (o *ModelsAgentsAgentPool) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *ModelsAgentsAgentPool) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *ModelsAgentsAgentPool) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *ModelsAgentsAgentPool) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetName

`func (o *ModelsAgentsAgentPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsAgentsAgentPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsAgentsAgentPool) SetName(v string)`

SetName sets Name field to given value.


### GetDiscoverAgentsCount

`func (o *ModelsAgentsAgentPool) GetDiscoverAgentsCount() int64`

GetDiscoverAgentsCount returns the DiscoverAgentsCount field if non-nil, zero value otherwise.

### GetDiscoverAgentsCountOk

`func (o *ModelsAgentsAgentPool) GetDiscoverAgentsCountOk() (*int64, bool)`

GetDiscoverAgentsCountOk returns a tuple with the DiscoverAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAgentsCount

`func (o *ModelsAgentsAgentPool) SetDiscoverAgentsCount(v int64)`

SetDiscoverAgentsCount sets DiscoverAgentsCount field to given value.

### HasDiscoverAgentsCount

`func (o *ModelsAgentsAgentPool) HasDiscoverAgentsCount() bool`

HasDiscoverAgentsCount returns a boolean if a field has been set.

### GetMonitorAgentsCount

`func (o *ModelsAgentsAgentPool) GetMonitorAgentsCount() int64`

GetMonitorAgentsCount returns the MonitorAgentsCount field if non-nil, zero value otherwise.

### GetMonitorAgentsCountOk

`func (o *ModelsAgentsAgentPool) GetMonitorAgentsCountOk() (*int64, bool)`

GetMonitorAgentsCountOk returns a tuple with the MonitorAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAgentsCount

`func (o *ModelsAgentsAgentPool) SetMonitorAgentsCount(v int64)`

SetMonitorAgentsCount sets MonitorAgentsCount field to given value.

### HasMonitorAgentsCount

`func (o *ModelsAgentsAgentPool) HasMonitorAgentsCount() bool`

HasMonitorAgentsCount returns a boolean if a field has been set.

### GetAgents

`func (o *ModelsAgentsAgentPool) GetAgents() []ModelsAgentsAgentPoolAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *ModelsAgentsAgentPool) GetAgentsOk() (*[]ModelsAgentsAgentPoolAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *ModelsAgentsAgentPool) SetAgents(v []ModelsAgentsAgentPoolAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *ModelsAgentsAgentPool) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


