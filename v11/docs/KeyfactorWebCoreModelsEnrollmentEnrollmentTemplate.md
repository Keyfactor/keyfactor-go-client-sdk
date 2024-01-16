# KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Forest** | Pointer to **NullableString** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**RfcEnforcement** | Pointer to **bool** |  | [optional] 
**CAs** | Pointer to [**[]KeyfactorWebCoreModelsEnrollmentEnrollmentCA**](KeyfactorWebCoreModelsEnrollmentEnrollmentCA.md) |  | [optional] 
**EnrollmentFields** | Pointer to [**[]ModelsTemplatesTemplateEnrollmentField**](ModelsTemplatesTemplateEnrollmentField.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]ModelsTemplatesTemplateMetadataField**](ModelsTemplatesTemplateMetadataField.md) |  | [optional] 
**Regexes** | Pointer to [**[]ModelsTemplatesTemplateRegex**](ModelsTemplatesTemplateRegex.md) |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]ModelsExtendedKeyUsage**](ModelsExtendedKeyUsage.md) |  | [optional] 
**EnrollmentTemplatePolicy** | Pointer to [**KeyfactorWebCoreModelsEnrollmentEnrollmentTemplatePolicy**](KeyfactorWebCoreModelsEnrollmentEnrollmentTemplatePolicy.md) |  | [optional] 
**KeySize** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplate

`func NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplate() *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate`

NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplate instantiates a new KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplateWithDefaults

`func NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplateWithDefaults() *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate`

NewKeyfactorWebCoreModelsEnrollmentEnrollmentTemplateWithDefaults instantiates a new KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetForest

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetRequiresApproval

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetRfcEnforcement

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRfcEnforcement() bool`

GetRfcEnforcement returns the RfcEnforcement field if non-nil, zero value otherwise.

### GetRfcEnforcementOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRfcEnforcementOk() (*bool, bool)`

GetRfcEnforcementOk returns a tuple with the RfcEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEnforcement

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetRfcEnforcement(v bool)`

SetRfcEnforcement sets RfcEnforcement field to given value.

### HasRfcEnforcement

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasRfcEnforcement() bool`

HasRfcEnforcement returns a boolean if a field has been set.

### GetCAs

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetCAs() []KeyfactorWebCoreModelsEnrollmentEnrollmentCA`

GetCAs returns the CAs field if non-nil, zero value otherwise.

### GetCAsOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetCAsOk() (*[]KeyfactorWebCoreModelsEnrollmentEnrollmentCA, bool)`

GetCAsOk returns a tuple with the CAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAs

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetCAs(v []KeyfactorWebCoreModelsEnrollmentEnrollmentCA)`

SetCAs sets CAs field to given value.

### HasCAs

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasCAs() bool`

HasCAs returns a boolean if a field has been set.

### SetCAsNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetCAsNil(b bool)`

 SetCAsNil sets the value for CAs to be an explicit nil

### UnsetCAs
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetCAs()`

UnsetCAs ensures that no value is present for CAs, not even an explicit nil
### GetEnrollmentFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFields() []ModelsTemplatesTemplateEnrollmentField`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFieldsOk() (*[]ModelsTemplatesTemplateEnrollmentField, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetEnrollmentFields(v []ModelsTemplatesTemplateEnrollmentField)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### SetEnrollmentFieldsNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetEnrollmentFieldsNil(b bool)`

 SetEnrollmentFieldsNil sets the value for EnrollmentFields to be an explicit nil

### UnsetEnrollmentFields
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetEnrollmentFields()`

UnsetEnrollmentFields ensures that no value is present for EnrollmentFields, not even an explicit nil
### GetMetadataFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetMetadataFields() []ModelsTemplatesTemplateMetadataField`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetMetadataFieldsOk() (*[]ModelsTemplatesTemplateMetadataField, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetMetadataFields(v []ModelsTemplatesTemplateMetadataField)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### SetMetadataFieldsNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetMetadataFieldsNil(b bool)`

 SetMetadataFieldsNil sets the value for MetadataFields to be an explicit nil

### UnsetMetadataFields
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetMetadataFields()`

UnsetMetadataFields ensures that no value is present for MetadataFields, not even an explicit nil
### GetRegexes

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRegexes() []ModelsTemplatesTemplateRegex`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetRegexesOk() (*[]ModelsTemplatesTemplateRegex, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetRegexes(v []ModelsTemplatesTemplateRegex)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### SetRegexesNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetRegexesNil(b bool)`

 SetRegexesNil sets the value for Regexes to be an explicit nil

### UnsetRegexes
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetRegexes()`

UnsetRegexes ensures that no value is present for Regexes, not even an explicit nil
### GetExtendedKeyUsages

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsages() []ModelsExtendedKeyUsage`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsagesOk() (*[]ModelsExtendedKeyUsage, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetExtendedKeyUsages(v []ModelsExtendedKeyUsage)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetEnrollmentTemplatePolicy

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentTemplatePolicy() KeyfactorWebCoreModelsEnrollmentEnrollmentTemplatePolicy`

GetEnrollmentTemplatePolicy returns the EnrollmentTemplatePolicy field if non-nil, zero value otherwise.

### GetEnrollmentTemplatePolicyOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentTemplatePolicyOk() (*KeyfactorWebCoreModelsEnrollmentEnrollmentTemplatePolicy, bool)`

GetEnrollmentTemplatePolicyOk returns a tuple with the EnrollmentTemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentTemplatePolicy

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetEnrollmentTemplatePolicy(v KeyfactorWebCoreModelsEnrollmentEnrollmentTemplatePolicy)`

SetEnrollmentTemplatePolicy sets EnrollmentTemplatePolicy field to given value.

### HasEnrollmentTemplatePolicy

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasEnrollmentTemplatePolicy() bool`

HasEnrollmentTemplatePolicy returns a boolean if a field has been set.

### GetKeySize

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### SetKeySizeNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetKeySizeNil(b bool)`

 SetKeySizeNil sets the value for KeySize to be an explicit nil

### UnsetKeySize
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetKeySize()`

UnsetKeySize ensures that no value is present for KeySize, not even an explicit nil
### GetKeyType

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetCurve

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *KeyfactorWebCoreModelsEnrollmentEnrollmentTemplate) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


