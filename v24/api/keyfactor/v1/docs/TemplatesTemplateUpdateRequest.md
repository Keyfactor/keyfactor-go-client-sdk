# TemplatesTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**KeyRetention** | Pointer to [**CSSCMSCoreEnumsKeyRetentionPolicy**](CSSCMSCoreEnumsKeyRetentionPolicy.md) |  | [optional] 
**KeyRetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**KeyArchival** | Pointer to **bool** |  | [optional] 
**EnrollmentFields** | Pointer to [**[]TemplatesTemplateEnrollmentFieldRequestResponseModel**](TemplatesTemplateEnrollmentFieldRequestResponseModel.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]TemplatesTemplateMetadataFieldRequestResponseModel**](TemplatesTemplateMetadataFieldRequestResponseModel.md) |  | [optional] 
**AllowedEnrollmentTypes** | Pointer to [**CSSCMSCoreEnumsEnrollmentType**](CSSCMSCoreEnumsEnrollmentType.md) |  | [optional] 
**TemplateRegexes** | Pointer to [**[]TemplatesTemplateRegexRequestResponseModel**](TemplatesTemplateRegexRequestResponseModel.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**[]TemplatesTemplateDefaultRequestResponseModel**](TemplatesTemplateDefaultRequestResponseModel.md) |  | [optional] 
**TemplatePolicy** | Pointer to [**TemplatesTemplatePolicyRequestModel**](TemplatesTemplatePolicyRequestModel.md) |  | [optional] 
**UseAllowedRequesters** | Pointer to **bool** |  | [optional] 
**AllowedRequesters** | Pointer to **[]string** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**AllowOneClickRenewals** | Pointer to **bool** |  | [optional] 

## Methods

### NewTemplatesTemplateUpdateRequest

`func NewTemplatesTemplateUpdateRequest() *TemplatesTemplateUpdateRequest`

NewTemplatesTemplateUpdateRequest instantiates a new TemplatesTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateUpdateRequestWithDefaults

`func NewTemplatesTemplateUpdateRequestWithDefaults() *TemplatesTemplateUpdateRequest`

NewTemplatesTemplateUpdateRequestWithDefaults instantiates a new TemplatesTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesTemplateUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesTemplateUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesTemplateUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesTemplateUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFriendlyName

`func (o *TemplatesTemplateUpdateRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TemplatesTemplateUpdateRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TemplatesTemplateUpdateRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TemplatesTemplateUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TemplatesTemplateUpdateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TemplatesTemplateUpdateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetKeyRetention

`func (o *TemplatesTemplateUpdateRequest) GetKeyRetention() CSSCMSCoreEnumsKeyRetentionPolicy`

GetKeyRetention returns the KeyRetention field if non-nil, zero value otherwise.

### GetKeyRetentionOk

`func (o *TemplatesTemplateUpdateRequest) GetKeyRetentionOk() (*CSSCMSCoreEnumsKeyRetentionPolicy, bool)`

GetKeyRetentionOk returns a tuple with the KeyRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetention

`func (o *TemplatesTemplateUpdateRequest) SetKeyRetention(v CSSCMSCoreEnumsKeyRetentionPolicy)`

SetKeyRetention sets KeyRetention field to given value.

### HasKeyRetention

`func (o *TemplatesTemplateUpdateRequest) HasKeyRetention() bool`

HasKeyRetention returns a boolean if a field has been set.

### GetKeyRetentionDays

`func (o *TemplatesTemplateUpdateRequest) GetKeyRetentionDays() int32`

GetKeyRetentionDays returns the KeyRetentionDays field if non-nil, zero value otherwise.

### GetKeyRetentionDaysOk

`func (o *TemplatesTemplateUpdateRequest) GetKeyRetentionDaysOk() (*int32, bool)`

GetKeyRetentionDaysOk returns a tuple with the KeyRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRetentionDays

`func (o *TemplatesTemplateUpdateRequest) SetKeyRetentionDays(v int32)`

SetKeyRetentionDays sets KeyRetentionDays field to given value.

### HasKeyRetentionDays

`func (o *TemplatesTemplateUpdateRequest) HasKeyRetentionDays() bool`

HasKeyRetentionDays returns a boolean if a field has been set.

### SetKeyRetentionDaysNil

`func (o *TemplatesTemplateUpdateRequest) SetKeyRetentionDaysNil(b bool)`

 SetKeyRetentionDaysNil sets the value for KeyRetentionDays to be an explicit nil

### UnsetKeyRetentionDays
`func (o *TemplatesTemplateUpdateRequest) UnsetKeyRetentionDays()`

UnsetKeyRetentionDays ensures that no value is present for KeyRetentionDays, not even an explicit nil
### GetKeyArchival

`func (o *TemplatesTemplateUpdateRequest) GetKeyArchival() bool`

GetKeyArchival returns the KeyArchival field if non-nil, zero value otherwise.

### GetKeyArchivalOk

`func (o *TemplatesTemplateUpdateRequest) GetKeyArchivalOk() (*bool, bool)`

GetKeyArchivalOk returns a tuple with the KeyArchival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArchival

`func (o *TemplatesTemplateUpdateRequest) SetKeyArchival(v bool)`

SetKeyArchival sets KeyArchival field to given value.

### HasKeyArchival

`func (o *TemplatesTemplateUpdateRequest) HasKeyArchival() bool`

HasKeyArchival returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *TemplatesTemplateUpdateRequest) GetEnrollmentFields() []TemplatesTemplateEnrollmentFieldRequestResponseModel`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *TemplatesTemplateUpdateRequest) GetEnrollmentFieldsOk() (*[]TemplatesTemplateEnrollmentFieldRequestResponseModel, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *TemplatesTemplateUpdateRequest) SetEnrollmentFields(v []TemplatesTemplateEnrollmentFieldRequestResponseModel)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *TemplatesTemplateUpdateRequest) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *TemplatesTemplateUpdateRequest) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *TemplatesTemplateUpdateRequest) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetMetadataFields

`func (o *TemplatesTemplateUpdateRequest) GetMetadataFields() []TemplatesTemplateMetadataFieldRequestResponseModel`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *TemplatesTemplateUpdateRequest) GetMetadataFieldsOk() (*[]TemplatesTemplateMetadataFieldRequestResponseModel, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *TemplatesTemplateUpdateRequest) SetMetadataFields(v []TemplatesTemplateMetadataFieldRequestResponseModel)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *TemplatesTemplateUpdateRequest) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *TemplatesTemplateUpdateRequest) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *TemplatesTemplateUpdateRequest) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetAllowedEnrollmentTypes

`func (o *TemplatesTemplateUpdateRequest) GetAllowedEnrollmentTypes() CSSCMSCoreEnumsEnrollmentType`

GetAllowedEnrollmentTypes returns the AllowedEnrollmentTypes field if non-nil, zero value otherwise.

### GetAllowedEnrollmentTypesOk

`func (o *TemplatesTemplateUpdateRequest) GetAllowedEnrollmentTypesOk() (*CSSCMSCoreEnumsEnrollmentType, bool)`

GetAllowedEnrollmentTypesOk returns a tuple with the AllowedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnrollmentTypes

`func (o *TemplatesTemplateUpdateRequest) SetAllowedEnrollmentTypes(v CSSCMSCoreEnumsEnrollmentType)`

SetAllowedEnrollmentTypes sets AllowedEnrollmentTypes field to given value.

### HasAllowedEnrollmentTypes

`func (o *TemplatesTemplateUpdateRequest) HasAllowedEnrollmentTypes() bool`

HasAllowedEnrollmentTypes returns a boolean if a field has been set.

### GetTemplateRegexes

`func (o *TemplatesTemplateUpdateRequest) GetTemplateRegexes() []TemplatesTemplateRegexRequestResponseModel`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *TemplatesTemplateUpdateRequest) GetTemplateRegexesOk() (*[]TemplatesTemplateRegexRequestResponseModel, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *TemplatesTemplateUpdateRequest) SetTemplateRegexes(v []TemplatesTemplateRegexRequestResponseModel)`

SetTemplateRegexes sets TemplateRegexes field to given value.

### HasTemplateRegexes

`func (o *TemplatesTemplateUpdateRequest) HasTemplateRegexes() bool`

HasTemplateRegexes returns a boolean if a field has been set.

### SetTemplateRegexesNil

`func (o *TemplatesTemplateUpdateRequest) SetTemplateRegexesNil(b bool)`

 SetTemplateRegexesNil sets the value for TemplateRegexes to be an explicit nil

### UnsetTemplateRegexes
`func (o *TemplatesTemplateUpdateRequest) UnsetTemplateRegexes()`

UnsetTemplateRegexes ensures that no value is present for TemplateRegexes, not even an explicit nil
### GetTemplateDefaults

`func (o *TemplatesTemplateUpdateRequest) GetTemplateDefaults() []TemplatesTemplateDefaultRequestResponseModel`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *TemplatesTemplateUpdateRequest) GetTemplateDefaultsOk() (*[]TemplatesTemplateDefaultRequestResponseModel, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *TemplatesTemplateUpdateRequest) SetTemplateDefaults(v []TemplatesTemplateDefaultRequestResponseModel)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *TemplatesTemplateUpdateRequest) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### SetTemplateDefaultsNil

`func (o *TemplatesTemplateUpdateRequest) SetTemplateDefaultsNil(b bool)`

 SetTemplateDefaultsNil sets the value for TemplateDefaults to be an explicit nil

### UnsetTemplateDefaults
`func (o *TemplatesTemplateUpdateRequest) UnsetTemplateDefaults()`

UnsetTemplateDefaults ensures that no value is present for TemplateDefaults, not even an explicit nil
### GetTemplatePolicy

`func (o *TemplatesTemplateUpdateRequest) GetTemplatePolicy() TemplatesTemplatePolicyRequestModel`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *TemplatesTemplateUpdateRequest) GetTemplatePolicyOk() (*TemplatesTemplatePolicyRequestModel, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *TemplatesTemplateUpdateRequest) SetTemplatePolicy(v TemplatesTemplatePolicyRequestModel)`

SetTemplatePolicy sets TemplatePolicy field to given value.

### HasTemplatePolicy

`func (o *TemplatesTemplateUpdateRequest) HasTemplatePolicy() bool`

HasTemplatePolicy returns a boolean if a field has been set.

### GetUseAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) GetUseAllowedRequesters() bool`

GetUseAllowedRequesters returns the UseAllowedRequesters field if non-nil, zero value otherwise.

### GetUseAllowedRequestersOk

`func (o *TemplatesTemplateUpdateRequest) GetUseAllowedRequestersOk() (*bool, bool)`

GetUseAllowedRequestersOk returns a tuple with the UseAllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) SetUseAllowedRequesters(v bool)`

SetUseAllowedRequesters sets UseAllowedRequesters field to given value.

### HasUseAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) HasUseAllowedRequesters() bool`

HasUseAllowedRequesters returns a boolean if a field has been set.

### GetAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) GetAllowedRequesters() []string`

GetAllowedRequesters returns the AllowedRequesters field if non-nil, zero value otherwise.

### GetAllowedRequestersOk

`func (o *TemplatesTemplateUpdateRequest) GetAllowedRequestersOk() (*[]string, bool)`

GetAllowedRequestersOk returns a tuple with the AllowedRequesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) SetAllowedRequesters(v []string)`

SetAllowedRequesters sets AllowedRequesters field to given value.

### HasAllowedRequesters

`func (o *TemplatesTemplateUpdateRequest) HasAllowedRequesters() bool`

HasAllowedRequesters returns a boolean if a field has been set.

### SetAllowedRequestersNil

`func (o *TemplatesTemplateUpdateRequest) SetAllowedRequestersNil(b bool)`

 SetAllowedRequestersNil sets the value for AllowedRequesters to be an explicit nil

### UnsetAllowedRequesters
`func (o *TemplatesTemplateUpdateRequest) UnsetAllowedRequesters()`

UnsetAllowedRequesters ensures that no value is present for AllowedRequesters, not even an explicit nil
### GetRequiresApproval

`func (o *TemplatesTemplateUpdateRequest) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *TemplatesTemplateUpdateRequest) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *TemplatesTemplateUpdateRequest) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *TemplatesTemplateUpdateRequest) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetKeyUsage

`func (o *TemplatesTemplateUpdateRequest) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TemplatesTemplateUpdateRequest) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TemplatesTemplateUpdateRequest) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TemplatesTemplateUpdateRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetAllowOneClickRenewals

`func (o *TemplatesTemplateUpdateRequest) GetAllowOneClickRenewals() bool`

GetAllowOneClickRenewals returns the AllowOneClickRenewals field if non-nil, zero value otherwise.

### GetAllowOneClickRenewalsOk

`func (o *TemplatesTemplateUpdateRequest) GetAllowOneClickRenewalsOk() (*bool, bool)`

GetAllowOneClickRenewalsOk returns a tuple with the AllowOneClickRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOneClickRenewals

`func (o *TemplatesTemplateUpdateRequest) SetAllowOneClickRenewals(v bool)`

SetAllowOneClickRenewals sets AllowOneClickRenewals field to given value.

### HasAllowOneClickRenewals

`func (o *TemplatesTemplateUpdateRequest) HasAllowOneClickRenewals() bool`

HasAllowOneClickRenewals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


