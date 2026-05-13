# TemplatesTemplateCollectionRetrievalResponse

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
**AllowedEnrollmentTypes** | Pointer to [**CSSCMSCoreEnumsEnrollmentType**](CSSCMSCoreEnumsEnrollmentType.md) |  | [optional] 
**TemplateRegexes** | Pointer to [**[]TemplatesTemplateRegexRequestResponseModel**](TemplatesTemplateRegexRequestResponseModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]TemplatesExtendedKeyUsageResponseModel**](TemplatesExtendedKeyUsageResponseModel.md) |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 
**KeyTypes** | Pointer to **NullableString** |  | [optional] 
**Manageability** | Pointer to [**KeyfactorPlatformExtensionsEnumsTemplateDetailsManageability**](KeyfactorPlatformExtensionsEnumsTemplateDetailsManageability.md) |  | [optional] 

## Methods

### NewTemplatesTemplateCollectionRetrievalResponse

`func NewTemplatesTemplateCollectionRetrievalResponse() *TemplatesTemplateCollectionRetrievalResponse`

NewTemplatesTemplateCollectionRetrievalResponse instantiates a new TemplatesTemplateCollectionRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateCollectionRetrievalResponseWithDefaults

`func NewTemplatesTemplateCollectionRetrievalResponseWithDefaults() *TemplatesTemplateCollectionRetrievalResponse`

NewTemplatesTemplateCollectionRetrievalResponseWithDefaults instantiates a new TemplatesTemplateCollectionRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommonName

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetTemplateName

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetOid

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetKeySize

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### SetKeySizeNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeySizeNil(b bool)`

 SetKeySizeNil sets the value for KeySize to be an explicit nil

### UnsetKeySize
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetKeySize()`

UnsetKeySize ensures that no value is present for KeySize, not even an explicit nil
### GetKeyType

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetForestRoot

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil
### GetFriendlyName

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetEnrollmentFields() []TemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetEnrollmentFieldsOk() (*[]TemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetEnrollmentFields(v []TemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetTemplateRegexes() []TemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetTemplateRegexesOk() (*[]TemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetTemplateRegexes(v []TemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetUseAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetDisplayName

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRequiresApproval

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetExtendedKeyUsages() []TemplatesExtendedKeyUsageResponseModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetExtendedKeyUsagesOk() (*[]TemplatesExtendedKeyUsageResponseModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetExtendedKeyUsages(v []TemplatesExtendedKeyUsageResponseModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetAllowOneClickRenewals

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.

### GetKeyTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyTypes() string`

GetKeyTypes returns the KeyTypes field if non-nil, zero value otherwise.

### GetKeyTypesOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetKeyTypesOk() (*string, bool)`

GetKeyTypesOk returns a tuple with the KeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyTypes(v string)`

SetKeyTypes sets KeyTypes field to given value.

### HasKeyTypes

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasKeyTypes() bool`

HasKeyTypes returns a boolean if a field has been set.

### SetKeyTypesNil

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetKeyTypesNil(b bool)`

 SetKeyTypesNil sets the value for KeyTypes to be an explicit nil

### UnsetKeyTypes
`func (o *TemplatesTemplateCollectionRetrievalResponse) UnsetKeyTypes()`

UnsetKeyTypes ensures that no value is present for KeyTypes, not even an explicit nil
### GetManageability

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetManageability() KeyfactorPlatformExtensionsEnumsTemplateDetailsManageability`

GetManageability returns the Manageability field if non-nil, zero value otherwise.

### GetManageabilityOk

`func (o *TemplatesTemplateCollectionRetrievalResponse) GetManageabilityOk() (*KeyfactorPlatformExtensionsEnumsTemplateDetailsManageability, bool)`

GetManageabilityOk returns a tuple with the Manageability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageability

`func (o *TemplatesTemplateCollectionRetrievalResponse) SetManageability(v KeyfactorPlatformExtensionsEnumsTemplateDetailsManageability)`

SetManageability sets Manageability field to given value.

### HasManageability

`func (o *TemplatesTemplateCollectionRetrievalResponse) HasManageability() bool`

HasManageability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


