# KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobAuditIds** | Pointer to **[]int64** | List of orchestrator job audit ids to be rescheduled | [optional] 
**Query** | Pointer to **string** | Query identifying orchestrator jobs to be rescheduled | [optional] 

## Methods

### NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest

`func NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest() *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest`

NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest instantiates a new KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequestWithDefaults

`func NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequestWithDefaults() *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest`

NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequestWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobAuditIds

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) GetJobAuditIds() []int64`

GetJobAuditIds returns the JobAuditIds field if non-nil, zero value otherwise.

### GetJobAuditIdsOk

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) GetJobAuditIdsOk() (*[]int64, bool)`

GetJobAuditIdsOk returns a tuple with the JobAuditIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAuditIds

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) SetJobAuditIds(v []int64)`

SetJobAuditIds sets JobAuditIds field to given value.

### HasJobAuditIds

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) HasJobAuditIds() bool`

HasJobAuditIds returns a boolean if a field has been set.

### GetQuery

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


