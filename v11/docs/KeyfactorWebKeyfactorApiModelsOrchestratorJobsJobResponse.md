# KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**OrchestratorId** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse() *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFields() []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetJobFields(v []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil
### GetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOrchestratorId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorId() string`

GetOrchestratorId returns the OrchestratorId field if non-nil, zero value otherwise.

### GetOrchestratorIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorIdOk() (*string, bool)`

GetOrchestratorIdOk returns a tuple with the OrchestratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) SetOrchestratorId(v string)`

SetOrchestratorId sets OrchestratorId field to given value.

### HasOrchestratorId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse) HasOrchestratorId() bool`

HasOrchestratorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


