# ModelsPkcs10CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**RequestDisposition** | Pointer to **NullableString** |  | [optional] 
**DispositionMessage** | Pointer to **NullableString** |  | [optional] 
**EnrollmentContext** | Pointer to **map[string]string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**KeyfactorID** | Pointer to **int32** |  | [optional] 
**Certificates** | Pointer to **[]string** |  | [optional] 
**WorkflowInstanceId** | Pointer to **string** |  | [optional] 
**WorkflowReferenceId** | Pointer to **int64** |  | [optional] 

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

### SetRequestDispositionNil

`func (o *ModelsPkcs10CertificateResponse) SetRequestDispositionNil(b bool)`

 SetRequestDispositionNil sets the value for RequestDisposition to be an explicit nil

### UnsetRequestDisposition
`func (o *ModelsPkcs10CertificateResponse) UnsetRequestDisposition()`

UnsetRequestDisposition ensures that no value is present for RequestDisposition, not even an explicit nil
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

### SetDispositionMessageNil

`func (o *ModelsPkcs10CertificateResponse) SetDispositionMessageNil(b bool)`

 SetDispositionMessageNil sets the value for DispositionMessage to be an explicit nil

### UnsetDispositionMessage
`func (o *ModelsPkcs10CertificateResponse) UnsetDispositionMessage()`

UnsetDispositionMessage ensures that no value is present for DispositionMessage, not even an explicit nil
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
### GetUrl

`func (o *ModelsPkcs10CertificateResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelsPkcs10CertificateResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelsPkcs10CertificateResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelsPkcs10CertificateResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ModelsPkcs10CertificateResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ModelsPkcs10CertificateResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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

### SetCertificatesNil

`func (o *ModelsPkcs10CertificateResponse) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ModelsPkcs10CertificateResponse) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetWorkflowInstanceId

`func (o *ModelsPkcs10CertificateResponse) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *ModelsPkcs10CertificateResponse) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *ModelsPkcs10CertificateResponse) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.

### HasWorkflowInstanceId

`func (o *ModelsPkcs10CertificateResponse) HasWorkflowInstanceId() bool`

HasWorkflowInstanceId returns a boolean if a field has been set.

### GetWorkflowReferenceId

`func (o *ModelsPkcs10CertificateResponse) GetWorkflowReferenceId() int64`

GetWorkflowReferenceId returns the WorkflowReferenceId field if non-nil, zero value otherwise.

### GetWorkflowReferenceIdOk

`func (o *ModelsPkcs10CertificateResponse) GetWorkflowReferenceIdOk() (*int64, bool)`

GetWorkflowReferenceIdOk returns a tuple with the WorkflowReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowReferenceId

`func (o *ModelsPkcs10CertificateResponse) SetWorkflowReferenceId(v int64)`

SetWorkflowReferenceId sets WorkflowReferenceId field to given value.

### HasWorkflowReferenceId

`func (o *ModelsPkcs10CertificateResponse) HasWorkflowReferenceId() bool`

HasWorkflowReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


