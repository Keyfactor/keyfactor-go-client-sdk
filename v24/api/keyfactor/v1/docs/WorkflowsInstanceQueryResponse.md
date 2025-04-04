# WorkflowsInstanceQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**KeyfactorWorkflowsWorkflowInstanceStatus**](KeyfactorWorkflowsWorkflowInstanceStatus.md) |  | [optional] 
**CurrentStepId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**Definition** | Pointer to [**WorkflowsInstanceDefinitionResponse**](WorkflowsInstanceDefinitionResponse.md) |  | [optional] 
**CurrentStepDisplayName** | Pointer to **NullableString** |  | [optional] 
**CurrentStepUniqueName** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableTime** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ReferenceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewWorkflowsInstanceQueryResponse

`func NewWorkflowsInstanceQueryResponse() *WorkflowsInstanceQueryResponse`

NewWorkflowsInstanceQueryResponse instantiates a new WorkflowsInstanceQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsInstanceQueryResponseWithDefaults

`func NewWorkflowsInstanceQueryResponseWithDefaults() *WorkflowsInstanceQueryResponse`

NewWorkflowsInstanceQueryResponseWithDefaults instantiates a new WorkflowsInstanceQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsInstanceQueryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsInstanceQueryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsInstanceQueryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsInstanceQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowsInstanceQueryResponse) GetStatus() KeyfactorWorkflowsWorkflowInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowsInstanceQueryResponse) GetStatusOk() (*KeyfactorWorkflowsWorkflowInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowsInstanceQueryResponse) SetStatus(v KeyfactorWorkflowsWorkflowInstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowsInstanceQueryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentStepId

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepId() string`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepIdOk() (*string, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *WorkflowsInstanceQueryResponse) SetCurrentStepId(v string)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *WorkflowsInstanceQueryResponse) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *WorkflowsInstanceQueryResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *WorkflowsInstanceQueryResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *WorkflowsInstanceQueryResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *WorkflowsInstanceQueryResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *WorkflowsInstanceQueryResponse) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *WorkflowsInstanceQueryResponse) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetDefinition

`func (o *WorkflowsInstanceQueryResponse) GetDefinition() WorkflowsInstanceDefinitionResponse`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WorkflowsInstanceQueryResponse) GetDefinitionOk() (*WorkflowsInstanceDefinitionResponse, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WorkflowsInstanceQueryResponse) SetDefinition(v WorkflowsInstanceDefinitionResponse)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *WorkflowsInstanceQueryResponse) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetCurrentStepDisplayName

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepDisplayName() string`

GetCurrentStepDisplayName returns the CurrentStepDisplayName field if non-nil, zero value otherwise.

### GetCurrentStepDisplayNameOk

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepDisplayNameOk() (*string, bool)`

GetCurrentStepDisplayNameOk returns a tuple with the CurrentStepDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepDisplayName

`func (o *WorkflowsInstanceQueryResponse) SetCurrentStepDisplayName(v string)`

SetCurrentStepDisplayName sets CurrentStepDisplayName field to given value.

### HasCurrentStepDisplayName

`func (o *WorkflowsInstanceQueryResponse) HasCurrentStepDisplayName() bool`

HasCurrentStepDisplayName returns a boolean if a field has been set.

### SetCurrentStepDisplayNameNil

`func (o *WorkflowsInstanceQueryResponse) SetCurrentStepDisplayNameNil(b bool)`

 SetCurrentStepDisplayNameNil sets the value for CurrentStepDisplayName to be an explicit nil

### UnsetCurrentStepDisplayName
`func (o *WorkflowsInstanceQueryResponse) UnsetCurrentStepDisplayName()`

UnsetCurrentStepDisplayName ensures that no value is present for CurrentStepDisplayName, not even an explicit nil
### GetCurrentStepUniqueName

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepUniqueName() string`

GetCurrentStepUniqueName returns the CurrentStepUniqueName field if non-nil, zero value otherwise.

### GetCurrentStepUniqueNameOk

`func (o *WorkflowsInstanceQueryResponse) GetCurrentStepUniqueNameOk() (*string, bool)`

GetCurrentStepUniqueNameOk returns a tuple with the CurrentStepUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepUniqueName

`func (o *WorkflowsInstanceQueryResponse) SetCurrentStepUniqueName(v string)`

SetCurrentStepUniqueName sets CurrentStepUniqueName field to given value.

### HasCurrentStepUniqueName

`func (o *WorkflowsInstanceQueryResponse) HasCurrentStepUniqueName() bool`

HasCurrentStepUniqueName returns a boolean if a field has been set.

### SetCurrentStepUniqueNameNil

`func (o *WorkflowsInstanceQueryResponse) SetCurrentStepUniqueNameNil(b bool)`

 SetCurrentStepUniqueNameNil sets the value for CurrentStepUniqueName to be an explicit nil

### UnsetCurrentStepUniqueName
`func (o *WorkflowsInstanceQueryResponse) UnsetCurrentStepUniqueName()`

UnsetCurrentStepUniqueName ensures that no value is present for CurrentStepUniqueName, not even an explicit nil
### GetTitle

`func (o *WorkflowsInstanceQueryResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowsInstanceQueryResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowsInstanceQueryResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowsInstanceQueryResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkflowsInstanceQueryResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkflowsInstanceQueryResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLastModified

`func (o *WorkflowsInstanceQueryResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WorkflowsInstanceQueryResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WorkflowsInstanceQueryResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *WorkflowsInstanceQueryResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *WorkflowsInstanceQueryResponse) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *WorkflowsInstanceQueryResponse) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetStartDate

`func (o *WorkflowsInstanceQueryResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WorkflowsInstanceQueryResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WorkflowsInstanceQueryResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WorkflowsInstanceQueryResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *WorkflowsInstanceQueryResponse) GetReferenceId() int64`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WorkflowsInstanceQueryResponse) GetReferenceIdOk() (*int64, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WorkflowsInstanceQueryResponse) SetReferenceId(v int64)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *WorkflowsInstanceQueryResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


