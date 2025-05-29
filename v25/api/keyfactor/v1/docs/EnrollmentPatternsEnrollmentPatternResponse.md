# EnrollmentPatternsEnrollmentPatternResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to [**EnrollmentPatternsEnrollmentPatternTemplateResponse**](EnrollmentPatternsEnrollmentPatternTemplateResponse.md) |  | [optional] 
**TemplateDefault** | Pointer to **bool** |  | [optional] 
**UseADPermissions** | Pointer to **bool** |  | [optional] 
**AssociatedRoles** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternAssociatedRoleResponse**](EnrollmentPatternsEnrollmentPatternAssociatedRoleResponse.md) |  | [optional] 
**CertificateAuthorities** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternCAResponse**](EnrollmentPatternsEnrollmentPatternCAResponse.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to [**CSSCMSCoreEnumsEnrollmentType**](CSSCMSCoreEnumsEnrollmentType.md) |  | [optional] 
**Regexes** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternRegexesResponse**](EnrollmentPatternsEnrollmentPatternRegexesResponse.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternMetadataFieldResponse**](EnrollmentPatternsEnrollmentPatternMetadataFieldResponse.md) |  | [optional] 
**RestrictCAs** | Pointer to **bool** |  | [optional] 
**Policies** | Pointer to [**EnrollmentPatternsEnrollmentPatternPolicyResponse**](EnrollmentPatternsEnrollmentPatternPolicyResponse.md) |  | [optional] 
**Defaults** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternDefaultResponse**](EnrollmentPatternsEnrollmentPatternDefaultResponse.md) |  | [optional] 
**EnrollmentFields** | Pointer to [**[]EnrollmentPatternsEnrollmentPatternFieldResponse**](EnrollmentPatternsEnrollmentPatternFieldResponse.md) |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternResponse

`func NewEnrollmentPatternsEnrollmentPatternResponse() *EnrollmentPatternsEnrollmentPatternResponse`

NewEnrollmentPatternsEnrollmentPatternResponse instantiates a new EnrollmentPatternsEnrollmentPatternResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternResponseWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternResponseWithDefaults() *EnrollmentPatternsEnrollmentPatternResponse`

NewEnrollmentPatternsEnrollmentPatternResponseWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplate

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetTemplate() EnrollmentPatternsEnrollmentPatternTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetTemplateOk() (*EnrollmentPatternsEnrollmentPatternTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetTemplate(v EnrollmentPatternsEnrollmentPatternTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetTemplateDefault() bool`

GetTemplateDefault returns the TemplateDefault field if non-nil, zero value otherwise.

### GetTemplateDefaultOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetTemplateDefaultOk() (*bool, bool)`

GetTemplateDefaultOk returns a tuple with the TemplateDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetTemplateDefault(v bool)`

SetTemplateDefault sets TemplateDefault field to given value.

### HasTemplateDefault

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasTemplateDefault() bool`

HasTemplateDefault returns a boolean if a field has been set.

### GetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetUseADPermissions() bool`

GetUseADPermissions returns the UseADPermissions field if non-nil, zero value otherwise.

### GetUseADPermissionsOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetUseADPermissionsOk() (*bool, bool)`

GetUseADPermissionsOk returns a tuple with the UseADPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetUseADPermissions(v bool)`

SetUseADPermissions sets UseADPermissions field to given value.

### HasUseADPermissions

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasUseADPermissions() bool`

HasUseADPermissions returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetAssociatedRoles() []EnrollmentPatternsEnrollmentPatternAssociatedRoleResponse`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetAssociatedRolesOk() (*[]EnrollmentPatternsEnrollmentPatternAssociatedRoleResponse, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetAssociatedRoles(v []EnrollmentPatternsEnrollmentPatternAssociatedRoleResponse)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetCertificateAuthorities() []EnrollmentPatternsEnrollmentPatternCAResponse`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetCertificateAuthoritiesOk() (*[]EnrollmentPatternsEnrollmentPatternCAResponse, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetCertificateAuthorities(v []EnrollmentPatternsEnrollmentPatternCAResponse)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### SetCertificateAuthoritiesNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetCertificateAuthoritiesNil(b bool)`

 SetCertificateAuthoritiesNil sets the value for CertificateAuthorities to be an explicit nil

### UnsetCertificateAuthorities
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetCertificateAuthorities()`

UnsetCertificateAuthorities ensures that no value is present for CertificateAuthorities, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetRegexes() []EnrollmentPatternsEnrollmentPatternRegexesResponse`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetRegexesOk() (*[]EnrollmentPatternsEnrollmentPatternRegexesResponse, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetRegexes(v []EnrollmentPatternsEnrollmentPatternRegexesResponse)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### SetRegexesNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetRegexesNil(b bool)`

 SetRegexesNil sets the value for Regexes to be an explicit nil

### UnsetRegexes
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetRegexes()`

UnsetRegexes ensures that no value is present for Regexes, not even an explicit nil
### GetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetMetadataFields() []EnrollmentPatternsEnrollmentPatternMetadataFieldResponse`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetMetadataFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternMetadataFieldResponse, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetMetadataFields(v []EnrollmentPatternsEnrollmentPatternMetadataFieldResponse)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetRestrictCAs() bool`

GetRestrictCAs returns the RestrictCAs field if non-nil, zero value otherwise.

### GetRestrictCAsOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetRestrictCAsOk() (*bool, bool)`

GetRestrictCAsOk returns a tuple with the RestrictCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetRestrictCAs(v bool)`

SetRestrictCAs sets RestrictCAs field to given value.

### HasRestrictCAs

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasRestrictCAs() bool`

HasRestrictCAs returns a boolean if a field has been set.

### GetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetPolicies() EnrollmentPatternsEnrollmentPatternPolicyResponse`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetPoliciesOk() (*EnrollmentPatternsEnrollmentPatternPolicyResponse, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetPolicies(v EnrollmentPatternsEnrollmentPatternPolicyResponse)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetDefaults() []EnrollmentPatternsEnrollmentPatternDefaultResponse`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetDefaultsOk() (*[]EnrollmentPatternsEnrollmentPatternDefaultResponse, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetDefaults(v []EnrollmentPatternsEnrollmentPatternDefaultResponse)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetEnrollmentFields() []EnrollmentPatternsEnrollmentPatternFieldResponse`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *EnrollmentPatternsEnrollmentPatternResponse) GetEnrollmentFieldsOk() (*[]EnrollmentPatternsEnrollmentPatternFieldResponse, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetEnrollmentFields(v []EnrollmentPatternsEnrollmentPatternFieldResponse)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *EnrollmentPatternsEnrollmentPatternResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *EnrollmentPatternsEnrollmentPatternResponse) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *EnrollmentPatternsEnrollmentPatternResponse) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


