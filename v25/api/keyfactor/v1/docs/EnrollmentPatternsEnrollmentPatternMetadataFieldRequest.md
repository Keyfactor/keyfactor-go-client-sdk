# EnrollmentPatternsEnrollmentPatternMetadataFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataId** | Pointer to **int32** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Enrollment** | Pointer to [**CSSCMSCoreEnumsMetadataTypeEnrollment**](CSSCMSCoreEnumsMetadataTypeEnrollment.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 

## Methods

### NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequest

`func NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequest() *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest`

NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequest instantiates a new EnrollmentPatternsEnrollmentPatternMetadataFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequestWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequestWithDefaults() *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest`

NewEnrollmentPatternsEnrollmentPatternMetadataFieldRequestWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternMetadataFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldRequest) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


