# EnrollmentPatternsEnrollmentPatternPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowKeyReuse** | Pointer to **NullableBool** |  | [optional] 
**AllowWildcards** | Pointer to **NullableBool** |  | [optional] 
**RFCEnforcement** | Pointer to **NullableBool** |  | [optional] 
**CertificateOwnerRole** | Pointer to [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | [optional] 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**KeyInfo** | Pointer to [**EnrollmentPatternsAlgorithmsKeyInfoRequest**](EnrollmentPatternsAlgorithmsKeyInfoRequest.md) |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternPolicyRequest

`func NewEnrollmentPatternsEnrollmentPatternPolicyRequest() *EnrollmentPatternsEnrollmentPatternPolicyRequest`

NewEnrollmentPatternsEnrollmentPatternPolicyRequest instantiates a new EnrollmentPatternsEnrollmentPatternPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternPolicyRequestWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternPolicyRequestWithDefaults() *EnrollmentPatternsEnrollmentPatternPolicyRequest`

NewEnrollmentPatternsEnrollmentPatternPolicyRequestWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.

### HasAllowKeyReuse

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasAllowKeyReuse() bool`

HasAllowKeyReuse returns a boolean if a field has been set.

### SetAllowKeyReuseNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetAllowKeyReuseNil(b bool)`

 SetAllowKeyReuseNil sets the value for AllowKeyReuse to be an explicit nil

### UnsetAllowKeyReuse
`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) UnsetAllowKeyReuse()`

UnsetAllowKeyReuse ensures that no value is present for AllowKeyReuse, not even an explicit nil
### GetAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.

### HasAllowWildcards

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasAllowWildcards() bool`

HasAllowWildcards returns a boolean if a field has been set.

### SetAllowWildcardsNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetAllowWildcardsNil(b bool)`

 SetAllowWildcardsNil sets the value for AllowWildcards to be an explicit nil

### UnsetAllowWildcards
`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) UnsetAllowWildcards()`

UnsetAllowWildcards ensures that no value is present for AllowWildcards, not even an explicit nil
### GetRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### SetRFCEnforcementNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetRFCEnforcementNil(b bool)`

 SetRFCEnforcementNil sets the value for RFCEnforcement to be an explicit nil

### UnsetRFCEnforcement
`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) UnsetRFCEnforcement()`

UnsetRFCEnforcement ensures that no value is present for RFCEnforcement, not even an explicit nil
### GetCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.

### HasCertificateOwnerRole

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasCertificateOwnerRole() bool`

HasCertificateOwnerRole returns a boolean if a field has been set.

### GetDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetKeyInfo() EnrollmentPatternsAlgorithmsKeyInfoRequest`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) GetKeyInfoOk() (*EnrollmentPatternsAlgorithmsKeyInfoRequest, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) SetKeyInfo(v EnrollmentPatternsAlgorithmsKeyInfoRequest)`

SetKeyInfo sets KeyInfo field to given value.

### HasKeyInfo

`func (o *EnrollmentPatternsEnrollmentPatternPolicyRequest) HasKeyInfo() bool`

HasKeyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


