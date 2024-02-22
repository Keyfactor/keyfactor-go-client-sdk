# KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**ScheduleType** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EntityId** | Pointer to **NullableInt64** |  | [optional] 
**LastRun** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse

`func NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse() *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse`

NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse instantiates a new KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse`

NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetScheduleType

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEntityId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetLastRun

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


