# ModelsSSHLogonsLogonCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**ServerId** | **int32** |  | 
**UserIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewModelsSSHLogonsLogonCreationRequest

`func NewModelsSSHLogonsLogonCreationRequest(username string, serverId int32, ) *ModelsSSHLogonsLogonCreationRequest`

NewModelsSSHLogonsLogonCreationRequest instantiates a new ModelsSSHLogonsLogonCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHLogonsLogonCreationRequestWithDefaults

`func NewModelsSSHLogonsLogonCreationRequestWithDefaults() *ModelsSSHLogonsLogonCreationRequest`

NewModelsSSHLogonsLogonCreationRequestWithDefaults instantiates a new ModelsSSHLogonsLogonCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ModelsSSHLogonsLogonCreationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHLogonsLogonCreationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHLogonsLogonCreationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetServerId

`func (o *ModelsSSHLogonsLogonCreationRequest) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelsSSHLogonsLogonCreationRequest) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelsSSHLogonsLogonCreationRequest) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetUserIds

`func (o *ModelsSSHLogonsLogonCreationRequest) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ModelsSSHLogonsLogonCreationRequest) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ModelsSSHLogonsLogonCreationRequest) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ModelsSSHLogonsLogonCreationRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *ModelsSSHLogonsLogonCreationRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *ModelsSSHLogonsLogonCreationRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


