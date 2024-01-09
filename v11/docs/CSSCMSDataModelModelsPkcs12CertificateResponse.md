# CSSCMSDataModelModelsPkcs12CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**RequestDisposition** | Pointer to **NullableString** |  | [optional] 
**DispositionMessage** | Pointer to **NullableString** |  | [optional] 
**EnrollmentContext** | Pointer to **map[string]string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**KeyfactorId** | Pointer to **int32** |  | [optional] 
**Pkcs12Blob** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**KeyStatus** | Pointer to [**CSSCMSDataModelEnumsPrivateKeyStatus**](CSSCMSDataModelEnumsPrivateKeyStatus.md) |  | [optional] 
**WorkflowInstanceId** | Pointer to **string** |  | [optional] 
**WorkflowReferenceId** | Pointer to **int64** |  | [optional] 
**StoreIdsInvalidForRenewal** | Pointer to **[]string** |  | [optional] 

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
### GetUrl

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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
### GetKeyStatus

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyStatus() CSSCMSDataModelEnumsPrivateKeyStatus`

GetKeyStatus returns the KeyStatus field if non-nil, zero value otherwise.

### GetKeyStatusOk

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) GetKeyStatusOk() (*CSSCMSDataModelEnumsPrivateKeyStatus, bool)`

GetKeyStatusOk returns a tuple with the KeyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStatus

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) SetKeyStatus(v CSSCMSDataModelEnumsPrivateKeyStatus)`

SetKeyStatus sets KeyStatus field to given value.

### HasKeyStatus

`func (o *CSSCMSDataModelModelsPkcs12CertificateResponse) HasKeyStatus() bool`

HasKeyStatus returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


