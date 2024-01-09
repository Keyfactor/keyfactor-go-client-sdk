# CSSCMSDataModelModelsCertificateStoreTypeProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**CSSCMSCoreEnumsCertificateStoreTypePropertyType**](CSSCMSCoreEnumsCertificateStoreTypePropertyType.md) |  | [optional] 
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreTypeProperty

`func NewCSSCMSDataModelModelsCertificateStoreTypeProperty() *CSSCMSDataModelModelsCertificateStoreTypeProperty`

NewCSSCMSDataModelModelsCertificateStoreTypeProperty instantiates a new CSSCMSDataModelModelsCertificateStoreTypeProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreTypePropertyWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreTypePropertyWithDefaults() *CSSCMSDataModelModelsCertificateStoreTypeProperty`

NewCSSCMSDataModelModelsCertificateStoreTypePropertyWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreTypeProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetType() CSSCMSCoreEnumsCertificateStoreTypePropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetTypeOk() (*CSSCMSCoreEnumsCertificateStoreTypePropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetType(v CSSCMSCoreEnumsCertificateStoreTypePropertyType)`

SetType sets Type field to given value.

### HasType

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRequired

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CSSCMSDataModelModelsCertificateStoreTypeProperty) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


