# CSSCMSDataModelModelsMetadataType

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
**DisplayOrder** | Pointer to **NullableInt32** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**ExemptFromActionedCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsMetadataType

`func NewCSSCMSDataModelModelsMetadataType() *CSSCMSDataModelModelsMetadataType`

NewCSSCMSDataModelModelsMetadataType instantiates a new CSSCMSDataModelModelsMetadataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsMetadataTypeWithDefaults

`func NewCSSCMSDataModelModelsMetadataTypeWithDefaults() *CSSCMSDataModelModelsMetadataType`

NewCSSCMSDataModelModelsMetadataTypeWithDefaults instantiates a new CSSCMSDataModelModelsMetadataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsMetadataType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsMetadataType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsMetadataType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsMetadataType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsMetadataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsMetadataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsMetadataType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsMetadataType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsMetadataType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsMetadataType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CSSCMSDataModelModelsMetadataType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsMetadataType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsMetadataType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsMetadataType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsMetadataType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsMetadataType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDataType

`func (o *CSSCMSDataModelModelsMetadataType) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CSSCMSDataModelModelsMetadataType) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CSSCMSDataModelModelsMetadataType) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *CSSCMSDataModelModelsMetadataType) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetHint

`func (o *CSSCMSDataModelModelsMetadataType) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CSSCMSDataModelModelsMetadataType) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CSSCMSDataModelModelsMetadataType) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CSSCMSDataModelModelsMetadataType) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *CSSCMSDataModelModelsMetadataType) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *CSSCMSDataModelModelsMetadataType) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *CSSCMSDataModelModelsMetadataType) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *CSSCMSDataModelModelsMetadataType) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *CSSCMSDataModelModelsMetadataType) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *CSSCMSDataModelModelsMetadataType) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *CSSCMSDataModelModelsMetadataType) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *CSSCMSDataModelModelsMetadataType) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *CSSCMSDataModelModelsMetadataType) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *CSSCMSDataModelModelsMetadataType) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *CSSCMSDataModelModelsMetadataType) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *CSSCMSDataModelModelsMetadataType) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *CSSCMSDataModelModelsMetadataType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CSSCMSDataModelModelsMetadataType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CSSCMSDataModelModelsMetadataType) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CSSCMSDataModelModelsMetadataType) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CSSCMSDataModelModelsMetadataType) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CSSCMSDataModelModelsMetadataType) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *CSSCMSDataModelModelsMetadataType) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CSSCMSDataModelModelsMetadataType) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CSSCMSDataModelModelsMetadataType) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CSSCMSDataModelModelsMetadataType) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CSSCMSDataModelModelsMetadataType) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CSSCMSDataModelModelsMetadataType) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *CSSCMSDataModelModelsMetadataType) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsMetadataType) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsMetadataType) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsMetadataType) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsMetadataType) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsMetadataType) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataType) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *CSSCMSDataModelModelsMetadataType) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataType) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *CSSCMSDataModelModelsMetadataType) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *CSSCMSDataModelModelsMetadataType) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *CSSCMSDataModelModelsMetadataType) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetCaseSensitive

`func (o *CSSCMSDataModelModelsMetadataType) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *CSSCMSDataModelModelsMetadataType) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *CSSCMSDataModelModelsMetadataType) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *CSSCMSDataModelModelsMetadataType) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetExemptFromActionedCount

`func (o *CSSCMSDataModelModelsMetadataType) GetExemptFromActionedCount() bool`

GetExemptFromActionedCount returns the ExemptFromActionedCount field if non-nil, zero value otherwise.

### GetExemptFromActionedCountOk

`func (o *CSSCMSDataModelModelsMetadataType) GetExemptFromActionedCountOk() (*bool, bool)`

GetExemptFromActionedCountOk returns a tuple with the ExemptFromActionedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptFromActionedCount

`func (o *CSSCMSDataModelModelsMetadataType) SetExemptFromActionedCount(v bool)`

SetExemptFromActionedCount sets ExemptFromActionedCount field to given value.

### HasExemptFromActionedCount

`func (o *CSSCMSDataModelModelsMetadataType) HasExemptFromActionedCount() bool`

HasExemptFromActionedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


