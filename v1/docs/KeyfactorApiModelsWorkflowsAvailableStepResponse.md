# KeyfactorApiModelsWorkflowsAvailableStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the step. | [optional] 
**ExtensionName** | Pointer to **string** | The name of the extension. | [optional] 
**Outputs** | Pointer to **[]string** | The possible outputs of the step. | [optional] 
**ConfigurationParametersDefinition** | Pointer to [**map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse**](KeyfactorApiModelsWorkflowsParameterDefinitionResponse.md) |  | [optional] 
**SignalsDefinition** | Pointer to [**map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse**](KeyfactorApiModelsWorkflowsSignalDefinitionResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsAvailableStepResponse

`func NewKeyfactorApiModelsWorkflowsAvailableStepResponse() *KeyfactorApiModelsWorkflowsAvailableStepResponse`

NewKeyfactorApiModelsWorkflowsAvailableStepResponse instantiates a new KeyfactorApiModelsWorkflowsAvailableStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults() *KeyfactorApiModelsWorkflowsAvailableStepResponse`

NewKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsAvailableStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetOutputs

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetConfigurationParametersDefinition() map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse`

GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field if non-nil, zero value otherwise.

### GetConfigurationParametersDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetConfigurationParametersDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse, bool)`

GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) SetConfigurationParametersDefinition(v map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse)`

SetConfigurationParametersDefinition sets ConfigurationParametersDefinition field to given value.

### HasConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) HasConfigurationParametersDefinition() bool`

HasConfigurationParametersDefinition returns a boolean if a field has been set.

### GetSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetSignalsDefinition() map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse`

GetSignalsDefinition returns the SignalsDefinition field if non-nil, zero value otherwise.

### GetSignalsDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) GetSignalsDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse, bool)`

GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) SetSignalsDefinition(v map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse)`

SetSignalsDefinition sets SignalsDefinition field to given value.

### HasSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepResponse) HasSignalsDefinition() bool`

HasSignalsDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


