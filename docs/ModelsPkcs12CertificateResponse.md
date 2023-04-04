# ModelsPkcs12CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**KeyfactorId** | Pointer to **int32** |  | [optional] 
**Pkcs12Blob** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** | The integer id of the certificate request in the Keyfactor database, if one exists. | [optional] 
**RequestDisposition** | Pointer to **string** |  | [optional] 
**DispositionMessage** | Pointer to **string** |  | [optional] 
**EnrollmentContext** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsPkcs12CertificateResponse

`func NewModelsPkcs12CertificateResponse() *ModelsPkcs12CertificateResponse`

NewModelsPkcs12CertificateResponse instantiates a new ModelsPkcs12CertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPkcs12CertificateResponseWithDefaults

`func NewModelsPkcs12CertificateResponseWithDefaults() *ModelsPkcs12CertificateResponse`

NewModelsPkcs12CertificateResponseWithDefaults instantiates a new ModelsPkcs12CertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *ModelsPkcs12CertificateResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsPkcs12CertificateResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsPkcs12CertificateResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsPkcs12CertificateResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsPkcs12CertificateResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsPkcs12CertificateResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsPkcs12CertificateResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsPkcs12CertificateResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsPkcs12CertificateResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsPkcs12CertificateResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsPkcs12CertificateResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsPkcs12CertificateResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsPkcs12CertificateResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsPkcs12CertificateResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetKeyfactorId

`func (o *ModelsPkcs12CertificateResponse) GetKeyfactorId() int32`

GetKeyfactorId returns the KeyfactorId field if non-nil, zero value otherwise.

### GetKeyfactorIdOk

`func (o *ModelsPkcs12CertificateResponse) GetKeyfactorIdOk() (*int32, bool)`

GetKeyfactorIdOk returns a tuple with the KeyfactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorId

`func (o *ModelsPkcs12CertificateResponse) SetKeyfactorId(v int32)`

SetKeyfactorId sets KeyfactorId field to given value.

### HasKeyfactorId

`func (o *ModelsPkcs12CertificateResponse) HasKeyfactorId() bool`

HasKeyfactorId returns a boolean if a field has been set.

### GetPkcs12Blob

`func (o *ModelsPkcs12CertificateResponse) GetPkcs12Blob() string`

GetPkcs12Blob returns the Pkcs12Blob field if non-nil, zero value otherwise.

### GetPkcs12BlobOk

`func (o *ModelsPkcs12CertificateResponse) GetPkcs12BlobOk() (*string, bool)`

GetPkcs12BlobOk returns a tuple with the Pkcs12Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12Blob

`func (o *ModelsPkcs12CertificateResponse) SetPkcs12Blob(v string)`

SetPkcs12Blob sets Pkcs12Blob field to given value.

### HasPkcs12Blob

`func (o *ModelsPkcs12CertificateResponse) HasPkcs12Blob() bool`

HasPkcs12Blob returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsPkcs12CertificateResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsPkcs12CertificateResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsPkcs12CertificateResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsPkcs12CertificateResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetKeyfactorRequestId

`func (o *ModelsPkcs12CertificateResponse) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *ModelsPkcs12CertificateResponse) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *ModelsPkcs12CertificateResponse) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *ModelsPkcs12CertificateResponse) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetRequestDisposition

`func (o *ModelsPkcs12CertificateResponse) GetRequestDisposition() string`

GetRequestDisposition returns the RequestDisposition field if non-nil, zero value otherwise.

### GetRequestDispositionOk

`func (o *ModelsPkcs12CertificateResponse) GetRequestDispositionOk() (*string, bool)`

GetRequestDispositionOk returns a tuple with the RequestDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDisposition

`func (o *ModelsPkcs12CertificateResponse) SetRequestDisposition(v string)`

SetRequestDisposition sets RequestDisposition field to given value.

### HasRequestDisposition

`func (o *ModelsPkcs12CertificateResponse) HasRequestDisposition() bool`

HasRequestDisposition returns a boolean if a field has been set.

### GetDispositionMessage

`func (o *ModelsPkcs12CertificateResponse) GetDispositionMessage() string`

GetDispositionMessage returns the DispositionMessage field if non-nil, zero value otherwise.

### GetDispositionMessageOk

`func (o *ModelsPkcs12CertificateResponse) GetDispositionMessageOk() (*string, bool)`

GetDispositionMessageOk returns a tuple with the DispositionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionMessage

`func (o *ModelsPkcs12CertificateResponse) SetDispositionMessage(v string)`

SetDispositionMessage sets DispositionMessage field to given value.

### HasDispositionMessage

`func (o *ModelsPkcs12CertificateResponse) HasDispositionMessage() bool`

HasDispositionMessage returns a boolean if a field has been set.

### GetEnrollmentContext

`func (o *ModelsPkcs12CertificateResponse) GetEnrollmentContext() map[string]string`

GetEnrollmentContext returns the EnrollmentContext field if non-nil, zero value otherwise.

### GetEnrollmentContextOk

`func (o *ModelsPkcs12CertificateResponse) GetEnrollmentContextOk() (*map[string]string, bool)`

GetEnrollmentContextOk returns a tuple with the EnrollmentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentContext

`func (o *ModelsPkcs12CertificateResponse) SetEnrollmentContext(v map[string]string)`

SetEnrollmentContext sets EnrollmentContext field to given value.

### HasEnrollmentContext

`func (o *ModelsPkcs12CertificateResponse) HasEnrollmentContext() bool`

HasEnrollmentContext returns a boolean if a field has been set.

### SetEnrollmentContextNil

`func (o *ModelsPkcs12CertificateResponse) SetEnrollmentContextNil(b bool)`

 SetEnrollmentContextNil sets the value for EnrollmentContext to be an explicit nil

### UnsetEnrollmentContext
`func (o *ModelsPkcs12CertificateResponse) UnsetEnrollmentContext()`

UnsetEnrollmentContext ensures that no value is present for EnrollmentContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


