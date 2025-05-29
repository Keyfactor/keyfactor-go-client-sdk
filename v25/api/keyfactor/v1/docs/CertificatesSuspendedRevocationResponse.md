# CertificatesSuspendedRevocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertId** | Pointer to **int32** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificatesSuspendedRevocationResponse

`func NewCertificatesSuspendedRevocationResponse() *CertificatesSuspendedRevocationResponse`

NewCertificatesSuspendedRevocationResponse instantiates a new CertificatesSuspendedRevocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesSuspendedRevocationResponseWithDefaults

`func NewCertificatesSuspendedRevocationResponseWithDefaults() *CertificatesSuspendedRevocationResponse`

NewCertificatesSuspendedRevocationResponseWithDefaults instantiates a new CertificatesSuspendedRevocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertId

`func (o *CertificatesSuspendedRevocationResponse) GetCertId() int32`

GetCertId returns the CertId field if non-nil, zero value otherwise.

### GetCertIdOk

`func (o *CertificatesSuspendedRevocationResponse) GetCertIdOk() (*int32, bool)`

GetCertIdOk returns a tuple with the CertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertId

`func (o *CertificatesSuspendedRevocationResponse) SetCertId(v int32)`

SetCertId sets CertId field to given value.

### HasCertId

`func (o *CertificatesSuspendedRevocationResponse) HasCertId() bool`

HasCertId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *CertificatesSuspendedRevocationResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CertificatesSuspendedRevocationResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CertificatesSuspendedRevocationResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CertificatesSuspendedRevocationResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetMessage

`func (o *CertificatesSuspendedRevocationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertificatesSuspendedRevocationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertificatesSuspendedRevocationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CertificatesSuspendedRevocationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CertificatesSuspendedRevocationResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CertificatesSuspendedRevocationResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


