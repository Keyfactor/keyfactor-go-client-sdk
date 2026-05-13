# WorkflowsParameterDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ParameterType** | Pointer to [**KeyfactorWorkflowsDataType**](KeyfactorWorkflowsDataType.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**ControlType** | Pointer to [**KeyfactorWorkflowsInputControlType**](KeyfactorWorkflowsInputControlType.md) |  | [optional] 
**PotentialValues** | Pointer to **map[string]string** |  | [optional] 
**SupportTokenReplacement** | Pointer to **bool** |  | [optional] 
**DependsOn** | Pointer to **map[string]string** |  | [optional] 
**ToolTipContents** | Pointer to **NullableString** |  | [optional] 
**QueryEndpoint** | Pointer to **NullableString** |  | [optional] 
**QueryFieldName** | Pointer to **NullableString** |  | [optional] 
**QueryGroupBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowsParameterDefinitionResponse

`func NewWorkflowsParameterDefinitionResponse() *WorkflowsParameterDefinitionResponse`

NewWorkflowsParameterDefinitionResponse instantiates a new WorkflowsParameterDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsParameterDefinitionResponseWithDefaults

`func NewWorkflowsParameterDefinitionResponseWithDefaults() *WorkflowsParameterDefinitionResponse`

NewWorkflowsParameterDefinitionResponseWithDefaults instantiates a new WorkflowsParameterDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsParameterDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsParameterDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsParameterDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsParameterDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsParameterDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsParameterDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParameterType

`func (o *WorkflowsParameterDefinitionResponse) GetParameterType() KeyfactorWorkflowsDataType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *WorkflowsParameterDefinitionResponse) GetParameterTypeOk() (*KeyfactorWorkflowsDataType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *WorkflowsParameterDefinitionResponse) SetParameterType(v KeyfactorWorkflowsDataType)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *WorkflowsParameterDefinitionResponse) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetRequired

`func (o *WorkflowsParameterDefinitionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *WorkflowsParameterDefinitionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *WorkflowsParameterDefinitionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *WorkflowsParameterDefinitionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *WorkflowsParameterDefinitionResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *WorkflowsParameterDefinitionResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *WorkflowsParameterDefinitionResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *WorkflowsParameterDefinitionResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *WorkflowsParameterDefinitionResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *WorkflowsParameterDefinitionResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetControlType

`func (o *WorkflowsParameterDefinitionResponse) GetControlType() KeyfactorWorkflowsInputControlType`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *WorkflowsParameterDefinitionResponse) GetControlTypeOk() (*KeyfactorWorkflowsInputControlType, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *WorkflowsParameterDefinitionResponse) SetControlType(v KeyfactorWorkflowsInputControlType)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *WorkflowsParameterDefinitionResponse) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetPotentialValues

`func (o *WorkflowsParameterDefinitionResponse) GetPotentialValues() map[string]string`

GetPotentialValues returns the PotentialValues field if non-nil, zero value otherwise.

### GetPotentialValuesOk

`func (o *WorkflowsParameterDefinitionResponse) GetPotentialValuesOk() (*map[string]string, bool)`

GetPotentialValuesOk returns a tuple with the PotentialValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialValues

`func (o *WorkflowsParameterDefinitionResponse) SetPotentialValues(v map[string]string)`

SetPotentialValues sets PotentialValues field to given value.

### HasPotentialValues

`func (o *WorkflowsParameterDefinitionResponse) HasPotentialValues() bool`

HasPotentialValues returns a boolean if a field has been set.

### SetPotentialValuesNil

`func (o *WorkflowsParameterDefinitionResponse) SetPotentialValuesNil(b bool)`

 SetPotentialValuesNil sets the value for PotentialValues to be an explicit nil

### UnsetPotentialValues
`func (o *WorkflowsParameterDefinitionResponse) UnsetPotentialValues()`

UnsetPotentialValues ensures that no value is present for PotentialValues, not even an explicit nil
### GetSupportTokenReplacement

`func (o *WorkflowsParameterDefinitionResponse) GetSupportTokenReplacement() bool`

GetSupportTokenReplacement returns the SupportTokenReplacement field if non-nil, zero value otherwise.

### GetSupportTokenReplacementOk

`func (o *WorkflowsParameterDefinitionResponse) GetSupportTokenReplacementOk() (*bool, bool)`

GetSupportTokenReplacementOk returns a tuple with the SupportTokenReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTokenReplacement

`func (o *WorkflowsParameterDefinitionResponse) SetSupportTokenReplacement(v bool)`

SetSupportTokenReplacement sets SupportTokenReplacement field to given value.

### HasSupportTokenReplacement

`func (o *WorkflowsParameterDefinitionResponse) HasSupportTokenReplacement() bool`

HasSupportTokenReplacement returns a boolean if a field has been set.

### GetDependsOn

`func (o *WorkflowsParameterDefinitionResponse) GetDependsOn() map[string]string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *WorkflowsParameterDefinitionResponse) GetDependsOnOk() (*map[string]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *WorkflowsParameterDefinitionResponse) SetDependsOn(v map[string]string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *WorkflowsParameterDefinitionResponse) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *WorkflowsParameterDefinitionResponse) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *WorkflowsParameterDefinitionResponse) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetToolTipContents

`func (o *WorkflowsParameterDefinitionResponse) GetToolTipContents() string`

GetToolTipContents returns the ToolTipContents field if non-nil, zero value otherwise.

### GetToolTipContentsOk

`func (o *WorkflowsParameterDefinitionResponse) GetToolTipContentsOk() (*string, bool)`

GetToolTipContentsOk returns a tuple with the ToolTipContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTipContents

`func (o *WorkflowsParameterDefinitionResponse) SetToolTipContents(v string)`

SetToolTipContents sets ToolTipContents field to given value.

### HasToolTipContents

`func (o *WorkflowsParameterDefinitionResponse) HasToolTipContents() bool`

HasToolTipContents returns a boolean if a field has been set.

### SetToolTipContentsNil

`func (o *WorkflowsParameterDefinitionResponse) SetToolTipContentsNil(b bool)`

 SetToolTipContentsNil sets the value for ToolTipContents to be an explicit nil

### UnsetToolTipContents
`func (o *WorkflowsParameterDefinitionResponse) UnsetToolTipContents()`

UnsetToolTipContents ensures that no value is present for ToolTipContents, not even an explicit nil
### GetQueryEndpoint

`func (o *WorkflowsParameterDefinitionResponse) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *WorkflowsParameterDefinitionResponse) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *WorkflowsParameterDefinitionResponse) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *WorkflowsParameterDefinitionResponse) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### SetQueryEndpointNil

`func (o *WorkflowsParameterDefinitionResponse) SetQueryEndpointNil(b bool)`

 SetQueryEndpointNil sets the value for QueryEndpoint to be an explicit nil

### UnsetQueryEndpoint
`func (o *WorkflowsParameterDefinitionResponse) UnsetQueryEndpoint()`

UnsetQueryEndpoint ensures that no value is present for QueryEndpoint, not even an explicit nil
### GetQueryFieldName

`func (o *WorkflowsParameterDefinitionResponse) GetQueryFieldName() string`

GetQueryFieldName returns the QueryFieldName field if non-nil, zero value otherwise.

### GetQueryFieldNameOk

`func (o *WorkflowsParameterDefinitionResponse) GetQueryFieldNameOk() (*string, bool)`

GetQueryFieldNameOk returns a tuple with the QueryFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFieldName

`func (o *WorkflowsParameterDefinitionResponse) SetQueryFieldName(v string)`

SetQueryFieldName sets QueryFieldName field to given value.

### HasQueryFieldName

`func (o *WorkflowsParameterDefinitionResponse) HasQueryFieldName() bool`

HasQueryFieldName returns a boolean if a field has been set.

### SetQueryFieldNameNil

`func (o *WorkflowsParameterDefinitionResponse) SetQueryFieldNameNil(b bool)`

 SetQueryFieldNameNil sets the value for QueryFieldName to be an explicit nil

### UnsetQueryFieldName
`func (o *WorkflowsParameterDefinitionResponse) UnsetQueryFieldName()`

UnsetQueryFieldName ensures that no value is present for QueryFieldName, not even an explicit nil
### GetQueryGroupBy

`func (o *WorkflowsParameterDefinitionResponse) GetQueryGroupBy() string`

GetQueryGroupBy returns the QueryGroupBy field if non-nil, zero value otherwise.

### GetQueryGroupByOk

`func (o *WorkflowsParameterDefinitionResponse) GetQueryGroupByOk() (*string, bool)`

GetQueryGroupByOk returns a tuple with the QueryGroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryGroupBy

`func (o *WorkflowsParameterDefinitionResponse) SetQueryGroupBy(v string)`

SetQueryGroupBy sets QueryGroupBy field to given value.

### HasQueryGroupBy

`func (o *WorkflowsParameterDefinitionResponse) HasQueryGroupBy() bool`

HasQueryGroupBy returns a boolean if a field has been set.

### SetQueryGroupByNil

`func (o *WorkflowsParameterDefinitionResponse) SetQueryGroupByNil(b bool)`

 SetQueryGroupByNil sets the value for QueryGroupBy to be an explicit nil

### UnsetQueryGroupBy
`func (o *WorkflowsParameterDefinitionResponse) UnsetQueryGroupBy()`

UnsetQueryGroupBy ensures that no value is present for QueryGroupBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


