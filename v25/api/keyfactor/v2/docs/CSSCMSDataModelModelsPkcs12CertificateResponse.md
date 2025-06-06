# CSSCMSDataModelModelsPkcs12CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**KeyfactorId** | Pointer to **int32** |  | [optional] 
**Pkcs12Blob** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**WorkflowInstanceId** | Pointer to **string** |  | [optional] 
**WorkflowReferenceId** | Pointer to **int64** |  | [optional] 
**StoreIdsInvalidForRenewal** | Pointer to **[]string** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**RequestDisposition** | Pointer to **NullableString** |  | [optional] 
**DispositionMessage** | Pointer to **NullableString** |  | [optional] 
**EnrollmentContext** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsPkcs12CertificateResponse

`func NewCSSCMSDataModelModelsPkcs12CertificateResponse() *CSSCMSDataModelModelsPkcs12CertificateResponse`

NewCSSCMSDataModelModelsPkcs12CertificateResponse instantiates a new CSSCMSDataModelModelsPkcs12CertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsPkcs12CertificateResponseWithDefaults

`func NewCSSCMSDataModelModelsPkcs12CertificateResponseWithDefaults() *CSSCMSDataModelModelsPkcs12CertificateResponse`

NewCSSCMSDataModelModelsPkcs12CertificateResponseWithDefaults instantiates a new CSSCMSDataModelModelsPkcs12CertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuerDN

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetKeyfactorId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyfactorId() int32`

GetKeyfactorId returns the KeyfactorId field if non-nil, zero value otherwise.

### GetKeyfactorIdOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyfactorIdOk() (*int32, bool)`

GetKeyfactorIdOk returns a tuple with the KeyfactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetKeyfactorId(v int32)`

SetKeyfactorId sets KeyfactorId field to given value.

### HasKeyfactorId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasKeyfactorId() bool`

HasKeyfactorId returns a boolean if a field has been set.

### GetPkcs12Blob

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetPkcs12Blob() string`

GetPkcs12Blob returns the Pkcs12Blob field if non-nil, zero value otherwise.

### GetPkcs12BlobOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetPkcs12BlobOk() (*string, bool)`

GetPkcs12BlobOk returns a tuple with the Pkcs12Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12Blob

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetPkcs12Blob(v string)`

SetPkcs12Blob sets Pkcs12Blob field to given value.

### HasPkcs12Blob

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasPkcs12Blob() bool`

HasPkcs12Blob returns a boolean if a field has been set.

### SetPkcs12BlobNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetPkcs12BlobNil(b bool)`

 SetPkcs12BlobNil sets the value for Pkcs12Blob to be an explicit nil

### UnsetPkcs12Blob
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetPkcs12Blob()`

UnsetPkcs12Blob ensures that no value is present for Pkcs12Blob, not even an explicit nil
### GetPassword

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetWorkflowInstanceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.

### HasWorkflowInstanceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasWorkflowInstanceId() bool`

HasWorkflowInstanceId returns a boolean if a field has been set.

### GetWorkflowReferenceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetWorkflowReferenceId() int64`

GetWorkflowReferenceId returns the WorkflowReferenceId field if non-nil, zero value otherwise.

### GetWorkflowReferenceIdOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetWorkflowReferenceIdOk() (*int64, bool)`

GetWorkflowReferenceIdOk returns a tuple with the WorkflowReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowReferenceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetWorkflowReferenceId(v int64)`

SetWorkflowReferenceId sets WorkflowReferenceId field to given value.

### HasWorkflowReferenceId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasWorkflowReferenceId() bool`

HasWorkflowReferenceId returns a boolean if a field has been set.

### GetStoreIdsInvalidForRenewal

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetStoreIdsInvalidForRenewal() []string`

GetStoreIdsInvalidForRenewal returns the StoreIdsInvalidForRenewal field if non-nil, zero value otherwise.

### GetStoreIdsInvalidForRenewalOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetStoreIdsInvalidForRenewalOk() (*[]string, bool)`

GetStoreIdsInvalidForRenewalOk returns a tuple with the StoreIdsInvalidForRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIdsInvalidForRenewal

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetStoreIdsInvalidForRenewal(v []string)`

SetStoreIdsInvalidForRenewal sets StoreIdsInvalidForRenewal field to given value.

### HasStoreIdsInvalidForRenewal

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasStoreIdsInvalidForRenewal() bool`

HasStoreIdsInvalidForRenewal returns a boolean if a field has been set.

### SetStoreIdsInvalidForRenewalNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetStoreIdsInvalidForRenewalNil(b bool)`

 SetStoreIdsInvalidForRenewalNil sets the value for StoreIdsInvalidForRenewal to be an explicit nil

### UnsetStoreIdsInvalidForRenewal
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetStoreIdsInvalidForRenewal()`

UnsetStoreIdsInvalidForRenewal ensures that no value is present for StoreIdsInvalidForRenewal, not even an explicit nil
### GetKeyfactorRequestId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetRequestDisposition

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetRequestDisposition() string`

GetRequestDisposition returns the RequestDisposition field if non-nil, zero value otherwise.

### GetRequestDispositionOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetRequestDispositionOk() (*string, bool)`

GetRequestDispositionOk returns a tuple with the RequestDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDisposition

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetRequestDisposition(v string)`

SetRequestDisposition sets RequestDisposition field to given value.

### HasRequestDisposition

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasRequestDisposition() bool`

HasRequestDisposition returns a boolean if a field has been set.

### SetRequestDispositionNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetRequestDispositionNil(b bool)`

 SetRequestDispositionNil sets the value for RequestDisposition to be an explicit nil

### UnsetRequestDisposition
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetRequestDisposition()`

UnsetRequestDisposition ensures that no value is present for RequestDisposition, not even an explicit nil
### GetDispositionMessage

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetDispositionMessage() string`

GetDispositionMessage returns the DispositionMessage field if non-nil, zero value otherwise.

### GetDispositionMessageOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetDispositionMessageOk() (*string, bool)`

GetDispositionMessageOk returns a tuple with the DispositionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionMessage

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetDispositionMessage(v string)`

SetDispositionMessage sets DispositionMessage field to given value.

### HasDispositionMessage

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasDispositionMessage() bool`

HasDispositionMessage returns a boolean if a field has been set.

### SetDispositionMessageNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetDispositionMessageNil(b bool)`

 SetDispositionMessageNil sets the value for DispositionMessage to be an explicit nil

### UnsetDispositionMessage
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetDispositionMessage()`

UnsetDispositionMessage ensures that no value is present for DispositionMessage, not even an explicit nil
### GetEnrollmentContext

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetEnrollmentContext() map[string]string`

GetEnrollmentContext returns the EnrollmentContext field if non-nil, zero value otherwise.

### GetEnrollmentContextOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetEnrollmentContextOk() (*map[string]string, bool)`

GetEnrollmentContextOk returns a tuple with the EnrollmentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentContext

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetEnrollmentContext(v map[string]string)`

SetEnrollmentContext sets EnrollmentContext field to given value.

### HasEnrollmentContext

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasEnrollmentContext() bool`

HasEnrollmentContext returns a boolean if a field has been set.

### SetEnrollmentContextNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetEnrollmentContextNil(b bool)`

 SetEnrollmentContextNil sets the value for EnrollmentContext to be an explicit nil

### UnsetEnrollmentContext
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetEnrollmentContext()`

UnsetEnrollmentContext ensures that no value is present for EnrollmentContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


