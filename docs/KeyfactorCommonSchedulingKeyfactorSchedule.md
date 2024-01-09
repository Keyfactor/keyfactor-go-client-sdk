# KeyfactorCommonSchedulingKeyfactorSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immediate** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to [**KeyfactorCommonSchedulingModelsIntervalModel**](KeyfactorCommonSchedulingModelsIntervalModel.md) |  | [optional] 
**Daily** | Pointer to [**KeyfactorCommonSchedulingModelsTimeModel**](KeyfactorCommonSchedulingModelsTimeModel.md) |  | [optional] 
**Weekly** | Pointer to [**KeyfactorCommonSchedulingModelsWeeklyModel**](KeyfactorCommonSchedulingModelsWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**KeyfactorCommonSchedulingModelsMonthlyModel**](KeyfactorCommonSchedulingModelsMonthlyModel.md) |  | [optional] 
**ExactlyOnce** | Pointer to [**KeyfactorCommonSchedulingModelsTimeModel**](KeyfactorCommonSchedulingModelsTimeModel.md) |  | [optional] 

## Methods

### NewKeyfactorCommonSchedulingKeyfactorSchedule

`func NewKeyfactorCommonSchedulingKeyfactorSchedule() *KeyfactorCommonSchedulingKeyfactorSchedule`

NewKeyfactorCommonSchedulingKeyfactorSchedule instantiates a new KeyfactorCommonSchedulingKeyfactorSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorCommonSchedulingKeyfactorScheduleWithDefaults

`func NewKeyfactorCommonSchedulingKeyfactorScheduleWithDefaults() *KeyfactorCommonSchedulingKeyfactorSchedule`

NewKeyfactorCommonSchedulingKeyfactorScheduleWithDefaults instantiates a new KeyfactorCommonSchedulingKeyfactorSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmediate

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetInterval

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetInterval() KeyfactorCommonSchedulingModelsIntervalModel`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetIntervalOk() (*KeyfactorCommonSchedulingModelsIntervalModel, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetInterval(v KeyfactorCommonSchedulingModelsIntervalModel)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDaily

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetDaily() KeyfactorCommonSchedulingModelsTimeModel`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetDailyOk() (*KeyfactorCommonSchedulingModelsTimeModel, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetDaily(v KeyfactorCommonSchedulingModelsTimeModel)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetWeekly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetWeekly() KeyfactorCommonSchedulingModelsWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetWeeklyOk() (*KeyfactorCommonSchedulingModelsWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetWeekly(v KeyfactorCommonSchedulingModelsWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetMonthly() KeyfactorCommonSchedulingModelsMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetMonthlyOk() (*KeyfactorCommonSchedulingModelsMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetMonthly(v KeyfactorCommonSchedulingModelsMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetExactlyOnce

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetExactlyOnce() KeyfactorCommonSchedulingModelsTimeModel`

GetExactlyOnce returns the ExactlyOnce field if non-nil, zero value otherwise.

### GetExactlyOnceOk

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) GetExactlyOnceOk() (*KeyfactorCommonSchedulingModelsTimeModel, bool)`

GetExactlyOnceOk returns a tuple with the ExactlyOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactlyOnce

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) SetExactlyOnce(v KeyfactorCommonSchedulingModelsTimeModel)`

SetExactlyOnce sets ExactlyOnce field to given value.

### HasExactlyOnce

`func (o *KeyfactorCommonSchedulingKeyfactorSchedule) HasExactlyOnce() bool`

HasExactlyOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


