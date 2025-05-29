# CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateInformation** | Pointer to [**CSSCMSDataModelModelsPkcs12CertificateResponse**](CSSCMSDataModelModelsPkcs12CertificateResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse

`func NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse() *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse`

NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse instantiates a new CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponseWithDefaults

`func NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponseWithDefaults() *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse`

NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentResponseWithDefaults instantiates a new CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) GetCertificateInformation() CSSCMSDataModelModelsPkcs12CertificateResponse`

GetCertificateInformation returns the CertificateInformation field if non-nil, zero value otherwise.

### GetCertificateInformationOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) GetCertificateInformationOk() (*CSSCMSDataModelModelsPkcs12CertificateResponse, bool)`

GetCertificateInformationOk returns a tuple with the CertificateInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) SetCertificateInformation(v CSSCMSDataModelModelsPkcs12CertificateResponse)`

SetCertificateInformation sets CertificateInformation field to given value.

### HasCertificateInformation

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) HasCertificateInformation() bool`

HasCertificateInformation returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


