# CertificateStoresTypesEntryParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**CSSCMSCoreEnumsCertStoreEntryParameterType**](CSSCMSCoreEnumsCertStoreEntryParameterType.md) |  | [optional] 
**ValidationOptions** | Pointer to [**CertificateStoresTypesEntryParametersValidationOptions**](CertificateStoresTypesEntryParametersValidationOptions.md) |  | [optional] 
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewCertificateStoresTypesEntryParameters

`func NewCertificateStoresTypesEntryParameters() *CertificateStoresTypesEntryParameters`

NewCertificateStoresTypesEntryParameters instantiates a new CertificateStoresTypesEntryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresTypesEntryParametersWithDefaults

`func NewCertificateStoresTypesEntryParametersWithDefaults() *CertificateStoresTypesEntryParameters`

NewCertificateStoresTypesEntryParametersWithDefaults instantiates a new CertificateStoresTypesEntryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreTypeId

`func (o *CertificateStoresTypesEntryParameters) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *CertificateStoresTypesEntryParameters) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *CertificateStoresTypesEntryParameters) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *CertificateStoresTypesEntryParameters) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetName

`func (o *CertificateStoresTypesEntryParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresTypesEntryParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresTypesEntryParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoresTypesEntryParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoresTypesEntryParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoresTypesEntryParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *CertificateStoresTypesEntryParameters) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateStoresTypesEntryParameters) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateStoresTypesEntryParameters) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateStoresTypesEntryParameters) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateStoresTypesEntryParameters) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateStoresTypesEntryParameters) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *CertificateStoresTypesEntryParameters) GetType() CSSCMSCoreEnumsCertStoreEntryParameterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateStoresTypesEntryParameters) GetTypeOk() (*CSSCMSCoreEnumsCertStoreEntryParameterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateStoresTypesEntryParameters) SetType(v CSSCMSCoreEnumsCertStoreEntryParameterType)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateStoresTypesEntryParameters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationOptions

`func (o *CertificateStoresTypesEntryParameters) GetValidationOptions() CertificateStoresTypesEntryParametersValidationOptions`

GetValidationOptions returns the ValidationOptions field if non-nil, zero value otherwise.

### GetValidationOptionsOk

`func (o *CertificateStoresTypesEntryParameters) GetValidationOptionsOk() (*CertificateStoresTypesEntryParametersValidationOptions, bool)`

GetValidationOptionsOk returns a tuple with the ValidationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOptions

`func (o *CertificateStoresTypesEntryParameters) SetValidationOptions(v CertificateStoresTypesEntryParametersValidationOptions)`

SetValidationOptions sets ValidationOptions field to given value.

### HasValidationOptions

`func (o *CertificateStoresTypesEntryParameters) HasValidationOptions() bool`

HasValidationOptions returns a boolean if a field has been set.

### GetDependsOn

`func (o *CertificateStoresTypesEntryParameters) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CertificateStoresTypesEntryParameters) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CertificateStoresTypesEntryParameters) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *CertificateStoresTypesEntryParameters) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *CertificateStoresTypesEntryParameters) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *CertificateStoresTypesEntryParameters) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDefaultValue

`func (o *CertificateStoresTypesEntryParameters) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CertificateStoresTypesEntryParameters) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CertificateStoresTypesEntryParameters) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CertificateStoresTypesEntryParameters) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CertificateStoresTypesEntryParameters) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CertificateStoresTypesEntryParameters) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptions

`func (o *CertificateStoresTypesEntryParameters) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CertificateStoresTypesEntryParameters) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CertificateStoresTypesEntryParameters) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CertificateStoresTypesEntryParameters) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CertificateStoresTypesEntryParameters) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CertificateStoresTypesEntryParameters) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *CertificateStoresTypesEntryParameters) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoresTypesEntryParameters) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoresTypesEntryParameters) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoresTypesEntryParameters) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


