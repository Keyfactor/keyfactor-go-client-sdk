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
	"time"
)

// checks if the KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}

// KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct for KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest
type KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct {
	AlertId        *int32     `json:"alertId,omitempty"`
	EvaluationDate *time.Time `json:"evaluationDate,omitempty"`
	SendAlerts     *bool      `json:"sendAlerts,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest() *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	this := KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	this := KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetAlertId() int32 {
	if o == nil || isNil(o.AlertId) {
		var ret int32
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetAlertIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasAlertId() bool {
	if o != nil && !isNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given int32 and assigns it to the AlertId field.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetAlertId(v int32) {
	o.AlertId = &v
}

// GetEvaluationDate returns the EvaluationDate field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetEvaluationDate() time.Time {
	if o == nil || isNil(o.EvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.EvaluationDate
}

// GetEvaluationDateOk returns a tuple with the EvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetEvaluationDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EvaluationDate) {
		return nil, false
	}
	return o.EvaluationDate, true
}

// HasEvaluationDate returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasEvaluationDate() bool {
	if o != nil && !isNil(o.EvaluationDate) {
		return true
	}

	return false
}

// SetEvaluationDate gets a reference to the given time.Time and assigns it to the EvaluationDate field.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetEvaluationDate(v time.Time) {
	o.EvaluationDate = &v
}

// GetSendAlerts returns the SendAlerts field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetSendAlerts() bool {
	if o == nil || isNil(o.SendAlerts) {
		var ret bool
		return ret
	}
	return *o.SendAlerts
}

// GetSendAlertsOk returns a tuple with the SendAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetSendAlertsOk() (*bool, bool) {
	if o == nil || isNil(o.SendAlerts) {
		return nil, false
	}
	return o.SendAlerts, true
}

// HasSendAlerts returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasSendAlerts() bool {
	if o != nil && !isNil(o.SendAlerts) {
		return true
	}

	return false
}

// SetSendAlerts gets a reference to the given bool and assigns it to the SendAlerts field.
func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetSendAlerts(v bool) {
	o.SendAlerts = &v
}

func (o KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !isNil(o.EvaluationDate) {
		toSerialize["evaluationDate"] = o.EvaluationDate
	}
	if !isNil(o.SendAlerts) {
		toSerialize["sendAlerts"] = o.SendAlerts
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct {
	value *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Get() *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Set(val *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest(val *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) *NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}