# OrchestratorJobsJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**OrchestratorId** | Pointer to **string** |  | [optional] 
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]OrchestratorJobsJobFieldResponse**](OrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrchestratorJobsJobResponse

`func NewOrchestratorJobsJobResponse() *OrchestratorJobsJobResponse`

NewOrchestratorJobsJobResponse instantiates a new OrchestratorJobsJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorJobsJobResponseWithDefaults

`func NewOrchestratorJobsJobResponseWithDefaults() *OrchestratorJobsJobResponse`

NewOrchestratorJobsJobResponseWithDefaults instantiates a new OrchestratorJobsJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *OrchestratorJobsJobResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *OrchestratorJobsJobResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *OrchestratorJobsJobResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *OrchestratorJobsJobResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOrchestratorId

`func (o *OrchestratorJobsJobResponse) GetOrchestratorId() string`

GetOrchestratorId returns the OrchestratorId field if non-nil, zero value otherwise.

### GetOrchestratorIdOk

`func (o *OrchestratorJobsJobResponse) GetOrchestratorIdOk() (*string, bool)`

GetOrchestratorIdOk returns a tuple with the OrchestratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorId

`func (o *OrchestratorJobsJobResponse) SetOrchestratorId(v string)`

SetOrchestratorId sets OrchestratorId field to given value.

### HasOrchestratorId

`func (o *OrchestratorJobsJobResponse) HasOrchestratorId() bool`

HasOrchestratorId returns a boolean if a field has been set.

### GetJobTypeName

`func (o *OrchestratorJobsJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *OrchestratorJobsJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *OrchestratorJobsJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *OrchestratorJobsJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *OrchestratorJobsJobResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *OrchestratorJobsJobResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetSchedule

`func (o *OrchestratorJobsJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *OrchestratorJobsJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *OrchestratorJobsJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *OrchestratorJobsJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *OrchestratorJobsJobResponse) GetJobFields() []OrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *OrchestratorJobsJobResponse) GetJobFieldsOk() (*[]OrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *OrchestratorJobsJobResponse) SetJobFields(v []OrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *OrchestratorJobsJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *OrchestratorJobsJobResponse) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *OrchestratorJobsJobResponse) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil
### GetRequestTimestamp

`func (o *OrchestratorJobsJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *OrchestratorJobsJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *OrchestratorJobsJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *OrchestratorJobsJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


