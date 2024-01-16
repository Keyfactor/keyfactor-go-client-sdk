# ModelsSSHServiceAccountsServiceAccountUserCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**LogonIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewModelsSSHServiceAccountsServiceAccountUserCreationRequest

`func NewModelsSSHServiceAccountsServiceAccountUserCreationRequest(username string, ) *ModelsSSHServiceAccountsServiceAccountUserCreationRequest`

NewModelsSSHServiceAccountsServiceAccountUserCreationRequest instantiates a new ModelsSSHServiceAccountsServiceAccountUserCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServiceAccountsServiceAccountUserCreationRequestWithDefaults

`func NewModelsSSHServiceAccountsServiceAccountUserCreationRequestWithDefaults() *ModelsSSHServiceAccountsServiceAccountUserCreationRequest`

NewModelsSSHServiceAccountsServiceAccountUserCreationRequestWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountUserCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetLogonIds

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetLogonIds() []int64`

GetLogonIds returns the LogonIds field if non-nil, zero value otherwise.

### GetLogonIdsOk

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetLogonIdsOk() (*[]int64, bool)`

GetLogonIdsOk returns a tuple with the LogonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonIds

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) SetLogonIds(v []int64)`

SetLogonIds sets LogonIds field to given value.

### HasLogonIds

`func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) HasLogonIds() bool`

HasLogonIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


