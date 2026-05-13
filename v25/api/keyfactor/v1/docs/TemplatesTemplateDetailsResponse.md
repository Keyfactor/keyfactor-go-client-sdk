# TemplatesTemplateDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**OID** | Pointer to **NullableString** |  | [optional] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**Validity** | Pointer to **string** |  | [optional] 
**AvailableCAs** | Pointer to **[]string** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]TemplatesTemplateExtendedKeyUsageResponse**](TemplatesTemplateExtendedKeyUsageResponse.md) |  | [optional] 
**CertificatePolicies** | Pointer to [**[]TemplatesTemplateCertificatePolicyResponse**](TemplatesTemplateCertificatePolicyResponse.md) |  | [optional] 
**KeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeResponse**](TemplatesTemplateKeyTypeResponse.md) |  | [optional] 
**AlternativeKeyInformation** | Pointer to [**[]TemplatesTemplateKeyTypeResponse**](TemplatesTemplateKeyTypeResponse.md) |  | [optional] 
**CertificateExtensions** | Pointer to [**[]TemplatesTemplateCertificateExtensionResponse**](TemplatesTemplateCertificateExtensionResponse.md) |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewTemplatesTemplateDetailsResponse

`func NewTemplatesTemplateDetailsResponse() *TemplatesTemplateDetailsResponse`

NewTemplatesTemplateDetailsResponse instantiates a new TemplatesTemplateDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateDetailsResponseWithDefaults

`func NewTemplatesTemplateDetailsResponseWithDefaults() *TemplatesTemplateDetailsResponse`

NewTemplatesTemplateDetailsResponseWithDefaults instantiates a new TemplatesTemplateDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplatesTemplateDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatesTemplateDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatesTemplateDetailsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplatesTemplateDetailsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplatesTemplateDetailsResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplatesTemplateDetailsResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOID

`func (o *TemplatesTemplateDetailsResponse) GetOID() string`

GetOID returns the OID field if non-nil, zero value otherwise.

### GetOIDOk

`func (o *TemplatesTemplateDetailsResponse) GetOIDOk() (*string, bool)`

GetOIDOk returns a tuple with the OID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOID

`func (o *TemplatesTemplateDetailsResponse) SetOID(v string)`

SetOID sets OID field to given value.

### HasOID

`func (o *TemplatesTemplateDetailsResponse) HasOID() bool`

HasOID returns a boolean if a field has been set.

### SetOIDNil

`func (o *TemplatesTemplateDetailsResponse) SetOIDNil(b bool)`

 SetOIDNil sets the value for OID to be an explicit nil

### UnsetOID
`func (o *TemplatesTemplateDetailsResponse) UnsetOID()`

UnsetOID ensures that no value is present for OID, not even an explicit nil
### GetConfigurationTenant

`func (o *TemplatesTemplateDetailsResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *TemplatesTemplateDetailsResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *TemplatesTemplateDetailsResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *TemplatesTemplateDetailsResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *TemplatesTemplateDetailsResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *TemplatesTemplateDetailsResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetValidity

`func (o *TemplatesTemplateDetailsResponse) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *TemplatesTemplateDetailsResponse) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *TemplatesTemplateDetailsResponse) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *TemplatesTemplateDetailsResponse) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetAvailableCAs

`func (o *TemplatesTemplateDetailsResponse) GetAvailableCAs() []string`

GetAvailableCAs returns the AvailableCAs field if non-nil, zero value otherwise.

### GetAvailableCAsOk

`func (o *TemplatesTemplateDetailsResponse) GetAvailableCAsOk() (*[]string, bool)`

GetAvailableCAsOk returns a tuple with the AvailableCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCAs

`func (o *TemplatesTemplateDetailsResponse) SetAvailableCAs(v []string)`

SetAvailableCAs sets AvailableCAs field to given value.

### HasAvailableCAs

`func (o *TemplatesTemplateDetailsResponse) HasAvailableCAs() bool`

HasAvailableCAs returns a boolean if a field has been set.

### SetAvailableCAsNil

`func (o *TemplatesTemplateDetailsResponse) SetAvailableCAsNil(b bool)`

 SetAvailableCAsNil sets the value for AvailableCAs to be an explicit nil

### UnsetAvailableCAs
`func (o *TemplatesTemplateDetailsResponse) UnsetAvailableCAs()`

UnsetAvailableCAs ensures that no value is present for AvailableCAs, not even an explicit nil
### GetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsResponse) GetExtendedKeyUsages() []TemplatesTemplateExtendedKeyUsageResponse`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *TemplatesTemplateDetailsResponse) GetExtendedKeyUsagesOk() (*[]TemplatesTemplateExtendedKeyUsageResponse, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *TemplatesTemplateDetailsResponse) SetExtendedKeyUsages(v []TemplatesTemplateExtendedKeyUsageResponse)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *TemplatesTemplateDetailsResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *TemplatesTemplateDetailsResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *TemplatesTemplateDetailsResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetCertificatePolicies

`func (o *TemplatesTemplateDetailsResponse) GetCertificatePolicies() []TemplatesTemplateCertificatePolicyResponse`

GetCertificatePolicies returns the CertificatePolicies field if non-nil, zero value otherwise.

### GetCertificatePoliciesOk

`func (o *TemplatesTemplateDetailsResponse) GetCertificatePoliciesOk() (*[]TemplatesTemplateCertificatePolicyResponse, bool)`

GetCertificatePoliciesOk returns a tuple with the CertificatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicies

`func (o *TemplatesTemplateDetailsResponse) SetCertificatePolicies(v []TemplatesTemplateCertificatePolicyResponse)`

SetCertificatePolicies sets CertificatePolicies field to given value.

### HasCertificatePolicies

`func (o *TemplatesTemplateDetailsResponse) HasCertificatePolicies() bool`

HasCertificatePolicies returns a boolean if a field has been set.

### SetCertificatePoliciesNil

`func (o *TemplatesTemplateDetailsResponse) SetCertificatePoliciesNil(b bool)`

 SetCertificatePoliciesNil sets the value for CertificatePolicies to be an explicit nil

### UnsetCertificatePolicies
`func (o *TemplatesTemplateDetailsResponse) UnsetCertificatePolicies()`

UnsetCertificatePolicies ensures that no value is present for CertificatePolicies, not even an explicit nil
### GetKeyInformation

`func (o *TemplatesTemplateDetailsResponse) GetKeyInformation() []TemplatesTemplateKeyTypeResponse`

GetKeyInformation returns the KeyInformation field if non-nil, zero value otherwise.

### GetKeyInformationOk

`func (o *TemplatesTemplateDetailsResponse) GetKeyInformationOk() (*[]TemplatesTemplateKeyTypeResponse, bool)`

GetKeyInformationOk returns a tuple with the KeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInformation

`func (o *TemplatesTemplateDetailsResponse) SetKeyInformation(v []TemplatesTemplateKeyTypeResponse)`

SetKeyInformation sets KeyInformation field to given value.

### HasKeyInformation

`func (o *TemplatesTemplateDetailsResponse) HasKeyInformation() bool`

HasKeyInformation returns a boolean if a field has been set.

### SetKeyInformationNil

`func (o *TemplatesTemplateDetailsResponse) SetKeyInformationNil(b bool)`

 SetKeyInformationNil sets the value for KeyInformation to be an explicit nil

### UnsetKeyInformation
`func (o *TemplatesTemplateDetailsResponse) UnsetKeyInformation()`

UnsetKeyInformation ensures that no value is present for KeyInformation, not even an explicit nil
### GetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsResponse) GetAlternativeKeyInformation() []TemplatesTemplateKeyTypeResponse`

GetAlternativeKeyInformation returns the AlternativeKeyInformation field if non-nil, zero value otherwise.

### GetAlternativeKeyInformationOk

`func (o *TemplatesTemplateDetailsResponse) GetAlternativeKeyInformationOk() (*[]TemplatesTemplateKeyTypeResponse, bool)`

GetAlternativeKeyInformationOk returns a tuple with the AlternativeKeyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsResponse) SetAlternativeKeyInformation(v []TemplatesTemplateKeyTypeResponse)`

SetAlternativeKeyInformation sets AlternativeKeyInformation field to given value.

### HasAlternativeKeyInformation

`func (o *TemplatesTemplateDetailsResponse) HasAlternativeKeyInformation() bool`

HasAlternativeKeyInformation returns a boolean if a field has been set.

### SetAlternativeKeyInformationNil

`func (o *TemplatesTemplateDetailsResponse) SetAlternativeKeyInformationNil(b bool)`

 SetAlternativeKeyInformationNil sets the value for AlternativeKeyInformation to be an explicit nil

### UnsetAlternativeKeyInformation
`func (o *TemplatesTemplateDetailsResponse) UnsetAlternativeKeyInformation()`

UnsetAlternativeKeyInformation ensures that no value is present for AlternativeKeyInformation, not even an explicit nil
### GetCertificateExtensions

`func (o *TemplatesTemplateDetailsResponse) GetCertificateExtensions() []TemplatesTemplateCertificateExtensionResponse`

GetCertificateExtensions returns the CertificateExtensions field if non-nil, zero value otherwise.

### GetCertificateExtensionsOk

`func (o *TemplatesTemplateDetailsResponse) GetCertificateExtensionsOk() (*[]TemplatesTemplateCertificateExtensionResponse, bool)`

GetCertificateExtensionsOk returns a tuple with the CertificateExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExtensions

`func (o *TemplatesTemplateDetailsResponse) SetCertificateExtensions(v []TemplatesTemplateCertificateExtensionResponse)`

SetCertificateExtensions sets CertificateExtensions field to given value.

### HasCertificateExtensions

`func (o *TemplatesTemplateDetailsResponse) HasCertificateExtensions() bool`

HasCertificateExtensions returns a boolean if a field has been set.

### SetCertificateExtensionsNil

`func (o *TemplatesTemplateDetailsResponse) SetCertificateExtensionsNil(b bool)`

 SetCertificateExtensionsNil sets the value for CertificateExtensions to be an explicit nil

### UnsetCertificateExtensions
`func (o *TemplatesTemplateDetailsResponse) UnsetCertificateExtensions()`

UnsetCertificateExtensions ensures that no value is present for CertificateExtensions, not even an explicit nil
### GetKeyUsage

`func (o *TemplatesTemplateDetailsResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateDetailsResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateDetailsResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateDetailsResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


