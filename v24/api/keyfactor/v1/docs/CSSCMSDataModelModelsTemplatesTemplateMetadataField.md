# CSSCMSDataModelModelsTemplatesTemplateMetadataField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**MetadataId** | Pointer to **int32** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**Enrollment** | Pointer to [**CSSCMSCoreEnumsMetadataTypeEnrollment**](CSSCMSCoreEnumsMetadataTypeEnrollment.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsTemplatesTemplateMetadataField

`func NewCSSCMSDataModelModelsTemplatesTemplateMetadataField() *CSSCMSDataModelModelsTemplatesTemplateMetadataField`

NewCSSCMSDataModelModelsTemplatesTemplateMetadataField instantiates a new CSSCMSDataModelModelsTemplatesTemplateMetadataField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsTemplatesTemplateMetadataFieldWithDefaults

`func NewCSSCMSDataModelModelsTemplatesTemplateMetadataFieldWithDefaults() *CSSCMSDataModelModelsTemplatesTemplateMetadataField`

NewCSSCMSDataModelModelsTemplatesTemplateMetadataFieldWithDefaults instantiates a new CSSCMSDataModelModelsTemplatesTemplateMetadataField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetTemplateId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetMetadataId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetValidation

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetMessage

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCaseSensitive

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetEnrollment

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *CSSCMSDataModelModelsTemplatesTemplateMetadataField) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


