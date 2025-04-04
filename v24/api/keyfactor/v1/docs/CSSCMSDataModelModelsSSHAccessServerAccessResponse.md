# CSSCMSDataModelModelsSSHAccessServerAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** |  | [optional] 
**LogonUsers** | Pointer to [**[]CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse**](CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHAccessServerAccessResponse

`func NewCSSCMSDataModelModelsSSHAccessServerAccessResponse() *CSSCMSDataModelModelsSSHAccessServerAccessResponse`

NewCSSCMSDataModelModelsSSHAccessServerAccessResponse instantiates a new CSSCMSDataModelModelsSSHAccessServerAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHAccessServerAccessResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHAccessServerAccessResponseWithDefaults() *CSSCMSDataModelModelsSSHAccessServerAccessResponse`

NewCSSCMSDataModelModelsSSHAccessServerAccessResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHAccessServerAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) GetLogonUsers() []CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) GetLogonUsersOk() (*[]CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) SetLogonUsers(v []CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse)`

SetLogonUsers sets LogonUsers field to given value.

### HasLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) HasLogonUsers() bool`

HasLogonUsers returns a boolean if a field has been set.

### SetLogonUsersNil

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) SetLogonUsersNil(b bool)`

 SetLogonUsersNil sets the value for LogonUsers to be an explicit nil

### UnsetLogonUsers
`func (o *CSSCMSDataModelModelsSSHAccessServerAccessResponse) UnsetLogonUsers()`

UnsetLogonUsers ensures that no value is present for LogonUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


