# ModelsSSHServerGroupsServerGroupCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerName** | **string** |  | 
**GroupName** | **string** |  | 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsSSHServerGroupsServerGroupCreationRequest

`func NewModelsSSHServerGroupsServerGroupCreationRequest(ownerName string, groupName string, ) *ModelsSSHServerGroupsServerGroupCreationRequest`

NewModelsSSHServerGroupsServerGroupCreationRequest instantiates a new ModelsSSHServerGroupsServerGroupCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServerGroupsServerGroupCreationRequestWithDefaults

`func NewModelsSSHServerGroupsServerGroupCreationRequestWithDefaults() *ModelsSSHServerGroupsServerGroupCreationRequest`

NewModelsSSHServerGroupsServerGroupCreationRequestWithDefaults instantiates a new ModelsSSHServerGroupsServerGroupCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerName

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetGroupName

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupCreationRequest) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


