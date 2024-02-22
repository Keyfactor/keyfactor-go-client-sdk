# ModelsEnrollmentPFXEnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **NullableString** |  | [optional] 
**SaNs** | Pointer to **map[string][]string** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**MetadataInput** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFieldsInput** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]string** |  | [optional] 
**CustomFriendlyName** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PopulateMissingValuesFromAD** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**RenewalCertificateId** | Pointer to **NullableInt32** |  | [optional] 
**ChainOrder** | Pointer to **NullableString** |  | [optional] 
**CertificateCollectionOrder** | Pointer to **int32** |  | [optional] 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsEnrollmentPFXEnrollmentRequest

`func NewModelsEnrollmentPFXEnrollmentRequest() *ModelsEnrollmentPFXEnrollmentRequest`

NewModelsEnrollmentPFXEnrollmentRequest instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentPFXEnrollmentRequestWithDefaults

`func NewModelsEnrollmentPFXEnrollmentRequestWithDefaults() *ModelsEnrollmentPFXEnrollmentRequest`

NewModelsEnrollmentPFXEnrollmentRequestWithDefaults instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetMetadataInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadataInput() map[string]interface{}`

GetMetadataInput returns the MetadataInput field if non-nil, zero value otherwise.

### GetMetadataInputOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadataInputOk() (*map[string]interface{}, bool)`

GetMetadataInputOk returns a tuple with the MetadataInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadataInput(v map[string]interface{})`

SetMetadataInput sets MetadataInput field to given value.

### HasMetadataInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasMetadataInput() bool`

HasMetadataInput returns a boolean if a field has been set.

### SetMetadataInputNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadataInputNil(b bool)`

 SetMetadataInputNil sets the value for MetadataInput to be an explicit nil

### UnsetMetadataInput
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetMetadataInput()`

UnsetMetadataInput ensures that no value is present for MetadataInput, not even an explicit nil
### GetAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsInput() map[string]interface{}`

GetAdditionalEnrollmentFieldsInput returns the AdditionalEnrollmentFieldsInput field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsInputOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsInputOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsInputOk returns a tuple with the AdditionalEnrollmentFieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsInput(v map[string]interface{})`

SetAdditionalEnrollmentFieldsInput sets AdditionalEnrollmentFieldsInput field to given value.

### HasAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFieldsInput() bool`

HasAdditionalEnrollmentFieldsInput returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsInputNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsInputNil(b bool)`

 SetAdditionalEnrollmentFieldsInputNil sets the value for AdditionalEnrollmentFieldsInput to be an explicit nil

### UnsetAdditionalEnrollmentFieldsInput
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetAdditionalEnrollmentFieldsInput()`

UnsetAdditionalEnrollmentFieldsInput ensures that no value is present for AdditionalEnrollmentFieldsInput, not even an explicit nil
### GetTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFields() map[string]string`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]string, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]string)`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyName() string`

GetCustomFriendlyName returns the CustomFriendlyName field if non-nil, zero value otherwise.

### GetCustomFriendlyNameOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyNameOk() (*string, bool)`

GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyName(v string)`

SetCustomFriendlyName sets CustomFriendlyName field to given value.

### HasCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCustomFriendlyName() bool`

HasCustomFriendlyName returns a boolean if a field has been set.

### SetCustomFriendlyNameNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyNameNil(b bool)`

 SetCustomFriendlyNameNil sets the value for CustomFriendlyName to be an explicit nil

### UnsetCustomFriendlyName
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetCustomFriendlyName()`

UnsetCustomFriendlyName ensures that no value is present for CustomFriendlyName, not even an explicit nil
### GetPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromAD() bool`

GetPopulateMissingValuesFromAD returns the PopulateMissingValuesFromAD field if non-nil, zero value otherwise.

### GetPopulateMissingValuesFromADOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromADOk() (*bool, bool)`

GetPopulateMissingValuesFromADOk returns a tuple with the PopulateMissingValuesFromAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPopulateMissingValuesFromAD(v bool)`

SetPopulateMissingValuesFromAD sets PopulateMissingValuesFromAD field to given value.

### HasPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPopulateMissingValuesFromAD() bool`

HasPopulateMissingValuesFromAD returns a boolean if a field has been set.

### GetSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateId() int32`

GetRenewalCertificateId returns the RenewalCertificateId field if non-nil, zero value otherwise.

### GetRenewalCertificateIdOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateIdOk() (*int32, bool)`

GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateId(v int32)`

SetRenewalCertificateId sets RenewalCertificateId field to given value.

### HasRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasRenewalCertificateId() bool`

HasRenewalCertificateId returns a boolean if a field has been set.

### SetRenewalCertificateIdNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateIdNil(b bool)`

 SetRenewalCertificateIdNil sets the value for RenewalCertificateId to be an explicit nil

### UnsetRenewalCertificateId
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetRenewalCertificateId()`

UnsetRenewalCertificateId ensures that no value is present for RenewalCertificateId, not even an explicit nil
### GetChainOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
### GetCertificateCollectionOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateCollectionOrder() int32`

GetCertificateCollectionOrder returns the CertificateCollectionOrder field if non-nil, zero value otherwise.

### GetCertificateCollectionOrderOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateCollectionOrderOk() (*int32, bool)`

GetCertificateCollectionOrderOk returns a tuple with the CertificateCollectionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCertificateCollectionOrder(v int32)`

SetCertificateCollectionOrder sets CertificateCollectionOrder field to given value.

### HasCertificateCollectionOrder

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCertificateCollectionOrder() bool`

HasCertificateCollectionOrder returns a boolean if a field has been set.

### GetUseLegacyEncryption

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
### GetKeyType

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *ModelsEnrollmentPFXEnrollmentRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


