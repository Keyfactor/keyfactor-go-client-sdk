# CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerName** | **string** |  | 
**GroupName** | **string** |  | 
**SyncSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**UnderManagement** | **bool** |  | 

## Methods

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest(id string, ownerName string, groupName string, underManagement bool, ) *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults() *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest`

NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSyncSchedule returns the SyncSchedule field if non-nil, zero value otherwise.

### GetSyncScheduleOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetSyncScheduleOk returns a tuple with the SyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSyncSchedule sets SyncSchedule field to given value.

### HasSyncSchedule

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) HasSyncSchedule() bool`

HasSyncSchedule returns a boolean if a field has been set.

### GetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagement() bool`

GetUnderManagement returns the UnderManagement field if non-nil, zero value otherwise.

### GetUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagementOk() (*bool, bool)`

GetUnderManagementOk returns a tuple with the UnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderManagement

`func (o *CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest) SetUnderManagement(v bool)`

SetUnderManagement sets UnderManagement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


