# ModelsTemplateRetrievalResponseTemplatePolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **NullableInt64** |  | [optional] 
**RSAValidKeySizes** | Pointer to **[]int64** |  | [optional] 
**ECCValidCurves** | Pointer to **[]string** |  | [optional] 
**AllowKeyReuse** | Pointer to **bool** |  | [optional] 
**AllowWildcards** | Pointer to **bool** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsTemplateRetrievalResponseTemplatePolicyModel

`func NewModelsTemplateRetrievalResponseTemplatePolicyModel() *ModelsTemplateRetrievalResponseTemplatePolicyModel`

NewModelsTemplateRetrievalResponseTemplatePolicyModel instantiates a new ModelsTemplateRetrievalResponseTemplatePolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateRetrievalResponseTemplatePolicyModelWithDefaults

`func NewModelsTemplateRetrievalResponseTemplatePolicyModelWithDefaults() *ModelsTemplateRetrievalResponseTemplatePolicyModel`

NewModelsTemplateRetrievalResponseTemplatePolicyModelWithDefaults instantiates a new ModelsTemplateRetrievalResponseTemplatePolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRSAValidKeySizes

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetRSAValidKeySizes() []int64`

GetRSAValidKeySizes returns the RSAValidKeySizes field if non-nil, zero value otherwise.

### GetRSAValidKeySizesOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetRSAValidKeySizesOk() (*[]int64, bool)`

GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRSAValidKeySizes

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetRSAValidKeySizes(v []int64)`

SetRSAValidKeySizes sets RSAValidKeySizes field to given value.

### HasRSAValidKeySizes

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasRSAValidKeySizes() bool`

HasRSAValidKeySizes returns a boolean if a field has been set.

### GetECCValidCurves

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetECCValidCurves() []string`

GetECCValidCurves returns the ECCValidCurves field if non-nil, zero value otherwise.

### GetECCValidCurvesOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetECCValidCurvesOk() (*[]string, bool)`

GetECCValidCurvesOk returns a tuple with the ECCValidCurves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCValidCurves

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetECCValidCurves(v []string)`

SetECCValidCurves sets ECCValidCurves field to given value.

### HasECCValidCurves

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasECCValidCurves() bool`

HasECCValidCurves returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *ModelsTemplateRetrievalResponseTemplatePolicyModel) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


