# KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse

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
**Result** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StorePath** | Pointer to **NullableString** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse() *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobHistoryId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryId() int64`

GetJobHistoryId returns the JobHistoryId field if non-nil, zero value otherwise.

### GetJobHistoryIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryIdOk() (*int64, bool)`

GetJobHistoryIdOk returns a tuple with the JobHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHistoryId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobHistoryId(v int64)`

SetJobHistoryId sets JobHistoryId field to given value.

### HasJobHistoryId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobHistoryId() bool`

HasJobHistoryId returns a boolean if a field has been set.

### GetAgentMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachine() string`

GetAgentMachine returns the AgentMachine field if non-nil, zero value otherwise.

### GetAgentMachineOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachineOk() (*string, bool)`

GetAgentMachineOk returns a tuple with the AgentMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetAgentMachine(v string)`

SetAgentMachine sets AgentMachine field to given value.

### HasAgentMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasAgentMachine() bool`

HasAgentMachine returns a boolean if a field has been set.

### SetAgentMachineNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetAgentMachineNil(b bool)`

 SetAgentMachineNil sets the value for AgentMachine to be an explicit nil

### UnsetAgentMachine
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetAgentMachine()`

UnsetAgentMachine ensures that no value is present for AgentMachine, not even an explicit nil
### GetJobId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetOperationStart

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStart() time.Time`

GetOperationStart returns the OperationStart field if non-nil, zero value otherwise.

### GetOperationStartOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStartOk() (*time.Time, bool)`

GetOperationStartOk returns a tuple with the OperationStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStart

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationStart(v time.Time)`

SetOperationStart sets OperationStart field to given value.

### HasOperationStart

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationStart() bool`

HasOperationStart returns a boolean if a field has been set.

### GetOperationEnd

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEnd() time.Time`

GetOperationEnd returns the OperationEnd field if non-nil, zero value otherwise.

### GetOperationEndOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEndOk() (*time.Time, bool)`

GetOperationEndOk returns a tuple with the OperationEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationEnd

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationEnd(v time.Time)`

SetOperationEnd sets OperationEnd field to given value.

### HasOperationEnd

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationEnd() bool`

HasOperationEnd returns a boolean if a field has been set.

### SetOperationEndNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationEndNil(b bool)`

 SetOperationEndNil sets the value for OperationEnd to be an explicit nil

### UnsetOperationEnd
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetOperationEnd()`

UnsetOperationEnd ensures that no value is present for OperationEnd, not even an explicit nil
### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetResult

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetResult(v int32)`

SetResult sets Result field to given value.

### HasResult

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorePath

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePath() string`

GetStorePath returns the StorePath field if non-nil, zero value otherwise.

### GetStorePathOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePathOk() (*string, bool)`

GetStorePathOk returns a tuple with the StorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePath

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStorePath(v string)`

SetStorePath sets StorePath field to given value.

### HasStorePath

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStorePath() bool`

HasStorePath returns a boolean if a field has been set.

### SetStorePathNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStorePathNil(b bool)`

 SetStorePathNil sets the value for StorePath to be an explicit nil

### UnsetStorePath
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetStorePath()`

UnsetStorePath ensures that no value is present for StorePath, not even an explicit nil
### GetClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


