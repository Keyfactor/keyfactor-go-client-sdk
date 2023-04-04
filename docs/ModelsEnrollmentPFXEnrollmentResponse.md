# ModelsEnrollmentPFXEnrollmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateInformation** | Pointer to [**ModelsPkcs12CertificateResponse**](ModelsPkcs12CertificateResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsEnrollmentPFXEnrollmentResponse

`func NewModelsEnrollmentPFXEnrollmentResponse() *ModelsEnrollmentPFXEnrollmentResponse`

NewModelsEnrollmentPFXEnrollmentResponse instantiates a new ModelsEnrollmentPFXEnrollmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentPFXEnrollmentResponseWithDefaults

`func NewModelsEnrollmentPFXEnrollmentResponseWithDefaults() *ModelsEnrollmentPFXEnrollmentResponse`

NewModelsEnrollmentPFXEnrollmentResponseWithDefaults instantiates a new ModelsEnrollmentPFXEnrollmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateInformation

`func (o *ModelsEnrollmentPFXEnrollmentResponse) GetCertificateInformation() ModelsPkcs12CertificateResponse`

GetCertificateInformation returns the CertificateInformation field if non-nil, zero value otherwise.

### GetCertificateInformationOk

`func (o *ModelsEnrollmentPFXEnrollmentResponse) GetCertificateInformationOk() (*ModelsPkcs12CertificateResponse, bool)`

GetCertificateInformationOk returns a tuple with the CertificateInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateInformation

`func (o *ModelsEnrollmentPFXEnrollmentResponse) SetCertificateInformation(v ModelsPkcs12CertificateResponse)`

SetCertificateInformation sets CertificateInformation field to given value.

### HasCertificateInformation

`func (o *ModelsEnrollmentPFXEnrollmentResponse) HasCertificateInformation() bool`

HasCertificateInformation returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentPFXEnrollmentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentPFXEnrollmentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


