# ModelsCertificateStoreTypeProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateStoreTypeProperty

`func NewModelsCertificateStoreTypeProperty() *ModelsCertificateStoreTypeProperty`

NewModelsCertificateStoreTypeProperty instantiates a new ModelsCertificateStoreTypeProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreTypePropertyWithDefaults

`func NewModelsCertificateStoreTypePropertyWithDefaults() *ModelsCertificateStoreTypeProperty`

NewModelsCertificateStoreTypePropertyWithDefaults instantiates a new ModelsCertificateStoreTypeProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreTypeId

`func (o *ModelsCertificateStoreTypeProperty) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *ModelsCertificateStoreTypeProperty) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *ModelsCertificateStoreTypeProperty) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *ModelsCertificateStoreTypeProperty) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetName

`func (o *ModelsCertificateStoreTypeProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertificateStoreTypeProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertificateStoreTypeProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsCertificateStoreTypeProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsCertificateStoreTypeProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsCertificateStoreTypeProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *ModelsCertificateStoreTypeProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelsCertificateStoreTypeProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelsCertificateStoreTypeProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelsCertificateStoreTypeProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ModelsCertificateStoreTypeProperty) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ModelsCertificateStoreTypeProperty) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *ModelsCertificateStoreTypeProperty) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsCertificateStoreTypeProperty) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsCertificateStoreTypeProperty) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsCertificateStoreTypeProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDependsOn

`func (o *ModelsCertificateStoreTypeProperty) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ModelsCertificateStoreTypeProperty) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ModelsCertificateStoreTypeProperty) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ModelsCertificateStoreTypeProperty) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *ModelsCertificateStoreTypeProperty) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *ModelsCertificateStoreTypeProperty) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDefaultValue

`func (o *ModelsCertificateStoreTypeProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ModelsCertificateStoreTypeProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ModelsCertificateStoreTypeProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ModelsCertificateStoreTypeProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ModelsCertificateStoreTypeProperty) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ModelsCertificateStoreTypeProperty) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRequired

`func (o *ModelsCertificateStoreTypeProperty) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ModelsCertificateStoreTypeProperty) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ModelsCertificateStoreTypeProperty) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ModelsCertificateStoreTypeProperty) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


