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

// checks if the KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse{}

// KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse struct for KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse
type KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse struct {
	KeyfactorVersion NullableString                                               `json:"keyfactorVersion,omitempty"`
	LicenseData      *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense `json:"licenseData,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse instantiates a new KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse() *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse {
	this := KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse {
	this := KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse{}
	return &this
}

// GetKeyfactorVersion returns the KeyfactorVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) GetKeyfactorVersion() string {
	if o == nil || isNil(o.KeyfactorVersion.Get()) {
		var ret string
		return ret
	}
	return *o.KeyfactorVersion.Get()
}

// GetKeyfactorVersionOk returns a tuple with the KeyfactorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) GetKeyfactorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyfactorVersion.Get(), o.KeyfactorVersion.IsSet()
}

// HasKeyfactorVersion returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) HasKeyfactorVersion() bool {
	if o != nil && o.KeyfactorVersion.IsSet() {
		return true
	}

	return false
}

// SetKeyfactorVersion gets a reference to the given NullableString and assigns it to the KeyfactorVersion field.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) SetKeyfactorVersion(v string) {
	o.KeyfactorVersion.Set(&v)
}

// SetKeyfactorVersionNil sets the value for KeyfactorVersion to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) SetKeyfactorVersionNil() {
	o.KeyfactorVersion.Set(nil)
}

// UnsetKeyfactorVersion ensures that no value is present for KeyfactorVersion, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) UnsetKeyfactorVersion() {
	o.KeyfactorVersion.Unset()
}

// GetLicenseData returns the LicenseData field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) GetLicenseData() KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense {
	if o == nil || isNil(o.LicenseData) {
		var ret KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense
		return ret
	}
	return *o.LicenseData
}

// GetLicenseDataOk returns a tuple with the LicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) GetLicenseDataOk() (*KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense, bool) {
	if o == nil || isNil(o.LicenseData) {
		return nil, false
	}
	return o.LicenseData, true
}

// HasLicenseData returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) HasLicenseData() bool {
	if o != nil && !isNil(o.LicenseData) {
		return true
	}

	return false
}

// SetLicenseData gets a reference to the given KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense and assigns it to the LicenseData field.
func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) SetLicenseData(v KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) {
	o.LicenseData = &v
}

func (o KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyfactorVersion.IsSet() {
		toSerialize["keyfactorVersion"] = o.KeyfactorVersion.Get()
	}
	if !isNil(o.LicenseData) {
		toSerialize["licenseData"] = o.LicenseData
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse struct {
	value *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) Get() *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) Set(val *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse(val *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) *NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsLicenseLicenseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}