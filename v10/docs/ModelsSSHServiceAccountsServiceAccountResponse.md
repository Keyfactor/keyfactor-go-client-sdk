# ModelsSSHServiceAccountsServiceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ClientHostname** | Pointer to **string** |  | [optional] 
**ServerGroup** | Pointer to [**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md) |  | [optional] 
**User** | Pointer to [**ModelsSSHUsersSshUserResponse**](ModelsSSHUsersSshUserResponse.md) |  | [optional] 

## Methods

### NewModelsSSHServiceAccountsServiceAccountResponse

`func NewModelsSSHServiceAccountsServiceAccountResponse() *ModelsSSHServiceAccountsServiceAccountResponse`

NewModelsSSHServiceAccountsServiceAccountResponse instantiates a new ModelsSSHServiceAccountsServiceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServiceAccountsServiceAccountResponseWithDefaults

`func NewModelsSSHServiceAccountsServiceAccountResponseWithDefaults() *ModelsSSHServiceAccountsServiceAccountResponse`

NewModelsSSHServiceAccountsServiceAccountResponseWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientHostname

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.

### HasClientHostname

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasClientHostname() bool`

HasClientHostname returns a boolean if a field has been set.

### GetServerGroup

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetServerGroup() ModelsSSHServerGroupsServerGroupResponse`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetServerGroupOk() (*ModelsSSHServerGroupsServerGroupResponse, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetServerGroup(v ModelsSSHServerGroupsServerGroupResponse)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetUser

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetUser() ModelsSSHUsersSshUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetUserOk() (*ModelsSSHUsersSshUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetUser(v ModelsSSHUsersSshUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


