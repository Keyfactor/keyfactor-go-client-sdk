# PAMProviderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Remote** | **bool** |  | 
**ProviderTypeParamValues** | Pointer to [**[]PAMProviderCreateRequestTypeParamValue**](PAMProviderCreateRequestTypeParamValue.md) |  | [optional] 

## Methods

### NewPAMProviderUpdateRequest

`func NewPAMProviderUpdateRequest(id int32, name string, remote bool, ) *PAMProviderUpdateRequest`

NewPAMProviderUpdateRequest instantiates a new PAMProviderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderUpdateRequestWithDefaults

`func NewPAMProviderUpdateRequestWithDefaults() *PAMProviderUpdateRequest`

NewPAMProviderUpdateRequestWithDefaults instantiates a new PAMProviderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PAMProviderUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRemote

`func (o *PAMProviderUpdateRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PAMProviderUpdateRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PAMProviderUpdateRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.


### GetProviderTypeParamValues

`func (o *PAMProviderUpdateRequest) GetProviderTypeParamValues() []PAMProviderCreateRequestTypeParamValue`

GetProviderTypeParamValues returns the ProviderTypeParamValues field if non-nil, zero value otherwise.

### GetProviderTypeParamValuesOk

`func (o *PAMProviderUpdateRequest) GetProviderTypeParamValuesOk() (*[]PAMProviderCreateRequestTypeParamValue, bool)`

GetProviderTypeParamValuesOk returns a tuple with the ProviderTypeParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeParamValues

`func (o *PAMProviderUpdateRequest) SetProviderTypeParamValues(v []PAMProviderCreateRequestTypeParamValue)`

SetProviderTypeParamValues sets ProviderTypeParamValues field to given value.

### HasProviderTypeParamValues

`func (o *PAMProviderUpdateRequest) HasProviderTypeParamValues() bool`

HasProviderTypeParamValues returns a boolean if a field has been set.

### SetProviderTypeParamValuesNil

`func (o *PAMProviderUpdateRequest) SetProviderTypeParamValuesNil(b bool)`

 SetProviderTypeParamValuesNil sets the value for ProviderTypeParamValues to be an explicit nil

### UnsetProviderTypeParamValues
`func (o *PAMProviderUpdateRequest) UnsetProviderTypeParamValues()`

UnsetProviderTypeParamValues ensures that no value is present for ProviderTypeParamValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


