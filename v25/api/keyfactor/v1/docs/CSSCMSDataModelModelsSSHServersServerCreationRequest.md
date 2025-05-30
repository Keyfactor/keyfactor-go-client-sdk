# CSSCMSDataModelModelsSSHServersServerCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**Hostname** | **string** |  | 
**ServerGroupId** | **string** |  | 
**UnderManagement** | Pointer to **NullableBool** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHServersServerCreationRequest

`func NewCSSCMSDataModelModelsSSHServersServerCreationRequest(agentId string, hostname string, serverGroupId string, ) *CSSCMSDataModelModelsSSHServersServerCreationRequest`

NewCSSCMSDataModelModelsSSHServersServerCreationRequest instantiates a new CSSCMSDataModelModelsSSHServersServerCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServersServerCreationRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHServersServerCreationRequestWithDefaults() *CSSCMSDataModelModelsSSHServersServerCreationRequest`

NewCSSCMSDataModelModelsSSHServersServerCreationRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHServersServerCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetHostname

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.


### GetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### SetUnderManagementNil

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetUnderManagementNil(b bool)`

 SetUnderManagementNil sets the value for UnderManagement to be an explicit nil

### UnsetUnderManagement
`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) UnsetUnderManagement()`

UnsetUnderManagement ensures that no value is present for UnderManagement, not even an explicit nil
### GetPort

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CSSCMSDataModelModelsSSHServersServerCreationRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


