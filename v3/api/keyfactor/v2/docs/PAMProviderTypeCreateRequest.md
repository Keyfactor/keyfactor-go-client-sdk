# PAMProviderTypeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Parameters** | Pointer to [**[]PAMProviderTypeParameterCreateRequest**](PAMProviderTypeParameterCreateRequest.md) |  | [optional] 

## Methods

### NewPAMProviderTypeCreateRequest

`func NewPAMProviderTypeCreateRequest(name string, ) *PAMProviderTypeCreateRequest`

NewPAMProviderTypeCreateRequest instantiates a new PAMProviderTypeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderTypeCreateRequestWithDefaults

`func NewPAMProviderTypeCreateRequestWithDefaults() *PAMProviderTypeCreateRequest`

NewPAMProviderTypeCreateRequestWithDefaults instantiates a new PAMProviderTypeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PAMProviderTypeCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderTypeCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderTypeCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *PAMProviderTypeCreateRequest) GetParameters() []PAMProviderTypeParameterCreateRequest`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PAMProviderTypeCreateRequest) GetParametersOk() (*[]PAMProviderTypeParameterCreateRequest, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PAMProviderTypeCreateRequest) SetParameters(v []PAMProviderTypeParameterCreateRequest)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PAMProviderTypeCreateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PAMProviderTypeCreateRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PAMProviderTypeCreateRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


