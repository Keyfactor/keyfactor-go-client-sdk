# WorkflowsDefinitionStepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]WorkflowsSignalConfigurationRequest**](WorkflowsSignalConfigurationRequest.md) |  | [optional] 
**Conditions** | Pointer to [**[]WorkflowsConditionConfigurationRequest**](WorkflowsConditionConfigurationRequest.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowsDefinitionStepRequest

`func NewWorkflowsDefinitionStepRequest() *WorkflowsDefinitionStepRequest`

NewWorkflowsDefinitionStepRequest instantiates a new WorkflowsDefinitionStepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionStepRequestWithDefaults

`func NewWorkflowsDefinitionStepRequestWithDefaults() *WorkflowsDefinitionStepRequest`

NewWorkflowsDefinitionStepRequestWithDefaults instantiates a new WorkflowsDefinitionStepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *WorkflowsDefinitionStepRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *WorkflowsDefinitionStepRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *WorkflowsDefinitionStepRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *WorkflowsDefinitionStepRequest) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *WorkflowsDefinitionStepRequest) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *WorkflowsDefinitionStepRequest) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetUniqueName

`func (o *WorkflowsDefinitionStepRequest) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowsDefinitionStepRequest) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowsDefinitionStepRequest) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowsDefinitionStepRequest) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowsDefinitionStepRequest) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowsDefinitionStepRequest) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetDisplayName

`func (o *WorkflowsDefinitionStepRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionStepRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionStepRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsDefinitionStepRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsDefinitionStepRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsDefinitionStepRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnabled

`func (o *WorkflowsDefinitionStepRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionStepRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionStepRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionStepRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *WorkflowsDefinitionStepRequest) GetConfigurationParameters() map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *WorkflowsDefinitionStepRequest) GetConfigurationParametersOk() (*map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *WorkflowsDefinitionStepRequest) SetConfigurationParameters(v map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *WorkflowsDefinitionStepRequest) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### SetConfigurationParametersNil

`func (o *WorkflowsDefinitionStepRequest) SetConfigurationParametersNil(b bool)`

 SetConfigurationParametersNil sets the value for ConfigurationParameters to be an explicit nil

### UnsetConfigurationParameters
`func (o *WorkflowsDefinitionStepRequest) UnsetConfigurationParameters()`

UnsetConfigurationParameters ensures that no value is present for ConfigurationParameters, not even an explicit nil
### GetSignals

`func (o *WorkflowsDefinitionStepRequest) GetSignals() []WorkflowsSignalConfigurationRequest`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WorkflowsDefinitionStepRequest) GetSignalsOk() (*[]WorkflowsSignalConfigurationRequest, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WorkflowsDefinitionStepRequest) SetSignals(v []WorkflowsSignalConfigurationRequest)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *WorkflowsDefinitionStepRequest) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *WorkflowsDefinitionStepRequest) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *WorkflowsDefinitionStepRequest) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetConditions

`func (o *WorkflowsDefinitionStepRequest) GetConditions() []WorkflowsConditionConfigurationRequest`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowsDefinitionStepRequest) GetConditionsOk() (*[]WorkflowsConditionConfigurationRequest, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowsDefinitionStepRequest) SetConditions(v []WorkflowsConditionConfigurationRequest)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowsDefinitionStepRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *WorkflowsDefinitionStepRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *WorkflowsDefinitionStepRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetOutputs

`func (o *WorkflowsDefinitionStepRequest) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowsDefinitionStepRequest) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowsDefinitionStepRequest) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowsDefinitionStepRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowsDefinitionStepRequest) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowsDefinitionStepRequest) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


