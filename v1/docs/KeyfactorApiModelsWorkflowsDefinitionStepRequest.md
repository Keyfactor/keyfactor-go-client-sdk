# KeyfactorApiModelsWorkflowsDefinitionStepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | Pointer to **string** |  | [optional] 
**UniqueName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]KeyfactorApiModelsWorkflowsSignalConfigurationRequest**](KeyfactorApiModelsWorkflowsSignalConfigurationRequest.md) |  | [optional] 
**Conditions** | Pointer to [**[]KeyfactorApiModelsWorkflowsConditionConfigurationRequest**](KeyfactorApiModelsWorkflowsConditionConfigurationRequest.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsDefinitionStepRequest

`func NewKeyfactorApiModelsWorkflowsDefinitionStepRequest() *KeyfactorApiModelsWorkflowsDefinitionStepRequest`

NewKeyfactorApiModelsWorkflowsDefinitionStepRequest instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults

`func NewKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults() *KeyfactorApiModelsWorkflowsDefinitionStepRequest`

NewKeyfactorApiModelsWorkflowsDefinitionStepRequestWithDefaults instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConfigurationParameters() map[string]map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConfigurationParametersOk() (*map[string]map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConfigurationParameters(v map[string]map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetSignals() []KeyfactorApiModelsWorkflowsSignalConfigurationRequest`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetSignalsOk() (*[]KeyfactorApiModelsWorkflowsSignalConfigurationRequest, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetSignals(v []KeyfactorApiModelsWorkflowsSignalConfigurationRequest)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConditions() []KeyfactorApiModelsWorkflowsConditionConfigurationRequest`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetConditionsOk() (*[]KeyfactorApiModelsWorkflowsConditionConfigurationRequest, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetConditions(v []KeyfactorApiModelsWorkflowsConditionConfigurationRequest)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


