# CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerGroupId** | **string** |  | 
**LogonUsers** | [**[]CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest**](CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest.md) |  | 

## Methods

### NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequest

`func NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequest(serverGroupId string, logonUsers []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest, ) *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest`

NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequest instantiates a new CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequestWithDefaults() *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest`

NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerGroupId

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.


### GetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) GetLogonUsers() []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest`

GetLogonUsers returns the LogonUsers field if non-nil, zero value otherwise.

### GetLogonUsersOk

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) GetLogonUsersOk() (*[]CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest, bool)`

GetLogonUsersOk returns a tuple with the LogonUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonUsers

`func (o *CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest) SetLogonUsers(v []CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest)`

SetLogonUsers sets LogonUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


