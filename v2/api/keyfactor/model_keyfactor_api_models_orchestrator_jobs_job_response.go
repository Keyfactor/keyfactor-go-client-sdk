/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor

import (
	"encoding/json"
	"time"
)

// checks if the KeyfactorApiModelsOrchestratorJobsJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsOrchestratorJobsJobResponse{}

// KeyfactorApiModelsOrchestratorJobsJobResponse struct for KeyfactorApiModelsOrchestratorJobsJobResponse
type KeyfactorApiModelsOrchestratorJobsJobResponse struct {
	JobId                *string                                              `json:"JobId,omitempty"`
	OrchestratorId       *string                                              `json:"OrchestratorId,omitempty"`
	JobTypeName          *string                                              `json:"JobTypeName,omitempty"`
	Schedule             *KeyfactorCommonSchedulingKeyfactorSchedule          `json:"Schedule,omitempty"`
	JobFields            []KeyfactorApiModelsOrchestratorJobsJobFieldResponse `json:"JobFields,omitempty"`
	RequestTimestamp     *time.Time                                           `json:"RequestTimestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsOrchestratorJobsJobResponse KeyfactorApiModelsOrchestratorJobsJobResponse

// NewKeyfactorApiModelsOrchestratorJobsJobResponse instantiates a new KeyfactorApiModelsOrchestratorJobsJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsOrchestratorJobsJobResponse() *KeyfactorApiModelsOrchestratorJobsJobResponse {
	this := KeyfactorApiModelsOrchestratorJobsJobResponse{}
	return &this
}

// NewKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsOrchestratorJobsJobResponseWithDefaults() *KeyfactorApiModelsOrchestratorJobsJobResponse {
	this := KeyfactorApiModelsOrchestratorJobsJobResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetOrchestratorId returns the OrchestratorId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorId() string {
	if o == nil || isNil(o.OrchestratorId) {
		var ret string
		return ret
	}
	return *o.OrchestratorId
}

// GetOrchestratorIdOk returns a tuple with the OrchestratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetOrchestratorIdOk() (*string, bool) {
	if o == nil || isNil(o.OrchestratorId) {
		return nil, false
	}
	return o.OrchestratorId, true
}

// HasOrchestratorId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasOrchestratorId() bool {
	if o != nil && !isNil(o.OrchestratorId) {
		return true
	}

	return false
}

// SetOrchestratorId gets a reference to the given string and assigns it to the OrchestratorId field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetOrchestratorId(v string) {
	o.OrchestratorId = &v
}

// GetJobTypeName returns the JobTypeName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeName() string {
	if o == nil || isNil(o.JobTypeName) {
		var ret string
		return ret
	}
	return *o.JobTypeName
}

// GetJobTypeNameOk returns a tuple with the JobTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobTypeNameOk() (*string, bool) {
	if o == nil || isNil(o.JobTypeName) {
		return nil, false
	}
	return o.JobTypeName, true
}

// HasJobTypeName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobTypeName() bool {
	if o != nil && !isNil(o.JobTypeName) {
		return true
	}

	return false
}

// SetJobTypeName gets a reference to the given string and assigns it to the JobTypeName field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobTypeName(v string) {
	o.JobTypeName = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.Schedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the Schedule field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.Schedule = &v
}

// GetJobFields returns the JobFields field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFields() []KeyfactorApiModelsOrchestratorJobsJobFieldResponse {
	if o == nil || isNil(o.JobFields) {
		var ret []KeyfactorApiModelsOrchestratorJobsJobFieldResponse
		return ret
	}
	return o.JobFields
}

// GetJobFieldsOk returns a tuple with the JobFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetJobFieldsOk() ([]KeyfactorApiModelsOrchestratorJobsJobFieldResponse, bool) {
	if o == nil || isNil(o.JobFields) {
		return nil, false
	}
	return o.JobFields, true
}

// HasJobFields returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasJobFields() bool {
	if o != nil && !isNil(o.JobFields) {
		return true
	}

	return false
}

// SetJobFields gets a reference to the given []KeyfactorApiModelsOrchestratorJobsJobFieldResponse and assigns it to the JobFields field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetJobFields(v []KeyfactorApiModelsOrchestratorJobsJobFieldResponse) {
	o.JobFields = v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestamp() time.Time {
	if o == nil || isNil(o.RequestTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) GetRequestTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestTimestamp) {
		return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) HasRequestTimestamp() bool {
	if o != nil && !isNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given time.Time and assigns it to the RequestTimestamp field.
func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) SetRequestTimestamp(v time.Time) {
	o.RequestTimestamp = &v
}

func (o KeyfactorApiModelsOrchestratorJobsJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsOrchestratorJobsJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobId) {
		toSerialize["JobId"] = o.JobId
	}
	if !isNil(o.OrchestratorId) {
		toSerialize["OrchestratorId"] = o.OrchestratorId
	}
	if !isNil(o.JobTypeName) {
		toSerialize["JobTypeName"] = o.JobTypeName
	}
	if !isNil(o.Schedule) {
		toSerialize["Schedule"] = o.Schedule
	}
	if !isNil(o.JobFields) {
		toSerialize["JobFields"] = o.JobFields
	}
	if !isNil(o.RequestTimestamp) {
		toSerialize["RequestTimestamp"] = o.RequestTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsOrchestratorJobsJobResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsOrchestratorJobsJobResponse := _KeyfactorApiModelsOrchestratorJobsJobResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsOrchestratorJobsJobResponse); err == nil {
		*o = KeyfactorApiModelsOrchestratorJobsJobResponse(varKeyfactorApiModelsOrchestratorJobsJobResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "JobId")
		delete(additionalProperties, "OrchestratorId")
		delete(additionalProperties, "JobTypeName")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "JobFields")
		delete(additionalProperties, "RequestTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsOrchestratorJobsJobResponse struct {
	value *KeyfactorApiModelsOrchestratorJobsJobResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobResponse) Get() *KeyfactorApiModelsOrchestratorJobsJobResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobResponse) Set(val *KeyfactorApiModelsOrchestratorJobsJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsOrchestratorJobsJobResponse(val *KeyfactorApiModelsOrchestratorJobsJobResponse) *NullableKeyfactorApiModelsOrchestratorJobsJobResponse {
	return &NullableKeyfactorApiModelsOrchestratorJobsJobResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}