# ModelsSSLEndpointHistoryResponseCertificateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**SigningAlgorithm** | Pointer to **string** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]ModelsSubjectAlternativeName**](ModelsSubjectAlternativeName.md) |  | [optional] 

## Methods

### NewModelsSSLEndpointHistoryResponseCertificateModel

`func NewModelsSSLEndpointHistoryResponseCertificateModel() *ModelsSSLEndpointHistoryResponseCertificateModel`

NewModelsSSLEndpointHistoryResponseCertificateModel instantiates a new ModelsSSLEndpointHistoryResponseCertificateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSLEndpointHistoryResponseCertificateModelWithDefaults

`func NewModelsSSLEndpointHistoryResponseCertificateModelWithDefaults() *ModelsSSLEndpointHistoryResponseCertificateModel`

NewModelsSSLEndpointHistoryResponseCertificateModelWithDefaults instantiates a new ModelsSSLEndpointHistoryResponseCertificateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetSerialNumber

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetNotBefore

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetThumbprint

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetIssuedCN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetSubjectAltNameElements

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSubjectAltNameElements() []ModelsSubjectAlternativeName`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) GetSubjectAltNameElementsOk() (*[]ModelsSubjectAlternativeName, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) SetSubjectAltNameElements(v []ModelsSubjectAlternativeName)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *ModelsSSLEndpointHistoryResponseCertificateModel) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


