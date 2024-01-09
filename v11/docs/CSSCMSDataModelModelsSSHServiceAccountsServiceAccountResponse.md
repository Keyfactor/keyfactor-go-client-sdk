# CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ClientHostname** | Pointer to **NullableString** |  | [optional] 
**ServerGroup** | Pointer to [**CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md) |  | [optional] 
**User** | Pointer to [**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse

`func NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse() *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse`

NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse instantiates a new CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponseWithDefaults() *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse`

NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientHostname

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.

### HasClientHostname

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) HasClientHostname() bool`

HasClientHostname returns a boolean if a field has been set.

### SetClientHostnameNil

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) SetClientHostnameNil(b bool)`

 SetClientHostnameNil sets the value for ClientHostname to be an explicit nil

### UnsetClientHostname
`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) UnsetClientHostname()`

UnsetClientHostname ensures that no value is present for ClientHostname, not even an explicit nil
### GetServerGroup

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetServerGroup() CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetServerGroupOk() (*CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) SetServerGroup(v CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetUser

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetUser() CSSCMSDataModelModelsSSHUsersSshUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) GetUserOk() (*CSSCMSDataModelModelsSSHUsersSshUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) SetUser(v CSSCMSDataModelModelsSSHUsersSshUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


