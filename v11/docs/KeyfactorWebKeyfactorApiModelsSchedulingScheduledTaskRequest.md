# KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScheduleType** | **int32** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to **NullableInt32** |  | [optional] 
**TimeOfDay** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest

`func NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest(scheduleType int32, ) *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest`

NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest instantiates a new KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest`

NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleType

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetTimeOfDay

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetTimeOfDay() time.Time`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) GetTimeOfDayOk() (*time.Time, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetTimeOfDay(v time.Time)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### SetTimeOfDayNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) SetTimeOfDayNil(b bool)`

 SetTimeOfDayNil sets the value for TimeOfDay to be an explicit nil

### UnsetTimeOfDay
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest) UnsetTimeOfDay()`

UnsetTimeOfDay ensures that no value is present for TimeOfDay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


