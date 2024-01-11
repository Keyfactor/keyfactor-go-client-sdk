# KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**DataType** | **int32** |  | 
**Hint** | Pointer to **NullableString** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Enrollment** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**AllowAPI** | Pointer to **bool** |  | [optional] 
**ExplicitUpdate** | Pointer to **bool** |  | [optional] 
**DisplayOrder** | **int32** |  | 

## Methods

### NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(id int32, name string, description string, dataType int32, displayOrder int32, ) *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDataType() int32`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDataTypeOk() (*int32, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDataType(v int32)`

SetDataType sets DataType field to given value.


### GetHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


