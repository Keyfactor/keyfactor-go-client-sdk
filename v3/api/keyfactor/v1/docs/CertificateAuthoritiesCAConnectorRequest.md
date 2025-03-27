# CertificateAuthoritiesCAConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**ClaimType** | Pointer to [**CSSCMSCoreEnumsClaimType**](CSSCMSCoreEnumsClaimType.md) |  | [optional] 
**ClaimValue** | Pointer to **NullableString** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**CAType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 

## Methods

### NewCertificateAuthoritiesCAConnectorRequest

`func NewCertificateAuthoritiesCAConnectorRequest() *CertificateAuthoritiesCAConnectorRequest`

NewCertificateAuthoritiesCAConnectorRequest instantiates a new CertificateAuthoritiesCAConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthoritiesCAConnectorRequestWithDefaults

`func NewCertificateAuthoritiesCAConnectorRequestWithDefaults() *CertificateAuthoritiesCAConnectorRequest`

NewCertificateAuthoritiesCAConnectorRequestWithDefaults instantiates a new CertificateAuthoritiesCAConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateAuthoritiesCAConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthoritiesCAConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateAuthoritiesCAConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateAuthoritiesCAConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateAuthoritiesCAConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *CertificateAuthoritiesCAConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateAuthoritiesCAConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CertificateAuthoritiesCAConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoolName

`func (o *CertificateAuthoritiesCAConnectorRequest) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *CertificateAuthoritiesCAConnectorRequest) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *CertificateAuthoritiesCAConnectorRequest) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *CertificateAuthoritiesCAConnectorRequest) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *CertificateAuthoritiesCAConnectorRequest) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetClaimType

`func (o *CertificateAuthoritiesCAConnectorRequest) GetClaimType() CSSCMSCoreEnumsClaimType`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetClaimTypeOk() (*CSSCMSCoreEnumsClaimType, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *CertificateAuthoritiesCAConnectorRequest) SetClaimType(v CSSCMSCoreEnumsClaimType)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *CertificateAuthoritiesCAConnectorRequest) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### GetClaimValue

`func (o *CertificateAuthoritiesCAConnectorRequest) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *CertificateAuthoritiesCAConnectorRequest) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *CertificateAuthoritiesCAConnectorRequest) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *CertificateAuthoritiesCAConnectorRequest) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *CertificateAuthoritiesCAConnectorRequest) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil
### GetProviderId

`func (o *CertificateAuthoritiesCAConnectorRequest) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CertificateAuthoritiesCAConnectorRequest) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CertificateAuthoritiesCAConnectorRequest) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetCAType

`func (o *CertificateAuthoritiesCAConnectorRequest) GetCAType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *CertificateAuthoritiesCAConnectorRequest) GetCATypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *CertificateAuthoritiesCAConnectorRequest) SetCAType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *CertificateAuthoritiesCAConnectorRequest) HasCAType() bool`

HasCAType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


