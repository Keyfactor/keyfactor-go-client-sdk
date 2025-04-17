# WorkflowsAvailableStepQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the step. | [optional] 
**ExtensionName** | Pointer to **NullableString** | The extension name of the step. | [optional] 
**SupportedWorkflowTypes** | Pointer to **[]string** | The workflow types which this step can be a part of. | [optional] 
**ConfigurationParametersDefinition** | Pointer to [**map[string]WorkflowsParameterDefinitionResponse**](WorkflowsParameterDefinitionResponse.md) |  | [optional] 
**SignalsDefinition** | Pointer to [**map[string]WorkflowsSignalDefinitionResponse**](WorkflowsSignalDefinitionResponse.md) |  | [optional] 

## Methods

### NewWorkflowsAvailableStepQueryResponse

`func NewWorkflowsAvailableStepQueryResponse() *WorkflowsAvailableStepQueryResponse`

NewWorkflowsAvailableStepQueryResponse instantiates a new WorkflowsAvailableStepQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsAvailableStepQueryResponseWithDefaults

`func NewWorkflowsAvailableStepQueryResponseWithDefaults() *WorkflowsAvailableStepQueryResponse`

NewWorkflowsAvailableStepQueryResponseWithDefaults instantiates a new WorkflowsAvailableStepQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsAvailableStepQueryResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsAvailableStepQueryResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsAvailableStepQueryResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsAvailableStepQueryResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsAvailableStepQueryResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsAvailableStepQueryResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExtensionName

`func (o *WorkflowsAvailableStepQueryResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *WorkflowsAvailableStepQueryResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *WorkflowsAvailableStepQueryResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *WorkflowsAvailableStepQueryResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### SetExtensionNameNil

`func (o *WorkflowsAvailableStepQueryResponse) SetExtensionNameNil(b bool)`

 SetExtensionNameNil sets the value for ExtensionName to be an explicit nil

### UnsetExtensionName
`func (o *WorkflowsAvailableStepQueryResponse) UnsetExtensionName()`

UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
### GetSupportedWorkflowTypes

`func (o *WorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypes() []string`

GetSupportedWorkflowTypes returns the SupportedWorkflowTypes field if non-nil, zero value otherwise.

### GetSupportedWorkflowTypesOk

`func (o *WorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypesOk() (*[]string, bool)`

GetSupportedWorkflowTypesOk returns a tuple with the SupportedWorkflowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWorkflowTypes

`func (o *WorkflowsAvailableStepQueryResponse) SetSupportedWorkflowTypes(v []string)`

SetSupportedWorkflowTypes sets SupportedWorkflowTypes field to given value.

### HasSupportedWorkflowTypes

`func (o *WorkflowsAvailableStepQueryResponse) HasSupportedWorkflowTypes() bool`

HasSupportedWorkflowTypes returns a boolean if a field has been set.

### SetSupportedWorkflowTypesNil

`func (o *WorkflowsAvailableStepQueryResponse) SetSupportedWorkflowTypesNil(b bool)`

 SetSupportedWorkflowTypesNil sets the value for SupportedWorkflowTypes to be an explicit nil

### UnsetSupportedWorkflowTypes
`func (o *WorkflowsAvailableStepQueryResponse) UnsetSupportedWorkflowTypes()`

UnsetSupportedWorkflowTypes ensures that no value is present for SupportedWorkflowTypes, not even an explicit nil
### GetConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinition() map[string]WorkflowsParameterDefinitionResponse`

GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field if non-nil, zero value otherwise.

### GetConfigurationParametersDefinitionOk

`func (o *WorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinitionOk() (*map[string]WorkflowsParameterDefinitionResponse, bool)`

GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepQueryResponse) SetConfigurationParametersDefinition(v map[string]WorkflowsParameterDefinitionResponse)`

SetConfigurationParametersDefinition sets ConfigurationParametersDefinition field to given value.

### HasConfigurationParametersDefinition

`func (o *WorkflowsAvailableStepQueryResponse) HasConfigurationParametersDefinition() bool`

HasConfigurationParametersDefinition returns a boolean if a field has been set.

### SetConfigurationParametersDefinitionNil

`func (o *WorkflowsAvailableStepQueryResponse) SetConfigurationParametersDefinitionNil(b bool)`

 SetConfigurationParametersDefinitionNil sets the value for ConfigurationParametersDefinition to be an explicit nil

### UnsetConfigurationParametersDefinition
`func (o *WorkflowsAvailableStepQueryResponse) UnsetConfigurationParametersDefinition()`

UnsetConfigurationParametersDefinition ensures that no value is present for ConfigurationParametersDefinition, not even an explicit nil
### GetSignalsDefinition

`func (o *WorkflowsAvailableStepQueryResponse) GetSignalsDefinition() map[string]WorkflowsSignalDefinitionResponse`

GetSignalsDefinition returns the SignalsDefinition field if non-nil, zero value otherwise.

### GetSignalsDefinitionOk

`func (o *WorkflowsAvailableStepQueryResponse) GetSignalsDefinitionOk() (*map[string]WorkflowsSignalDefinitionResponse, bool)`

GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDefinition

`func (o *WorkflowsAvailableStepQueryResponse) SetSignalsDefinition(v map[string]WorkflowsSignalDefinitionResponse)`

SetSignalsDefinition sets SignalsDefinition field to given value.

### HasSignalsDefinition

`func (o *WorkflowsAvailableStepQueryResponse) HasSignalsDefinition() bool`

HasSignalsDefinition returns a boolean if a field has been set.

### SetSignalsDefinitionNil

`func (o *WorkflowsAvailableStepQueryResponse) SetSignalsDefinitionNil(b bool)`

 SetSignalsDefinitionNil sets the value for SignalsDefinition to be an explicit nil

### UnsetSignalsDefinition
`func (o *WorkflowsAvailableStepQueryResponse) UnsetSignalsDefinition()`

UnsetSignalsDefinition ensures that no value is present for SignalsDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


