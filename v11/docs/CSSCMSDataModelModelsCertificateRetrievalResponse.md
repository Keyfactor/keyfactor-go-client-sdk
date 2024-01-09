# CSSCMSDataModelModelsCertificateRetrievalResponse

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
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**CertState** | Pointer to [**KeyfactorPKIEnumsCertificateState**](KeyfactorPKIEnumsCertificateState.md) |  | [optional] 
**KeySizeInBits** | Pointer to **int32** |  | [optional] 
**KeyType** | Pointer to [**CSSCMSCoreEnumsEncryptionKeyType**](CSSCMSCoreEnumsEncryptionKeyType.md) |  | [optional] 
**RequesterId** | Pointer to **NullableInt32** |  | [optional] 
**IssuedOU** | Pointer to **NullableString** |  | [optional] 
**IssuedEmail** | Pointer to **NullableString** |  | [optional] 
**KeyUsage** | Pointer to **NullableInt32** |  | [optional] 
**SigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CertStateString** | Pointer to **NullableString** |  | [optional] 
**KeyTypeString** | Pointer to **NullableString** |  | [optional] 
**RevocationEffDate** | Pointer to **NullableTime** |  | [optional] 
**RevocationReason** | Pointer to [**KeyfactorPKIEnumsRevokeCode**](KeyfactorPKIEnumsRevokeCode.md) |  | [optional] 
**RevocationComment** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthorityId** | Pointer to **NullableInt32** |  | [optional] 
**CertificateAuthorityName** | Pointer to **NullableString** |  | [optional] 
**TemplateName** | Pointer to **NullableString** |  | [optional] 
**ArchivedKey** | Pointer to **bool** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] 
**PrincipalName** | Pointer to **NullableString** |  | [optional] 
**CertRequestId** | Pointer to **NullableInt32** |  | [optional] 
**RequesterName** | Pointer to **NullableString** |  | [optional] 
**ContentBytes** | Pointer to **NullableString** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseExtendedKeyUsageModel**](CSSCMSDataModelModelsCertificateRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseSubjectAlternativeNameModel**](CSSCMSDataModelModelsCertificateRetrievalResponseSubjectAlternativeNameModel.md) |  | [optional] 
**CrlDistributionPoints** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseCRLDistributionPointModel**](CSSCMSDataModelModelsCertificateRetrievalResponseCRLDistributionPointModel.md) |  | [optional] 
**LocationsCount** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseLocationCountModel**](CSSCMSDataModelModelsCertificateRetrievalResponseLocationCountModel.md) |  | [optional] 
**SslLocations** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel**](CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel.md) |  | [optional] 
**Locations** | Pointer to [**[]CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel**](CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CertificateKeyId** | Pointer to **int32** |  | [optional] 
**CaRowIndex** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**CaRecordId** | Pointer to **NullableString** |  | [optional] 
**DetailedKeyUsage** | Pointer to [**CSSCMSDataModelModelsCertificateRetrievalResponseDetailedKeyUsageModel**](CSSCMSDataModelModelsCertificateRetrievalResponseDetailedKeyUsageModel.md) |  | [optional] 
**KeyRecoverable** | Pointer to **bool** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateRetrievalResponse

`func NewCSSCMSDataModelModelsCertificateRetrievalResponse() *CSSCMSDataModelModelsCertificateRetrievalResponse`

NewCSSCMSDataModelModelsCertificateRetrievalResponse instantiates a new CSSCMSDataModelModelsCertificateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateRetrievalResponseWithDefaults

`func NewCSSCMSDataModelModelsCertificateRetrievalResponseWithDefaults() *CSSCMSDataModelModelsCertificateRetrievalResponse`

NewCSSCMSDataModelModelsCertificateRetrievalResponseWithDefaults instantiates a new CSSCMSDataModelModelsCertificateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThumbprint

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetSerialNumber

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuedCN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetImportDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetImportDate() time.Time`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetImportDateOk() (*time.Time, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetImportDate(v time.Time)`

SetImportDate sets ImportDate field to given value.

### HasImportDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasImportDate() bool`

HasImportDate returns a boolean if a field has been set.

### GetNotBefore

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetPrincipalId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetPrincipalId() int32`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetPrincipalIdOk() (*int32, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetPrincipalId(v int32)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetTemplateId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetCertState

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertState() KeyfactorPKIEnumsCertificateState`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertStateOk() (*KeyfactorPKIEnumsCertificateState, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertState(v KeyfactorPKIEnumsCertificateState)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetKeySizeInBits

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeySizeInBits() int32`

GetKeySizeInBits returns the KeySizeInBits field if non-nil, zero value otherwise.

### GetKeySizeInBitsOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeySizeInBitsOk() (*int32, bool)`

GetKeySizeInBitsOk returns a tuple with the KeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeInBits

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeySizeInBits(v int32)`

SetKeySizeInBits sets KeySizeInBits field to given value.

### HasKeySizeInBits

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasKeySizeInBits() bool`

HasKeySizeInBits returns a boolean if a field has been set.

### GetKeyType

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyType() CSSCMSCoreEnumsEncryptionKeyType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyTypeOk() (*CSSCMSCoreEnumsEncryptionKeyType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyType(v CSSCMSCoreEnumsEncryptionKeyType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetRequesterId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRequesterId() int32`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRequesterIdOk() (*int32, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRequesterId(v int32)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### SetRequesterIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRequesterIdNil(b bool)`

 SetRequesterIdNil sets the value for RequesterId to be an explicit nil

### UnsetRequesterId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetRequesterId()`

UnsetRequesterId ensures that no value is present for RequesterId, not even an explicit nil
### GetIssuedOU

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedOU() string`

GetIssuedOU returns the IssuedOU field if non-nil, zero value otherwise.

### GetIssuedOUOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedOUOk() (*string, bool)`

GetIssuedOUOk returns a tuple with the IssuedOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOU

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedOU(v string)`

SetIssuedOU sets IssuedOU field to given value.

### HasIssuedOU

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasIssuedOU() bool`

HasIssuedOU returns a boolean if a field has been set.

### SetIssuedOUNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedOUNil(b bool)`

 SetIssuedOUNil sets the value for IssuedOU to be an explicit nil

### UnsetIssuedOU
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetIssuedOU()`

UnsetIssuedOU ensures that no value is present for IssuedOU, not even an explicit nil
### GetIssuedEmail

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedEmail() string`

GetIssuedEmail returns the IssuedEmail field if non-nil, zero value otherwise.

### GetIssuedEmailOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetIssuedEmailOk() (*string, bool)`

GetIssuedEmailOk returns a tuple with the IssuedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedEmail

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedEmail(v string)`

SetIssuedEmail sets IssuedEmail field to given value.

### HasIssuedEmail

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasIssuedEmail() bool`

HasIssuedEmail returns a boolean if a field has been set.

### SetIssuedEmailNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetIssuedEmailNil(b bool)`

 SetIssuedEmailNil sets the value for IssuedEmail to be an explicit nil

### UnsetIssuedEmail
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetIssuedEmail()`

UnsetIssuedEmail ensures that no value is present for IssuedEmail, not even an explicit nil
### GetKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### SetKeyUsageNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyUsageNil(b bool)`

 SetKeyUsageNil sets the value for KeyUsage to be an explicit nil

### UnsetKeyUsage
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetKeyUsage()`

UnsetKeyUsage ensures that no value is present for KeyUsage, not even an explicit nil
### GetSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetCertStateString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertStateString() string`

GetCertStateString returns the CertStateString field if non-nil, zero value otherwise.

### GetCertStateStringOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertStateStringOk() (*string, bool)`

GetCertStateStringOk returns a tuple with the CertStateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStateString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertStateString(v string)`

SetCertStateString sets CertStateString field to given value.

### HasCertStateString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertStateString() bool`

HasCertStateString returns a boolean if a field has been set.

### SetCertStateStringNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertStateStringNil(b bool)`

 SetCertStateStringNil sets the value for CertStateString to be an explicit nil

### UnsetCertStateString
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCertStateString()`

UnsetCertStateString ensures that no value is present for CertStateString, not even an explicit nil
### GetKeyTypeString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyTypeString() string`

GetKeyTypeString returns the KeyTypeString field if non-nil, zero value otherwise.

### GetKeyTypeStringOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyTypeStringOk() (*string, bool)`

GetKeyTypeStringOk returns a tuple with the KeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyTypeString(v string)`

SetKeyTypeString sets KeyTypeString field to given value.

### HasKeyTypeString

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasKeyTypeString() bool`

HasKeyTypeString returns a boolean if a field has been set.

### SetKeyTypeStringNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyTypeStringNil(b bool)`

 SetKeyTypeStringNil sets the value for KeyTypeString to be an explicit nil

### UnsetKeyTypeString
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetKeyTypeString()`

UnsetKeyTypeString ensures that no value is present for KeyTypeString, not even an explicit nil
### GetRevocationEffDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationEffDate() time.Time`

GetRevocationEffDate returns the RevocationEffDate field if non-nil, zero value otherwise.

### GetRevocationEffDateOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationEffDateOk() (*time.Time, bool)`

GetRevocationEffDateOk returns a tuple with the RevocationEffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEffDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRevocationEffDate(v time.Time)`

SetRevocationEffDate sets RevocationEffDate field to given value.

### HasRevocationEffDate

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasRevocationEffDate() bool`

HasRevocationEffDate returns a boolean if a field has been set.

### SetRevocationEffDateNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRevocationEffDateNil(b bool)`

 SetRevocationEffDateNil sets the value for RevocationEffDate to be an explicit nil

### UnsetRevocationEffDate
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetRevocationEffDate()`

UnsetRevocationEffDate ensures that no value is present for RevocationEffDate, not even an explicit nil
### GetRevocationReason

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationReason() KeyfactorPKIEnumsRevokeCode`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationReasonOk() (*KeyfactorPKIEnumsRevokeCode, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRevocationReason(v KeyfactorPKIEnumsRevokeCode)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### GetRevocationComment

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationComment() string`

GetRevocationComment returns the RevocationComment field if non-nil, zero value otherwise.

### GetRevocationCommentOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRevocationCommentOk() (*string, bool)`

GetRevocationCommentOk returns a tuple with the RevocationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationComment

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRevocationComment(v string)`

SetRevocationComment sets RevocationComment field to given value.

### HasRevocationComment

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasRevocationComment() bool`

HasRevocationComment returns a boolean if a field has been set.

### SetRevocationCommentNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRevocationCommentNil(b bool)`

 SetRevocationCommentNil sets the value for RevocationComment to be an explicit nil

### UnsetRevocationComment
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetRevocationComment()`

UnsetRevocationComment ensures that no value is present for RevocationComment, not even an explicit nil
### GetCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### SetCertificateAuthorityIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertificateAuthorityIdNil(b bool)`

 SetCertificateAuthorityIdNil sets the value for CertificateAuthorityId to be an explicit nil

### UnsetCertificateAuthorityId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCertificateAuthorityId()`

UnsetCertificateAuthorityId ensures that no value is present for CertificateAuthorityId, not even an explicit nil
### GetCertificateAuthorityName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateAuthorityName() string`

GetCertificateAuthorityName returns the CertificateAuthorityName field if non-nil, zero value otherwise.

### GetCertificateAuthorityNameOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateAuthorityNameOk() (*string, bool)`

GetCertificateAuthorityNameOk returns a tuple with the CertificateAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertificateAuthorityName(v string)`

SetCertificateAuthorityName sets CertificateAuthorityName field to given value.

### HasCertificateAuthorityName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertificateAuthorityName() bool`

HasCertificateAuthorityName returns a boolean if a field has been set.

### SetCertificateAuthorityNameNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertificateAuthorityNameNil(b bool)`

 SetCertificateAuthorityNameNil sets the value for CertificateAuthorityName to be an explicit nil

### UnsetCertificateAuthorityName
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCertificateAuthorityName()`

UnsetCertificateAuthorityName ensures that no value is present for CertificateAuthorityName, not even an explicit nil
### GetTemplateName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetArchivedKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetArchivedKey() bool`

GetArchivedKey returns the ArchivedKey field if non-nil, zero value otherwise.

### GetArchivedKeyOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetArchivedKeyOk() (*bool, bool)`

GetArchivedKeyOk returns a tuple with the ArchivedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetArchivedKey(v bool)`

SetArchivedKey sets ArchivedKey field to given value.

### HasArchivedKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasArchivedKey() bool`

HasArchivedKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrincipalName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetCertRequestId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertRequestId() int32`

GetCertRequestId returns the CertRequestId field if non-nil, zero value otherwise.

### GetCertRequestIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertRequestIdOk() (*int32, bool)`

GetCertRequestIdOk returns a tuple with the CertRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRequestId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertRequestId(v int32)`

SetCertRequestId sets CertRequestId field to given value.

### HasCertRequestId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertRequestId() bool`

HasCertRequestId returns a boolean if a field has been set.

### SetCertRequestIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertRequestIdNil(b bool)`

 SetCertRequestIdNil sets the value for CertRequestId to be an explicit nil

### UnsetCertRequestId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCertRequestId()`

UnsetCertRequestId ensures that no value is present for CertRequestId, not even an explicit nil
### GetRequesterName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetContentBytes

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetContentBytes() string`

GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.

### GetContentBytesOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetContentBytesOk() (*string, bool)`

GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBytes

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetContentBytes(v string)`

SetContentBytes sets ContentBytes field to given value.

### HasContentBytes

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasContentBytes() bool`

HasContentBytes returns a boolean if a field has been set.

### SetContentBytesNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetContentBytesNil(b bool)`

 SetContentBytesNil sets the value for ContentBytes to be an explicit nil

### UnsetContentBytes
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetContentBytes()`

UnsetContentBytes ensures that no value is present for ContentBytes, not even an explicit nil
### GetExtendedKeyUsages

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetExtendedKeyUsages() []CSSCMSDataModelModelsCertificateRetrievalResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetExtendedKeyUsages(v []CSSCMSDataModelModelsCertificateRetrievalResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### SetExtendedKeyUsagesNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetExtendedKeyUsagesNil(b bool)`

 SetExtendedKeyUsagesNil sets the value for ExtendedKeyUsages to be an explicit nil

### UnsetExtendedKeyUsages
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetExtendedKeyUsages()`

UnsetExtendedKeyUsages ensures that no value is present for ExtendedKeyUsages, not even an explicit nil
### GetSubjectAltNameElements

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSubjectAltNameElements() []CSSCMSDataModelModelsCertificateRetrievalResponseSubjectAlternativeNameModel`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSubjectAltNameElementsOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseSubjectAlternativeNameModel, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSubjectAltNameElements(v []CSSCMSDataModelModelsCertificateRetrievalResponseSubjectAlternativeNameModel)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.

### SetSubjectAltNameElementsNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSubjectAltNameElementsNil(b bool)`

 SetSubjectAltNameElementsNil sets the value for SubjectAltNameElements to be an explicit nil

### UnsetSubjectAltNameElements
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetSubjectAltNameElements()`

UnsetSubjectAltNameElements ensures that no value is present for SubjectAltNameElements, not even an explicit nil
### GetCrlDistributionPoints

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCrlDistributionPoints() []CSSCMSDataModelModelsCertificateRetrievalResponseCRLDistributionPointModel`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCrlDistributionPointsOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseCRLDistributionPointModel, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCrlDistributionPoints(v []CSSCMSDataModelModelsCertificateRetrievalResponseCRLDistributionPointModel)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.

### HasCrlDistributionPoints

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.

### SetCrlDistributionPointsNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCrlDistributionPointsNil(b bool)`

 SetCrlDistributionPointsNil sets the value for CrlDistributionPoints to be an explicit nil

### UnsetCrlDistributionPoints
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCrlDistributionPoints()`

UnsetCrlDistributionPoints ensures that no value is present for CrlDistributionPoints, not even an explicit nil
### GetLocationsCount

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetLocationsCount() []CSSCMSDataModelModelsCertificateRetrievalResponseLocationCountModel`

GetLocationsCount returns the LocationsCount field if non-nil, zero value otherwise.

### GetLocationsCountOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetLocationsCountOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseLocationCountModel, bool)`

GetLocationsCountOk returns a tuple with the LocationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsCount

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetLocationsCount(v []CSSCMSDataModelModelsCertificateRetrievalResponseLocationCountModel)`

SetLocationsCount sets LocationsCount field to given value.

### HasLocationsCount

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasLocationsCount() bool`

HasLocationsCount returns a boolean if a field has been set.

### SetLocationsCountNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetLocationsCountNil(b bool)`

 SetLocationsCountNil sets the value for LocationsCount to be an explicit nil

### UnsetLocationsCount
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetLocationsCount()`

UnsetLocationsCount ensures that no value is present for LocationsCount, not even an explicit nil
### GetSslLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSslLocations() []CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel`

GetSslLocations returns the SslLocations field if non-nil, zero value otherwise.

### GetSslLocationsOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetSslLocationsOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel, bool)`

GetSslLocationsOk returns a tuple with the SslLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSslLocations(v []CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel)`

SetSslLocations sets SslLocations field to given value.

### HasSslLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasSslLocations() bool`

HasSslLocations returns a boolean if a field has been set.

### SetSslLocationsNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetSslLocationsNil(b bool)`

 SetSslLocationsNil sets the value for SslLocations to be an explicit nil

### UnsetSslLocations
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetSslLocations()`

UnsetSslLocations ensures that no value is present for SslLocations, not even an explicit nil
### GetLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetLocations() []CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetLocationsOk() (*[]CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetLocations(v []CSSCMSDataModelModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetMetadata

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCertificateKeyId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateKeyId() int32`

GetCertificateKeyId returns the CertificateKeyId field if non-nil, zero value otherwise.

### GetCertificateKeyIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCertificateKeyIdOk() (*int32, bool)`

GetCertificateKeyIdOk returns a tuple with the CertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKeyId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCertificateKeyId(v int32)`

SetCertificateKeyId sets CertificateKeyId field to given value.

### HasCertificateKeyId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCertificateKeyId() bool`

HasCertificateKeyId returns a boolean if a field has been set.

### GetCaRowIndex

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCaRowIndex() int64`

GetCaRowIndex returns the CaRowIndex field if non-nil, zero value otherwise.

### GetCaRowIndexOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCaRowIndexOk() (*int64, bool)`

GetCaRowIndexOk returns a tuple with the CaRowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRowIndex

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCaRowIndex(v int64)`

SetCaRowIndex sets CaRowIndex field to given value.

### HasCaRowIndex

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCaRowIndex() bool`

HasCaRowIndex returns a boolean if a field has been set.

### SetCaRowIndexNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCaRowIndexNil(b bool)`

 SetCaRowIndexNil sets the value for CaRowIndex to be an explicit nil

### UnsetCaRowIndex
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCaRowIndex()`

UnsetCaRowIndex ensures that no value is present for CaRowIndex, not even an explicit nil
### GetCaRecordId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCaRecordId() string`

GetCaRecordId returns the CaRecordId field if non-nil, zero value otherwise.

### GetCaRecordIdOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCaRecordIdOk() (*string, bool)`

GetCaRecordIdOk returns a tuple with the CaRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRecordId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCaRecordId(v string)`

SetCaRecordId sets CaRecordId field to given value.

### HasCaRecordId

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCaRecordId() bool`

HasCaRecordId returns a boolean if a field has been set.

### SetCaRecordIdNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCaRecordIdNil(b bool)`

 SetCaRecordIdNil sets the value for CaRecordId to be an explicit nil

### UnsetCaRecordId
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCaRecordId()`

UnsetCaRecordId ensures that no value is present for CaRecordId, not even an explicit nil
### GetDetailedKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetDetailedKeyUsage() CSSCMSDataModelModelsCertificateRetrievalResponseDetailedKeyUsageModel`

GetDetailedKeyUsage returns the DetailedKeyUsage field if non-nil, zero value otherwise.

### GetDetailedKeyUsageOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetDetailedKeyUsageOk() (*CSSCMSDataModelModelsCertificateRetrievalResponseDetailedKeyUsageModel, bool)`

GetDetailedKeyUsageOk returns a tuple with the DetailedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetDetailedKeyUsage(v CSSCMSDataModelModelsCertificateRetrievalResponseDetailedKeyUsageModel)`

SetDetailedKeyUsage sets DetailedKeyUsage field to given value.

### HasDetailedKeyUsage

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasDetailedKeyUsage() bool`

HasDetailedKeyUsage returns a boolean if a field has been set.

### GetKeyRecoverable

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyRecoverable() bool`

GetKeyRecoverable returns the KeyRecoverable field if non-nil, zero value otherwise.

### GetKeyRecoverableOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetKeyRecoverableOk() (*bool, bool)`

GetKeyRecoverableOk returns a tuple with the KeyRecoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRecoverable

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetKeyRecoverable(v bool)`

SetKeyRecoverable sets KeyRecoverable field to given value.

### HasKeyRecoverable

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasKeyRecoverable() bool`

HasKeyRecoverable returns a boolean if a field has been set.

### GetCurve

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CSSCMSDataModelModelsCertificateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


