# KeyfactorApiModelsOrchestratorJobsJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**OrchestratorId** | Pointer to **string** |  | [optional] 
**JobTypeName** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]KeyfactorApiModelsOrchestratorJobsJobFieldResponse**](KeyfactorApiModelsOrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKeyfactorApiModelsOrchestratorJobsJobResponse

`func NewKeyfactorApiModelsOrchestratorJobsJobResponse() *KeyfactorApiModelsOrchestratorJobsJobResponse`

NewKeyfactorApiModelsOrchestratorJobsJobResponse instantiates a new KeyfactorApiModelsOrchestratorJobsJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults

`func NewKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults() *KeyfactorApiModelsOrchestratorJobsJobResponse`

NewKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOrchestratorId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorId() string`

GetOrchestratorId returns the OrchestratorId field if non-nil, zero value otherwise.

### GetOrchestratorIdOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorIdOk() (*string, bool)`

GetOrchestratorIdOk returns a tuple with the OrchestratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetOrchestratorId(v string)`

SetOrchestratorId sets OrchestratorId field to given value.

### HasOrchestratorId

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasOrchestratorId() bool`

HasOrchestratorId returns a boolean if a field has been set.

### GetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFields() []KeyfactorApiModelsOrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFieldsOk() (*[]KeyfactorApiModelsOrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobFields(v []KeyfactorApiModelsOrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


