# CertificatesCertificateIdentityAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Id of the account represented by the audit response | [optional] 
**AccountName** | Pointer to **NullableString** | Name of the account represented by the audit response | [optional] 
**IdentityType** | Pointer to **NullableString** | The type of account represented by the audit response (User or Group) | [optional] 
**SID** | Pointer to **NullableString** | The SID of the account represented by the audit reponse | [optional] 
**Permissions** | Pointer to [**[]CertificatesCertificateIdentityAuditResponseCertificatePermission**](CertificatesCertificateIdentityAuditResponseCertificatePermission.md) | Permissions granted to the account represented by the audit reponse on the specified certifcate | [optional] [readonly] 

## Methods

### NewCertificatesCertificateIdentityAuditResponse

`func NewCertificatesCertificateIdentityAuditResponse() *CertificatesCertificateIdentityAuditResponse`

NewCertificatesCertificateIdentityAuditResponse instantiates a new CertificatesCertificateIdentityAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateIdentityAuditResponseWithDefaults

`func NewCertificatesCertificateIdentityAuditResponseWithDefaults() *CertificatesCertificateIdentityAuditResponse`

NewCertificatesCertificateIdentityAuditResponseWithDefaults instantiates a new CertificatesCertificateIdentityAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertificateIdentityAuditResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertificateIdentityAuditResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertificateIdentityAuditResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertificateIdentityAuditResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *CertificatesCertificateIdentityAuditResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CertificatesCertificateIdentityAuditResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CertificatesCertificateIdentityAuditResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *CertificatesCertificateIdentityAuditResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *CertificatesCertificateIdentityAuditResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *CertificatesCertificateIdentityAuditResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetIdentityType

`func (o *CertificatesCertificateIdentityAuditResponse) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *CertificatesCertificateIdentityAuditResponse) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *CertificatesCertificateIdentityAuditResponse) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *CertificatesCertificateIdentityAuditResponse) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *CertificatesCertificateIdentityAuditResponse) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *CertificatesCertificateIdentityAuditResponse) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetSID

`func (o *CertificatesCertificateIdentityAuditResponse) GetSID() string`

GetSID returns the SID field if non-nil, zero value otherwise.

### GetSIDOk

`func (o *CertificatesCertificateIdentityAuditResponse) GetSIDOk() (*string, bool)`

GetSIDOk returns a tuple with the SID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSID

`func (o *CertificatesCertificateIdentityAuditResponse) SetSID(v string)`

SetSID sets SID field to given value.

### HasSID

`func (o *CertificatesCertificateIdentityAuditResponse) HasSID() bool`

HasSID returns a boolean if a field has been set.

### SetSIDNil

`func (o *CertificatesCertificateIdentityAuditResponse) SetSIDNil(b bool)`

 SetSIDNil sets the value for SID to be an explicit nil

### UnsetSID
`func (o *CertificatesCertificateIdentityAuditResponse) UnsetSID()`

UnsetSID ensures that no value is present for SID, not even an explicit nil
### GetPermissions

`func (o *CertificatesCertificateIdentityAuditResponse) GetPermissions() []CertificatesCertificateIdentityAuditResponseCertificatePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CertificatesCertificateIdentityAuditResponse) GetPermissionsOk() (*[]CertificatesCertificateIdentityAuditResponseCertificatePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CertificatesCertificateIdentityAuditResponse) SetPermissions(v []CertificatesCertificateIdentityAuditResponseCertificatePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CertificatesCertificateIdentityAuditResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CertificatesCertificateIdentityAuditResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CertificatesCertificateIdentityAuditResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


