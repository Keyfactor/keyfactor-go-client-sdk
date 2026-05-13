# WorkflowsDefinitionStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**ExtensionName** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]WorkflowsDefinitionStepSignalResponse**](WorkflowsDefinitionStepSignalResponse.md) |  | [optional] 
**Conditions** | Pointer to [**[]WorkflowsConditionConfigurationResponse**](WorkflowsConditionConfigurationResponse.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowsDefinitionStepResponse

`func NewWorkflowsDefinitionStepResponse() *WorkflowsDefinitionStepResponse`

NewWorkflowsDefinitionStepResponse instantiates a new WorkflowsDefinitionStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionStepResponseWithDefaults

`func NewWorkflowsDefinitionStepResponseWithDefaults() *WorkflowsDefinitionStepResponse`

NewWorkflowsDefinitionStepResponseWithDefaults instantiates a new WorkflowsDefinitionStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsDefinitionStepResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsDefinitionStepResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsDefinitionStepResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsDefinitionStepResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *WorkflowsDefinitionStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsDefinitionStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsDefinitionStepResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsDefinitionStepResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUniqueName

`func (o *WorkflowsDefinitionStepResponse) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowsDefinitionStepResponse) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowsDefinitionStepResponse) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowsDefinitionStepResponse) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowsDefinitionStepResponse) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowsDefinitionStepResponse) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetExtensionName

`func (o *WorkflowsDefinitionStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *WorkflowsDefinitionStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *WorkflowsDefinitionStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *WorkflowsDefinitionStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *WorkflowsDefinitionStepResponse) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *WorkflowsDefinitionStepResponse) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetEnabled

`func (o *WorkflowsDefinitionStepResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionStepResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionStepResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionStepResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *WorkflowsDefinitionStepResponse) GetConfigurationParameters() map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *WorkflowsDefinitionStepResponse) GetConfigurationParametersOk() (*map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *WorkflowsDefinitionStepResponse) SetConfigurationParameters(v map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *WorkflowsDefinitionStepResponse) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### SetConfigurationParametersNil

`func (o *WorkflowsDefinitionStepResponse) SetConfigurationParametersNil(b bool)`

 SetConfigurationParametersNil sets the value for ConfigurationParameters to be an explicit nil

### UnsetConfigurationParameters
`func (o *WorkflowsDefinitionStepResponse) UnsetConfigurationParameters()`

UnsetConfigurationParameters ensures that no value is present for ConfigurationParameters, not even an explicit nil
### GetSignals

`func (o *WorkflowsDefinitionStepResponse) GetSignals() []WorkflowsDefinitionStepSignalResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WorkflowsDefinitionStepResponse) GetSignalsOk() (*[]WorkflowsDefinitionStepSignalResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WorkflowsDefinitionStepResponse) SetSignals(v []WorkflowsDefinitionStepSignalResponse)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *WorkflowsDefinitionStepResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *WorkflowsDefinitionStepResponse) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *WorkflowsDefinitionStepResponse) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetConditions

`func (o *WorkflowsDefinitionStepResponse) GetConditions() []WorkflowsConditionConfigurationResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowsDefinitionStepResponse) GetConditionsOk() (*[]WorkflowsConditionConfigurationResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowsDefinitionStepResponse) SetConditions(v []WorkflowsConditionConfigurationResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowsDefinitionStepResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *WorkflowsDefinitionStepResponse) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *WorkflowsDefinitionStepResponse) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetOutputs

`func (o *WorkflowsDefinitionStepResponse) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowsDefinitionStepResponse) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowsDefinitionStepResponse) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowsDefinitionStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowsDefinitionStepResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowsDefinitionStepResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


