# CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyGenerationRequest** | [**CSSCMSDataModelModelsSSHKeysKeyGenerationRequest**](CSSCMSDataModelModelsSSHKeysKeyGenerationRequest.md) |  | 
**User** | [**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest.md) |  | 
**ClientHostname** | **string** |  | 
**ServerGroupId** | **string** |  | 

## Methods

### NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest

`func NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest(keyGenerationRequest CSSCMSDataModelModelsSSHKeysKeyGenerationRequest, user CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest, clientHostname string, serverGroupId string, ) *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest`

NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest instantiates a new CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults() *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest`

NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyGenerationRequest

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetKeyGenerationRequest() CSSCMSDataModelModelsSSHKeysKeyGenerationRequest`

GetKeyGenerationRequest returns the KeyGenerationRequest field if non-nil, zero value otherwise.

### GetKeyGenerationRequestOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetKeyGenerationRequestOk() (*CSSCMSDataModelModelsSSHKeysKeyGenerationRequest, bool)`

GetKeyGenerationRequestOk returns a tuple with the KeyGenerationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationRequest

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) SetKeyGenerationRequest(v CSSCMSDataModelModelsSSHKeysKeyGenerationRequest)`

SetKeyGenerationRequest sets KeyGenerationRequest field to given value.


### GetUser

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetUser() CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetUserOk() (*CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) SetUser(v CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest)`

SetUser sets User field to given value.


### GetClientHostname

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.


### GetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


