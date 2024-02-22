# ModelsOrchestratorJobsJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Target** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Requested** | Pointer to **NullableString** |  | [optional] 
**JobType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsOrchestratorJobsJob

`func NewModelsOrchestratorJobsJob() *ModelsOrchestratorJobsJob`

NewModelsOrchestratorJobsJob instantiates a new ModelsOrchestratorJobsJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrchestratorJobsJobWithDefaults

`func NewModelsOrchestratorJobsJobWithDefaults() *ModelsOrchestratorJobsJob`

NewModelsOrchestratorJobsJobWithDefaults instantiates a new ModelsOrchestratorJobsJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsOrchestratorJobsJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsOrchestratorJobsJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsOrchestratorJobsJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsOrchestratorJobsJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientMachine

`func (o *ModelsOrchestratorJobsJob) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *ModelsOrchestratorJobsJob) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *ModelsOrchestratorJobsJob) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *ModelsOrchestratorJobsJob) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *ModelsOrchestratorJobsJob) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *ModelsOrchestratorJobsJob) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetTarget

`func (o *ModelsOrchestratorJobsJob) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ModelsOrchestratorJobsJob) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ModelsOrchestratorJobsJob) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ModelsOrchestratorJobsJob) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ModelsOrchestratorJobsJob) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ModelsOrchestratorJobsJob) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetSchedule

`func (o *ModelsOrchestratorJobsJob) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsOrchestratorJobsJob) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsOrchestratorJobsJob) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsOrchestratorJobsJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRequested

`func (o *ModelsOrchestratorJobsJob) GetRequested() string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *ModelsOrchestratorJobsJob) GetRequestedOk() (*string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *ModelsOrchestratorJobsJob) SetRequested(v string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *ModelsOrchestratorJobsJob) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### SetRequestedNil

`func (o *ModelsOrchestratorJobsJob) SetRequestedNil(b bool)`

 SetRequestedNil sets the value for Requested to be an explicit nil

### UnsetRequested
`func (o *ModelsOrchestratorJobsJob) UnsetRequested()`

UnsetRequested ensures that no value is present for Requested, not even an explicit nil
### GetJobType

`func (o *ModelsOrchestratorJobsJob) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ModelsOrchestratorJobsJob) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ModelsOrchestratorJobsJob) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ModelsOrchestratorJobsJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *ModelsOrchestratorJobsJob) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *ModelsOrchestratorJobsJob) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


