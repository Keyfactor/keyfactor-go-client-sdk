# TemplatesGlobalGlobalTemplatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the global template policy. | [optional] 
**AllowKeyReuse** | **bool** | Whether or not keys can be reused. | 
**AllowWildcards** | **bool** | Whether or not wildcards can be used. | 
**RFCEnforcement** | **bool** | Whether or not RFC 2818 compliance should be enforced. | 
**CertificateOwnerRole** | [**CSSCMSCoreEnumsTemplateCertificateOwnerRole**](CSSCMSCoreEnumsTemplateCertificateOwnerRole.md) |  | 
**DefaultCertificateOwnerRoleId** | Pointer to **NullableInt32** | The id of the security role that should be set as the owner of the cert during import of new certificates | [optional] 
**DefaultCertificateOwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**KeyInfo** | [**CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo**](CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo.md) |  | 

## Methods

### NewTemplatesGlobalGlobalTemplatePolicyRequest

`func NewTemplatesGlobalGlobalTemplatePolicyRequest(allowKeyReuse bool, allowWildcards bool, rFCEnforcement bool, certificateOwnerRole CSSCMSCoreEnumsTemplateCertificateOwnerRole, keyInfo CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo, ) *TemplatesGlobalGlobalTemplatePolicyRequest`

NewTemplatesGlobalGlobalTemplatePolicyRequest instantiates a new TemplatesGlobalGlobalTemplatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults

`func NewTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults() *TemplatesGlobalGlobalTemplatePolicyRequest`

NewTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults instantiates a new TemplatesGlobalGlobalTemplatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.


### GetAllowWildcards

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.


### GetRFCEnforcement

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.


### GetCertificateOwnerRole

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetCertificateOwnerRole() CSSCMSCoreEnumsTemplateCertificateOwnerRole`

GetCertificateOwnerRole returns the CertificateOwnerRole field if non-nil, zero value otherwise.

### GetCertificateOwnerRoleOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetCertificateOwnerRoleOk() (*CSSCMSCoreEnumsTemplateCertificateOwnerRole, bool)`

GetCertificateOwnerRoleOk returns a tuple with the CertificateOwnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOwnerRole

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetCertificateOwnerRole(v CSSCMSCoreEnumsTemplateCertificateOwnerRole)`

SetCertificateOwnerRole sets CertificateOwnerRole field to given value.


### GetDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetDefaultCertificateOwnerRoleId() int32`

GetDefaultCertificateOwnerRoleId returns the DefaultCertificateOwnerRoleId field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleIdOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetDefaultCertificateOwnerRoleIdOk() (*int32, bool)`

GetDefaultCertificateOwnerRoleIdOk returns a tuple with the DefaultCertificateOwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetDefaultCertificateOwnerRoleId(v int32)`

SetDefaultCertificateOwnerRoleId sets DefaultCertificateOwnerRoleId field to given value.

### HasDefaultCertificateOwnerRoleId

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) HasDefaultCertificateOwnerRoleId() bool`

HasDefaultCertificateOwnerRoleId returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleIdNil

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetDefaultCertificateOwnerRoleIdNil(b bool)`

 SetDefaultCertificateOwnerRoleIdNil sets the value for DefaultCertificateOwnerRoleId to be an explicit nil

### UnsetDefaultCertificateOwnerRoleId
`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) UnsetDefaultCertificateOwnerRoleId()`

UnsetDefaultCertificateOwnerRoleId ensures that no value is present for DefaultCertificateOwnerRoleId, not even an explicit nil
### GetDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetDefaultCertificateOwnerRoleName() string`

GetDefaultCertificateOwnerRoleName returns the DefaultCertificateOwnerRoleName field if non-nil, zero value otherwise.

### GetDefaultCertificateOwnerRoleNameOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetDefaultCertificateOwnerRoleNameOk() (*string, bool)`

GetDefaultCertificateOwnerRoleNameOk returns a tuple with the DefaultCertificateOwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetDefaultCertificateOwnerRoleName(v string)`

SetDefaultCertificateOwnerRoleName sets DefaultCertificateOwnerRoleName field to given value.

### HasDefaultCertificateOwnerRoleName

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) HasDefaultCertificateOwnerRoleName() bool`

HasDefaultCertificateOwnerRoleName returns a boolean if a field has been set.

### SetDefaultCertificateOwnerRoleNameNil

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetDefaultCertificateOwnerRoleNameNil(b bool)`

 SetDefaultCertificateOwnerRoleNameNil sets the value for DefaultCertificateOwnerRoleName to be an explicit nil

### UnsetDefaultCertificateOwnerRoleName
`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) UnsetDefaultCertificateOwnerRoleName()`

UnsetDefaultCertificateOwnerRoleName ensures that no value is present for DefaultCertificateOwnerRoleName, not even an explicit nil
### GetKeyInfo

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetKeyInfo() CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) GetKeyInfoOk() (*CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *TemplatesGlobalGlobalTemplatePolicyRequest) SetKeyInfo(v CSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo)`

SetKeyInfo sets KeyInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


