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

// checks if the KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest{}

// KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest struct for KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest
type KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest struct {
	CertID              NullableInt32  `json:"certID,omitempty"`
	SerialNumber        NullableString `json:"serialNumber,omitempty"`
	IssuerDN            NullableString `json:"issuerDN,omitempty"`
	Thumbprint          NullableString `json:"thumbprint,omitempty"`
	IncludeChain        *bool          `json:"includeChain,omitempty"`
	ChainOrder          NullableString `json:"chainOrder,omitempty"`
	Password            string         `json:"password"`
	UseLegacyEncryption NullableBool   `json:"useLegacyEncryption,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest(password string) *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest {
	this := KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest{}
	this.Password = password
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest {
	this := KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest{}
	return &this
}

// GetCertID returns the CertID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetCertID() int32 {
	if o == nil || isNil(o.CertID.Get()) {
		var ret int32
		return ret
	}
	return *o.CertID.Get()
}

// GetCertIDOk returns a tuple with the CertID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetCertIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertID.Get(), o.CertID.IsSet()
}

// HasCertID returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasCertID() bool {
	if o != nil && o.CertID.IsSet() {
		return true
	}

	return false
}

// SetCertID gets a reference to the given NullableInt32 and assigns it to the CertID field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetCertID(v int32) {
	o.CertID.Set(&v)
}

// SetCertIDNil sets the value for CertID to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetCertIDNil() {
	o.CertID.Set(nil)
}

// UnsetCertID ensures that no value is present for CertID, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetCertID() {
	o.CertID.Unset()
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SerialNumber.Get()
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerialNumber.Get(), o.SerialNumber.IsSet()
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber.IsSet() {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given NullableString and assigns it to the SerialNumber field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetSerialNumber(v string) {
	o.SerialNumber.Set(&v)
}

// SetSerialNumberNil sets the value for SerialNumber to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetSerialNumberNil() {
	o.SerialNumber.Set(nil)
}

// UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetSerialNumber() {
	o.SerialNumber.Unset()
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}

// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbprint.Get()
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbprint.Get(), o.Thumbprint.IsSet()
}

// HasThumbprint returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasThumbprint() bool {
	if o != nil && o.Thumbprint.IsSet() {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given NullableString and assigns it to the Thumbprint field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetThumbprint(v string) {
	o.Thumbprint.Set(&v)
}

// SetThumbprintNil sets the value for Thumbprint to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetThumbprintNil() {
	o.Thumbprint.Set(nil)
}

// UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetThumbprint() {
	o.Thumbprint.Unset()
}

// GetIncludeChain returns the IncludeChain field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIncludeChain() bool {
	if o == nil || isNil(o.IncludeChain) {
		var ret bool
		return ret
	}
	return *o.IncludeChain
}

// GetIncludeChainOk returns a tuple with the IncludeChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIncludeChainOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeChain) {
		return nil, false
	}
	return o.IncludeChain, true
}

// HasIncludeChain returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasIncludeChain() bool {
	if o != nil && !isNil(o.IncludeChain) {
		return true
	}

	return false
}

// SetIncludeChain gets a reference to the given bool and assigns it to the IncludeChain field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIncludeChain(v bool) {
	o.IncludeChain = &v
}

// GetChainOrder returns the ChainOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetChainOrder() string {
	if o == nil || isNil(o.ChainOrder.Get()) {
		var ret string
		return ret
	}
	return *o.ChainOrder.Get()
}

// GetChainOrderOk returns a tuple with the ChainOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetChainOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainOrder.Get(), o.ChainOrder.IsSet()
}

// HasChainOrder returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasChainOrder() bool {
	if o != nil && o.ChainOrder.IsSet() {
		return true
	}

	return false
}

// SetChainOrder gets a reference to the given NullableString and assigns it to the ChainOrder field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetChainOrder(v string) {
	o.ChainOrder.Set(&v)
}

// SetChainOrderNil sets the value for ChainOrder to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetChainOrderNil() {
	o.ChainOrder.Set(nil)
}

// UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetChainOrder() {
	o.ChainOrder.Unset()
}

// GetPassword returns the Password field value
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetPassword(v string) {
	o.Password = v
}

// GetUseLegacyEncryption returns the UseLegacyEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetUseLegacyEncryption() bool {
	if o == nil || isNil(o.UseLegacyEncryption.Get()) {
		var ret bool
		return ret
	}
	return *o.UseLegacyEncryption.Get()
}

// GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetUseLegacyEncryptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseLegacyEncryption.Get(), o.UseLegacyEncryption.IsSet()
}

// HasUseLegacyEncryption returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasUseLegacyEncryption() bool {
	if o != nil && o.UseLegacyEncryption.IsSet() {
		return true
	}

	return false
}

// SetUseLegacyEncryption gets a reference to the given NullableBool and assigns it to the UseLegacyEncryption field.
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetUseLegacyEncryption(v bool) {
	o.UseLegacyEncryption.Set(&v)
}

// SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetUseLegacyEncryptionNil() {
	o.UseLegacyEncryption.Set(nil)
}

// UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetUseLegacyEncryption() {
	o.UseLegacyEncryption.Unset()
}

func (o KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["password"] = o.Password
	if o.UseLegacyEncryption.IsSet() {
		toSerialize["useLegacyEncryption"] = o.UseLegacyEncryption.Get()
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest struct {
	value *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) Get() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) Set(val *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest(val *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}