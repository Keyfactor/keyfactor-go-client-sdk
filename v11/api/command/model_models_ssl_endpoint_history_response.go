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

// checks if the ModelsSSLEndpointHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSLEndpointHistoryResponse{}

// ModelsSSLEndpointHistoryResponse struct for ModelsSSLEndpointHistoryResponse
type ModelsSSLEndpointHistoryResponse struct {
	HistoryId           *string                                            `json:"historyId,omitempty"`
	EndpointId          *string                                            `json:"endpointId,omitempty"`
	AuditId             *int64                                             `json:"auditId,omitempty"`
	Timestamp           *time.Time                                         `json:"timestamp,omitempty"`
	Status              *int32                                             `json:"status,omitempty"`
	JobType             *int32                                             `json:"jobType,omitempty"`
	ProbeType           *int32                                             `json:"probeType,omitempty"`
	ReverseDNS          NullableString                                     `json:"reverseDNS,omitempty"`
	HistoryCertificates []ModelsSSLEndpointHistoryResponseCertificateModel `json:"historyCertificates,omitempty"`
}

// NewModelsSSLEndpointHistoryResponse instantiates a new ModelsSSLEndpointHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSLEndpointHistoryResponse() *ModelsSSLEndpointHistoryResponse {
	this := ModelsSSLEndpointHistoryResponse{}
	return &this
}

// NewModelsSSLEndpointHistoryResponseWithDefaults instantiates a new ModelsSSLEndpointHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSLEndpointHistoryResponseWithDefaults() *ModelsSSLEndpointHistoryResponse {
	this := ModelsSSLEndpointHistoryResponse{}
	return &this
}

// GetHistoryId returns the HistoryId field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetHistoryId() string {
	if o == nil || isNil(o.HistoryId) {
		var ret string
		return ret
	}
	return *o.HistoryId
}

// GetHistoryIdOk returns a tuple with the HistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetHistoryIdOk() (*string, bool) {
	if o == nil || isNil(o.HistoryId) {
		return nil, false
	}
	return o.HistoryId, true
}

// HasHistoryId returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasHistoryId() bool {
	if o != nil && !isNil(o.HistoryId) {
		return true
	}

	return false
}

// SetHistoryId gets a reference to the given string and assigns it to the HistoryId field.
func (o *ModelsSSLEndpointHistoryResponse) SetHistoryId(v string) {
	o.HistoryId = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetEndpointId() string {
	if o == nil || isNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetEndpointIdOk() (*string, bool) {
	if o == nil || isNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasEndpointId() bool {
	if o != nil && !isNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *ModelsSSLEndpointHistoryResponse) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetAuditId returns the AuditId field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetAuditId() int64 {
	if o == nil || isNil(o.AuditId) {
		var ret int64
		return ret
	}
	return *o.AuditId
}

// GetAuditIdOk returns a tuple with the AuditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetAuditIdOk() (*int64, bool) {
	if o == nil || isNil(o.AuditId) {
		return nil, false
	}
	return o.AuditId, true
}

// HasAuditId returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasAuditId() bool {
	if o != nil && !isNil(o.AuditId) {
		return true
	}

	return false
}

// SetAuditId gets a reference to the given int64 and assigns it to the AuditId field.
func (o *ModelsSSLEndpointHistoryResponse) SetAuditId(v int64) {
	o.AuditId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ModelsSSLEndpointHistoryResponse) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ModelsSSLEndpointHistoryResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetJobType() int32 {
	if o == nil || isNil(o.JobType) {
		var ret int32
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetJobTypeOk() (*int32, bool) {
	if o == nil || isNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasJobType() bool {
	if o != nil && !isNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given int32 and assigns it to the JobType field.
func (o *ModelsSSLEndpointHistoryResponse) SetJobType(v int32) {
	o.JobType = &v
}

// GetProbeType returns the ProbeType field value if set, zero value otherwise.
func (o *ModelsSSLEndpointHistoryResponse) GetProbeType() int32 {
	if o == nil || isNil(o.ProbeType) {
		var ret int32
		return ret
	}
	return *o.ProbeType
}

// GetProbeTypeOk returns a tuple with the ProbeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointHistoryResponse) GetProbeTypeOk() (*int32, bool) {
	if o == nil || isNil(o.ProbeType) {
		return nil, false
	}
	return o.ProbeType, true
}

// HasProbeType returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasProbeType() bool {
	if o != nil && !isNil(o.ProbeType) {
		return true
	}

	return false
}

// SetProbeType gets a reference to the given int32 and assigns it to the ProbeType field.
func (o *ModelsSSLEndpointHistoryResponse) SetProbeType(v int32) {
	o.ProbeType = &v
}

// GetReverseDNS returns the ReverseDNS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSLEndpointHistoryResponse) GetReverseDNS() string {
	if o == nil || isNil(o.ReverseDNS.Get()) {
		var ret string
		return ret
	}
	return *o.ReverseDNS.Get()
}

// GetReverseDNSOk returns a tuple with the ReverseDNS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSLEndpointHistoryResponse) GetReverseDNSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseDNS.Get(), o.ReverseDNS.IsSet()
}

// HasReverseDNS returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasReverseDNS() bool {
	if o != nil && o.ReverseDNS.IsSet() {
		return true
	}

	return false
}

// SetReverseDNS gets a reference to the given NullableString and assigns it to the ReverseDNS field.
func (o *ModelsSSLEndpointHistoryResponse) SetReverseDNS(v string) {
	o.ReverseDNS.Set(&v)
}

// SetReverseDNSNil sets the value for ReverseDNS to be an explicit nil
func (o *ModelsSSLEndpointHistoryResponse) SetReverseDNSNil() {
	o.ReverseDNS.Set(nil)
}

// UnsetReverseDNS ensures that no value is present for ReverseDNS, not even an explicit nil
func (o *ModelsSSLEndpointHistoryResponse) UnsetReverseDNS() {
	o.ReverseDNS.Unset()
}

// GetHistoryCertificates returns the HistoryCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSLEndpointHistoryResponse) GetHistoryCertificates() []ModelsSSLEndpointHistoryResponseCertificateModel {
	if o == nil {
		var ret []ModelsSSLEndpointHistoryResponseCertificateModel
		return ret
	}
	return o.HistoryCertificates
}

// GetHistoryCertificatesOk returns a tuple with the HistoryCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSLEndpointHistoryResponse) GetHistoryCertificatesOk() ([]ModelsSSLEndpointHistoryResponseCertificateModel, bool) {
	if o == nil || isNil(o.HistoryCertificates) {
		return nil, false
	}
	return o.HistoryCertificates, true
}

// HasHistoryCertificates returns a boolean if a field has been set.
func (o *ModelsSSLEndpointHistoryResponse) HasHistoryCertificates() bool {
	if o != nil && isNil(o.HistoryCertificates) {
		return true
	}

	return false
}

// SetHistoryCertificates gets a reference to the given []ModelsSSLEndpointHistoryResponseCertificateModel and assigns it to the HistoryCertificates field.
func (o *ModelsSSLEndpointHistoryResponse) SetHistoryCertificates(v []ModelsSSLEndpointHistoryResponseCertificateModel) {
	o.HistoryCertificates = v
}

func (o ModelsSSLEndpointHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSLEndpointHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HistoryId) {
		toSerialize["historyId"] = o.HistoryId
	}
	if !isNil(o.EndpointId) {
		toSerialize["endpointId"] = o.EndpointId
	}
	if !isNil(o.AuditId) {
		toSerialize["auditId"] = o.AuditId
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.JobType) {
		toSerialize["jobType"] = o.JobType
	}
	if !isNil(o.ProbeType) {
		toSerialize["probeType"] = o.ProbeType
	}
	if o.ReverseDNS.IsSet() {
		toSerialize["reverseDNS"] = o.ReverseDNS.Get()
	}
	if o.HistoryCertificates != nil {
		toSerialize["historyCertificates"] = o.HistoryCertificates
	}
	return toSerialize, nil
}

type NullableModelsSSLEndpointHistoryResponse struct {
	value *ModelsSSLEndpointHistoryResponse
	isSet bool
}

func (v NullableModelsSSLEndpointHistoryResponse) Get() *ModelsSSLEndpointHistoryResponse {
	return v.value
}

func (v *NullableModelsSSLEndpointHistoryResponse) Set(val *ModelsSSLEndpointHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSLEndpointHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSLEndpointHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSLEndpointHistoryResponse(val *ModelsSSLEndpointHistoryResponse) *NullableModelsSSLEndpointHistoryResponse {
	return &NullableModelsSSLEndpointHistoryResponse{value: val, isSet: true}
}

func (v NullableModelsSSLEndpointHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSLEndpointHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}