# MetadataFieldMetadataFieldUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
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
**DisplayOrder** | **int32** |  | 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**ExemptFromActionedCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetadataFieldMetadataFieldUpdateRequest

`func NewMetadataFieldMetadataFieldUpdateRequest(id int32, name string, description string, dataType CSSCMSCoreEnumsMetadataDataType, displayOrder int32, ) *MetadataFieldMetadataFieldUpdateRequest`

NewMetadataFieldMetadataFieldUpdateRequest instantiates a new MetadataFieldMetadataFieldUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldMetadataFieldUpdateRequestWithDefaults

`func NewMetadataFieldMetadataFieldUpdateRequestWithDefaults() *MetadataFieldMetadataFieldUpdateRequest`

NewMetadataFieldMetadataFieldUpdateRequestWithDefaults instantiates a new MetadataFieldMetadataFieldUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.


### GetHint

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *MetadataFieldMetadataFieldUpdateRequest) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *MetadataFieldMetadataFieldUpdateRequest) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MetadataFieldMetadataFieldUpdateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MetadataFieldMetadataFieldUpdateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *MetadataFieldMetadataFieldUpdateRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetCaseSensitive

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetExemptFromActionedCount() bool`

GetExemptFromActionedCount returns the ExemptFromActionedCount field if non-nil, zero value otherwise.

### GetExemptFromActionedCountOk

`func (o *MetadataFieldMetadataFieldUpdateRequest) GetExemptFromActionedCountOk() (*bool, bool)`

GetExemptFromActionedCountOk returns a tuple with the ExemptFromActionedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldUpdateRequest) SetExemptFromActionedCount(v bool)`

SetExemptFromActionedCount sets ExemptFromActionedCount field to given value.

### HasExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldUpdateRequest) HasExemptFromActionedCount() bool`

HasExemptFromActionedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


