# KeyfactorApiModelsWorkflowsParameterDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**ParameterType** | Pointer to **int64** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**ControlType** | Pointer to **int64** |  | [optional] 
**PotentialValues** | Pointer to **map[string]string** |  | [optional] 
**SupportTokenReplacement** | Pointer to **bool** |  | [optional] 
**DependsOn** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsParameterDefinitionResponse

`func NewKeyfactorApiModelsWorkflowsParameterDefinitionResponse() *KeyfactorApiModelsWorkflowsParameterDefinitionResponse`

NewKeyfactorApiModelsWorkflowsParameterDefinitionResponse instantiates a new KeyfactorApiModelsWorkflowsParameterDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults() *KeyfactorApiModelsWorkflowsParameterDefinitionResponse`

NewKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsParameterDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParameterType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetParameterType() int64`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetParameterTypeOk() (*int64, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetParameterType(v int64)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetRequired

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetControlType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetControlType() int64`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetControlTypeOk() (*int64, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetControlType(v int64)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetPotentialValues

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetPotentialValues() map[string]string`

GetPotentialValues returns the PotentialValues field if non-nil, zero value otherwise.

### GetPotentialValuesOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetPotentialValuesOk() (*map[string]string, bool)`

GetPotentialValuesOk returns a tuple with the PotentialValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialValues

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetPotentialValues(v map[string]string)`

SetPotentialValues sets PotentialValues field to given value.

### HasPotentialValues

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasPotentialValues() bool`

HasPotentialValues returns a boolean if a field has been set.

### GetSupportTokenReplacement

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetSupportTokenReplacement() bool`

GetSupportTokenReplacement returns the SupportTokenReplacement field if non-nil, zero value otherwise.

### GetSupportTokenReplacementOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetSupportTokenReplacementOk() (*bool, bool)`

GetSupportTokenReplacementOk returns a tuple with the SupportTokenReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTokenReplacement

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetSupportTokenReplacement(v bool)`

SetSupportTokenReplacement sets SupportTokenReplacement field to given value.

### HasSupportTokenReplacement

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasSupportTokenReplacement() bool`

HasSupportTokenReplacement returns a boolean if a field has been set.

### GetDependsOn

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDependsOn() map[string]string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDependsOnOk() (*map[string]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDependsOn(v map[string]string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *KeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


