# KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Id of the account represented by the audit response | [optional] 
**AccountName** | Pointer to **string** | Name of the account represented by the audit response | [optional] 
**IdentityType** | Pointer to **string** | The type of account represented by the audit response (User or Group) | [optional] 
**SID** | Pointer to **string** | The SID of the account represented by the audit reponse | [optional] 
**Permissions** | Pointer to [**[]KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission**](KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission.md) | Permissions granted to the account represented by the audit reponse on the specified certifcate | [optional] [readonly] 

## Methods

### NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse

`func NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse() *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse`

NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse instantiates a new KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults

`func NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults() *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse`

NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults instantiates a new KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetIdentityType

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetSID

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSID() string`

GetSID returns the SID field if non-nil, zero value otherwise.

### GetSIDOk

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSIDOk() (*string, bool)`

GetSIDOk returns a tuple with the SID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSID

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetSID(v string)`

SetSID sets SID field to given value.

### HasSID

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasSID() bool`

HasSID returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissions() []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissionsOk() (*[]KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetPermissions(v []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


