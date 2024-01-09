# CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestratorIds** | **[]string** |  | 
**JobTypeName** | **string** |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest

`func NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest(orchestratorIds []string, jobTypeName string, ) *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest`

NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest instantiates a new CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults

`func NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults() *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest`

NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequestWithDefaults instantiates a new CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestratorIds

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetOrchestratorIds() []string`

GetOrchestratorIds returns the OrchestratorIds field if non-nil, zero value otherwise.

### GetOrchestratorIdsOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetOrchestratorIdsOk() (*[]string, bool)`

GetOrchestratorIdsOk returns a tuple with the OrchestratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorIds

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) SetOrchestratorIds(v []string)`

SetOrchestratorIds sets OrchestratorIds field to given value.


### GetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetJobFields() map[string]string`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) GetJobFieldsOk() (*map[string]string, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) SetJobFields(v map[string]string)`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


