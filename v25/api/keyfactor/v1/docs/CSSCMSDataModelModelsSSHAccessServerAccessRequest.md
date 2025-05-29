# CSSCMSDataModelModelsSSHAccessServerAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int32** |  | 
**LogonUsers** | [**[]CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest**](CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest.md) |  | 

## Methods

### NewCSSCMSDataModelModelsSSHAccessServerAccessRequest

`func NewCSSCMSDataModelModelsSSHAccessServerAccessRequest(serverId int32, logonUsers []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest, ) *CSSCMSDataModelModelsSSHAccessServerAccessRequest`

NewCSSCMSDataModelModelsSSHAccessServerAccessRequest instantiates a new CSSCMSDataModelModelsSSHAccessServerAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHAccessServerAccessRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHAccessServerAccessRequestWithDefaults() *CSSCMSDataModelModelsSSHAccessServerAccessRequest`

NewCSSCMSDataModelModelsSSHAccessServerAccessRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHAccessServerAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) GetLogonUsers() []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) GetLogonUsersOk() (*[]CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerAccessRequest) SetLogonUsers(v []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest)`

SetLogonUsers sets LogonUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


