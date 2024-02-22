# KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RSAValidKeySizes** | Pointer to **[]int64** | The allowed RSA key sizes. | [optional] 
**ECCValidCurves** | Pointer to **[]string** | The allowed ECC curves. | [optional] 
**AllowKeyReuse** | Pointer to **bool** | Whether or not keys can be reused. | [optional] 
**AllowWildcards** | Pointer to **bool** | Whether or not wildcards can be used. | [optional] 
**RFCEnforcement** | Pointer to **bool** | Whether or not RFC 2818 compliance should be enforced. | [optional] 

## Methods

### NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse

`func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse() *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse`

NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponseWithDefaults

`func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponseWithDefaults() *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse`

NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyResponseWithDefaults instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetRSAValidKeySizes() []int64`

GetRSAValidKeySizes returns the RSAValidKeySizes field if non-nil, zero value otherwise.

### GetRSAValidKeySizesOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetRSAValidKeySizesOk() (*[]int64, bool)`

GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) SetRSAValidKeySizes(v []int64)`

SetRSAValidKeySizes sets RSAValidKeySizes field to given value.

### HasRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) HasRSAValidKeySizes() bool`

HasRSAValidKeySizes returns a boolean if a field has been set.

### GetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetECCValidCurves() []string`

GetECCValidCurves returns the ECCValidCurves field if non-nil, zero value otherwise.

### GetECCValidCurvesOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetECCValidCurvesOk() (*[]string, bool)`

GetECCValidCurvesOk returns a tuple with the ECCValidCurves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) SetECCValidCurves(v []string)`

SetECCValidCurves sets ECCValidCurves field to given value.

### HasECCValidCurves

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) HasECCValidCurves() bool`

HasECCValidCurves returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


