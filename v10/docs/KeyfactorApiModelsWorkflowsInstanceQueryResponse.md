# KeyfactorApiModelsWorkflowsInstanceQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**CurrentStepId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**KeyfactorApiModelsWorkflowsInstanceDefinitionResponse**](KeyfactorApiModelsWorkflowsInstanceDefinitionResponse.md) |  | [optional] 
**CurrentStepDisplayName** | Pointer to **string** |  | [optional] 
**CurrentStepUniqueName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ReferenceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsInstanceQueryResponse

`func NewKeyfactorApiModelsWorkflowsInstanceQueryResponse() *KeyfactorApiModelsWorkflowsInstanceQueryResponse`

NewKeyfactorApiModelsWorkflowsInstanceQueryResponse instantiates a new KeyfactorApiModelsWorkflowsInstanceQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsInstanceQueryResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsInstanceQueryResponseWithDefaults() *KeyfactorApiModelsWorkflowsInstanceQueryResponse`

NewKeyfactorApiModelsWorkflowsInstanceQueryResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsInstanceQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepId() string`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepIdOk() (*string, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetCurrentStepId(v string)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetDefinition() KeyfactorApiModelsWorkflowsInstanceDefinitionResponse`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetDefinitionOk() (*KeyfactorApiModelsWorkflowsInstanceDefinitionResponse, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetDefinition(v KeyfactorApiModelsWorkflowsInstanceDefinitionResponse)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepDisplayName() string`

GetCurrentStepDisplayName returns the CurrentStepDisplayName field if non-nil, zero value otherwise.

### GetCurrentStepDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepDisplayNameOk() (*string, bool)`

GetCurrentStepDisplayNameOk returns a tuple with the CurrentStepDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetCurrentStepDisplayName(v string)`

SetCurrentStepDisplayName sets CurrentStepDisplayName field to given value.

### HasCurrentStepDisplayName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasCurrentStepDisplayName() bool`

HasCurrentStepDisplayName returns a boolean if a field has been set.

### GetCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepUniqueName() string`

GetCurrentStepUniqueName returns the CurrentStepUniqueName field if non-nil, zero value otherwise.

### GetCurrentStepUniqueNameOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetCurrentStepUniqueNameOk() (*string, bool)`

GetCurrentStepUniqueNameOk returns a tuple with the CurrentStepUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetCurrentStepUniqueName(v string)`

SetCurrentStepUniqueName sets CurrentStepUniqueName field to given value.

### HasCurrentStepUniqueName

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasCurrentStepUniqueName() bool`

HasCurrentStepUniqueName returns a boolean if a field has been set.

### GetTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetReferenceId() int64`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) GetReferenceIdOk() (*int64, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) SetReferenceId(v int64)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KeyfactorApiModelsWorkflowsInstanceQueryResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


