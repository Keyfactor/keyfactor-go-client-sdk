# KeyfactorApiModelsWorkflowsAvailableStepQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the step. | [optional] 
**ExtensionName** | Pointer to **string** | The extension name of the step. | [optional] 
**SupportedWorkflowTypes** | Pointer to **[]string** | The workflow types which this step can be a part of. | [optional] 
**ConfigurationParametersDefinition** | Pointer to [**map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse**](KeyfactorApiModelsWorkflowsParameterDefinitionResponse.md) |  | [optional] 
**SignalsDefinition** | Pointer to [**map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse**](KeyfactorApiModelsWorkflowsSignalDefinitionResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponse

`func NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponse() *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse`

NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponse instantiates a new KeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults() *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse`

NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetSupportedWorkflowTypes

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypes() []string`

GetSupportedWorkflowTypes returns the SupportedWorkflowTypes field if non-nil, zero value otherwise.

### GetSupportedWorkflowTypesOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypesOk() (*[]string, bool)`

GetSupportedWorkflowTypesOk returns a tuple with the SupportedWorkflowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWorkflowTypes

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSupportedWorkflowTypes(v []string)`

SetSupportedWorkflowTypes sets SupportedWorkflowTypes field to given value.

### HasSupportedWorkflowTypes

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSupportedWorkflowTypes() bool`

HasSupportedWorkflowTypes returns a boolean if a field has been set.

### GetConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinition() map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse`

GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field if non-nil, zero value otherwise.

### GetConfigurationParametersDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse, bool)`

GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetConfigurationParametersDefinition(v map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse)`

SetConfigurationParametersDefinition sets ConfigurationParametersDefinition field to given value.

### HasConfigurationParametersDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasConfigurationParametersDefinition() bool`

HasConfigurationParametersDefinition returns a boolean if a field has been set.

### GetSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinition() map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse`

GetSignalsDefinition returns the SignalsDefinition field if non-nil, zero value otherwise.

### GetSignalsDefinitionOk

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse, bool)`

GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSignalsDefinition(v map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse)`

SetSignalsDefinition sets SignalsDefinition field to given value.

### HasSignalsDefinition

`func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSignalsDefinition() bool`

HasSignalsDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


