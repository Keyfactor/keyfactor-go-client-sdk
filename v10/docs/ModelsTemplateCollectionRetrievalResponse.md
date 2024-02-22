# ModelsTemplateCollectionRetrievalResponse

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
**EnrollmentFields** | Pointer to [**[]ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel**](ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | **int64** |  | 
**TemplateRegexes** | Pointer to [**[]ModelsTemplateCollectionRetrievalResponseTemplateRegexModel**](ModelsTemplateCollectionRetrievalResponseTemplateRegexModel.md) |  | [optional] 
**UseAllowedRequesters** | **bool** |  | 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**DisplayName** | **string** |  | 
**RequiresApproval** | **bool** |  | 
**KeyUsage** | **int64** |  | 
**ExtendedKeyUsages** | Pointer to [**[]ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel**](ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 

## Methods

### NewModelsTemplateCollectionRetrievalResponse

`func NewModelsTemplateCollectionRetrievalResponse(id int64, commonName string, templateName string, oid string, keySize string, keyType string, configurationTenant string, keyRetention string, keyArchival bool, allowedEnrollmentTypes int64, useAllowedRequesters bool, displayName string, requiresApproval bool, keyUsage int64, ) *ModelsTemplateCollectionRetrievalResponse`

NewModelsTemplateCollectionRetrievalResponse instantiates a new ModelsTemplateCollectionRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateCollectionRetrievalResponseWithDefaults

`func NewModelsTemplateCollectionRetrievalResponseWithDefaults() *ModelsTemplateCollectionRetrievalResponse`

NewModelsTemplateCollectionRetrievalResponseWithDefaults instantiates a new ModelsTemplateCollectionRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsTemplateCollectionRetrievalResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTemplateCollectionRetrievalResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetCommonName

`func (o *ModelsTemplateCollectionRetrievalResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *ModelsTemplateCollectionRetrievalResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetTemplateName

`func (o *ModelsTemplateCollectionRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ModelsTemplateCollectionRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetOid

`func (o *ModelsTemplateCollectionRetrievalResponse) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *ModelsTemplateCollectionRetrievalResponse) SetOid(v string)`

SetOid sets Oid field to given value.


### GetKeySize

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.


### GetKeyType

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetForestRoot

`func (o *ModelsTemplateCollectionRetrievalResponse) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *ModelsTemplateCollectionRetrievalResponse) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *ModelsTemplateCollectionRetrievalResponse) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### GetConfigurationTenant

`func (o *ModelsTemplateCollectionRetrievalResponse) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *ModelsTemplateCollectionRetrievalResponse) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.


### GetFriendlyName

`func (o *ModelsTemplateCollectionRetrievalResponse) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ModelsTemplateCollectionRetrievalResponse) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ModelsTemplateCollectionRetrievalResponse) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetKeyRetention

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyRetention() string`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyRetentionOk() (*string, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeyRetention(v string)`

SetKeyRetention sets KeyRetention field to given value.


### GetKeyRetentionDays

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyRetentionDays() int64`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyRetentionDaysOk() (*int64, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeyRetentionDays(v int64)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *ModelsTemplateCollectionRetrievalResponse) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### GetKeyArchival

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.


### GetEnrollmentFields

`func (o *ModelsTemplateCollectionRetrievalResponse) GetEnrollmentFields() []ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetEnrollmentFieldsOk() (*[]ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *ModelsTemplateCollectionRetrievalResponse) SetEnrollmentFields(v []ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *ModelsTemplateCollectionRetrievalResponse) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### GetAllowedEnrollmentTypes

`func (o *ModelsTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypes() int64`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetAllowedEnrollmentTypesOk() (*int64, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsTemplateCollectionRetrievalResponse) SetAllowedEnrollmentTypes(v int64)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.


### GetTemplateRegexes

`func (o *ModelsTemplateCollectionRetrievalResponse) GetTemplateRegexes() []ModelsTemplateCollectionRetrievalResponseTemplateRegexModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetTemplateRegexesOk() (*[]ModelsTemplateCollectionRetrievalResponseTemplateRegexModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *ModelsTemplateCollectionRetrievalResponse) SetTemplateRegexes(v []ModelsTemplateCollectionRetrievalResponseTemplateRegexModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *ModelsTemplateCollectionRetrievalResponse) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *ModelsTemplateCollectionRetrievalResponse) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *ModelsTemplateCollectionRetrievalResponse) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.


### GetAllowedRequesters

`func (o *ModelsTemplateCollectionRetrievalResponse) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *ModelsTemplateCollectionRetrievalResponse) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *ModelsTemplateCollectionRetrievalResponse) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelsTemplateCollectionRetrievalResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelsTemplateCollectionRetrievalResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRequiresApproval

`func (o *ModelsTemplateCollectionRetrievalResponse) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *ModelsTemplateCollectionRetrievalResponse) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.


### GetKeyUsage

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyUsage() int64`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetKeyUsageOk() (*int64, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *ModelsTemplateCollectionRetrievalResponse) SetKeyUsage(v int64)`

SetKeyUsage sets KeyUsage field to given value.


### GetExtendedKeyUsages

`func (o *ModelsTemplateCollectionRetrievalResponse) GetExtendedKeyUsages() []ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *ModelsTemplateCollectionRetrievalResponse) GetExtendedKeyUsagesOk() (*[]ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *ModelsTemplateCollectionRetrievalResponse) SetExtendedKeyUsages(v []ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *ModelsTemplateCollectionRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


