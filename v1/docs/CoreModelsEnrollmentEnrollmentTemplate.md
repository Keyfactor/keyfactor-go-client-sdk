# CoreModelsEnrollmentEnrollmentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Forest** | Pointer to **string** |  | [optional] 
**KeySize** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**RequiresApproval** | Pointer to **bool** |  | [optional] 
**RFCEnforcement** | Pointer to **bool** |  | [optional] 
**CAs** | Pointer to [**[]CoreModelsEnrollmentEnrollmentCA**](CoreModelsEnrollmentEnrollmentCA.md) |  | [optional] 
**EnrollmentFields** | Pointer to [**[]ModelsTemplateEnrollmentField**](ModelsTemplateEnrollmentField.md) |  | [optional] 
**MetadataFields** | Pointer to [**[]ModelsTemplateMetadataField**](ModelsTemplateMetadataField.md) |  | [optional] 
**Regexes** | Pointer to [**[]ModelsTemplateRegex**](ModelsTemplateRegex.md) |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]ModelsExtendedKeyUsage**](ModelsExtendedKeyUsage.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCoreModelsEnrollmentEnrollmentTemplate

`func NewCoreModelsEnrollmentEnrollmentTemplate() *CoreModelsEnrollmentEnrollmentTemplate`

NewCoreModelsEnrollmentEnrollmentTemplate instantiates a new CoreModelsEnrollmentEnrollmentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreModelsEnrollmentEnrollmentTemplateWithDefaults

`func NewCoreModelsEnrollmentEnrollmentTemplateWithDefaults() *CoreModelsEnrollmentEnrollmentTemplate`

NewCoreModelsEnrollmentEnrollmentTemplateWithDefaults instantiates a new CoreModelsEnrollmentEnrollmentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetForest

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetKeySize

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetKeyType

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.

### GetRFCEnforcement

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRFCEnforcement() bool`

GetRFCEnforcement returns the RFCEnforcement field if non-nil, zero value otherwise.

### GetRFCEnforcementOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRFCEnforcementOk() (*bool, bool)`

GetRFCEnforcementOk returns a tuple with the RFCEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRFCEnforcement

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRFCEnforcement(v bool)`

SetRFCEnforcement sets RFCEnforcement field to given value.

### HasRFCEnforcement

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRFCEnforcement() bool`

HasRFCEnforcement returns a boolean if a field has been set.

### GetCAs

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCAs() []CoreModelsEnrollmentEnrollmentCA`

GetCAs returns the CAs field if non-nil, zero value otherwise.

### GetCAsOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCAsOk() (*[]CoreModelsEnrollmentEnrollmentCA, bool)`

GetCAsOk returns a tuple with the CAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAs

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCAs(v []CoreModelsEnrollmentEnrollmentCA)`

SetCAs sets CAs field to given value.

### HasCAs

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasCAs() bool`

HasCAs returns a boolean if a field has been set.

### GetEnrollmentFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFields() []ModelsTemplateEnrollmentField`

GetEnrollmentFields returns the EnrollmentFields field if non-nil, zero value otherwise.

### GetEnrollmentFieldsOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFieldsOk() (*[]ModelsTemplateEnrollmentField, bool)`

GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetEnrollmentFields(v []ModelsTemplateEnrollmentField)`

SetEnrollmentFields sets EnrollmentFields field to given value.

### HasEnrollmentFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasEnrollmentFields() bool`

HasEnrollmentFields returns a boolean if a field has been set.

### GetMetadataFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetMetadataFields() []ModelsTemplateMetadataField`

GetMetadataFields returns the MetadataFields field if non-nil, zero value otherwise.

### GetMetadataFieldsOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetMetadataFieldsOk() (*[]ModelsTemplateMetadataField, bool)`

GetMetadataFieldsOk returns a tuple with the MetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetMetadataFields(v []ModelsTemplateMetadataField)`

SetMetadataFields sets MetadataFields field to given value.

### HasMetadataFields

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasMetadataFields() bool`

HasMetadataFields returns a boolean if a field has been set.

### GetRegexes

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRegexes() []ModelsTemplateRegex`

GetRegexes returns the Regexes field if non-nil, zero value otherwise.

### GetRegexesOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRegexesOk() (*[]ModelsTemplateRegex, bool)`

GetRegexesOk returns a tuple with the Regexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexes

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRegexes(v []ModelsTemplateRegex)`

SetRegexes sets Regexes field to given value.

### HasRegexes

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRegexes() bool`

HasRegexes returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsages() []ModelsExtendedKeyUsage`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsagesOk() (*[]ModelsExtendedKeyUsage, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetExtendedKeyUsages(v []ModelsExtendedKeyUsage)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### GetCurve

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CoreModelsEnrollmentEnrollmentTemplate) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CoreModelsEnrollmentEnrollmentTemplate) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


