# ModelsOrchestratorJobsScheduleBulkJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestratorIds** | **[]string** |  | 
**JobTypeName** | **string** |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsOrchestratorJobsScheduleBulkJobRequest

`func NewModelsOrchestratorJobsScheduleBulkJobRequest(orchestratorIds []string, jobTypeName string, ) *ModelsOrchestratorJobsScheduleBulkJobRequest`

NewModelsOrchestratorJobsScheduleBulkJobRequest instantiates a new ModelsOrchestratorJobsScheduleBulkJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults

`func NewModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults() *ModelsOrchestratorJobsScheduleBulkJobRequest`

NewModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults instantiates a new ModelsOrchestratorJobsScheduleBulkJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestratorIds

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetOrchestratorIds() []string`

GetOrchestratorIds returns the OrchestratorIds field if non-nil, zero value otherwise.

### GetOrchestratorIdsOk

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetOrchestratorIdsOk() (*[]string, bool)`

GetOrchestratorIdsOk returns a tuple with the OrchestratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorIds

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) SetOrchestratorIds(v []string)`

SetOrchestratorIds sets OrchestratorIds field to given value.


### GetJobTypeName

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetSchedule

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetJobFields() map[string]string`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) GetJobFieldsOk() (*map[string]string, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) SetJobFields(v map[string]string)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *ModelsOrchestratorJobsScheduleBulkJobRequest) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


