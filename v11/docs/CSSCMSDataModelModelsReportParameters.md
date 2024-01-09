# CSSCMSDataModelModelsReportParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ParameterName** | Pointer to **NullableString** |  | [optional] 
**ParameterType** | Pointer to [**CSSCMSCoreEnumsReportParameterType**](CSSCMSCoreEnumsReportParameterType.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**ParameterVisibility** | Pointer to [**CSSCMSCoreEnumsReportParameterVisibility**](CSSCMSCoreEnumsReportParameterVisibility.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsReportParameters

`func NewCSSCMSDataModelModelsReportParameters() *CSSCMSDataModelModelsReportParameters`

NewCSSCMSDataModelModelsReportParameters instantiates a new CSSCMSDataModelModelsReportParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsReportParametersWithDefaults

`func NewCSSCMSDataModelModelsReportParametersWithDefaults() *CSSCMSDataModelModelsReportParameters`

NewCSSCMSDataModelModelsReportParametersWithDefaults instantiates a new CSSCMSDataModelModelsReportParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsReportParameters) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsReportParameters) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsReportParameters) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsReportParameters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameterName

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *CSSCMSDataModelModelsReportParameters) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *CSSCMSDataModelModelsReportParameters) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### SetParameterNameNil

`func (o *CSSCMSDataModelModelsReportParameters) SetParameterNameNil(b bool)`

 SetParameterNameNil sets the value for ParameterName to be an explicit nil

### UnsetParameterName
`func (o *CSSCMSDataModelModelsReportParameters) UnsetParameterName()`

UnsetParameterName ensures that no value is present for ParameterName, not even an explicit nil
### GetParameterType

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterType() CSSCMSCoreEnumsReportParameterType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterTypeOk() (*CSSCMSCoreEnumsReportParameterType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *CSSCMSDataModelModelsReportParameters) SetParameterType(v CSSCMSCoreEnumsReportParameterType)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *CSSCMSDataModelModelsReportParameters) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetDisplayName

`func (o *CSSCMSDataModelModelsReportParameters) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CSSCMSDataModelModelsReportParameters) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CSSCMSDataModelModelsReportParameters) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CSSCMSDataModelModelsReportParameters) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CSSCMSDataModelModelsReportParameters) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CSSCMSDataModelModelsReportParameters) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CSSCMSDataModelModelsReportParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsReportParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsReportParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsReportParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsReportParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsReportParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefaultValue

`func (o *CSSCMSDataModelModelsReportParameters) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsReportParameters) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsReportParameters) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsReportParameters) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsReportParameters) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsReportParameters) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDisplayOrder

`func (o *CSSCMSDataModelModelsReportParameters) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *CSSCMSDataModelModelsReportParameters) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *CSSCMSDataModelModelsReportParameters) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *CSSCMSDataModelModelsReportParameters) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetParameterVisibility

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterVisibility() CSSCMSCoreEnumsReportParameterVisibility`

GetParameterVisibility returns the ParameterVisibility field if non-nil, zero value otherwise.

### GetParameterVisibilityOk

`func (o *CSSCMSDataModelModelsReportParameters) GetParameterVisibilityOk() (*CSSCMSCoreEnumsReportParameterVisibility, bool)`

GetParameterVisibilityOk returns a tuple with the ParameterVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterVisibility

`func (o *CSSCMSDataModelModelsReportParameters) SetParameterVisibility(v CSSCMSCoreEnumsReportParameterVisibility)`

SetParameterVisibility sets ParameterVisibility field to given value.

### HasParameterVisibility

`func (o *CSSCMSDataModelModelsReportParameters) HasParameterVisibility() bool`

HasParameterVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


