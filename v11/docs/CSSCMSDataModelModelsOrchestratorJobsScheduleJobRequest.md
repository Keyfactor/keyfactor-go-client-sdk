# CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**JobTypeName** | **string** |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest

`func NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest(agentId string, jobTypeName string, ) *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest`

NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest instantiates a new CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequestWithDefaults

`func NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequestWithDefaults() *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest`

NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequestWithDefaults instantiates a new CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetJobFields() map[string]string`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) GetJobFieldsOk() (*map[string]string, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) SetJobFields(v map[string]string)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


