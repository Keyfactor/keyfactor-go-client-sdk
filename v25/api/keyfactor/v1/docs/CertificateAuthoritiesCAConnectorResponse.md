# CertificateAuthoritiesCAConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**LastSeen** | Pointer to **NullableTime** |  | [optional] 
**ClaimType** | Pointer to [**CSSCMSCoreEnumsClaimType**](CSSCMSCoreEnumsClaimType.md) |  | [optional] 
**ClaimValue** | Pointer to **NullableString** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**CAType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 

## Methods

### NewCertificateAuthoritiesCAConnectorResponse

`func NewCertificateAuthoritiesCAConnectorResponse() *CertificateAuthoritiesCAConnectorResponse`

NewCertificateAuthoritiesCAConnectorResponse instantiates a new CertificateAuthoritiesCAConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthoritiesCAConnectorResponseWithDefaults

`func NewCertificateAuthoritiesCAConnectorResponseWithDefaults() *CertificateAuthoritiesCAConnectorResponse`

NewCertificateAuthoritiesCAConnectorResponseWithDefaults instantiates a new CertificateAuthoritiesCAConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateAuthoritiesCAConnectorResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthoritiesCAConnectorResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateAuthoritiesCAConnectorResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthoritiesCAConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthoritiesCAConnectorResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateAuthoritiesCAConnectorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateAuthoritiesCAConnectorResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateAuthoritiesCAConnectorResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *CertificateAuthoritiesCAConnectorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateAuthoritiesCAConnectorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CertificateAuthoritiesCAConnectorResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoolName

`func (o *CertificateAuthoritiesCAConnectorResponse) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *CertificateAuthoritiesCAConnectorResponse) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *CertificateAuthoritiesCAConnectorResponse) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *CertificateAuthoritiesCAConnectorResponse) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *CertificateAuthoritiesCAConnectorResponse) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetLastSeen

`func (o *CertificateAuthoritiesCAConnectorResponse) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *CertificateAuthoritiesCAConnectorResponse) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *CertificateAuthoritiesCAConnectorResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeenNil

`func (o *CertificateAuthoritiesCAConnectorResponse) SetLastSeenNil(b bool)`

 SetLastSeenNil sets the value for LastSeen to be an explicit nil

### UnsetLastSeen
`func (o *CertificateAuthoritiesCAConnectorResponse) UnsetLastSeen()`

UnsetLastSeen ensures that no value is present for LastSeen, not even an explicit nil
### GetClaimType

`func (o *CertificateAuthoritiesCAConnectorResponse) GetClaimType() CSSCMSCoreEnumsClaimType`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetClaimTypeOk() (*CSSCMSCoreEnumsClaimType, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *CertificateAuthoritiesCAConnectorResponse) SetClaimType(v CSSCMSCoreEnumsClaimType)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *CertificateAuthoritiesCAConnectorResponse) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### GetClaimValue

`func (o *CertificateAuthoritiesCAConnectorResponse) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *CertificateAuthoritiesCAConnectorResponse) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *CertificateAuthoritiesCAConnectorResponse) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *CertificateAuthoritiesCAConnectorResponse) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *CertificateAuthoritiesCAConnectorResponse) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil
### GetProviderId

`func (o *CertificateAuthoritiesCAConnectorResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CertificateAuthoritiesCAConnectorResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CertificateAuthoritiesCAConnectorResponse) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderName

`func (o *CertificateAuthoritiesCAConnectorResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CertificateAuthoritiesCAConnectorResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *CertificateAuthoritiesCAConnectorResponse) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *CertificateAuthoritiesCAConnectorResponse) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *CertificateAuthoritiesCAConnectorResponse) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetConnected

`func (o *CertificateAuthoritiesCAConnectorResponse) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *CertificateAuthoritiesCAConnectorResponse) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *CertificateAuthoritiesCAConnectorResponse) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCAType

`func (o *CertificateAuthoritiesCAConnectorResponse) GetCAType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *CertificateAuthoritiesCAConnectorResponse) GetCATypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *CertificateAuthoritiesCAConnectorResponse) SetCAType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *CertificateAuthoritiesCAConnectorResponse) HasCAType() bool`

HasCAType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


