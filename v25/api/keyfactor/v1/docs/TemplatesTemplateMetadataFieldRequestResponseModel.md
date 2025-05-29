# TemplatesTemplateMetadataFieldRequestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**MetadataId** | Pointer to **int32** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Enrollment** | Pointer to [**CSSCMSCoreEnumsMetadataTypeEnrollment**](CSSCMSCoreEnumsMetadataTypeEnrollment.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 

## Methods

### NewTemplatesTemplateMetadataFieldRequestResponseModel

`func NewTemplatesTemplateMetadataFieldRequestResponseModel() *TemplatesTemplateMetadataFieldRequestResponseModel`

NewTemplatesTemplateMetadataFieldRequestResponseModel instantiates a new TemplatesTemplateMetadataFieldRequestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateMetadataFieldRequestResponseModelWithDefaults

`func NewTemplatesTemplateMetadataFieldRequestResponseModelWithDefaults() *TemplatesTemplateMetadataFieldRequestResponseModel`

NewTemplatesTemplateMetadataFieldRequestResponseModelWithDefaults instantiates a new TemplatesTemplateMetadataFieldRequestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMetadataId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetValidation

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCaseSensitive

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplatesTemplateMetadataFieldRequestResponseModel) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


