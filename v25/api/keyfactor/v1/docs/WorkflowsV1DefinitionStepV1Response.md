# WorkflowsV1DefinitionStepV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**ExtensionName** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]WorkflowsV1DefinitionStepSignalV1Response**](WorkflowsV1DefinitionStepSignalV1Response.md) |  | [optional] 
**Conditions** | Pointer to [**[]WorkflowsConditionConfigurationResponse**](WorkflowsConditionConfigurationResponse.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowsV1DefinitionStepV1Response

`func NewWorkflowsV1DefinitionStepV1Response() *WorkflowsV1DefinitionStepV1Response`

NewWorkflowsV1DefinitionStepV1Response instantiates a new WorkflowsV1DefinitionStepV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsV1DefinitionStepV1ResponseWithDefaults

`func NewWorkflowsV1DefinitionStepV1ResponseWithDefaults() *WorkflowsV1DefinitionStepV1Response`

NewWorkflowsV1DefinitionStepV1ResponseWithDefaults instantiates a new WorkflowsV1DefinitionStepV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsV1DefinitionStepV1Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsV1DefinitionStepV1Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsV1DefinitionStepV1Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *WorkflowsV1DefinitionStepV1Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsV1DefinitionStepV1Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsV1DefinitionStepV1Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUniqueName

`func (o *WorkflowsV1DefinitionStepV1Response) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowsV1DefinitionStepV1Response) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowsV1DefinitionStepV1Response) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetExtensionName

`func (o *WorkflowsV1DefinitionStepV1Response) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *WorkflowsV1DefinitionStepV1Response) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *WorkflowsV1DefinitionStepV1Response) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetEnabled

`func (o *WorkflowsV1DefinitionStepV1Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsV1DefinitionStepV1Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsV1DefinitionStepV1Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *WorkflowsV1DefinitionStepV1Response) GetConfigurationParameters() map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetConfigurationParametersOk() (*map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *WorkflowsV1DefinitionStepV1Response) SetConfigurationParameters(v map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *WorkflowsV1DefinitionStepV1Response) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### SetConfigurationParametersNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetConfigurationParametersNil(b bool)`

 SetConfigurationParametersNil sets the value for ConfigurationParameters to be an explicit nil

### UnsetConfigurationParameters
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetConfigurationParameters()`

UnsetConfigurationParameters ensures that no value is present for ConfigurationParameters, not even an explicit nil
### GetSignals

`func (o *WorkflowsV1DefinitionStepV1Response) GetSignals() []WorkflowsV1DefinitionStepSignalV1Response`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetSignalsOk() (*[]WorkflowsV1DefinitionStepSignalV1Response, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WorkflowsV1DefinitionStepV1Response) SetSignals(v []WorkflowsV1DefinitionStepSignalV1Response)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *WorkflowsV1DefinitionStepV1Response) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### SetSignalsNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetSignalsNil(b bool)`

 SetSignalsNil sets the value for Signals to be an explicit nil

### UnsetSignals
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetSignals()`

UnsetSignals ensures that no value is present for Signals, not even an explicit nil
### GetConditions

`func (o *WorkflowsV1DefinitionStepV1Response) GetConditions() []WorkflowsConditionConfigurationResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetConditionsOk() (*[]WorkflowsConditionConfigurationResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowsV1DefinitionStepV1Response) SetConditions(v []WorkflowsConditionConfigurationResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowsV1DefinitionStepV1Response) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetOutputs

`func (o *WorkflowsV1DefinitionStepV1Response) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowsV1DefinitionStepV1Response) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowsV1DefinitionStepV1Response) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowsV1DefinitionStepV1Response) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowsV1DefinitionStepV1Response) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowsV1DefinitionStepV1Response) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


