# EnrollmentPatternsEnrollmentPatternRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**TemplateDefault** | Pointer to **bool** |  | [optional] 
**AssociatedRoles** | Pointer to **[]string** |  | [optional] 
**UseADPermissions** | Pointer to **bool** |  | [optional] 
**CertificateAuthorities** | Pointer to **[]int32** |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**Regexes** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternRegexesRequest**](EnrollmentPatternsEnrollmentPatternRegexesRequest.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternMetadataFieldRequest**](EnrollmentPatternsEnrollmentPatternMetadataFieldRequest.md) |  | [optional] 
**RestrictCAs** | Pointer to **bool** |  | [optional] 
**Policies** | Pointer to [**EnrollmentPatternsEnrollmentPatternPolicyRequest**](EnrollmentPatternsEnrollmentPatternPolicyRequest.md) |  | [optional] 
**Defaults** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternDefaultRequest**](EnrollmentPatternsEnrollmentPatternDefaultRequest.md) |  | [optional] 
**EnrollmentFields** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternFieldRequest**](EnrollmentPatternsEnrollmentPatternFieldRequest.md) |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternRequest

`func NewEnrollmentPatternsEnrollmentPatternRequest(name string, ) *EnrollmentPatternsEnrollmentPatternRequest`

NewEnrollmentPatternsEnrollmentPatternRequest instantiates a new EnrollmentPatternsEnrollmentPatternRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternRequestWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternRequestWithDefaults() *EnrollmentPatternsEnrollmentPatternRequest`

NewEnrollmentPatternsEnrollmentPatternRequestWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetTemplateDefault() bool`

GetTemplateDefault returns the TemplateDefault field if non-nil, zero value otherwise.

### GetTemplateDefaultOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetTemplateDefaultOk() (*bool, bool)`

GetTemplateDefaultOk returns a tuple with the TemplateDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetTemplateDefault(v bool)`

SetTemplateDefault sets TemplateDefault field to given value.

### HasTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasTemplateDefault() bool`

HasTemplateDefault returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetAssociatedRoles() []string`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetAssociatedRolesOk() (*[]string, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetAssociatedRoles(v []string)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetUseADPermissions() bool`

GetUseADPermissions returns the UseADPermissions field if non-nil, zero value otherwise.

### GetUseADPermissionsOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetUseADPermissionsOk() (*bool, bool)`

GetUseADPermissionsOk returns a tuple with the UseADPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetUseADPermissions(v bool)`

SetUseADPermissions sets UseADPermissions field to given value.

### HasUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasUseADPermissions() bool`

HasUseADPermissions returns a boolean if a field has been set.

### GetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetCertificateAuthorities() []int32`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetCertificateAuthoritiesOk() (*[]int32, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetCertificateAuthorities(v []int32)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### SetCertificateAuthoritiesNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetCertificateAuthoritiesNil(b bool)`

 SetCertificateAuthoritiesNil sets the value for CertificateAuthorities to be an explicit nil

### UnsetCertificateAuthorities
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetCertificateAuthorities()`

UnsetCertificateAuthorities ensures that no value is present for CertificateAuthorities, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetRegexes() []EnrollmentPatternsEnrollmentPatternRegexesRequest`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetRegexesOk() (*[]EnrollmentPatternsEnrollmentPatternRegexesRequest, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetRegexes(v []EnrollmentPatternsEnrollmentPatternRegexesRequest)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### SetRegexesNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetRegexesNil(b bool)`

 SetRegexesNil sets the value for Regexes to be an explicit nil

### UnsetRegexes
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetRegexes()`

UnsetRegexes ensures that no value is present for Regexes, not even an explicit nil
### GetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetMetadataFields() []EnrollmentPatternsEnrollmentPatternMetadataFieldRequest`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetMetadataFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternMetadataFieldRequest, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetMetadataFields(v []EnrollmentPatternsEnrollmentPatternMetadataFieldRequest)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetRestrictCAs() bool`

GetRestrictCAs returns the RestrictCAs field if non-nil, zero value otherwise.

### GetRestrictCAsOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetRestrictCAsOk() (*bool, bool)`

GetRestrictCAsOk returns a tuple with the RestrictCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetRestrictCAs(v bool)`

SetRestrictCAs sets RestrictCAs field to given value.

### HasRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasRestrictCAs() bool`

HasRestrictCAs returns a boolean if a field has been set.

### GetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetPolicies() EnrollmentPatternsEnrollmentPatternPolicyRequest`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetPoliciesOk() (*EnrollmentPatternsEnrollmentPatternPolicyRequest, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetPolicies(v EnrollmentPatternsEnrollmentPatternPolicyRequest)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetDefaults() []EnrollmentPatternsEnrollmentPatternDefaultRequest`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetDefaultsOk() (*[]EnrollmentPatternsEnrollmentPatternDefaultRequest, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetDefaults(v []EnrollmentPatternsEnrollmentPatternDefaultRequest)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetEnrollmentFields() []EnrollmentPatternsEnrollmentPatternFieldRequest`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternRequest) GetEnrollmentFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternFieldRequest, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetEnrollmentFields(v []EnrollmentPatternsEnrollmentPatternFieldRequest)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternRequest) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternRequest) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *EnrollmentPatternsEnrollmentPatternRequest) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


