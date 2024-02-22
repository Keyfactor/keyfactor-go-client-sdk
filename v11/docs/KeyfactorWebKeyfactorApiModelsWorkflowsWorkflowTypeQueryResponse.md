# KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowType** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**ContextParameters** | Pointer to **[]string** |  | [optional] 
**BuiltInSteps** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
### GetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetContextParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetContextParameters() []string`

GetContextParameters returns the ContextParameters field if non-nil, zero value otherwise.

### GetContextParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetContextParametersOk() (*[]string, bool)`

GetContextParametersOk returns a tuple with the ContextParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetContextParameters(v []string)`

SetContextParameters sets ContextParameters field to given value.

### HasContextParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasContextParameters() bool`

HasContextParameters returns a boolean if a field has been set.

### SetContextParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetContextParametersNil(b bool)`

 SetContextParametersNil sets the value for ContextParameters to be an explicit nil

### UnsetContextParameters
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) UnsetContextParameters()`

UnsetContextParameters ensures that no value is present for ContextParameters, not even an explicit nil
### GetBuiltInSteps

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetBuiltInSteps() []KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse`

GetBuiltInSteps returns the BuiltInSteps field if non-nil, zero value otherwise.

### GetBuiltInStepsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetBuiltInStepsOk() (*[]KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse, bool)`

GetBuiltInStepsOk returns a tuple with the BuiltInSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInSteps

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetBuiltInSteps(v []KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse)`

SetBuiltInSteps sets BuiltInSteps field to given value.

### HasBuiltInSteps

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasBuiltInSteps() bool`

HasBuiltInSteps returns a boolean if a field has been set.

### SetBuiltInStepsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetBuiltInStepsNil(b bool)`

 SetBuiltInStepsNil sets the value for BuiltInSteps to be an explicit nil

### UnsetBuiltInSteps
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) UnsetBuiltInSteps()`

UnsetBuiltInSteps ensures that no value is present for BuiltInSteps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


