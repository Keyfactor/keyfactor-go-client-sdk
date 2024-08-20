# ModelsPkcs10CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**KeyfactorID** | Pointer to **int32** | The integer ID of the certificate in the keyfactor database, if Issued | [optional] 
**Certificates** | Pointer to **[]string** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** | The integer id of the certificate request in the Keyfactor database, if one exists. | [optional] 
**RequestDisposition** | Pointer to **string** |  | [optional] 
**DispositionMessage** | Pointer to **string** |  | [optional] 
**EnrollmentContext** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsPkcs10CertificateResponse

`func NewModelsPkcs10CertificateResponse() *ModelsPkcs10CertificateResponse`

NewModelsPkcs10CertificateResponse instantiates a new ModelsPkcs10CertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPkcs10CertificateResponseWithDefaults

`func NewModelsPkcs10CertificateResponseWithDefaults() *ModelsPkcs10CertificateResponse`

NewModelsPkcs10CertificateResponseWithDefaults instantiates a new ModelsPkcs10CertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *ModelsPkcs10CertificateResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsPkcs10CertificateResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsPkcs10CertificateResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsPkcs10CertificateResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsPkcs10CertificateResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsPkcs10CertificateResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsPkcs10CertificateResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsPkcs10CertificateResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsPkcs10CertificateResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsPkcs10CertificateResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsPkcs10CertificateResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsPkcs10CertificateResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsPkcs10CertificateResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsPkcs10CertificateResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetKeyfactorID

`func (o *ModelsPkcs10CertificateResponse) GetKeyfactorID() int32`

GetKeyfactorID returns the KeyfactorID field if non-nil, zero value otherwise.

### GetKeyfactorIDOk

`func (o *ModelsPkcs10CertificateResponse) GetKeyfactorIDOk() (*int32, bool)`

GetKeyfactorIDOk returns a tuple with the KeyfactorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorID

`func (o *ModelsPkcs10CertificateResponse) SetKeyfactorID(v int32)`

SetKeyfactorID sets KeyfactorID field to given value.

### HasKeyfactorID

`func (o *ModelsPkcs10CertificateResponse) HasKeyfactorID() bool`

HasKeyfactorID returns a boolean if a field has been set.

### GetCertificates

`func (o *ModelsPkcs10CertificateResponse) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ModelsPkcs10CertificateResponse) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ModelsPkcs10CertificateResponse) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ModelsPkcs10CertificateResponse) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetKeyfactorRequestId

`func (o *ModelsPkcs10CertificateResponse) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *ModelsPkcs10CertificateResponse) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *ModelsPkcs10CertificateResponse) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *ModelsPkcs10CertificateResponse) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetRequestDisposition

`func (o *ModelsPkcs10CertificateResponse) GetRequestDisposition() string`

GetRequestDisposition returns the RequestDisposition field if non-nil, zero value otherwise.

### GetRequestDispositionOk

`func (o *ModelsPkcs10CertificateResponse) GetRequestDispositionOk() (*string, bool)`

GetRequestDispositionOk returns a tuple with the RequestDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDisposition

`func (o *ModelsPkcs10CertificateResponse) SetRequestDisposition(v string)`

SetRequestDisposition sets RequestDisposition field to given value.

### HasRequestDisposition

`func (o *ModelsPkcs10CertificateResponse) HasRequestDisposition() bool`

HasRequestDisposition returns a boolean if a field has been set.

### GetDispositionMessage

`func (o *ModelsPkcs10CertificateResponse) GetDispositionMessage() string`

GetDispositionMessage returns the DispositionMessage field if non-nil, zero value otherwise.

### GetDispositionMessageOk

`func (o *ModelsPkcs10CertificateResponse) GetDispositionMessageOk() (*string, bool)`

GetDispositionMessageOk returns a tuple with the DispositionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionMessage

`func (o *ModelsPkcs10CertificateResponse) SetDispositionMessage(v string)`

SetDispositionMessage sets DispositionMessage field to given value.

### HasDispositionMessage

`func (o *ModelsPkcs10CertificateResponse) HasDispositionMessage() bool`

HasDispositionMessage returns a boolean if a field has been set.

### GetEnrollmentContext

`func (o *ModelsPkcs10CertificateResponse) GetEnrollmentContext() map[string]string`

GetEnrollmentContext returns the EnrollmentContext field if non-nil, zero value otherwise.

### GetEnrollmentContextOk

`func (o *ModelsPkcs10CertificateResponse) GetEnrollmentContextOk() (*map[string]string, bool)`

GetEnrollmentContextOk returns a tuple with the EnrollmentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentContext

`func (o *ModelsPkcs10CertificateResponse) SetEnrollmentContext(v map[string]string)`

SetEnrollmentContext sets EnrollmentContext field to given value.

### HasEnrollmentContext

`func (o *ModelsPkcs10CertificateResponse) HasEnrollmentContext() bool`

HasEnrollmentContext returns a boolean if a field has been set.

### SetEnrollmentContextNil

`func (o *ModelsPkcs10CertificateResponse) SetEnrollmentContextNil(b bool)`

 SetEnrollmentContextNil sets the value for EnrollmentContext to be an explicit nil

### UnsetEnrollmentContext
`func (o *ModelsPkcs10CertificateResponse) UnsetEnrollmentContext()`

UnsetEnrollmentContext ensures that no value is present for EnrollmentContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


