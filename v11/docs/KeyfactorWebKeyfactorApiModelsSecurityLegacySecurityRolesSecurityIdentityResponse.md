# KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the security identity. | [optional] 
**AccountName** | Pointer to **NullableString** | The username associated with the account. | [optional] 
**IdentityType** | Pointer to **NullableString** | The type of the identity. | [optional] 
**Roles** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md) | The roles this identity belongs to. | [optional] 
**Valid** | Pointer to **NullableBool** | Whether or not the identity&#39;s role XML is valid. | [optional] 
**Sid** | Pointer to **NullableString** | The security identifier for the identity. | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetRoles

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetRoles() []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetRolesOk() (*[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetRoles(v []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetSid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


