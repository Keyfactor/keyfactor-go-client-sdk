# EnrollmentRenewalApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyfactorId** | Pointer to **int32** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**RequestDisposition** | Pointer to **NullableString** |  | [optional] 
**DispositionMessage** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnrollmentRenewalApiResponse

`func NewEnrollmentRenewalApiResponse() *EnrollmentRenewalApiResponse`

NewEnrollmentRenewalApiResponse instantiates a new EnrollmentRenewalApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentRenewalApiResponseWithDefaults

`func NewEnrollmentRenewalApiResponseWithDefaults() *EnrollmentRenewalApiResponse`

NewEnrollmentRenewalApiResponseWithDefaults instantiates a new EnrollmentRenewalApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyfactorId

`func (o *EnrollmentRenewalApiResponse) GetKeyfactorId() int32`

GetKeyfactorId returns the KeyfactorId field if non-nil, zero value otherwise.

### GetKeyfactorIdOk

`func (o *EnrollmentRenewalApiResponse) GetKeyfactorIdOk() (*int32, bool)`

GetKeyfactorIdOk returns a tuple with the KeyfactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorId

`func (o *EnrollmentRenewalApiResponse) SetKeyfactorId(v int32)`

SetKeyfactorId sets KeyfactorId field to given value.

### HasKeyfactorId

`func (o *EnrollmentRenewalApiResponse) HasKeyfactorId() bool`

HasKeyfactorId returns a boolean if a field has been set.

### GetKeyfactorRequestId

`func (o *EnrollmentRenewalApiResponse) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *EnrollmentRenewalApiResponse) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *EnrollmentRenewalApiResponse) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *EnrollmentRenewalApiResponse) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetThumbprint

`func (o *EnrollmentRenewalApiResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *EnrollmentRenewalApiResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *EnrollmentRenewalApiResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *EnrollmentRenewalApiResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *EnrollmentRenewalApiResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *EnrollmentRenewalApiResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetSerialNumber

`func (o *EnrollmentRenewalApiResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *EnrollmentRenewalApiResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *EnrollmentRenewalApiResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *EnrollmentRenewalApiResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *EnrollmentRenewalApiResponse) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *EnrollmentRenewalApiResponse) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuerDN

`func (o *EnrollmentRenewalApiResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *EnrollmentRenewalApiResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *EnrollmentRenewalApiResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *EnrollmentRenewalApiResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *EnrollmentRenewalApiResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *EnrollmentRenewalApiResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetRequestDisposition

`func (o *EnrollmentRenewalApiResponse) GetRequestDisposition() string`

GetRequestDisposition returns the RequestDisposition field if non-nil, zero value otherwise.

### GetRequestDispositionOk

`func (o *EnrollmentRenewalApiResponse) GetRequestDispositionOk() (*string, bool)`

GetRequestDispositionOk returns a tuple with the RequestDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDisposition

`func (o *EnrollmentRenewalApiResponse) SetRequestDisposition(v string)`

SetRequestDisposition sets RequestDisposition field to given value.

### HasRequestDisposition

`func (o *EnrollmentRenewalApiResponse) HasRequestDisposition() bool`

HasRequestDisposition returns a boolean if a field has been set.

### SetRequestDispositionNil

`func (o *EnrollmentRenewalApiResponse) SetRequestDispositionNil(b bool)`

 SetRequestDispositionNil sets the value for RequestDisposition to be an explicit nil

### UnsetRequestDisposition
`func (o *EnrollmentRenewalApiResponse) UnsetRequestDisposition()`

UnsetRequestDisposition ensures that no value is present for RequestDisposition, not even an explicit nil
### GetDispositionMessage

`func (o *EnrollmentRenewalApiResponse) GetDispositionMessage() string`

GetDispositionMessage returns the DispositionMessage field if non-nil, zero value otherwise.

### GetDispositionMessageOk

`func (o *EnrollmentRenewalApiResponse) GetDispositionMessageOk() (*string, bool)`

GetDispositionMessageOk returns a tuple with the DispositionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionMessage

`func (o *EnrollmentRenewalApiResponse) SetDispositionMessage(v string)`

SetDispositionMessage sets DispositionMessage field to given value.

### HasDispositionMessage

`func (o *EnrollmentRenewalApiResponse) HasDispositionMessage() bool`

HasDispositionMessage returns a boolean if a field has been set.

### SetDispositionMessageNil

`func (o *EnrollmentRenewalApiResponse) SetDispositionMessageNil(b bool)`

 SetDispositionMessageNil sets the value for DispositionMessage to be an explicit nil

### UnsetDispositionMessage
`func (o *EnrollmentRenewalApiResponse) UnsetDispositionMessage()`

UnsetDispositionMessage ensures that no value is present for DispositionMessage, not even an explicit nil
### GetPassword

`func (o *EnrollmentRenewalApiResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollmentRenewalApiResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollmentRenewalApiResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnrollmentRenewalApiResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EnrollmentRenewalApiResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EnrollmentRenewalApiResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


