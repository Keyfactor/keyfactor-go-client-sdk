# KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse

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
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**TemplateRegexes** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel**](KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel.md) |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**KeyTypes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### SetKeySizeNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeySizeNil(b bool)`

 SetKeySizeNil sets the value for KeySize to be an explicit nil

### UnsetKeySize
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetKeySize()`

UnsetKeySize ensures that no value is present for KeySize, not even an explicit nil
### GetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetEnrollmentFields() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetEnrollmentFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetEnrollmentFields(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetTemplateRegexes() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetTemplateRegexesOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetTemplateRegexes(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetExtendedKeyUsages() []KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetExtendedKeyUsagesOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetExtendedKeyUsages(v []KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyTypes() string`

GetKeyTypes returns the KeyTypes field if non-nil, zero value otherwise.

### GetKeyTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) GetKeyTypesOk() (*string, bool)`

GetKeyTypesOk returns a tuple with the KeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyTypes(v string)`

SetKeyTypes sets KeyTypes field to given value.

### HasKeyTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) HasKeyTypes() bool`

HasKeyTypes returns a boolean if a field has been set.

### SetKeyTypesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) SetKeyTypesNil(b bool)`

 SetKeyTypesNil sets the value for KeyTypes to be an explicit nil

### UnsetKeyTypes
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse) UnsetKeyTypes()`

UnsetKeyTypes ensures that no value is present for KeyTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


