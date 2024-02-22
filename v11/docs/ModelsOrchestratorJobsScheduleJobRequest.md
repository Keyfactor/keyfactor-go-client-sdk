# ModelsOrchestratorJobsScheduleJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**JobTypeName** | **string** |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsOrchestratorJobsScheduleJobRequest

`func NewModelsOrchestratorJobsScheduleJobRequest(agentId string, jobTypeName string, ) *ModelsOrchestratorJobsScheduleJobRequest`

NewModelsOrchestratorJobsScheduleJobRequest instantiates a new ModelsOrchestratorJobsScheduleJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrchestratorJobsScheduleJobRequestWithDefaults

`func NewModelsOrchestratorJobsScheduleJobRequestWithDefaults() *ModelsOrchestratorJobsScheduleJobRequest`

NewModelsOrchestratorJobsScheduleJobRequestWithDefaults instantiates a new ModelsOrchestratorJobsScheduleJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsOrchestratorJobsScheduleJobRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetJobTypeName

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *ModelsOrchestratorJobsScheduleJobRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetSchedule

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsOrchestratorJobsScheduleJobRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsOrchestratorJobsScheduleJobRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetJobFields() map[string]string`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *ModelsOrchestratorJobsScheduleJobRequest) GetJobFieldsOk() (*map[string]string, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *ModelsOrchestratorJobsScheduleJobRequest) SetJobFields(v map[string]string)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *ModelsOrchestratorJobsScheduleJobRequest) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *ModelsOrchestratorJobsScheduleJobRequest) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *ModelsOrchestratorJobsScheduleJobRequest) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


