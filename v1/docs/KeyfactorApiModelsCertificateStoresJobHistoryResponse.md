# KeyfactorApiModelsCertificateStoresJobHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobHistoryId** | Pointer to **int64** |  | [optional] 
**AgentMachine** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**OperationStart** | Pointer to **time.Time** |  | [optional] 
**OperationEnd** | Pointer to **time.Time** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StorePath** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresJobHistoryResponse

`func NewKeyfactorApiModelsCertificateStoresJobHistoryResponse() *KeyfactorApiModelsCertificateStoresJobHistoryResponse`

NewKeyfactorApiModelsCertificateStoresJobHistoryResponse instantiates a new KeyfactorApiModelsCertificateStoresJobHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults

`func NewKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults() *KeyfactorApiModelsCertificateStoresJobHistoryResponse`

NewKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresJobHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobHistoryId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryId() int64`

GetJobHistoryId returns the JobHistoryId field if non-nil, zero value otherwise.

### GetJobHistoryIdOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryIdOk() (*int64, bool)`

GetJobHistoryIdOk returns a tuple with the JobHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHistoryId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobHistoryId(v int64)`

SetJobHistoryId sets JobHistoryId field to given value.

### HasJobHistoryId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobHistoryId() bool`

HasJobHistoryId returns a boolean if a field has been set.

### GetAgentMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachine() string`

GetAgentMachine returns the AgentMachine field if non-nil, zero value otherwise.

### GetAgentMachineOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachineOk() (*string, bool)`

GetAgentMachineOk returns a tuple with the AgentMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetAgentMachine(v string)`

SetAgentMachine sets AgentMachine field to given value.

### HasAgentMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasAgentMachine() bool`

HasAgentMachine returns a boolean if a field has been set.

### GetJobId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobType

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetOperationStart

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStart() time.Time`

GetOperationStart returns the OperationStart field if non-nil, zero value otherwise.

### GetOperationStartOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStartOk() (*time.Time, bool)`

GetOperationStartOk returns a tuple with the OperationStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStart

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationStart(v time.Time)`

SetOperationStart sets OperationStart field to given value.

### HasOperationStart

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationStart() bool`

HasOperationStart returns a boolean if a field has been set.

### GetOperationEnd

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEnd() time.Time`

GetOperationEnd returns the OperationEnd field if non-nil, zero value otherwise.

### GetOperationEndOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEndOk() (*time.Time, bool)`

GetOperationEndOk returns a tuple with the OperationEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationEnd

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationEnd(v time.Time)`

SetOperationEnd sets OperationEnd field to given value.

### HasOperationEnd

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationEnd() bool`

HasOperationEnd returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetResult(v int32)`

SetResult sets Result field to given value.

### HasResult

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorePath

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePath() string`

GetStorePath returns the StorePath field if non-nil, zero value otherwise.

### GetStorePathOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePathOk() (*string, bool)`

GetStorePathOk returns a tuple with the StorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePath

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStorePath(v string)`

SetStorePath sets StorePath field to given value.

### HasStorePath

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStorePath() bool`

HasStorePath returns a boolean if a field has been set.

### GetClientMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


