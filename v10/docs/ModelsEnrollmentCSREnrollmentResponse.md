# ModelsEnrollmentCSREnrollmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateInformation** | Pointer to [**ModelsPkcs10CertificateResponse**](ModelsPkcs10CertificateResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsEnrollmentCSREnrollmentResponse

`func NewModelsEnrollmentCSREnrollmentResponse() *ModelsEnrollmentCSREnrollmentResponse`

NewModelsEnrollmentCSREnrollmentResponse instantiates a new ModelsEnrollmentCSREnrollmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentCSREnrollmentResponseWithDefaults

`func NewModelsEnrollmentCSREnrollmentResponseWithDefaults() *ModelsEnrollmentCSREnrollmentResponse`

NewModelsEnrollmentCSREnrollmentResponseWithDefaults instantiates a new ModelsEnrollmentCSREnrollmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateInformation

`func (o *ModelsEnrollmentCSREnrollmentResponse) GetCertificateInformation() ModelsPkcs10CertificateResponse`

GetCertificateInformation returns the CertificateInformation field if non-nil, zero value otherwise.

### GetCertificateInformationOk

`func (o *ModelsEnrollmentCSREnrollmentResponse) GetCertificateInformationOk() (*ModelsPkcs10CertificateResponse, bool)`

GetCertificateInformationOk returns a tuple with the CertificateInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateInformation

`func (o *ModelsEnrollmentCSREnrollmentResponse) SetCertificateInformation(v ModelsPkcs10CertificateResponse)`

SetCertificateInformation sets CertificateInformation field to given value.

### HasCertificateInformation

`func (o *ModelsEnrollmentCSREnrollmentResponse) HasCertificateInformation() bool`

HasCertificateInformation returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsEnrollmentCSREnrollmentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentCSREnrollmentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentCSREnrollmentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentCSREnrollmentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


