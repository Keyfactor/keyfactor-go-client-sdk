# KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the step. | [optional] 
**ExtensionName** | Pointer to **NullableString** | The name of the extension. | [optional] 
**Outputs** | Pointer to **[]string** | The possible outputs of the step. | [optional] 
**ConfigurationParametersDefinition** | Pointer to [**map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse.md) |  | [optional] 
**SignalsDefinition** | Pointer to [**map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetConfigurationParametersDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetConfigurationParametersDefinition() map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse`

GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field if non-nil, zero value otherwise.

### GetConfigurationParametersDefinitionOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetConfigurationParametersDefinitionOk() (*map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse, bool)`

GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParametersDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetConfigurationParametersDefinition(v map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse)`

SetConfigurationParametersDefinition sets ConfigurationParametersDefinition field to given value.

### HasConfigurationParametersDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) HasConfigurationParametersDefinition() bool`

HasConfigurationParametersDefinition returns a boolean if a field has been set.

### SetConfigurationParametersDefinitionNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetConfigurationParametersDefinitionNil(b bool)`

 SetConfigurationParametersDefinitionNil sets the value for ConfigurationParametersDefinition to be an explicit nil

### UnsetConfigurationParametersDefinition
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) UnsetConfigurationParametersDefinition()`

UnsetConfigurationParametersDefinition ensures that no value is present for ConfigurationParametersDefinition, not even an explicit nil
### GetSignalsDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetSignalsDefinition() map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse`

GetSignalsDefinition returns the SignalsDefinition field if non-nil, zero value otherwise.

### GetSignalsDefinitionOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) GetSignalsDefinitionOk() (*map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse, bool)`

GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetSignalsDefinition(v map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse)`

SetSignalsDefinition sets SignalsDefinition field to given value.

### HasSignalsDefinition

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) HasSignalsDefinition() bool`

HasSignalsDefinition returns a boolean if a field has been set.

### SetSignalsDefinitionNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) SetSignalsDefinitionNil(b bool)`

 SetSignalsDefinitionNil sets the value for SignalsDefinition to be an explicit nil

### UnsetSignalsDefinition
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse) UnsetSignalsDefinition()`

UnsetSignalsDefinition ensures that no value is present for SignalsDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


