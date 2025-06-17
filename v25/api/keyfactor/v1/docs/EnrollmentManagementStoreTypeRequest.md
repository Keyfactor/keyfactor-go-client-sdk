# EnrollmentManagementStoreTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreTypeId** | Pointer to **int32** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewEnrollmentManagementStoreTypeRequest

`func NewEnrollmentManagementStoreTypeRequest() *EnrollmentManagementStoreTypeRequest`

NewEnrollmentManagementStoreTypeRequest instantiates a new EnrollmentManagementStoreTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentManagementStoreTypeRequestWithDefaults

`func NewEnrollmentManagementStoreTypeRequestWithDefaults() *EnrollmentManagementStoreTypeRequest`

NewEnrollmentManagementStoreTypeRequestWithDefaults instantiates a new EnrollmentManagementStoreTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreTypeId

`func (o *EnrollmentManagementStoreTypeRequest) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *EnrollmentManagementStoreTypeRequest) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *EnrollmentManagementStoreTypeRequest) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *EnrollmentManagementStoreTypeRequest) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetAlias

`func (o *EnrollmentManagementStoreTypeRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EnrollmentManagementStoreTypeRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EnrollmentManagementStoreTypeRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *EnrollmentManagementStoreTypeRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *EnrollmentManagementStoreTypeRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *EnrollmentManagementStoreTypeRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetOverwrite

`func (o *EnrollmentManagementStoreTypeRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *EnrollmentManagementStoreTypeRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *EnrollmentManagementStoreTypeRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *EnrollmentManagementStoreTypeRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetProperties

`func (o *EnrollmentManagementStoreTypeRequest) GetProperties() []interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EnrollmentManagementStoreTypeRequest) GetPropertiesOk() (*[]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EnrollmentManagementStoreTypeRequest) SetProperties(v []interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EnrollmentManagementStoreTypeRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EnrollmentManagementStoreTypeRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EnrollmentManagementStoreTypeRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


