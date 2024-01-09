# CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerName** | **string** |  | 
**GroupName** | **string** |  | 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest(ownerName string, groupName string, ) *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequestWithDefaults() *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.

### HasUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest) HasUnderManagement() bool`

HasUnderManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


