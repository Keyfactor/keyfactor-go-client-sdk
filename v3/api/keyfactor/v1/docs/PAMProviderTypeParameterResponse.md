# PAMProviderTypeParameterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**DataType** | Pointer to [**CSSCMSDataModelEnumsPamParameterDataType**](CSSCMSDataModelEnumsPamParameterDataType.md) |  | [optional] 
**InstanceLevel** | Pointer to **bool** |  | [optional] 

## Methods

### NewPAMProviderTypeParameterResponse

`func NewPAMProviderTypeParameterResponse() *PAMProviderTypeParameterResponse`

NewPAMProviderTypeParameterResponse instantiates a new PAMProviderTypeParameterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMProviderTypeParameterResponseWithDefaults

`func NewPAMProviderTypeParameterResponseWithDefaults() *PAMProviderTypeParameterResponse`

NewPAMProviderTypeParameterResponseWithDefaults instantiates a new PAMProviderTypeParameterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PAMProviderTypeParameterResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PAMProviderTypeParameterResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PAMProviderTypeParameterResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PAMProviderTypeParameterResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PAMProviderTypeParameterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PAMProviderTypeParameterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PAMProviderTypeParameterResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PAMProviderTypeParameterResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PAMProviderTypeParameterResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PAMProviderTypeParameterResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *PAMProviderTypeParameterResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PAMProviderTypeParameterResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PAMProviderTypeParameterResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PAMProviderTypeParameterResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PAMProviderTypeParameterResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PAMProviderTypeParameterResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDataType

`func (o *PAMProviderTypeParameterResponse) GetDataType() CSSCMSDataModelEnumsPamParameterDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PAMProviderTypeParameterResponse) GetDataTypeOk() (*CSSCMSDataModelEnumsPamParameterDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PAMProviderTypeParameterResponse) SetDataType(v CSSCMSDataModelEnumsPamParameterDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *PAMProviderTypeParameterResponse) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetInstanceLevel

`func (o *PAMProviderTypeParameterResponse) GetInstanceLevel() bool`

GetInstanceLevel returns the InstanceLevel field if non-nil, zero value otherwise.

### GetInstanceLevelOk

`func (o *PAMProviderTypeParameterResponse) GetInstanceLevelOk() (*bool, bool)`

GetInstanceLevelOk returns a tuple with the InstanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLevel

`func (o *PAMProviderTypeParameterResponse) SetInstanceLevel(v bool)`

SetInstanceLevel sets InstanceLevel field to given value.

### HasInstanceLevel

`func (o *PAMProviderTypeParameterResponse) HasInstanceLevel() bool`

HasInstanceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


