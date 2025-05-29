# IdentityProviderProviderTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TypeParameters** | Pointer to [**[]IdentityProviderProviderTypeParameterResponse**](IdentityProviderProviderTypeParameterResponse.md) |  | [optional] 

## Methods

### NewIdentityProviderProviderTypeResponse

`func NewIdentityProviderProviderTypeResponse() *IdentityProviderProviderTypeResponse`

NewIdentityProviderProviderTypeResponse instantiates a new IdentityProviderProviderTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderProviderTypeResponseWithDefaults

`func NewIdentityProviderProviderTypeResponseWithDefaults() *IdentityProviderProviderTypeResponse`

NewIdentityProviderProviderTypeResponseWithDefaults instantiates a new IdentityProviderProviderTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderProviderTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderProviderTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderProviderTypeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderProviderTypeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderProviderTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderProviderTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderProviderTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProviderProviderTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityProviderProviderTypeResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityProviderProviderTypeResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTypeParameters

`func (o *IdentityProviderProviderTypeResponse) GetTypeParameters() []IdentityProviderProviderTypeParameterResponse`

GetTypeParameters returns the TypeParameters field if non-nil, zero value otherwise.

### GetTypeParametersOk

`func (o *IdentityProviderProviderTypeResponse) GetTypeParametersOk() (*[]IdentityProviderProviderTypeParameterResponse, bool)`

GetTypeParametersOk returns a tuple with the TypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeParameters

`func (o *IdentityProviderProviderTypeResponse) SetTypeParameters(v []IdentityProviderProviderTypeParameterResponse)`

SetTypeParameters sets TypeParameters field to given value.

### HasTypeParameters

`func (o *IdentityProviderProviderTypeResponse) HasTypeParameters() bool`

HasTypeParameters returns a boolean if a field has been set.

### SetTypeParametersNil

`func (o *IdentityProviderProviderTypeResponse) SetTypeParametersNil(b bool)`

 SetTypeParametersNil sets the value for TypeParameters to be an explicit nil

### UnsetTypeParameters
`func (o *IdentityProviderProviderTypeResponse) UnsetTypeParameters()`

UnsetTypeParameters ensures that no value is present for TypeParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


