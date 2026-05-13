# TemplatesGlobalGlobalTemplatePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowKeyReuse** | Pointer to **bool** | Whether or not keys can be reused. | [optional] 
**AllowWildcards** | Pointer to **bool** | Whether or not wildcards can be used. | [optional] 
**RFCEnforcement** | Pointer to **bool** | Whether or not RFC 2818 compliance should be enforced. | [optional] 
**CertificateOwnerRole** | Pointer to [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | [optional] 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** | The id of the security role that should be set as the owner of the cert during import of new certificates | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**KeyInfo** | Pointer to [**EnrollmentPatternsAlgorithmsKeyInfoResponse**](EnrollmentPatternsAlgorithmsKeyInfoResponse.md) |  | [optional] 

## Methods

### NewTemplatesGlobalGlobalTemplatePolicyResponse

`func NewTemplatesGlobalGlobalTemplatePolicyResponse() *TemplatesGlobalGlobalTemplatePolicyResponse`

NewTemplatesGlobalGlobalTemplatePolicyResponse instantiates a new TemplatesGlobalGlobalTemplatePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesGlobalGlobalTemplatePolicyResponseWithDefaults

`func NewTemplatesGlobalGlobalTemplatePolicyResponseWithDefaults() *TemplatesGlobalGlobalTemplatePolicyResponse`

NewTemplatesGlobalGlobalTemplatePolicyResponseWithDefaults instantiates a new TemplatesGlobalGlobalTemplatePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowKeyReuse

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### GetAllowWildcards

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetCertificateOwnerRole

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetKeyInfo

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetKeyInfo() EnrollmentPatternsAlgorithmsKeyInfoResponse`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) GetKeyInfoOk() (*EnrollmentPatternsAlgorithmsKeyInfoResponse, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) SetKeyInfo(v EnrollmentPatternsAlgorithmsKeyInfoResponse)`

SetKeyInfo sets KeyInfo field to given value.

### HasKeyInfo

`func (o *TemplatesGlobalGlobalTemplatePolicyResponse) HasKeyInfo() bool`

HasKeyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


