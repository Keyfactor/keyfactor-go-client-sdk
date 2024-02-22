# ModelsSSHServerGroupsServerGroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerName** | **string** |  | 
**GroupName** | **string** |  | 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | **bool** |  | 

## Methods

### NewModelsSSHServerGroupsServerGroupUpdateRequest

`func NewModelsSSHServerGroupsServerGroupUpdateRequest(id string, ownerName string, groupName string, underManagement bool, ) *ModelsSSHServerGroupsServerGroupUpdateRequest`

NewModelsSSHServerGroupsServerGroupUpdateRequest instantiates a new ModelsSSHServerGroupsServerGroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults

`func NewModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults() *ModelsSSHServerGroupsServerGroupUpdateRequest`

NewModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults instantiates a new ModelsSSHServerGroupsServerGroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerName

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetGroupName

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


