# ModelsTemplateRetrievalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**Oid** | Pointer to **string** |  | [optional] 
**KeySize** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**ForestRoot** | Pointer to **string** |  | [optional] [readonly] 
**ConfigurationTenant** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**KeyRetention** | Pointer to **int32** |  | [optional] 
**KeyRetentionDays** | Pointer to **int32** |  | [optional] 
**KeyArchival** | Pointer to **bool** |  | [optional] 
**EnrollmentFields** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel**](ModelsTemplateRetrievalResponseTemplateEnrollmentFieldModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateMetadataFieldModel**](ModelsTemplateRetrievalResponseTemplateMetadataFieldModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**TemplateRegexes** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateRegexModel**](ModelsTemplateRetrievalResponseTemplateRegexModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]ModelsTemplateRetrievalResponseTemplateDefaultModel**](ModelsTemplateRetrievalResponseTemplateDefaultModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**ModelsTemplateRetrievalResponseTemplatePolicyModel**](ModelsTemplateRetrievalResponseTemplatePolicyModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]ModelsTemplateRetrievalResponseExtendedKeyUsageModel**](ModelsTemplateRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsTemplateRetrievalResponse

`func NewModelsTemplateRetrievalResponse() *ModelsTemplateRetrievalResponse`

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

`func (o *ModelsTemplateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTemplateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTemplateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTemplateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasCommonName

`func (o *ModelsTemplateRetrievalResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

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

### HasTemplateName

`func (o *ModelsTemplateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

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

### HasOid

`func (o *ModelsTemplateRetrievalResponse) HasOid() bool`

HasOid returns a boolean if a field has been set.

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

### HasKeySize

`func (o *ModelsTemplateRetrievalResponse) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

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

### HasKeyType

`func (o *ModelsTemplateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

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

### HasConfigurationTenant

`func (o *ModelsTemplateRetrievalResponse) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

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

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsTemplateRetrievalResponse) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *ModelsTemplateRetrievalResponse) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsTemplateRetrievalResponse) SetKeyRetentionDays(v int32)`

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

### HasKeyArchival

`func (o *ModelsTemplateRetrievalResponse) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

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

`func (o *ModelsTemplateRetrievalResponse) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsTemplateRetrievalResponse) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsTemplateRetrievalResponse) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *ModelsTemplateRetrievalResponse) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

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

### HasUseAllowedRequesters

`func (o *ModelsTemplateRetrievalResponse) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

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

### HasDisplayName

`func (o *ModelsTemplateRetrievalResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

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

### HasRequiresApproval

`func (o *ModelsTemplateRetrievalResponse) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *ModelsTemplateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *ModelsTemplateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *ModelsTemplateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *ModelsTemplateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

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


