# KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**ExtensionName** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse.md) |  | [optional] 
**Conditions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationResponse.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConfigurationParameters() map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConfigurationParametersOk() (*map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConfigurationParameters(v map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### SetConfigurationParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConfigurationParametersNil(b bool)`

 SetConfigurationParametersNil sets the value for ConfigurationParameters to be an explicit nil

### UnsetConfigurationParameters
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetConfigurationParameters()`

UnsetConfigurationParameters ensures that no value is present for ConfigurationParameters, not even an explicit nil
### GetSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetSignals() []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetSignalsOk() (*[]KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetSignals(v []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConditions() []KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConditionsOk() (*[]KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConditions(v []KeyfactorWebKeyfactorApiModelsWorkflowsConditionConfigurationResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


