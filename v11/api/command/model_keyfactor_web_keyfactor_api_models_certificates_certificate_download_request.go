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

// checks if the KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest{}

// KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest struct for KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest
type KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest struct {
	CertID       NullableInt32  `json:"certID,omitempty"`
	SerialNumber NullableString `json:"serialNumber,omitempty"`
	IssuerDN     NullableString `json:"issuerDN,omitempty"`
	Thumbprint   NullableString `json:"thumbprint,omitempty"`
	IncludeChain *bool          `json:"includeChain,omitempty"`
	ChainOrder   NullableString `json:"chainOrder,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest {
	this := KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest {
	this := KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest{}
	return &this
}

// GetCertID returns the CertID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetCertID() int32 {
	if o == nil || isNil(o.CertID.Get()) {
		var ret int32
		return ret
	}
	return *o.CertID.Get()
}

// GetCertIDOk returns a tuple with the CertID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetCertIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertID.Get(), o.CertID.IsSet()
}

// HasCertID returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasCertID() bool {
	if o != nil && o.CertID.IsSet() {
		return true
	}

	return false
}

// SetCertID gets a reference to the given NullableInt32 and assigns it to the CertID field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetCertID(v int32) {
	o.CertID.Set(&v)
}

// SetCertIDNil sets the value for CertID to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetCertIDNil() {
	o.CertID.Set(nil)
}

// UnsetCertID ensures that no value is present for CertID, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnsetCertID() {
	o.CertID.Unset()
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SerialNumber.Get()
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerialNumber.Get(), o.SerialNumber.IsSet()
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber.IsSet() {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given NullableString and assigns it to the SerialNumber field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetSerialNumber(v string) {
	o.SerialNumber.Set(&v)
}

// SetSerialNumberNil sets the value for SerialNumber to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetSerialNumberNil() {
	o.SerialNumber.Set(nil)
}

// UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnsetSerialNumber() {
	o.SerialNumber.Unset()
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}

// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbprint.Get()
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbprint.Get(), o.Thumbprint.IsSet()
}

// HasThumbprint returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasThumbprint() bool {
	if o != nil && o.Thumbprint.IsSet() {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given NullableString and assigns it to the Thumbprint field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetThumbprint(v string) {
	o.Thumbprint.Set(&v)
}

// SetThumbprintNil sets the value for Thumbprint to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetThumbprintNil() {
	o.Thumbprint.Set(nil)
}

// UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnsetThumbprint() {
	o.Thumbprint.Unset()
}

// GetIncludeChain returns the IncludeChain field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetIncludeChain() bool {
	if o == nil || isNil(o.IncludeChain) {
		var ret bool
		return ret
	}
	return *o.IncludeChain
}

// GetIncludeChainOk returns a tuple with the IncludeChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetIncludeChainOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeChain) {
		return nil, false
	}
	return o.IncludeChain, true
}

// HasIncludeChain returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasIncludeChain() bool {
	if o != nil && !isNil(o.IncludeChain) {
		return true
	}

	return false
}

// SetIncludeChain gets a reference to the given bool and assigns it to the IncludeChain field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetIncludeChain(v bool) {
	o.IncludeChain = &v
}

// GetChainOrder returns the ChainOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetChainOrder() string {
	if o == nil || isNil(o.ChainOrder.Get()) {
		var ret string
		return ret
	}
	return *o.ChainOrder.Get()
}

// GetChainOrderOk returns a tuple with the ChainOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) GetChainOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainOrder.Get(), o.ChainOrder.IsSet()
}

// HasChainOrder returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) HasChainOrder() bool {
	if o != nil && o.ChainOrder.IsSet() {
		return true
	}

	return false
}

// SetChainOrder gets a reference to the given NullableString and assigns it to the ChainOrder field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetChainOrder(v string) {
	o.ChainOrder.Set(&v)
}

// SetChainOrderNil sets the value for ChainOrder to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) SetChainOrderNil() {
	o.ChainOrder.Set(nil)
}

// UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnsetChainOrder() {
	o.ChainOrder.Unset()
}

func (o KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CertID.IsSet() {
		toSerialize["certID"] = o.CertID.Get()
	}
	if o.SerialNumber.IsSet() {
		toSerialize["serialNumber"] = o.SerialNumber.Get()
	}
	if o.IssuerDN.IsSet() {
		toSerialize["issuerDN"] = o.IssuerDN.Get()
	}
	if o.Thumbprint.IsSet() {
		toSerialize["thumbprint"] = o.Thumbprint.Get()
	}
	if !isNil(o.IncludeChain) {
		toSerialize["includeChain"] = o.IncludeChain
	}
	if o.ChainOrder.IsSet() {
		toSerialize["chainOrder"] = o.ChainOrder.Get()
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest struct {
	value *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) Get() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) Set(val *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest(val *KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}