# EnrollmentPFXEnrollmentWithStoresRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | Pointer to [**[]EnrollmentManagementStoreRequest**](EnrollmentManagementStoreRequest.md) | A list of certificate stores a successfully issued PKCS12 should be installed into. | [optional] 
**CustomFriendlyName** | Pointer to **NullableString** | The Friendly name of the issued certificate | [optional] 
**Password** | Pointer to **NullableString** | The custom password for the resulting Private key | [optional] 
**Subject** | Pointer to **NullableString** | The certificte&#39;s subject | [optional] 
**IncludeChain** | Pointer to **bool** | An option to include the certificate chain in the response | [optional] 
**IncludeSubjectHeader** | Pointer to **bool** | An option to remove the subject header comment in pem file | [optional] 
**RenewalCertificateId** | Pointer to **NullableInt32** | An optional Id of a certificate being renewed | [optional] 
**CertificateAuthority** | Pointer to **NullableString** | The Certifcate Authority configuration in the format hostname\\logicalname | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The metadata to be assigned to the certificate after issuance | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]interface{}** | Any Additional Enrollment Fields to be included in the request | [optional] 
**Timestamp** | Pointer to **time.Time** | A timestamp to help ensure a legitimate request | [optional] 
**Template** | Pointer to **NullableString** | The template for the certificate | [optional] 
**SANs** | Pointer to **map[string][]string** | Any Subject Alternative Names to be included in the request | [optional] 
**InstallIntoExistingCertificateStores** | Pointer to **bool** | Allows the new certificate to be installed into all certificate stores the renewed certificate is currently installed into. | [optional] 
**JobTime** | Pointer to **NullableTime** | The time any management jobs should be scheduled. | [optional] 
**ChainOrder** | Pointer to **NullableString** | The order of the certificates in the certificate chain | [optional] 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 
**KeyType** | Pointer to **NullableString** | Certificate key type [RSA, ECC] | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**Curve** | Pointer to **NullableString** | The curve being used that will be sent in when the Key Algorithm is of type ECC | [optional] 
**AlternativeKeyType** | Pointer to **NullableString** | Alternative Certificate key type [ML-DSA-44, ML-DSA-65, ML-DSA-87] | [optional] 
**AlternativeKeyLength** | Pointer to **NullableInt32** | Size of the alternative certificate key. | [optional] 
**MicrosoftTargetCSP** | Pointer to **NullableString** | Optionally pass a microsoft cryptographic service provider to the generated certificate | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** | Id or name of the security role that will have ownership of the generated certificate | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** | Id of pattern for the enrollment | [optional] 
**FileExtension** | Pointer to **NullableString** | The file extension that the downloaded certificate will have | [optional] 

## Methods

### NewEnrollmentPFXEnrollmentWithStoresRequest

`func NewEnrollmentPFXEnrollmentWithStoresRequest() *EnrollmentPFXEnrollmentWithStoresRequest`

NewEnrollmentPFXEnrollmentWithStoresRequest instantiates a new EnrollmentPFXEnrollmentWithStoresRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPFXEnrollmentWithStoresRequestWithDefaults

`func NewEnrollmentPFXEnrollmentWithStoresRequestWithDefaults() *EnrollmentPFXEnrollmentWithStoresRequest`

NewEnrollmentPFXEnrollmentWithStoresRequestWithDefaults instantiates a new EnrollmentPFXEnrollmentWithStoresRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetStores() []EnrollmentManagementStoreRequest`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetStoresOk() (*[]EnrollmentManagementStoreRequest, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetStores(v []EnrollmentManagementStoreRequest)`

SetStores sets Stores field to given value.

### HasStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasStores() bool`

HasStores returns a boolean if a field has been set.

### SetStoresNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetStoresNil(b bool)`

 SetStoresNil sets the value for Stores to be an explicit nil

### UnsetStores
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetStores()`

UnsetStores ensures that no value is present for Stores, not even an explicit nil
### GetCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCustomFriendlyName() string`

GetCustomFriendlyName returns the CustomFriendlyName field if non-nil, zero value otherwise.

### GetCustomFriendlyNameOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCustomFriendlyNameOk() (*string, bool)`

GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCustomFriendlyName(v string)`

SetCustomFriendlyName sets CustomFriendlyName field to given value.

### HasCustomFriendlyName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasCustomFriendlyName() bool`

HasCustomFriendlyName returns a boolean if a field has been set.

### SetCustomFriendlyNameNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCustomFriendlyNameNil(b bool)`

 SetCustomFriendlyNameNil sets the value for CustomFriendlyName to be an explicit nil

### UnsetCustomFriendlyName
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetCustomFriendlyName()`

UnsetCustomFriendlyName ensures that no value is present for CustomFriendlyName, not even an explicit nil
### GetPassword

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSubject

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetIncludeChain

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetIncludeSubjectHeader() bool`

GetIncludeSubjectHeader returns the IncludeSubjectHeader field if non-nil, zero value otherwise.

### GetIncludeSubjectHeaderOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetIncludeSubjectHeaderOk() (*bool, bool)`

GetIncludeSubjectHeaderOk returns a tuple with the IncludeSubjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetIncludeSubjectHeader(v bool)`

SetIncludeSubjectHeader sets IncludeSubjectHeader field to given value.

### HasIncludeSubjectHeader

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasIncludeSubjectHeader() bool`

HasIncludeSubjectHeader returns a boolean if a field has been set.

### GetRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetRenewalCertificateId() int32`

GetRenewalCertificateId returns the RenewalCertificateId field if non-nil, zero value otherwise.

### GetRenewalCertificateIdOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetRenewalCertificateIdOk() (*int32, bool)`

GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetRenewalCertificateId(v int32)`

SetRenewalCertificateId sets RenewalCertificateId field to given value.

### HasRenewalCertificateId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasRenewalCertificateId() bool`

HasRenewalCertificateId returns a boolean if a field has been set.

### SetRenewalCertificateIdNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetRenewalCertificateIdNil(b bool)`

 SetRenewalCertificateIdNil sets the value for RenewalCertificateId to be an explicit nil

### UnsetRenewalCertificateId
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetRenewalCertificateId()`

UnsetRenewalCertificateId ensures that no value is present for RenewalCertificateId, not even an explicit nil
### GetCertificateAuthority

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetMetadata

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAdditionalEnrollmentFields() map[string]interface{}`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAdditionalEnrollmentFields(v map[string]interface{})`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetTimestamp

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTemplate

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSANs

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil
### GetInstallIntoExistingCertificateStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetInstallIntoExistingCertificateStores() bool`

GetInstallIntoExistingCertificateStores returns the InstallIntoExistingCertificateStores field if non-nil, zero value otherwise.

### GetInstallIntoExistingCertificateStoresOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetInstallIntoExistingCertificateStoresOk() (*bool, bool)`

GetInstallIntoExistingCertificateStoresOk returns a tuple with the InstallIntoExistingCertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallIntoExistingCertificateStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetInstallIntoExistingCertificateStores(v bool)`

SetInstallIntoExistingCertificateStores sets InstallIntoExistingCertificateStores field to given value.

### HasInstallIntoExistingCertificateStores

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasInstallIntoExistingCertificateStores() bool`

HasInstallIntoExistingCertificateStores returns a boolean if a field has been set.

### GetJobTime

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetJobTime() time.Time`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetJobTimeOk() (*time.Time, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetJobTime(v time.Time)`

SetJobTime sets JobTime field to given value.

### HasJobTime

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasJobTime() bool`

HasJobTime returns a boolean if a field has been set.

### SetJobTimeNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetJobTimeNil(b bool)`

 SetJobTimeNil sets the value for JobTime to be an explicit nil

### UnsetJobTime
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetJobTime()`

UnsetJobTime ensures that no value is present for JobTime, not even an explicit nil
### GetChainOrder

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
### GetUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
### GetKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCurve

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetAlternativeKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAlternativeKeyType() string`

GetAlternativeKeyType returns the AlternativeKeyType field if non-nil, zero value otherwise.

### GetAlternativeKeyTypeOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAlternativeKeyTypeOk() (*string, bool)`

GetAlternativeKeyTypeOk returns a tuple with the AlternativeKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAlternativeKeyType(v string)`

SetAlternativeKeyType sets AlternativeKeyType field to given value.

### HasAlternativeKeyType

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasAlternativeKeyType() bool`

HasAlternativeKeyType returns a boolean if a field has been set.

### SetAlternativeKeyTypeNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAlternativeKeyTypeNil(b bool)`

 SetAlternativeKeyTypeNil sets the value for AlternativeKeyType to be an explicit nil

### UnsetAlternativeKeyType
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetAlternativeKeyType()`

UnsetAlternativeKeyType ensures that no value is present for AlternativeKeyType, not even an explicit nil
### GetAlternativeKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAlternativeKeyLength() int32`

GetAlternativeKeyLength returns the AlternativeKeyLength field if non-nil, zero value otherwise.

### GetAlternativeKeyLengthOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetAlternativeKeyLengthOk() (*int32, bool)`

GetAlternativeKeyLengthOk returns a tuple with the AlternativeKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAlternativeKeyLength(v int32)`

SetAlternativeKeyLength sets AlternativeKeyLength field to given value.

### HasAlternativeKeyLength

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasAlternativeKeyLength() bool`

HasAlternativeKeyLength returns a boolean if a field has been set.

### SetAlternativeKeyLengthNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetAlternativeKeyLengthNil(b bool)`

 SetAlternativeKeyLengthNil sets the value for AlternativeKeyLength to be an explicit nil

### UnsetAlternativeKeyLength
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetAlternativeKeyLength()`

UnsetAlternativeKeyLength ensures that no value is present for AlternativeKeyLength, not even an explicit nil
### GetMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetMicrosoftTargetCSP() string`

GetMicrosoftTargetCSP returns the MicrosoftTargetCSP field if non-nil, zero value otherwise.

### GetMicrosoftTargetCSPOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetMicrosoftTargetCSPOk() (*string, bool)`

GetMicrosoftTargetCSPOk returns a tuple with the MicrosoftTargetCSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetMicrosoftTargetCSP(v string)`

SetMicrosoftTargetCSP sets MicrosoftTargetCSP field to given value.

### HasMicrosoftTargetCSP

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasMicrosoftTargetCSP() bool`

HasMicrosoftTargetCSP returns a boolean if a field has been set.

### SetMicrosoftTargetCSPNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetMicrosoftTargetCSPNil(b bool)`

 SetMicrosoftTargetCSPNil sets the value for MicrosoftTargetCSP to be an explicit nil

### UnsetMicrosoftTargetCSP
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetMicrosoftTargetCSP()`

UnsetMicrosoftTargetCSP ensures that no value is present for MicrosoftTargetCSP, not even an explicit nil
### GetOwnerRoleId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil
### GetEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil
### GetFileExtension

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### SetFileExtensionNil

`func (o *EnrollmentPFXEnrollmentWithStoresRequest) SetFileExtensionNil(b bool)`

 SetFileExtensionNil sets the value for FileExtension to be an explicit nil

### UnsetFileExtension
`func (o *EnrollmentPFXEnrollmentWithStoresRequest) UnsetFileExtension()`

UnsetFileExtension ensures that no value is present for FileExtension, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


