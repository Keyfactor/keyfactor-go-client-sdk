# OrchestratorJobsBulkJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestratorJobPairs** | Pointer to [**[]CSSCMSDataModelModelsOrchestratorJobsBulkOrchestratorJobPair**](CSSCMSDataModelModelsOrchestratorJobsBulkOrchestratorJobPair.md) |  | [optional] 
**FailedOrchestratorIds** | Pointer to **[]string** |  | [optional] 
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to [**[]OrchestratorJobsJobFieldResponse**](OrchestratorJobsJobFieldResponse.md) |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrchestratorJobsBulkJobResponse

`func NewOrchestratorJobsBulkJobResponse() *OrchestratorJobsBulkJobResponse`

NewOrchestratorJobsBulkJobResponse instantiates a new OrchestratorJobsBulkJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorJobsBulkJobResponseWithDefaults

`func NewOrchestratorJobsBulkJobResponseWithDefaults() *OrchestratorJobsBulkJobResponse`

NewOrchestratorJobsBulkJobResponseWithDefaults instantiates a new OrchestratorJobsBulkJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestratorJobPairs

`func (o *OrchestratorJobsBulkJobResponse) GetOrchestratorJobPairs() []CSSCMSDataModelModelsOrchestratorJobsBulkOrchestratorJobPair`

GetOrchestratorJobPairs returns the OrchestratorJobPairs field if non-nil, zero value otherwise.

### GetOrchestratorJobPairsOk

`func (o *OrchestratorJobsBulkJobResponse) GetOrchestratorJobPairsOk() (*[]CSSCMSDataModelModelsOrchestratorJobsBulkOrchestratorJobPair, bool)`

GetOrchestratorJobPairsOk returns a tuple with the OrchestratorJobPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorJobPairs

`func (o *OrchestratorJobsBulkJobResponse) SetOrchestratorJobPairs(v []CSSCMSDataModelModelsOrchestratorJobsBulkOrchestratorJobPair)`

SetOrchestratorJobPairs sets OrchestratorJobPairs field to given value.

### HasOrchestratorJobPairs

`func (o *OrchestratorJobsBulkJobResponse) HasOrchestratorJobPairs() bool`

HasOrchestratorJobPairs returns a boolean if a field has been set.

### SetOrchestratorJobPairsNil

`func (o *OrchestratorJobsBulkJobResponse) SetOrchestratorJobPairsNil(b bool)`

 SetOrchestratorJobPairsNil sets the value for OrchestratorJobPairs to be an explicit nil

### UnsetOrchestratorJobPairs
`func (o *OrchestratorJobsBulkJobResponse) UnsetOrchestratorJobPairs()`

UnsetOrchestratorJobPairs ensures that no value is present for OrchestratorJobPairs, not even an explicit nil
### GetFailedOrchestratorIds

`func (o *OrchestratorJobsBulkJobResponse) GetFailedOrchestratorIds() []string`

GetFailedOrchestratorIds returns the FailedOrchestratorIds field if non-nil, zero value otherwise.

### GetFailedOrchestratorIdsOk

`func (o *OrchestratorJobsBulkJobResponse) GetFailedOrchestratorIdsOk() (*[]string, bool)`

GetFailedOrchestratorIdsOk returns a tuple with the FailedOrchestratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOrchestratorIds

`func (o *OrchestratorJobsBulkJobResponse) SetFailedOrchestratorIds(v []string)`

SetFailedOrchestratorIds sets FailedOrchestratorIds field to given value.

### HasFailedOrchestratorIds

`func (o *OrchestratorJobsBulkJobResponse) HasFailedOrchestratorIds() bool`

HasFailedOrchestratorIds returns a boolean if a field has been set.

### SetFailedOrchestratorIdsNil

`func (o *OrchestratorJobsBulkJobResponse) SetFailedOrchestratorIdsNil(b bool)`

 SetFailedOrchestratorIdsNil sets the value for FailedOrchestratorIds to be an explicit nil

### UnsetFailedOrchestratorIds
`func (o *OrchestratorJobsBulkJobResponse) UnsetFailedOrchestratorIds()`

UnsetFailedOrchestratorIds ensures that no value is present for FailedOrchestratorIds, not even an explicit nil
### GetJobTypeName

`func (o *OrchestratorJobsBulkJobResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *OrchestratorJobsBulkJobResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *OrchestratorJobsBulkJobResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *OrchestratorJobsBulkJobResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *OrchestratorJobsBulkJobResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *OrchestratorJobsBulkJobResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetSchedule

`func (o *OrchestratorJobsBulkJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *OrchestratorJobsBulkJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *OrchestratorJobsBulkJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *OrchestratorJobsBulkJobResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *OrchestratorJobsBulkJobResponse) GetJobFields() []OrchestratorJobsJobFieldResponse`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *OrchestratorJobsBulkJobResponse) GetJobFieldsOk() (*[]OrchestratorJobsJobFieldResponse, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *OrchestratorJobsBulkJobResponse) SetJobFields(v []OrchestratorJobsJobFieldResponse)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *OrchestratorJobsBulkJobResponse) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *OrchestratorJobsBulkJobResponse) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *OrchestratorJobsBulkJobResponse) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil
### GetRequestTimestamp

`func (o *OrchestratorJobsBulkJobResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *OrchestratorJobsBulkJobResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *OrchestratorJobsBulkJobResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *OrchestratorJobsBulkJobResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


