# ModelsTemplateRetrievalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CommonName** | **string** |  | 
**TemplateName** | **string** |  | 
**Oid** | **string** |  | 
**KeySize** | **string** |  | 
**KeyType** | **string** |  | 
**ForestRoot** | Pointer to **string** |  | [optional] [readonly] 
**ConfigurationTenant** | **string** |  | 
**FriendlyName** | Pointer to **string** |  | [optional] 
**KeyRetention** | **string** |  | 
**KeyRetentionDays** | Pointer to **int64** |  | [optional] 
**KeyArchival** | **bool** |  | 
**EnrollmentFields** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel**](ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateMetadataFieldModel**](ModelsTemplateRetrievalResponseTemplateMetadataFieldModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | **int64** |  | 
**TemplateRegexes** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateRegexModel**](ModelsTemplateRetrievalResponseTemplateRegexModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateDefaultModel**](ModelsTemplateRetrievalResponseTemplateDefaultModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**ModelsTemplateRetrievalResponseTemplatePolicyModel**](ModelsTemplateRetrievalResponseTemplatePolicyModel.md) |  | [optional] 
**UseAllowedRequesters** | **bool** |  | 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | **string** |  | 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
**RequiresApproval** | **bool** |  | 
**KeyUsage** | **int64** |  | 
**ExtendedKeyUsages** | Pointer to [**[]ModelsTemplateRetrievalResponseExtendedKeyUsageModel**](ModelsTemplateRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsTemplateRetrievalResponse

`func NewModelsTemplateRetrievalResponse(id int64, commonName string, templateName string, oid string, keySize string, keyType string, configurationTenant string, keyRetention string, keyArchival bool, allowedEnrollmentTypes int64, useAllowedRequesters bool, displayName string, requiresApproval bool, keyUsage int64, ) *ModelsTemplateRetrievalResponse`

NewModelsTemplateRetrievalResponse instantiates a new ModelsTemplateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateRetrievalResponseWithDefaults

`func NewModelsTemplateRetrievalResponseWithDefaults() *ModelsTemplateRetrievalResponse`

NewModelsTemplateRetrievalResponseWithDefaults instantiates a new ModelsTemplateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsTemplateRetrievalResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTemplateRetrievalResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTemplateRetrievalResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetCommonName

`func (o *ModelsTemplateRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *ModelsTemplateRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *ModelsTemplateRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetTemplateName

`func (o *ModelsTemplateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ModelsTemplateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ModelsTemplateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetOid

`func (o *ModelsTemplateRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *ModelsTemplateRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *ModelsTemplateRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.


### GetKeySize

`func (o *ModelsTemplateRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *ModelsTemplateRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *ModelsTemplateRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.


### GetKeyType

`func (o *ModelsTemplateRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsTemplateRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetForestRoot

`func (o *ModelsTemplateRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *ModelsTemplateRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *ModelsTemplateRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *ModelsTemplateRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### GetConfigurationTenant

`func (o *ModelsTemplateRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *ModelsTemplateRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *ModelsTemplateRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.


### GetFriendlyName

`func (o *ModelsTemplateRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ModelsTemplateRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ModelsTemplateRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ModelsTemplateRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetKeyRetention

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetention() string`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionOk() (*string, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsTemplateRetrievalResponse) SetKeyRetention(v string)`

SetKeyRetention sets KeyRetention field to given value.


### GetKeyRetentionDays

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionDays() int64`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionDaysOk() (*int64, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsTemplateRetrievalResponse) SetKeyRetentionDays(v int64)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *ModelsTemplateRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### GetKeyArchival

`func (o *ModelsTemplateRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *ModelsTemplateRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.


### GetEnrollmentFields

`func (o *ModelsTemplateRetrievalResponse) GetEnrollmentFields() []ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *ModelsTemplateRetrievalResponse) GetEnrollmentFieldsOk() (*[]ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *ModelsTemplateRetrievalResponse) SetEnrollmentFields(v []ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *ModelsTemplateRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### GetMetadataFields

`func (o *ModelsTemplateRetrievalResponse) GetMetadataFields() []ModelsTemplateRetrievalResponseTemplateMetadataFieldModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *ModelsTemplateRetrievalResponse) GetMetadataFieldsOk() (*[]ModelsTemplateRetrievalResponseTemplateMetadataFieldModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *ModelsTemplateRetrievalResponse) SetMetadataFields(v []ModelsTemplateRetrievalResponseTemplateMetadataFieldModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *ModelsTemplateRetrievalResponse) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### GetAllowedEnrollmentTypes

`func (o *ModelsTemplateRetrievalResponse) GetAllowedEnrollmentTypes() int64`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsTemplateRetrievalResponse) GetAllowedEnrollmentTypesOk() (*int64, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsTemplateRetrievalResponse) SetAllowedEnrollmentTypes(v int64)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.


### GetTemplateRegexes

`func (o *ModelsTemplateRetrievalResponse) GetTemplateRegexes() []ModelsTemplateRetrievalResponseTemplateRegexModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *ModelsTemplateRetrievalResponse) GetTemplateRegexesOk() (*[]ModelsTemplateRetrievalResponseTemplateRegexModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *ModelsTemplateRetrievalResponse) SetTemplateRegexes(v []ModelsTemplateRetrievalResponseTemplateRegexModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *ModelsTemplateRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### GetTemplateDefaults

`func (o *ModelsTemplateRetrievalResponse) GetTemplateDefaults() []ModelsTemplateRetrievalResponseTemplateDefaultModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *ModelsTemplateRetrievalResponse) GetTemplateDefaultsOk() (*[]ModelsTemplateRetrievalResponseTemplateDefaultModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *ModelsTemplateRetrievalResponse) SetTemplateDefaults(v []ModelsTemplateRetrievalResponseTemplateDefaultModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *ModelsTemplateRetrievalResponse) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### GetTemplatePolicy

`func (o *ModelsTemplateRetrievalResponse) GetTemplatePolicy() ModelsTemplateRetrievalResponseTemplatePolicyModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *ModelsTemplateRetrievalResponse) GetTemplatePolicyOk() (*ModelsTemplateRetrievalResponseTemplatePolicyModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *ModelsTemplateRetrievalResponse) SetTemplatePolicy(v ModelsTemplateRetrievalResponseTemplatePolicyModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *ModelsTemplateRetrievalResponse) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *ModelsTemplateRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.


### GetAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *ModelsTemplateRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelsTemplateRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelsTemplateRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelsTemplateRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRFCEnforcement

`func (o *ModelsTemplateRetrievalResponse) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *ModelsTemplateRetrievalResponse) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *ModelsTemplateRetrievalResponse) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *ModelsTemplateRetrievalResponse) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *ModelsTemplateRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *ModelsTemplateRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *ModelsTemplateRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.


### GetKeyUsage

`func (o *ModelsTemplateRetrievalResponse) GetKeyUsage() int64`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyUsageOk() (*int64, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *ModelsTemplateRetrievalResponse) SetKeyUsage(v int64)`

SetKeyUsage sets KeyUsage field to given value.


### GetExtendedKeyUsages

`func (o *ModelsTemplateRetrievalResponse) GetExtendedKeyUsages() []ModelsTemplateRetrievalResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *ModelsTemplateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]ModelsTemplateRetrievalResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *ModelsTemplateRetrievalResponse) SetExtendedKeyUsages(v []ModelsTemplateRetrievalResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *ModelsTemplateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### GetCurve

`func (o *ModelsTemplateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *ModelsTemplateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *ModelsTemplateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *ModelsTemplateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *ModelsTemplateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *ModelsTemplateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


