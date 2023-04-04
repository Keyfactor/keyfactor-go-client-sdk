# ModelsSSHAccessServerAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int32** |  | 
**LogonUsers** | [**[]ModelsSSHAccessLogonUserAccessRequest**](ModelsSSHAccessLogonUserAccessRequest.md) |  | 

## Methods

### NewModelsSSHAccessServerAccessRequest

`func NewModelsSSHAccessServerAccessRequest(serverId int32, logonUsers []ModelsSSHAccessLogonUserAccessRequest, ) *ModelsSSHAccessServerAccessRequest`

NewModelsSSHAccessServerAccessRequest instantiates a new ModelsSSHAccessServerAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessServerAccessRequestWithDefaults

`func NewModelsSSHAccessServerAccessRequestWithDefaults() *ModelsSSHAccessServerAccessRequest`

NewModelsSSHAccessServerAccessRequestWithDefaults instantiates a new ModelsSSHAccessServerAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ModelsSSHAccessServerAccessRequest) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelsSSHAccessServerAccessRequest) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelsSSHAccessServerAccessRequest) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetLogonUsers

`func (o *ModelsSSHAccessServerAccessRequest) GetLogonUsers() []ModelsSSHAccessLogonUserAccessRequest`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *ModelsSSHAccessServerAccessRequest) GetLogonUsersOk() (*[]ModelsSSHAccessLogonUserAccessRequest, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *ModelsSSHAccessServerAccessRequest) SetLogonUsers(v []ModelsSSHAccessLogonUserAccessRequest)`

SetLogonUsers sets LogonUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


