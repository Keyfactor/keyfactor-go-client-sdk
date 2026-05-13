# TemplatesTemplatePolicyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **int32** |  | [optional] 
**AllowKeyReuse** | Pointer to **NullableBool** |  | [optional] 
**AllowWildcards** | Pointer to **NullableBool** |  | [optional] 
**RFCEnforcement** | Pointer to **NullableBool** |  | [optional] 
**CertificateOwnerRole** | Pointer to [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | [optional] 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**PrimaryKeyAlgorithms** | Pointer to [**[]EnrollmentPatternsAlgorithmsAlgorithmDataResponse**](EnrollmentPatternsAlgorithmsAlgorithmDataResponse.md) |  | [optional] 
**AlternativeKeyAlgorithms** | Pointer to [**[]EnrollmentPatternsAlgorithmsAlgorithmDataResponse**](EnrollmentPatternsAlgorithmsAlgorithmDataResponse.md) |  | [optional] 

## Methods

### NewTemplatesTemplatePolicyResponseModel

`func NewTemplatesTemplatePolicyResponseModel() *TemplatesTemplatePolicyResponseModel`

NewTemplatesTemplatePolicyResponseModel instantiates a new TemplatesTemplatePolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplatePolicyResponseModelWithDefaults

`func NewTemplatesTemplatePolicyResponseModelWithDefaults() *TemplatesTemplatePolicyResponseModel`

NewTemplatesTemplatePolicyResponseModelWithDefaults instantiates a new TemplatesTemplatePolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplatesTemplatePolicyResponseModel) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatesTemplatePolicyResponseModel) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatesTemplatePolicyResponseModel) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplatesTemplatePolicyResponseModel) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *TemplatesTemplatePolicyResponseModel) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *TemplatesTemplatePolicyResponseModel) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *TemplatesTemplatePolicyResponseModel) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *TemplatesTemplatePolicyResponseModel) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### SetAllowKeyReuseNil

`func (o *TemplatesTemplatePolicyResponseModel) SetAllowKeyReuseNil(b bool)`

 SetAllowKeyReuseNil sets the value for AllowKeyReuse to be an explicit nil

### UnsetAllowKeyReuse
`func (o *TemplatesTemplatePolicyResponseModel) UnsetAllowKeyReuse()`

UnsetAllowKeyReuse ensures that no value is present for AllowKeyReuse, not even an explicit nil
### GetAllowWildcards

`func (o *TemplatesTemplatePolicyResponseModel) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *TemplatesTemplatePolicyResponseModel) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *TemplatesTemplatePolicyResponseModel) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *TemplatesTemplatePolicyResponseModel) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### SetAllowWildcardsNil

`func (o *TemplatesTemplatePolicyResponseModel) SetAllowWildcardsNil(b bool)`

 SetAllowWildcardsNil sets the value for AllowWildcards to be an explicit nil

### UnsetAllowWildcards
`func (o *TemplatesTemplatePolicyResponseModel) UnsetAllowWildcards()`

UnsetAllowWildcards ensures that no value is present for AllowWildcards, not even an explicit nil
### GetRFCEnforcement

`func (o *TemplatesTemplatePolicyResponseModel) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesTemplatePolicyResponseModel) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesTemplatePolicyResponseModel) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *TemplatesTemplatePolicyResponseModel) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### SetRFCEnforcementNil

`func (o *TemplatesTemplatePolicyResponseModel) SetRFCEnforcementNil(b bool)`

 SetRFCEnforcementNil sets the value for RFCEnforcement to be an explicit nil

### UnsetRFCEnforcement
`func (o *TemplatesTemplatePolicyResponseModel) UnsetRFCEnforcement()`

UnsetRFCEnforcement ensures that no value is present for RFCEnforcement, not even an explicit nil
### GetCertificateOwnerRole

`func (o *TemplatesTemplatePolicyResponseModel) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *TemplatesTemplatePolicyResponseModel) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *TemplatesTemplatePolicyResponseModel) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *TemplatesTemplatePolicyResponseModel) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyResponseModel) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *TemplatesTemplatePolicyResponseModel) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyResponseModel) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyResponseModel) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *TemplatesTemplatePolicyResponseModel) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *TemplatesTemplatePolicyResponseModel) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyResponseModel) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *TemplatesTemplatePolicyResponseModel) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyResponseModel) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyResponseModel) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *TemplatesTemplatePolicyResponseModel) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *TemplatesTemplatePolicyResponseModel) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) GetPrimaryKeyAlgorithms() []EnrollmentPatternsAlgorithmsAlgorithmDataResponse`

GetPrimaryKeyAlgorithms returns the PrimaryKeyAlgorithms field if non-nil, zero value otherwise.

### GetPrimaryKeyAlgorithmsOk

`func (o *TemplatesTemplatePolicyResponseModel) GetPrimaryKeyAlgorithmsOk() (*[]EnrollmentPatternsAlgorithmsAlgorithmDataResponse, bool)`

GetPrimaryKeyAlgorithmsOk returns a tuple with the PrimaryKeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) SetPrimaryKeyAlgorithms(v []EnrollmentPatternsAlgorithmsAlgorithmDataResponse)`

SetPrimaryKeyAlgorithms sets PrimaryKeyAlgorithms field to given value.

### HasPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) HasPrimaryKeyAlgorithms() bool`

HasPrimaryKeyAlgorithms returns a boolean if a field has been set.

### SetPrimaryKeyAlgorithmsNil

`func (o *TemplatesTemplatePolicyResponseModel) SetPrimaryKeyAlgorithmsNil(b bool)`

 SetPrimaryKeyAlgorithmsNil sets the value for PrimaryKeyAlgorithms to be an explicit nil

### UnsetPrimaryKeyAlgorithms
`func (o *TemplatesTemplatePolicyResponseModel) UnsetPrimaryKeyAlgorithms()`

UnsetPrimaryKeyAlgorithms ensures that no value is present for PrimaryKeyAlgorithms, not even an explicit nil
### GetAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) GetAlternativeKeyAlgorithms() []EnrollmentPatternsAlgorithmsAlgorithmDataResponse`

GetAlternativeKeyAlgorithms returns the AlternativeKeyAlgorithms field if non-nil, zero value otherwise.

### GetAlternativeKeyAlgorithmsOk

`func (o *TemplatesTemplatePolicyResponseModel) GetAlternativeKeyAlgorithmsOk() (*[]EnrollmentPatternsAlgorithmsAlgorithmDataResponse, bool)`

GetAlternativeKeyAlgorithmsOk returns a tuple with the AlternativeKeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) SetAlternativeKeyAlgorithms(v []EnrollmentPatternsAlgorithmsAlgorithmDataResponse)`

SetAlternativeKeyAlgorithms sets AlternativeKeyAlgorithms field to given value.

### HasAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyResponseModel) HasAlternativeKeyAlgorithms() bool`

HasAlternativeKeyAlgorithms returns a boolean if a field has been set.

### SetAlternativeKeyAlgorithmsNil

`func (o *TemplatesTemplatePolicyResponseModel) SetAlternativeKeyAlgorithmsNil(b bool)`

 SetAlternativeKeyAlgorithmsNil sets the value for AlternativeKeyAlgorithms to be an explicit nil

### UnsetAlternativeKeyAlgorithms
`func (o *TemplatesTemplatePolicyResponseModel) UnsetAlternativeKeyAlgorithms()`

UnsetAlternativeKeyAlgorithms ensures that no value is present for AlternativeKeyAlgorithms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


