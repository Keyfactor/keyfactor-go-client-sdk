# ModelsMetadataFieldTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DataType** | Pointer to **int32** |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Enrollment** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**AllowAPI** | Pointer to **bool** |  | [optional] 
**ExplicitUpdate** | Pointer to **bool** |  | [optional] 
**DisplayOrder** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewModelsMetadataFieldTypeModel

`func NewModelsMetadataFieldTypeModel() *ModelsMetadataFieldTypeModel`

NewModelsMetadataFieldTypeModel instantiates a new ModelsMetadataFieldTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMetadataFieldTypeModelWithDefaults

`func NewModelsMetadataFieldTypeModelWithDefaults() *ModelsMetadataFieldTypeModel`

NewModelsMetadataFieldTypeModelWithDefaults instantiates a new ModelsMetadataFieldTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsMetadataFieldTypeModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMetadataFieldTypeModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMetadataFieldTypeModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMetadataFieldTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsMetadataFieldTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsMetadataFieldTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsMetadataFieldTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsMetadataFieldTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsMetadataFieldTypeModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsMetadataFieldTypeModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ModelsMetadataFieldTypeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsMetadataFieldTypeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsMetadataFieldTypeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsMetadataFieldTypeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelsMetadataFieldTypeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelsMetadataFieldTypeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDataType

`func (o *ModelsMetadataFieldTypeModel) GetDataType() int32`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ModelsMetadataFieldTypeModel) GetDataTypeOk() (*int32, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ModelsMetadataFieldTypeModel) SetDataType(v int32)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ModelsMetadataFieldTypeModel) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetHint

`func (o *ModelsMetadataFieldTypeModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ModelsMetadataFieldTypeModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ModelsMetadataFieldTypeModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ModelsMetadataFieldTypeModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *ModelsMetadataFieldTypeModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *ModelsMetadataFieldTypeModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *ModelsMetadataFieldTypeModel) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ModelsMetadataFieldTypeModel) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ModelsMetadataFieldTypeModel) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *ModelsMetadataFieldTypeModel) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *ModelsMetadataFieldTypeModel) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *ModelsMetadataFieldTypeModel) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *ModelsMetadataFieldTypeModel) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ModelsMetadataFieldTypeModel) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ModelsMetadataFieldTypeModel) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *ModelsMetadataFieldTypeModel) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *ModelsMetadataFieldTypeModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsMetadataFieldTypeModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsMetadataFieldTypeModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelsMetadataFieldTypeModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ModelsMetadataFieldTypeModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ModelsMetadataFieldTypeModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *ModelsMetadataFieldTypeModel) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModelsMetadataFieldTypeModel) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModelsMetadataFieldTypeModel) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModelsMetadataFieldTypeModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ModelsMetadataFieldTypeModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ModelsMetadataFieldTypeModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *ModelsMetadataFieldTypeModel) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ModelsMetadataFieldTypeModel) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ModelsMetadataFieldTypeModel) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ModelsMetadataFieldTypeModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ModelsMetadataFieldTypeModel) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ModelsMetadataFieldTypeModel) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *ModelsMetadataFieldTypeModel) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *ModelsMetadataFieldTypeModel) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *ModelsMetadataFieldTypeModel) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *ModelsMetadataFieldTypeModel) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *ModelsMetadataFieldTypeModel) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *ModelsMetadataFieldTypeModel) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *ModelsMetadataFieldTypeModel) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *ModelsMetadataFieldTypeModel) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ModelsMetadataFieldTypeModel) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ModelsMetadataFieldTypeModel) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ModelsMetadataFieldTypeModel) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ModelsMetadataFieldTypeModel) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *ModelsMetadataFieldTypeModel) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *ModelsMetadataFieldTypeModel) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


