# KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ParameterType** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**ControlType** | Pointer to **int32** |  | [optional] 
**PotentialValues** | Pointer to **map[string]string** |  | [optional] 
**SupportTokenReplacement** | Pointer to **bool** |  | [optional] 
**DependsOn** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParameterType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetParameterType() int32`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetParameterTypeOk() (*int32, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetParameterType(v int32)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetRequired

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetControlType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetControlType() int32`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetControlTypeOk() (*int32, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetControlType(v int32)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetPotentialValues

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetPotentialValues() map[string]string`

GetPotentialValues returns the PotentialValues field if non-nil, zero value otherwise.

### GetPotentialValuesOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetPotentialValuesOk() (*map[string]string, bool)`

GetPotentialValuesOk returns a tuple with the PotentialValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialValues

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetPotentialValues(v map[string]string)`

SetPotentialValues sets PotentialValues field to given value.

### HasPotentialValues

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasPotentialValues() bool`

HasPotentialValues returns a boolean if a field has been set.

### SetPotentialValuesNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetPotentialValuesNil(b bool)`

 SetPotentialValuesNil sets the value for PotentialValues to be an explicit nil

### UnsetPotentialValues
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) UnsetPotentialValues()`

UnsetPotentialValues ensures that no value is present for PotentialValues, not even an explicit nil
### GetSupportTokenReplacement

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetSupportTokenReplacement() bool`

GetSupportTokenReplacement returns the SupportTokenReplacement field if non-nil, zero value otherwise.

### GetSupportTokenReplacementOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetSupportTokenReplacementOk() (*bool, bool)`

GetSupportTokenReplacementOk returns a tuple with the SupportTokenReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTokenReplacement

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetSupportTokenReplacement(v bool)`

SetSupportTokenReplacement sets SupportTokenReplacement field to given value.

### HasSupportTokenReplacement

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasSupportTokenReplacement() bool`

HasSupportTokenReplacement returns a boolean if a field has been set.

### GetDependsOn

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDependsOn() map[string]string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) GetDependsOnOk() (*map[string]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDependsOn(v map[string]string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


