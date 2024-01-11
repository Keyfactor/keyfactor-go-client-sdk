# KeyfactorApiModelsWorkflowsDefinitionStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**UniqueName** | Pointer to **string** |  | [optional] 
**ExtensionName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigurationParameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Signals** | Pointer to [**[]KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse**](KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse.md) |  | [optional] 
**Conditions** | Pointer to [**[]KeyfactorApiModelsWorkflowsConditionConfigurationResponse**](KeyfactorApiModelsWorkflowsConditionConfigurationResponse.md) |  | [optional] 
**Outputs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsDefinitionStepResponse

`func NewKeyfactorApiModelsWorkflowsDefinitionStepResponse() *KeyfactorApiModelsWorkflowsDefinitionStepResponse`

NewKeyfactorApiModelsWorkflowsDefinitionStepResponse instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults() *KeyfactorApiModelsWorkflowsDefinitionStepResponse`

NewKeyfactorApiModelsWorkflowsDefinitionStepResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConfigurationParameters() map[string]map[string]interface{}`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConfigurationParametersOk() (*map[string]map[string]interface{}, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConfigurationParameters(v map[string]map[string]interface{})`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetSignals() []KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetSignalsOk() (*[]KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetSignals(v []KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConditions() []KeyfactorApiModelsWorkflowsConditionConfigurationResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetConditionsOk() (*[]KeyfactorApiModelsWorkflowsConditionConfigurationResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetConditions(v []KeyfactorApiModelsWorkflowsConditionConfigurationResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetOutputs() map[string]string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) GetOutputsOk() (*map[string]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) SetOutputs(v map[string]string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *KeyfactorApiModelsWorkflowsDefinitionStepResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


