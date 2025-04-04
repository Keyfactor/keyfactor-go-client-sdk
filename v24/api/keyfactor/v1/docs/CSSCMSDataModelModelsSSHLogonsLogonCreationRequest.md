# CSSCMSDataModelModelsSSHLogonsLogonCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**ServerId** | **int32** |  | 
**UserIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequest

`func NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequest(username string, serverId int32, ) *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest`

NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequest instantiates a new CSSCMSDataModelModelsSSHLogonsLogonCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequestWithDefaults() *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest`

NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHLogonsLogonCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetServerId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *CSSCMSDataModelModelsSSHLogonsLogonCreationRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


