# EnrollmentRenewAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **NullableInt32** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEnrollmentRenewAPIRequest

`func NewEnrollmentRenewAPIRequest() *EnrollmentRenewAPIRequest`

NewEnrollmentRenewAPIRequest instantiates a new EnrollmentRenewAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentRenewAPIRequestWithDefaults

`func NewEnrollmentRenewAPIRequestWithDefaults() *EnrollmentRenewAPIRequest`

NewEnrollmentRenewAPIRequestWithDefaults instantiates a new EnrollmentRenewAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *EnrollmentRenewAPIRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *EnrollmentRenewAPIRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *EnrollmentRenewAPIRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *EnrollmentRenewAPIRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *EnrollmentRenewAPIRequest) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *EnrollmentRenewAPIRequest) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetThumbprint

`func (o *EnrollmentRenewAPIRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *EnrollmentRenewAPIRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *EnrollmentRenewAPIRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *EnrollmentRenewAPIRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *EnrollmentRenewAPIRequest) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *EnrollmentRenewAPIRequest) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetCertificateAuthority

`func (o *EnrollmentRenewAPIRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *EnrollmentRenewAPIRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *EnrollmentRenewAPIRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *EnrollmentRenewAPIRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *EnrollmentRenewAPIRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *EnrollmentRenewAPIRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetTemplate

`func (o *EnrollmentRenewAPIRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentRenewAPIRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentRenewAPIRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentRenewAPIRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EnrollmentRenewAPIRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EnrollmentRenewAPIRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetTimestamp

`func (o *EnrollmentRenewAPIRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnrollmentRenewAPIRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnrollmentRenewAPIRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnrollmentRenewAPIRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEnrollmentPatternId

`func (o *EnrollmentRenewAPIRequest) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *EnrollmentRenewAPIRequest) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *EnrollmentRenewAPIRequest) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *EnrollmentRenewAPIRequest) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *EnrollmentRenewAPIRequest) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *EnrollmentRenewAPIRequest) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


