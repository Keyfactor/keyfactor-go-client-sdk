# WorkflowsInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**KeyfactorWorkflowsWorkflowInstanceStatus**](KeyfactorWorkflowsWorkflowInstanceStatus.md) |  | [optional] 
**CurrentStepId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**Signals** | Pointer to [**[]WorkflowsAvailableSignalResponse**](WorkflowsAvailableSignalResponse.md) |  | [optional] 
**Definition** | Pointer to [**WorkflowsInstanceDefinitionResponse**](WorkflowsInstanceDefinitionResponse.md) |  | [optional] 
**CurrentStepDisplayName** | Pointer to **NullableString** |  | [optional] 
**CurrentStepUniqueName** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableTime** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**InitialData** | Pointer to **map[string]interface{}** |  | [optional] 
**CurrentStateData** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewWorkflowsInstanceResponse

`func NewWorkflowsInstanceResponse() *WorkflowsInstanceResponse`

NewWorkflowsInstanceResponse instantiates a new WorkflowsInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsInstanceResponseWithDefaults

`func NewWorkflowsInstanceResponseWithDefaults() *WorkflowsInstanceResponse`

NewWorkflowsInstanceResponseWithDefaults instantiates a new WorkflowsInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsInstanceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsInstanceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowsInstanceResponse) GetStatus() KeyfactorWorkflowsWorkflowInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowsInstanceResponse) GetStatusOk() (*KeyfactorWorkflowsWorkflowInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowsInstanceResponse) SetStatus(v KeyfactorWorkflowsWorkflowInstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowsInstanceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentStepId

`func (o *WorkflowsInstanceResponse) GetCurrentStepId() string`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *WorkflowsInstanceResponse) GetCurrentStepIdOk() (*string, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *WorkflowsInstanceResponse) SetCurrentStepId(v string)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *WorkflowsInstanceResponse) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *WorkflowsInstanceResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *WorkflowsInstanceResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *WorkflowsInstanceResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *WorkflowsInstanceResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *WorkflowsInstanceResponse) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *WorkflowsInstanceResponse) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetSignals

`func (o *WorkflowsInstanceResponse) GetSignals() []WorkflowsAvailableSignalResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WorkflowsInstanceResponse) GetSignalsOk() (*[]WorkflowsAvailableSignalResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WorkflowsInstanceResponse) SetSignals(v []WorkflowsAvailableSignalResponse)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *WorkflowsInstanceResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *WorkflowsInstanceResponse) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *WorkflowsInstanceResponse) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetDefinition

`func (o *WorkflowsInstanceResponse) GetDefinition() WorkflowsInstanceDefinitionResponse`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WorkflowsInstanceResponse) GetDefinitionOk() (*WorkflowsInstanceDefinitionResponse, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WorkflowsInstanceResponse) SetDefinition(v WorkflowsInstanceDefinitionResponse)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *WorkflowsInstanceResponse) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetCurrentStepDisplayName

`func (o *WorkflowsInstanceResponse) GetCurrentStepDisplayName() string`

GetCurrentStepDisplayName returns the CurrentStepDisplayName field if non-nil, zero value otherwise.

### GetCurrentStepDisplayNameOk

`func (o *WorkflowsInstanceResponse) GetCurrentStepDisplayNameOk() (*string, bool)`

GetCurrentStepDisplayNameOk returns a tuple with the CurrentStepDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepDisplayName

`func (o *WorkflowsInstanceResponse) SetCurrentStepDisplayName(v string)`

SetCurrentStepDisplayName sets CurrentStepDisplayName field to given value.

### HasCurrentStepDisplayName

`func (o *WorkflowsInstanceResponse) HasCurrentStepDisplayName() bool`

HasCurrentStepDisplayName returns a boolean if a field has been set.

### SetCurrentStepDisplayNameNil

`func (o *WorkflowsInstanceResponse) SetCurrentStepDisplayNameNil(b bool)`

 SetCurrentStepDisplayNameNil sets the value for CurrentStepDisplayName to be an explicit nil

### UnsetCurrentStepDisplayName
`func (o *WorkflowsInstanceResponse) UnsetCurrentStepDisplayName()`

UnsetCurrentStepDisplayName ensures that no value is present for CurrentStepDisplayName, not even an explicit nil
### GetCurrentStepUniqueName

`func (o *WorkflowsInstanceResponse) GetCurrentStepUniqueName() string`

GetCurrentStepUniqueName returns the CurrentStepUniqueName field if non-nil, zero value otherwise.

### GetCurrentStepUniqueNameOk

`func (o *WorkflowsInstanceResponse) GetCurrentStepUniqueNameOk() (*string, bool)`

GetCurrentStepUniqueNameOk returns a tuple with the CurrentStepUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepUniqueName

`func (o *WorkflowsInstanceResponse) SetCurrentStepUniqueName(v string)`

SetCurrentStepUniqueName sets CurrentStepUniqueName field to given value.

### HasCurrentStepUniqueName

`func (o *WorkflowsInstanceResponse) HasCurrentStepUniqueName() bool`

HasCurrentStepUniqueName returns a boolean if a field has been set.

### SetCurrentStepUniqueNameNil

`func (o *WorkflowsInstanceResponse) SetCurrentStepUniqueNameNil(b bool)`

 SetCurrentStepUniqueNameNil sets the value for CurrentStepUniqueName to be an explicit nil

### UnsetCurrentStepUniqueName
`func (o *WorkflowsInstanceResponse) UnsetCurrentStepUniqueName()`

UnsetCurrentStepUniqueName ensures that no value is present for CurrentStepUniqueName, not even an explicit nil
### GetTitle

`func (o *WorkflowsInstanceResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowsInstanceResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowsInstanceResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowsInstanceResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkflowsInstanceResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkflowsInstanceResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLastModified

`func (o *WorkflowsInstanceResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WorkflowsInstanceResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WorkflowsInstanceResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *WorkflowsInstanceResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *WorkflowsInstanceResponse) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *WorkflowsInstanceResponse) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetStartDate

`func (o *WorkflowsInstanceResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WorkflowsInstanceResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WorkflowsInstanceResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WorkflowsInstanceResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetInitialData

`func (o *WorkflowsInstanceResponse) GetInitialData() map[string]interface{}`

GetInitialData returns the InitialData field if non-nil, zero value otherwise.

### GetInitialDataOk

`func (o *WorkflowsInstanceResponse) GetInitialDataOk() (*map[string]interface{}, bool)`

GetInitialDataOk returns a tuple with the InitialData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialData

`func (o *WorkflowsInstanceResponse) SetInitialData(v map[string]interface{})`

SetInitialData sets InitialData field to given value.

### HasInitialData

`func (o *WorkflowsInstanceResponse) HasInitialData() bool`

HasInitialData returns a boolean if a field has been set.

### SetInitialDataNil

`func (o *WorkflowsInstanceResponse) SetInitialDataNil(b bool)`

 SetInitialDataNil sets the value for InitialData to be an explicit nil

### UnsetInitialData
`func (o *WorkflowsInstanceResponse) UnsetInitialData()`

UnsetInitialData ensures that no value is present for InitialData, not even an explicit nil
### GetCurrentStateData

`func (o *WorkflowsInstanceResponse) GetCurrentStateData() map[string]interface{}`

GetCurrentStateData returns the CurrentStateData field if non-nil, zero value otherwise.

### GetCurrentStateDataOk

`func (o *WorkflowsInstanceResponse) GetCurrentStateDataOk() (*map[string]interface{}, bool)`

GetCurrentStateDataOk returns a tuple with the CurrentStateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStateData

`func (o *WorkflowsInstanceResponse) SetCurrentStateData(v map[string]interface{})`

SetCurrentStateData sets CurrentStateData field to given value.

### HasCurrentStateData

`func (o *WorkflowsInstanceResponse) HasCurrentStateData() bool`

HasCurrentStateData returns a boolean if a field has been set.

### SetCurrentStateDataNil

`func (o *WorkflowsInstanceResponse) SetCurrentStateDataNil(b bool)`

 SetCurrentStateDataNil sets the value for CurrentStateData to be an explicit nil

### UnsetCurrentStateData
`func (o *WorkflowsInstanceResponse) UnsetCurrentStateData()`

UnsetCurrentStateData ensures that no value is present for CurrentStateData, not even an explicit nil
### GetReferenceId

`func (o *WorkflowsInstanceResponse) GetReferenceId() int64`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WorkflowsInstanceResponse) GetReferenceIdOk() (*int64, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WorkflowsInstanceResponse) SetReferenceId(v int64)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *WorkflowsInstanceResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


