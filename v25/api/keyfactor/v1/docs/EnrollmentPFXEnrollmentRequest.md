# EnrollmentPFXEnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFriendlyName** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PopulateMissingValuesFromAD** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**RenewalCertificateId** | Pointer to **NullableInt32** |  | [optional] 
**ChainOrder** | Pointer to **NullableString** |  | [optional] 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**MicrosoftTargetCSP** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**IncludeSubjectHeader** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEnrollmentPFXEnrollmentRequest

`func NewEnrollmentPFXEnrollmentRequest() *EnrollmentPFXEnrollmentRequest`

NewEnrollmentPFXEnrollmentRequest instantiates a new EnrollmentPFXEnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPFXEnrollmentRequestWithDefaults

`func NewEnrollmentPFXEnrollmentRequestWithDefaults() *EnrollmentPFXEnrollmentRequest`

NewEnrollmentPFXEnrollmentRequestWithDefaults instantiates a new EnrollmentPFXEnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentRequest) GetCustomFriendlyName() string`

GetCustomFriendlyName returns the CustomFriendlyName field if non-nil, zero value otherwise.

### GetCustomFriendlyNameOk

`func (o *EnrollmentPFXEnrollmentRequest) GetCustomFriendlyNameOk() (*string, bool)`

GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentRequest) SetCustomFriendlyName(v string)`

SetCustomFriendlyName sets CustomFriendlyName field to given value.

### HasCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentRequest) HasCustomFriendlyName() bool`

HasCustomFriendlyName returns a boolean if a field has been set.

### SetCustomFriendlyNameNil

`func (o *EnrollmentPFXEnrollmentRequest) SetCustomFriendlyNameNil(b bool)`

 SetCustomFriendlyNameNil sets the value for CustomFriendlyName to be an explicit nil

### UnsetCustomFriendlyName
`func (o *EnrollmentPFXEnrollmentRequest) UnsetCustomFriendlyName()`

UnsetCustomFriendlyName ensures that no value is present for CustomFriendlyName, not even an explicit nil
### GetPassword

`func (o *EnrollmentPFXEnrollmentRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollmentPFXEnrollmentRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollmentPFXEnrollmentRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnrollmentPFXEnrollmentRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EnrollmentPFXEnrollmentRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EnrollmentPFXEnrollmentRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPopulateMissingValuesFromAD

`func (o *EnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromAD() bool`

GetPopulateMissingValuesFromAD returns the PopulateMissingValuesFromAD field if non-nil, zero value otherwise.

### GetPopulateMissingValuesFromADOk

`func (o *EnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromADOk() (*bool, bool)`

GetPopulateMissingValuesFromADOk returns a tuple with the PopulateMissingValuesFromAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateMissingValuesFromAD

`func (o *EnrollmentPFXEnrollmentRequest) SetPopulateMissingValuesFromAD(v bool)`

SetPopulateMissingValuesFromAD sets PopulateMissingValuesFromAD field to given value.

### HasPopulateMissingValuesFromAD

`func (o *EnrollmentPFXEnrollmentRequest) HasPopulateMissingValuesFromAD() bool`

HasPopulateMissingValuesFromAD returns a boolean if a field has been set.

### GetSubject

`func (o *EnrollmentPFXEnrollmentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EnrollmentPFXEnrollmentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EnrollmentPFXEnrollmentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EnrollmentPFXEnrollmentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EnrollmentPFXEnrollmentRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EnrollmentPFXEnrollmentRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentRequest) GetRenewalCertificateId() int32`

GetRenewalCertificateId returns the RenewalCertificateId field if non-nil, zero value otherwise.

### GetRenewalCertificateIdOk

`func (o *EnrollmentPFXEnrollmentRequest) GetRenewalCertificateIdOk() (*int32, bool)`

GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentRequest) SetRenewalCertificateId(v int32)`

SetRenewalCertificateId sets RenewalCertificateId field to given value.

### HasRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentRequest) HasRenewalCertificateId() bool`

HasRenewalCertificateId returns a boolean if a field has been set.

### SetRenewalCertificateIdNil

`func (o *EnrollmentPFXEnrollmentRequest) SetRenewalCertificateIdNil(b bool)`

 SetRenewalCertificateIdNil sets the value for RenewalCertificateId to be an explicit nil

### UnsetRenewalCertificateId
`func (o *EnrollmentPFXEnrollmentRequest) UnsetRenewalCertificateId()`

UnsetRenewalCertificateId ensures that no value is present for RenewalCertificateId, not even an explicit nil
### GetChainOrder

`func (o *EnrollmentPFXEnrollmentRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *EnrollmentPFXEnrollmentRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *EnrollmentPFXEnrollmentRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *EnrollmentPFXEnrollmentRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *EnrollmentPFXEnrollmentRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *EnrollmentPFXEnrollmentRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
### GetUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *EnrollmentPFXEnrollmentRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *EnrollmentPFXEnrollmentRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *EnrollmentPFXEnrollmentRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
### GetKeyType

`func (o *EnrollmentPFXEnrollmentRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *EnrollmentPFXEnrollmentRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *EnrollmentPFXEnrollmentRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *EnrollmentPFXEnrollmentRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *EnrollmentPFXEnrollmentRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *EnrollmentPFXEnrollmentRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *EnrollmentPFXEnrollmentRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *EnrollmentPFXEnrollmentRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *EnrollmentPFXEnrollmentRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *EnrollmentPFXEnrollmentRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *EnrollmentPFXEnrollmentRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *EnrollmentPFXEnrollmentRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *EnrollmentPFXEnrollmentRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *EnrollmentPFXEnrollmentRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *EnrollmentPFXEnrollmentRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *EnrollmentPFXEnrollmentRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentRequest) GetMicrosoftTargetCSP() string`

GetMicrosoftTargetCSP returns the MicrosoftTargetCSP field if non-nil, zero value otherwise.

### GetMicrosoftTargetCSPOk

`func (o *EnrollmentPFXEnrollmentRequest) GetMicrosoftTargetCSPOk() (*string, bool)`

GetMicrosoftTargetCSPOk returns a tuple with the MicrosoftTargetCSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentRequest) SetMicrosoftTargetCSP(v string)`

SetMicrosoftTargetCSP sets MicrosoftTargetCSP field to given value.

### HasMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentRequest) HasMicrosoftTargetCSP() bool`

HasMicrosoftTargetCSP returns a boolean if a field has been set.

### SetMicrosoftTargetCSPNil

`func (o *EnrollmentPFXEnrollmentRequest) SetMicrosoftTargetCSPNil(b bool)`

 SetMicrosoftTargetCSPNil sets the value for MicrosoftTargetCSP to be an explicit nil

### UnsetMicrosoftTargetCSP
`func (o *EnrollmentPFXEnrollmentRequest) UnsetMicrosoftTargetCSP()`

UnsetMicrosoftTargetCSP ensures that no value is present for MicrosoftTargetCSP, not even an explicit nil
### GetCertificateAuthority

`func (o *EnrollmentPFXEnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *EnrollmentPFXEnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *EnrollmentPFXEnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *EnrollmentPFXEnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *EnrollmentPFXEnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *EnrollmentPFXEnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *EnrollmentPFXEnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *EnrollmentPFXEnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *EnrollmentPFXEnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *EnrollmentPFXEnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentRequest) GetIncludeSubjectHeader() bool`

GetIncludeSubjectHeader returns the IncludeSubjectHeader field if non-nil, zero value otherwise.

### GetIncludeSubjectHeaderOk

`func (o *EnrollmentPFXEnrollmentRequest) GetIncludeSubjectHeaderOk() (*bool, bool)`

GetIncludeSubjectHeaderOk returns a tuple with the IncludeSubjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentRequest) SetIncludeSubjectHeader(v bool)`

SetIncludeSubjectHeader sets IncludeSubjectHeader field to given value.

### HasIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentRequest) HasIncludeSubjectHeader() bool`

HasIncludeSubjectHeader returns a boolean if a field has been set.

### GetMetadata

`func (o *EnrollmentPFXEnrollmentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnrollmentPFXEnrollmentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnrollmentPFXEnrollmentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnrollmentPFXEnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EnrollmentPFXEnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EnrollmentPFXEnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFields() map[string]interface{}`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *EnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]interface{})`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *EnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *EnrollmentPFXEnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetTimestamp

`func (o *EnrollmentPFXEnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnrollmentPFXEnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnrollmentPFXEnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnrollmentPFXEnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOwnerRoleId

`func (o *EnrollmentPFXEnrollmentRequest) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *EnrollmentPFXEnrollmentRequest) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *EnrollmentPFXEnrollmentRequest) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *EnrollmentPFXEnrollmentRequest) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *EnrollmentPFXEnrollmentRequest) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *EnrollmentPFXEnrollmentRequest) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *EnrollmentPFXEnrollmentRequest) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *EnrollmentPFXEnrollmentRequest) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *EnrollmentPFXEnrollmentRequest) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *EnrollmentPFXEnrollmentRequest) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *EnrollmentPFXEnrollmentRequest) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *EnrollmentPFXEnrollmentRequest) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil
### GetTemplate

`func (o *EnrollmentPFXEnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentPFXEnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentPFXEnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentPFXEnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EnrollmentPFXEnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EnrollmentPFXEnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSANs

`func (o *EnrollmentPFXEnrollmentRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *EnrollmentPFXEnrollmentRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *EnrollmentPFXEnrollmentRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *EnrollmentPFXEnrollmentRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *EnrollmentPFXEnrollmentRequest) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *EnrollmentPFXEnrollmentRequest) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil
### GetEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentRequest) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *EnrollmentPFXEnrollmentRequest) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentRequest) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentRequest) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *EnrollmentPFXEnrollmentRequest) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *EnrollmentPFXEnrollmentRequest) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


