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

// checks if the KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse{}

// KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse struct for KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
type KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse struct {
	ExpirationAlerts []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse `json:"expirationAlerts,omitempty"`
	AlertBuildResult *int32                                                                  `json:"alertBuildResult,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse {
	this := KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse {
	this := KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse{}
	return &this
}

// GetExpirationAlerts returns the ExpirationAlerts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) GetExpirationAlerts() []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse
		return ret
	}
	return o.ExpirationAlerts
}

// GetExpirationAlertsOk returns a tuple with the ExpirationAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) GetExpirationAlertsOk() ([]KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse, bool) {
	if o == nil || isNil(o.ExpirationAlerts) {
		return nil, false
	}
	return o.ExpirationAlerts, true
}

// HasExpirationAlerts returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) HasExpirationAlerts() bool {
	if o != nil && isNil(o.ExpirationAlerts) {
		return true
	}

	return false
}

// SetExpirationAlerts gets a reference to the given []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse and assigns it to the ExpirationAlerts field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) SetExpirationAlerts(v []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) {
	o.ExpirationAlerts = v
}

// GetAlertBuildResult returns the AlertBuildResult field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) GetAlertBuildResult() int32 {
	if o == nil || isNil(o.AlertBuildResult) {
		var ret int32
		return ret
	}
	return *o.AlertBuildResult
}

// GetAlertBuildResultOk returns a tuple with the AlertBuildResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) GetAlertBuildResultOk() (*int32, bool) {
	if o == nil || isNil(o.AlertBuildResult) {
		return nil, false
	}
	return o.AlertBuildResult, true
}

// HasAlertBuildResult returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) HasAlertBuildResult() bool {
	if o != nil && !isNil(o.AlertBuildResult) {
		return true
	}

	return false
}

// SetAlertBuildResult gets a reference to the given int32 and assigns it to the AlertBuildResult field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) SetAlertBuildResult(v int32) {
	o.AlertBuildResult = &v
}

func (o KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationAlerts != nil {
		toSerialize["expirationAlerts"] = o.ExpirationAlerts
	}
	if !isNil(o.AlertBuildResult) {
		toSerialize["alertBuildResult"] = o.AlertBuildResult
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse struct {
	value *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) Get() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) Set(val *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse(val *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) *NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}