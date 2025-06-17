# MetadataFieldMetadataFieldResponse

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
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**ExemptFromActionedCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetadataFieldMetadataFieldResponse

`func NewMetadataFieldMetadataFieldResponse() *MetadataFieldMetadataFieldResponse`

NewMetadataFieldMetadataFieldResponse instantiates a new MetadataFieldMetadataFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldMetadataFieldResponseWithDefaults

`func NewMetadataFieldMetadataFieldResponseWithDefaults() *MetadataFieldMetadataFieldResponse`

NewMetadataFieldMetadataFieldResponseWithDefaults instantiates a new MetadataFieldMetadataFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataFieldMetadataFieldResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataFieldMetadataFieldResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataFieldMetadataFieldResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataFieldMetadataFieldResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataFieldMetadataFieldResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFieldMetadataFieldResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFieldMetadataFieldResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataFieldMetadataFieldResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataFieldMetadataFieldResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataFieldMetadataFieldResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *MetadataFieldMetadataFieldResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetadataFieldMetadataFieldResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetadataFieldMetadataFieldResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetadataFieldMetadataFieldResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MetadataFieldMetadataFieldResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MetadataFieldMetadataFieldResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDataType

`func (o *MetadataFieldMetadataFieldResponse) GetDataType() CSSCMSCoreEnumsMetadataDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetadataFieldMetadataFieldResponse) GetDataTypeOk() (*CSSCMSCoreEnumsMetadataDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetadataFieldMetadataFieldResponse) SetDataType(v CSSCMSCoreEnumsMetadataDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MetadataFieldMetadataFieldResponse) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetHint

`func (o *MetadataFieldMetadataFieldResponse) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *MetadataFieldMetadataFieldResponse) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *MetadataFieldMetadataFieldResponse) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *MetadataFieldMetadataFieldResponse) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *MetadataFieldMetadataFieldResponse) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *MetadataFieldMetadataFieldResponse) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetValidation

`func (o *MetadataFieldMetadataFieldResponse) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *MetadataFieldMetadataFieldResponse) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *MetadataFieldMetadataFieldResponse) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *MetadataFieldMetadataFieldResponse) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *MetadataFieldMetadataFieldResponse) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *MetadataFieldMetadataFieldResponse) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *MetadataFieldMetadataFieldResponse) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *MetadataFieldMetadataFieldResponse) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *MetadataFieldMetadataFieldResponse) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *MetadataFieldMetadataFieldResponse) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *MetadataFieldMetadataFieldResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetadataFieldMetadataFieldResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetadataFieldMetadataFieldResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetadataFieldMetadataFieldResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MetadataFieldMetadataFieldResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MetadataFieldMetadataFieldResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOptions

`func (o *MetadataFieldMetadataFieldResponse) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MetadataFieldMetadataFieldResponse) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MetadataFieldMetadataFieldResponse) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MetadataFieldMetadataFieldResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MetadataFieldMetadataFieldResponse) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MetadataFieldMetadataFieldResponse) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultValue

`func (o *MetadataFieldMetadataFieldResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetadataFieldMetadataFieldResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetadataFieldMetadataFieldResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetadataFieldMetadataFieldResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *MetadataFieldMetadataFieldResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *MetadataFieldMetadataFieldResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetAllowAPI

`func (o *MetadataFieldMetadataFieldResponse) GetAllowAPI() bool`

GetAllowAPI returns the AllowAPI field if non-nil, zero value otherwise.

### GetAllowAPIOk

`func (o *MetadataFieldMetadataFieldResponse) GetAllowAPIOk() (*bool, bool)`

GetAllowAPIOk returns a tuple with the AllowAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAPI

`func (o *MetadataFieldMetadataFieldResponse) SetAllowAPI(v bool)`

SetAllowAPI sets AllowAPI field to given value.

### HasAllowAPI

`func (o *MetadataFieldMetadataFieldResponse) HasAllowAPI() bool`

HasAllowAPI returns a boolean if a field has been set.

### GetExplicitUpdate

`func (o *MetadataFieldMetadataFieldResponse) GetExplicitUpdate() bool`

GetExplicitUpdate returns the ExplicitUpdate field if non-nil, zero value otherwise.

### GetExplicitUpdateOk

`func (o *MetadataFieldMetadataFieldResponse) GetExplicitUpdateOk() (*bool, bool)`

GetExplicitUpdateOk returns a tuple with the ExplicitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUpdate

`func (o *MetadataFieldMetadataFieldResponse) SetExplicitUpdate(v bool)`

SetExplicitUpdate sets ExplicitUpdate field to given value.

### HasExplicitUpdate

`func (o *MetadataFieldMetadataFieldResponse) HasExplicitUpdate() bool`

HasExplicitUpdate returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *MetadataFieldMetadataFieldResponse) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *MetadataFieldMetadataFieldResponse) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *MetadataFieldMetadataFieldResponse) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *MetadataFieldMetadataFieldResponse) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *MetadataFieldMetadataFieldResponse) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *MetadataFieldMetadataFieldResponse) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetCaseSensitive

`func (o *MetadataFieldMetadataFieldResponse) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *MetadataFieldMetadataFieldResponse) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *MetadataFieldMetadataFieldResponse) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *MetadataFieldMetadataFieldResponse) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldResponse) GetExemptFromActionedCount() bool`

GetExemptFromActionedCount returns the ExemptFromActionedCount field if non-nil, zero value otherwise.

### GetExemptFromActionedCountOk

`func (o *MetadataFieldMetadataFieldResponse) GetExemptFromActionedCountOk() (*bool, bool)`

GetExemptFromActionedCountOk returns a tuple with the ExemptFromActionedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldResponse) SetExemptFromActionedCount(v bool)`

SetExemptFromActionedCount sets ExemptFromActionedCount field to given value.

### HasExemptFromActionedCount

`func (o *MetadataFieldMetadataFieldResponse) HasExemptFromActionedCount() bool`

HasExemptFromActionedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


