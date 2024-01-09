# KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**DataType** | [**CSSCMSCoreEnumsMetadataDataType**](CSSCMSCoreEnumsMetadataDataType.md) |  | 
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

### NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest

`func NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest(name string, description string, dataType CSSCMSCoreEnumsMetadataDataType, ) *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest`

NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest instantiates a new KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest`

NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.


### GetHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


