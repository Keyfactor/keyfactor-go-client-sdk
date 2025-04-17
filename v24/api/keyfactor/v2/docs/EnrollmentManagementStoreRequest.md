# EnrollmentManagementStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnrollmentManagementStoreRequest

`func NewEnrollmentManagementStoreRequest() *EnrollmentManagementStoreRequest`

NewEnrollmentManagementStoreRequest instantiates a new EnrollmentManagementStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentManagementStoreRequestWithDefaults

`func NewEnrollmentManagementStoreRequestWithDefaults() *EnrollmentManagementStoreRequest`

NewEnrollmentManagementStoreRequestWithDefaults instantiates a new EnrollmentManagementStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *EnrollmentManagementStoreRequest) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *EnrollmentManagementStoreRequest) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *EnrollmentManagementStoreRequest) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *EnrollmentManagementStoreRequest) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetAlias

`func (o *EnrollmentManagementStoreRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EnrollmentManagementStoreRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EnrollmentManagementStoreRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *EnrollmentManagementStoreRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *EnrollmentManagementStoreRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *EnrollmentManagementStoreRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetOverwrite

`func (o *EnrollmentManagementStoreRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *EnrollmentManagementStoreRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *EnrollmentManagementStoreRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *EnrollmentManagementStoreRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetProperties

`func (o *EnrollmentManagementStoreRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EnrollmentManagementStoreRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EnrollmentManagementStoreRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EnrollmentManagementStoreRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EnrollmentManagementStoreRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EnrollmentManagementStoreRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


