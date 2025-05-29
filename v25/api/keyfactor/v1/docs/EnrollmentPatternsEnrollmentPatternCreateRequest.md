# EnrollmentPatternsEnrollmentPatternCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **int32** |  | 
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

### NewEnrollmentPatternsEnrollmentPatternCreateRequest

`func NewEnrollmentPatternsEnrollmentPatternCreateRequest(template int32, name string, ) *EnrollmentPatternsEnrollmentPatternCreateRequest`

NewEnrollmentPatternsEnrollmentPatternCreateRequest instantiates a new EnrollmentPatternsEnrollmentPatternCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternCreateRequestWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternCreateRequestWithDefaults() *EnrollmentPatternsEnrollmentPatternCreateRequest`

NewEnrollmentPatternsEnrollmentPatternCreateRequestWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetTemplate(v int32)`

SetTemplate sets Template field to given value.


### GetName

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetTemplateDefault() bool`

GetTemplateDefault returns the TemplateDefault field if non-nil, zero value otherwise.

### GetTemplateDefaultOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetTemplateDefaultOk() (*bool, bool)`

GetTemplateDefaultOk returns a tuple with the TemplateDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetTemplateDefault(v bool)`

SetTemplateDefault sets TemplateDefault field to given value.

### HasTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasTemplateDefault() bool`

HasTemplateDefault returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetAssociatedRoles() []string`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetAssociatedRolesOk() (*[]string, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetAssociatedRoles(v []string)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetUseADPermissions() bool`

GetUseADPermissions returns the UseADPermissions field if non-nil, zero value otherwise.

### GetUseADPermissionsOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetUseADPermissionsOk() (*bool, bool)`

GetUseADPermissionsOk returns a tuple with the UseADPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetUseADPermissions(v bool)`

SetUseADPermissions sets UseADPermissions field to given value.

### HasUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasUseADPermissions() bool`

HasUseADPermissions returns a boolean if a field has been set.

### GetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetCertificateAuthorities() []int32`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetCertificateAuthoritiesOk() (*[]int32, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetCertificateAuthorities(v []int32)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### SetCertificateAuthoritiesNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetCertificateAuthoritiesNil(b bool)`

 SetCertificateAuthoritiesNil sets the value for CertificateAuthorities to be an explicit nil

### UnsetCertificateAuthorities
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetCertificateAuthorities()`

UnsetCertificateAuthorities ensures that no value is present for CertificateAuthorities, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetRegexes() []EnrollmentPatternsEnrollmentPatternRegexesRequest`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetRegexesOk() (*[]EnrollmentPatternsEnrollmentPatternRegexesRequest, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetRegexes(v []EnrollmentPatternsEnrollmentPatternRegexesRequest)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### SetRegexesNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetRegexesNil(b bool)`

 SetRegexesNil sets the value for Regexes to be an explicit nil

### UnsetRegexes
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetRegexes()`

UnsetRegexes ensures that no value is present for Regexes, not even an explicit nil
### GetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetMetadataFields() []EnrollmentPatternsEnrollmentPatternMetadataFieldRequest`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetMetadataFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternMetadataFieldRequest, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetMetadataFields(v []EnrollmentPatternsEnrollmentPatternMetadataFieldRequest)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetRestrictCAs() bool`

GetRestrictCAs returns the RestrictCAs field if non-nil, zero value otherwise.

### GetRestrictCAsOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetRestrictCAsOk() (*bool, bool)`

GetRestrictCAsOk returns a tuple with the RestrictCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetRestrictCAs(v bool)`

SetRestrictCAs sets RestrictCAs field to given value.

### HasRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasRestrictCAs() bool`

HasRestrictCAs returns a boolean if a field has been set.

### GetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetPolicies() EnrollmentPatternsEnrollmentPatternPolicyRequest`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetPoliciesOk() (*EnrollmentPatternsEnrollmentPatternPolicyRequest, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetPolicies(v EnrollmentPatternsEnrollmentPatternPolicyRequest)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetDefaults() []EnrollmentPatternsEnrollmentPatternDefaultRequest`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetDefaultsOk() (*[]EnrollmentPatternsEnrollmentPatternDefaultRequest, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetDefaults(v []EnrollmentPatternsEnrollmentPatternDefaultRequest)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetEnrollmentFields() []EnrollmentPatternsEnrollmentPatternFieldRequest`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) GetEnrollmentFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternFieldRequest, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetEnrollmentFields(v []EnrollmentPatternsEnrollmentPatternFieldRequest)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *EnrollmentPatternsEnrollmentPatternCreateRequest) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


