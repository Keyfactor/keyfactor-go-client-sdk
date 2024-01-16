# ModelsSSHServersServerCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**Hostname** | **string** |  | 
**ServerGroupId** | **string** |  | 
**UnderManagement** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelsSSHServersServerCreationRequest

`func NewModelsSSHServersServerCreationRequest(agentId string, hostname string, serverGroupId string, ) *ModelsSSHServersServerCreationRequest`

NewModelsSSHServersServerCreationRequest instantiates a new ModelsSSHServersServerCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServersServerCreationRequestWithDefaults

`func NewModelsSSHServersServerCreationRequestWithDefaults() *ModelsSSHServersServerCreationRequest`

NewModelsSSHServersServerCreationRequestWithDefaults instantiates a new ModelsSSHServersServerCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ModelsSSHServersServerCreationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsSSHServersServerCreationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsSSHServersServerCreationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetHostname

`func (o *ModelsSSHServersServerCreationRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ModelsSSHServersServerCreationRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ModelsSSHServersServerCreationRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetServerGroupId

`func (o *ModelsSSHServersServerCreationRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ModelsSSHServersServerCreationRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ModelsSSHServersServerCreationRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.


### GetUnderManagement

`func (o *ModelsSSHServersServerCreationRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *ModelsSSHServersServerCreationRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *ModelsSSHServersServerCreationRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *ModelsSSHServersServerCreationRequest) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### GetPort

`func (o *ModelsSSHServersServerCreationRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelsSSHServersServerCreationRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelsSSHServersServerCreationRequest) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelsSSHServersServerCreationRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


