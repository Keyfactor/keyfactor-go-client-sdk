# KeyfactorApiModelsOrchestratorJobsBulkJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestratorJobPairs** | Pointer to [**[]ModelsOrchestratorJobsBulkOrchestratorJobPair**](ModelsOrchestratorJobsBulkOrchestratorJobPair.md) |  | [optional] 
**FailedOrchestratorIds** | Pointer to **[]string** |  | [optional] 
**JobTypeName** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]KeyfactorApiModelsOrchestratorJobsJobFieldResponse**](KeyfactorApiModelsOrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKeyfactorApiModelsOrchestratorJobsBulkJobResponse

`func NewKeyfactorApiModelsOrchestratorJobsBulkJobResponse() *KeyfactorApiModelsOrchestratorJobsBulkJobResponse`

NewKeyfactorApiModelsOrchestratorJobsBulkJobResponse instantiates a new KeyfactorApiModelsOrchestratorJobsBulkJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults

`func NewKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults() *KeyfactorApiModelsOrchestratorJobsBulkJobResponse`

NewKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsBulkJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestratorJobPairs

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetOrchestratorJobPairs() []ModelsOrchestratorJobsBulkOrchestratorJobPair`

GetOrchestratorJobPairs returns the OrchestratorJobPairs field if non-nil, zero value otherwise.

### GetOrchestratorJobPairsOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetOrchestratorJobPairsOk() (*[]ModelsOrchestratorJobsBulkOrchestratorJobPair, bool)`

GetOrchestratorJobPairsOk returns a tuple with the OrchestratorJobPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorJobPairs

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetOrchestratorJobPairs(v []ModelsOrchestratorJobsBulkOrchestratorJobPair)`

SetOrchestratorJobPairs sets OrchestratorJobPairs field to given value.

### HasOrchestratorJobPairs

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasOrchestratorJobPairs() bool`

HasOrchestratorJobPairs returns a boolean if a field has been set.

### GetFailedOrchestratorIds

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetFailedOrchestratorIds() []string`

GetFailedOrchestratorIds returns the FailedOrchestratorIds field if non-nil, zero value otherwise.

### GetFailedOrchestratorIdsOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetFailedOrchestratorIdsOk() (*[]string, bool)`

GetFailedOrchestratorIdsOk returns a tuple with the FailedOrchestratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOrchestratorIds

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetFailedOrchestratorIds(v []string)`

SetFailedOrchestratorIds sets FailedOrchestratorIds field to given value.

### HasFailedOrchestratorIds

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasFailedOrchestratorIds() bool`

HasFailedOrchestratorIds returns a boolean if a field has been set.

### GetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobFields() []KeyfactorApiModelsOrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobFieldsOk() (*[]KeyfactorApiModelsOrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobFields(v []KeyfactorApiModelsOrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


