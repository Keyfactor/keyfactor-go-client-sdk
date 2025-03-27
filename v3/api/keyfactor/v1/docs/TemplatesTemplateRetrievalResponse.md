# TemplatesTemplateRetrievalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**TemplateName** | Pointer to **NullableString** |  | [optional] 
**Oid** | Pointer to **NullableString** |  | [optional] 
**KeySize** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] [readonly] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**KeyRetention** | Pointer to [**CSSCMSCoreEnumsKeyRetentionPolicy**](CSSCMSCoreEnumsKeyRetentionPolicy.md) |  | [optional] 
**KeyRetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**KeyArchival** | Pointer to **bool** |  | [optional] 
**EnrollmentFields** | Pointer to [**[]TemplatesTemplateEnrollmentFieldRequestResponseModel**](TemplatesTemplateEnrollmentFieldRequestResponseModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]TemplatesTemplateMetadataFieldRequestResponseModel**](TemplatesTemplateMetadataFieldRequestResponseModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to [**CSSCMSCoreEnumsEnrollmentType**](CSSCMSCoreEnumsEnrollmentType.md) |  | [optional] 
**TemplateRegexes** | Pointer to [**[]TemplatesTemplateRegexRequestResponseModel**](TemplatesTemplateRegexRequestResponseModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]TemplatesTemplateDefaultRequestResponseModel**](TemplatesTemplateDefaultRequestResponseModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**TemplatesTemplatePolicyResponseModel**](TemplatesTemplatePolicyResponseModel.md) |  | [optional] 
**KeyAlgorithms** | Pointer to [**TemplatesKeyAlgorithmsResponseModel**](TemplatesKeyAlgorithmsResponseModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]TemplatesExtendedKeyUsageResponseModel**](TemplatesExtendedKeyUsageResponseModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**KeyTypes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTemplatesTemplateRetrievalResponse

`func NewTemplatesTemplateRetrievalResponse() *TemplatesTemplateRetrievalResponse`

NewTemplatesTemplateRetrievalResponse instantiates a new TemplatesTemplateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateRetrievalResponseWithDefaults

`func NewTemplatesTemplateRetrievalResponseWithDefaults() *TemplatesTemplateRetrievalResponse`

NewTemplatesTemplateRetrievalResponseWithDefaults instantiates a new TemplatesTemplateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesTemplateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesTemplateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesTemplateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesTemplateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommonName

`func (o *TemplatesTemplateRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *TemplatesTemplateRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *TemplatesTemplateRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *TemplatesTemplateRetrievalResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *TemplatesTemplateRetrievalResponse) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *TemplatesTemplateRetrievalResponse) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetTemplateName

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplatesTemplateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *TemplatesTemplateRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetOid

`func (o *TemplatesTemplateRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *TemplatesTemplateRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *TemplatesTemplateRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *TemplatesTemplateRetrievalResponse) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *TemplatesTemplateRetrievalResponse) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *TemplatesTemplateRetrievalResponse) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetKeySize

`func (o *TemplatesTemplateRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *TemplatesTemplateRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *TemplatesTemplateRetrievalResponse) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### SetKeySizeNil

`func (o *TemplatesTemplateRetrievalResponse) SetKeySizeNil(b bool)`

 SetKeySizeNil sets the value for KeySize to be an explicit nil

### UnsetKeySize
`func (o *TemplatesTemplateRetrievalResponse) UnsetKeySize()`

UnsetKeySize ensures that no value is present for KeySize, not even an explicit nil
### GetKeyType

`func (o *TemplatesTemplateRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TemplatesTemplateRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TemplatesTemplateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *TemplatesTemplateRetrievalResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *TemplatesTemplateRetrievalResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetForestRoot

`func (o *TemplatesTemplateRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *TemplatesTemplateRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *TemplatesTemplateRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *TemplatesTemplateRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *TemplatesTemplateRetrievalResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *TemplatesTemplateRetrievalResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *TemplatesTemplateRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *TemplatesTemplateRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *TemplatesTemplateRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *TemplatesTemplateRetrievalResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *TemplatesTemplateRetrievalResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *TemplatesTemplateRetrievalResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetFriendlyName

`func (o *TemplatesTemplateRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TemplatesTemplateRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TemplatesTemplateRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TemplatesTemplateRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TemplatesTemplateRetrievalResponse) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TemplatesTemplateRetrievalResponse) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *TemplatesTemplateRetrievalResponse) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *TemplatesTemplateRetrievalResponse) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *TemplatesTemplateRetrievalResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *TemplatesTemplateRetrievalResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *TemplatesTemplateRetrievalResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *TemplatesTemplateRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *TemplatesTemplateRetrievalResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *TemplatesTemplateRetrievalResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *TemplatesTemplateRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *TemplatesTemplateRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *TemplatesTemplateRetrievalResponse) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *TemplatesTemplateRetrievalResponse) GetEnrollmentFields() []TemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *TemplatesTemplateRetrievalResponse) GetEnrollmentFieldsOk() (*[]TemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *TemplatesTemplateRetrievalResponse) SetEnrollmentFields(v []TemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *TemplatesTemplateRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *TemplatesTemplateRetrievalResponse) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *TemplatesTemplateRetrievalResponse) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetMetadataFields

`func (o *TemplatesTemplateRetrievalResponse) GetMetadataFields() []TemplatesTemplateMetadataFieldRequestResponseModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *TemplatesTemplateRetrievalResponse) GetMetadataFieldsOk() (*[]TemplatesTemplateMetadataFieldRequestResponseModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *TemplatesTemplateRetrievalResponse) SetMetadataFields(v []TemplatesTemplateMetadataFieldRequestResponseModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *TemplatesTemplateRetrievalResponse) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *TemplatesTemplateRetrievalResponse) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *TemplatesTemplateRetrievalResponse) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *TemplatesTemplateRetrievalResponse) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *TemplatesTemplateRetrievalResponse) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *TemplatesTemplateRetrievalResponse) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *TemplatesTemplateRetrievalResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateRegexes() []TemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateRegexesOk() (*[]TemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateRegexes(v []TemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *TemplatesTemplateRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *TemplatesTemplateRetrievalResponse) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetTemplateDefaults

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateDefaults() []TemplatesTemplateDefaultRequestResponseModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *TemplatesTemplateRetrievalResponse) GetTemplateDefaultsOk() (*[]TemplatesTemplateDefaultRequestResponseModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateDefaults(v []TemplatesTemplateDefaultRequestResponseModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *TemplatesTemplateRetrievalResponse) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### SetTemplateDefaultsNil

`func (o *TemplatesTemplateRetrievalResponse) SetTemplateDefaultsNil(b bool)`

 SetTemplateDefaultsNil sets the value for TemplateDefaults to be an explicit nil

### UnsetTemplateDefaults
`func (o *TemplatesTemplateRetrievalResponse) UnsetTemplateDefaults()`

UnsetTemplateDefaults ensures that no value is present for TemplateDefaults, not even an explicit nil
### GetTemplatePolicy

`func (o *TemplatesTemplateRetrievalResponse) GetTemplatePolicy() TemplatesTemplatePolicyResponseModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *TemplatesTemplateRetrievalResponse) GetTemplatePolicyOk() (*TemplatesTemplatePolicyResponseModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *TemplatesTemplateRetrievalResponse) SetTemplatePolicy(v TemplatesTemplatePolicyResponseModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *TemplatesTemplateRetrievalResponse) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetKeyAlgorithms

`func (o *TemplatesTemplateRetrievalResponse) GetKeyAlgorithms() TemplatesKeyAlgorithmsResponseModel`

GetKeyAlgorithms returns the KeyAlgorithms field if non-nil, zero value otherwise.

### GetKeyAlgorithmsOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyAlgorithmsOk() (*TemplatesKeyAlgorithmsResponseModel, bool)`

GetKeyAlgorithmsOk returns a tuple with the KeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithms

`func (o *TemplatesTemplateRetrievalResponse) SetKeyAlgorithms(v TemplatesKeyAlgorithmsResponseModel)`

SetKeyAlgorithms sets KeyAlgorithms field to given value.

### HasKeyAlgorithms

`func (o *TemplatesTemplateRetrievalResponse) HasKeyAlgorithms() bool`

HasKeyAlgorithms returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *TemplatesTemplateRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *TemplatesTemplateRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *TemplatesTemplateRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *TemplatesTemplateRetrievalResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *TemplatesTemplateRetrievalResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetDisplayName

`func (o *TemplatesTemplateRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TemplatesTemplateRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TemplatesTemplateRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TemplatesTemplateRetrievalResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TemplatesTemplateRetrievalResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TemplatesTemplateRetrievalResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRFCEnforcement

`func (o *TemplatesTemplateRetrievalResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *TemplatesTemplateRetrievalResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *TemplatesTemplateRetrievalResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *TemplatesTemplateRetrievalResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *TemplatesTemplateRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *TemplatesTemplateRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *TemplatesTemplateRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *TemplatesTemplateRetrievalResponse) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *TemplatesTemplateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *TemplatesTemplateRetrievalResponse) GetExtendedKeyUsages() []TemplatesExtendedKeyUsageResponseModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *TemplatesTemplateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]TemplatesExtendedKeyUsageResponseModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *TemplatesTemplateRetrievalResponse) SetExtendedKeyUsages(v []TemplatesExtendedKeyUsageResponseModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *TemplatesTemplateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *TemplatesTemplateRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *TemplatesTemplateRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetCurve

`func (o *TemplatesTemplateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TemplatesTemplateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TemplatesTemplateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TemplatesTemplateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *TemplatesTemplateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *TemplatesTemplateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetAllowOneClickRenewals

`func (o *TemplatesTemplateRetrievalResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *TemplatesTemplateRetrievalResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *TemplatesTemplateRetrievalResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *TemplatesTemplateRetrievalResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetKeyTypes

`func (o *TemplatesTemplateRetrievalResponse) GetKeyTypes() string`

GetKeyTypes returns the KeyTypes field if non-nil, zero value otherwise.

### GetKeyTypesOk

`func (o *TemplatesTemplateRetrievalResponse) GetKeyTypesOk() (*string, bool)`

GetKeyTypesOk returns a tuple with the KeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypes

`func (o *TemplatesTemplateRetrievalResponse) SetKeyTypes(v string)`

SetKeyTypes sets KeyTypes field to given value.

### HasKeyTypes

`func (o *TemplatesTemplateRetrievalResponse) HasKeyTypes() bool`

HasKeyTypes returns a boolean if a field has been set.

### SetKeyTypesNil

`func (o *TemplatesTemplateRetrievalResponse) SetKeyTypesNil(b bool)`

 SetKeyTypesNil sets the value for KeyTypes to be an explicit nil

### UnsetKeyTypes
`func (o *TemplatesTemplateRetrievalResponse) UnsetKeyTypes()`

UnsetKeyTypes ensures that no value is present for KeyTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


