# ModelsEnrollmentRenewalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyfactorId** | Pointer to **int32** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**RequestDisposition** | Pointer to **string** |  | [optional] 
**DispositionMessage** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsEnrollmentRenewalResponse

`func NewModelsEnrollmentRenewalResponse() *ModelsEnrollmentRenewalResponse`

NewModelsEnrollmentRenewalResponse instantiates a new ModelsEnrollmentRenewalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentRenewalResponseWithDefaults

`func NewModelsEnrollmentRenewalResponseWithDefaults() *ModelsEnrollmentRenewalResponse`

NewModelsEnrollmentRenewalResponseWithDefaults instantiates a new ModelsEnrollmentRenewalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyfactorId

`func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorId() int32`

GetKeyfactorId returns the KeyfactorId field if non-nil, zero value otherwise.

### GetKeyfactorIdOk

`func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorIdOk() (*int32, bool)`

GetKeyfactorIdOk returns a tuple with the KeyfactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorId

`func (o *ModelsEnrollmentRenewalResponse) SetKeyfactorId(v int32)`

SetKeyfactorId sets KeyfactorId field to given value.

### HasKeyfactorId

`func (o *ModelsEnrollmentRenewalResponse) HasKeyfactorId() bool`

HasKeyfactorId returns a boolean if a field has been set.

### GetKeyfactorRequestId

`func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *ModelsEnrollmentRenewalResponse) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *ModelsEnrollmentRenewalResponse) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetThumbprint

`func (o *ModelsEnrollmentRenewalResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsEnrollmentRenewalResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsEnrollmentRenewalResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsEnrollmentRenewalResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ModelsEnrollmentRenewalResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsEnrollmentRenewalResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsEnrollmentRenewalResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsEnrollmentRenewalResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsEnrollmentRenewalResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsEnrollmentRenewalResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsEnrollmentRenewalResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsEnrollmentRenewalResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsEnrollmentRenewalResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsEnrollmentRenewalResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetRequestDisposition

`func (o *ModelsEnrollmentRenewalResponse) GetRequestDisposition() string`

GetRequestDisposition returns the RequestDisposition field if non-nil, zero value otherwise.

### GetRequestDispositionOk

`func (o *ModelsEnrollmentRenewalResponse) GetRequestDispositionOk() (*string, bool)`

GetRequestDispositionOk returns a tuple with the RequestDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDisposition

`func (o *ModelsEnrollmentRenewalResponse) SetRequestDisposition(v string)`

SetRequestDisposition sets RequestDisposition field to given value.

### HasRequestDisposition

`func (o *ModelsEnrollmentRenewalResponse) HasRequestDisposition() bool`

HasRequestDisposition returns a boolean if a field has been set.

### GetDispositionMessage

`func (o *ModelsEnrollmentRenewalResponse) GetDispositionMessage() string`

GetDispositionMessage returns the DispositionMessage field if non-nil, zero value otherwise.

### GetDispositionMessageOk

`func (o *ModelsEnrollmentRenewalResponse) GetDispositionMessageOk() (*string, bool)`

GetDispositionMessageOk returns a tuple with the DispositionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionMessage

`func (o *ModelsEnrollmentRenewalResponse) SetDispositionMessage(v string)`

SetDispositionMessage sets DispositionMessage field to given value.

### HasDispositionMessage

`func (o *ModelsEnrollmentRenewalResponse) HasDispositionMessage() bool`

HasDispositionMessage returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsEnrollmentRenewalResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsEnrollmentRenewalResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsEnrollmentRenewalResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsEnrollmentRenewalResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


