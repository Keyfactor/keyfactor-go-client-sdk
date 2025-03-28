# CertificateAuthoritiesCAJobQueueTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Whether the test succeeded or failed. | [optional] 
**Message** | Pointer to **NullableString** | The message returned by the test. | [optional] 

## Methods

### NewCertificateAuthoritiesCAJobQueueTestResponse

`func NewCertificateAuthoritiesCAJobQueueTestResponse() *CertificateAuthoritiesCAJobQueueTestResponse`

NewCertificateAuthoritiesCAJobQueueTestResponse instantiates a new CertificateAuthoritiesCAJobQueueTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthoritiesCAJobQueueTestResponseWithDefaults

`func NewCertificateAuthoritiesCAJobQueueTestResponseWithDefaults() *CertificateAuthoritiesCAJobQueueTestResponse`

NewCertificateAuthoritiesCAJobQueueTestResponseWithDefaults instantiates a new CertificateAuthoritiesCAJobQueueTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CertificateAuthoritiesCAJobQueueTestResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CertificateAuthoritiesCAJobQueueTestResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


