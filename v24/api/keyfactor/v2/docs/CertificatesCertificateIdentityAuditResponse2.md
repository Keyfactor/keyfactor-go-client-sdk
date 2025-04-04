# CertificatesCertificateIdentityAuditResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**CertificatesCertificateIdentityAuditResponse2IdentityResponse**](CertificatesCertificateIdentityAuditResponse2IdentityResponse.md) |  | [optional] 
**Permissions** | Pointer to [**[]CertificatesCertificateIdentityAuditResponse2CertificatePermission**](CertificatesCertificateIdentityAuditResponse2CertificatePermission.md) | Permissions granted to the account represented by the audit reponse on the specified certifcate. | [optional] 

## Methods

### NewCertificatesCertificateIdentityAuditResponse2

`func NewCertificatesCertificateIdentityAuditResponse2() *CertificatesCertificateIdentityAuditResponse2`

NewCertificatesCertificateIdentityAuditResponse2 instantiates a new CertificatesCertificateIdentityAuditResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateIdentityAuditResponse2WithDefaults

`func NewCertificatesCertificateIdentityAuditResponse2WithDefaults() *CertificatesCertificateIdentityAuditResponse2`

NewCertificatesCertificateIdentityAuditResponse2WithDefaults instantiates a new CertificatesCertificateIdentityAuditResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *CertificatesCertificateIdentityAuditResponse2) GetIdentity() CertificatesCertificateIdentityAuditResponse2IdentityResponse`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CertificatesCertificateIdentityAuditResponse2) GetIdentityOk() (*CertificatesCertificateIdentityAuditResponse2IdentityResponse, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CertificatesCertificateIdentityAuditResponse2) SetIdentity(v CertificatesCertificateIdentityAuditResponse2IdentityResponse)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CertificatesCertificateIdentityAuditResponse2) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetPermissions

`func (o *CertificatesCertificateIdentityAuditResponse2) GetPermissions() []CertificatesCertificateIdentityAuditResponse2CertificatePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CertificatesCertificateIdentityAuditResponse2) GetPermissionsOk() (*[]CertificatesCertificateIdentityAuditResponse2CertificatePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CertificatesCertificateIdentityAuditResponse2) SetPermissions(v []CertificatesCertificateIdentityAuditResponse2CertificatePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CertificatesCertificateIdentityAuditResponse2) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CertificatesCertificateIdentityAuditResponse2) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CertificatesCertificateIdentityAuditResponse2) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


