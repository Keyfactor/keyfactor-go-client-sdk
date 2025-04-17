# OrchestratorPoolsAgentPoolAgentGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**EnableDiscover** | Pointer to **bool** |  | [optional] 
**EnableMonitor** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**AllowsDiscover** | Pointer to **bool** |  | [optional] 
**AllowsMonitor** | Pointer to **bool** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrchestratorPoolsAgentPoolAgentGetResponse

`func NewOrchestratorPoolsAgentPoolAgentGetResponse() *OrchestratorPoolsAgentPoolAgentGetResponse`

NewOrchestratorPoolsAgentPoolAgentGetResponse instantiates a new OrchestratorPoolsAgentPoolAgentGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorPoolsAgentPoolAgentGetResponseWithDefaults

`func NewOrchestratorPoolsAgentPoolAgentGetResponseWithDefaults() *OrchestratorPoolsAgentPoolAgentGetResponse`

NewOrchestratorPoolsAgentPoolAgentGetResponseWithDefaults instantiates a new OrchestratorPoolsAgentPoolAgentGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetEnableDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetEnableDiscover() bool`

GetEnableDiscover returns the EnableDiscover field if non-nil, zero value otherwise.

### GetEnableDiscoverOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetEnableDiscoverOk() (*bool, bool)`

GetEnableDiscoverOk returns a tuple with the EnableDiscover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetEnableDiscover(v bool)`

SetEnableDiscover sets EnableDiscover field to given value.

### HasEnableDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasEnableDiscover() bool`

HasEnableDiscover returns a boolean if a field has been set.

### GetEnableMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetEnableMonitor() bool`

GetEnableMonitor returns the EnableMonitor field if non-nil, zero value otherwise.

### GetEnableMonitorOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetEnableMonitorOk() (*bool, bool)`

GetEnableMonitorOk returns a tuple with the EnableMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetEnableMonitor(v bool)`

SetEnableMonitor sets EnableMonitor field to given value.

### HasEnableMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasEnableMonitor() bool`

HasEnableMonitor returns a boolean if a field has been set.

### GetVersion

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAllowsDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAllowsDiscover() bool`

GetAllowsDiscover returns the AllowsDiscover field if non-nil, zero value otherwise.

### GetAllowsDiscoverOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAllowsDiscoverOk() (*bool, bool)`

GetAllowsDiscoverOk returns a tuple with the AllowsDiscover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetAllowsDiscover(v bool)`

SetAllowsDiscover sets AllowsDiscover field to given value.

### HasAllowsDiscover

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasAllowsDiscover() bool`

HasAllowsDiscover returns a boolean if a field has been set.

### GetAllowsMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAllowsMonitor() bool`

GetAllowsMonitor returns the AllowsMonitor field if non-nil, zero value otherwise.

### GetAllowsMonitorOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetAllowsMonitorOk() (*bool, bool)`

GetAllowsMonitorOk returns a tuple with the AllowsMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetAllowsMonitor(v bool)`

SetAllowsMonitor sets AllowsMonitor field to given value.

### HasAllowsMonitor

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasAllowsMonitor() bool`

HasAllowsMonitor returns a boolean if a field has been set.

### GetClientMachine

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *OrchestratorPoolsAgentPoolAgentGetResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


