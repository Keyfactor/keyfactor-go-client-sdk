# SchedulingScheduledTaskScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**LastEstimated** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSchedulingScheduledTaskScheduleResponse

`func NewSchedulingScheduledTaskScheduleResponse() *SchedulingScheduledTaskScheduleResponse`

NewSchedulingScheduledTaskScheduleResponse instantiates a new SchedulingScheduledTaskScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingScheduledTaskScheduleResponseWithDefaults

`func NewSchedulingScheduledTaskScheduleResponseWithDefaults() *SchedulingScheduledTaskScheduleResponse`

NewSchedulingScheduledTaskScheduleResponseWithDefaults instantiates a new SchedulingScheduledTaskScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *SchedulingScheduledTaskScheduleResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SchedulingScheduledTaskScheduleResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SchedulingScheduledTaskScheduleResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *SchedulingScheduledTaskScheduleResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetLastEstimated

`func (o *SchedulingScheduledTaskScheduleResponse) GetLastEstimated() time.Time`

GetLastEstimated returns the LastEstimated field if non-nil, zero value otherwise.

### GetLastEstimatedOk

`func (o *SchedulingScheduledTaskScheduleResponse) GetLastEstimatedOk() (*time.Time, bool)`

GetLastEstimatedOk returns a tuple with the LastEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEstimated

`func (o *SchedulingScheduledTaskScheduleResponse) SetLastEstimated(v time.Time)`

SetLastEstimated sets LastEstimated field to given value.

### HasLastEstimated

`func (o *SchedulingScheduledTaskScheduleResponse) HasLastEstimated() bool`

HasLastEstimated returns a boolean if a field has been set.

### SetLastEstimatedNil

`func (o *SchedulingScheduledTaskScheduleResponse) SetLastEstimatedNil(b bool)`

 SetLastEstimatedNil sets the value for LastEstimated to be an explicit nil

### UnsetLastEstimated
`func (o *SchedulingScheduledTaskScheduleResponse) UnsetLastEstimated()`

UnsetLastEstimated ensures that no value is present for LastEstimated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


