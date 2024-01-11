# KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 
**OrchestratorJobPairs** | Pointer to [**[]ModelsOrchestratorJobsBulkOrchestratorJobPair**](ModelsOrchestratorJobsBulkOrchestratorJobPair.md) |  | [optional] 
**FailedOrchestratorIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse() *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobFields() []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetJobFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobFields(v []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil
### GetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetOrchestratorJobPairs

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetOrchestratorJobPairs() []ModelsOrchestratorJobsBulkOrchestratorJobPair`

GetOrchestratorJobPairs returns the OrchestratorJobPairs field if non-nil, zero value otherwise.

### GetOrchestratorJobPairsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetOrchestratorJobPairsOk() (*[]ModelsOrchestratorJobsBulkOrchestratorJobPair, bool)`

GetOrchestratorJobPairsOk returns a tuple with the OrchestratorJobPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorJobPairs

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetOrchestratorJobPairs(v []ModelsOrchestratorJobsBulkOrchestratorJobPair)`

SetOrchestratorJobPairs sets OrchestratorJobPairs field to given value.

### HasOrchestratorJobPairs

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasOrchestratorJobPairs() bool`

HasOrchestratorJobPairs returns a boolean if a field has been set.

### SetOrchestratorJobPairsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetOrchestratorJobPairsNil(b bool)`

 SetOrchestratorJobPairsNil sets the value for OrchestratorJobPairs to be an explicit nil

### UnsetOrchestratorJobPairs
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) UnsetOrchestratorJobPairs()`

UnsetOrchestratorJobPairs ensures that no value is present for OrchestratorJobPairs, not even an explicit nil
### GetFailedOrchestratorIds

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetFailedOrchestratorIds() []string`

GetFailedOrchestratorIds returns the FailedOrchestratorIds field if non-nil, zero value otherwise.

### GetFailedOrchestratorIdsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) GetFailedOrchestratorIdsOk() (*[]string, bool)`

GetFailedOrchestratorIdsOk returns a tuple with the FailedOrchestratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOrchestratorIds

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetFailedOrchestratorIds(v []string)`

SetFailedOrchestratorIds sets FailedOrchestratorIds field to given value.

### HasFailedOrchestratorIds

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) HasFailedOrchestratorIds() bool`

HasFailedOrchestratorIds returns a boolean if a field has been set.

### SetFailedOrchestratorIdsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) SetFailedOrchestratorIdsNil(b bool)`

 SetFailedOrchestratorIdsNil sets the value for FailedOrchestratorIds to be an explicit nil

### UnsetFailedOrchestratorIds
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse) UnsetFailedOrchestratorIds()`

UnsetFailedOrchestratorIds ensures that no value is present for FailedOrchestratorIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


