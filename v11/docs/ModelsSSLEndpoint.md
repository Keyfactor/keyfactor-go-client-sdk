# ModelsSSLEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**NetworkId** | Pointer to **string** |  | [optional] 
**LastHistoryId** | Pointer to **NullableString** |  | [optional] 
**IpAddressBytes** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**SniName** | Pointer to **NullableString** |  | [optional] 
**EnableMonitor** | Pointer to **bool** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsSSLEndpoint

`func NewModelsSSLEndpoint() *ModelsSSLEndpoint`

NewModelsSSLEndpoint instantiates a new ModelsSSLEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSLEndpointWithDefaults

`func NewModelsSSLEndpointWithDefaults() *ModelsSSLEndpoint`

NewModelsSSLEndpointWithDefaults instantiates a new ModelsSSLEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *ModelsSSLEndpoint) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ModelsSSLEndpoint) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ModelsSSLEndpoint) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ModelsSSLEndpoint) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetNetworkId

`func (o *ModelsSSLEndpoint) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ModelsSSLEndpoint) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ModelsSSLEndpoint) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ModelsSSLEndpoint) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetLastHistoryId

`func (o *ModelsSSLEndpoint) GetLastHistoryId() string`

GetLastHistoryId returns the LastHistoryId field if non-nil, zero value otherwise.

### GetLastHistoryIdOk

`func (o *ModelsSSLEndpoint) GetLastHistoryIdOk() (*string, bool)`

GetLastHistoryIdOk returns a tuple with the LastHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHistoryId

`func (o *ModelsSSLEndpoint) SetLastHistoryId(v string)`

SetLastHistoryId sets LastHistoryId field to given value.

### HasLastHistoryId

`func (o *ModelsSSLEndpoint) HasLastHistoryId() bool`

HasLastHistoryId returns a boolean if a field has been set.

### SetLastHistoryIdNil

`func (o *ModelsSSLEndpoint) SetLastHistoryIdNil(b bool)`

 SetLastHistoryIdNil sets the value for LastHistoryId to be an explicit nil

### UnsetLastHistoryId
`func (o *ModelsSSLEndpoint) UnsetLastHistoryId()`

UnsetLastHistoryId ensures that no value is present for LastHistoryId, not even an explicit nil
### GetIpAddressBytes

`func (o *ModelsSSLEndpoint) GetIpAddressBytes() string`

GetIpAddressBytes returns the IpAddressBytes field if non-nil, zero value otherwise.

### GetIpAddressBytesOk

`func (o *ModelsSSLEndpoint) GetIpAddressBytesOk() (*string, bool)`

GetIpAddressBytesOk returns a tuple with the IpAddressBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressBytes

`func (o *ModelsSSLEndpoint) SetIpAddressBytes(v string)`

SetIpAddressBytes sets IpAddressBytes field to given value.

### HasIpAddressBytes

`func (o *ModelsSSLEndpoint) HasIpAddressBytes() bool`

HasIpAddressBytes returns a boolean if a field has been set.

### SetIpAddressBytesNil

`func (o *ModelsSSLEndpoint) SetIpAddressBytesNil(b bool)`

 SetIpAddressBytesNil sets the value for IpAddressBytes to be an explicit nil

### UnsetIpAddressBytes
`func (o *ModelsSSLEndpoint) UnsetIpAddressBytes()`

UnsetIpAddressBytes ensures that no value is present for IpAddressBytes, not even an explicit nil
### GetPort

`func (o *ModelsSSLEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelsSSLEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelsSSLEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelsSSLEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSniName

`func (o *ModelsSSLEndpoint) GetSniName() string`

GetSniName returns the SniName field if non-nil, zero value otherwise.

### GetSniNameOk

`func (o *ModelsSSLEndpoint) GetSniNameOk() (*string, bool)`

GetSniNameOk returns a tuple with the SniName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniName

`func (o *ModelsSSLEndpoint) SetSniName(v string)`

SetSniName sets SniName field to given value.

### HasSniName

`func (o *ModelsSSLEndpoint) HasSniName() bool`

HasSniName returns a boolean if a field has been set.

### SetSniNameNil

`func (o *ModelsSSLEndpoint) SetSniNameNil(b bool)`

 SetSniNameNil sets the value for SniName to be an explicit nil

### UnsetSniName
`func (o *ModelsSSLEndpoint) UnsetSniName()`

UnsetSniName ensures that no value is present for SniName, not even an explicit nil
### GetEnableMonitor

`func (o *ModelsSSLEndpoint) GetEnableMonitor() bool`

GetEnableMonitor returns the EnableMonitor field if non-nil, zero value otherwise.

### GetEnableMonitorOk

`func (o *ModelsSSLEndpoint) GetEnableMonitorOk() (*bool, bool)`

GetEnableMonitorOk returns a tuple with the EnableMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitor

`func (o *ModelsSSLEndpoint) SetEnableMonitor(v bool)`

SetEnableMonitor sets EnableMonitor field to given value.

### HasEnableMonitor

`func (o *ModelsSSLEndpoint) HasEnableMonitor() bool`

HasEnableMonitor returns a boolean if a field has been set.

### GetReviewed

`func (o *ModelsSSLEndpoint) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *ModelsSSLEndpoint) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *ModelsSSLEndpoint) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *ModelsSSLEndpoint) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


