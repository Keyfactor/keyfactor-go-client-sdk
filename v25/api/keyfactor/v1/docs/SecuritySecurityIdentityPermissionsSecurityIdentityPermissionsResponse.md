# SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | Pointer to [**SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md) |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**SecuredAreaPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**CollectionPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**ContainerPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**PamProviderPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**IdentityProviderPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**PamPermissions** | Pointer to [**[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](SecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 

## Methods

### NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse

`func NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse() *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse`

NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse instantiates a new SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults

`func NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults() *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse`

NewSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults instantiates a new SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaim() SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaimOk() (*SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetClaim(v SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### GetIdentity

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetSecuredAreaPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetSecuredAreaPermissions returns the SecuredAreaPermissions field if non-nil, zero value otherwise.

### GetSecuredAreaPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetSecuredAreaPermissionsOk returns a tuple with the SecuredAreaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetSecuredAreaPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetSecuredAreaPermissions sets SecuredAreaPermissions field to given value.

### HasSecuredAreaPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasSecuredAreaPermissions() bool`

HasSecuredAreaPermissions returns a boolean if a field has been set.

### SetSecuredAreaPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetSecuredAreaPermissionsNil(b bool)`

 SetSecuredAreaPermissionsNil sets the value for SecuredAreaPermissions to be an explicit nil

### UnsetSecuredAreaPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetSecuredAreaPermissions()`

UnsetSecuredAreaPermissions ensures that no value is present for SecuredAreaPermissions, not even an explicit nil
### GetCollectionPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetCollectionPermissions returns the CollectionPermissions field if non-nil, zero value otherwise.

### GetCollectionPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetCollectionPermissionsOk returns a tuple with the CollectionPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetCollectionPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetCollectionPermissions sets CollectionPermissions field to given value.

### HasCollectionPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasCollectionPermissions() bool`

HasCollectionPermissions returns a boolean if a field has been set.

### SetCollectionPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetCollectionPermissionsNil(b bool)`

 SetCollectionPermissionsNil sets the value for CollectionPermissions to be an explicit nil

### UnsetCollectionPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetCollectionPermissions()`

UnsetCollectionPermissions ensures that no value is present for CollectionPermissions, not even an explicit nil
### GetContainerPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetContainerPermissions returns the ContainerPermissions field if non-nil, zero value otherwise.

### GetContainerPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetContainerPermissionsOk returns a tuple with the ContainerPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetContainerPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetContainerPermissions sets ContainerPermissions field to given value.

### HasContainerPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasContainerPermissions() bool`

HasContainerPermissions returns a boolean if a field has been set.

### SetContainerPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetContainerPermissionsNil(b bool)`

 SetContainerPermissionsNil sets the value for ContainerPermissions to be an explicit nil

### UnsetContainerPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetContainerPermissions()`

UnsetContainerPermissions ensures that no value is present for ContainerPermissions, not even an explicit nil
### GetPamProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetPamProviderPermissions returns the PamProviderPermissions field if non-nil, zero value otherwise.

### GetPamProviderPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetPamProviderPermissionsOk returns a tuple with the PamProviderPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPamProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamProviderPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetPamProviderPermissions sets PamProviderPermissions field to given value.

### HasPamProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamProviderPermissions() bool`

HasPamProviderPermissions returns a boolean if a field has been set.

### SetPamProviderPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamProviderPermissionsNil(b bool)`

 SetPamProviderPermissionsNil sets the value for PamProviderPermissions to be an explicit nil

### UnsetPamProviderPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetPamProviderPermissions()`

UnsetPamProviderPermissions ensures that no value is present for PamProviderPermissions, not even an explicit nil
### GetIdentityProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetIdentityProviderPermissions returns the IdentityProviderPermissions field if non-nil, zero value otherwise.

### GetIdentityProviderPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetIdentityProviderPermissionsOk returns a tuple with the IdentityProviderPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityProviderPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetIdentityProviderPermissions sets IdentityProviderPermissions field to given value.

### HasIdentityProviderPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentityProviderPermissions() bool`

HasIdentityProviderPermissions returns a boolean if a field has been set.

### SetIdentityProviderPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityProviderPermissionsNil(b bool)`

 SetIdentityProviderPermissionsNil sets the value for IdentityProviderPermissions to be an explicit nil

### UnsetIdentityProviderPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetIdentityProviderPermissions()`

UnsetIdentityProviderPermissions ensures that no value is present for IdentityProviderPermissions, not even an explicit nil
### GetPamPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissions() []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetPamPermissions returns the PamPermissions field if non-nil, zero value otherwise.

### GetPamPermissionsOk

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissionsOk() (*[]SecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetPamPermissionsOk returns a tuple with the PamPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPamPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamPermissions(v []SecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetPamPermissions sets PamPermissions field to given value.

### HasPamPermissions

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamPermissions() bool`

HasPamPermissions returns a boolean if a field has been set.

### SetPamPermissionsNil

`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamPermissionsNil(b bool)`

 SetPamPermissionsNil sets the value for PamPermissions to be an explicit nil

### UnsetPamPermissions
`func (o *SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetPamPermissions()`

UnsetPamPermissions ensures that no value is present for PamPermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


