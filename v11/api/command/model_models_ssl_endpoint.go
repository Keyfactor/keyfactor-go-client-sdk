/*

Copyright 2023 Keyfactor
Licensed under the Apache License, Version 2.0 (the License); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an AS IS BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

Keyfactor API Reference and Utility

<p>This page provides a utility through which the Keyfactor API endpoints can be called and results returned.                                                           It is intended to be used primarily for validation, testing and workflow development.                                                           It also serves secondarily as documentation for the API.</p>                                                          <p>If you would like to view documentation containing details on the Keyfactor API and endpoints,                                                           please refer to the Web API section of the Keyfactor Command documentation.</p>

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"encoding/json"
)

// checks if the ModelsSSLEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSLEndpoint{}

// ModelsSSLEndpoint struct for ModelsSSLEndpoint
type ModelsSSLEndpoint struct {
	EndpointId     *string        `json:"endpointId,omitempty"`
	NetworkId      *string        `json:"networkId,omitempty"`
	LastHistoryId  NullableString `json:"lastHistoryId,omitempty"`
	IpAddressBytes NullableString `json:"ipAddressBytes,omitempty"`
	Port           *int32         `json:"port,omitempty"`
	SniName        NullableString `json:"sniName,omitempty"`
	EnableMonitor  *bool          `json:"enableMonitor,omitempty"`
	Reviewed       *bool          `json:"reviewed,omitempty"`
}

// NewModelsSSLEndpoint instantiates a new ModelsSSLEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSLEndpoint() *ModelsSSLEndpoint {
	this := ModelsSSLEndpoint{}
	return &this
}

// NewModelsSSLEndpointWithDefaults instantiates a new ModelsSSLEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSLEndpointWithDefaults() *ModelsSSLEndpoint {
	this := ModelsSSLEndpoint{}
	return &this
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *ModelsSSLEndpoint) GetEndpointId() string {
	if o == nil || isNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpoint) GetEndpointIdOk() (*string, bool) {
	if o == nil || isNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasEndpointId() bool {
	if o != nil && !isNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *ModelsSSLEndpoint) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *ModelsSSLEndpoint) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpoint) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *ModelsSSLEndpoint) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetLastHistoryId returns the LastHistoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSLEndpoint) GetLastHistoryId() string {
	if o == nil || isNil(o.LastHistoryId.Get()) {
		var ret string
		return ret
	}
	return *o.LastHistoryId.Get()
}

// GetLastHistoryIdOk returns a tuple with the LastHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSLEndpoint) GetLastHistoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastHistoryId.Get(), o.LastHistoryId.IsSet()
}

// HasLastHistoryId returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasLastHistoryId() bool {
	if o != nil && o.LastHistoryId.IsSet() {
		return true
	}

	return false
}

// SetLastHistoryId gets a reference to the given NullableString and assigns it to the LastHistoryId field.
func (o *ModelsSSLEndpoint) SetLastHistoryId(v string) {
	o.LastHistoryId.Set(&v)
}

// SetLastHistoryIdNil sets the value for LastHistoryId to be an explicit nil
func (o *ModelsSSLEndpoint) SetLastHistoryIdNil() {
	o.LastHistoryId.Set(nil)
}

// UnsetLastHistoryId ensures that no value is present for LastHistoryId, not even an explicit nil
func (o *ModelsSSLEndpoint) UnsetLastHistoryId() {
	o.LastHistoryId.Unset()
}

// GetIpAddressBytes returns the IpAddressBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSLEndpoint) GetIpAddressBytes() string {
	if o == nil || isNil(o.IpAddressBytes.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddressBytes.Get()
}

// GetIpAddressBytesOk returns a tuple with the IpAddressBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSLEndpoint) GetIpAddressBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddressBytes.Get(), o.IpAddressBytes.IsSet()
}

// HasIpAddressBytes returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasIpAddressBytes() bool {
	if o != nil && o.IpAddressBytes.IsSet() {
		return true
	}

	return false
}

// SetIpAddressBytes gets a reference to the given NullableString and assigns it to the IpAddressBytes field.
func (o *ModelsSSLEndpoint) SetIpAddressBytes(v string) {
	o.IpAddressBytes.Set(&v)
}

// SetIpAddressBytesNil sets the value for IpAddressBytes to be an explicit nil
func (o *ModelsSSLEndpoint) SetIpAddressBytesNil() {
	o.IpAddressBytes.Set(nil)
}

// UnsetIpAddressBytes ensures that no value is present for IpAddressBytes, not even an explicit nil
func (o *ModelsSSLEndpoint) UnsetIpAddressBytes() {
	o.IpAddressBytes.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelsSSLEndpoint) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpoint) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ModelsSSLEndpoint) SetPort(v int32) {
	o.Port = &v
}

// GetSniName returns the SniName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSLEndpoint) GetSniName() string {
	if o == nil || isNil(o.SniName.Get()) {
		var ret string
		return ret
	}
	return *o.SniName.Get()
}

// GetSniNameOk returns a tuple with the SniName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSLEndpoint) GetSniNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SniName.Get(), o.SniName.IsSet()
}

// HasSniName returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasSniName() bool {
	if o != nil && o.SniName.IsSet() {
		return true
	}

	return false
}

// SetSniName gets a reference to the given NullableString and assigns it to the SniName field.
func (o *ModelsSSLEndpoint) SetSniName(v string) {
	o.SniName.Set(&v)
}

// SetSniNameNil sets the value for SniName to be an explicit nil
func (o *ModelsSSLEndpoint) SetSniNameNil() {
	o.SniName.Set(nil)
}

// UnsetSniName ensures that no value is present for SniName, not even an explicit nil
func (o *ModelsSSLEndpoint) UnsetSniName() {
	o.SniName.Unset()
}

// GetEnableMonitor returns the EnableMonitor field value if set, zero value otherwise.
func (o *ModelsSSLEndpoint) GetEnableMonitor() bool {
	if o == nil || isNil(o.EnableMonitor) {
		var ret bool
		return ret
	}
	return *o.EnableMonitor
}

// GetEnableMonitorOk returns a tuple with the EnableMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpoint) GetEnableMonitorOk() (*bool, bool) {
	if o == nil || isNil(o.EnableMonitor) {
		return nil, false
	}
	return o.EnableMonitor, true
}

// HasEnableMonitor returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasEnableMonitor() bool {
	if o != nil && !isNil(o.EnableMonitor) {
		return true
	}

	return false
}

// SetEnableMonitor gets a reference to the given bool and assigns it to the EnableMonitor field.
func (o *ModelsSSLEndpoint) SetEnableMonitor(v bool) {
	o.EnableMonitor = &v
}

// GetReviewed returns the Reviewed field value if set, zero value otherwise.
func (o *ModelsSSLEndpoint) GetReviewed() bool {
	if o == nil || isNil(o.Reviewed) {
		var ret bool
		return ret
	}
	return *o.Reviewed
}

// GetReviewedOk returns a tuple with the Reviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpoint) GetReviewedOk() (*bool, bool) {
	if o == nil || isNil(o.Reviewed) {
		return nil, false
	}
	return o.Reviewed, true
}

// HasReviewed returns a boolean if a field has been set.
func (o *ModelsSSLEndpoint) HasReviewed() bool {
	if o != nil && !isNil(o.Reviewed) {
		return true
	}

	return false
}

// SetReviewed gets a reference to the given bool and assigns it to the Reviewed field.
func (o *ModelsSSLEndpoint) SetReviewed(v bool) {
	o.Reviewed = &v
}

func (o ModelsSSLEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSLEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EndpointId) {
		toSerialize["endpointId"] = o.EndpointId
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if o.LastHistoryId.IsSet() {
		toSerialize["lastHistoryId"] = o.LastHistoryId.Get()
	}
	if o.IpAddressBytes.IsSet() {
		toSerialize["ipAddressBytes"] = o.IpAddressBytes.Get()
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if o.SniName.IsSet() {
		toSerialize["sniName"] = o.SniName.Get()
	}
	if !isNil(o.EnableMonitor) {
		toSerialize["enableMonitor"] = o.EnableMonitor
	}
	if !isNil(o.Reviewed) {
		toSerialize["reviewed"] = o.Reviewed
	}
	return toSerialize, nil
}

type NullableModelsSSLEndpoint struct {
	value *ModelsSSLEndpoint
	isSet bool
}

func (v NullableModelsSSLEndpoint) Get() *ModelsSSLEndpoint {
	return v.value
}

func (v *NullableModelsSSLEndpoint) Set(val *ModelsSSLEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSLEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSLEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSLEndpoint(val *ModelsSSLEndpoint) *NullableModelsSSLEndpoint {
	return &NullableModelsSSLEndpoint{value: val, isSet: true}
}

func (v NullableModelsSSLEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSLEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}