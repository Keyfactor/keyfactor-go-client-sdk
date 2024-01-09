# CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md) |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 
**ServerCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponse

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponse() *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponse instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponseWithDefaults() *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetOwner() CSSCMSDataModelModelsSSHUsersSshUserResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetOwnerOk() (*CSSCMSDataModelModelsSSHUsersSshUserResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetOwner(v CSSCMSDataModelModelsSSHUsersSshUserResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### GetServerCount

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetServerCount() int32`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) GetServerCountOk() (*int32, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetServerCount(v int32)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### SetServerCountNil

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) SetServerCountNil(b bool)`

 SetServerCountNil sets the value for ServerCount to be an explicit nil

### UnsetServerCount
`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse) UnsetServerCount()`

UnsetServerCount ensures that no value is present for ServerCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


