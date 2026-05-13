# TemplatesTemplateDetailsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Validity** | **string** |  | 
**ConfigurationTenant** | **string** |  | 
**AvailableCAs** | Pointer to **[]string** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]TemplatesTemplateExtendedKeyUsageRequest**](TemplatesTemplateExtendedKeyUsageRequest.md) |  | [optional] 
**CertificatePolicies** | Pointer to [**[]TemplatesTemplateCertificatePolicyRequest**](TemplatesTemplateCertificatePolicyRequest.md) |  | [optional] 
**KeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeRequest**](TemplatesTemplateKeyTypeRequest.md) |  | [optional] 
**AlternativeKeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeRequest**](TemplatesTemplateKeyTypeRequest.md) |  | [optional] 
**CertificateExtensions** | Pointer to [**[]TemplatesTemplateCertificateExtensionRequest**](TemplatesTemplateCertificateExtensionRequest.md) |  | [optional] 

## Methods

### NewTemplatesTemplateDetailsCreateRequest

`func NewTemplatesTemplateDetailsCreateRequest(name string, validity string, configurationTenant string, ) *TemplatesTemplateDetailsCreateRequest`

NewTemplatesTemplateDetailsCreateRequest instantiates a new TemplatesTemplateDetailsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateDetailsCreateRequestWithDefaults

`func NewTemplatesTemplateDetailsCreateRequestWithDefaults() *TemplatesTemplateDetailsCreateRequest`

NewTemplatesTemplateDetailsCreateRequestWithDefaults instantiates a new TemplatesTemplateDetailsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplatesTemplateDetailsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatesTemplateDetailsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValidity

`func (o *TemplatesTemplateDetailsCreateRequest) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *TemplatesTemplateDetailsCreateRequest) SetValidity(v string)`

SetValidity sets Validity field to given value.


### GetConfigurationTenant

`func (o *TemplatesTemplateDetailsCreateRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *TemplatesTemplateDetailsCreateRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.


### GetAvailableCAs

`func (o *TemplatesTemplateDetailsCreateRequest) GetAvailableCAs() []string`

GetAvailableCAs returns the AvailableCAs field if non-nil, zero value otherwise.

### GetAvailableCAsOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetAvailableCAsOk() (*[]string, bool)`

GetAvailableCAsOk returns a tuple with the AvailableCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCAs

`func (o *TemplatesTemplateDetailsCreateRequest) SetAvailableCAs(v []string)`

SetAvailableCAs sets AvailableCAs field to given value.

### HasAvailableCAs

`func (o *TemplatesTemplateDetailsCreateRequest) HasAvailableCAs() bool`

HasAvailableCAs returns a boolean if a field has been set.

### SetAvailableCAsNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetAvailableCAsNil(b bool)`

 SetAvailableCAsNil sets the value for AvailableCAs to be an explicit nil

### UnsetAvailableCAs
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetAvailableCAs()`

UnsetAvailableCAs ensures that no value is present for AvailableCAs, not even an explicit nil
### GetKeyUsage

`func (o *TemplatesTemplateDetailsCreateRequest) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateDetailsCreateRequest) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateDetailsCreateRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsCreateRequest) GetExtendedKeyUsages() []TemplatesTemplateExtendedKeyUsageRequest`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetExtendedKeyUsagesOk() (*[]TemplatesTemplateExtendedKeyUsageRequest, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsCreateRequest) SetExtendedKeyUsages(v []TemplatesTemplateExtendedKeyUsageRequest)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *TemplatesTemplateDetailsCreateRequest) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetCertificatePolicies

`func (o *TemplatesTemplateDetailsCreateRequest) GetCertificatePolicies() []TemplatesTemplateCertificatePolicyRequest`

GetCertificatePolicies returns the CertificatePolicies field if non-nil, zero value otherwise.

### GetCertificatePoliciesOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetCertificatePoliciesOk() (*[]TemplatesTemplateCertificatePolicyRequest, bool)`

GetCertificatePoliciesOk returns a tuple with the CertificatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicies

`func (o *TemplatesTemplateDetailsCreateRequest) SetCertificatePolicies(v []TemplatesTemplateCertificatePolicyRequest)`

SetCertificatePolicies sets CertificatePolicies field to given value.

### HasCertificatePolicies

`func (o *TemplatesTemplateDetailsCreateRequest) HasCertificatePolicies() bool`

HasCertificatePolicies returns a boolean if a field has been set.

### SetCertificatePoliciesNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetCertificatePoliciesNil(b bool)`

 SetCertificatePoliciesNil sets the value for CertificatePolicies to be an explicit nil

### UnsetCertificatePolicies
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetCertificatePolicies()`

UnsetCertificatePolicies ensures that no value is present for CertificatePolicies, not even an explicit nil
### GetKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) GetKeyInformation() []TemplatesTemplateKeyTypeRequest`

GetKeyInformation returns the KeyInformation field if non-nil, zero value otherwise.

### GetKeyInformationOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetKeyInformationOk() (*[]TemplatesTemplateKeyTypeRequest, bool)`

GetKeyInformationOk returns a tuple with the KeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) SetKeyInformation(v []TemplatesTemplateKeyTypeRequest)`

SetKeyInformation sets KeyInformation field to given value.

### HasKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) HasKeyInformation() bool`

HasKeyInformation returns a boolean if a field has been set.

### SetKeyInformationNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetKeyInformationNil(b bool)`

 SetKeyInformationNil sets the value for KeyInformation to be an explicit nil

### UnsetKeyInformation
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetKeyInformation()`

UnsetKeyInformation ensures that no value is present for KeyInformation, not even an explicit nil
### GetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) GetAlternativeKeyInformation() []TemplatesTemplateKeyTypeRequest`

GetAlternativeKeyInformation returns the AlternativeKeyInformation field if non-nil, zero value otherwise.

### GetAlternativeKeyInformationOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetAlternativeKeyInformationOk() (*[]TemplatesTemplateKeyTypeRequest, bool)`

GetAlternativeKeyInformationOk returns a tuple with the AlternativeKeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) SetAlternativeKeyInformation(v []TemplatesTemplateKeyTypeRequest)`

SetAlternativeKeyInformation sets AlternativeKeyInformation field to given value.

### HasAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsCreateRequest) HasAlternativeKeyInformation() bool`

HasAlternativeKeyInformation returns a boolean if a field has been set.

### SetAlternativeKeyInformationNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetAlternativeKeyInformationNil(b bool)`

 SetAlternativeKeyInformationNil sets the value for AlternativeKeyInformation to be an explicit nil

### UnsetAlternativeKeyInformation
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetAlternativeKeyInformation()`

UnsetAlternativeKeyInformation ensures that no value is present for AlternativeKeyInformation, not even an explicit nil
### GetCertificateExtensions

`func (o *TemplatesTemplateDetailsCreateRequest) GetCertificateExtensions() []TemplatesTemplateCertificateExtensionRequest`

GetCertificateExtensions returns the CertificateExtensions field if non-nil, zero value otherwise.

### GetCertificateExtensionsOk

`func (o *TemplatesTemplateDetailsCreateRequest) GetCertificateExtensionsOk() (*[]TemplatesTemplateCertificateExtensionRequest, bool)`

GetCertificateExtensionsOk returns a tuple with the CertificateExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExtensions

`func (o *TemplatesTemplateDetailsCreateRequest) SetCertificateExtensions(v []TemplatesTemplateCertificateExtensionRequest)`

SetCertificateExtensions sets CertificateExtensions field to given value.

### HasCertificateExtensions

`func (o *TemplatesTemplateDetailsCreateRequest) HasCertificateExtensions() bool`

HasCertificateExtensions returns a boolean if a field has been set.

### SetCertificateExtensionsNil

`func (o *TemplatesTemplateDetailsCreateRequest) SetCertificateExtensionsNil(b bool)`

 SetCertificateExtensionsNil sets the value for CertificateExtensions to be an explicit nil

### UnsetCertificateExtensions
`func (o *TemplatesTemplateDetailsCreateRequest) UnsetCertificateExtensions()`

UnsetCertificateExtensions ensures that no value is present for CertificateExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


