# CSSCMSDataModelModelsMetadataFieldTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DataType** | Pointer to [**CSSCMSCoreEnumsMetadataDataType**](CSSCMSCoreEnumsMetadataDataType.md) |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Enrollment** | Pointer to [**CSSCMSCoreEnumsMetadataTypeEnrollment**](CSSCMSCoreEnumsMetadataTypeEnrollment.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**AllowAPI** | Pointer to **bool** |  | [optional] 
**ExplicitUpdate** | Pointer to **bool** |  | [optional] 
**DisplayOrder** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsMetadataFieldTypeModel

`func NewCSSCMSDataModelModelsMetadataFieldTypeModel() *CSSCMSDataModelModelsMetadataFieldTypeModel`

NewCSSCMSDataModelModelsMetadataFieldTypeModel instantiates a new CSSCMSDataModelModelsMetadataFieldTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsMetadataFieldTypeModelWithDefaults

`func NewCSSCMSDataModelModelsMetadataFieldTypeModelWithDefaults() *CSSCMSDataModelModelsMetadataFieldTypeModel`

NewCSSCMSDataModelModelsMetadataFieldTypeModelWithDefaults instantiates a new CSSCMSDataModelModelsMetadataFieldTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDataType

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetHint

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *CSSCMSDataModelModelsMetadataFieldTypeModel) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


