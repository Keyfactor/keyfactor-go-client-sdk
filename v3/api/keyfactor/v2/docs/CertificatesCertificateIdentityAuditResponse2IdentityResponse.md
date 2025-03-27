# CertificatesCertificateIdentityAuditResponse2IdentityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Id of the identity. | [optional] 
**Description** | Pointer to **NullableString** | The description of the identity. | [optional] 
**ClaimType** | Pointer to **NullableString** | The claim type of the identity. | [optional] 
**ClaimValue** | Pointer to **NullableString** | The claim value of the identity. | [optional] 
**Provider** | Pointer to [**CertificatesCertificateIdentityAuditResponse2IdentityProvider**](CertificatesCertificateIdentityAuditResponse2IdentityProvider.md) |  | [optional] 

## Methods

### NewCertificatesCertificateIdentityAuditResponse2IdentityResponse

`func NewCertificatesCertificateIdentityAuditResponse2IdentityResponse() *CertificatesCertificateIdentityAuditResponse2IdentityResponse`

NewCertificatesCertificateIdentityAuditResponse2IdentityResponse instantiates a new CertificatesCertificateIdentityAuditResponse2IdentityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateIdentityAuditResponse2IdentityResponseWithDefaults

`func NewCertificatesCertificateIdentityAuditResponse2IdentityResponseWithDefaults() *CertificatesCertificateIdentityAuditResponse2IdentityResponse`

NewCertificatesCertificateIdentityAuditResponse2IdentityResponseWithDefaults instantiates a new CertificatesCertificateIdentityAuditResponse2IdentityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetClaimType

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetClaimType() string`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetClaimTypeOk() (*string, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetClaimType(v string)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### SetClaimTypeNil

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetClaimTypeNil(b bool)`

 SetClaimTypeNil sets the value for ClaimType to be an explicit nil

### UnsetClaimType
`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) UnsetClaimType()`

UnsetClaimType ensures that no value is present for ClaimType, not even an explicit nil
### GetClaimValue

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil
### GetProvider

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetProvider() CertificatesCertificateIdentityAuditResponse2IdentityProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) GetProviderOk() (*CertificatesCertificateIdentityAuditResponse2IdentityProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) SetProvider(v CertificatesCertificateIdentityAuditResponse2IdentityProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CertificatesCertificateIdentityAuditResponse2IdentityResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


