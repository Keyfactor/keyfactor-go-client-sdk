# KeyfactorCommonSchedulingModelsWeeklyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Days** | Pointer to [**[]SystemDayOfWeek**](SystemDayOfWeek.md) |  | [optional] 

## Methods

### NewKeyfactorCommonSchedulingModelsWeeklyModel

`func NewKeyfactorCommonSchedulingModelsWeeklyModel() *KeyfactorCommonSchedulingModelsWeeklyModel`

NewKeyfactorCommonSchedulingModelsWeeklyModel instantiates a new KeyfactorCommonSchedulingModelsWeeklyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorCommonSchedulingModelsWeeklyModelWithDefaults

`func NewKeyfactorCommonSchedulingModelsWeeklyModelWithDefaults() *KeyfactorCommonSchedulingModelsWeeklyModel`

NewKeyfactorCommonSchedulingModelsWeeklyModelWithDefaults instantiates a new KeyfactorCommonSchedulingModelsWeeklyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDays

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) GetDays() []SystemDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) GetDaysOk() (*[]SystemDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) SetDays(v []SystemDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *KeyfactorCommonSchedulingModelsWeeklyModel) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


