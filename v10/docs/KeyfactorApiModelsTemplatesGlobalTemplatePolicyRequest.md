# KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RSAValidKeySizes** | **[]int64** | The allowed RSA key sizes. | 
**ECCValidCurves** | **[]string** | The allowed ECC curves. | 
**AllowKeyReuse** | **bool** | Whether or not keys can be reused. | 
**AllowWildcards** | **bool** | Whether or not wildcards can be used. | 
**RFCEnforcement** | **bool** | Whether or not RFC 2818 compliance should be enforced. | 

## Methods

### NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest

`func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest(rSAValidKeySizes []int64, eCCValidCurves []string, allowKeyReuse bool, allowWildcards bool, rFCEnforcement bool, ) *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequestWithDefaults

`func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequestWithDefaults() *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequestWithDefaults instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRSAValidKeySizes() []int64`

GetRSAValidKeySizes returns the RSAValidKeySizes field if non-nil, zero value otherwise.

### GetRSAValidKeySizesOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRSAValidKeySizesOk() (*[]int64, bool)`

GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetRSAValidKeySizes(v []int64)`

SetRSAValidKeySizes sets RSAValidKeySizes field to given value.


### GetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetECCValidCurves() []string`

GetECCValidCurves returns the ECCValidCurves field if non-nil, zero value otherwise.

### GetECCValidCurvesOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetECCValidCurvesOk() (*[]string, bool)`

GetECCValidCurvesOk returns a tuple with the ECCValidCurves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetECCValidCurves(v []string)`

SetECCValidCurves sets ECCValidCurves field to given value.


### GetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.


### GetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.


### GetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


