# KeyfactorApiPAMProviderTypeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Parameters** | Pointer to [**[]KeyfactorApiPAMProviderTypeParameterCreateRequest**](KeyfactorApiPAMProviderTypeParameterCreateRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiPAMProviderTypeCreateRequest

`func NewKeyfactorApiPAMProviderTypeCreateRequest(name string, ) *KeyfactorApiPAMProviderTypeCreateRequest`

NewKeyfactorApiPAMProviderTypeCreateRequest instantiates a new KeyfactorApiPAMProviderTypeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiPAMProviderTypeCreateRequestWithDefaults

`func NewKeyfactorApiPAMProviderTypeCreateRequestWithDefaults() *KeyfactorApiPAMProviderTypeCreateRequest`

NewKeyfactorApiPAMProviderTypeCreateRequestWithDefaults instantiates a new KeyfactorApiPAMProviderTypeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) GetParameters() []KeyfactorApiPAMProviderTypeParameterCreateRequest`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) GetParametersOk() (*[]KeyfactorApiPAMProviderTypeParameterCreateRequest, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) SetParameters(v []KeyfactorApiPAMProviderTypeParameterCreateRequest)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *KeyfactorApiPAMProviderTypeCreateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


