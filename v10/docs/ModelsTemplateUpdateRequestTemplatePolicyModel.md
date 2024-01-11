# ModelsTemplateUpdateRequestTemplatePolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**RSAValidKeySizes** | Pointer to **[]int32** |  | [optional] 
**ECCValidCurves** | Pointer to **[]string** |  | [optional] 
**AllowKeyReuse** | Pointer to **bool** |  | [optional] 
**AllowWildcards** | Pointer to **bool** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsTemplateUpdateRequestTemplatePolicyModel

`func NewModelsTemplateUpdateRequestTemplatePolicyModel() *ModelsTemplateUpdateRequestTemplatePolicyModel`

NewModelsTemplateUpdateRequestTemplatePolicyModel instantiates a new ModelsTemplateUpdateRequestTemplatePolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateUpdateRequestTemplatePolicyModelWithDefaults

`func NewModelsTemplateUpdateRequestTemplatePolicyModelWithDefaults() *ModelsTemplateUpdateRequestTemplatePolicyModel`

NewModelsTemplateUpdateRequestTemplatePolicyModelWithDefaults instantiates a new ModelsTemplateUpdateRequestTemplatePolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRSAValidKeySizes

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetRSAValidKeySizes() []int32`

GetRSAValidKeySizes returns the RSAValidKeySizes field if non-nil, zero value otherwise.

### GetRSAValidKeySizesOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetRSAValidKeySizesOk() (*[]int32, bool)`

GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRSAValidKeySizes

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetRSAValidKeySizes(v []int32)`

SetRSAValidKeySizes sets RSAValidKeySizes field to given value.

### HasRSAValidKeySizes

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasRSAValidKeySizes() bool`

HasRSAValidKeySizes returns a boolean if a field has been set.

### GetECCValidCurves

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetECCValidCurves() []string`

GetECCValidCurves returns the ECCValidCurves field if non-nil, zero value otherwise.

### GetECCValidCurvesOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetECCValidCurvesOk() (*[]string, bool)`

GetECCValidCurvesOk returns a tuple with the ECCValidCurves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCValidCurves

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetECCValidCurves(v []string)`

SetECCValidCurves sets ECCValidCurves field to given value.

### HasECCValidCurves

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasECCValidCurves() bool`

HasECCValidCurves returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *ModelsTemplateUpdateRequestTemplatePolicyModel) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


