# CertificateStoresTypesStoreTypeProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**CSSCMSCoreEnumsCertificateStoreTypePropertyType**](CSSCMSCoreEnumsCertificateStoreTypePropertyType.md) |  | [optional] 
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**ValidationOptions** | Pointer to [**CertificateStoresTypesPropertiesValidationOptions**](CertificateStoresTypesPropertiesValidationOptions.md) |  | [optional] 

## Methods

### NewCertificateStoresTypesStoreTypeProperty

`func NewCertificateStoresTypesStoreTypeProperty() *CertificateStoresTypesStoreTypeProperty`

NewCertificateStoresTypesStoreTypeProperty instantiates a new CertificateStoresTypesStoreTypeProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresTypesStoreTypePropertyWithDefaults

`func NewCertificateStoresTypesStoreTypePropertyWithDefaults() *CertificateStoresTypesStoreTypeProperty`

NewCertificateStoresTypesStoreTypePropertyWithDefaults instantiates a new CertificateStoresTypesStoreTypeProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateStoresTypesStoreTypeProperty) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoresTypesStoreTypeProperty) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoresTypesStoreTypeProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStoreTypeId

`func (o *CertificateStoresTypesStoreTypeProperty) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *CertificateStoresTypesStoreTypeProperty) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *CertificateStoresTypesStoreTypeProperty) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetName

`func (o *CertificateStoresTypesStoreTypeProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresTypesStoreTypeProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoresTypesStoreTypeProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoresTypesStoreTypeProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoresTypesStoreTypeProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *CertificateStoresTypesStoreTypeProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateStoresTypesStoreTypeProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateStoresTypesStoreTypeProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateStoresTypesStoreTypeProperty) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateStoresTypesStoreTypeProperty) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *CertificateStoresTypesStoreTypeProperty) GetType() CSSCMSCoreEnumsCertificateStoreTypePropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetTypeOk() (*CSSCMSCoreEnumsCertificateStoreTypePropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateStoresTypesStoreTypeProperty) SetType(v CSSCMSCoreEnumsCertificateStoreTypePropertyType)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateStoresTypesStoreTypeProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDependsOn

`func (o *CertificateStoresTypesStoreTypeProperty) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CertificateStoresTypesStoreTypeProperty) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *CertificateStoresTypesStoreTypeProperty) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *CertificateStoresTypesStoreTypeProperty) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *CertificateStoresTypesStoreTypeProperty) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDefaultValue

`func (o *CertificateStoresTypesStoreTypeProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CertificateStoresTypesStoreTypeProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CertificateStoresTypesStoreTypeProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CertificateStoresTypesStoreTypeProperty) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CertificateStoresTypesStoreTypeProperty) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRequired

`func (o *CertificateStoresTypesStoreTypeProperty) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CertificateStoresTypesStoreTypeProperty) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CertificateStoresTypesStoreTypeProperty) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *CertificateStoresTypesStoreTypeProperty) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *CertificateStoresTypesStoreTypeProperty) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetValidationOptions

`func (o *CertificateStoresTypesStoreTypeProperty) GetValidationOptions() CertificateStoresTypesPropertiesValidationOptions`

GetValidationOptions returns the ValidationOptions field if non-nil, zero value otherwise.

### GetValidationOptionsOk

`func (o *CertificateStoresTypesStoreTypeProperty) GetValidationOptionsOk() (*CertificateStoresTypesPropertiesValidationOptions, bool)`

GetValidationOptionsOk returns a tuple with the ValidationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOptions

`func (o *CertificateStoresTypesStoreTypeProperty) SetValidationOptions(v CertificateStoresTypesPropertiesValidationOptions)`

SetValidationOptions sets ValidationOptions field to given value.

### HasValidationOptions

`func (o *CertificateStoresTypesStoreTypeProperty) HasValidationOptions() bool`

HasValidationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


