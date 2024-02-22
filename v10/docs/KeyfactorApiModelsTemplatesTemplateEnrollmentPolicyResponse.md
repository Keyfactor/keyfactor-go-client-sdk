# KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RSAValidKeySizes** | Pointer to **[]int64** | The allowed RSA key sizes. | [optional] 
**ECCValidCurves** | Pointer to **[]string** | The allowed ECC curves. | [optional] 
**AllowKeyReuse** | Pointer to **bool** | Whether or not keys can be reused. | [optional] 
**AllowWildcards** | Pointer to **bool** | Whether or not wildcards can be used. | [optional] 
**RFCEnforcement** | Pointer to **bool** | Whether or not RFC 2818 compliance should be enforced. | [optional] 

## Methods

### NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse

`func NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse() *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse`

NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse instantiates a new KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponseWithDefaults

`func NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponseWithDefaults() *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse`

NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponseWithDefaults instantiates a new KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRSAValidKeySizes() []int64`

GetRSAValidKeySizes returns the RSAValidKeySizes field if non-nil, zero value otherwise.

### GetRSAValidKeySizesOk

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRSAValidKeySizesOk() (*[]int64, bool)`

GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetRSAValidKeySizes(v []int64)`

SetRSAValidKeySizes sets RSAValidKeySizes field to given value.

### HasRSAValidKeySizes

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasRSAValidKeySizes() bool`

HasRSAValidKeySizes returns a boolean if a field has been set.

### GetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetECCValidCurves() []string`

GetECCValidCurves returns the ECCValidCurves field if non-nil, zero value otherwise.

### GetECCValidCurvesOk

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetECCValidCurvesOk() (*[]string, bool)`

GetECCValidCurvesOk returns a tuple with the ECCValidCurves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCValidCurves

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetECCValidCurves(v []string)`

SetECCValidCurves sets ECCValidCurves field to given value.

### HasECCValidCurves

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasECCValidCurves() bool`

HasECCValidCurves returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


