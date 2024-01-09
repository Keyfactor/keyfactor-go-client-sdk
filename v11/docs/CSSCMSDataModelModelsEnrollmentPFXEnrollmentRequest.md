# CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest

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
**CertificateCollectionOrder** | Pointer to [**KeyfactorPKIEnumsCertificateCollectionOrder**](KeyfactorPKIEnumsCertificateCollectionOrder.md) |  | [optional] 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest

`func NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest() *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest`

NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest instantiates a new CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequestWithDefaults

`func NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequestWithDefaults() *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest`

NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequestWithDefaults instantiates a new CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetMetadataInput() map[string]interface{}`

GetMetadataInput returns the MetadataInput field if non-nil, zero value otherwise.

### GetMetadataInputOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetMetadataInputOk() (*map[string]interface{}, bool)`

GetMetadataInputOk returns a tuple with the MetadataInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetMetadataInput(v map[string]interface{})`

SetMetadataInput sets MetadataInput field to given value.

### HasMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasMetadataInput() bool`

HasMetadataInput returns a boolean if a field has been set.

### SetMetadataInputNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetMetadataInputNil(b bool)`

 SetMetadataInputNil sets the value for MetadataInput to be an explicit nil

### UnsetMetadataInput
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetMetadataInput()`

UnsetMetadataInput ensures that no value is present for MetadataInput, not even an explicit nil
### GetAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsInput() map[string]interface{}`

GetAdditionalEnrollmentFieldsInput returns the AdditionalEnrollmentFieldsInput field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsInputOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsInputOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsInputOk returns a tuple with the AdditionalEnrollmentFieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsInput(v map[string]interface{})`

SetAdditionalEnrollmentFieldsInput sets AdditionalEnrollmentFieldsInput field to given value.

### HasAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFieldsInput() bool`

HasAdditionalEnrollmentFieldsInput returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsInputNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsInputNil(b bool)`

 SetAdditionalEnrollmentFieldsInputNil sets the value for AdditionalEnrollmentFieldsInput to be an explicit nil

### UnsetAdditionalEnrollmentFieldsInput
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetAdditionalEnrollmentFieldsInput()`

UnsetAdditionalEnrollmentFieldsInput ensures that no value is present for AdditionalEnrollmentFieldsInput, not even an explicit nil
### GetTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFields() map[string]string`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]string, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]string)`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetCustomFriendlyName

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyName() string`

GetCustomFriendlyName returns the CustomFriendlyName field if non-nil, zero value otherwise.

### GetCustomFriendlyNameOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyNameOk() (*string, bool)`

GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFriendlyName

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyName(v string)`

SetCustomFriendlyName sets CustomFriendlyName field to given value.

### HasCustomFriendlyName

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasCustomFriendlyName() bool`

HasCustomFriendlyName returns a boolean if a field has been set.

### SetCustomFriendlyNameNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyNameNil(b bool)`

 SetCustomFriendlyNameNil sets the value for CustomFriendlyName to be an explicit nil

### UnsetCustomFriendlyName
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetCustomFriendlyName()`

UnsetCustomFriendlyName ensures that no value is present for CustomFriendlyName, not even an explicit nil
### GetPassword

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPopulateMissingValuesFromAD

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromAD() bool`

GetPopulateMissingValuesFromAD returns the PopulateMissingValuesFromAD field if non-nil, zero value otherwise.

### GetPopulateMissingValuesFromADOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromADOk() (*bool, bool)`

GetPopulateMissingValuesFromADOk returns a tuple with the PopulateMissingValuesFromAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateMissingValuesFromAD

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetPopulateMissingValuesFromAD(v bool)`

SetPopulateMissingValuesFromAD sets PopulateMissingValuesFromAD field to given value.

### HasPopulateMissingValuesFromAD

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasPopulateMissingValuesFromAD() bool`

HasPopulateMissingValuesFromAD returns a boolean if a field has been set.

### GetSubject

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetRenewalCertificateId

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateId() int32`

GetRenewalCertificateId returns the RenewalCertificateId field if non-nil, zero value otherwise.

### GetRenewalCertificateIdOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateIdOk() (*int32, bool)`

GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalCertificateId

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateId(v int32)`

SetRenewalCertificateId sets RenewalCertificateId field to given value.

### HasRenewalCertificateId

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasRenewalCertificateId() bool`

HasRenewalCertificateId returns a boolean if a field has been set.

### SetRenewalCertificateIdNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateIdNil(b bool)`

 SetRenewalCertificateIdNil sets the value for RenewalCertificateId to be an explicit nil

### UnsetRenewalCertificateId
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetRenewalCertificateId()`

UnsetRenewalCertificateId ensures that no value is present for RenewalCertificateId, not even an explicit nil
### GetChainOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
### GetCertificateCollectionOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCertificateCollectionOrder() KeyfactorPKIEnumsCertificateCollectionOrder`

GetCertificateCollectionOrder returns the CertificateCollectionOrder field if non-nil, zero value otherwise.

### GetCertificateCollectionOrderOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCertificateCollectionOrderOk() (*KeyfactorPKIEnumsCertificateCollectionOrder, bool)`

GetCertificateCollectionOrderOk returns a tuple with the CertificateCollectionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCertificateCollectionOrder(v KeyfactorPKIEnumsCertificateCollectionOrder)`

SetCertificateCollectionOrder sets CertificateCollectionOrder field to given value.

### HasCertificateCollectionOrder

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasCertificateCollectionOrder() bool`

HasCertificateCollectionOrder returns a boolean if a field has been set.

### GetUseLegacyEncryption

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
### GetKeyType

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


