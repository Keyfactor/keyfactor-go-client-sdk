# KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | Pointer to [**KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md) |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**SecuredAreaPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**CollectionPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**ContainerPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**PamProviderPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**IdentityProviderPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 
**PamPermissions** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse

`func NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse() *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse`

NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse instantiates a new KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse`

NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaim() KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaimOk() (*KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetClaim(v KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### GetIdentity

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetSecuredAreaPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetSecuredAreaPermissions returns the SecuredAreaPermissions field if non-nil, zero value otherwise.

### GetSecuredAreaPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetSecuredAreaPermissionsOk returns a tuple with the SecuredAreaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuredAreaPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetSecuredAreaPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetSecuredAreaPermissions sets SecuredAreaPermissions field to given value.

### HasSecuredAreaPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasSecuredAreaPermissions() bool`

HasSecuredAreaPermissions returns a boolean if a field has been set.

### SetSecuredAreaPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetSecuredAreaPermissionsNil(b bool)`

 SetSecuredAreaPermissionsNil sets the value for SecuredAreaPermissions to be an explicit nil

### UnsetSecuredAreaPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetSecuredAreaPermissions()`

UnsetSecuredAreaPermissions ensures that no value is present for SecuredAreaPermissions, not even an explicit nil
### GetCollectionPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetCollectionPermissions returns the CollectionPermissions field if non-nil, zero value otherwise.

### GetCollectionPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetCollectionPermissionsOk returns a tuple with the CollectionPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetCollectionPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetCollectionPermissions sets CollectionPermissions field to given value.

### HasCollectionPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasCollectionPermissions() bool`

HasCollectionPermissions returns a boolean if a field has been set.

### SetCollectionPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetCollectionPermissionsNil(b bool)`

 SetCollectionPermissionsNil sets the value for CollectionPermissions to be an explicit nil

### UnsetCollectionPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetCollectionPermissions()`

UnsetCollectionPermissions ensures that no value is present for CollectionPermissions, not even an explicit nil
### GetContainerPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetContainerPermissions returns the ContainerPermissions field if non-nil, zero value otherwise.

### GetContainerPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetContainerPermissionsOk returns a tuple with the ContainerPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetContainerPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetContainerPermissions sets ContainerPermissions field to given value.

### HasContainerPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasContainerPermissions() bool`

HasContainerPermissions returns a boolean if a field has been set.

### SetContainerPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetContainerPermissionsNil(b bool)`

 SetContainerPermissionsNil sets the value for ContainerPermissions to be an explicit nil

### UnsetContainerPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetContainerPermissions()`

UnsetContainerPermissions ensures that no value is present for ContainerPermissions, not even an explicit nil
### GetPamProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetPamProviderPermissions returns the PamProviderPermissions field if non-nil, zero value otherwise.

### GetPamProviderPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetPamProviderPermissionsOk returns a tuple with the PamProviderPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPamProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamProviderPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetPamProviderPermissions sets PamProviderPermissions field to given value.

### HasPamProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamProviderPermissions() bool`

HasPamProviderPermissions returns a boolean if a field has been set.

### SetPamProviderPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamProviderPermissionsNil(b bool)`

 SetPamProviderPermissionsNil sets the value for PamProviderPermissions to be an explicit nil

### UnsetPamProviderPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetPamProviderPermissions()`

UnsetPamProviderPermissions ensures that no value is present for PamProviderPermissions, not even an explicit nil
### GetIdentityProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetIdentityProviderPermissions returns the IdentityProviderPermissions field if non-nil, zero value otherwise.

### GetIdentityProviderPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetIdentityProviderPermissionsOk returns a tuple with the IdentityProviderPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityProviderPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetIdentityProviderPermissions sets IdentityProviderPermissions field to given value.

### HasIdentityProviderPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentityProviderPermissions() bool`

HasIdentityProviderPermissions returns a boolean if a field has been set.

### SetIdentityProviderPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityProviderPermissionsNil(b bool)`

 SetIdentityProviderPermissionsNil sets the value for IdentityProviderPermissions to be an explicit nil

### UnsetIdentityProviderPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetIdentityProviderPermissions()`

UnsetIdentityProviderPermissions ensures that no value is present for IdentityProviderPermissions, not even an explicit nil
### GetPamPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse`

GetPamPermissions returns the PamPermissions field if non-nil, zero value otherwise.

### GetPamPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissionsOk() (*[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool)`

GetPamPermissionsOk returns a tuple with the PamPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPamPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse)`

SetPamPermissions sets PamPermissions field to given value.

### HasPamPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamPermissions() bool`

HasPamPermissions returns a boolean if a field has been set.

### SetPamPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamPermissionsNil(b bool)`

 SetPamPermissionsNil sets the value for PamPermissions to be an explicit nil

### UnsetPamPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetPamPermissions()`

UnsetPamPermissions ensures that no value is present for PamPermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


