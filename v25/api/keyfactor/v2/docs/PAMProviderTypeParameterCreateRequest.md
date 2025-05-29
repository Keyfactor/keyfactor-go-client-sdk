# PAMProviderTypeParameterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**DataType** | Pointer to [**CSSCMSDataModelEnumsPamParameterDataType**](CSSCMSDataModelEnumsPamParameterDataType.md) |  | [optional] 
**InstanceLevel** | Pointer to **bool** |  | [optional] 

## Methods

### NewPAMProviderTypeParameterCreateRequest

`func NewPAMProviderTypeParameterCreateRequest(name string, ) *PAMProviderTypeParameterCreateRequest`

NewPAMProviderTypeParameterCreateRequest instantiates a new PAMProviderTypeParameterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderTypeParameterCreateRequestWithDefaults

`func NewPAMProviderTypeParameterCreateRequestWithDefaults() *PAMProviderTypeParameterCreateRequest`

NewPAMProviderTypeParameterCreateRequestWithDefaults instantiates a new PAMProviderTypeParameterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PAMProviderTypeParameterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderTypeParameterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderTypeParameterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *PAMProviderTypeParameterCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PAMProviderTypeParameterCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PAMProviderTypeParameterCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PAMProviderTypeParameterCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PAMProviderTypeParameterCreateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PAMProviderTypeParameterCreateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDataType

`func (o *PAMProviderTypeParameterCreateRequest) GetDataType() CSSCMSDataModelEnumsPamParameterDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PAMProviderTypeParameterCreateRequest) GetDataTypeOk() (*CSSCMSDataModelEnumsPamParameterDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PAMProviderTypeParameterCreateRequest) SetDataType(v CSSCMSDataModelEnumsPamParameterDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *PAMProviderTypeParameterCreateRequest) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetInstanceLevel

`func (o *PAMProviderTypeParameterCreateRequest) GetInstanceLevel() bool`

GetInstanceLevel returns the InstanceLevel field if non-nil, zero value otherwise.

### GetInstanceLevelOk

`func (o *PAMProviderTypeParameterCreateRequest) GetInstanceLevelOk() (*bool, bool)`

GetInstanceLevelOk returns a tuple with the InstanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLevel

`func (o *PAMProviderTypeParameterCreateRequest) SetInstanceLevel(v bool)`

SetInstanceLevel sets InstanceLevel field to given value.

### HasInstanceLevel

`func (o *PAMProviderTypeParameterCreateRequest) HasInstanceLevel() bool`

HasInstanceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


