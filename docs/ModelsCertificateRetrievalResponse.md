# ModelsCertificateRetrievalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**ImportDate** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**PrincipalId** | Pointer to **NullableInt32** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**CertState** | Pointer to **int32** |  | [optional] 
**KeySizeInBits** | Pointer to **int32** |  | [optional] 
**KeyType** | Pointer to **int32** |  | [optional] 
**RequesterId** | Pointer to **int32** |  | [optional] 
**IssuedOU** | Pointer to **NullableString** |  | [optional] 
**IssuedEmail** | Pointer to **NullableString** |  | [optional] 
**KeyUsage** | Pointer to **int32** |  | [optional] 
**SigningAlgorithm** | Pointer to **string** |  | [optional] 
**CertStateString** | Pointer to **string** |  | [optional] 
**KeyTypeString** | Pointer to **string** |  | [optional] 
**RevocationEffDate** | Pointer to **NullableString** |  | [optional] 
**RevocationReason** | Pointer to **NullableInt32** |  | [optional] 
**RevocationComment** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthorityId** | Pointer to **int32** |  | [optional] 
**CertificateAuthorityName** | Pointer to **string** |  | [optional] 
**TemplateName** | Pointer to **string** | Full template display name. | [optional] 
**ArchivedKey** | Pointer to **bool** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] 
**PrincipalName** | Pointer to **NullableString** |  | [optional] 
**CertRequestId** | Pointer to **int32** |  | [optional] 
**RequesterName** | Pointer to **string** |  | [optional] 
**ContentBytes** | Pointer to **string** |  | [optional] 
**ExtendedKeyUsages** | Pointer to [**[]ModelsCertificateRetrievalResponseExtendedKeyUsageModel**](ModelsCertificateRetrievalResponseExtendedKeyUsageModel.md) |  | [optional] 
**SubjectAltNameElements** | Pointer to [**[]ModelsCertificateRetrievalResponseSubjectAlternativeNameModel**](ModelsCertificateRetrievalResponseSubjectAlternativeNameModel.md) |  | [optional] 
**CRLDistributionPoints** | Pointer to [**[]ModelsCertificateRetrievalResponseCRLDistributionPointModel**](ModelsCertificateRetrievalResponseCRLDistributionPointModel.md) |  | [optional] 
**LocationsCount** | Pointer to [**[]ModelsCertificateRetrievalResponseLocationCountModel**](ModelsCertificateRetrievalResponseLocationCountModel.md) |  | [optional] 
**SSLLocations** | Pointer to [**[]ModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel**](ModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel.md) |  | [optional] 
**Locations** | Pointer to [**[]ModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel**](ModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CertificateKeyId** | Pointer to **int32** |  | [optional] 
**CARowIndex** | Pointer to **int64** |  | [optional] [readonly] 
**CARecordId** | Pointer to **string** |  | [optional] 
**DetailedKeyUsage** | Pointer to [**ModelsCertificateRetrievalResponseDetailedKeyUsageModel**](ModelsCertificateRetrievalResponseDetailedKeyUsageModel.md) |  | [optional] 
**KeyRecoverable** | Pointer to **bool** |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsCertificateRetrievalResponse

`func NewModelsCertificateRetrievalResponse() *ModelsCertificateRetrievalResponse`

NewModelsCertificateRetrievalResponse instantiates a new ModelsCertificateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateRetrievalResponseWithDefaults

`func NewModelsCertificateRetrievalResponseWithDefaults() *ModelsCertificateRetrievalResponse`

NewModelsCertificateRetrievalResponseWithDefaults instantiates a new ModelsCertificateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThumbprint

`func (o *ModelsCertificateRetrievalResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateRetrievalResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateRetrievalResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateRetrievalResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ModelsCertificateRetrievalResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsCertificateRetrievalResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsCertificateRetrievalResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsCertificateRetrievalResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuedDN

`func (o *ModelsCertificateRetrievalResponse) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *ModelsCertificateRetrievalResponse) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *ModelsCertificateRetrievalResponse) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *ModelsCertificateRetrievalResponse) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *ModelsCertificateRetrievalResponse) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *ModelsCertificateRetrievalResponse) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuedCN

`func (o *ModelsCertificateRetrievalResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *ModelsCertificateRetrievalResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *ModelsCertificateRetrievalResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *ModelsCertificateRetrievalResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *ModelsCertificateRetrievalResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *ModelsCertificateRetrievalResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetImportDate

`func (o *ModelsCertificateRetrievalResponse) GetImportDate() string`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *ModelsCertificateRetrievalResponse) GetImportDateOk() (*string, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *ModelsCertificateRetrievalResponse) SetImportDate(v string)`

SetImportDate sets ImportDate field to given value.

### HasImportDate

`func (o *ModelsCertificateRetrievalResponse) HasImportDate() bool`

HasImportDate returns a boolean if a field has been set.

### GetNotBefore

`func (o *ModelsCertificateRetrievalResponse) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ModelsCertificateRetrievalResponse) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ModelsCertificateRetrievalResponse) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ModelsCertificateRetrievalResponse) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *ModelsCertificateRetrievalResponse) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ModelsCertificateRetrievalResponse) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ModelsCertificateRetrievalResponse) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *ModelsCertificateRetrievalResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsCertificateRetrievalResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsCertificateRetrievalResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsCertificateRetrievalResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsCertificateRetrievalResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsCertificateRetrievalResponse) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsCertificateRetrievalResponse) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetPrincipalId

`func (o *ModelsCertificateRetrievalResponse) GetPrincipalId() int32`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *ModelsCertificateRetrievalResponse) GetPrincipalIdOk() (*int32, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *ModelsCertificateRetrievalResponse) SetPrincipalId(v int32)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *ModelsCertificateRetrievalResponse) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *ModelsCertificateRetrievalResponse) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *ModelsCertificateRetrievalResponse) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetTemplateId

`func (o *ModelsCertificateRetrievalResponse) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsCertificateRetrievalResponse) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsCertificateRetrievalResponse) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ModelsCertificateRetrievalResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetCertState

`func (o *ModelsCertificateRetrievalResponse) GetCertState() int32`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *ModelsCertificateRetrievalResponse) GetCertStateOk() (*int32, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *ModelsCertificateRetrievalResponse) SetCertState(v int32)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *ModelsCertificateRetrievalResponse) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetKeySizeInBits

`func (o *ModelsCertificateRetrievalResponse) GetKeySizeInBits() int32`

GetKeySizeInBits returns the KeySizeInBits field if non-nil, zero value otherwise.

### GetKeySizeInBitsOk

`func (o *ModelsCertificateRetrievalResponse) GetKeySizeInBitsOk() (*int32, bool)`

GetKeySizeInBitsOk returns a tuple with the KeySizeInBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizeInBits

`func (o *ModelsCertificateRetrievalResponse) SetKeySizeInBits(v int32)`

SetKeySizeInBits sets KeySizeInBits field to given value.

### HasKeySizeInBits

`func (o *ModelsCertificateRetrievalResponse) HasKeySizeInBits() bool`

HasKeySizeInBits returns a boolean if a field has been set.

### GetKeyType

`func (o *ModelsCertificateRetrievalResponse) GetKeyType() int32`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsCertificateRetrievalResponse) GetKeyTypeOk() (*int32, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsCertificateRetrievalResponse) SetKeyType(v int32)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ModelsCertificateRetrievalResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetRequesterId

`func (o *ModelsCertificateRetrievalResponse) GetRequesterId() int32`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *ModelsCertificateRetrievalResponse) GetRequesterIdOk() (*int32, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *ModelsCertificateRetrievalResponse) SetRequesterId(v int32)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *ModelsCertificateRetrievalResponse) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetIssuedOU

`func (o *ModelsCertificateRetrievalResponse) GetIssuedOU() string`

GetIssuedOU returns the IssuedOU field if non-nil, zero value otherwise.

### GetIssuedOUOk

`func (o *ModelsCertificateRetrievalResponse) GetIssuedOUOk() (*string, bool)`

GetIssuedOUOk returns a tuple with the IssuedOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOU

`func (o *ModelsCertificateRetrievalResponse) SetIssuedOU(v string)`

SetIssuedOU sets IssuedOU field to given value.

### HasIssuedOU

`func (o *ModelsCertificateRetrievalResponse) HasIssuedOU() bool`

HasIssuedOU returns a boolean if a field has been set.

### SetIssuedOUNil

`func (o *ModelsCertificateRetrievalResponse) SetIssuedOUNil(b bool)`

 SetIssuedOUNil sets the value for IssuedOU to be an explicit nil

### UnsetIssuedOU
`func (o *ModelsCertificateRetrievalResponse) UnsetIssuedOU()`

UnsetIssuedOU ensures that no value is present for IssuedOU, not even an explicit nil
### GetIssuedEmail

`func (o *ModelsCertificateRetrievalResponse) GetIssuedEmail() string`

GetIssuedEmail returns the IssuedEmail field if non-nil, zero value otherwise.

### GetIssuedEmailOk

`func (o *ModelsCertificateRetrievalResponse) GetIssuedEmailOk() (*string, bool)`

GetIssuedEmailOk returns a tuple with the IssuedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedEmail

`func (o *ModelsCertificateRetrievalResponse) SetIssuedEmail(v string)`

SetIssuedEmail sets IssuedEmail field to given value.

### HasIssuedEmail

`func (o *ModelsCertificateRetrievalResponse) HasIssuedEmail() bool`

HasIssuedEmail returns a boolean if a field has been set.

### SetIssuedEmailNil

`func (o *ModelsCertificateRetrievalResponse) SetIssuedEmailNil(b bool)`

 SetIssuedEmailNil sets the value for IssuedEmail to be an explicit nil

### UnsetIssuedEmail
`func (o *ModelsCertificateRetrievalResponse) UnsetIssuedEmail()`

UnsetIssuedEmail ensures that no value is present for IssuedEmail, not even an explicit nil
### GetKeyUsage

`func (o *ModelsCertificateRetrievalResponse) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *ModelsCertificateRetrievalResponse) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *ModelsCertificateRetrievalResponse) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *ModelsCertificateRetrievalResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *ModelsCertificateRetrievalResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *ModelsCertificateRetrievalResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *ModelsCertificateRetrievalResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *ModelsCertificateRetrievalResponse) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetCertStateString

`func (o *ModelsCertificateRetrievalResponse) GetCertStateString() string`

GetCertStateString returns the CertStateString field if non-nil, zero value otherwise.

### GetCertStateStringOk

`func (o *ModelsCertificateRetrievalResponse) GetCertStateStringOk() (*string, bool)`

GetCertStateStringOk returns a tuple with the CertStateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStateString

`func (o *ModelsCertificateRetrievalResponse) SetCertStateString(v string)`

SetCertStateString sets CertStateString field to given value.

### HasCertStateString

`func (o *ModelsCertificateRetrievalResponse) HasCertStateString() bool`

HasCertStateString returns a boolean if a field has been set.

### GetKeyTypeString

`func (o *ModelsCertificateRetrievalResponse) GetKeyTypeString() string`

GetKeyTypeString returns the KeyTypeString field if non-nil, zero value otherwise.

### GetKeyTypeStringOk

`func (o *ModelsCertificateRetrievalResponse) GetKeyTypeStringOk() (*string, bool)`

GetKeyTypeStringOk returns a tuple with the KeyTypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeString

`func (o *ModelsCertificateRetrievalResponse) SetKeyTypeString(v string)`

SetKeyTypeString sets KeyTypeString field to given value.

### HasKeyTypeString

`func (o *ModelsCertificateRetrievalResponse) HasKeyTypeString() bool`

HasKeyTypeString returns a boolean if a field has been set.

### GetRevocationEffDate

`func (o *ModelsCertificateRetrievalResponse) GetRevocationEffDate() string`

GetRevocationEffDate returns the RevocationEffDate field if non-nil, zero value otherwise.

### GetRevocationEffDateOk

`func (o *ModelsCertificateRetrievalResponse) GetRevocationEffDateOk() (*string, bool)`

GetRevocationEffDateOk returns a tuple with the RevocationEffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEffDate

`func (o *ModelsCertificateRetrievalResponse) SetRevocationEffDate(v string)`

SetRevocationEffDate sets RevocationEffDate field to given value.

### HasRevocationEffDate

`func (o *ModelsCertificateRetrievalResponse) HasRevocationEffDate() bool`

HasRevocationEffDate returns a boolean if a field has been set.

### SetRevocationEffDateNil

`func (o *ModelsCertificateRetrievalResponse) SetRevocationEffDateNil(b bool)`

 SetRevocationEffDateNil sets the value for RevocationEffDate to be an explicit nil

### UnsetRevocationEffDate
`func (o *ModelsCertificateRetrievalResponse) UnsetRevocationEffDate()`

UnsetRevocationEffDate ensures that no value is present for RevocationEffDate, not even an explicit nil
### GetRevocationReason

`func (o *ModelsCertificateRetrievalResponse) GetRevocationReason() int32`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *ModelsCertificateRetrievalResponse) GetRevocationReasonOk() (*int32, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *ModelsCertificateRetrievalResponse) SetRevocationReason(v int32)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *ModelsCertificateRetrievalResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *ModelsCertificateRetrievalResponse) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *ModelsCertificateRetrievalResponse) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil
### GetRevocationComment

`func (o *ModelsCertificateRetrievalResponse) GetRevocationComment() string`

GetRevocationComment returns the RevocationComment field if non-nil, zero value otherwise.

### GetRevocationCommentOk

`func (o *ModelsCertificateRetrievalResponse) GetRevocationCommentOk() (*string, bool)`

GetRevocationCommentOk returns a tuple with the RevocationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationComment

`func (o *ModelsCertificateRetrievalResponse) SetRevocationComment(v string)`

SetRevocationComment sets RevocationComment field to given value.

### HasRevocationComment

`func (o *ModelsCertificateRetrievalResponse) HasRevocationComment() bool`

HasRevocationComment returns a boolean if a field has been set.

### SetRevocationCommentNil

`func (o *ModelsCertificateRetrievalResponse) SetRevocationCommentNil(b bool)`

 SetRevocationCommentNil sets the value for RevocationComment to be an explicit nil

### UnsetRevocationComment
`func (o *ModelsCertificateRetrievalResponse) UnsetRevocationComment()`

UnsetRevocationComment ensures that no value is present for RevocationComment, not even an explicit nil
### GetCertificateAuthorityId

`func (o *ModelsCertificateRetrievalResponse) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *ModelsCertificateRetrievalResponse) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *ModelsCertificateRetrievalResponse) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *ModelsCertificateRetrievalResponse) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetCertificateAuthorityName

`func (o *ModelsCertificateRetrievalResponse) GetCertificateAuthorityName() string`

GetCertificateAuthorityName returns the CertificateAuthorityName field if non-nil, zero value otherwise.

### GetCertificateAuthorityNameOk

`func (o *ModelsCertificateRetrievalResponse) GetCertificateAuthorityNameOk() (*string, bool)`

GetCertificateAuthorityNameOk returns a tuple with the CertificateAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityName

`func (o *ModelsCertificateRetrievalResponse) SetCertificateAuthorityName(v string)`

SetCertificateAuthorityName sets CertificateAuthorityName field to given value.

### HasCertificateAuthorityName

`func (o *ModelsCertificateRetrievalResponse) HasCertificateAuthorityName() bool`

HasCertificateAuthorityName returns a boolean if a field has been set.

### GetTemplateName

`func (o *ModelsCertificateRetrievalResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ModelsCertificateRetrievalResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ModelsCertificateRetrievalResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ModelsCertificateRetrievalResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetArchivedKey

`func (o *ModelsCertificateRetrievalResponse) GetArchivedKey() bool`

GetArchivedKey returns the ArchivedKey field if non-nil, zero value otherwise.

### GetArchivedKeyOk

`func (o *ModelsCertificateRetrievalResponse) GetArchivedKeyOk() (*bool, bool)`

GetArchivedKeyOk returns a tuple with the ArchivedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedKey

`func (o *ModelsCertificateRetrievalResponse) SetArchivedKey(v bool)`

SetArchivedKey sets ArchivedKey field to given value.

### HasArchivedKey

`func (o *ModelsCertificateRetrievalResponse) HasArchivedKey() bool`

HasArchivedKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *ModelsCertificateRetrievalResponse) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *ModelsCertificateRetrievalResponse) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *ModelsCertificateRetrievalResponse) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *ModelsCertificateRetrievalResponse) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrincipalName

`func (o *ModelsCertificateRetrievalResponse) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *ModelsCertificateRetrievalResponse) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *ModelsCertificateRetrievalResponse) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *ModelsCertificateRetrievalResponse) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *ModelsCertificateRetrievalResponse) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *ModelsCertificateRetrievalResponse) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetCertRequestId

`func (o *ModelsCertificateRetrievalResponse) GetCertRequestId() int32`

GetCertRequestId returns the CertRequestId field if non-nil, zero value otherwise.

### GetCertRequestIdOk

`func (o *ModelsCertificateRetrievalResponse) GetCertRequestIdOk() (*int32, bool)`

GetCertRequestIdOk returns a tuple with the CertRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRequestId

`func (o *ModelsCertificateRetrievalResponse) SetCertRequestId(v int32)`

SetCertRequestId sets CertRequestId field to given value.

### HasCertRequestId

`func (o *ModelsCertificateRetrievalResponse) HasCertRequestId() bool`

HasCertRequestId returns a boolean if a field has been set.

### GetRequesterName

`func (o *ModelsCertificateRetrievalResponse) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *ModelsCertificateRetrievalResponse) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *ModelsCertificateRetrievalResponse) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *ModelsCertificateRetrievalResponse) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetContentBytes

`func (o *ModelsCertificateRetrievalResponse) GetContentBytes() string`

GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.

### GetContentBytesOk

`func (o *ModelsCertificateRetrievalResponse) GetContentBytesOk() (*string, bool)`

GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBytes

`func (o *ModelsCertificateRetrievalResponse) SetContentBytes(v string)`

SetContentBytes sets ContentBytes field to given value.

### HasContentBytes

`func (o *ModelsCertificateRetrievalResponse) HasContentBytes() bool`

HasContentBytes returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *ModelsCertificateRetrievalResponse) GetExtendedKeyUsages() []ModelsCertificateRetrievalResponseExtendedKeyUsageModel`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *ModelsCertificateRetrievalResponse) GetExtendedKeyUsagesOk() (*[]ModelsCertificateRetrievalResponseExtendedKeyUsageModel, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *ModelsCertificateRetrievalResponse) SetExtendedKeyUsages(v []ModelsCertificateRetrievalResponseExtendedKeyUsageModel)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *ModelsCertificateRetrievalResponse) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### GetSubjectAltNameElements

`func (o *ModelsCertificateRetrievalResponse) GetSubjectAltNameElements() []ModelsCertificateRetrievalResponseSubjectAlternativeNameModel`

GetSubjectAltNameElements returns the SubjectAltNameElements field if non-nil, zero value otherwise.

### GetSubjectAltNameElementsOk

`func (o *ModelsCertificateRetrievalResponse) GetSubjectAltNameElementsOk() (*[]ModelsCertificateRetrievalResponseSubjectAlternativeNameModel, bool)`

GetSubjectAltNameElementsOk returns a tuple with the SubjectAltNameElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNameElements

`func (o *ModelsCertificateRetrievalResponse) SetSubjectAltNameElements(v []ModelsCertificateRetrievalResponseSubjectAlternativeNameModel)`

SetSubjectAltNameElements sets SubjectAltNameElements field to given value.

### HasSubjectAltNameElements

`func (o *ModelsCertificateRetrievalResponse) HasSubjectAltNameElements() bool`

HasSubjectAltNameElements returns a boolean if a field has been set.

### GetCRLDistributionPoints

`func (o *ModelsCertificateRetrievalResponse) GetCRLDistributionPoints() []ModelsCertificateRetrievalResponseCRLDistributionPointModel`

GetCRLDistributionPoints returns the CRLDistributionPoints field if non-nil, zero value otherwise.

### GetCRLDistributionPointsOk

`func (o *ModelsCertificateRetrievalResponse) GetCRLDistributionPointsOk() (*[]ModelsCertificateRetrievalResponseCRLDistributionPointModel, bool)`

GetCRLDistributionPointsOk returns a tuple with the CRLDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCRLDistributionPoints

`func (o *ModelsCertificateRetrievalResponse) SetCRLDistributionPoints(v []ModelsCertificateRetrievalResponseCRLDistributionPointModel)`

SetCRLDistributionPoints sets CRLDistributionPoints field to given value.

### HasCRLDistributionPoints

`func (o *ModelsCertificateRetrievalResponse) HasCRLDistributionPoints() bool`

HasCRLDistributionPoints returns a boolean if a field has been set.

### GetLocationsCount

`func (o *ModelsCertificateRetrievalResponse) GetLocationsCount() []ModelsCertificateRetrievalResponseLocationCountModel`

GetLocationsCount returns the LocationsCount field if non-nil, zero value otherwise.

### GetLocationsCountOk

`func (o *ModelsCertificateRetrievalResponse) GetLocationsCountOk() (*[]ModelsCertificateRetrievalResponseLocationCountModel, bool)`

GetLocationsCountOk returns a tuple with the LocationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsCount

`func (o *ModelsCertificateRetrievalResponse) SetLocationsCount(v []ModelsCertificateRetrievalResponseLocationCountModel)`

SetLocationsCount sets LocationsCount field to given value.

### HasLocationsCount

`func (o *ModelsCertificateRetrievalResponse) HasLocationsCount() bool`

HasLocationsCount returns a boolean if a field has been set.

### GetSSLLocations

`func (o *ModelsCertificateRetrievalResponse) GetSSLLocations() []ModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel`

GetSSLLocations returns the SSLLocations field if non-nil, zero value otherwise.

### GetSSLLocationsOk

`func (o *ModelsCertificateRetrievalResponse) GetSSLLocationsOk() (*[]ModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel, bool)`

GetSSLLocationsOk returns a tuple with the SSLLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLLocations

`func (o *ModelsCertificateRetrievalResponse) SetSSLLocations(v []ModelsCertificateRetrievalResponseCertificateStoreLocationDetailModel)`

SetSSLLocations sets SSLLocations field to given value.

### HasSSLLocations

`func (o *ModelsCertificateRetrievalResponse) HasSSLLocations() bool`

HasSSLLocations returns a boolean if a field has been set.

### GetLocations

`func (o *ModelsCertificateRetrievalResponse) GetLocations() []ModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ModelsCertificateRetrievalResponse) GetLocationsOk() (*[]ModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ModelsCertificateRetrievalResponse) SetLocations(v []ModelsCertificateRetrievalResponseCertificateStoreInventoryItemModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ModelsCertificateRetrievalResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsCertificateRetrievalResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsCertificateRetrievalResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsCertificateRetrievalResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsCertificateRetrievalResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCertificateKeyId

`func (o *ModelsCertificateRetrievalResponse) GetCertificateKeyId() int32`

GetCertificateKeyId returns the CertificateKeyId field if non-nil, zero value otherwise.

### GetCertificateKeyIdOk

`func (o *ModelsCertificateRetrievalResponse) GetCertificateKeyIdOk() (*int32, bool)`

GetCertificateKeyIdOk returns a tuple with the CertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKeyId

`func (o *ModelsCertificateRetrievalResponse) SetCertificateKeyId(v int32)`

SetCertificateKeyId sets CertificateKeyId field to given value.

### HasCertificateKeyId

`func (o *ModelsCertificateRetrievalResponse) HasCertificateKeyId() bool`

HasCertificateKeyId returns a boolean if a field has been set.

### GetCARowIndex

`func (o *ModelsCertificateRetrievalResponse) GetCARowIndex() int64`

GetCARowIndex returns the CARowIndex field if non-nil, zero value otherwise.

### GetCARowIndexOk

`func (o *ModelsCertificateRetrievalResponse) GetCARowIndexOk() (*int64, bool)`

GetCARowIndexOk returns a tuple with the CARowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARowIndex

`func (o *ModelsCertificateRetrievalResponse) SetCARowIndex(v int64)`

SetCARowIndex sets CARowIndex field to given value.

### HasCARowIndex

`func (o *ModelsCertificateRetrievalResponse) HasCARowIndex() bool`

HasCARowIndex returns a boolean if a field has been set.

### GetCARecordId

`func (o *ModelsCertificateRetrievalResponse) GetCARecordId() string`

GetCARecordId returns the CARecordId field if non-nil, zero value otherwise.

### GetCARecordIdOk

`func (o *ModelsCertificateRetrievalResponse) GetCARecordIdOk() (*string, bool)`

GetCARecordIdOk returns a tuple with the CARecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARecordId

`func (o *ModelsCertificateRetrievalResponse) SetCARecordId(v string)`

SetCARecordId sets CARecordId field to given value.

### HasCARecordId

`func (o *ModelsCertificateRetrievalResponse) HasCARecordId() bool`

HasCARecordId returns a boolean if a field has been set.

### GetDetailedKeyUsage

`func (o *ModelsCertificateRetrievalResponse) GetDetailedKeyUsage() ModelsCertificateRetrievalResponseDetailedKeyUsageModel`

GetDetailedKeyUsage returns the DetailedKeyUsage field if non-nil, zero value otherwise.

### GetDetailedKeyUsageOk

`func (o *ModelsCertificateRetrievalResponse) GetDetailedKeyUsageOk() (*ModelsCertificateRetrievalResponseDetailedKeyUsageModel, bool)`

GetDetailedKeyUsageOk returns a tuple with the DetailedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedKeyUsage

`func (o *ModelsCertificateRetrievalResponse) SetDetailedKeyUsage(v ModelsCertificateRetrievalResponseDetailedKeyUsageModel)`

SetDetailedKeyUsage sets DetailedKeyUsage field to given value.

### HasDetailedKeyUsage

`func (o *ModelsCertificateRetrievalResponse) HasDetailedKeyUsage() bool`

HasDetailedKeyUsage returns a boolean if a field has been set.

### GetKeyRecoverable

`func (o *ModelsCertificateRetrievalResponse) GetKeyRecoverable() bool`

GetKeyRecoverable returns the KeyRecoverable field if non-nil, zero value otherwise.

### GetKeyRecoverableOk

`func (o *ModelsCertificateRetrievalResponse) GetKeyRecoverableOk() (*bool, bool)`

GetKeyRecoverableOk returns a tuple with the KeyRecoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRecoverable

`func (o *ModelsCertificateRetrievalResponse) SetKeyRecoverable(v bool)`

SetKeyRecoverable sets KeyRecoverable field to given value.

### HasKeyRecoverable

`func (o *ModelsCertificateRetrievalResponse) HasKeyRecoverable() bool`

HasKeyRecoverable returns a boolean if a field has been set.

### GetCurve

`func (o *ModelsCertificateRetrievalResponse) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *ModelsCertificateRetrievalResponse) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *ModelsCertificateRetrievalResponse) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *ModelsCertificateRetrievalResponse) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *ModelsCertificateRetrievalResponse) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *ModelsCertificateRetrievalResponse) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


