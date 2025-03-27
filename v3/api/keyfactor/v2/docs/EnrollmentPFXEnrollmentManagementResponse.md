# EnrollmentPFXEnrollmentManagementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessfulStores** | Pointer to **[]string** |  | [optional] 
**CertificateInformation** | Pointer to [**CSSCMSDataModelModelsPkcs12CertificateResponse**](CSSCMSDataModelModelsPkcs12CertificateResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEnrollmentPFXEnrollmentManagementResponse

`func NewEnrollmentPFXEnrollmentManagementResponse() *EnrollmentPFXEnrollmentManagementResponse`

NewEnrollmentPFXEnrollmentManagementResponse instantiates a new EnrollmentPFXEnrollmentManagementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPFXEnrollmentManagementResponseWithDefaults

`func NewEnrollmentPFXEnrollmentManagementResponseWithDefaults() *EnrollmentPFXEnrollmentManagementResponse`

NewEnrollmentPFXEnrollmentManagementResponseWithDefaults instantiates a new EnrollmentPFXEnrollmentManagementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessfulStores

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetSuccessfulStores() []string`

GetSuccessfulStores returns the SuccessfulStores field if non-nil, zero value otherwise.

### GetSuccessfulStoresOk

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetSuccessfulStoresOk() (*[]string, bool)`

GetSuccessfulStoresOk returns a tuple with the SuccessfulStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulStores

`func (o *EnrollmentPFXEnrollmentManagementResponse) SetSuccessfulStores(v []string)`

SetSuccessfulStores sets SuccessfulStores field to given value.

### HasSuccessfulStores

`func (o *EnrollmentPFXEnrollmentManagementResponse) HasSuccessfulStores() bool`

HasSuccessfulStores returns a boolean if a field has been set.

### SetSuccessfulStoresNil

`func (o *EnrollmentPFXEnrollmentManagementResponse) SetSuccessfulStoresNil(b bool)`

 SetSuccessfulStoresNil sets the value for SuccessfulStores to be an explicit nil

### UnsetSuccessfulStores
`func (o *EnrollmentPFXEnrollmentManagementResponse) UnsetSuccessfulStores()`

UnsetSuccessfulStores ensures that no value is present for SuccessfulStores, not even an explicit nil
### GetCertificateInformation

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetCertificateInformation() CSSCMSDataModelModelsPkcs12CertificateResponse`

GetCertificateInformation returns the CertificateInformation field if non-nil, zero value otherwise.

### GetCertificateInformationOk

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetCertificateInformationOk() (*CSSCMSDataModelModelsPkcs12CertificateResponse, bool)`

GetCertificateInformationOk returns a tuple with the CertificateInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateInformation

`func (o *EnrollmentPFXEnrollmentManagementResponse) SetCertificateInformation(v CSSCMSDataModelModelsPkcs12CertificateResponse)`

SetCertificateInformation sets CertificateInformation field to given value.

### HasCertificateInformation

`func (o *EnrollmentPFXEnrollmentManagementResponse) HasCertificateInformation() bool`

HasCertificateInformation returns a boolean if a field has been set.

### GetMetadata

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnrollmentPFXEnrollmentManagementResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnrollmentPFXEnrollmentManagementResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnrollmentPFXEnrollmentManagementResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EnrollmentPFXEnrollmentManagementResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EnrollmentPFXEnrollmentManagementResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


