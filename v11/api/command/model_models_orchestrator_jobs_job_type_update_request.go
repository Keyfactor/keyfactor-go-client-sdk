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

// checks if the ModelsOrchestratorJobsJobTypeUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsOrchestratorJobsJobTypeUpdateRequest{}

// ModelsOrchestratorJobsJobTypeUpdateRequest struct for ModelsOrchestratorJobsJobTypeUpdateRequest
type ModelsOrchestratorJobsJobTypeUpdateRequest struct {
	Id            string                                      `json:"id"`
	JobTypeName   string                                      `json:"jobTypeName"`
	Description   NullableString                              `json:"description,omitempty"`
	JobTypeFields []ModelsOrchestratorJobsJobTypeFieldRequest `json:"jobTypeFields,omitempty"`
}

// NewModelsOrchestratorJobsJobTypeUpdateRequest instantiates a new ModelsOrchestratorJobsJobTypeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsOrchestratorJobsJobTypeUpdateRequest(id string, jobTypeName string) *ModelsOrchestratorJobsJobTypeUpdateRequest {
	this := ModelsOrchestratorJobsJobTypeUpdateRequest{}
	this.Id = id
	this.JobTypeName = jobTypeName
	return &this
}

// NewModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults instantiates a new ModelsOrchestratorJobsJobTypeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults() *ModelsOrchestratorJobsJobTypeUpdateRequest {
	this := ModelsOrchestratorJobsJobTypeUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetId(v string) {
	o.Id = v
}

// GetJobTypeName returns the JobTypeName field value
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobTypeName
}

// GetJobTypeNameOk returns a tuple with the JobTypeName field value
// and a boolean to check if the value has been set.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobTypeName, true
}

// SetJobTypeName sets field value
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeName(v string) {
	o.JobTypeName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetJobTypeFields returns the JobTypeFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFields() []ModelsOrchestratorJobsJobTypeFieldRequest {
	if o == nil {
		var ret []ModelsOrchestratorJobsJobTypeFieldRequest
		return ret
	}
	return o.JobTypeFields
}

// GetJobTypeFieldsOk returns a tuple with the JobTypeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFieldsOk() ([]ModelsOrchestratorJobsJobTypeFieldRequest, bool) {
	if o == nil || isNil(o.JobTypeFields) {
		return nil, false
	}
	return o.JobTypeFields, true
}

// HasJobTypeFields returns a boolean if a field has been set.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) HasJobTypeFields() bool {
	if o != nil && isNil(o.JobTypeFields) {
		return true
	}

	return false
}

// SetJobTypeFields gets a reference to the given []ModelsOrchestratorJobsJobTypeFieldRequest and assigns it to the JobTypeFields field.
func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeFields(v []ModelsOrchestratorJobsJobTypeFieldRequest) {
	o.JobTypeFields = v
}

func (o ModelsOrchestratorJobsJobTypeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsOrchestratorJobsJobTypeUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["jobTypeName"] = o.JobTypeName
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.JobTypeFields != nil {
		toSerialize["jobTypeFields"] = o.JobTypeFields
	}
	return toSerialize, nil
}

type NullableModelsOrchestratorJobsJobTypeUpdateRequest struct {
	value *ModelsOrchestratorJobsJobTypeUpdateRequest
	isSet bool
}

func (v NullableModelsOrchestratorJobsJobTypeUpdateRequest) Get() *ModelsOrchestratorJobsJobTypeUpdateRequest {
	return v.value
}

func (v *NullableModelsOrchestratorJobsJobTypeUpdateRequest) Set(val *ModelsOrchestratorJobsJobTypeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsOrchestratorJobsJobTypeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsOrchestratorJobsJobTypeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsOrchestratorJobsJobTypeUpdateRequest(val *ModelsOrchestratorJobsJobTypeUpdateRequest) *NullableModelsOrchestratorJobsJobTypeUpdateRequest {
	return &NullableModelsOrchestratorJobsJobTypeUpdateRequest{value: val, isSet: true}
}

func (v NullableModelsOrchestratorJobsJobTypeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsOrchestratorJobsJobTypeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}