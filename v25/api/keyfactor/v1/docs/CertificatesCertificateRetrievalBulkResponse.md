# CertificatesCertificateRetrievalBulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**ImportDate** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**PrincipalId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**CertState** | Pointer to [**KeyfactorPKIEnumsCertificateState**](KeyfactorPKIEnumsCertificateState.md) |  | [optional] 
**KeySizeInBits** | Pointer to **int32** |  | [optional] 
**KeyType** | Pointer to [**KeyfactorPKIEnumsEncryptionKeyType**](KeyfactorPKIEnumsEncryptionKeyType.md) |  | [optional] 
**KeyAlgorithm** | Pointer to **NullableString** |  | [optional] 
**AltKeyAlgorithm** | Pointer to **NullableString** |  | [optional] 
**AltKeySizeInBits** | Pointer to **int32** |  | [optional] 
**AltKeyType** | Pointer to [**KeyfactorPKIEnumsEncryptionKeyType**](KeyfactorPKIEnumsEncryptionKeyType.md) |  | [optional] 
**RequesterId** | Pointer to **NullableInt32** |  | [optional] 
**IssuedOU** | Pointer to **NullableString** |  | [optional] 
**IssuedEmail** | Pointer to **NullableString** |  | [optional] 
**KeyUsage** | Pointer to **NullableInt32** |  | [optional] 
**SigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**AltSigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CertStateString** | Pointer to **NullableString** |  | [optional] 
**KeyTypeString** | Pointer to **NullableString** |  | [optional] 
**AltKeyTypeString** | Pointer to **NullableString** |  | [optional] 
**RevocationEffDate** | Pointer to **NullableTime** |  | [optional] 
**RevocationReason** | Pointer to [**KeyfactorPKIEnumsRevokeCode**](KeyfactorPKIEnumsRevokeCode.md) |  | [optional] 
**RevocationComment** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthorityId** | Pointer to **NullableInt32** |  | [optional] 
**CertificateAuthorityName** | Pointer to **NullableString** |  | [optional] 
**TemplateName** | Pointer to **NullableString** | Full template display name. | [optional] 
**ArchivedKey** | Pointer to **bool** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] 
**HasAltPrivateKey** | Pointer to **bool** |  | [optional] 
**PrincipalName** | Pointer to **NullableString** |  | [optional] 
**CertRequestId** | Pointer to **NullableInt32** |  | [optional] 
**RequesterName** | Pointer to **NullableString** |  | [optional] 
**ContentBytes** | Pointer to **NullableString** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseExtendedKeyUsageModel**](CertificatesCertificateRetrievalBulkResponseExtendedKeyUsageModel.md) |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseSubjectAlternativeNameModel**](CertificatesCertificateRetrievalBulkResponseSubjectAlternativeNameModel.md) |  | [optional] 
**CRLDistributionPoints** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseCRLDistributionPointModel**](CertificatesCertificateRetrievalBulkResponseCRLDistributionPointModel.md) |  | [optional] 
**LocationsCount** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseLocationCountModel**](CertificatesCertificateRetrievalBulkResponseLocationCountModel.md) |  | [optional] 
**SSLLocations** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseCertificateStoreLocationDetailModel**](CertificatesCertificateRetrievalBulkResponseCertificateStoreLocationDetailModel.md) |  | [optional] 
**Locations** | Pointer to [**[]CertificatesCertificateRetrievalBulkResponseCertificateStoreInventoryItemModel**](CertificatesCertificateRetrievalBulkResponseCertificateStoreInventoryItemModel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CARowIndex** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**CARecordId** | Pointer to **NullableString** |  | [optional] 
**DetailedKeyUsage** | Pointer to [**CertificatesCertificateRetrievalBulkResponseDetailedKeyUsageModel**](CertificatesCertificateRetrievalBulkResponseDetailedKeyUsageModel.md) |  | [optional] 
**KeyRecoverable** | Pointer to **bool** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** |  | [optional] 
**Lifespan** | Pointer to **int32** |  | [optional] 

## Methods

### NewCertificatesCertificateRetrievalBulkResponse

`func NewCertificatesCertificateRetrievalBulkResponse() *CertificatesCertificateRetrievalBulkResponse`

NewCertificatesCertificateRetrievalBulkResponse instantiates a new CertificatesCertificateRetrievalBulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateRetrievalBulkResponseWithDefaults

`func NewCertificatesCertificateRetrievalBulkResponseWithDefaults() *CertificatesCertificateRetrievalBulkResponse`

NewCertificatesCertificateRetrievalBulkResponseWithDefaults instantiates a new CertificatesCertificateRetrievalBulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThumbprint

`func (o *CertificatesCertificateRetrievalBulkResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificatesCertificateRetrievalBulkResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificatesCertificateRetrievalBulkResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetSerialNumber

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificatesCertificateRetrievalBulkResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuedDN

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CertificatesCertificateRetrievalBulkResponse) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuedCN

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *CertificatesCertificateRetrievalBulkResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetImportDate

`func (o *CertificatesCertificateRetrievalBulkResponse) GetImportDate() time.Time`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetImportDateOk() (*time.Time, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *CertificatesCertificateRetrievalBulkResponse) SetImportDate(v time.Time)`

SetImportDate sets ImportDate field to given value.

### HasImportDate

`func (o *CertificatesCertificateRetrievalBulkResponse) HasImportDate() bool`

HasImportDate returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificatesCertificateRetrievalBulkResponse) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificatesCertificateRetrievalBulkResponse) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificatesCertificateRetrievalBulkResponse) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificatesCertificateRetrievalBulkResponse) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificatesCertificateRetrievalBulkResponse) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificatesCertificateRetrievalBulkResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetIssuerDN

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificatesCertificateRetrievalBulkResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetPrincipalId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetPrincipalId() int32`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetPrincipalIdOk() (*int32, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetPrincipalId(v int32)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetOwnerRoleId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *CertificatesCertificateRetrievalBulkResponse) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *CertificatesCertificateRetrievalBulkResponse) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *CertificatesCertificateRetrievalBulkResponse) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil
### GetTemplateId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetCertState

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertState() KeyfactorPKIEnumsCertificateState`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertStateOk() (*KeyfactorPKIEnumsCertificateState, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertState(v KeyfactorPKIEnumsCertificateState)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeySizeInBits() int32`

GetKeySizeInBits returns the KeySizeInBits field if non-nil, zero value otherwise.

### GetKeySizeInBitsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeySizeInBitsOk() (*int32, bool)`

GetKeySizeInBitsOk returns a tuple with the KeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeySizeInBits(v int32)`

SetKeySizeInBits sets KeySizeInBits field to given value.

### HasKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeySizeInBits() bool`

HasKeySizeInBits returns a boolean if a field has been set.

### GetKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyType() KeyfactorPKIEnumsEncryptionKeyType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyTypeOk() (*KeyfactorPKIEnumsEncryptionKeyType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyType(v KeyfactorPKIEnumsEncryptionKeyType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### SetKeyAlgorithmNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyAlgorithmNil(b bool)`

 SetKeyAlgorithmNil sets the value for KeyAlgorithm to be an explicit nil

### UnsetKeyAlgorithm
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetKeyAlgorithm()`

UnsetKeyAlgorithm ensures that no value is present for KeyAlgorithm, not even an explicit nil
### GetAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyAlgorithm() string`

GetAltKeyAlgorithm returns the AltKeyAlgorithm field if non-nil, zero value otherwise.

### GetAltKeyAlgorithmOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyAlgorithmOk() (*string, bool)`

GetAltKeyAlgorithmOk returns a tuple with the AltKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeyAlgorithm(v string)`

SetAltKeyAlgorithm sets AltKeyAlgorithm field to given value.

### HasAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) HasAltKeyAlgorithm() bool`

HasAltKeyAlgorithm returns a boolean if a field has been set.

### SetAltKeyAlgorithmNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeyAlgorithmNil(b bool)`

 SetAltKeyAlgorithmNil sets the value for AltKeyAlgorithm to be an explicit nil

### UnsetAltKeyAlgorithm
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetAltKeyAlgorithm()`

UnsetAltKeyAlgorithm ensures that no value is present for AltKeyAlgorithm, not even an explicit nil
### GetAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeySizeInBits() int32`

GetAltKeySizeInBits returns the AltKeySizeInBits field if non-nil, zero value otherwise.

### GetAltKeySizeInBitsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeySizeInBitsOk() (*int32, bool)`

GetAltKeySizeInBitsOk returns a tuple with the AltKeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeySizeInBits(v int32)`

SetAltKeySizeInBits sets AltKeySizeInBits field to given value.

### HasAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalBulkResponse) HasAltKeySizeInBits() bool`

HasAltKeySizeInBits returns a boolean if a field has been set.

### GetAltKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyType() KeyfactorPKIEnumsEncryptionKeyType`

GetAltKeyType returns the AltKeyType field if non-nil, zero value otherwise.

### GetAltKeyTypeOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyTypeOk() (*KeyfactorPKIEnumsEncryptionKeyType, bool)`

GetAltKeyTypeOk returns a tuple with the AltKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeyType(v KeyfactorPKIEnumsEncryptionKeyType)`

SetAltKeyType sets AltKeyType field to given value.

### HasAltKeyType

`func (o *CertificatesCertificateRetrievalBulkResponse) HasAltKeyType() bool`

HasAltKeyType returns a boolean if a field has been set.

### GetRequesterId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRequesterId() int32`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRequesterIdOk() (*int32, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRequesterId(v int32)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### SetRequesterIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRequesterIdNil(b bool)`

 SetRequesterIdNil sets the value for RequesterId to be an explicit nil

### UnsetRequesterId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetRequesterId()`

UnsetRequesterId ensures that no value is present for RequesterId, not even an explicit nil
### GetIssuedOU

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedOU() string`

GetIssuedOU returns the IssuedOU field if non-nil, zero value otherwise.

### GetIssuedOUOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedOUOk() (*string, bool)`

GetIssuedOUOk returns a tuple with the IssuedOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOU

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedOU(v string)`

SetIssuedOU sets IssuedOU field to given value.

### HasIssuedOU

`func (o *CertificatesCertificateRetrievalBulkResponse) HasIssuedOU() bool`

HasIssuedOU returns a boolean if a field has been set.

### SetIssuedOUNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedOUNil(b bool)`

 SetIssuedOUNil sets the value for IssuedOU to be an explicit nil

### UnsetIssuedOU
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetIssuedOU()`

UnsetIssuedOU ensures that no value is present for IssuedOU, not even an explicit nil
### GetIssuedEmail

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedEmail() string`

GetIssuedEmail returns the IssuedEmail field if non-nil, zero value otherwise.

### GetIssuedEmailOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetIssuedEmailOk() (*string, bool)`

GetIssuedEmailOk returns a tuple with the IssuedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedEmail

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedEmail(v string)`

SetIssuedEmail sets IssuedEmail field to given value.

### HasIssuedEmail

`func (o *CertificatesCertificateRetrievalBulkResponse) HasIssuedEmail() bool`

HasIssuedEmail returns a boolean if a field has been set.

### SetIssuedEmailNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetIssuedEmailNil(b bool)`

 SetIssuedEmailNil sets the value for IssuedEmail to be an explicit nil

### UnsetIssuedEmail
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetIssuedEmail()`

UnsetIssuedEmail ensures that no value is present for IssuedEmail, not even an explicit nil
### GetKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### SetKeyUsageNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyUsageNil(b bool)`

 SetKeyUsageNil sets the value for KeyUsage to be an explicit nil

### UnsetKeyUsage
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetKeyUsage()`

UnsetKeyUsage ensures that no value is present for KeyUsage, not even an explicit nil
### GetSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltSigningAlgorithm() string`

GetAltSigningAlgorithm returns the AltSigningAlgorithm field if non-nil, zero value otherwise.

### GetAltSigningAlgorithmOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltSigningAlgorithmOk() (*string, bool)`

GetAltSigningAlgorithmOk returns a tuple with the AltSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltSigningAlgorithm(v string)`

SetAltSigningAlgorithm sets AltSigningAlgorithm field to given value.

### HasAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalBulkResponse) HasAltSigningAlgorithm() bool`

HasAltSigningAlgorithm returns a boolean if a field has been set.

### SetAltSigningAlgorithmNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltSigningAlgorithmNil(b bool)`

 SetAltSigningAlgorithmNil sets the value for AltSigningAlgorithm to be an explicit nil

### UnsetAltSigningAlgorithm
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetAltSigningAlgorithm()`

UnsetAltSigningAlgorithm ensures that no value is present for AltSigningAlgorithm, not even an explicit nil
### GetCertStateString

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertStateString() string`

GetCertStateString returns the CertStateString field if non-nil, zero value otherwise.

### GetCertStateStringOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertStateStringOk() (*string, bool)`

GetCertStateStringOk returns a tuple with the CertStateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStateString

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertStateString(v string)`

SetCertStateString sets CertStateString field to given value.

### HasCertStateString

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCertStateString() bool`

HasCertStateString returns a boolean if a field has been set.

### SetCertStateStringNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertStateStringNil(b bool)`

 SetCertStateStringNil sets the value for CertStateString to be an explicit nil

### UnsetCertStateString
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCertStateString()`

UnsetCertStateString ensures that no value is present for CertStateString, not even an explicit nil
### GetKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyTypeString() string`

GetKeyTypeString returns the KeyTypeString field if non-nil, zero value otherwise.

### GetKeyTypeStringOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyTypeStringOk() (*string, bool)`

GetKeyTypeStringOk returns a tuple with the KeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyTypeString(v string)`

SetKeyTypeString sets KeyTypeString field to given value.

### HasKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeyTypeString() bool`

HasKeyTypeString returns a boolean if a field has been set.

### SetKeyTypeStringNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyTypeStringNil(b bool)`

 SetKeyTypeStringNil sets the value for KeyTypeString to be an explicit nil

### UnsetKeyTypeString
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetKeyTypeString()`

UnsetKeyTypeString ensures that no value is present for KeyTypeString, not even an explicit nil
### GetAltKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyTypeString() string`

GetAltKeyTypeString returns the AltKeyTypeString field if non-nil, zero value otherwise.

### GetAltKeyTypeStringOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetAltKeyTypeStringOk() (*string, bool)`

GetAltKeyTypeStringOk returns a tuple with the AltKeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeyTypeString(v string)`

SetAltKeyTypeString sets AltKeyTypeString field to given value.

### HasAltKeyTypeString

`func (o *CertificatesCertificateRetrievalBulkResponse) HasAltKeyTypeString() bool`

HasAltKeyTypeString returns a boolean if a field has been set.

### SetAltKeyTypeStringNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetAltKeyTypeStringNil(b bool)`

 SetAltKeyTypeStringNil sets the value for AltKeyTypeString to be an explicit nil

### UnsetAltKeyTypeString
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetAltKeyTypeString()`

UnsetAltKeyTypeString ensures that no value is present for AltKeyTypeString, not even an explicit nil
### GetRevocationEffDate

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationEffDate() time.Time`

GetRevocationEffDate returns the RevocationEffDate field if non-nil, zero value otherwise.

### GetRevocationEffDateOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationEffDateOk() (*time.Time, bool)`

GetRevocationEffDateOk returns a tuple with the RevocationEffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEffDate

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRevocationEffDate(v time.Time)`

SetRevocationEffDate sets RevocationEffDate field to given value.

### HasRevocationEffDate

`func (o *CertificatesCertificateRetrievalBulkResponse) HasRevocationEffDate() bool`

HasRevocationEffDate returns a boolean if a field has been set.

### SetRevocationEffDateNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRevocationEffDateNil(b bool)`

 SetRevocationEffDateNil sets the value for RevocationEffDate to be an explicit nil

### UnsetRevocationEffDate
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetRevocationEffDate()`

UnsetRevocationEffDate ensures that no value is present for RevocationEffDate, not even an explicit nil
### GetRevocationReason

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationReason() KeyfactorPKIEnumsRevokeCode`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationReasonOk() (*KeyfactorPKIEnumsRevokeCode, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRevocationReason(v KeyfactorPKIEnumsRevokeCode)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CertificatesCertificateRetrievalBulkResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### GetRevocationComment

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationComment() string`

GetRevocationComment returns the RevocationComment field if non-nil, zero value otherwise.

### GetRevocationCommentOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRevocationCommentOk() (*string, bool)`

GetRevocationCommentOk returns a tuple with the RevocationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationComment

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRevocationComment(v string)`

SetRevocationComment sets RevocationComment field to given value.

### HasRevocationComment

`func (o *CertificatesCertificateRetrievalBulkResponse) HasRevocationComment() bool`

HasRevocationComment returns a boolean if a field has been set.

### SetRevocationCommentNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRevocationCommentNil(b bool)`

 SetRevocationCommentNil sets the value for RevocationComment to be an explicit nil

### UnsetRevocationComment
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetRevocationComment()`

UnsetRevocationComment ensures that no value is present for RevocationComment, not even an explicit nil
### GetCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### SetCertificateAuthorityIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertificateAuthorityIdNil(b bool)`

 SetCertificateAuthorityIdNil sets the value for CertificateAuthorityId to be an explicit nil

### UnsetCertificateAuthorityId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCertificateAuthorityId()`

UnsetCertificateAuthorityId ensures that no value is present for CertificateAuthorityId, not even an explicit nil
### GetCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertificateAuthorityName() string`

GetCertificateAuthorityName returns the CertificateAuthorityName field if non-nil, zero value otherwise.

### GetCertificateAuthorityNameOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertificateAuthorityNameOk() (*string, bool)`

GetCertificateAuthorityNameOk returns a tuple with the CertificateAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertificateAuthorityName(v string)`

SetCertificateAuthorityName sets CertificateAuthorityName field to given value.

### HasCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCertificateAuthorityName() bool`

HasCertificateAuthorityName returns a boolean if a field has been set.

### SetCertificateAuthorityNameNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertificateAuthorityNameNil(b bool)`

 SetCertificateAuthorityNameNil sets the value for CertificateAuthorityName to be an explicit nil

### UnsetCertificateAuthorityName
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCertificateAuthorityName()`

UnsetCertificateAuthorityName ensures that no value is present for CertificateAuthorityName, not even an explicit nil
### GetTemplateName

`func (o *CertificatesCertificateRetrievalBulkResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CertificatesCertificateRetrievalBulkResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CertificatesCertificateRetrievalBulkResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetArchivedKey

`func (o *CertificatesCertificateRetrievalBulkResponse) GetArchivedKey() bool`

GetArchivedKey returns the ArchivedKey field if non-nil, zero value otherwise.

### GetArchivedKeyOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetArchivedKeyOk() (*bool, bool)`

GetArchivedKeyOk returns a tuple with the ArchivedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedKey

`func (o *CertificatesCertificateRetrievalBulkResponse) SetArchivedKey(v bool)`

SetArchivedKey sets ArchivedKey field to given value.

### HasArchivedKey

`func (o *CertificatesCertificateRetrievalBulkResponse) HasArchivedKey() bool`

HasArchivedKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) GetHasAltPrivateKey() bool`

GetHasAltPrivateKey returns the HasAltPrivateKey field if non-nil, zero value otherwise.

### GetHasAltPrivateKeyOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetHasAltPrivateKeyOk() (*bool, bool)`

GetHasAltPrivateKeyOk returns a tuple with the HasAltPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) SetHasAltPrivateKey(v bool)`

SetHasAltPrivateKey sets HasAltPrivateKey field to given value.

### HasHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalBulkResponse) HasHasAltPrivateKey() bool`

HasHasAltPrivateKey returns a boolean if a field has been set.

### GetPrincipalName

`func (o *CertificatesCertificateRetrievalBulkResponse) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *CertificatesCertificateRetrievalBulkResponse) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *CertificatesCertificateRetrievalBulkResponse) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetCertRequestId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertRequestId() int32`

GetCertRequestId returns the CertRequestId field if non-nil, zero value otherwise.

### GetCertRequestIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCertRequestIdOk() (*int32, bool)`

GetCertRequestIdOk returns a tuple with the CertRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRequestId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertRequestId(v int32)`

SetCertRequestId sets CertRequestId field to given value.

### HasCertRequestId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCertRequestId() bool`

HasCertRequestId returns a boolean if a field has been set.

### SetCertRequestIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCertRequestIdNil(b bool)`

 SetCertRequestIdNil sets the value for CertRequestId to be an explicit nil

### UnsetCertRequestId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCertRequestId()`

UnsetCertRequestId ensures that no value is present for CertRequestId, not even an explicit nil
### GetRequesterName

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *CertificatesCertificateRetrievalBulkResponse) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetContentBytes

`func (o *CertificatesCertificateRetrievalBulkResponse) GetContentBytes() string`

GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.

### GetContentBytesOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetContentBytesOk() (*string, bool)`

GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBytes

`func (o *CertificatesCertificateRetrievalBulkResponse) SetContentBytes(v string)`

SetContentBytes sets ContentBytes field to given value.

### HasContentBytes

`func (o *CertificatesCertificateRetrievalBulkResponse) HasContentBytes() bool`

HasContentBytes returns a boolean if a field has been set.

### SetContentBytesNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetContentBytesNil(b bool)`

 SetContentBytesNil sets the value for ContentBytes to be an explicit nil

### UnsetContentBytes
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetContentBytes()`

UnsetContentBytes ensures that no value is present for ContentBytes, not even an explicit nil
### GetExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalBulkResponse) GetExtendedKeyUsages() []CertificatesCertificateRetrievalBulkResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetExtendedKeyUsagesOk() (*[]CertificatesCertificateRetrievalBulkResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalBulkResponse) SetExtendedKeyUsages(v []CertificatesCertificateRetrievalBulkResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalBulkResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSubjectAltNameElements() []CertificatesCertificateRetrievalBulkResponseSubjectAlternativeNameModel`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSubjectAltNameElementsOk() (*[]CertificatesCertificateRetrievalBulkResponseSubjectAlternativeNameModel, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSubjectAltNameElements(v []CertificatesCertificateRetrievalBulkResponseSubjectAlternativeNameModel)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalBulkResponse) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.

### SetSubjectAltNameElementsNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSubjectAltNameElementsNil(b bool)`

 SetSubjectAltNameElementsNil sets the value for SubjectAltNameElements to be an explicit nil

### UnsetSubjectAltNameElements
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetSubjectAltNameElements()`

UnsetSubjectAltNameElements ensures that no value is present for SubjectAltNameElements, not even an explicit nil
### GetCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCRLDistributionPoints() []CertificatesCertificateRetrievalBulkResponseCRLDistributionPointModel`

GetCRLDistributionPoints returns the CRLDistributionPoints field if non-nil, zero value otherwise.

### GetCRLDistributionPointsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCRLDistributionPointsOk() (*[]CertificatesCertificateRetrievalBulkResponseCRLDistributionPointModel, bool)`

GetCRLDistributionPointsOk returns a tuple with the CRLDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCRLDistributionPoints(v []CertificatesCertificateRetrievalBulkResponseCRLDistributionPointModel)`

SetCRLDistributionPoints sets CRLDistributionPoints field to given value.

### HasCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCRLDistributionPoints() bool`

HasCRLDistributionPoints returns a boolean if a field has been set.

### SetCRLDistributionPointsNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCRLDistributionPointsNil(b bool)`

 SetCRLDistributionPointsNil sets the value for CRLDistributionPoints to be an explicit nil

### UnsetCRLDistributionPoints
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCRLDistributionPoints()`

UnsetCRLDistributionPoints ensures that no value is present for CRLDistributionPoints, not even an explicit nil
### GetLocationsCount

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLocationsCount() []CertificatesCertificateRetrievalBulkResponseLocationCountModel`

GetLocationsCount returns the LocationsCount field if non-nil, zero value otherwise.

### GetLocationsCountOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLocationsCountOk() (*[]CertificatesCertificateRetrievalBulkResponseLocationCountModel, bool)`

GetLocationsCountOk returns a tuple with the LocationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsCount

`func (o *CertificatesCertificateRetrievalBulkResponse) SetLocationsCount(v []CertificatesCertificateRetrievalBulkResponseLocationCountModel)`

SetLocationsCount sets LocationsCount field to given value.

### HasLocationsCount

`func (o *CertificatesCertificateRetrievalBulkResponse) HasLocationsCount() bool`

HasLocationsCount returns a boolean if a field has been set.

### SetLocationsCountNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetLocationsCountNil(b bool)`

 SetLocationsCountNil sets the value for LocationsCount to be an explicit nil

### UnsetLocationsCount
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetLocationsCount()`

UnsetLocationsCount ensures that no value is present for LocationsCount, not even an explicit nil
### GetSSLLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSSLLocations() []CertificatesCertificateRetrievalBulkResponseCertificateStoreLocationDetailModel`

GetSSLLocations returns the SSLLocations field if non-nil, zero value otherwise.

### GetSSLLocationsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetSSLLocationsOk() (*[]CertificatesCertificateRetrievalBulkResponseCertificateStoreLocationDetailModel, bool)`

GetSSLLocationsOk returns a tuple with the SSLLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSSLLocations(v []CertificatesCertificateRetrievalBulkResponseCertificateStoreLocationDetailModel)`

SetSSLLocations sets SSLLocations field to given value.

### HasSSLLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) HasSSLLocations() bool`

HasSSLLocations returns a boolean if a field has been set.

### SetSSLLocationsNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetSSLLocationsNil(b bool)`

 SetSSLLocationsNil sets the value for SSLLocations to be an explicit nil

### UnsetSSLLocations
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetSSLLocations()`

UnsetSSLLocations ensures that no value is present for SSLLocations, not even an explicit nil
### GetLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLocations() []CertificatesCertificateRetrievalBulkResponseCertificateStoreInventoryItemModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLocationsOk() (*[]CertificatesCertificateRetrievalBulkResponseCertificateStoreInventoryItemModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) SetLocations(v []CertificatesCertificateRetrievalBulkResponseCertificateStoreInventoryItemModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *CertificatesCertificateRetrievalBulkResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetMetadata

`func (o *CertificatesCertificateRetrievalBulkResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificatesCertificateRetrievalBulkResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificatesCertificateRetrievalBulkResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCARowIndex

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCARowIndex() int64`

GetCARowIndex returns the CARowIndex field if non-nil, zero value otherwise.

### GetCARowIndexOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCARowIndexOk() (*int64, bool)`

GetCARowIndexOk returns a tuple with the CARowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARowIndex

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCARowIndex(v int64)`

SetCARowIndex sets CARowIndex field to given value.

### HasCARowIndex

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCARowIndex() bool`

HasCARowIndex returns a boolean if a field has been set.

### SetCARowIndexNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCARowIndexNil(b bool)`

 SetCARowIndexNil sets the value for CARowIndex to be an explicit nil

### UnsetCARowIndex
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCARowIndex()`

UnsetCARowIndex ensures that no value is present for CARowIndex, not even an explicit nil
### GetCARecordId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCARecordId() string`

GetCARecordId returns the CARecordId field if non-nil, zero value otherwise.

### GetCARecordIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCARecordIdOk() (*string, bool)`

GetCARecordIdOk returns a tuple with the CARecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARecordId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCARecordId(v string)`

SetCARecordId sets CARecordId field to given value.

### HasCARecordId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCARecordId() bool`

HasCARecordId returns a boolean if a field has been set.

### SetCARecordIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCARecordIdNil(b bool)`

 SetCARecordIdNil sets the value for CARecordId to be an explicit nil

### UnsetCARecordId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCARecordId()`

UnsetCARecordId ensures that no value is present for CARecordId, not even an explicit nil
### GetDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) GetDetailedKeyUsage() CertificatesCertificateRetrievalBulkResponseDetailedKeyUsageModel`

GetDetailedKeyUsage returns the DetailedKeyUsage field if non-nil, zero value otherwise.

### GetDetailedKeyUsageOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetDetailedKeyUsageOk() (*CertificatesCertificateRetrievalBulkResponseDetailedKeyUsageModel, bool)`

GetDetailedKeyUsageOk returns a tuple with the DetailedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) SetDetailedKeyUsage(v CertificatesCertificateRetrievalBulkResponseDetailedKeyUsageModel)`

SetDetailedKeyUsage sets DetailedKeyUsage field to given value.

### HasDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalBulkResponse) HasDetailedKeyUsage() bool`

HasDetailedKeyUsage returns a boolean if a field has been set.

### GetKeyRecoverable

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyRecoverable() bool`

GetKeyRecoverable returns the KeyRecoverable field if non-nil, zero value otherwise.

### GetKeyRecoverableOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetKeyRecoverableOk() (*bool, bool)`

GetKeyRecoverableOk returns a tuple with the KeyRecoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRecoverable

`func (o *CertificatesCertificateRetrievalBulkResponse) SetKeyRecoverable(v bool)`

SetKeyRecoverable sets KeyRecoverable field to given value.

### HasKeyRecoverable

`func (o *CertificatesCertificateRetrievalBulkResponse) HasKeyRecoverable() bool`

HasKeyRecoverable returns a boolean if a field has been set.

### GetCurve

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CertificatesCertificateRetrievalBulkResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalBulkResponse) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalBulkResponse) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalBulkResponse) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *CertificatesCertificateRetrievalBulkResponse) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *CertificatesCertificateRetrievalBulkResponse) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil
### GetLifespan

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLifespan() int32`

GetLifespan returns the Lifespan field if non-nil, zero value otherwise.

### GetLifespanOk

`func (o *CertificatesCertificateRetrievalBulkResponse) GetLifespanOk() (*int32, bool)`

GetLifespanOk returns a tuple with the Lifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifespan

`func (o *CertificatesCertificateRetrievalBulkResponse) SetLifespan(v int32)`

SetLifespan sets Lifespan field to given value.

### HasLifespan

`func (o *CertificatesCertificateRetrievalBulkResponse) HasLifespan() bool`

HasLifespan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


