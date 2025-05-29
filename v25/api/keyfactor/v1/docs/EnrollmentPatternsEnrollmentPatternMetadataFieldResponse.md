# EnrollmentPatternsEnrollmentPatternMetadataFieldResponse

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

### NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponse

`func NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponse() *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse`

NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponse instantiates a new EnrollmentPatternsEnrollmentPatternMetadataFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponseWithDefaults

`func NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponseWithDefaults() *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse`

NewEnrollmentPatternsEnrollmentPatternMetadataFieldResponseWithDefaults instantiates a new EnrollmentPatternsEnrollmentPatternMetadataFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetEnrollment() CSSCMSCoreEnumsMetadataTypeEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetEnrollmentOk() (*CSSCMSCoreEnumsMetadataTypeEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetEnrollment(v CSSCMSCoreEnumsMetadataTypeEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *EnrollmentPatternsEnrollmentPatternMetadataFieldResponse) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


