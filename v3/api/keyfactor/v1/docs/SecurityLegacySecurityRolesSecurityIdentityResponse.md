# SecurityLegacySecurityRolesSecurityIdentityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the security identity. | [optional] 
**AccountName** | Pointer to **NullableString** | The username associated with the account. | [optional] 
**IdentityType** | Pointer to **NullableString** | The type of the identity. | [optional] 
**Roles** | Pointer to [**[]SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md) | The roles this identity belongs to. | [optional] 
**Valid** | Pointer to **NullableBool** | Whether or not the identity&#39;s role XML is valid. | [optional] 
**SID** | Pointer to **NullableString** | The security identifier for the identity. | [optional] 

## Methods

### NewSecurityLegacySecurityRolesSecurityIdentityResponse

`func NewSecurityLegacySecurityRolesSecurityIdentityResponse() *SecurityLegacySecurityRolesSecurityIdentityResponse`

NewSecurityLegacySecurityRolesSecurityIdentityResponse instantiates a new SecurityLegacySecurityRolesSecurityIdentityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults

`func NewSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults() *SecurityLegacySecurityRolesSecurityIdentityResponse`

NewSecurityLegacySecurityRolesSecurityIdentityResponseWithDefaults instantiates a new SecurityLegacySecurityRolesSecurityIdentityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetIdentityType

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetRoles

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetRoles() []SecurityLegacySecurityRolesSecurityRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetRolesOk() (*[]SecurityLegacySecurityRolesSecurityRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetRoles(v []SecurityLegacySecurityRolesSecurityRoleResponse)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetValid

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetSID

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetSID() string`

GetSID returns the SID field if non-nil, zero value otherwise.

### GetSIDOk

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) GetSIDOk() (*string, bool)`

GetSIDOk returns a tuple with the SID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSID

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetSID(v string)`

SetSID sets SID field to given value.

### HasSID

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) HasSID() bool`

HasSID returns a boolean if a field has been set.

### SetSIDNil

`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) SetSIDNil(b bool)`

 SetSIDNil sets the value for SID to be an explicit nil

### UnsetSID
`func (o *SecurityLegacySecurityRolesSecurityIdentityResponse) UnsetSID()`

UnsetSID ensures that no value is present for SID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


