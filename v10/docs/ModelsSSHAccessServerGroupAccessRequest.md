# ModelsSSHAccessServerGroupAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerGroupId** | **string** |  | 
**LogonUsers** | [**[]ModelsSSHAccessLogonUserAccessRequest**](ModelsSSHAccessLogonUserAccessRequest.md) |  | 

## Methods

### NewModelsSSHAccessServerGroupAccessRequest

`func NewModelsSSHAccessServerGroupAccessRequest(serverGroupId string, logonUsers []ModelsSSHAccessLogonUserAccessRequest, ) *ModelsSSHAccessServerGroupAccessRequest`

NewModelsSSHAccessServerGroupAccessRequest instantiates a new ModelsSSHAccessServerGroupAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessServerGroupAccessRequestWithDefaults

`func NewModelsSSHAccessServerGroupAccessRequestWithDefaults() *ModelsSSHAccessServerGroupAccessRequest`

NewModelsSSHAccessServerGroupAccessRequestWithDefaults instantiates a new ModelsSSHAccessServerGroupAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerGroupId

`func (o *ModelsSSHAccessServerGroupAccessRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ModelsSSHAccessServerGroupAccessRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ModelsSSHAccessServerGroupAccessRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.


### GetLogonUsers

`func (o *ModelsSSHAccessServerGroupAccessRequest) GetLogonUsers() []ModelsSSHAccessLogonUserAccessRequest`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *ModelsSSHAccessServerGroupAccessRequest) GetLogonUsersOk() (*[]ModelsSSHAccessLogonUserAccessRequest, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *ModelsSSHAccessServerGroupAccessRequest) SetLogonUsers(v []ModelsSSHAccessLogonUserAccessRequest)`

SetLogonUsers sets LogonUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


