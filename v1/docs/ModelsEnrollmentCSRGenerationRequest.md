# ModelsEnrollmentCSRGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Subject for the requested certificate | 
**KeyType** | **string** | Certificate key type [RSA, ECC] | 
**KeyLength** | **int32** | Size of the certificate key (ex: RSA 1024, 2048, 4096/ECC 256, 384, 521) | 
**Template** | Pointer to **string** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewModelsEnrollmentCSRGenerationRequest

`func NewModelsEnrollmentCSRGenerationRequest(subject string, keyType string, keyLength int32, ) *ModelsEnrollmentCSRGenerationRequest`

NewModelsEnrollmentCSRGenerationRequest instantiates a new ModelsEnrollmentCSRGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentCSRGenerationRequestWithDefaults

`func NewModelsEnrollmentCSRGenerationRequestWithDefaults() *ModelsEnrollmentCSRGenerationRequest`

NewModelsEnrollmentCSRGenerationRequestWithDefaults instantiates a new ModelsEnrollmentCSRGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetSANs

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *ModelsEnrollmentCSRGenerationRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *ModelsEnrollmentCSRGenerationRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *ModelsEnrollmentCSRGenerationRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


