# CertificatesCertificateRetrievalResponse

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
**ExtendedKeyUsages** | Pointer to [**[]CertificatesCertificateRetrievalResponseExtendedKeyUsageModel**](CertificatesCertificateRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]CertificatesCertificateRetrievalResponseSubjectAlternativeNameModel**](CertificatesCertificateRetrievalResponseSubjectAlternativeNameModel.md) |  | [optional] 
**CRLDistributionPoints** | Pointer to [**[]CertificatesCertificateRetrievalResponseCRLDistributionPointModel**](CertificatesCertificateRetrievalResponseCRLDistributionPointModel.md) |  | [optional] 
**LocationsCount** | Pointer to [**[]CertificatesCertificateRetrievalResponseLocationCountModel**](CertificatesCertificateRetrievalResponseLocationCountModel.md) |  | [optional] 
**SSLLocations** | Pointer to [**[]CertificatesCertificateRetrievalResponseCertificateStoreLocationDetailModel**](CertificatesCertificateRetrievalResponseCertificateStoreLocationDetailModel.md) |  | [optional] 
**Locations** | Pointer to [**[]CertificatesCertificateRetrievalResponseCertificateStoreInventoryItemModel**](CertificatesCertificateRetrievalResponseCertificateStoreInventoryItemModel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CARowIndex** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**CARecordId** | Pointer to **NullableString** |  | [optional] 
**DetailedKeyUsage** | Pointer to [**CertificatesCertificateRetrievalResponseDetailedKeyUsageModel**](CertificatesCertificateRetrievalResponseDetailedKeyUsageModel.md) |  | [optional] 
**KeyRecoverable** | Pointer to **bool** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCertificatesCertificateRetrievalResponse

`func NewCertificatesCertificateRetrievalResponse() *CertificatesCertificateRetrievalResponse`

NewCertificatesCertificateRetrievalResponse instantiates a new CertificatesCertificateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateRetrievalResponseWithDefaults

`func NewCertificatesCertificateRetrievalResponseWithDefaults() *CertificatesCertificateRetrievalResponse`

NewCertificatesCertificateRetrievalResponseWithDefaults instantiates a new CertificatesCertificateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertificateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertificateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertificateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThumbprint

`func (o *CertificatesCertificateRetrievalResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificatesCertificateRetrievalResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificatesCertificateRetrievalResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificatesCertificateRetrievalResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CertificatesCertificateRetrievalResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CertificatesCertificateRetrievalResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetSerialNumber

`func (o *CertificatesCertificateRetrievalResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificatesCertificateRetrievalResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificatesCertificateRetrievalResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificatesCertificateRetrievalResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CertificatesCertificateRetrievalResponse) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CertificatesCertificateRetrievalResponse) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuedDN

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CertificatesCertificateRetrievalResponse) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CertificatesCertificateRetrievalResponse) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuedCN

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *CertificatesCertificateRetrievalResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *CertificatesCertificateRetrievalResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetImportDate

`func (o *CertificatesCertificateRetrievalResponse) GetImportDate() time.Time`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *CertificatesCertificateRetrievalResponse) GetImportDateOk() (*time.Time, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *CertificatesCertificateRetrievalResponse) SetImportDate(v time.Time)`

SetImportDate sets ImportDate field to given value.

### HasImportDate

`func (o *CertificatesCertificateRetrievalResponse) HasImportDate() bool`

HasImportDate returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificatesCertificateRetrievalResponse) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificatesCertificateRetrievalResponse) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificatesCertificateRetrievalResponse) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificatesCertificateRetrievalResponse) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificatesCertificateRetrievalResponse) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificatesCertificateRetrievalResponse) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificatesCertificateRetrievalResponse) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificatesCertificateRetrievalResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetIssuerDN

`func (o *CertificatesCertificateRetrievalResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificatesCertificateRetrievalResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificatesCertificateRetrievalResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificatesCertificateRetrievalResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CertificatesCertificateRetrievalResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CertificatesCertificateRetrievalResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetPrincipalId

`func (o *CertificatesCertificateRetrievalResponse) GetPrincipalId() int32`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetPrincipalIdOk() (*int32, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CertificatesCertificateRetrievalResponse) SetPrincipalId(v int32)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *CertificatesCertificateRetrievalResponse) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *CertificatesCertificateRetrievalResponse) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetOwnerRoleId

`func (o *CertificatesCertificateRetrievalResponse) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *CertificatesCertificateRetrievalResponse) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *CertificatesCertificateRetrievalResponse) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *CertificatesCertificateRetrievalResponse) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *CertificatesCertificateRetrievalResponse) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *CertificatesCertificateRetrievalResponse) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *CertificatesCertificateRetrievalResponse) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *CertificatesCertificateRetrievalResponse) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *CertificatesCertificateRetrievalResponse) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *CertificatesCertificateRetrievalResponse) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil
### GetTemplateId

`func (o *CertificatesCertificateRetrievalResponse) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CertificatesCertificateRetrievalResponse) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CertificatesCertificateRetrievalResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *CertificatesCertificateRetrievalResponse) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetCertState

`func (o *CertificatesCertificateRetrievalResponse) GetCertState() KeyfactorPKIEnumsCertificateState`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *CertificatesCertificateRetrievalResponse) GetCertStateOk() (*KeyfactorPKIEnumsCertificateState, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *CertificatesCertificateRetrievalResponse) SetCertState(v KeyfactorPKIEnumsCertificateState)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *CertificatesCertificateRetrievalResponse) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) GetKeySizeInBits() int32`

GetKeySizeInBits returns the KeySizeInBits field if non-nil, zero value otherwise.

### GetKeySizeInBitsOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeySizeInBitsOk() (*int32, bool)`

GetKeySizeInBitsOk returns a tuple with the KeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) SetKeySizeInBits(v int32)`

SetKeySizeInBits sets KeySizeInBits field to given value.

### HasKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) HasKeySizeInBits() bool`

HasKeySizeInBits returns a boolean if a field has been set.

### GetKeyType

`func (o *CertificatesCertificateRetrievalResponse) GetKeyType() KeyfactorPKIEnumsEncryptionKeyType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeyTypeOk() (*KeyfactorPKIEnumsEncryptionKeyType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificatesCertificateRetrievalResponse) SetKeyType(v KeyfactorPKIEnumsEncryptionKeyType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CertificatesCertificateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### SetKeyAlgorithmNil

`func (o *CertificatesCertificateRetrievalResponse) SetKeyAlgorithmNil(b bool)`

 SetKeyAlgorithmNil sets the value for KeyAlgorithm to be an explicit nil

### UnsetKeyAlgorithm
`func (o *CertificatesCertificateRetrievalResponse) UnsetKeyAlgorithm()`

UnsetKeyAlgorithm ensures that no value is present for KeyAlgorithm, not even an explicit nil
### GetAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyAlgorithm() string`

GetAltKeyAlgorithm returns the AltKeyAlgorithm field if non-nil, zero value otherwise.

### GetAltKeyAlgorithmOk

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyAlgorithmOk() (*string, bool)`

GetAltKeyAlgorithmOk returns a tuple with the AltKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeyAlgorithm(v string)`

SetAltKeyAlgorithm sets AltKeyAlgorithm field to given value.

### HasAltKeyAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) HasAltKeyAlgorithm() bool`

HasAltKeyAlgorithm returns a boolean if a field has been set.

### SetAltKeyAlgorithmNil

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeyAlgorithmNil(b bool)`

 SetAltKeyAlgorithmNil sets the value for AltKeyAlgorithm to be an explicit nil

### UnsetAltKeyAlgorithm
`func (o *CertificatesCertificateRetrievalResponse) UnsetAltKeyAlgorithm()`

UnsetAltKeyAlgorithm ensures that no value is present for AltKeyAlgorithm, not even an explicit nil
### GetAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeySizeInBits() int32`

GetAltKeySizeInBits returns the AltKeySizeInBits field if non-nil, zero value otherwise.

### GetAltKeySizeInBitsOk

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeySizeInBitsOk() (*int32, bool)`

GetAltKeySizeInBitsOk returns a tuple with the AltKeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeySizeInBits(v int32)`

SetAltKeySizeInBits sets AltKeySizeInBits field to given value.

### HasAltKeySizeInBits

`func (o *CertificatesCertificateRetrievalResponse) HasAltKeySizeInBits() bool`

HasAltKeySizeInBits returns a boolean if a field has been set.

### GetAltKeyType

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyType() KeyfactorPKIEnumsEncryptionKeyType`

GetAltKeyType returns the AltKeyType field if non-nil, zero value otherwise.

### GetAltKeyTypeOk

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyTypeOk() (*KeyfactorPKIEnumsEncryptionKeyType, bool)`

GetAltKeyTypeOk returns a tuple with the AltKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyType

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeyType(v KeyfactorPKIEnumsEncryptionKeyType)`

SetAltKeyType sets AltKeyType field to given value.

### HasAltKeyType

`func (o *CertificatesCertificateRetrievalResponse) HasAltKeyType() bool`

HasAltKeyType returns a boolean if a field has been set.

### GetRequesterId

`func (o *CertificatesCertificateRetrievalResponse) GetRequesterId() int32`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetRequesterIdOk() (*int32, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *CertificatesCertificateRetrievalResponse) SetRequesterId(v int32)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *CertificatesCertificateRetrievalResponse) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### SetRequesterIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetRequesterIdNil(b bool)`

 SetRequesterIdNil sets the value for RequesterId to be an explicit nil

### UnsetRequesterId
`func (o *CertificatesCertificateRetrievalResponse) UnsetRequesterId()`

UnsetRequesterId ensures that no value is present for RequesterId, not even an explicit nil
### GetIssuedOU

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedOU() string`

GetIssuedOU returns the IssuedOU field if non-nil, zero value otherwise.

### GetIssuedOUOk

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedOUOk() (*string, bool)`

GetIssuedOUOk returns a tuple with the IssuedOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOU

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedOU(v string)`

SetIssuedOU sets IssuedOU field to given value.

### HasIssuedOU

`func (o *CertificatesCertificateRetrievalResponse) HasIssuedOU() bool`

HasIssuedOU returns a boolean if a field has been set.

### SetIssuedOUNil

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedOUNil(b bool)`

 SetIssuedOUNil sets the value for IssuedOU to be an explicit nil

### UnsetIssuedOU
`func (o *CertificatesCertificateRetrievalResponse) UnsetIssuedOU()`

UnsetIssuedOU ensures that no value is present for IssuedOU, not even an explicit nil
### GetIssuedEmail

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedEmail() string`

GetIssuedEmail returns the IssuedEmail field if non-nil, zero value otherwise.

### GetIssuedEmailOk

`func (o *CertificatesCertificateRetrievalResponse) GetIssuedEmailOk() (*string, bool)`

GetIssuedEmailOk returns a tuple with the IssuedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedEmail

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedEmail(v string)`

SetIssuedEmail sets IssuedEmail field to given value.

### HasIssuedEmail

`func (o *CertificatesCertificateRetrievalResponse) HasIssuedEmail() bool`

HasIssuedEmail returns a boolean if a field has been set.

### SetIssuedEmailNil

`func (o *CertificatesCertificateRetrievalResponse) SetIssuedEmailNil(b bool)`

 SetIssuedEmailNil sets the value for IssuedEmail to be an explicit nil

### UnsetIssuedEmail
`func (o *CertificatesCertificateRetrievalResponse) UnsetIssuedEmail()`

UnsetIssuedEmail ensures that no value is present for IssuedEmail, not even an explicit nil
### GetKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### SetKeyUsageNil

`func (o *CertificatesCertificateRetrievalResponse) SetKeyUsageNil(b bool)`

 SetKeyUsageNil sets the value for KeyUsage to be an explicit nil

### UnsetKeyUsage
`func (o *CertificatesCertificateRetrievalResponse) UnsetKeyUsage()`

UnsetKeyUsage ensures that no value is present for KeyUsage, not even an explicit nil
### GetSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CertificatesCertificateRetrievalResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CertificatesCertificateRetrievalResponse) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CertificatesCertificateRetrievalResponse) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) GetAltSigningAlgorithm() string`

GetAltSigningAlgorithm returns the AltSigningAlgorithm field if non-nil, zero value otherwise.

### GetAltSigningAlgorithmOk

`func (o *CertificatesCertificateRetrievalResponse) GetAltSigningAlgorithmOk() (*string, bool)`

GetAltSigningAlgorithmOk returns a tuple with the AltSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) SetAltSigningAlgorithm(v string)`

SetAltSigningAlgorithm sets AltSigningAlgorithm field to given value.

### HasAltSigningAlgorithm

`func (o *CertificatesCertificateRetrievalResponse) HasAltSigningAlgorithm() bool`

HasAltSigningAlgorithm returns a boolean if a field has been set.

### SetAltSigningAlgorithmNil

`func (o *CertificatesCertificateRetrievalResponse) SetAltSigningAlgorithmNil(b bool)`

 SetAltSigningAlgorithmNil sets the value for AltSigningAlgorithm to be an explicit nil

### UnsetAltSigningAlgorithm
`func (o *CertificatesCertificateRetrievalResponse) UnsetAltSigningAlgorithm()`

UnsetAltSigningAlgorithm ensures that no value is present for AltSigningAlgorithm, not even an explicit nil
### GetCertStateString

`func (o *CertificatesCertificateRetrievalResponse) GetCertStateString() string`

GetCertStateString returns the CertStateString field if non-nil, zero value otherwise.

### GetCertStateStringOk

`func (o *CertificatesCertificateRetrievalResponse) GetCertStateStringOk() (*string, bool)`

GetCertStateStringOk returns a tuple with the CertStateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStateString

`func (o *CertificatesCertificateRetrievalResponse) SetCertStateString(v string)`

SetCertStateString sets CertStateString field to given value.

### HasCertStateString

`func (o *CertificatesCertificateRetrievalResponse) HasCertStateString() bool`

HasCertStateString returns a boolean if a field has been set.

### SetCertStateStringNil

`func (o *CertificatesCertificateRetrievalResponse) SetCertStateStringNil(b bool)`

 SetCertStateStringNil sets the value for CertStateString to be an explicit nil

### UnsetCertStateString
`func (o *CertificatesCertificateRetrievalResponse) UnsetCertStateString()`

UnsetCertStateString ensures that no value is present for CertStateString, not even an explicit nil
### GetKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) GetKeyTypeString() string`

GetKeyTypeString returns the KeyTypeString field if non-nil, zero value otherwise.

### GetKeyTypeStringOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeyTypeStringOk() (*string, bool)`

GetKeyTypeStringOk returns a tuple with the KeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) SetKeyTypeString(v string)`

SetKeyTypeString sets KeyTypeString field to given value.

### HasKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) HasKeyTypeString() bool`

HasKeyTypeString returns a boolean if a field has been set.

### SetKeyTypeStringNil

`func (o *CertificatesCertificateRetrievalResponse) SetKeyTypeStringNil(b bool)`

 SetKeyTypeStringNil sets the value for KeyTypeString to be an explicit nil

### UnsetKeyTypeString
`func (o *CertificatesCertificateRetrievalResponse) UnsetKeyTypeString()`

UnsetKeyTypeString ensures that no value is present for KeyTypeString, not even an explicit nil
### GetAltKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyTypeString() string`

GetAltKeyTypeString returns the AltKeyTypeString field if non-nil, zero value otherwise.

### GetAltKeyTypeStringOk

`func (o *CertificatesCertificateRetrievalResponse) GetAltKeyTypeStringOk() (*string, bool)`

GetAltKeyTypeStringOk returns a tuple with the AltKeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeyTypeString(v string)`

SetAltKeyTypeString sets AltKeyTypeString field to given value.

### HasAltKeyTypeString

`func (o *CertificatesCertificateRetrievalResponse) HasAltKeyTypeString() bool`

HasAltKeyTypeString returns a boolean if a field has been set.

### SetAltKeyTypeStringNil

`func (o *CertificatesCertificateRetrievalResponse) SetAltKeyTypeStringNil(b bool)`

 SetAltKeyTypeStringNil sets the value for AltKeyTypeString to be an explicit nil

### UnsetAltKeyTypeString
`func (o *CertificatesCertificateRetrievalResponse) UnsetAltKeyTypeString()`

UnsetAltKeyTypeString ensures that no value is present for AltKeyTypeString, not even an explicit nil
### GetRevocationEffDate

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationEffDate() time.Time`

GetRevocationEffDate returns the RevocationEffDate field if non-nil, zero value otherwise.

### GetRevocationEffDateOk

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationEffDateOk() (*time.Time, bool)`

GetRevocationEffDateOk returns a tuple with the RevocationEffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEffDate

`func (o *CertificatesCertificateRetrievalResponse) SetRevocationEffDate(v time.Time)`

SetRevocationEffDate sets RevocationEffDate field to given value.

### HasRevocationEffDate

`func (o *CertificatesCertificateRetrievalResponse) HasRevocationEffDate() bool`

HasRevocationEffDate returns a boolean if a field has been set.

### SetRevocationEffDateNil

`func (o *CertificatesCertificateRetrievalResponse) SetRevocationEffDateNil(b bool)`

 SetRevocationEffDateNil sets the value for RevocationEffDate to be an explicit nil

### UnsetRevocationEffDate
`func (o *CertificatesCertificateRetrievalResponse) UnsetRevocationEffDate()`

UnsetRevocationEffDate ensures that no value is present for RevocationEffDate, not even an explicit nil
### GetRevocationReason

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationReason() KeyfactorPKIEnumsRevokeCode`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationReasonOk() (*KeyfactorPKIEnumsRevokeCode, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CertificatesCertificateRetrievalResponse) SetRevocationReason(v KeyfactorPKIEnumsRevokeCode)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CertificatesCertificateRetrievalResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### GetRevocationComment

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationComment() string`

GetRevocationComment returns the RevocationComment field if non-nil, zero value otherwise.

### GetRevocationCommentOk

`func (o *CertificatesCertificateRetrievalResponse) GetRevocationCommentOk() (*string, bool)`

GetRevocationCommentOk returns a tuple with the RevocationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationComment

`func (o *CertificatesCertificateRetrievalResponse) SetRevocationComment(v string)`

SetRevocationComment sets RevocationComment field to given value.

### HasRevocationComment

`func (o *CertificatesCertificateRetrievalResponse) HasRevocationComment() bool`

HasRevocationComment returns a boolean if a field has been set.

### SetRevocationCommentNil

`func (o *CertificatesCertificateRetrievalResponse) SetRevocationCommentNil(b bool)`

 SetRevocationCommentNil sets the value for RevocationComment to be an explicit nil

### UnsetRevocationComment
`func (o *CertificatesCertificateRetrievalResponse) UnsetRevocationComment()`

UnsetRevocationComment ensures that no value is present for RevocationComment, not even an explicit nil
### GetCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalResponse) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalResponse) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *CertificatesCertificateRetrievalResponse) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### SetCertificateAuthorityIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetCertificateAuthorityIdNil(b bool)`

 SetCertificateAuthorityIdNil sets the value for CertificateAuthorityId to be an explicit nil

### UnsetCertificateAuthorityId
`func (o *CertificatesCertificateRetrievalResponse) UnsetCertificateAuthorityId()`

UnsetCertificateAuthorityId ensures that no value is present for CertificateAuthorityId, not even an explicit nil
### GetCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalResponse) GetCertificateAuthorityName() string`

GetCertificateAuthorityName returns the CertificateAuthorityName field if non-nil, zero value otherwise.

### GetCertificateAuthorityNameOk

`func (o *CertificatesCertificateRetrievalResponse) GetCertificateAuthorityNameOk() (*string, bool)`

GetCertificateAuthorityNameOk returns a tuple with the CertificateAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalResponse) SetCertificateAuthorityName(v string)`

SetCertificateAuthorityName sets CertificateAuthorityName field to given value.

### HasCertificateAuthorityName

`func (o *CertificatesCertificateRetrievalResponse) HasCertificateAuthorityName() bool`

HasCertificateAuthorityName returns a boolean if a field has been set.

### SetCertificateAuthorityNameNil

`func (o *CertificatesCertificateRetrievalResponse) SetCertificateAuthorityNameNil(b bool)`

 SetCertificateAuthorityNameNil sets the value for CertificateAuthorityName to be an explicit nil

### UnsetCertificateAuthorityName
`func (o *CertificatesCertificateRetrievalResponse) UnsetCertificateAuthorityName()`

UnsetCertificateAuthorityName ensures that no value is present for CertificateAuthorityName, not even an explicit nil
### GetTemplateName

`func (o *CertificatesCertificateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CertificatesCertificateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CertificatesCertificateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CertificatesCertificateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *CertificatesCertificateRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *CertificatesCertificateRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetArchivedKey

`func (o *CertificatesCertificateRetrievalResponse) GetArchivedKey() bool`

GetArchivedKey returns the ArchivedKey field if non-nil, zero value otherwise.

### GetArchivedKeyOk

`func (o *CertificatesCertificateRetrievalResponse) GetArchivedKeyOk() (*bool, bool)`

GetArchivedKeyOk returns a tuple with the ArchivedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedKey

`func (o *CertificatesCertificateRetrievalResponse) SetArchivedKey(v bool)`

SetArchivedKey sets ArchivedKey field to given value.

### HasArchivedKey

`func (o *CertificatesCertificateRetrievalResponse) HasArchivedKey() bool`

HasArchivedKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *CertificatesCertificateRetrievalResponse) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) GetHasAltPrivateKey() bool`

GetHasAltPrivateKey returns the HasAltPrivateKey field if non-nil, zero value otherwise.

### GetHasAltPrivateKeyOk

`func (o *CertificatesCertificateRetrievalResponse) GetHasAltPrivateKeyOk() (*bool, bool)`

GetHasAltPrivateKeyOk returns a tuple with the HasAltPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) SetHasAltPrivateKey(v bool)`

SetHasAltPrivateKey sets HasAltPrivateKey field to given value.

### HasHasAltPrivateKey

`func (o *CertificatesCertificateRetrievalResponse) HasHasAltPrivateKey() bool`

HasHasAltPrivateKey returns a boolean if a field has been set.

### GetPrincipalName

`func (o *CertificatesCertificateRetrievalResponse) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *CertificatesCertificateRetrievalResponse) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *CertificatesCertificateRetrievalResponse) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *CertificatesCertificateRetrievalResponse) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *CertificatesCertificateRetrievalResponse) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *CertificatesCertificateRetrievalResponse) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetCertRequestId

`func (o *CertificatesCertificateRetrievalResponse) GetCertRequestId() int32`

GetCertRequestId returns the CertRequestId field if non-nil, zero value otherwise.

### GetCertRequestIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetCertRequestIdOk() (*int32, bool)`

GetCertRequestIdOk returns a tuple with the CertRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRequestId

`func (o *CertificatesCertificateRetrievalResponse) SetCertRequestId(v int32)`

SetCertRequestId sets CertRequestId field to given value.

### HasCertRequestId

`func (o *CertificatesCertificateRetrievalResponse) HasCertRequestId() bool`

HasCertRequestId returns a boolean if a field has been set.

### SetCertRequestIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetCertRequestIdNil(b bool)`

 SetCertRequestIdNil sets the value for CertRequestId to be an explicit nil

### UnsetCertRequestId
`func (o *CertificatesCertificateRetrievalResponse) UnsetCertRequestId()`

UnsetCertRequestId ensures that no value is present for CertRequestId, not even an explicit nil
### GetRequesterName

`func (o *CertificatesCertificateRetrievalResponse) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *CertificatesCertificateRetrievalResponse) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *CertificatesCertificateRetrievalResponse) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *CertificatesCertificateRetrievalResponse) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *CertificatesCertificateRetrievalResponse) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *CertificatesCertificateRetrievalResponse) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetContentBytes

`func (o *CertificatesCertificateRetrievalResponse) GetContentBytes() string`

GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.

### GetContentBytesOk

`func (o *CertificatesCertificateRetrievalResponse) GetContentBytesOk() (*string, bool)`

GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBytes

`func (o *CertificatesCertificateRetrievalResponse) SetContentBytes(v string)`

SetContentBytes sets ContentBytes field to given value.

### HasContentBytes

`func (o *CertificatesCertificateRetrievalResponse) HasContentBytes() bool`

HasContentBytes returns a boolean if a field has been set.

### SetContentBytesNil

`func (o *CertificatesCertificateRetrievalResponse) SetContentBytesNil(b bool)`

 SetContentBytesNil sets the value for ContentBytes to be an explicit nil

### UnsetContentBytes
`func (o *CertificatesCertificateRetrievalResponse) UnsetContentBytes()`

UnsetContentBytes ensures that no value is present for ContentBytes, not even an explicit nil
### GetExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalResponse) GetExtendedKeyUsages() []CertificatesCertificateRetrievalResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CertificatesCertificateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]CertificatesCertificateRetrievalResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalResponse) SetExtendedKeyUsages(v []CertificatesCertificateRetrievalResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *CertificatesCertificateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *CertificatesCertificateRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *CertificatesCertificateRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalResponse) GetSubjectAltNameElements() []CertificatesCertificateRetrievalResponseSubjectAlternativeNameModel`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *CertificatesCertificateRetrievalResponse) GetSubjectAltNameElementsOk() (*[]CertificatesCertificateRetrievalResponseSubjectAlternativeNameModel, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalResponse) SetSubjectAltNameElements(v []CertificatesCertificateRetrievalResponseSubjectAlternativeNameModel)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *CertificatesCertificateRetrievalResponse) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.

### SetSubjectAltNameElementsNil

`func (o *CertificatesCertificateRetrievalResponse) SetSubjectAltNameElementsNil(b bool)`

 SetSubjectAltNameElementsNil sets the value for SubjectAltNameElements to be an explicit nil

### UnsetSubjectAltNameElements
`func (o *CertificatesCertificateRetrievalResponse) UnsetSubjectAltNameElements()`

UnsetSubjectAltNameElements ensures that no value is present for SubjectAltNameElements, not even an explicit nil
### GetCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalResponse) GetCRLDistributionPoints() []CertificatesCertificateRetrievalResponseCRLDistributionPointModel`

GetCRLDistributionPoints returns the CRLDistributionPoints field if non-nil, zero value otherwise.

### GetCRLDistributionPointsOk

`func (o *CertificatesCertificateRetrievalResponse) GetCRLDistributionPointsOk() (*[]CertificatesCertificateRetrievalResponseCRLDistributionPointModel, bool)`

GetCRLDistributionPointsOk returns a tuple with the CRLDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalResponse) SetCRLDistributionPoints(v []CertificatesCertificateRetrievalResponseCRLDistributionPointModel)`

SetCRLDistributionPoints sets CRLDistributionPoints field to given value.

### HasCRLDistributionPoints

`func (o *CertificatesCertificateRetrievalResponse) HasCRLDistributionPoints() bool`

HasCRLDistributionPoints returns a boolean if a field has been set.

### SetCRLDistributionPointsNil

`func (o *CertificatesCertificateRetrievalResponse) SetCRLDistributionPointsNil(b bool)`

 SetCRLDistributionPointsNil sets the value for CRLDistributionPoints to be an explicit nil

### UnsetCRLDistributionPoints
`func (o *CertificatesCertificateRetrievalResponse) UnsetCRLDistributionPoints()`

UnsetCRLDistributionPoints ensures that no value is present for CRLDistributionPoints, not even an explicit nil
### GetLocationsCount

`func (o *CertificatesCertificateRetrievalResponse) GetLocationsCount() []CertificatesCertificateRetrievalResponseLocationCountModel`

GetLocationsCount returns the LocationsCount field if non-nil, zero value otherwise.

### GetLocationsCountOk

`func (o *CertificatesCertificateRetrievalResponse) GetLocationsCountOk() (*[]CertificatesCertificateRetrievalResponseLocationCountModel, bool)`

GetLocationsCountOk returns a tuple with the LocationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsCount

`func (o *CertificatesCertificateRetrievalResponse) SetLocationsCount(v []CertificatesCertificateRetrievalResponseLocationCountModel)`

SetLocationsCount sets LocationsCount field to given value.

### HasLocationsCount

`func (o *CertificatesCertificateRetrievalResponse) HasLocationsCount() bool`

HasLocationsCount returns a boolean if a field has been set.

### SetLocationsCountNil

`func (o *CertificatesCertificateRetrievalResponse) SetLocationsCountNil(b bool)`

 SetLocationsCountNil sets the value for LocationsCount to be an explicit nil

### UnsetLocationsCount
`func (o *CertificatesCertificateRetrievalResponse) UnsetLocationsCount()`

UnsetLocationsCount ensures that no value is present for LocationsCount, not even an explicit nil
### GetSSLLocations

`func (o *CertificatesCertificateRetrievalResponse) GetSSLLocations() []CertificatesCertificateRetrievalResponseCertificateStoreLocationDetailModel`

GetSSLLocations returns the SSLLocations field if non-nil, zero value otherwise.

### GetSSLLocationsOk

`func (o *CertificatesCertificateRetrievalResponse) GetSSLLocationsOk() (*[]CertificatesCertificateRetrievalResponseCertificateStoreLocationDetailModel, bool)`

GetSSLLocationsOk returns a tuple with the SSLLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLLocations

`func (o *CertificatesCertificateRetrievalResponse) SetSSLLocations(v []CertificatesCertificateRetrievalResponseCertificateStoreLocationDetailModel)`

SetSSLLocations sets SSLLocations field to given value.

### HasSSLLocations

`func (o *CertificatesCertificateRetrievalResponse) HasSSLLocations() bool`

HasSSLLocations returns a boolean if a field has been set.

### SetSSLLocationsNil

`func (o *CertificatesCertificateRetrievalResponse) SetSSLLocationsNil(b bool)`

 SetSSLLocationsNil sets the value for SSLLocations to be an explicit nil

### UnsetSSLLocations
`func (o *CertificatesCertificateRetrievalResponse) UnsetSSLLocations()`

UnsetSSLLocations ensures that no value is present for SSLLocations, not even an explicit nil
### GetLocations

`func (o *CertificatesCertificateRetrievalResponse) GetLocations() []CertificatesCertificateRetrievalResponseCertificateStoreInventoryItemModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *CertificatesCertificateRetrievalResponse) GetLocationsOk() (*[]CertificatesCertificateRetrievalResponseCertificateStoreInventoryItemModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *CertificatesCertificateRetrievalResponse) SetLocations(v []CertificatesCertificateRetrievalResponseCertificateStoreInventoryItemModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *CertificatesCertificateRetrievalResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *CertificatesCertificateRetrievalResponse) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *CertificatesCertificateRetrievalResponse) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetMetadata

`func (o *CertificatesCertificateRetrievalResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificatesCertificateRetrievalResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificatesCertificateRetrievalResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificatesCertificateRetrievalResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CertificatesCertificateRetrievalResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CertificatesCertificateRetrievalResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCARowIndex

`func (o *CertificatesCertificateRetrievalResponse) GetCARowIndex() int64`

GetCARowIndex returns the CARowIndex field if non-nil, zero value otherwise.

### GetCARowIndexOk

`func (o *CertificatesCertificateRetrievalResponse) GetCARowIndexOk() (*int64, bool)`

GetCARowIndexOk returns a tuple with the CARowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARowIndex

`func (o *CertificatesCertificateRetrievalResponse) SetCARowIndex(v int64)`

SetCARowIndex sets CARowIndex field to given value.

### HasCARowIndex

`func (o *CertificatesCertificateRetrievalResponse) HasCARowIndex() bool`

HasCARowIndex returns a boolean if a field has been set.

### SetCARowIndexNil

`func (o *CertificatesCertificateRetrievalResponse) SetCARowIndexNil(b bool)`

 SetCARowIndexNil sets the value for CARowIndex to be an explicit nil

### UnsetCARowIndex
`func (o *CertificatesCertificateRetrievalResponse) UnsetCARowIndex()`

UnsetCARowIndex ensures that no value is present for CARowIndex, not even an explicit nil
### GetCARecordId

`func (o *CertificatesCertificateRetrievalResponse) GetCARecordId() string`

GetCARecordId returns the CARecordId field if non-nil, zero value otherwise.

### GetCARecordIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetCARecordIdOk() (*string, bool)`

GetCARecordIdOk returns a tuple with the CARecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARecordId

`func (o *CertificatesCertificateRetrievalResponse) SetCARecordId(v string)`

SetCARecordId sets CARecordId field to given value.

### HasCARecordId

`func (o *CertificatesCertificateRetrievalResponse) HasCARecordId() bool`

HasCARecordId returns a boolean if a field has been set.

### SetCARecordIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetCARecordIdNil(b bool)`

 SetCARecordIdNil sets the value for CARecordId to be an explicit nil

### UnsetCARecordId
`func (o *CertificatesCertificateRetrievalResponse) UnsetCARecordId()`

UnsetCARecordId ensures that no value is present for CARecordId, not even an explicit nil
### GetDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) GetDetailedKeyUsage() CertificatesCertificateRetrievalResponseDetailedKeyUsageModel`

GetDetailedKeyUsage returns the DetailedKeyUsage field if non-nil, zero value otherwise.

### GetDetailedKeyUsageOk

`func (o *CertificatesCertificateRetrievalResponse) GetDetailedKeyUsageOk() (*CertificatesCertificateRetrievalResponseDetailedKeyUsageModel, bool)`

GetDetailedKeyUsageOk returns a tuple with the DetailedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) SetDetailedKeyUsage(v CertificatesCertificateRetrievalResponseDetailedKeyUsageModel)`

SetDetailedKeyUsage sets DetailedKeyUsage field to given value.

### HasDetailedKeyUsage

`func (o *CertificatesCertificateRetrievalResponse) HasDetailedKeyUsage() bool`

HasDetailedKeyUsage returns a boolean if a field has been set.

### GetKeyRecoverable

`func (o *CertificatesCertificateRetrievalResponse) GetKeyRecoverable() bool`

GetKeyRecoverable returns the KeyRecoverable field if non-nil, zero value otherwise.

### GetKeyRecoverableOk

`func (o *CertificatesCertificateRetrievalResponse) GetKeyRecoverableOk() (*bool, bool)`

GetKeyRecoverableOk returns a tuple with the KeyRecoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRecoverable

`func (o *CertificatesCertificateRetrievalResponse) SetKeyRecoverable(v bool)`

SetKeyRecoverable sets KeyRecoverable field to given value.

### HasKeyRecoverable

`func (o *CertificatesCertificateRetrievalResponse) HasKeyRecoverable() bool`

HasKeyRecoverable returns a boolean if a field has been set.

### GetCurve

`func (o *CertificatesCertificateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CertificatesCertificateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CertificatesCertificateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CertificatesCertificateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CertificatesCertificateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CertificatesCertificateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalResponse) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *CertificatesCertificateRetrievalResponse) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalResponse) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *CertificatesCertificateRetrievalResponse) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *CertificatesCertificateRetrievalResponse) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *CertificatesCertificateRetrievalResponse) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


