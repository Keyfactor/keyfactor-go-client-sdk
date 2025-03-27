# CertificatesRevocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokedIds** | Pointer to **[]int32** |  | [optional] 
**SuspendedCerts** | Pointer to [**[]CertificatesSuspendedRevocationResponse**](CertificatesSuspendedRevocationResponse.md) |  | [optional] 

## Methods

### NewCertificatesRevocationResponse

`func NewCertificatesRevocationResponse() *CertificatesRevocationResponse`

NewCertificatesRevocationResponse instantiates a new CertificatesRevocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesRevocationResponseWithDefaults

`func NewCertificatesRevocationResponseWithDefaults() *CertificatesRevocationResponse`

NewCertificatesRevocationResponseWithDefaults instantiates a new CertificatesRevocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokedIds

`func (o *CertificatesRevocationResponse) GetRevokedIds() []int32`

GetRevokedIds returns the RevokedIds field if non-nil, zero value otherwise.

### GetRevokedIdsOk

`func (o *CertificatesRevocationResponse) GetRevokedIdsOk() (*[]int32, bool)`

GetRevokedIdsOk returns a tuple with the RevokedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedIds

`func (o *CertificatesRevocationResponse) SetRevokedIds(v []int32)`

SetRevokedIds sets RevokedIds field to given value.

### HasRevokedIds

`func (o *CertificatesRevocationResponse) HasRevokedIds() bool`

HasRevokedIds returns a boolean if a field has been set.

### SetRevokedIdsNil

`func (o *CertificatesRevocationResponse) SetRevokedIdsNil(b bool)`

 SetRevokedIdsNil sets the value for RevokedIds to be an explicit nil

### UnsetRevokedIds
`func (o *CertificatesRevocationResponse) UnsetRevokedIds()`

UnsetRevokedIds ensures that no value is present for RevokedIds, not even an explicit nil
### GetSuspendedCerts

`func (o *CertificatesRevocationResponse) GetSuspendedCerts() []CertificatesSuspendedRevocationResponse`

GetSuspendedCerts returns the SuspendedCerts field if non-nil, zero value otherwise.

### GetSuspendedCertsOk

`func (o *CertificatesRevocationResponse) GetSuspendedCertsOk() (*[]CertificatesSuspendedRevocationResponse, bool)`

GetSuspendedCertsOk returns a tuple with the SuspendedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedCerts

`func (o *CertificatesRevocationResponse) SetSuspendedCerts(v []CertificatesSuspendedRevocationResponse)`

SetSuspendedCerts sets SuspendedCerts field to given value.

### HasSuspendedCerts

`func (o *CertificatesRevocationResponse) HasSuspendedCerts() bool`

HasSuspendedCerts returns a boolean if a field has been set.

### SetSuspendedCertsNil

`func (o *CertificatesRevocationResponse) SetSuspendedCertsNil(b bool)`

 SetSuspendedCertsNil sets the value for SuspendedCerts to be an explicit nil

### UnsetSuspendedCerts
`func (o *CertificatesRevocationResponse) UnsetSuspendedCerts()`

UnsetSuspendedCerts ensures that no value is present for SuspendedCerts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


