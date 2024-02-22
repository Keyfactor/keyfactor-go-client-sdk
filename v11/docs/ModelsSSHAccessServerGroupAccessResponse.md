# ModelsSSHAccessServerGroupAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerGroupId** | Pointer to **string** |  | [optional] 
**LogonUsers** | Pointer to [**[]ModelsSSHAccessLogonUserAccessResponse**](ModelsSSHAccessLogonUserAccessResponse.md) |  | [optional] 

## Methods

### NewModelsSSHAccessServerGroupAccessResponse

`func NewModelsSSHAccessServerGroupAccessResponse() *ModelsSSHAccessServerGroupAccessResponse`

NewModelsSSHAccessServerGroupAccessResponse instantiates a new ModelsSSHAccessServerGroupAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessServerGroupAccessResponseWithDefaults

`func NewModelsSSHAccessServerGroupAccessResponseWithDefaults() *ModelsSSHAccessServerGroupAccessResponse`

NewModelsSSHAccessServerGroupAccessResponseWithDefaults instantiates a new ModelsSSHAccessServerGroupAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerGroupId

`func (o *ModelsSSHAccessServerGroupAccessResponse) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ModelsSSHAccessServerGroupAccessResponse) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ModelsSSHAccessServerGroupAccessResponse) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ModelsSSHAccessServerGroupAccessResponse) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### GetLogonUsers

`func (o *ModelsSSHAccessServerGroupAccessResponse) GetLogonUsers() []ModelsSSHAccessLogonUserAccessResponse`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *ModelsSSHAccessServerGroupAccessResponse) GetLogonUsersOk() (*[]ModelsSSHAccessLogonUserAccessResponse, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *ModelsSSHAccessServerGroupAccessResponse) SetLogonUsers(v []ModelsSSHAccessLogonUserAccessResponse)`

SetLogonUsers sets LogonUsers field to given value.

### HasLogonUsers

`func (o *ModelsSSHAccessServerGroupAccessResponse) HasLogonUsers() bool`

HasLogonUsers returns a boolean if a field has been set.

### SetLogonUsersNil

`func (o *ModelsSSHAccessServerGroupAccessResponse) SetLogonUsersNil(b bool)`

 SetLogonUsersNil sets the value for LogonUsers to be an explicit nil

### UnsetLogonUsers
`func (o *ModelsSSHAccessServerGroupAccessResponse) UnsetLogonUsers()`

UnsetLogonUsers ensures that no value is present for LogonUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


