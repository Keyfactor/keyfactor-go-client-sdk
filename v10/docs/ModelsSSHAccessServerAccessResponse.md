# ModelsSSHAccessServerAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int64** |  | [optional] 
**LogonUsers** | Pointer to [**[]ModelsSSHAccessLogonUserAccessResponse**](ModelsSSHAccessLogonUserAccessResponse.md) |  | [optional] 

## Methods

### NewModelsSSHAccessServerAccessResponse

`func NewModelsSSHAccessServerAccessResponse() *ModelsSSHAccessServerAccessResponse`

NewModelsSSHAccessServerAccessResponse instantiates a new ModelsSSHAccessServerAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessServerAccessResponseWithDefaults

`func NewModelsSSHAccessServerAccessResponseWithDefaults() *ModelsSSHAccessServerAccessResponse`

NewModelsSSHAccessServerAccessResponseWithDefaults instantiates a new ModelsSSHAccessServerAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ModelsSSHAccessServerAccessResponse) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelsSSHAccessServerAccessResponse) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelsSSHAccessServerAccessResponse) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelsSSHAccessServerAccessResponse) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetLogonUsers

`func (o *ModelsSSHAccessServerAccessResponse) GetLogonUsers() []ModelsSSHAccessLogonUserAccessResponse`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *ModelsSSHAccessServerAccessResponse) GetLogonUsersOk() (*[]ModelsSSHAccessLogonUserAccessResponse, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *ModelsSSHAccessServerAccessResponse) SetLogonUsers(v []ModelsSSHAccessLogonUserAccessResponse)`

SetLogonUsers sets LogonUsers field to given value.

### HasLogonUsers

`func (o *ModelsSSHAccessServerAccessResponse) HasLogonUsers() bool`

HasLogonUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


