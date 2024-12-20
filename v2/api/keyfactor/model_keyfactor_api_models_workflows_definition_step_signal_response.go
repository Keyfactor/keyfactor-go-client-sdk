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
)

// checks if the KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse{}

// KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse struct for KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse
type KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse struct {
	// The roles that are allowed to send this signal.
	RoleIds              []int32 `json:"RoleIds,omitempty"`
	SignalName           *string `json:"SignalName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse

// NewKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse() *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse {
	this := KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsDefinitionStepSignalResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsDefinitionStepSignalResponseWithDefaults() *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse {
	this := KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse{}
	return &this
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) GetRoleIds() []int32 {
	if o == nil || isNil(o.RoleIds) {
		var ret []int32
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) GetRoleIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []int32 and assigns it to the RoleIds field.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) SetRoleIds(v []int32) {
	o.RoleIds = v
}

// GetSignalName returns the SignalName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) GetSignalName() string {
	if o == nil || isNil(o.SignalName) {
		var ret string
		return ret
	}
	return *o.SignalName
}

// GetSignalNameOk returns a tuple with the SignalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) GetSignalNameOk() (*string, bool) {
	if o == nil || isNil(o.SignalName) {
		return nil, false
	}
	return o.SignalName, true
}

// HasSignalName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) HasSignalName() bool {
	if o != nil && !isNil(o.SignalName) {
		return true
	}

	return false
}

// SetSignalName gets a reference to the given string and assigns it to the SignalName field.
func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) SetSignalName(v string) {
	o.SignalName = &v
}

func (o KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoleIds) {
		toSerialize["RoleIds"] = o.RoleIds
	}
	if !isNil(o.SignalName) {
		toSerialize["SignalName"] = o.SignalName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse := _KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse); err == nil {
		*o = KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse(varKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "RoleIds")
		delete(additionalProperties, "SignalName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse struct {
	value *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) Get() *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) Set(val *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse(val *KeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) *NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse {
	return &NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsDefinitionStepSignalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
