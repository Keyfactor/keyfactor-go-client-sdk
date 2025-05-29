# TemplatesEnrollmentTemplateEnrollmentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowKeyReuse** | Pointer to **bool** | Whether or not keys can be reused. | [optional] 
**AllowWildcards** | Pointer to **bool** | Whether or not wildcards can be used. | [optional] 
**RFCEnforcement** | Pointer to **bool** | Whether or not RFC 2818 compliance should be enforced. | [optional] 
**KeyInfo** | Pointer to [**CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo**](CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo.md) |  | [optional] 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** | The id of the security role that should be set as the owner of the cert during import of new certificates | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**CertificateOwnerRole** | Pointer to [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | [optional] 

## Methods

### NewTemplatesEnrollmentTemplateEnrollmentPolicyResponse

`func NewTemplatesEnrollmentTemplateEnrollmentPolicyResponse() *TemplatesEnrollmentTemplateEnrollmentPolicyResponse`

NewTemplatesEnrollmentTemplateEnrollmentPolicyResponse instantiates a new TemplatesEnrollmentTemplateEnrollmentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesEnrollmentTemplateEnrollmentPolicyResponseWithDefaults

`func NewTemplatesEnrollmentTemplateEnrollmentPolicyResponseWithDefaults() *TemplatesEnrollmentTemplateEnrollmentPolicyResponse`

NewTemplatesEnrollmentTemplateEnrollmentPolicyResponseWithDefaults instantiates a new TemplatesEnrollmentTemplateEnrollmentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowKeyReuse

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetKeyInfo

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetKeyInfo() CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetKeyInfoOk() (*CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetKeyInfo(v CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo)`

SetKeyInfo sets KeyInfo field to given value.

### HasKeyInfo

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasKeyInfo() bool`

HasKeyInfo returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetCertificateOwnerRole

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *TemplatesEnrollmentTemplateEnrollmentPolicyResponse) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


