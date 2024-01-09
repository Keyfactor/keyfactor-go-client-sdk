# CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**CSSCMSCoreEnumsCertStoreEntryParameterType**](CSSCMSCoreEnumsCertStoreEntryParameterType.md) |  | [optional] 
**RequiredWhen** | Pointer to [**CSSCMSCoreEnumsEntryParameterUsageFlags**](CSSCMSCoreEnumsEntryParameterUsageFlags.md) |  | [optional] 
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter

`func NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter() *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter instantiates a new CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameterWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameterWithDefaults() *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

NewCSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameterWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetType() CSSCMSCoreEnumsCertStoreEntryParameterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetTypeOk() (*CSSCMSCoreEnumsCertStoreEntryParameterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetType(v CSSCMSCoreEnumsCertStoreEntryParameterType)`

SetType sets Type field to given value.

### HasType

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiredWhen

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetRequiredWhen() CSSCMSCoreEnumsEntryParameterUsageFlags`

GetRequiredWhen returns the RequiredWhen field if non-nil, zero value otherwise.

### GetRequiredWhenOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetRequiredWhenOk() (*CSSCMSCoreEnumsEntryParameterUsageFlags, bool)`

GetRequiredWhenOk returns a tuple with the RequiredWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredWhen

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetRequiredWhen(v CSSCMSCoreEnumsEntryParameterUsageFlags)`

SetRequiredWhen sets RequiredWhen field to given value.

### HasRequiredWhen

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasRequiredWhen() bool`

HasRequiredWhen returns a boolean if a field has been set.

### GetDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptions

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


