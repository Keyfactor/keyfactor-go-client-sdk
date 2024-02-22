# ModelsEnrollmentCSRGenerationRequest

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

### NewModelsEnrollmentCSRGenerationRequest

`func NewModelsEnrollmentCSRGenerationRequest(subject string, keyType string, ) *ModelsEnrollmentCSRGenerationRequest`

NewModelsEnrollmentCSRGenerationRequest instantiates a new ModelsEnrollmentCSRGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentCSRGenerationRequestWithDefaults

`func NewModelsEnrollmentCSRGenerationRequestWithDefaults() *ModelsEnrollmentCSRGenerationRequest`

NewModelsEnrollmentCSRGenerationRequestWithDefaults instantiates a new ModelsEnrollmentCSRGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ModelsEnrollmentCSRGenerationRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsEnrollmentCSRGenerationRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsEnrollmentCSRGenerationRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *ModelsEnrollmentCSRGenerationRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *ModelsEnrollmentCSRGenerationRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *ModelsEnrollmentCSRGenerationRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *ModelsEnrollmentCSRGenerationRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *ModelsEnrollmentCSRGenerationRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *ModelsEnrollmentCSRGenerationRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetSubject

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelsEnrollmentCSRGenerationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetKeyType

`func (o *ModelsEnrollmentCSRGenerationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsEnrollmentCSRGenerationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyLength

`func (o *ModelsEnrollmentCSRGenerationRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsEnrollmentCSRGenerationRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *ModelsEnrollmentCSRGenerationRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *ModelsEnrollmentCSRGenerationRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *ModelsEnrollmentCSRGenerationRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *ModelsEnrollmentCSRGenerationRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *ModelsEnrollmentCSRGenerationRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *ModelsEnrollmentCSRGenerationRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


