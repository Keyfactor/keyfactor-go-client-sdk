# ModelsSSHServiceAccountsServiceAccountCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyGenerationRequest** | [**ModelsSSHKeysKeyGenerationRequest**](ModelsSSHKeysKeyGenerationRequest.md) |  | 
**User** | [**ModelsSSHServiceAccountsServiceAccountUserCreationRequest**](ModelsSSHServiceAccountsServiceAccountUserCreationRequest.md) |  | 
**ClientHostname** | **string** |  | 
**ServerGroupId** | **string** |  | 

## Methods

### NewModelsSSHServiceAccountsServiceAccountCreationRequest

`func NewModelsSSHServiceAccountsServiceAccountCreationRequest(keyGenerationRequest ModelsSSHKeysKeyGenerationRequest, user ModelsSSHServiceAccountsServiceAccountUserCreationRequest, clientHostname string, serverGroupId string, ) *ModelsSSHServiceAccountsServiceAccountCreationRequest`

NewModelsSSHServiceAccountsServiceAccountCreationRequest instantiates a new ModelsSSHServiceAccountsServiceAccountCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults

`func NewModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults() *ModelsSSHServiceAccountsServiceAccountCreationRequest`

NewModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyGenerationRequest

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetKeyGenerationRequest() ModelsSSHKeysKeyGenerationRequest`

GetKeyGenerationRequest returns the KeyGenerationRequest field if non-nil, zero value otherwise.

### GetKeyGenerationRequestOk

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetKeyGenerationRequestOk() (*ModelsSSHKeysKeyGenerationRequest, bool)`

GetKeyGenerationRequestOk returns a tuple with the KeyGenerationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationRequest

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) SetKeyGenerationRequest(v ModelsSSHKeysKeyGenerationRequest)`

SetKeyGenerationRequest sets KeyGenerationRequest field to given value.


### GetUser

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetUser() ModelsSSHServiceAccountsServiceAccountUserCreationRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetUserOk() (*ModelsSSHServiceAccountsServiceAccountUserCreationRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) SetUser(v ModelsSSHServiceAccountsServiceAccountUserCreationRequest)`

SetUser sets User field to given value.


### GetClientHostname

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.


### GetServerGroupId

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ModelsSSHServiceAccountsServiceAccountCreationRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


