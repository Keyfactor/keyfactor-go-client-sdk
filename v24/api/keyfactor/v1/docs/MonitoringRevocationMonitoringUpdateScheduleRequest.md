# MonitoringRevocationMonitoringUpdateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 

## Methods

### NewMonitoringRevocationMonitoringUpdateScheduleRequest

`func NewMonitoringRevocationMonitoringUpdateScheduleRequest(id int32, ) *MonitoringRevocationMonitoringUpdateScheduleRequest`

NewMonitoringRevocationMonitoringUpdateScheduleRequest instantiates a new MonitoringRevocationMonitoringUpdateScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringRevocationMonitoringUpdateScheduleRequestWithDefaults

`func NewMonitoringRevocationMonitoringUpdateScheduleRequestWithDefaults() *MonitoringRevocationMonitoringUpdateScheduleRequest`

NewMonitoringRevocationMonitoringUpdateScheduleRequestWithDefaults instantiates a new MonitoringRevocationMonitoringUpdateScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSchedule

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MonitoringRevocationMonitoringUpdateScheduleRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


