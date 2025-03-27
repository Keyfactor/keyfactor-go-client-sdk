# SchedulingScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**ScheduleType** | Pointer to [**CSSCMSCoreEnumsScheduledTaskType**](CSSCMSCoreEnumsScheduledTaskType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EntityId** | Pointer to **NullableInt64** |  | [optional] 
**LastRun** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSchedulingScheduledTaskResponse

`func NewSchedulingScheduledTaskResponse() *SchedulingScheduledTaskResponse`

NewSchedulingScheduledTaskResponse instantiates a new SchedulingScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingScheduledTaskResponseWithDefaults

`func NewSchedulingScheduledTaskResponseWithDefaults() *SchedulingScheduledTaskResponse`

NewSchedulingScheduledTaskResponseWithDefaults instantiates a new SchedulingScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchedulingScheduledTaskResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchedulingScheduledTaskResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchedulingScheduledTaskResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SchedulingScheduledTaskResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchedule

`func (o *SchedulingScheduledTaskResponse) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SchedulingScheduledTaskResponse) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SchedulingScheduledTaskResponse) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *SchedulingScheduledTaskResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *SchedulingScheduledTaskResponse) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *SchedulingScheduledTaskResponse) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetScheduleType

`func (o *SchedulingScheduledTaskResponse) GetScheduleType() CSSCMSCoreEnumsScheduledTaskType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *SchedulingScheduledTaskResponse) GetScheduleTypeOk() (*CSSCMSCoreEnumsScheduledTaskType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *SchedulingScheduledTaskResponse) SetScheduleType(v CSSCMSCoreEnumsScheduledTaskType)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *SchedulingScheduledTaskResponse) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetEnabled

`func (o *SchedulingScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchedulingScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchedulingScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SchedulingScheduledTaskResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *SchedulingScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulingScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulingScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulingScheduledTaskResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SchedulingScheduledTaskResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SchedulingScheduledTaskResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEntityId

`func (o *SchedulingScheduledTaskResponse) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SchedulingScheduledTaskResponse) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SchedulingScheduledTaskResponse) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SchedulingScheduledTaskResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *SchedulingScheduledTaskResponse) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *SchedulingScheduledTaskResponse) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetLastRun

`func (o *SchedulingScheduledTaskResponse) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *SchedulingScheduledTaskResponse) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *SchedulingScheduledTaskResponse) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *SchedulingScheduledTaskResponse) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *SchedulingScheduledTaskResponse) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *SchedulingScheduledTaskResponse) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


