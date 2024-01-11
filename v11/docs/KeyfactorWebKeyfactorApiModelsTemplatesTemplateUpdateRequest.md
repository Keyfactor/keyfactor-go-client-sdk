# KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
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
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyRetention() int32`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyRetentionOk() (*int32, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetKeyRetention(v int32)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetEnrollmentFields() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetEnrollmentFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetEnrollmentFields(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetMetadataFields() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetMetadataFieldsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetMetadataFields(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateMetadataFieldRequestResponseModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowedEnrollmentTypes() int32`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowedEnrollmentTypesOk() (*int32, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetAllowedEnrollmentTypes(v int32)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplateRegexes() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplateRegexesOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetTemplateRegexes(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplateDefaults() []KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplateDefaultsOk() (*[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetTemplateDefaults(v []KeyfactorWebKeyfactorApiModelsTemplatesTemplateDefaultRequestResponseModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### SetTemplateDefaultsNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetTemplateDefaultsNil(b bool)`

 SetTemplateDefaultsNil sets the value for TemplateDefaults to be an explicit nil

### UnsetTemplateDefaults
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetTemplateDefaults()`

UnsetTemplateDefaults ensures that no value is present for TemplateDefaults, not even an explicit nil
### GetTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplatePolicy() KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetTemplatePolicyOk() (*KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetTemplatePolicy(v KeyfactorWebKeyfactorApiModelsTemplatesTemplatePolicyRequestResponseModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


