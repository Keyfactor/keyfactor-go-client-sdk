# CertificateStoresJobHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobHistoryId** | Pointer to **int64** |  | [optional] 
**AgentMachine** | Pointer to **NullableString** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobType** | Pointer to **NullableString** |  | [optional] 
**OperationStart** | Pointer to **time.Time** |  | [optional] 
**OperationEnd** | Pointer to **NullableTime** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Result** | Pointer to [**KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobResult**](KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobResult.md) |  | [optional] 
**Status** | Pointer to [**KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobStatus**](KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobStatus.md) |  | [optional] 
**StorePath** | Pointer to **NullableString** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateStoresJobHistoryResponse

`func NewCertificateStoresJobHistoryResponse() *CertificateStoresJobHistoryResponse`

NewCertificateStoresJobHistoryResponse instantiates a new CertificateStoresJobHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresJobHistoryResponseWithDefaults

`func NewCertificateStoresJobHistoryResponseWithDefaults() *CertificateStoresJobHistoryResponse`

NewCertificateStoresJobHistoryResponseWithDefaults instantiates a new CertificateStoresJobHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobHistoryId

`func (o *CertificateStoresJobHistoryResponse) GetJobHistoryId() int64`

GetJobHistoryId returns the JobHistoryId field if non-nil, zero value otherwise.

### GetJobHistoryIdOk

`func (o *CertificateStoresJobHistoryResponse) GetJobHistoryIdOk() (*int64, bool)`

GetJobHistoryIdOk returns a tuple with the JobHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHistoryId

`func (o *CertificateStoresJobHistoryResponse) SetJobHistoryId(v int64)`

SetJobHistoryId sets JobHistoryId field to given value.

### HasJobHistoryId

`func (o *CertificateStoresJobHistoryResponse) HasJobHistoryId() bool`

HasJobHistoryId returns a boolean if a field has been set.

### GetAgentMachine

`func (o *CertificateStoresJobHistoryResponse) GetAgentMachine() string`

GetAgentMachine returns the AgentMachine field if non-nil, zero value otherwise.

### GetAgentMachineOk

`func (o *CertificateStoresJobHistoryResponse) GetAgentMachineOk() (*string, bool)`

GetAgentMachineOk returns a tuple with the AgentMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachine

`func (o *CertificateStoresJobHistoryResponse) SetAgentMachine(v string)`

SetAgentMachine sets AgentMachine field to given value.

### HasAgentMachine

`func (o *CertificateStoresJobHistoryResponse) HasAgentMachine() bool`

HasAgentMachine returns a boolean if a field has been set.

### SetAgentMachineNil

`func (o *CertificateStoresJobHistoryResponse) SetAgentMachineNil(b bool)`

 SetAgentMachineNil sets the value for AgentMachine to be an explicit nil

### UnsetAgentMachine
`func (o *CertificateStoresJobHistoryResponse) UnsetAgentMachine()`

UnsetAgentMachine ensures that no value is present for AgentMachine, not even an explicit nil
### GetJobId

`func (o *CertificateStoresJobHistoryResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CertificateStoresJobHistoryResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CertificateStoresJobHistoryResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CertificateStoresJobHistoryResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSchedule

`func (o *CertificateStoresJobHistoryResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CertificateStoresJobHistoryResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CertificateStoresJobHistoryResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CertificateStoresJobHistoryResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobType

`func (o *CertificateStoresJobHistoryResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CertificateStoresJobHistoryResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CertificateStoresJobHistoryResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CertificateStoresJobHistoryResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *CertificateStoresJobHistoryResponse) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *CertificateStoresJobHistoryResponse) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetOperationStart

`func (o *CertificateStoresJobHistoryResponse) GetOperationStart() time.Time`

GetOperationStart returns the OperationStart field if non-nil, zero value otherwise.

### GetOperationStartOk

`func (o *CertificateStoresJobHistoryResponse) GetOperationStartOk() (*time.Time, bool)`

GetOperationStartOk returns a tuple with the OperationStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStart

`func (o *CertificateStoresJobHistoryResponse) SetOperationStart(v time.Time)`

SetOperationStart sets OperationStart field to given value.

### HasOperationStart

`func (o *CertificateStoresJobHistoryResponse) HasOperationStart() bool`

HasOperationStart returns a boolean if a field has been set.

### GetOperationEnd

`func (o *CertificateStoresJobHistoryResponse) GetOperationEnd() time.Time`

GetOperationEnd returns the OperationEnd field if non-nil, zero value otherwise.

### GetOperationEndOk

`func (o *CertificateStoresJobHistoryResponse) GetOperationEndOk() (*time.Time, bool)`

GetOperationEndOk returns a tuple with the OperationEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationEnd

`func (o *CertificateStoresJobHistoryResponse) SetOperationEnd(v time.Time)`

SetOperationEnd sets OperationEnd field to given value.

### HasOperationEnd

`func (o *CertificateStoresJobHistoryResponse) HasOperationEnd() bool`

HasOperationEnd returns a boolean if a field has been set.

### SetOperationEndNil

`func (o *CertificateStoresJobHistoryResponse) SetOperationEndNil(b bool)`

 SetOperationEndNil sets the value for OperationEnd to be an explicit nil

### UnsetOperationEnd
`func (o *CertificateStoresJobHistoryResponse) UnsetOperationEnd()`

UnsetOperationEnd ensures that no value is present for OperationEnd, not even an explicit nil
### GetMessage

`func (o *CertificateStoresJobHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertificateStoresJobHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertificateStoresJobHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CertificateStoresJobHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CertificateStoresJobHistoryResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CertificateStoresJobHistoryResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetResult

`func (o *CertificateStoresJobHistoryResponse) GetResult() KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CertificateStoresJobHistoryResponse) GetResultOk() (*KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CertificateStoresJobHistoryResponse) SetResult(v KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CertificateStoresJobHistoryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateStoresJobHistoryResponse) GetStatus() KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateStoresJobHistoryResponse) GetStatusOk() (*KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateStoresJobHistoryResponse) SetStatus(v KeyfactorOrchestratorsCommonEnumsOrchestratorJobStatusJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateStoresJobHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorePath

`func (o *CertificateStoresJobHistoryResponse) GetStorePath() string`

GetStorePath returns the StorePath field if non-nil, zero value otherwise.

### GetStorePathOk

`func (o *CertificateStoresJobHistoryResponse) GetStorePathOk() (*string, bool)`

GetStorePathOk returns a tuple with the StorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePath

`func (o *CertificateStoresJobHistoryResponse) SetStorePath(v string)`

SetStorePath sets StorePath field to given value.

### HasStorePath

`func (o *CertificateStoresJobHistoryResponse) HasStorePath() bool`

HasStorePath returns a boolean if a field has been set.

### SetStorePathNil

`func (o *CertificateStoresJobHistoryResponse) SetStorePathNil(b bool)`

 SetStorePathNil sets the value for StorePath to be an explicit nil

### UnsetStorePath
`func (o *CertificateStoresJobHistoryResponse) UnsetStorePath()`

UnsetStorePath ensures that no value is present for StorePath, not even an explicit nil
### GetClientMachine

`func (o *CertificateStoresJobHistoryResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CertificateStoresJobHistoryResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CertificateStoresJobHistoryResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CertificateStoresJobHistoryResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CertificateStoresJobHistoryResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CertificateStoresJobHistoryResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


