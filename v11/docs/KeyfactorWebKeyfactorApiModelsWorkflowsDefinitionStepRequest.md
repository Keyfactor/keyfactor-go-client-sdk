# KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsWorkflowsSignalConfigurationRequest**](KeyfactorWebKeyfactorApiModelsWorkflowsSignalConfigurationRequest.md) |  | [optional] 
**Conditions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationRequest**](KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationRequest.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest`

NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest`

NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConfigurationParameters() map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConfigurationParametersOk() (*map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConfigurationParameters(v map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### SetConfigurationParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConfigurationParametersNil(b bool)`

 SetConfigurationParametersNil sets the value for ConfigurationParameters to be an explicit nil

### UnsetConfigurationParameters
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetConfigurationParameters()`

UnsetConfigurationParameters ensures that no value is present for ConfigurationParameters, not even an explicit nil
### GetSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetSignals() []KeyfactorWebKeyfactorApiModelsWorkflowsSignalConfigurationRequest`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetSignalsOk() (*[]KeyfactorWebKeyfactorApiModelsWorkflowsSignalConfigurationRequest, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetSignals(v []KeyfactorWebKeyfactorApiModelsWorkflowsSignalConfigurationRequest)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConditions() []KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationRequest`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConditionsOk() (*[]KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationRequest, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConditions(v []KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationRequest)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


