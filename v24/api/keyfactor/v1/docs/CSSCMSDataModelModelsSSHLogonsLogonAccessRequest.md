# CSSCMSDataModelModelsSSHLogonsLogonAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogonId** | **int32** |  | 
**UserIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequest

`func NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequest(logonId int32, ) *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest`

NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequest instantiates a new CSSCMSDataModelModelsSSHLogonsLogonAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequestWithDefaults() *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest`

NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHLogonsLogonAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogonId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) GetLogonId() int32`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) GetLogonIdOk() (*int32, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) SetLogonId(v int32)`

SetLogonId sets LogonId field to given value.


### GetUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *CSSCMSDataModelModelsSSHLogonsLogonAccessRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


