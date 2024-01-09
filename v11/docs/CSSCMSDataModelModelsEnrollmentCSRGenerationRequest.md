# CSSCMSDataModelModelsEnrollmentCSRGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **NullableString** |  | [optional] 
**SaNs** | Pointer to **map[string][]string** |  | [optional] 
**Subject** | **string** |  | 
**KeyType** | **string** |  | 
**KeyLength** | Pointer to **int32** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequest

`func NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequest(subject string, keyType string, ) *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest`

NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequest instantiates a new CSSCMSDataModelModelsEnrollmentCSRGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequestWithDefaults

`func NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequestWithDefaults() *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest`

NewCSSCMSDataModelModelsEnrollmentCSRGenerationRequestWithDefaults instantiates a new CSSCMSDataModelModelsEnrollmentCSRGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetSubject

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetKeyType

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CSSCMSDataModelModelsEnrollmentCSRGenerationRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


