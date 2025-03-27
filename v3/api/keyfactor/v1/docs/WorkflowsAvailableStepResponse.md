# WorkflowsAvailableStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the step. | [optional] 
**ExtensionName** | Pointer to **NullableString** | The name of the extension. | [optional] 
**Outputs** | Pointer to **[]string** | The possible outputs of the step. | [optional] 
**ConfigurationParametersDefinition** | Pointer to [**map[string]WorkflowsParameterDefinitionResponse**](WorkflowsParameterDefinitionResponse.md) |  | [optional] 
**SignalsDefinition** | Pointer to [**map[string]WorkflowsSignalDefinitionResponse**](WorkflowsSignalDefinitionResponse.md) |  | [optional] 

## Methods

### NewWorkflowsAvailableStepResponse

`func NewWorkflowsAvailableStepResponse() *WorkflowsAvailableStepResponse`

NewWorkflowsAvailableStepResponse instantiates a new WorkflowsAvailableStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsAvailableStepResponseWithDefaults

`func NewWorkflowsAvailableStepResponseWithDefaults() *WorkflowsAvailableStepResponse`

NewWorkflowsAvailableStepResponseWithDefaults instantiates a new WorkflowsAvailableStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsAvailableStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsAvailableStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsAvailableStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsAvailableStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsAvailableStepResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsAvailableStepResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExtensionName

`func (o *WorkflowsAvailableStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *WorkflowsAvailableStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *WorkflowsAvailableStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *WorkflowsAvailableStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *WorkflowsAvailableStepResponse) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *WorkflowsAvailableStepResponse) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetOutputs

`func (o *WorkflowsAvailableStepResponse) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowsAvailableStepResponse) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowsAvailableStepResponse) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowsAvailableStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowsAvailableStepResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowsAvailableStepResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepResponse) GetConfigurationParametersDefinition() map[string]WorkflowsParameterDefinitionResponse`

GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field if non-nil, zero value otherwise.

### GetConfigurationParametersDefinitionOk

`func (o *WorkflowsAvailableStepResponse) GetConfigurationParametersDefinitionOk() (*map[string]WorkflowsParameterDefinitionResponse, bool)`

GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepResponse) SetConfigurationParametersDefinition(v map[string]WorkflowsParameterDefinitionResponse)`

SetConfigurationParametersDefinition sets ConfigurationParametersDefinition field to given value.

### HasConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepResponse) HasConfigurationParametersDefinition() bool`

HasConfigurationParametersDefinition returns a boolean if a field has been set.

### SetConfigurationParametersDefinitionNil

`func (o *WorkflowsAvailableStepResponse) SetConfigurationParametersDefinitionNil(b bool)`

 SetConfigurationParametersDefinitionNil sets the value for ConfigurationParametersDefinition to be an explicit nil

### UnsetConfigurationParametersDefinition
`func (o *WorkflowsAvailableStepResponse) UnsetConfigurationParametersDefinition()`

UnsetConfigurationParametersDefinition ensures that no value is present for ConfigurationParametersDefinition, not even an explicit nil
### GetSignalsDefinition

`func (o *WorkflowsAvailableStepResponse) GetSignalsDefinition() map[string]WorkflowsSignalDefinitionResponse`

GetSignalsDefinition returns the SignalsDefinition field if non-nil, zero value otherwise.

### GetSignalsDefinitionOk

`func (o *WorkflowsAvailableStepResponse) GetSignalsDefinitionOk() (*map[string]WorkflowsSignalDefinitionResponse, bool)`

GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDefinition

`func (o *WorkflowsAvailableStepResponse) SetSignalsDefinition(v map[string]WorkflowsSignalDefinitionResponse)`

SetSignalsDefinition sets SignalsDefinition field to given value.

### HasSignalsDefinition

`func (o *WorkflowsAvailableStepResponse) HasSignalsDefinition() bool`

HasSignalsDefinition returns a boolean if a field has been set.

### SetSignalsDefinitionNil

`func (o *WorkflowsAvailableStepResponse) SetSignalsDefinitionNil(b bool)`

 SetSignalsDefinitionNil sets the value for SignalsDefinition to be an explicit nil

### UnsetSignalsDefinition
`func (o *WorkflowsAvailableStepResponse) UnsetSignalsDefinition()`

UnsetSignalsDefinition ensures that no value is present for SignalsDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


