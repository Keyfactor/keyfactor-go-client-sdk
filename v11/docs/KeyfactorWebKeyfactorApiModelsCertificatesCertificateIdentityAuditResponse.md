# KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Id of the account represented by the audit response | [optional] 
**AccountName** | Pointer to **NullableString** | Name of the account represented by the audit response | [optional] 
**IdentityType** | Pointer to **NullableString** | The type of account represented by the audit response (User or Group) | [optional] 
**Sid** | Pointer to **NullableString** | The SID of the account represented by the audit reponse | [optional] 
**Permissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission**](KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission.md) | Permissions granted to the account represented by the audit reponse on the specified certifcate | [optional] [readonly] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetSid

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissions() []KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetPermissions(v []KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


