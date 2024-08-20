# ModelsSSHServerGroupsServerGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ModelsSSHUsersSshUserResponse**](ModelsSSHUsersSshUserResponse.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelsSSHServerGroupsServerGroupResponse

`func NewModelsSSHServerGroupsServerGroupResponse() *ModelsSSHServerGroupsServerGroupResponse`

NewModelsSSHServerGroupsServerGroupResponse instantiates a new ModelsSSHServerGroupsServerGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServerGroupsServerGroupResponseWithDefaults

`func NewModelsSSHServerGroupsServerGroupResponseWithDefaults() *ModelsSSHServerGroupsServerGroupResponse`

NewModelsSSHServerGroupsServerGroupResponseWithDefaults instantiates a new ModelsSSHServerGroupsServerGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetOwner() ModelsSSHUsersSshUserResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetOwnerOk() (*ModelsSSHUsersSshUserResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetOwner(v ModelsSSHUsersSshUserResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroupName

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.

### GetServerCount

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *ModelsSSHServerGroupsServerGroupResponse) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *ModelsSSHServerGroupsServerGroupResponse) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *ModelsSSHServerGroupsServerGroupResponse) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


