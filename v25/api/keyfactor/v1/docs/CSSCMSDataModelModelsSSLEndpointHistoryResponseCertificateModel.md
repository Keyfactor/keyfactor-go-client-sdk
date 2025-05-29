# CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**SigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]CSSCMSDataModelModelsSubjectAlternativeName**](CSSCMSDataModelModelsSubjectAlternativeName.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel

`func NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel() *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel`

NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel instantiates a new CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModelWithDefaults

`func NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModelWithDefaults() *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel`

NewCSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModelWithDefaults instantiates a new CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetSerialNumber

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetNotBefore

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetThumbprint

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetIssuerDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetIssuedCN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetSubjectAltNameElements

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSubjectAltNameElements() []CSSCMSDataModelModelsSubjectAlternativeName`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) GetSubjectAltNameElementsOk() (*[]CSSCMSDataModelModelsSubjectAlternativeName, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSubjectAltNameElements(v []CSSCMSDataModelModelsSubjectAlternativeName)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.

### SetSubjectAltNameElementsNil

`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) SetSubjectAltNameElementsNil(b bool)`

 SetSubjectAltNameElementsNil sets the value for SubjectAltNameElements to be an explicit nil

### UnsetSubjectAltNameElements
`func (o *CSSCMSDataModelModelsSSLEndpointHistoryResponseCertificateModel) UnsetSubjectAltNameElements()`

UnsetSubjectAltNameElements ensures that no value is present for SubjectAltNameElements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


