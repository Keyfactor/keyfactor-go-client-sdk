# TemplatesTemplatePolicyRequestModel

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
**PrimaryKeyAlgorithms** | Pointer to [**[]EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2**](EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2.md) |  | [optional] 
**AlternativeKeyAlgorithms** | Pointer to [**[]EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2**](EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2.md) |  | [optional] 

## Methods

### NewTemplatesTemplatePolicyRequestModel

`func NewTemplatesTemplatePolicyRequestModel() *TemplatesTemplatePolicyRequestModel`

NewTemplatesTemplatePolicyRequestModel instantiates a new TemplatesTemplatePolicyRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplatePolicyRequestModelWithDefaults

`func NewTemplatesTemplatePolicyRequestModelWithDefaults() *TemplatesTemplatePolicyRequestModel`

NewTemplatesTemplatePolicyRequestModelWithDefaults instantiates a new TemplatesTemplatePolicyRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplatesTemplatePolicyRequestModel) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatesTemplatePolicyRequestModel) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatesTemplatePolicyRequestModel) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplatesTemplatePolicyRequestModel) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *TemplatesTemplatePolicyRequestModel) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *TemplatesTemplatePolicyRequestModel) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *TemplatesTemplatePolicyRequestModel) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *TemplatesTemplatePolicyRequestModel) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### SetAllowKeyReuseNil

`func (o *TemplatesTemplatePolicyRequestModel) SetAllowKeyReuseNil(b bool)`

 SetAllowKeyReuseNil sets the value for AllowKeyReuse to be an explicit nil

### UnsetAllowKeyReuse
`func (o *TemplatesTemplatePolicyRequestModel) UnsetAllowKeyReuse()`

UnsetAllowKeyReuse ensures that no value is present for AllowKeyReuse, not even an explicit nil
### GetAllowWildcards

`func (o *TemplatesTemplatePolicyRequestModel) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *TemplatesTemplatePolicyRequestModel) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *TemplatesTemplatePolicyRequestModel) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *TemplatesTemplatePolicyRequestModel) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### SetAllowWildcardsNil

`func (o *TemplatesTemplatePolicyRequestModel) SetAllowWildcardsNil(b bool)`

 SetAllowWildcardsNil sets the value for AllowWildcards to be an explicit nil

### UnsetAllowWildcards
`func (o *TemplatesTemplatePolicyRequestModel) UnsetAllowWildcards()`

UnsetAllowWildcards ensures that no value is present for AllowWildcards, not even an explicit nil
### GetRFCEnforcement

`func (o *TemplatesTemplatePolicyRequestModel) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesTemplatePolicyRequestModel) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesTemplatePolicyRequestModel) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *TemplatesTemplatePolicyRequestModel) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### SetRFCEnforcementNil

`func (o *TemplatesTemplatePolicyRequestModel) SetRFCEnforcementNil(b bool)`

 SetRFCEnforcementNil sets the value for RFCEnforcement to be an explicit nil

### UnsetRFCEnforcement
`func (o *TemplatesTemplatePolicyRequestModel) UnsetRFCEnforcement()`

UnsetRFCEnforcement ensures that no value is present for RFCEnforcement, not even an explicit nil
### GetCertificateOwnerRole

`func (o *TemplatesTemplatePolicyRequestModel) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *TemplatesTemplatePolicyRequestModel) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *TemplatesTemplatePolicyRequestModel) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *TemplatesTemplatePolicyRequestModel) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyRequestModel) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *TemplatesTemplatePolicyRequestModel) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyRequestModel) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *TemplatesTemplatePolicyRequestModel) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *TemplatesTemplatePolicyRequestModel) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *TemplatesTemplatePolicyRequestModel) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyRequestModel) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *TemplatesTemplatePolicyRequestModel) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyRequestModel) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *TemplatesTemplatePolicyRequestModel) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *TemplatesTemplatePolicyRequestModel) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *TemplatesTemplatePolicyRequestModel) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) GetPrimaryKeyAlgorithms() []EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2`

GetPrimaryKeyAlgorithms returns the PrimaryKeyAlgorithms field if non-nil, zero value otherwise.

### GetPrimaryKeyAlgorithmsOk

`func (o *TemplatesTemplatePolicyRequestModel) GetPrimaryKeyAlgorithmsOk() (*[]EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2, bool)`

GetPrimaryKeyAlgorithmsOk returns a tuple with the PrimaryKeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) SetPrimaryKeyAlgorithms(v []EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2)`

SetPrimaryKeyAlgorithms sets PrimaryKeyAlgorithms field to given value.

### HasPrimaryKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) HasPrimaryKeyAlgorithms() bool`

HasPrimaryKeyAlgorithms returns a boolean if a field has been set.

### SetPrimaryKeyAlgorithmsNil

`func (o *TemplatesTemplatePolicyRequestModel) SetPrimaryKeyAlgorithmsNil(b bool)`

 SetPrimaryKeyAlgorithmsNil sets the value for PrimaryKeyAlgorithms to be an explicit nil

### UnsetPrimaryKeyAlgorithms
`func (o *TemplatesTemplatePolicyRequestModel) UnsetPrimaryKeyAlgorithms()`

UnsetPrimaryKeyAlgorithms ensures that no value is present for PrimaryKeyAlgorithms, not even an explicit nil
### GetAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) GetAlternativeKeyAlgorithms() []EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2`

GetAlternativeKeyAlgorithms returns the AlternativeKeyAlgorithms field if non-nil, zero value otherwise.

### GetAlternativeKeyAlgorithmsOk

`func (o *TemplatesTemplatePolicyRequestModel) GetAlternativeKeyAlgorithmsOk() (*[]EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2, bool)`

GetAlternativeKeyAlgorithmsOk returns a tuple with the AlternativeKeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) SetAlternativeKeyAlgorithms(v []EnrollmentPatternsAlgorithmsAlgorithmDataRequestV2)`

SetAlternativeKeyAlgorithms sets AlternativeKeyAlgorithms field to given value.

### HasAlternativeKeyAlgorithms

`func (o *TemplatesTemplatePolicyRequestModel) HasAlternativeKeyAlgorithms() bool`

HasAlternativeKeyAlgorithms returns a boolean if a field has been set.

### SetAlternativeKeyAlgorithmsNil

`func (o *TemplatesTemplatePolicyRequestModel) SetAlternativeKeyAlgorithmsNil(b bool)`

 SetAlternativeKeyAlgorithmsNil sets the value for AlternativeKeyAlgorithms to be an explicit nil

### UnsetAlternativeKeyAlgorithms
`func (o *TemplatesTemplatePolicyRequestModel) UnsetAlternativeKeyAlgorithms()`

UnsetAlternativeKeyAlgorithms ensures that no value is present for AlternativeKeyAlgorithms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


