# KeyfactorApiModelsWorkflowsInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**CurrentStepId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Signals** | Pointer to [**[]KeyfactorApiModelsWorkflowsAvailableSignalResponse**](KeyfactorApiModelsWorkflowsAvailableSignalResponse.md) |  | [optional] 
**Definition** | Pointer to [**KeyfactorApiModelsWorkflowsInstanceDefinitionResponse**](KeyfactorApiModelsWorkflowsInstanceDefinitionResponse.md) |  | [optional] 
**CurrentStepDisplayName** | Pointer to **string** |  | [optional] 
**CurrentStepUniqueName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**InitialData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CurrentStateData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsInstanceResponse

`func NewKeyfactorApiModelsWorkflowsInstanceResponse() *KeyfactorApiModelsWorkflowsInstanceResponse`

NewKeyfactorApiModelsWorkflowsInstanceResponse instantiates a new KeyfactorApiModelsWorkflowsInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsInstanceResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsInstanceResponseWithDefaults() *KeyfactorApiModelsWorkflowsInstanceResponse`

NewKeyfactorApiModelsWorkflowsInstanceResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepId() string`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepIdOk() (*string, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepId(v string)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSignals

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetSignals() []KeyfactorApiModelsWorkflowsAvailableSignalResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetSignalsOk() (*[]KeyfactorApiModelsWorkflowsAvailableSignalResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetSignals(v []KeyfactorApiModelsWorkflowsAvailableSignalResponse)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetDefinition() KeyfactorApiModelsWorkflowsInstanceDefinitionResponse`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetDefinitionOk() (*KeyfactorApiModelsWorkflowsInstanceDefinitionResponse, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetDefinition(v KeyfactorApiModelsWorkflowsInstanceDefinitionResponse)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepDisplayName() string`

GetCurrentStepDisplayName returns the CurrentStepDisplayName field if non-nil, zero value otherwise.

### GetCurrentStepDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepDisplayNameOk() (*string, bool)`

GetCurrentStepDisplayNameOk returns a tuple with the CurrentStepDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepDisplayName(v string)`

SetCurrentStepDisplayName sets CurrentStepDisplayName field to given value.

### HasCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepDisplayName() bool`

HasCurrentStepDisplayName returns a boolean if a field has been set.

### GetCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepUniqueName() string`

GetCurrentStepUniqueName returns the CurrentStepUniqueName field if non-nil, zero value otherwise.

### GetCurrentStepUniqueNameOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepUniqueNameOk() (*string, bool)`

GetCurrentStepUniqueNameOk returns a tuple with the CurrentStepUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepUniqueName(v string)`

SetCurrentStepUniqueName sets CurrentStepUniqueName field to given value.

### HasCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepUniqueName() bool`

HasCurrentStepUniqueName returns a boolean if a field has been set.

### GetTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetInitialData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetInitialData() map[string]map[string]interface{}`

GetInitialData returns the InitialData field if non-nil, zero value otherwise.

### GetInitialDataOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetInitialDataOk() (*map[string]map[string]interface{}, bool)`

GetInitialDataOk returns a tuple with the InitialData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetInitialData(v map[string]map[string]interface{})`

SetInitialData sets InitialData field to given value.

### HasInitialData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasInitialData() bool`

HasInitialData returns a boolean if a field has been set.

### GetCurrentStateData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStateData() map[string]map[string]interface{}`

GetCurrentStateData returns the CurrentStateData field if non-nil, zero value otherwise.

### GetCurrentStateDataOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStateDataOk() (*map[string]map[string]interface{}, bool)`

GetCurrentStateDataOk returns a tuple with the CurrentStateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStateData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStateData(v map[string]map[string]interface{})`

SetCurrentStateData sets CurrentStateData field to given value.

### HasCurrentStateData

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStateData() bool`

HasCurrentStateData returns a boolean if a field has been set.

### GetReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetReferenceId() int64`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetReferenceIdOk() (*int64, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetReferenceId(v int64)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


