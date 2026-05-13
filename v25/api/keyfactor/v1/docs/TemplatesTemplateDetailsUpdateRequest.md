# TemplatesTemplateDetailsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OID** | **string** |  | 
**Validity** | **string** |  | 
**AvailableCAs** | Pointer to **[]string** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]TemplatesTemplateExtendedKeyUsageRequest**](TemplatesTemplateExtendedKeyUsageRequest.md) |  | [optional] 
**CertificatePolicies** | Pointer to [**[]TemplatesTemplateCertificatePolicyRequest**](TemplatesTemplateCertificatePolicyRequest.md) |  | [optional] 
**KeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeRequest**](TemplatesTemplateKeyTypeRequest.md) |  | [optional] 
**AlternativeKeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeRequest**](TemplatesTemplateKeyTypeRequest.md) |  | [optional] 
**CertificateExtensions** | Pointer to [**[]TemplatesTemplateCertificateExtensionRequest**](TemplatesTemplateCertificateExtensionRequest.md) |  | [optional] 

## Methods

### NewTemplatesTemplateDetailsUpdateRequest

`func NewTemplatesTemplateDetailsUpdateRequest(name string, oID string, validity string, ) *TemplatesTemplateDetailsUpdateRequest`

NewTemplatesTemplateDetailsUpdateRequest instantiates a new TemplatesTemplateDetailsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateDetailsUpdateRequestWithDefaults

`func NewTemplatesTemplateDetailsUpdateRequestWithDefaults() *TemplatesTemplateDetailsUpdateRequest`

NewTemplatesTemplateDetailsUpdateRequestWithDefaults instantiates a new TemplatesTemplateDetailsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplatesTemplateDetailsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatesTemplateDetailsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOID

`func (o *TemplatesTemplateDetailsUpdateRequest) GetOID() string`

GetOID returns the OID field if non-nil, zero value otherwise.

### GetOIDOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetOIDOk() (*string, bool)`

GetOIDOk returns a tuple with the OID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOID

`func (o *TemplatesTemplateDetailsUpdateRequest) SetOID(v string)`

SetOID sets OID field to given value.


### GetValidity

`func (o *TemplatesTemplateDetailsUpdateRequest) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *TemplatesTemplateDetailsUpdateRequest) SetValidity(v string)`

SetValidity sets Validity field to given value.


### GetAvailableCAs

`func (o *TemplatesTemplateDetailsUpdateRequest) GetAvailableCAs() []string`

GetAvailableCAs returns the AvailableCAs field if non-nil, zero value otherwise.

### GetAvailableCAsOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetAvailableCAsOk() (*[]string, bool)`

GetAvailableCAsOk returns a tuple with the AvailableCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCAs

`func (o *TemplatesTemplateDetailsUpdateRequest) SetAvailableCAs(v []string)`

SetAvailableCAs sets AvailableCAs field to given value.

### HasAvailableCAs

`func (o *TemplatesTemplateDetailsUpdateRequest) HasAvailableCAs() bool`

HasAvailableCAs returns a boolean if a field has been set.

### SetAvailableCAsNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetAvailableCAsNil(b bool)`

 SetAvailableCAsNil sets the value for AvailableCAs to be an explicit nil

### UnsetAvailableCAs
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetAvailableCAs()`

UnsetAvailableCAs ensures that no value is present for AvailableCAs, not even an explicit nil
### GetKeyUsage

`func (o *TemplatesTemplateDetailsUpdateRequest) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateDetailsUpdateRequest) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateDetailsUpdateRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsUpdateRequest) GetExtendedKeyUsages() []TemplatesTemplateExtendedKeyUsageRequest`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetExtendedKeyUsagesOk() (*[]TemplatesTemplateExtendedKeyUsageRequest, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsUpdateRequest) SetExtendedKeyUsages(v []TemplatesTemplateExtendedKeyUsageRequest)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *TemplatesTemplateDetailsUpdateRequest) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetCertificatePolicies

`func (o *TemplatesTemplateDetailsUpdateRequest) GetCertificatePolicies() []TemplatesTemplateCertificatePolicyRequest`

GetCertificatePolicies returns the CertificatePolicies field if non-nil, zero value otherwise.

### GetCertificatePoliciesOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetCertificatePoliciesOk() (*[]TemplatesTemplateCertificatePolicyRequest, bool)`

GetCertificatePoliciesOk returns a tuple with the CertificatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicies

`func (o *TemplatesTemplateDetailsUpdateRequest) SetCertificatePolicies(v []TemplatesTemplateCertificatePolicyRequest)`

SetCertificatePolicies sets CertificatePolicies field to given value.

### HasCertificatePolicies

`func (o *TemplatesTemplateDetailsUpdateRequest) HasCertificatePolicies() bool`

HasCertificatePolicies returns a boolean if a field has been set.

### SetCertificatePoliciesNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetCertificatePoliciesNil(b bool)`

 SetCertificatePoliciesNil sets the value for CertificatePolicies to be an explicit nil

### UnsetCertificatePolicies
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetCertificatePolicies()`

UnsetCertificatePolicies ensures that no value is present for CertificatePolicies, not even an explicit nil
### GetKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) GetKeyInformation() []TemplatesTemplateKeyTypeRequest`

GetKeyInformation returns the KeyInformation field if non-nil, zero value otherwise.

### GetKeyInformationOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetKeyInformationOk() (*[]TemplatesTemplateKeyTypeRequest, bool)`

GetKeyInformationOk returns a tuple with the KeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) SetKeyInformation(v []TemplatesTemplateKeyTypeRequest)`

SetKeyInformation sets KeyInformation field to given value.

### HasKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) HasKeyInformation() bool`

HasKeyInformation returns a boolean if a field has been set.

### SetKeyInformationNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetKeyInformationNil(b bool)`

 SetKeyInformationNil sets the value for KeyInformation to be an explicit nil

### UnsetKeyInformation
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetKeyInformation()`

UnsetKeyInformation ensures that no value is present for KeyInformation, not even an explicit nil
### GetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) GetAlternativeKeyInformation() []TemplatesTemplateKeyTypeRequest`

GetAlternativeKeyInformation returns the AlternativeKeyInformation field if non-nil, zero value otherwise.

### GetAlternativeKeyInformationOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetAlternativeKeyInformationOk() (*[]TemplatesTemplateKeyTypeRequest, bool)`

GetAlternativeKeyInformationOk returns a tuple with the AlternativeKeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) SetAlternativeKeyInformation(v []TemplatesTemplateKeyTypeRequest)`

SetAlternativeKeyInformation sets AlternativeKeyInformation field to given value.

### HasAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsUpdateRequest) HasAlternativeKeyInformation() bool`

HasAlternativeKeyInformation returns a boolean if a field has been set.

### SetAlternativeKeyInformationNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetAlternativeKeyInformationNil(b bool)`

 SetAlternativeKeyInformationNil sets the value for AlternativeKeyInformation to be an explicit nil

### UnsetAlternativeKeyInformation
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetAlternativeKeyInformation()`

UnsetAlternativeKeyInformation ensures that no value is present for AlternativeKeyInformation, not even an explicit nil
### GetCertificateExtensions

`func (o *TemplatesTemplateDetailsUpdateRequest) GetCertificateExtensions() []TemplatesTemplateCertificateExtensionRequest`

GetCertificateExtensions returns the CertificateExtensions field if non-nil, zero value otherwise.

### GetCertificateExtensionsOk

`func (o *TemplatesTemplateDetailsUpdateRequest) GetCertificateExtensionsOk() (*[]TemplatesTemplateCertificateExtensionRequest, bool)`

GetCertificateExtensionsOk returns a tuple with the CertificateExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExtensions

`func (o *TemplatesTemplateDetailsUpdateRequest) SetCertificateExtensions(v []TemplatesTemplateCertificateExtensionRequest)`

SetCertificateExtensions sets CertificateExtensions field to given value.

### HasCertificateExtensions

`func (o *TemplatesTemplateDetailsUpdateRequest) HasCertificateExtensions() bool`

HasCertificateExtensions returns a boolean if a field has been set.

### SetCertificateExtensionsNil

`func (o *TemplatesTemplateDetailsUpdateRequest) SetCertificateExtensionsNil(b bool)`

 SetCertificateExtensionsNil sets the value for CertificateExtensions to be an explicit nil

### UnsetCertificateExtensions
`func (o *TemplatesTemplateDetailsUpdateRequest) UnsetCertificateExtensions()`

UnsetCertificateExtensions ensures that no value is present for CertificateExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


