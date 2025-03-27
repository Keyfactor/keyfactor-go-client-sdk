# PAMProviderTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]PAMProviderTypeParameterResponse**](PAMProviderTypeParameterResponse.md) |  | [optional] 

## Methods

### NewPAMProviderTypeResponse

`func NewPAMProviderTypeResponse() *PAMProviderTypeResponse`

NewPAMProviderTypeResponse instantiates a new PAMProviderTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderTypeResponseWithDefaults

`func NewPAMProviderTypeResponseWithDefaults() *PAMProviderTypeResponse`

NewPAMProviderTypeResponseWithDefaults instantiates a new PAMProviderTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderTypeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PAMProviderTypeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PAMProviderTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PAMProviderTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PAMProviderTypeResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PAMProviderTypeResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParameters

`func (o *PAMProviderTypeResponse) GetParameters() []PAMProviderTypeParameterResponse`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PAMProviderTypeResponse) GetParametersOk() (*[]PAMProviderTypeParameterResponse, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PAMProviderTypeResponse) SetParameters(v []PAMProviderTypeParameterResponse)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PAMProviderTypeResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PAMProviderTypeResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PAMProviderTypeResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


