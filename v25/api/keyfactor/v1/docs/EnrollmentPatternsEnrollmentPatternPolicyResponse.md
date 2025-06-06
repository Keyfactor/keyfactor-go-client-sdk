# EnrollmentPatternsEnrollmentPatternPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowKeyReuse** | Pointer to **NullableBool** | Whether or not keys can be reused. | [optional] 
**AllowWildcards** | Pointer to **NullableBool** | Whether or not wildcards can be used. | [optional] 
**RFCEnforcement** | Pointer to **NullableBool** | Whether or not RFC 2818 compliance should be enforced. | [optional] 
**CertificateOwnerRole** | Pointer to [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | [optional] 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** | The id of the security role that should be set as the owner of the cert during import of new certificates | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** | The name of the security role that should be set as the owner of the cert during import of new certificates | [optional] 
**KeyInfo** | Pointer to [**EnrollmentPatternsAlgorithmsKeyInfoResponse**](EnrollmentPatternsAlgorithmsKeyInfoResponse.md) |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternPolicyResponse

`func NewEnrollmentPatternsEnrollmentPatternPolicyResponse() *EnrollmentPatternsEnrollmentPatternPolicyResponse`

NewEnrollmentPatternsEnrollmentPatternPolicyResponse instantiates a new EnrollmentPatternsEnrollmentPatternPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternPolicyResponseWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternPolicyResponseWithDefaults() *EnrollmentPatternsEnrollmentPatternPolicyResponse`

NewEnrollmentPatternsEnrollmentPatternPolicyResponseWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### SetAllowKeyReuseNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetAllowKeyReuseNil(b bool)`

 SetAllowKeyReuseNil sets the value for AllowKeyReuse to be an explicit nil

### UnsetAllowKeyReuse
`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) UnsetAllowKeyReuse()`

UnsetAllowKeyReuse ensures that no value is present for AllowKeyReuse, not even an explicit nil
### GetAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### SetAllowWildcardsNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetAllowWildcardsNil(b bool)`

 SetAllowWildcardsNil sets the value for AllowWildcards to be an explicit nil

### UnsetAllowWildcards
`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) UnsetAllowWildcards()`

UnsetAllowWildcards ensures that no value is present for AllowWildcards, not even an explicit nil
### GetRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### SetRFCEnforcementNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetRFCEnforcementNil(b bool)`

 SetRFCEnforcementNil sets the value for RFCEnforcement to be an explicit nil

### UnsetRFCEnforcement
`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) UnsetRFCEnforcement()`

UnsetRFCEnforcement ensures that no value is present for RFCEnforcement, not even an explicit nil
### GetCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetKeyInfo() EnrollmentPatternsAlgorithmsKeyInfoResponse`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) GetKeyInfoOk() (*EnrollmentPatternsAlgorithmsKeyInfoResponse, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) SetKeyInfo(v EnrollmentPatternsAlgorithmsKeyInfoResponse)`

SetKeyInfo sets KeyInfo field to given value.

### HasKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyResponse) HasKeyInfo() bool`

HasKeyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


