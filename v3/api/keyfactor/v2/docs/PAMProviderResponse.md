# PAMProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProviderType** | Pointer to [**CSSCMSDataModelModelsProviderType**](CSSCMSDataModelModelsProviderType.md) |  | [optional] 
**ProviderTypeParamValues** | Pointer to [**[]PAMPamProviderTypeParamValueResponse**](PAMPamProviderTypeParamValueResponse.md) |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 

## Methods

### NewPAMProviderResponse

`func NewPAMProviderResponse() *PAMProviderResponse`

NewPAMProviderResponse instantiates a new PAMProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderResponseWithDefaults

`func NewPAMProviderResponseWithDefaults() *PAMProviderResponse`

NewPAMProviderResponseWithDefaults instantiates a new PAMProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PAMProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PAMProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PAMProviderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PAMProviderResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PAMProviderResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderType

`func (o *PAMProviderResponse) GetProviderType() CSSCMSDataModelModelsProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PAMProviderResponse) GetProviderTypeOk() (*CSSCMSDataModelModelsProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PAMProviderResponse) SetProviderType(v CSSCMSDataModelModelsProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *PAMProviderResponse) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetProviderTypeParamValues

`func (o *PAMProviderResponse) GetProviderTypeParamValues() []PAMPamProviderTypeParamValueResponse`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *PAMProviderResponse) GetProviderTypeParamValuesOk() (*[]PAMPamProviderTypeParamValueResponse, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *PAMProviderResponse) SetProviderTypeParamValues(v []PAMPamProviderTypeParamValueResponse)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *PAMProviderResponse) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *PAMProviderResponse) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *PAMProviderResponse) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil
### GetRemote

`func (o *PAMProviderResponse) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PAMProviderResponse) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PAMProviderResponse) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PAMProviderResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


