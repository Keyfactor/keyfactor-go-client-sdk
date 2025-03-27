# MetadataFieldMetadataFieldCreateRequest

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
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**ExemptFromActionedCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetadataFieldMetadataFieldCreateRequest

`func NewMetadataFieldMetadataFieldCreateRequest(name string, description string, dataType CSSCMSCoreEnumsMetadataDataType, ) *MetadataFieldMetadataFieldCreateRequest`

NewMetadataFieldMetadataFieldCreateRequest instantiates a new MetadataFieldMetadataFieldCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldMetadataFieldCreateRequestWithDefaults

`func NewMetadataFieldMetadataFieldCreateRequestWithDefaults() *MetadataFieldMetadataFieldCreateRequest`

NewMetadataFieldMetadataFieldCreateRequestWithDefaults instantiates a new MetadataFieldMetadataFieldCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetadataFieldMetadataFieldCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFieldMetadataFieldCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.


### GetHint

`func (o *MetadataFieldMetadataFieldCreateRequest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *MetadataFieldMetadataFieldCreateRequest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *MetadataFieldMetadataFieldCreateRequest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *MetadataFieldMetadataFieldCreateRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *MetadataFieldMetadataFieldCreateRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *MetadataFieldMetadataFieldCreateRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *MetadataFieldMetadataFieldCreateRequest) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *MetadataFieldMetadataFieldCreateRequest) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *MetadataFieldMetadataFieldCreateRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *MetadataFieldMetadataFieldCreateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetadataFieldMetadataFieldCreateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetadataFieldMetadataFieldCreateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *MetadataFieldMetadataFieldCreateRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MetadataFieldMetadataFieldCreateRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MetadataFieldMetadataFieldCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetadataFieldMetadataFieldCreateRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *MetadataFieldMetadataFieldCreateRequest) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *MetadataFieldMetadataFieldCreateRequest) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *MetadataFieldMetadataFieldCreateRequest) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *MetadataFieldMetadataFieldCreateRequest) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *MetadataFieldMetadataFieldCreateRequest) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *MetadataFieldMetadataFieldCreateRequest) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *MetadataFieldMetadataFieldCreateRequest) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *MetadataFieldMetadataFieldCreateRequest) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *MetadataFieldMetadataFieldCreateRequest) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetCaseSensitive

`func (o *MetadataFieldMetadataFieldCreateRequest) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *MetadataFieldMetadataFieldCreateRequest) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *MetadataFieldMetadataFieldCreateRequest) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldCreateRequest) GetExemptFromActionedCount() bool`

GetExemptFromActionedCount returns the ExemptFromActionedCount field if non-nil, zero value otherwise.

### GetExemptFromActionedCountOk

`func (o *MetadataFieldMetadataFieldCreateRequest) GetExemptFromActionedCountOk() (*bool, bool)`

GetExemptFromActionedCountOk returns a tuple with the ExemptFromActionedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldCreateRequest) SetExemptFromActionedCount(v bool)`

SetExemptFromActionedCount sets ExemptFromActionedCount field to given value.

### HasExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldCreateRequest) HasExemptFromActionedCount() bool`

HasExemptFromActionedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


