# EnrollmentCSRGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Subject for the requested certificate | 
**KeyType** | **string** | Certificate key type [RSA, ECC, ED448, ED25519] | 
**KeyLength** | Pointer to **int32** | Size of the certificate key (ex: RSA 1024, 2048, 4096/ECC 256, 384, 521) | [optional] 
**Curve** | Pointer to **NullableString** | The curve used to generate a CSR. | [optional] 
**AlternativeKeyType** | Pointer to **NullableString** | Alternative Certificate key type [ML-DSA-44, ML-DSA-65, ML-DSA-87] | [optional] 
**AlternativeKeyLength** | Pointer to **NullableInt32** | Size of the alternative certificate key. | [optional] 
**AlternativeCurve** | Pointer to **NullableString** | The alternative curve used to generate a CSR. | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewEnrollmentCSRGenerationRequest

`func NewEnrollmentCSRGenerationRequest(subject string, keyType string, ) *EnrollmentCSRGenerationRequest`

NewEnrollmentCSRGenerationRequest instantiates a new EnrollmentCSRGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentCSRGenerationRequestWithDefaults

`func NewEnrollmentCSRGenerationRequestWithDefaults() *EnrollmentCSRGenerationRequest`

NewEnrollmentCSRGenerationRequestWithDefaults instantiates a new EnrollmentCSRGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EnrollmentCSRGenerationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EnrollmentCSRGenerationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EnrollmentCSRGenerationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetKeyType

`func (o *EnrollmentCSRGenerationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *EnrollmentCSRGenerationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *EnrollmentCSRGenerationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyLength

`func (o *EnrollmentCSRGenerationRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *EnrollmentCSRGenerationRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *EnrollmentCSRGenerationRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *EnrollmentCSRGenerationRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *EnrollmentCSRGenerationRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *EnrollmentCSRGenerationRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *EnrollmentCSRGenerationRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *EnrollmentCSRGenerationRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *EnrollmentCSRGenerationRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *EnrollmentCSRGenerationRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetAlternativeKeyType

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeKeyType() string`

GetAlternativeKeyType returns the AlternativeKeyType field if non-nil, zero value otherwise.

### GetAlternativeKeyTypeOk

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeKeyTypeOk() (*string, bool)`

GetAlternativeKeyTypeOk returns a tuple with the AlternativeKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyType

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeKeyType(v string)`

SetAlternativeKeyType sets AlternativeKeyType field to given value.

### HasAlternativeKeyType

`func (o *EnrollmentCSRGenerationRequest) HasAlternativeKeyType() bool`

HasAlternativeKeyType returns a boolean if a field has been set.

### SetAlternativeKeyTypeNil

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeKeyTypeNil(b bool)`

 SetAlternativeKeyTypeNil sets the value for AlternativeKeyType to be an explicit nil

### UnsetAlternativeKeyType
`func (o *EnrollmentCSRGenerationRequest) UnsetAlternativeKeyType()`

UnsetAlternativeKeyType ensures that no value is present for AlternativeKeyType, not even an explicit nil
### GetAlternativeKeyLength

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeKeyLength() int32`

GetAlternativeKeyLength returns the AlternativeKeyLength field if non-nil, zero value otherwise.

### GetAlternativeKeyLengthOk

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeKeyLengthOk() (*int32, bool)`

GetAlternativeKeyLengthOk returns a tuple with the AlternativeKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyLength

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeKeyLength(v int32)`

SetAlternativeKeyLength sets AlternativeKeyLength field to given value.

### HasAlternativeKeyLength

`func (o *EnrollmentCSRGenerationRequest) HasAlternativeKeyLength() bool`

HasAlternativeKeyLength returns a boolean if a field has been set.

### SetAlternativeKeyLengthNil

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeKeyLengthNil(b bool)`

 SetAlternativeKeyLengthNil sets the value for AlternativeKeyLength to be an explicit nil

### UnsetAlternativeKeyLength
`func (o *EnrollmentCSRGenerationRequest) UnsetAlternativeKeyLength()`

UnsetAlternativeKeyLength ensures that no value is present for AlternativeKeyLength, not even an explicit nil
### GetAlternativeCurve

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeCurve() string`

GetAlternativeCurve returns the AlternativeCurve field if non-nil, zero value otherwise.

### GetAlternativeCurveOk

`func (o *EnrollmentCSRGenerationRequest) GetAlternativeCurveOk() (*string, bool)`

GetAlternativeCurveOk returns a tuple with the AlternativeCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeCurve

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeCurve(v string)`

SetAlternativeCurve sets AlternativeCurve field to given value.

### HasAlternativeCurve

`func (o *EnrollmentCSRGenerationRequest) HasAlternativeCurve() bool`

HasAlternativeCurve returns a boolean if a field has been set.

### SetAlternativeCurveNil

`func (o *EnrollmentCSRGenerationRequest) SetAlternativeCurveNil(b bool)`

 SetAlternativeCurveNil sets the value for AlternativeCurve to be an explicit nil

### UnsetAlternativeCurve
`func (o *EnrollmentCSRGenerationRequest) UnsetAlternativeCurve()`

UnsetAlternativeCurve ensures that no value is present for AlternativeCurve, not even an explicit nil
### GetTemplate

`func (o *EnrollmentCSRGenerationRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentCSRGenerationRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentCSRGenerationRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentCSRGenerationRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EnrollmentCSRGenerationRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EnrollmentCSRGenerationRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSANs

`func (o *EnrollmentCSRGenerationRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *EnrollmentCSRGenerationRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *EnrollmentCSRGenerationRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *EnrollmentCSRGenerationRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *EnrollmentCSRGenerationRequest) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *EnrollmentCSRGenerationRequest) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


