# KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse

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
**KeyRetention** | Pointer to **int32** |  | [optional] 
**KeyRetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**KeyArchival** | Pointer to **bool** |  | [optional] 
**EnrollmentFields** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**TemplateRegexes** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel.md) |  | [optional] 
**KeyAlgorithms** | Pointer to [**KeyfactorWebKeyfactorApiModelsTemplatesKeyAlgorithmsResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesKeyAlgorithmsResponseModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RfcEnforcement** | Pointer to **bool** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**KeyTypes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### SetKeySizeNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeySizeNil(b bool)`

 SetKeySizeNil sets the value for KeySize to be an explicit nil

### UnsetKeySize
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetKeySize()`

UnsetKeySize ensures that no value is present for KeySize, not even an explicit nil
### GetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetEnrollmentFields() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetEnrollmentFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetEnrollmentFields(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetMetadataFields() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetMetadataFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetMetadataFields(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateRegexes() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateRegexesOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateRegexes(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateDefaults() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplateDefaultsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateDefaults(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### SetTemplateDefaultsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplateDefaultsNil(b bool)`

 SetTemplateDefaultsNil sets the value for TemplateDefaults to be an explicit nil

### UnsetTemplateDefaults
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetTemplateDefaults()`

UnsetTemplateDefaults ensures that no value is present for TemplateDefaults, not even an explicit nil
### GetTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplatePolicy() KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetTemplatePolicyOk() (*KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetTemplatePolicy(v KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetKeyAlgorithms

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyAlgorithms() KeyfactorWebKeyfactorApiModelsTemplatesKeyAlgorithmsResponseModel`

GetKeyAlgorithms returns the KeyAlgorithms field if non-nil, zero value otherwise.

### GetKeyAlgorithmsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyAlgorithmsOk() (*KeyfactorWebKeyfactorApiModelsTemplatesKeyAlgorithmsResponseModel, bool)`

GetKeyAlgorithmsOk returns a tuple with the KeyAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithms

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyAlgorithms(v KeyfactorWebKeyfactorApiModelsTemplatesKeyAlgorithmsResponseModel)`

SetKeyAlgorithms sets KeyAlgorithms field to given value.

### HasKeyAlgorithms

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyAlgorithms() bool`

HasKeyAlgorithms returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetRfcEnforcement() bool`

GetRfcEnforcement returns the RfcEnforcement field if non-nil, zero value otherwise.

### GetRfcEnforcementOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetRfcEnforcementOk() (*bool, bool)`

GetRfcEnforcementOk returns a tuple with the RfcEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetRfcEnforcement(v bool)`

SetRfcEnforcement sets RfcEnforcement field to given value.

### HasRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasRfcEnforcement() bool`

HasRfcEnforcement returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetExtendedKeyUsages() []KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetExtendedKeyUsages(v []KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetCurve

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyTypes() string`

GetKeyTypes returns the KeyTypes field if non-nil, zero value otherwise.

### GetKeyTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) GetKeyTypesOk() (*string, bool)`

GetKeyTypesOk returns a tuple with the KeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyTypes(v string)`

SetKeyTypes sets KeyTypes field to given value.

### HasKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) HasKeyTypes() bool`

HasKeyTypes returns a boolean if a field has been set.

### SetKeyTypesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) SetKeyTypesNil(b bool)`

 SetKeyTypesNil sets the value for KeyTypes to be an explicit nil

### UnsetKeyTypes
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse) UnsetKeyTypes()`

UnsetKeyTypes ensures that no value is present for KeyTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


