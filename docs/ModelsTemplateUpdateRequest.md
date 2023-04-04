# ModelsTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**KeySize** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**KeyRetention** | Pointer to **int32** |  | [optional] 
**KeyRetentionDays** | Pointer to **int32** |  | [optional] 
**KeyArchival** | Pointer to **bool** |  | [optional] 
**EnrollmentFields** | Pointer to [**[]ModelsTemplateUpdateRequestTemplateEnrollmentFieldModel**](ModelsTemplateUpdateRequestTemplateEnrollmentFieldModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]ModelsTemplateUpdateRequestTemplateMetadataFieldModel**](ModelsTemplateUpdateRequestTemplateMetadataFieldModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to **int32** |  | [optional] 
**TemplateRegexes** | Pointer to [**[]ModelsTemplateUpdateRequestTemplateRegexModel**](ModelsTemplateUpdateRequestTemplateRegexModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]ModelsTemplateUpdateRequestTemplateDefaultModel**](ModelsTemplateUpdateRequestTemplateDefaultModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**ModelsTemplateUpdateRequestTemplatePolicyModel**](ModelsTemplateUpdateRequestTemplatePolicyModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsTemplateUpdateRequest

`func NewModelsTemplateUpdateRequest() *ModelsTemplateUpdateRequest`

NewModelsTemplateUpdateRequest instantiates a new ModelsTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateUpdateRequestWithDefaults

`func NewModelsTemplateUpdateRequestWithDefaults() *ModelsTemplateUpdateRequest`

NewModelsTemplateUpdateRequestWithDefaults instantiates a new ModelsTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsTemplateUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTemplateUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTemplateUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsTemplateUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeySize

`func (o *ModelsTemplateUpdateRequest) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *ModelsTemplateUpdateRequest) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *ModelsTemplateUpdateRequest) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *ModelsTemplateUpdateRequest) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetKeyType

`func (o *ModelsTemplateUpdateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsTemplateUpdateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsTemplateUpdateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ModelsTemplateUpdateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetFriendlyName

`func (o *ModelsTemplateUpdateRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ModelsTemplateUpdateRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ModelsTemplateUpdateRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ModelsTemplateUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetKeyRetention

`func (o *ModelsTemplateUpdateRequest) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *ModelsTemplateUpdateRequest) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *ModelsTemplateUpdateRequest) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *ModelsTemplateUpdateRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *ModelsTemplateUpdateRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *ModelsTemplateUpdateRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *ModelsTemplateUpdateRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *ModelsTemplateUpdateRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### GetKeyArchival

`func (o *ModelsTemplateUpdateRequest) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *ModelsTemplateUpdateRequest) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *ModelsTemplateUpdateRequest) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *ModelsTemplateUpdateRequest) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *ModelsTemplateUpdateRequest) GetEnrollmentFields() []ModelsTemplateUpdateRequestTemplateEnrollmentFieldModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *ModelsTemplateUpdateRequest) GetEnrollmentFieldsOk() (*[]ModelsTemplateUpdateRequestTemplateEnrollmentFieldModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *ModelsTemplateUpdateRequest) SetEnrollmentFields(v []ModelsTemplateUpdateRequestTemplateEnrollmentFieldModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *ModelsTemplateUpdateRequest) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### GetMetadataFields

`func (o *ModelsTemplateUpdateRequest) GetMetadataFields() []ModelsTemplateUpdateRequestTemplateMetadataFieldModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *ModelsTemplateUpdateRequest) GetMetadataFieldsOk() (*[]ModelsTemplateUpdateRequestTemplateMetadataFieldModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *ModelsTemplateUpdateRequest) SetMetadataFields(v []ModelsTemplateUpdateRequestTemplateMetadataFieldModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *ModelsTemplateUpdateRequest) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### GetAllowedEnrollmentTypes

`func (o *ModelsTemplateUpdateRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *ModelsTemplateUpdateRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *ModelsTemplateUpdateRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *ModelsTemplateUpdateRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *ModelsTemplateUpdateRequest) GetTemplateRegexes() []ModelsTemplateUpdateRequestTemplateRegexModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *ModelsTemplateUpdateRequest) GetTemplateRegexesOk() (*[]ModelsTemplateUpdateRequestTemplateRegexModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *ModelsTemplateUpdateRequest) SetTemplateRegexes(v []ModelsTemplateUpdateRequestTemplateRegexModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *ModelsTemplateUpdateRequest) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### GetTemplateDefaults

`func (o *ModelsTemplateUpdateRequest) GetTemplateDefaults() []ModelsTemplateUpdateRequestTemplateDefaultModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *ModelsTemplateUpdateRequest) GetTemplateDefaultsOk() (*[]ModelsTemplateUpdateRequestTemplateDefaultModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *ModelsTemplateUpdateRequest) SetTemplateDefaults(v []ModelsTemplateUpdateRequestTemplateDefaultModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *ModelsTemplateUpdateRequest) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### GetTemplatePolicy

`func (o *ModelsTemplateUpdateRequest) GetTemplatePolicy() ModelsTemplateUpdateRequestTemplatePolicyModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *ModelsTemplateUpdateRequest) GetTemplatePolicyOk() (*ModelsTemplateUpdateRequestTemplatePolicyModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *ModelsTemplateUpdateRequest) SetTemplatePolicy(v ModelsTemplateUpdateRequestTemplatePolicyModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *ModelsTemplateUpdateRequest) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *ModelsTemplateUpdateRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *ModelsTemplateUpdateRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *ModelsTemplateUpdateRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *ModelsTemplateUpdateRequest) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *ModelsTemplateUpdateRequest) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *ModelsTemplateUpdateRequest) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *ModelsTemplateUpdateRequest) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *ModelsTemplateUpdateRequest) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *ModelsTemplateUpdateRequest) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *ModelsTemplateUpdateRequest) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *ModelsTemplateUpdateRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


