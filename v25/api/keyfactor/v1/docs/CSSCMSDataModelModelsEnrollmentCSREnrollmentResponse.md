# CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateInformation** | Pointer to [**CSSCMSDataModelModelsPkcs10CertificateResponse**](CSSCMSDataModelModelsPkcs10CertificateResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponse

`func NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponse() *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse`

NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponse instantiates a new CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponseWithDefaults

`func NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponseWithDefaults() *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse`

NewCSSCMSDataModelModelsEnrollmentCSREnrollmentResponseWithDefaults instantiates a new CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) GetCertificateInformation() CSSCMSDataModelModelsPkcs10CertificateResponse`

GetCertificateInformation returns the CertificateInformation field if non-nil, zero value otherwise.

### GetCertificateInformationOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) GetCertificateInformationOk() (*CSSCMSDataModelModelsPkcs10CertificateResponse, bool)`

GetCertificateInformationOk returns a tuple with the CertificateInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) SetCertificateInformation(v CSSCMSDataModelModelsPkcs10CertificateResponse)`

SetCertificateInformation sets CertificateInformation field to given value.

### HasCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) HasCertificateInformation() bool`

HasCertificateInformation returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


