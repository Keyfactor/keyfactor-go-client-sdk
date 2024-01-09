# CSSCMSDataModelModelsOrchestratorJobsJob

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

### NewCSSCMSDataModelModelsOrchestratorJobsJob

`func NewCSSCMSDataModelModelsOrchestratorJobsJob() *CSSCMSDataModelModelsOrchestratorJobsJob`

NewCSSCMSDataModelModelsOrchestratorJobsJob instantiates a new CSSCMSDataModelModelsOrchestratorJobsJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsOrchestratorJobsJobWithDefaults

`func NewCSSCMSDataModelModelsOrchestratorJobsJobWithDefaults() *CSSCMSDataModelModelsOrchestratorJobsJob`

NewCSSCMSDataModelModelsOrchestratorJobsJobWithDefaults instantiates a new CSSCMSDataModelModelsOrchestratorJobsJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientMachine

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetTarget

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRequested

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetRequested() string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetRequestedOk() (*string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetRequested(v string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### SetRequestedNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetRequestedNil(b bool)`

 SetRequestedNil sets the value for Requested to be an explicit nil

### UnsetRequested
`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) UnsetRequested()`

UnsetRequested ensures that no value is present for Requested, not even an explicit nil
### GetJobType

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *CSSCMSDataModelModelsOrchestratorJobsJob) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


